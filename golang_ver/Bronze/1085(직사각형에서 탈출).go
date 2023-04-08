package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd              *bufio.Reader
	wr              *bufio.Writer
	x, y, w, h, ans int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	fmt.Fscan(rd, &x, &y, &w, &h)
	ans = x
	if ans > y {
		ans = y
	}
	if ans > w-x {
		ans = w - x
	}
	if ans > h-y {
		ans = h - y
	}
	fmt.Fprintln(wr, ans)
}
