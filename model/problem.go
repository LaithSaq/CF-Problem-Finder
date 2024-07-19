package models

/*
Represents a problem in Codeforces.

contestId	Integer. Can be absent. Id of the contest, containing the problem.
problemsetName	String. Can be absent. Short name of the problemset the problem belongs to.
index	String. Usually, a letter or letter with digit(s) indicating the problem index in a contest.
name	String. Localized.
type	Enum: PROGRAMMING, QUESTION.
points	Floating point number. Can be absent. Maximum amount of points for the problem.
rating	Integer. Can be absent. Problem rating (difficulty).
tags	String list. Problem tags.
*/

type Problem struct {
	ContestId      int      `json:"contestId"`
	ProblemsetName string   `json:"problemsetName"`
	Index          string   `json:"index"`
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	Points         float64  `json:"points"`
	Rating         int      `json:"rating"`
	Tags           []string `json:"tags"`
}
