package models

/*
Member
Represents a member of a party.

Field	Description
-------------------
handle	String. Codeforces user handle.
name	String. Can be absent. User's name if available.
*/

type Member struct {
	Handle string `json:"handle"`
	Name   string `json:"name"`
}
