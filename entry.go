package main

import (
	"fmt"
	"regexp"
	"time"
)

type Entry struct {
	title     string
	summary   string
	category  string
	date      time.Time
	permalink string
}

func NewEntry(file string) *Entry {
	r, _ := regexp.Compile("([0-9]{4})-([0-9]{2})-([0-9]{2})\\.md")
	m := r.FindAllStringSubmatch(file, -1)[0]
	year, month, day := m[0], m[1], m[2]
	fmt.Println(year, month, day)
	return &Entry{}
}
