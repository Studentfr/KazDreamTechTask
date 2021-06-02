package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func sequentialSolution() {
	start := time.Now()
	file, err := ioutil.ReadFile("mobydick.txt")

	if err != nil {
		log.Fatal(err)
	}
	var word []byte
	var wordsSlice [][]byte

	for i := 0; i < len(file)-1; i++ {
		if file[i] >= 97 && file[i] <= 122 {
			word = append(word, file[i])
			continue
		}
		if file[i] >= 65 && file[i] <= 90{
			word = append(word, file[i] + 32)
			continue
		}
		if len(word) > 0 {
			wordsSlice = append(wordsSlice, word)
		}
		word = []byte{}
	}

	var occurrencesCount []uint
	var occurredWords [][]byte
	var index int

	for i := 0; i < len(wordsSlice); i++ {
		if occurredWords != nil {
			index = isOccurred(&occurredWords, &wordsSlice[i])
			if index == -1 {
				occurredWords = append(occurredWords, wordsSlice[i])
				occurrencesCount = append(occurrencesCount, 1)
				continue
			}
			occurrencesCount[index] += 1
			continue
		}else{
			occurredWords = append(occurredWords, wordsSlice[i])
			occurrencesCount = append(occurrencesCount, 1)
		}
	}

	for i := 0; i<len(occurrencesCount)-1; i++{
		for j:= 0; j<len(occurrencesCount)-i-1; j++{
			if occurrencesCount[j] > occurrencesCount[j + 1] {
				cache1 := occurrencesCount[j]
				occurrencesCount[j] = occurrencesCount[j+1]
				occurrencesCount[j+1] = cache1

				cache2 := occurredWords[j]
				occurredWords[j] = occurredWords[j+1]
				occurredWords[j+1] = cache2
			}
		}
	}

	for i := 0; i < 20; i++ {
		print(occurrencesCount[len(occurrencesCount)-i-1])
		println(" " + string(occurredWords[len(occurrencesCount)-i-1]))
	}
	fmt.Printf("Process took %s\n", time.Since(start))
}

func isOccurred(arr *[][]byte, word *[]byte) int{
	for i := 0; i < len(*arr); i++ {
		if bytes.Equal((*arr)[i], *word) == true {
			return i
		}
	}
	return -1
}
