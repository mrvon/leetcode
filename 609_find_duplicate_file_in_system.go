package main

import (
	"fmt"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	contents := make(map[string][]string)
	for i := 0; i < len(paths); i++ {
		path := paths[i]
		items := strings.Split(path, " ")
		if len(items) <= 1 {
			continue
		}
		root := items[0]
		for i := 1; i < len(items); i++ {
			pair := strings.Split(items[i], "(")
			name := pair[0]
			content := pair[1]
			content = content[:len(content)-1] // remove ')'
			contents[content] = append(contents[content], root+"/"+name)
		}
	}
	results := [][]string{}
	for _, paths := range contents {
		if len(paths) >= 2 {
			results = append(results, paths)
		}
	}
	return results
}

func main() {
	paths := []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}
	fmt.Println(findDuplicate(paths))
	paths = []string{"root/a 1.txt(abcd) 2.txt(efsfgh)", "root/c 3.txt(abdfcd)", "root/c/d 4.txt(efggdfh)"}
	fmt.Println(findDuplicate(paths))
}
