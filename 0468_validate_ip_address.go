package main

import (
	"fmt"
	"strconv"
	"strings"
)

func check4(c string) bool {
	// leading zero or sign
	for i := 0; i < len(c)-1; i++ {
		if c[i] == '+' || c[i] == '-' {
			return false
		}

		if c[i] != '0' {
			break
		} else {
			return false
		}
	}

	i, err := strconv.Atoi(c)
	if err != nil {
		return false
	}

	return i >= 0 && i <= 255
}

func check6(c string) bool {
	// extra leading zero
	if len(c) > 4 {
		return false
	}

	// leading sign
	for i := 0; i < len(c)-1; i++ {
		if c[i] == '+' || c[i] == '-' {
			return false
		}
	}

	i, err := strconv.ParseInt(c, 16, 0)
	if err != nil {
		return false
	}

	return i >= 0 && i <= 0xffff
}

func validIPAddress(IP string) string {
	if ipv4 := strings.Split(IP, "."); len(ipv4) == 4 {
		for i := 0; i < len(ipv4); i++ {
			if !check4(ipv4[i]) {
				return "Neither"
			}
		}
		return "IPv4"
	} else {
		ipv6 := strings.Split(IP, ":")
		if len(ipv6) == 8 {
			for i := 0; i < len(ipv6); i++ {
				if !check6(ipv6[i]) {
					return "Neither"
				}
			}
			return "IPv6"
		} else {
			return "Neither"
		}
	}
}

func main() {
	fmt.Println(validIPAddress("172.16.254.1"))
	fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("ffff:0db8:85a3:0:0:8A2E:0370:7334"))

	fmt.Println("-----------------")

	fmt.Println(validIPAddress("-0.1.1.1"))
	fmt.Println(validIPAddress("+1.1.1.1"))
	fmt.Println(validIPAddress("-1.1.1.1"))
	fmt.Println(validIPAddress("01.01.01.01"))
	fmt.Println(validIPAddress("256.256.256.256"))
	fmt.Println(validIPAddress("2001:0db8:85a3:00000:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("002001:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("-001:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("-001:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("1ffff:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress(":0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress(""))
}
