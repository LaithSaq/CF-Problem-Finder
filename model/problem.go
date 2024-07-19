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
	contestId int
	problemsetName string
	index string
	name string
	type string
	points float64
	rating int
	tags []string
}