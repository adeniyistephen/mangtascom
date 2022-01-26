package business

// WordCount is a struct that stores the word and count
type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

// Word is the struct that holds the text to be counted
type Word struct {
	Text string `json:"text"`
}
