package main

import "fmt"

type Dependency struct {
	Job          string
	Prerequisite string
}

func plan(jobs []string, dependencies []Dependency) ([]string, error) {
	// TODO: implement deterministic topological sort.
	return nil, nil
}

func main() {
	jobs := []string{"generate", "compile", "test", "deploy"}
	dependencies := []Dependency{
		{Job: "compile", Prerequisite: "generate"},
		{Job: "test", Prerequisite: "compile"},
		{Job: "deploy", Prerequisite: "test"},
	}
	fmt.Println(plan(jobs, dependencies))
}
