package userinput

import (
	"fmt"
)

func GetProgrammingLanguage() string {
	var programmingLanguage string

	fmt.Println("Please specify the programming language:")

	fmt.Scanf("%s", &programmingLanguage)

	fmt.Println()

	return programmingLanguage
}
