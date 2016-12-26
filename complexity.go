package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("file.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	//fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	//fmt.Println(str) // print the content as a 'string'
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\r\n", "", -1)
	fmt.Println(str)

	//words := strings.Fields(str)

	//fmt.Println(words, len(words)) // [one two three four] 4
	var length int64
	length = int64(len(str))
	fmt.Println("Characters: ", length)
	bs_reps := 100
	for i := 0; i <= bs_reps; i++ {
		ran := gen_cryp_num(length)
		end := ran + 3
		fmt.Println("Ran: ", ran, " End: ", end)
		fmt.Printf("%d, %s , %d\n", i, str[ran:end], ran)

	}
}

// Crypto rand int number
func gen_cryp_num(input int64) (n int64) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(input))
	if err != nil {
		panic(err)
	}
	n = nBig.Int64()
	//fmt.Printf("Here is a random %T in [0,27) : %d\n", n, n)
	return
}
