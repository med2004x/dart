package main

import "fmt"

type Folder struct {
	Name        string
	DirectFiles int
	Children    []Folder
}

func totalFiles(folder Folder) int {
	// TODO: sum direct files and every descendant.
	return 0
}

func main() {
	root := Folder{
		Name:        "root",
		DirectFiles: 1,
		Children: []Folder{
			{Name: "documents", DirectFiles: 2},
			{Name: "pictures", DirectFiles: 3},
		},
	}
	fmt.Println(totalFiles(root))
}
