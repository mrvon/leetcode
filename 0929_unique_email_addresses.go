package main

import (
	"fmt"
	"strings"
)

func split(s string, split string) (string, string) {
	parts := strings.Split(s, split)
	if len(parts) >= 2 {
		return parts[0], parts[1]
	}
	return parts[0], ""
}

func translate(origin string) string {
	localName, domainName := split(origin, "@")
	header, _ := split(localName, "+")
	header = strings.Replace(header, ".", "", -1)
	return header + "@" + domainName
}

func numUniqueEmails(emails []string) int {
	uniqueEmails := make(map[string]bool)
	uniqueCount := 0
	for _, email := range emails {
		e := translate(email)
		if !uniqueEmails[e] {
			uniqueCount++
			uniqueEmails[e] = true
		}
	}
	return uniqueCount
}

func main() {
	mails := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}
	fmt.Println(numUniqueEmails(mails))
}
