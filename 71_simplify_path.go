package main

import "fmt"

func next(path string, cur int) (string, int) {
	i := cur + 1
	for i < len(path) && path[i] == '/' {
		i++
	}
	j := i + 1
	for j < len(path) && path[j] != '/' {
		j++
	}
	if j > len(path) {
		return "", j
	} else {
		return path[i:j], j
	}
}

func split(path string) []string {
	var components []string

	str := ""
	cur := -1
	for {
		str, cur = next(path, cur)
		if len(str) == 0 {
			break
		}
		components = append(components, str)
	}

	return components
}

func simplifyPath(path string) string {
	components := split(path)
	simplify := []string{}

	for i := 0; i < len(components); i++ {
		if components[i] == "." {
			continue
		} else if components[i] == ".." {
			if len(simplify) > 0 {
				simplify = simplify[0 : len(simplify)-1]
			}
		} else {
			simplify = append(simplify, components[i])
		}
	}

	new_path := []byte("/")
	for i := 0; i < len(simplify); i++ {
		new_path = append(new_path, []byte(simplify[i])...)
		if i < len(simplify)-1 {
			new_path = append(new_path, []byte("/")...)
		}
	}

	return string(new_path)
}

func main() {
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/hell/.././hello////goo"))
}
