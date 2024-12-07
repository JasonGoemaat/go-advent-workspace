package main

import "fmt"

type (
	ParentS struct {
		Title string
	}

	ChildS struct {
		ParentS
		Content string
	}

	ParentI interface {
		Init(title string)
		GetTitle() string
	}

	ChildI interface {
		Init(title string, content string)
	}
)

func main() {
	child := ChildS{ParentS: ParentS{Title: "ChildTitle"}, Content: "ChildContent"}
	fmt.Printf("%v\n", child)
}
