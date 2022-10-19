package main

import (
	"strings"
)

func convert(s string, numRows int) string {
	slicedS := strings.Split(s, "")
	plane := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		plane[i] = make([]string, len(s))
	}
	currentCol := 0
	for len(slicedS) > 0 {
		for i := 0; i < numRows; i++ {
			if len(slicedS) == 0 {
				break
			}
			plane[i][currentCol] = slicedS[0]
			slicedS = slicedS[1:]
		}
		currentCol++
		for i := numRows - 2; i > 0; i-- {
			if len(slicedS) == 0 {
				break
			}
			plane[i][currentCol] = slicedS[0]
			slicedS = slicedS[1:]
			currentCol++
		}
	}
	str := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(plane[i]); j++ {
			str += plane[i][j]
		}
	}
	return str
}
