package wiz

import "path/filepath"

type options struct {
	DbPath    string
	NotesPath string
}

type Option func(option *options)

func DbPath(path string) Option {
	return func(option *options) {
		option.DbPath = path
	}
}
func NotesPath(path string) Option {
	return func(option *options) {
		option.NotesPath = path
	}
}

func RootPath(path string) Option {
	return func(option *options) {
		option.DbPath = filepath.Join(path, DefaultDbPath)
		option.NotesPath = filepath.Join(path, DefaultNotesPath)
	}
}
