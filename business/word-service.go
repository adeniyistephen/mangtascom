package business

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

// NewWordCount manages the business logic for word counter
type NewWordCount struct {
	log *log.Logger
}

// New returns a new instance of the business package.
func New(log *log.Logger) NewWordCount {
	return NewWordCount{log}
}

func  (nwc NewWordCount) WordCounter(text Word) ([]WordCount, error) {
	// splits the string text around each instance of one or more white space characters, returning a slice.
	list := strings.Fields(text.Text)

	// Create a map of words and their counts
	counts := make(map[string]int)
	for _, word := range list {
		_, ok := counts[word] // check if current key is in counts?
		if ok {               // if yes, increment
			counts[word] += 1
		} else {
			// if not, add it to the map
			counts[word] = 1
		}
	}
	fmt.Println(counts)

	// we create a slice of WordCount structs to store the key and value of counts results for easy sorting
	wordcounts := make([]WordCount, 0, len(counts))
	for key, value := range counts {
		wordcounts = append(wordcounts, WordCount{key, value})
	}

	// sort the slice of WordCount structs by the value of Count
	sort.Slice(wordcounts, func(i, j int) bool {
		return wordcounts[i].Count > wordcounts[j].Count
	})

	// print 3 of the most frequent words
	for i := 0; i < len(wordcounts) && i < 10; i++ {
		fmt.Printf("%s: %d\n", wordcounts[i].Word, wordcounts[i].Count)
	}

	return wordcounts, nil
}
