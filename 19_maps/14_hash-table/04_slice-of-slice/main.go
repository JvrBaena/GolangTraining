package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([][]string, 12)

	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucketRemainder(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	fmt.Println(buckets[1])

}

func HashBucketRemainder(word string, numBuckets int) int {
	letter := int(word[0])
	bucket := letter % numBuckets
	return bucket
}
