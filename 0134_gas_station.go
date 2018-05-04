package main

import "fmt"

type Pair struct {
	index int
	gas   int
}

func sign(gas int) int {
	if gas >= 0 {
		return 1
	} else {
		return -1
	}
}

func collapse(gas []int) []Pair {
	list := []Pair{}

	for i := 0; i < len(gas); i++ {
		if len(list) == 0 {
			p := Pair{
				index: i,
				gas:   gas[i],
			}
			list = append(list, p)
		} else {
			p := &list[len(list)-1]
			if sign(p.gas) == sign(gas[i]) {
				p.gas += gas[i]
			} else {
				p := Pair{
					index: i,
					gas:   gas[i],
				}
				list = append(list, p)
			}
		}
	}

	return list
}

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		gas[i] -= cost[i]
	}

	list := collapse(gas)

	for i := 0; i < len(list); i++ {
		is_ok := true
		sum := list[i].gas

		if sum < 0 {
			continue
		}

		for j := i + 1; j < len(list); j++ {
			sum += list[j].gas
			if sum < 0 {
				is_ok = false
				break
			}
		}

		if !is_ok {
			continue
		}

		for j := 0; j < i; j++ {
			sum += list[j].gas
			if sum < 0 {
				is_ok = false
				break
			}
		}

		if !is_ok {
			continue
		}

		return list[i].index
	}

	return -1
}

func main() {
	fmt.Println(canCompleteCircuit(
		[]int{5},
		[]int{4},
	))

	fmt.Println(canCompleteCircuit(
		[]int{4},
		[]int{5},
	))

	fmt.Println(canCompleteCircuit(
		[]int{4, 1, 2, 3, 1, 1},
		[]int{5, 2, 2, 1, 2, 2},
	))

	fmt.Println(canCompleteCircuit(
		[]int{1, 2, 3, 4, 5},
		[]int{3, 4, 5, 1, 2},
	))
}
