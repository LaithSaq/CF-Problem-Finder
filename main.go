package main

import (
	services "cf-problem-finder/service"
	"fmt"
)

func main() {
	problem := services.FetchProblems()
	fmt.Println(problem)
}
