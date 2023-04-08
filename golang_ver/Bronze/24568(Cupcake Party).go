package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd                                        *bufio.Reader
	wr                                        *bufio.Writer
	regular_box, small_box, sum_cupcakes, ans int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	fmt.Fscan(rd, &regular_box, &small_box)
	sum_cupcakes = (8 * regular_box) + (3 * small_box)
	ans = sum_cupcakes - 28
	fmt.Fprintln(wr, ans)
}
