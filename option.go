package wiz

import "path/filepath"

type option struct {
	DbPath    string
	NotesPath string
}

type Option func(option *option)

func DbPath(path string) Option {
	return func(option *option) {
		option.DbPath = path
	}
}
func NotesPath(path string) Option {
	return func(option *option) {
		option.NotesPath = path
	}
}

func RootPath(path string) Option {
	return func(option *option) {
		option.DbPath = filepath.Join(path, DefaultDbPath)
		option.NotesPath = filepath.Join(path, DefaultNotesPath)
	}
}
