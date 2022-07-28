package main

import (
	"fmt"
	"locvis/userinput"
)

func main() {
	var programmingLanguage string = userinput.GetProgrammingLanguage()
	fmt.Println(programmingLanguage)
}
