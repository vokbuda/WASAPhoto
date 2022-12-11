package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrFountainDoesNotExist = errors.New("fountain does not exist")
var ErrPostDoesNotExist = errors.New("post doesn't exist")
var ErrCommentDoesNotExist = errors.New("comment doesn't exist")
var ErrLikePostDoesNotExist = errors.New("like post doesn't exist")
var ErrDislikePostDoesNotExist = errors.New("dislike post doesn't exist")
var ErrLikeCommentDoesNotExist = errors.New("like comment doesn't exist")
var ErrDislikeCommentDoesNotExist = errors.New("dislike comment doesn't exist")
var ErrNotAuthorized = errors.New("U don't have authorization")

// then when it is necessary check for the rest of components inside of your data

// Fountain struct represent a fountain in every API call between this package and the outside world.
// Note that the internal representation of fountain in the database might be different.
type Fountain struct {
	ID        uint64
	Latitude  float64
	Longitude float64
	Status    string
}

// adjust post component below and check for re
type Post struct {
	ID         uint64 `json:"postid"`
	Text       string `json:"text"`
	Image      string `json:"image"`
	Authorid   uint64 `json:"authorid"`
	LastChange string `json:"lastChange"`
}
type Comment struct {
	Postid           uint64 `json:"postid"`
	Commentid        uint64 `json:"commentid"`
	Text             string `json:"text"`
	QuantityLikes    string `json:"quantityLikes"`
	QuantityDislikes string `json:"quantityDislikes"`
	LastChange       string `json:"lastChange"`
	Authorid         string `json:"authorid"`
}
type Profile struct {
	Userid                uint64         `json:"userid"`
	Username              string         `json:"username"`
	Avatar                sql.NullString `json:"avatar"`
	QuantitySubscribers   string         `json:"quantitySubscribers"`
	QuantitySubscriptions string         `json:"quantitySubscriptions"`
	Me                    bool           `json:"me"`
}
type SessionData struct {
	Token     string `json:"token"`
	Userid    uint64 `json:"userid"`
	Created   string `json:"created"`
	Lastlogin string `json:"lastlogin"`
}

type AppDatabase interface {
	// UpdatePost(Post) (string, error) is equaivalent of updatepost my profile check for the values inside of your component

	GetBannedUsers(userid uint64, offset uint64) ([]Profile, error)
	BanUser(banninguserid uint64, banneduserid uint64) error
	UnbanUser(banninguserid uint64, banneduserid uint64) error
	UserSearch(searchedData string, offset uint64) ([]Profile, error)
	GetMyStream(userid uint64, offset uint64) ([]Post, error)
	// then implement data for getting current stream inside
	GetProfile(userid uint64) (uint64, uint64, Profile, error)
	GetProfilePosts(userid uint64, offset uint64) ([]Post, error)

	CreateProfilePost(text string, image string, authorid uint64) (uint64, error)
	UpdateProfilePost(postid uint64, text string, image string, uid uint64) (uint64, uint64, error)
	DeleteProfilePost(postid uint64, authorid uint64) error
	AddLikePost(idPostLiked uint64, idUserLiked uint64) error
	DeleteLikePost(idPostLiked uint64, idUserLiked uint64) error
	AddDislikePost(idPostDisliked uint64, idUserDislike uint64) error
	DeleteDislikePost(idPostDisliked uint64, idUserDislike uint64) error
	FollowUser(idFolloweduser uint64, idFollowingUser uint64) error
	UnfollowUser(idFolloweduser uint64, idFollowingUser uint64) error
	GetCommentsRelatedToPost(postid uint64, offset uint64) ([]Comment, error)
	CreateCommentRelatedToPost(commentText string, authorid uint64, postid uint64) (uint64, error)
	UpdateCommentRelatedToPost(commentid uint64, postid uint64, authorid uint64, text string) (uint64, uint64, error)
	DeleteCommentRelatedToPost(idForPost uint64, idForComment uint64, authorid uint64) error

	// u should implement methods below written inside of your database activity

	AddLikeToCommentRelatedToPost(commentid uint64, userid uint64) error
	RemoveLikeFromCommentRelatedToPost(commentid uint64, userid uint64) error
	AddDislikeToCommentRelatedToPost(commentid uint64, userid uint64) error
	RemoveDislikeFromCommentRelatedToPost(commentid uint64, userid uint64) error

	ChangePassword(userid uint64, password string) error
	ChangeAvatar(userid uint64, avatar string) error
	ChangeUsername(userid uint64, username string) error
	DeleteAccount(userid uint64) error
	PostAuthUidCheck(postauthorid uint64) (uint64, error)

	// below u have component for register different users
	Session(username string, password string, bearerToken string) (uint64, string, error)
	// then in case of username u should remove data from current accoutn

	Auth(token string) error
	AuthUid(token string) (uint64, error)

	// and then after all implement token inside of your component

	// ListFountainsWithFilter returns the list of fountains with their status, filtered using the specified parameters

	// Ping checks whether the database is available or not (in that case, an error will be returned)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='fountains';`).Scan(&tableName)
	// below u should create data for your component and check for the rest

	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE fountains (
    id INTEGER NOT NULL PRIMARY KEY,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    status TEXT NOT NULL
	);`

		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
