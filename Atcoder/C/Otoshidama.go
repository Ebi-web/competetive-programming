package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	num_bills, _ := strconv.Atoi(input[0])
	amount, _ := strconv.Atoi(input[1])
	bills := map[int]int{
		10000: -1,
		5000:  -1,
		1000:  -1,
	}
	for i := 0; i <= num_bills; i++ {
		for j := 0; j <= num_bills-i; j++ {
			k := num_bills - i - j
			if 10000*i+5000*j+1000*k == amount {
				bills[10000] = i
				bills[5000] = j
				bills[1000] = k
				break
			}
		}
	}
	//	stdout bills with space
	fmt.Print(bills[10000], " ", bills[5000], " ", bills[1000])
}
