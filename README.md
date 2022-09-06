# Countle

This is a solution for countle.org
It uses brute force and some minor heuristics.

# Build

go build

# Run

Update the  call to `calc` in main with required arguments

```
func main() {
        // 6,6,8,2,1,3 are the options in the puzzle
	// 526 is the required result from the puzzle
        // leave other arguments as they are
	calc([]int{6,6,8,2,1,3}, 526,[]*Operation{},0)
}
```
