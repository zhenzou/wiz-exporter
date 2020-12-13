package wiz

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/mattn/godown"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DefaultDbPath    = "/data/index.db"
	DefaultNotesPath = "/data/notes"

	ContentFileName = "index.html"
	ContentFilesDir = "index_files"
)

const (
	tsFormat = "2006-01-02 15:04:05"
)

func New(opts ...Option) (*Wiz, error) {
	opt := &options{}

	for _, o := range opts {
		o(opt)
	}

	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("file://%s", opt.DbPath)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	wiz := Wiz{
		db:    db,
		notes: opt.NotesPath,
	}
	return &wiz, nil
}

type Wiz struct {
	db    *gorm.DB
	notes string
}

type WalkFunc func(doc Document) error

type FilesFunc func(path string, reader io.Reader) error

func (w *Wiz) Walk(f WalkFunc) error {
	var docs []*documentEntity
	if err := w.db.Model(&documentEntity{}).Find(&docs).Error; err != nil {
		return err
	}
	for _, doc := range docs {
		if err := f(Document{
			documentEntity: *doc,
			wiz:            w,
		}); err != nil {
			return err
		}
	}
	return nil
}

type Document struct {
	documentEntity
	wiz *Wiz
}

func (d Document) CreatedAt() time.Time {
	ts, err := time.Parse(tsFormat, d.DTCreated)
	if err != nil {
		panic(err)
	}
	return ts
}

func (d Document) UpdatedAt() time.Time {
	ts, err := time.Parse(tsFormat, d.DTModified)
	if err != nil {
		panic(err)
	}
	return ts
}

func (d Document) AccessedAt() time.Time {
	ts, err := time.Parse(tsFormat, d.DTAccessed)
	if err != nil {
		panic(err)
	}
	return ts
}

func (d *Document) Tags() []string {
	var relations []*documentTagEntity
	_ = d.wiz.db.Model(&documentTagEntity{}).Where(documentTagEntity{DocumentGuid: d.Guid}).Find(&relations)
	tags := make([]string, 0, len(relations))
	if len(relations) > 0 {
		tag := &tagEntity{}
		for _, dt := range relations {
			if err := d.wiz.db.Where(tagEntity{Guid: dt.TagGuid}).Find(tag).Error; err != nil {
				tags = append(tags, tag.Name)
			}
		}
	}
	return tags
}

func (d *Document) Path() string {
	return filepath.Join(d.wiz.notes, fmt.Sprintf("{%s}", d.Guid))
}

func (d *Document) Markdown() (string, error) {
	content, err := d.Raw()
	if err != nil {
		return "", err
	}

	if strings.HasSuffix(d.Title, ".md") {
		return content, nil
	}
	buf := bytes.Buffer{}
	err = godown.Convert(&buf, strings.NewReader(content), nil)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (d *Document) Raw() (content string, err error) {
	path := d.Path()
	reader, err := zip.OpenReader(path)
	if err != nil {
		return "", err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if file.Name == ContentFileName {
			fd, _ := file.Open()

			buff, err := ioutil.ReadAll(fd)
			if err != nil {
				return "", err
			}
			fd.Close()
			return string(buff), err
		}
	}
	return "", errors.New("not found")
}

func (d *Document) Files(f FilesFunc) (err error) {
	path := d.Path()
	reader, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if file.Name != ContentFileName {
			fd, _ := file.Open()
			//noinspection GoDeferInLoop
			defer fd.Close()
			if err = f(file.Name, fd); err != nil {
				return err
			}
		}
	}
	return err
}
