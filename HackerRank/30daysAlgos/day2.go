package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	var _ = strconv.Atoi // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "
	day2(i, d, s)

}
func day2(i uint64, d float64, s string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	i22 := scanner.Text()
	i2, _ := strconv.ParseUint(i22, 10, 64)
	scanner.Scan()
	d2 := scanner.Text()
	d22, _ := strconv.ParseFloat(d2, 64)

	scanner.Scan()
	d3 := scanner.Text()
	fmt.Println(i2 + i)
	// fmt.Println(d22 + d)
	fmt.Println(fmt.Sprintf("%.1f", d22+d))
	fmt.Println(s + d3)
}
