package services

import (
	models "cf-problem-finder/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchProblems() string {
	resp, _ := http.Get("https://codeforces.com/api/problemset.problems")
	jsonData := map[string]interface{}{}
	json.NewDecoder(resp.Body).Decode(&jsonData)
	problemsInterface, ok := jsonData["result"].(map[string]interface{})["problems"].([]interface{})

	if !ok {
		fmt.Println("Type assertion failed")
		return "Error"
	}

	var problems []models.Problem
	for _, problemInterface := range problemsInterface {
		problemJSON, err := json.Marshal(problemInterface)
		if err != nil {
			fmt.Println("Error marshaling problem:", err)
			continue
		}

		var problem models.Problem
		err = json.Unmarshal(problemJSON, &problem)
		if err != nil {
			// Handle the error during unmarshaling
			fmt.Println("Error unmarshaling problem:", err)
			continue
		}

		problems = append(problems, problem)
	}

	fmt.Println("-----------------")
	fmt.Println(len(problems))
	fmt.Println("-----------------")
	fmt.Println(problems[0])
	return "Hello"
}
