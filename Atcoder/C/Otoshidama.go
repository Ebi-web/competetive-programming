package C

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
	available_bills := [3]int{10000, 5000, 1000}
	bills := map[int]int{
		10000: 0,
		5000:  0,
		1000:  0,
	}
	for available_bill := range available_bills {
		if amount < available_bill {
			continue
		}

	}
	for bill := range bills {
		fmt.Print(bill)
	}
}
