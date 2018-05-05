/*
A character in UTF8 can be from 1 to 4 bytes long, subjected to the following
rules:

1. For 1-byte character, the first bit is a 0, followed by its unicode code.

2. For n-bytes character, the first n-bits are all one's, the n+1 bit is 0,
followed by n-1 bytes with most significant 2 bits being 10.

This is how the UTF-8 encoding would work:

   Char. number range  |        UTF-8 octet sequence
      (hexadecimal)    |              (binary)
   --------------------+---------------------------------------------
   0000 0000-0000 007F | 0xxxxxxx
   0000 0080-0000 07FF | 110xxxxx 10xxxxxx
   0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
   0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx

Given an array of integers representing the data, return whether it is a valid
utf-8 encoding.
*/
package main

import "fmt"

func oneB(i int) bool {
	b := byte(i)
	return (b & (0x1 << 7)) == 0
}

func twoB(i int) bool {
	b := byte(i)
	return ((b&(0x1<<7)) > 0 &&
		(b&(0x1<<6)) > 0 &&
		(b&(0x1<<5)) == 0)
}

func threeB(i int) bool {
	b := byte(i)
	return ((b&(0x1<<7)) > 0 &&
		(b&(0x1<<6)) > 0 &&
		(b&(0x1<<5)) > 0 &&
		(b&(0x1<<4)) == 0)
}

func fourB(i int) bool {
	b := byte(i)
	return ((b&(0x1<<7)) > 0 &&
		(b&(0x1<<6)) > 0 &&
		(b&(0x1<<5)) > 0 &&
		(b&(0x1<<4)) > 0 &&
		(b&(0x1<<3)) == 0)
}

func conB(i int) bool {
	b := byte(i)
	return ((b&(0x1<<7)) > 0 &&
		(b&(0x1<<6)) == 0)
}

func parseItem(data []int, index *int) bool {
	if oneB(data[*index]) {
		(*index)++
		return true
	} else {
		count := 0
		if twoB(data[*index]) {
			count = 1
		} else if threeB(data[*index]) {
			count = 2
		} else if fourB(data[*index]) {
			count = 3
		} else {
			return false
		}

		(*index)++

		n := len(data)

		for i := 0; i < count; i++ {
			if *index >= n {
				return false
			}
			if conB(data[*index]) {
				(*index)++
			} else {
				return false
			}
		}
		return true
	}
}

func validUtf8(data []int) bool {
	index := 0
	for index < len(data) {
		if !parseItem(data, &index) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(validUtf8([]int{1, 1, 1}))
	fmt.Println(validUtf8([]int{197, 130, 1}))
	fmt.Println(validUtf8([]int{197, 130, 197, 130}))
	fmt.Println(validUtf8([]int{197, 130, 197}))
	fmt.Println(validUtf8([]int{235, 140, 4}))
	fmt.Println(validUtf8([]int{235, 140}))
}
