// Top down recursive
package main

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

func isDigit(b byte) bool {
	return b == '-' || b >= '0' && b <= '9'
}

func parseList(s string, index *int) *NestedInteger {
	n := &NestedInteger{}

	(*index)++ // skip '['

	for s[*index] != ']' {
		if s[*index] == '[' {
			l := parseList(s, index)
			n.Add(*l)
		} else if isDigit(s[*index]) {
			l := parseNumber(s, index)
			n.Add(*l)
		} else {
			// ','
			(*index)++
		}
	}

	(*index)++ // skip ']'

	return n
}

func parseNumber(s string, index *int) *NestedInteger {
	n := &NestedInteger{}
	start := *index
	for *index < len(s) && isDigit(s[*index]) {
		(*index)++
	}
	number, _ := strconv.Atoi(s[start:*index])
	n.SetInteger(number)
	return n
}

func deserialize(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	}
	index := 0
	if s[index] == '[' {
		return parseList(s, &index)
	} else if isDigit(s[index]) {
		return parseNumber(s, &index)
	} else {
		return nil
	}
}
