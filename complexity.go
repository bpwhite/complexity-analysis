package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"sort"
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

	// initialize counter hash to count n-mers
	var counter map[string]int
	counter = make(map[string]int)

	// number of bs reps to extract
	bs_reps := 500
	for i := 0; i <= bs_reps; i++ {
		// generate random number for start point of substring
		ran := gen_cryp_num(length)
		// add length of n-mer to substring position
		end := ran + 3
		// check to make sure substring does not extend beyond string length
		if end < length {
			// select substring
			substr := str[ran:end]
			// print substr stats
			//fmt.Println("Ran: ", ran, " End: ", end)
			//fmt.Printf("%d, %s , %d\n", i, substr, ran)

			// check if substr has been extracted previously
			_, check_sub := counter[substr]
			// if exists, add 1. if not, set to 1
			if check_sub == true {
				counter[substr] = counter[substr] + 1
			} else {
				counter[substr] = 1
			}
		}
	}
	var nmer_counts []int
	nmer_counts = make([]int, 1)
	// print results of hash collection
	for _, value := range counter {
		nmer_counts = append(nmer_counts, value)

		//fmt.Println("Key:", key, "Value:", value)
	}

	//sort.Ints(nmer_counts)
	sort.Sort(sort.Reverse(sort.IntSlice(nmer_counts)))
	fmt.Println(nmer_counts)
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
