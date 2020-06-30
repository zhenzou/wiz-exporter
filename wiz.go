package wiz

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DefaultDbPath    = "/data/index.db"
	DefaultNotesPath = "/data/notes"

	ContentFileName = "index.html"
	ContentFilesDir = "index_files"
)

var (
	baseDocumentSQL    string = extractBaseQuerySQL(Document{}, DocumentTableName)
	baseTagSQL         string = extractBaseQuerySQL(Tag{}, TagTableName)
	baseDocumentTagSQL string = extractBaseQuerySQL(DocumentTag{}, DocumentTagTableName)
)

func New(opts ...Option) (*Wiz, error) {
	opt := &options{}

	for _, o := range opts {
		o(opt)
	}

	db, err := sqlx.Open("sqlite3", opt.DbPath)
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
	db    *sqlx.DB
	notes string
}

type WalkFunc func(doc Document) error

type FilesFunc func(path string, reader io.Reader) error

func (w *Wiz) Walk(f WalkFunc) error {
	var docs []*Document
	err := w.db.Select(&docs, baseDocumentSQL)
	if err != nil {
		return err
	}
	for _, doc := range docs {
		if err := f(*doc); err != nil {
			return err
		}
	}
	return nil
}

func (w *Wiz) Tags(guid string) []string {
	var relations []*DocumentTag
	err := w.db.Select(&relations, baseDocumentSQL+"WHERE DOCUMENT_GUID = ?", guid)
	if err != nil {
		return []string{}
	}
	tags := make([]string, 0, len(relations))
	if len(relations) > 0 {
		tag := &Tag{}
		for _, dt := range relations {
			if err := w.db.Select(tag, baseTagSQL+"WHERE TAG_GUID = ?", dt.TagGuid); err == nil {

				tags = append(tags, tag.Name)
			}
		}
	}
	return tags
}

func (w *Wiz) Path(guid string) string {
	return filepath.Join(w.notes, fmt.Sprintf("{%s}", guid))
}

func (w *Wiz) Content(guid string) (content string, err error) {
	path := w.Path(guid)
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

func (w *Wiz) Files(guid string, f FilesFunc) (err error) {
	path := w.Path(guid)
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
