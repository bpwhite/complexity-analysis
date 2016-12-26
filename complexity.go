package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("file.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	fmt.Println(str) // print the content as a 'string'
	str = strings.Replace(str, " ", "", -1)
	fmt.Println(strings.Replace(str, "\r\n", "", -1))

	words := strings.Fields(str)

	fmt.Println(words, len(words)) // [one two three four] 4

	for i := 0; i < len(words); i++ {

		fmt.Printf("%d, %s\n", i, words[i])
	}
}
