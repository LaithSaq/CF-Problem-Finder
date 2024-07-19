package models

/*
Represents a Codeforces user.

Field	Description
-------------------
handle	String. Codeforces user handle.
email	String. Shown only if user allowed to share his contact info.
vkId	String. User id for VK social network. Shown only if user allowed to share his contact info.
openId	String. Shown only if user allowed to share his contact info.
firstName	String. Localized. Can be absent.
lastName	String. Localized. Can be absent.
country	String. Localized. Can be absent.
city	String. Localized. Can be absent.
organization	String. Localized. Can be absent.
contribution	Integer. User contribution.
rank	String. Localized.
rating	Integer.
maxRank	String. Localized.
maxRating	Integer.
lastOnlineTimeSeconds	Integer. Time, when user was last seen online, in unix format.
registrationTimeSeconds	Integer. Time, when user was registered, in unix format.
friendOfCount	Integer. Amount of users who have this user in friends.
avatar	String. User's avatar URL.
titlePhoto	String. User's title photo URL.
*/

type User struct {
	Handle                  string `json:"handle"`
	Email                   string `json:"email"`
	VkId                    string `json:"vkId"`
	OpenId                  string `json:"openId"`
	FirstName               string `json:"firstName"`
	LastName                string `json:"lastName"`
	Country                 string `json:"country"`
	City                    string `json:"city"`
	Organization            string `json:"organization"`
	Contribution            int    `json:"contribution"`
	Rank                    string `json:"rank"`
	Rating                  int    `json:"rating"`
	MaxRank                 string `json:"maxRank"`
	MaxRating               int    `json:"maxRating"`
	LastOnlineTimeSeconds   int    `json:"lastOnlineTimeSeconds"`
	RegistrationTimeSeconds int    `json:"registrationTimeSeconds"`
	FriendOfCount           int    `json:"friendOfCount"`
	Avatar                  string `json:"avatar"`
	TitlePhoto              string `json:"titlePhoto"`
}
