package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"

	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// parameters
	// number of bs reps to extract
	bs_reps := 10
	subs_samples := 5000
	// nmer size (initial/default)
	var nmer_size int64
	nmer_size = 3

	// input file text
	b, err := ioutil.ReadFile("file2.txt")
	if err != nil {
		fmt.Print(err)
	}

	// output logs
	f, err := os.Create("temp.txt")
	check(err)
	defer f.Close()

	// convert content to a 'string'
	str := string(b)

	// capture words as fields
	words := strings.Fields(str)
	words_cleared := []string{""}

	block_list := []string{"and", "the", "are", "for", "of", "is", "to"}
	for _, word := range words {
		block := 0
		for _, b := range block_list {
			if word == b || len(word) <= 2 {
				//fmt.Println(word)
				//words = words[:i+copy(words[i:], words[i+1:])]
				block = 1
				break
			}
		}
		if block == 0 {
			words_cleared = append(words_cleared, word)
		}
	}
	//fmt.Println(words_cleared)
	words = words_cleared
	//os.Exit(3)

	// clean string for processing
	str = strings.Join(words, " ")
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\r\n", "", -1)
	str = str + "\n"
	f.WriteString(str)

	fmt.Println("Words: ", len(words))
	var length int64
	length = int64(len(str))
	fmt.Println("Characters: ", length)

	// initialize word map to count word frequency
	var word_map map[string]int
	word_map = make(map[string]int)

	for rep := 0; rep <= bs_reps; rep++ {
		// initialize counter map to count n-mers
		var counter_map map[string]int
		counter_map = make(map[string]int)

		for i := 0; i <= subs_samples; i++ {
			// generate random number for start point of substring
			ran := gen_cryp_num(length)
			// add length of n-mer to substring position
			end := ran + nmer_size
			// check to make sure substring does not extend beyond string length
			if end < length {
				// select substring
				substr := str[ran:end]
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

		// store results of hash collection in array
		for _, value := range counter_map {
			nmer_counts = append(nmer_counts, value)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(nmer_counts)))

		word_pos := make([]int, len(words))
		// loop through sorted nmer count ints
		for _, count := range nmer_counts[0:25] {
			// loop through map containing nmer and count
			for key, value := range counter_map {
				// if value matches current count, print mapping
				if value == count {
					v_outp := []string{"(", key, ")\t", strconv.Itoa(value), "\n"}
					f.WriteString(strings.Join(v_outp, ""))

					delete(counter_map, key)
					// check through the list of words to see what words
					// the nmer is found in
					for pos, word := range words {
						if strings.Contains(word, key) {
							pos_outp := []string{"\t", word, "(", strconv.Itoa(pos), ")", "\n"}
							f.WriteString(strings.Join(pos_outp, ""))
							word_pos[pos] = 1
						}
						// only report nmers found in words //
					}
					break
				}
				//fmt.Println("Key:", key, "Value:", value)
			}
		}

		for i, word := range words {
			for j, word_pos := range word_pos {
				if i == j {
					fmt.Println(word_pos, "\t", word)
					if word_pos == 1 {
						// check if word has been found previously
						_, check_word := word_map[word]
						// if exists, add 1. if not, set to 1
						if check_word == true {
							word_map[word] = word_map[word] + 1
						} else {
							word_map[word] = 1
						}
					}
				}
			}
		}
		fmt.Println(word_pos)
	}
	//f.WriteString("(", pos, ")", word)
	//}

	fmt.Println(word_map)
	f.Sync()
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
