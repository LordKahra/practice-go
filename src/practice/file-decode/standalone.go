package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//debugging
	// Uses the Walk function to walk the directory tree
	/*err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Prints the name of each file or directory
		fmt.Println(path)
		return nil
	})*/

	// Open the file.

	file, err := os.Open("practice/file-decode/stuff.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Hello World")
	// Decode the file.
	message := decode(file)
	fmt.Println("Hello World")
	// Print the message.
	fmt.Println(message)
}

func decode(file *os.File) string {
	// Create the scanner and the storage array.
	scanner := bufio.NewScanner(file)
	words := make(map[int]string)
	var keys []int

	// Handle each line.
	for scanner.Scan() {
		// Split the string.
		segments := strings.SplitN(scanner.Text(), " ", 2)
		if len(segments) != 2 {
			continue
		}

		fmt.Println("segments:", segments)

		// Get the number.
		n, err := strconv.ParseInt(segments[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		// If it's triangular, store the result.
		if isTriangular(n) {
			fmt.Println("storing segment " + segments[1])
			words[int(n)] = segments[1]
			keys = append(keys, int(n))
			fmt.Println("words[", n, "]:", words[int(n)])
		} else {
			fmt.Println("not storing?")
		}
	}

	fmt.Println("words", words)
	fmt.Println("keys", keys)

	sort.Ints(keys)

	fmt.Println("keys", keys)

	// Join the string.
	message := ""
	for _, index := range keys {
		fmt.Println(index)
	}
	//message := strings.Join(words, " ")

	return message
}

func isTriangular(n int64) bool {
	// n is triangular if 8n+1 is a perfect square.
	result := 8*n + 1

	// Get the square root of 8n+1, and round it.
	round := int64(math.Sqrt(float64(8*n + 1)))

	// If the square root was an int, return true.
	return (round * round) == result

}
