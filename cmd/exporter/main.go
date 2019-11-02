package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/lunny/html2md"
	"github.com/zhenzou/wiz"
)

var (
	in  string
	out string
	w   *wiz.Wiz
)

func init() {
	flag.StringVar(&in, "in", "", "wiz data directory")
	flag.StringVar(&out, "out", "", "directory save exported note file  ")
}

func mkdirs(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func markdown(doc wiz.Document) (string, string, error) {
	content, err := w.Content(doc.Guid)
	if err != nil {
		return "", content, err
	}

	if strings.HasSuffix(doc.Title, ".md") {
		return doc.Title, content, nil
	}
	return doc.Title + ".md", html2md.Convert(content), nil
}

func writeFile(path string, content [] byte) error {
	base := filepath.Dir(path)
	err := mkdirs(base)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return ioutil.WriteFile(path, content, os.ModePerm)
}

func main() {

	flag.Parse()
	if in == "" || out == "" {
		println("in and out directory must not be empty!")
		os.Exit(-1)
	}
	var err error

	w, err = wiz.New(wiz.RootPath(in))
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}
	err = w.Walk(func(doc wiz.Document) error {
		println(fmt.Sprintf("exporting %s%s", doc.Location, doc.Title))

		name, md, err := markdown(doc)
		if err != nil {
			return err
		}
		dir := filepath.Join(out, doc.Location)
		path := filepath.Join(dir, name)

		err = writeFile(path, []byte(md))
		if err != nil {
			return err
		}

		return w.Files(doc.Guid, func(path string, reader io.Reader) error {
			path = filepath.Join(dir, path)
			content, err := ioutil.ReadAll(reader)
			if err != nil {
				return err
			}
			return writeFile(path, content)
		})
	})

	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}
}
