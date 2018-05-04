package main

import "fmt"

func padding(n int) []byte {
	s := []byte{}
	for i := 0; i < n; i++ {
		s = append(s, ' ')
	}
	return s
}

func justify(words []string, maxWidth int) string {
	word_count := len(words)

	character_count := 0
	for i := 0; i < len(words); i++ {
		character_count += len(words[i])
	}

	space_count := maxWidth - character_count

	new_line := []byte{}
	if word_count == 1 {
		new_line = append(new_line, words[0]...)
		new_line = append(new_line, padding(space_count)...)
	} else {
		slot_count := word_count - 1
		slot_size := space_count / slot_count
		slot_remainder := space_count % slot_count
		slot := []int{}
		for i := 1; i <= slot_count; i++ {
			if i <= slot_remainder {
				slot = append(slot, slot_size+1)
			} else {
				slot = append(slot, slot_size)
			}
		}
		for i := 0; i < len(words); i++ {
			new_line = append(new_line, words[i]...)
			if i < len(words)-1 {
				new_line = append(new_line, padding(slot[i])...)
			}
		}
	}
	return string(new_line)
}

func concat(words []string, maxWidth int) string {
	new_line := []byte{}
	for i := 0; i < len(words); i++ {
		new_line = append(new_line, words[i]...)
		if i < len(words)-1 {
			new_line = append(new_line, ' ')
		}
	}
	for len(new_line) < maxWidth {
		new_line = append(new_line, ' ')
	}
	return string(new_line)
}

func fullJustify(words []string, maxWidth int) []string {
	lines := []string{}
	line_words := []string{}
	length := 0

	i := 0
	for i < len(words) {
		word := words[i]
		space := 0
		if len(line_words) > 0 {
			space = 1
		}

		if length+space+len(word) > maxWidth {
			// full
			lines = append(lines, justify(line_words, maxWidth))
			// clear
			line_words = []string{}
			length = 0
		} else {
			line_words = append(line_words, word)
			length += len(word)
			length += space
			// next word
			i++
		}
	}

	// last line_words needn't justify
	lines = append(lines, concat(line_words, maxWidth))

	return lines
}

func main() {
	lines := fullJustify(
		[]string{"This", "is", "an", "example", "of", "text", "jus", "Goo"},
		16,
	)
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}
