package main

import (
	"fmt"
	"sort"
	"strings"
)

type Dir struct {
	items map[string]interface{}
}

type File struct {
	contents []byte
}

type FileSystem struct {
	root *Dir
}

func parsePath(path string) []string {
	l := strings.Split(path, "/")
	n := []string{}
	for i := 0; i < len(l); i++ {
		if l[i] == "" {
			continue
		}
		n = append(n, l[i])
	}
	return n
}

func Constructor() FileSystem {
	return FileSystem{&Dir{
		items: make(map[string]interface{}),
	}}
}

func (this *FileSystem) Ls(path string) []string {
	list := parsePath(path)
	cwd := this.root
	for i := 0; i < len(list)-1; i++ {
		item := cwd.items[list[i]]
		if item == nil {
			return nil
		}
		cwd = item.(*Dir)
	}

	var item interface{}
	if len(list) > 0 {
		item = cwd.items[list[len(list)-1]]
	} else {
		item = cwd
	}
	results := []string{}

	switch i := item.(type) {
	case *Dir:
		{
			for name := range i.items {
				results = append(results, name)
			}
		}
	case *File:
		results = append(results, list[len(list)-1])
	}

	sort.Strings(results)
	return results
}

func (this *FileSystem) Mkdir(path string) {
	this.AuxMkdir(this.root, parsePath(path))
}

func (this *FileSystem) AuxMkdir(cwd *Dir, path []string) {
	if len(path) == 0 {
		return
	}
	name := path[0]
	if cwd.items[name] == nil {
		cwd.items[name] = &Dir{
			items: make(map[string]interface{}),
		}
	}
	this.AuxMkdir(cwd.items[name].(*Dir), path[1:])
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	file := this.getFile(filePath)
	file.contents = append(file.contents, content...)
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	file := this.getFile(filePath)
	return string(file.contents)
}

func (this *FileSystem) getFile(path string) *File {
	list := parsePath(path)
	cwd := this.root
	for i := 0; i < len(list)-1; i++ {
		item := cwd.items[list[i]]
		if item == nil {
			return nil
		}
		cwd = item.(*Dir)
	}
	file := cwd.items[list[len(list)-1]]
	if file == nil {
		file = &File{}
		cwd.items[list[len(list)-1]] = file
	}
	return file.(*File)
}

func main() {
	obj := Constructor()
	obj.Mkdir("/goowmfn")
	fmt.Println(obj.Ls("/goowmfn"))
	fmt.Println(obj.Ls("/"))
	obj.Mkdir("/z")
	fmt.Println(obj.Ls("/z"))
	fmt.Println(obj.Ls("/"))
	obj.AddContentToFile("/goowmfn/c", "shetopcy")
	fmt.Println(obj.Ls("/z"))
	fmt.Println(obj.Ls("/goowmfn/c"))
	fmt.Println(obj.Ls("/"))

	obj = Constructor()
	fmt.Println(obj.Ls("/"))
	obj.Mkdir("/a/b/c")
	obj.AddContentToFile("/a/b/c/d", "hello")
	fmt.Println(obj.Ls("/"))
	fmt.Println(obj.ReadContentFromFile("/a/b/c/d"))
}
