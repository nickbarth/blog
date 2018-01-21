package main

import (
	_ "fmt"
	"path/filepath"
	"sort"
)

func loadEntry(file string) {
	NewEntry(file)
}

type Blog struct {
	entries []Entry
}

func NewBlog() *Blog {
	blog := Blog{}

	files := blog.getEntryFiles()

	for _, file := range files {
		loadEntry(file)
	}

	return &blog
}

func (b Blog) getEntryFiles() []string {
	files, _ := filepath.Glob("entries/*.md")
	sort.Strings(files)

	return files
}
