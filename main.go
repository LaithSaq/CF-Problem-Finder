package main

import (
	services "cf-problem-finder/service"
	"fmt"
)

func main() {
	problems := services.FetchAllProblems()
	fmt.Print(problems[0].Id())

}
