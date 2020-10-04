package main

import (
	"fmt"

	"github.com/childult-programmer/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"First": "First word"}

	// test Search method
	def, errSearch := dictionary.Search("First") // next run, change "First" to "Second"
	// Error occurred
	if errSearch != nil {
		fmt.Println(errSearch)
	} else {
		fmt.Println(def)
	}

	// test Add method
	errAdd := dictionary.Add("hello", "Greeting")
	if errAdd != nil {
		fmt.Println(errAdd)
	}
	fmt.Println(dictionary)

	// test Update method
	word := "Bye"
	dictionary.Add(word, "Say Bye")
	dictionary.Update(word, "See you tomorrow")
	fmt.Println(dictionary)

	// test Delete method => delete "First"
	dictionary.Delete("First")
	fmt.Println(dictionary)
}
