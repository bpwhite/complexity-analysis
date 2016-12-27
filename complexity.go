package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"os/exec"
	"sort"
	"strconv"
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

	// initialize counter map to count n-mers
	var counter_map map[string]int
	counter_map = make(map[string]int)

	// number of bs reps to extract
	bs_reps := 5000
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
			_, check_sub := counter_map[substr]
			// if exists, add 1. if not, set to 1
			if check_sub == true {
				counter_map[substr] = counter_map[substr] + 1
			} else {
				counter_map[substr] = 1
			}
		}
	}
	// hold the counts of nmers
	var nmer_counts []int
	nmer_counts = make([]int, 1)

	// print results of hash collection
	for _, value := range counter_map {
		nmer_counts = append(nmer_counts, value)

		//fmt.Println("Key:", key, "Value:", value)
	}

	//sort.Ints(nmer_counts)
	sort.Sort(sort.Reverse(sort.IntSlice(nmer_counts)))
	fmt.Println(nmer_counts)

	// output nmer counts for histogram
	f, err := os.Create("temp.txt")
	check(err)
	defer f.Close()

	for _, nmer := range nmer_counts {
		fmt.Println(nmer)
		t := strconv.Itoa(nmer)
		t = t + "\n"
		f.WriteString(t)
		//fmt.Printf("wrote %d bytes\n", n3)
	}

	f.Sync()
	// end nmer output

	os.Exit(3)

	for index, count := range nmer_counts[0:25] {
		for key, value := range counter_map {
			if value == count {
				fmt.Println("[", index, "] (", key, ")\t", value)
				delete(counter_map, key)
				break
			}

			//fmt.Println("Key:", key, "Value:", value)
		}
	}

	out, err := exec.Command("cmd", "/C", "dir", "C:\\").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
