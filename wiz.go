package wiz

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"github.com/jinzhu/gorm"
)

const (
	DefaultDbPath   = "data/index.db"
	DefaultDataPath = "data/notes"

	ContentFileName = "index.html"
	ContentFilesDir = "index_files"
)

func New(root string) (*Wiz, error) {
	db, err := gorm.Open("sqlite3", filepath.Join(root, DefaultDbPath))
	if err != nil {
		return nil, err
	}

	wiz := Wiz{
		root:  root,
		db:    db,
		notes: filepath.Join(root, DefaultDataPath),
	}
	return &wiz, nil
}

type Wiz struct {
	root  string
	db    *gorm.DB
	notes string
}

type WalkFunc func(path string, doc Document, tags []string) error

func (w *Wiz) Walk(f WalkFunc) error {
	var docs []*Document
	if err := w.db.Model(&Document{}).Find(&docs).Error; err != nil {
		return err
	}
	for _, doc := range docs {
		path := filepath.Join(w.notes, fmt.Sprintf("{%s}", doc.Guid))
		tags := w.Tags(doc.Guid)
		if err := f(path, *doc, tags); err != nil {
			return err
		}
	}
	return nil
}

func (w *Wiz) Tags(guid string) []string {
	var relations []*DocumentTag
	_ = w.db.Model(&DocumentTag{}).Where(DocumentTag{DocumentGuid: guid}).Find(&relations)
	tags := make([]string, 0, len(relations))
	if len(relations) > 0 {
		tag := &Tag{}
		for _, dt := range relations {
			if err := w.db.Where(Tag{Guid: dt.TagGuid}).Find(tag).Error; err != nil {
				tags = append(tags, tag.Name)
			}
		}
	}
	return tags
}

func (w *Wiz) Path(guid string) string {
	return filepath.Join(w.notes, fmt.Sprintf("{%s}", guid))
}

func (w *Wiz) Content(guid string) (content []byte, err error) {
	path := w.Path(guid)
	reader, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if (*file).Name == ContentFileName {
			fd, _ := file.Open()
			buff := make([]byte, file.UncompressedSize64)
			_, err = fd.Read(buff)
			fd.Close()
			return buff, err
		}
	}
	return nil, errors.New("not found")
}

func (w *Wiz) Files(guid string, f func(string, io.Reader) error) (err error) {
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
