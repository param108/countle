package main

import (
	"fmt"
)

type Operation struct {
	X int
	Y int
	Op string
	Val int
}

type Pair struct {
	X int
	Y int
	R []int
}

func (p *Pair) String() string {
	return fmt.Sprintf("{%d,%d,%v}", p.X, p.Y, p.R)
}

func getRemaining(x, y int, r []int) []int {
	foundx := false
	foundy := false

	ret := []int{}

	for _, v := range (r) {
		if x == v && !foundx {
			foundx = true
			continue
		}
		if y == v && !foundy {
			foundy = true
			continue
		}
		ret = append(ret, v)
	}

	return ret
}

// returns []{X,Y,Remaining}
func pairs(numbers []int) []*Pair {
	ret := []*Pair{}
	for ix, x := range(numbers) {
		for iy, y := range(numbers) {
			if ix == iy {
				continue
			}
			ret = append(ret, &Pair{x,y, getRemaining(x,y,numbers)})
		}
	}

	return ret
}

var done = false

func calc(numbers []int, result int, l[]*Operation, depth int) {
	if done {
		return
	}

	if depth == 10 {
		return
	}
	ops := []string{"+","-","*","/"}

	allPairs := pairs(numbers)

	for _, pair := range(allPairs) {
		for _, op := range(ops) {
			if done {
				return
			}
			newL := l
			var remaining []int
			switch(op) {
			case "+":
				val:= pair.X + pair.Y
				remaining = pair.R
				remaining = append(remaining, val)
				newL = append(newL, &Operation{ pair.X, pair.Y, "+", val})
			case "-":
				if pair.X < pair.Y {
					continue
				}
				val:= pair.X - pair.Y
				remaining = pair.R
				remaining = append(remaining, val)
				newL = append(newL, &Operation{ pair.X, pair.Y, "-", val})
			case "*":
				val:= pair.X * pair.Y
				remaining = pair.R
				remaining = append(remaining, val)
				newL = append(newL, &Operation{ pair.X, pair.Y, "*", val})
			case "/":
				if pair.Y == 0 {
					continue
				}
				if pair.X % pair.Y != 0 {
					continue
				}
				val:= pair.X / pair.Y
				remaining = pair.R
				remaining = append(remaining, val)
				newL = append(newL, &Operation{ pair.X, pair.Y, "/", val})
			}
			if len(remaining) == 1 {
				if remaining[0] == result {
					for _, o := range(newL) {
						fmt.Println(o.X,o.Op,o.Y,"=",o.Val)
					}
					done = true
					return
				}
				continue
			}
			calc(remaining, result, newL, depth+1)

		}
	}
}


func main() {
	calc([]int{6,6,8,2,1,3}, 526,[]*Operation{},0)
}
