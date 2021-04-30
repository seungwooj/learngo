package main

import (
	"fmt"

	"github.com/seungwooj/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition", hello) // found hello definitnio Greeting
	
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2) // That word already exists
	}

	err3 := dictionary.Update("first", "Not First word")
	if err3 != nil {
		fmt.Println(err3)
	}
	err4 := dictionary.Delete("second")
	if err4 != nil {
		fmt.Println(err4) // Cant delete non-existing word
	}
}
