package services

import (
	models "cf-problem-finder/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchAllProblems() []models.Problem {
	resp, _ := http.Get("https://codeforces.com/api/problemset.problems")
	jsonData := map[string]interface{}{}
	json.NewDecoder(resp.Body).Decode(&jsonData)
	problemInterfaces := jsonData["result"].(map[string]interface{})["problems"].([]interface{})
	var problems []models.Problem

	for _, problemInterface := range problemInterfaces {
		var problem models.Problem
		parseInterfaceIntoObject(problemInterface, &problem)
		problems = append(problems, problem)
	}

	problemStatisticsInterfaces := jsonData["result"].(map[string]interface{})["problemStatistics"].([]interface{})

	for i, problemStatisticsInterface := range problemStatisticsInterfaces {
		problems[i].SolvedCount = int(problemStatisticsInterface.(map[string]interface{})["solvedCount"].(float64))
	}

	return problems
}

func parseInterfaceIntoObject(interfaceToParse interface{}, object interface{}) error {
	interfaceJSON, err := json.Marshal(interfaceToParse)
	if err != nil {
		fmt.Printf("Error marshaling %T", interfaceToParse)
		return err
	}
	err = json.Unmarshal(interfaceJSON, object)
	if err != nil {
		// Handle the error during unmarshaling
		fmt.Printf("Error unmarshaling %T", object)
		return err
	}

	return err
}
