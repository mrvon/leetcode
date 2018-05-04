package main

import (
	"fmt"
	"strconv"
)

type Domain struct {
	count int
	name  string
}

func parse(s string) []Domain {
	count := 0
	url := ""
	fmt.Sscanf(s, "%d %s", &count, &url)
	domains := []Domain{}
	domains = append(domains, Domain{count: count, name: url})
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '.' {
			domains = append(domains, Domain{
				count: count,
				name:  url[i+1:],
			})
		}
	}
	return domains
}

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, domain := range cpdomains {
		for _, d := range parse(domain) {
			m[d.name] += d.count
		}
	}

	l := []string{}
	for name, count := range m {
		l = append(l, strconv.Itoa(count)+" "+name)
	}
	return l
}

func main() {
	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
