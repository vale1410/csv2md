package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var filename = flag.String("f", ";", "Name of File.")
var separator = flag.String("s", ";", "Column Separator.")

func main() {

	flag.Parse()

	input, err := ioutil.ReadFile(*filename)

	if err != nil {
		panic("Please specifiy correct path to instance. Does not exist")
	}

	lines := strings.Split(string(input), "\n")

	if len(lines) == 0 {
		return
	}

	nc := len(strings.FieldsFunc(lines[0], sep))
	width := make([]int, nc)

	for _, l := range lines {
		if l == "" {
			continue
		}
		entries := strings.FieldsFunc(l, sep)
		for i, e := range entries {

			if width[i] < len(e)+2 {
				width[i] = len(e) + 2
			}
		}
	}

	for i, l := range lines {
		if l == "" {
			continue
		}

		//if i == 0 {
		//	for _, x := range width {
		//		printN("-", x)
		//		fmt.Print("--")
		//	}
		//	fmt.Println("-")
		//}

		fmt.Print("| ")

		if i == 1 {
			for _, x := range width {
				printN("-", x)
				fmt.Print("| ")
			}
			fmt.Println()
			fmt.Print("| ")
		}

		entries := strings.FieldsFunc(l, sep)

		for j, e := range entries {
			var space int
			if i == 0 {
				fmt.Print("*", e, "*")
				space = width[j] - len(e) - 2
			} else {
				fmt.Print(e)
				space = width[j] - len(e)
			}
			printN(" ", space)
			fmt.Print("| ")
		}
		fmt.Println()
		//if i == len(lines)-2 {
		//	for _, x := range width {
		//		printN("-", x)
		//		fmt.Print("--")
		//	}
		//	fmt.Println("-")
		//}
	}
	return
}

func printN(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(s)
	}
}

func sep(c rune) bool {
	return string(c) == *separator
}
