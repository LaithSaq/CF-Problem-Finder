package models

/*
Represents a submission to a problem.

Field	Description
-------------------
id	Integer.
contestId	Integer. Can be absent.
creationTimeSeconds	Integer. Time, when submission was created, in unix-format.
relativeTimeSeconds	Integer. Number of seconds, passed after the start of the contest (or a virtual start for virtual parties), before the submission.
problem	Problem object.
author	Party object.
programmingLanguage	String.
verdict	Enum: FAILED, OK, PARTIAL, COMPILATION_ERROR, RUNTIME_ERROR, WRONG_ANSWER, PRESENTATION_ERROR, TIME_LIMIT_EXCEEDED, MEMORY_LIMIT_EXCEEDED, IDLENESS_LIMIT_EXCEEDED, SECURITY_VIOLATED, CRASHED, INPUT_PREPARATION_CRASHED, CHALLENGED, SKIPPED, TESTING, REJECTED. Can be absent.
testset	Enum: SAMPLES, PRETESTS, TESTS, CHALLENGES, TESTS1, ..., TESTS10. Testset used for judging the submission.
passedTestCount	Integer. Number of passed tests.
timeConsumedMillis	Integer. Maximum time in milliseconds, consumed by solution for one test.
memoryConsumedBytes	Integer. Maximum memory in bytes, consumed by solution for one test.
points	Floating point number. Can be absent. Number of scored points for IOI-like contests.
*/

type Submission struct {
	Id                  int     `json:"id"`
	ContestId           int     `json:"contestId"`
	CreationTimeSeconds int     `json:"creationTimeSeconds"`
	RelativeTimeSeconds int     `json:"relativeTimeSeconds"`
	Problem             Problem `json:"problem"`
	Author              Party   `json:"author"`
	ProgrammingLanguage string  `json:"programmingLanguage"`
	Verdict             string  `json:"verdict"`
	Testset             string  `json:"testset"`
	PassedTestCount     int     `json:"passedTestCount"`
	TimeConsumedMillis  int     `json:"timeConsumedMillis"`
	MemoryConsumedBytes int     `json:"memoryConsumedBytes"`
	Points              float64 `json:"points"`
}
