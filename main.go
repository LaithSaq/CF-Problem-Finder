package main

import (
	services "cf-problem-finder/service"
	"fmt"
)

func main() {
	allProblems := services.FetchAllProblems()
	fmt.Println(allProblems[0].Id())
	mySolvedProblems := services.FetchSolvedProblemsForUser("NapSaq")
	fmt.Println(len(mySolvedProblems))

}
