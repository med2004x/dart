package main

import "fmt"

type Comment struct {
	Text    string
	Replies []Comment
}

func countComments(comment Comment) int {
	// TODO: count current comment and all recursive replies.
	return 0
}

func main() {
	root := Comment{
		Text: "A",
		Replies: []Comment{
			{Text: "B"},
			{Text: "C", Replies: []Comment{{Text: "D"}}},
		},
	}
	fmt.Println(countComments(root))
}
