package main

import (
	"flag"
	"fmt"
	"strconv"
)

func IntToString(a [3]int) [3]string {
	var numbers = []string{ //initialization of numbers array
		"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
	}

	var words [len(a)]string

	for i := 0; i<len(a); i++{

		if a[i] <= 9{
			words[i] = numbers[a[i]]
		} else{
			var div = 1
			for k := a[i]; k>0; k=k/10{
				div *=10

			}
			var mod = div
			div = div/10


			for j:= a[i]; div>0; j = j % mod {
				mod /=10

				words[i] += numbers[j/div]
				div /= 10
			}
		}
	}
	return words
}

var one string

func main() {
	flag.StringVar(&one, "", "123 2 3", "arg one")
	flag.Parse()
	tail := flag.Args()
	var a [3]int
	a[0], _ = strconv.Atoi(tail[0])
	a[1], _ = strconv.Atoi(tail[1])
	a[2], _ = strconv.Atoi(tail[2])

	fmt.Println(IntToString(a))

}



