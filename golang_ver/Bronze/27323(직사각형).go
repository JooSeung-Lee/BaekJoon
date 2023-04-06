package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd                  *bufio.Reader
	wr                  *bufio.Writer
	width, height, area int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	fmt.Fscan(rd, &width, &height)
	area = width * height
	fmt.Fprintln(wr, area)
}
