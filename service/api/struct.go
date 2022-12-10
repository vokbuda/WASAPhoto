package api

import "git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"

const (
	FountainStatusGood   string = "good"
	FountainStatusFaulty string = "faulty"
)

// Fountain struct represent a fountain in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See Fountain.FromDatabase (below) to understand why.
type Fountain struct {
	ID        uint64  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status    string  `json:"status"`
}
type PostCreate struct {
	Postid   uint64 `json:"postid"`
	Authorid uint64 `json:"authorid"`
	Text     string `json:"text"`
	Image    string `json:"image"`
}

type DataAccountUpdate struct {
	NewValue string `json:"newValue"`
}
type DataAccountUpdated struct {
	Entity string `json:"entity"`
}

type BanningUser struct {
	BanningUser uint64 `json:"banningUser"`
	BannedUser  uint64 `json:"bannedUser"`
}

type PostToChange struct {
	Postid uint64 `json:"postid"`
	Text   string `json:"text"`
	Image  string `json:"image"`
}
type CommentToCreate struct {
	CommentId uint64 `json:"commentid"`
	PostId    uint64 `json:"postid"`
	Text      string `json:"text"`
	Authorid  uint64 `json:"authorid"`
}

type CommentToChange struct {
	CommentId uint64 `json:"commentid"`
	PostId    uint64 `json:"postid"`
	Text      string `json:"text"`
}
type CommentToDelete struct {
	CommentId uint64 `json:"commentid"`
	PostId    uint64 `json:"postid"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SessionUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
type SessionResponse struct {
	Session string `json:"session"`
	Uid     uint64 `json:"uid"`
}

type Subscription struct {
	FollowedUserId  uint64 `json:"followedUserId"`
	FollowingUserId uint64 `json:"followingUserId"`
}
type RequestEmotionToPost struct {
	IdPostEmotion uint64 `json:"idPostEmotion"`
	IdUser        uint64 `json:"idUser"`
}
type RequestEmotionToComment struct {
	IdPost           uint64 `json:"idPost"`
	IdCommentEmotion uint64 `json:"idCommentEmotion"`
	IdUser           uint64 `json:"idUser"`
}
type CommentToUpdate struct {
	CommentId uint64 `json:"CommentId"`
	PostId    uint64 `json:"PostId"`
	text      string `text:"text"`
}

type CommentChanged struct {
	Postid           uint64 `json:"postid"`
	Commentid        uint64 `json:"commentid"`
	QuantityLikes    string `json:"quantityLikes"`
	QuantityDislikes string `json:"quantityDislikes"`
}
type PostChanged struct {
	Postid           uint64 `json:"postid"`
	QuantityLikes    string `json:"quantityLikes"`
	QuantityDislikes string `json:"quantityDislikes"`
}
type PostCreated struct {
	Postid uint64 `json:"postid"`
}
type CommentCreated struct {
	Commentid uint64 `json:"commentid"`
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.Fountain" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
func (f *Fountain) FromDatabase(fountain database.Fountain) {
	f.ID = fountain.ID
	f.Latitude = fountain.Latitude
	f.Longitude = fountain.Longitude
	f.Status = fountain.Status
}

/*
func (p *Profile) profileToDatabase() database.Profile{
	return database.Profile{
		Username: p.username,
		Avatar: p.Avatar,
		QuantitySubscribers: p.QuantitySubscribers,
		QuantitySubscriptions: p.QuantitySubscriptions,
	}
}*/
// ToDatabase returns the fountain in a database-compatible representation
func (f *Fountain) ToDatabase() database.Fountain {
	return database.Fountain{
		ID:        f.ID,
		Latitude:  f.Latitude,
		Longitude: f.Longitude,
		Status:    f.Status,
	}
}

// IsValid checks the validity of the content. In particular, coordinates should be in their range of validity, and the
// status should be either FountainStatusGood or FountainStatusFaulty. Note that the ID is not checked, as fountains
// read from requests have zero IDs as the user won't send us the ID in that way.
func (f *Fountain) IsValid() bool {
	return -90 <= f.Latitude && f.Latitude <= 90 &&
		-180 <= f.Longitude && f.Longitude <= 180 &&
		(f.Status == FountainStatusGood || f.Status == FountainStatusFaulty)
}
