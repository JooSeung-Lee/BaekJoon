package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd                         *bufio.Reader
	wr                         *bufio.Writer
	num1, num2, resGcd, resLcm int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func gcd(m, n int) int {
	if m < n {
		m, n = n, m
	}
	if n == 0 {
		return m
	}
	if m%n == 0 {
		return n
	}
	return gcd(n, m%n)

}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func main() {
	defer wr.Flush()
	fmt.Fscan(rd, &num1, &num2)
	resGcd = gcd(num1, num2)
	resLcm = lcm(num1, num2)
	fmt.Fprintln(wr, resGcd)
	fmt.Fprintln(wr, resLcm)
}
