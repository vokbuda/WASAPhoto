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
type PostDatabase struct {
	ID               uint64        `json:"postid"`
	Text             string        `json:"text"`
	Image            string        `json:"image"`
	Authorid         uint64        `json:"authorid"`
	LastChange       string        `json:"lastChange"`
	Me               bool          `json:"me"`
	QuantityLikes    sql.NullInt64 `json:"quantityLikes"`
	QuantityDislikes sql.NullInt64 `json:"quantityDislikes"`
	CurrentEmotion   sql.NullInt64 `json:"currentemotion"`
}
type Post struct {
	ID               uint64 `json:"postid"`
	Text             string `json:"text"`
	Image            string `json:"image"`
	Authorid         uint64 `json:"authorid"`
	LastChange       string `json:"lastChange"`
	Me               bool   `json:"me"`
	QuantityLikes    int64  `json:"quantityLikes"`
	QuantityDislikes int64  `json:"quantityDislikes"`
	CurrentEmotion   int64  `json:"currentemotion"`
}

type CommentDatabase struct {
	Postid           uint64         `json:"postid"`
	Commentid        uint64         `json:"commentid"`
	Text             string         `json:"text"`
	QuantityLikes    sql.NullInt64  `json:"quantityLikes"`
	QuantityDislikes sql.NullInt64  `json:"quantityDislikes"`
	LastChange       string         `json:"lastChange"`
	Authorid         string         `json:"authorid"`
	Me               bool           `json:"me"`
	Avatar           sql.NullString `json:"avatar"`
	Username         string         `json:"username"`
	CurrentEmotion   sql.NullInt64  `json:"currentemotion"`
}
type Comment struct {
	Postid           uint64 `json:"postid"`
	Commentid        uint64 `json:"commentid"`
	Text             string `json:"text"`
	QuantityLikes    int64  `json:"quantityLikes"`
	QuantityDislikes int64  `json:"quantityDislikes"`
	LastChange       string `json:"lastChange"`
	Authorid         string `json:"authorid"`
	Me               bool   `json:"me"`
	Avatar           string `json:"avatar"`
	Username         string `json:"username"`
	CurrentEmotion   int64  `json:"currentemotion"`
}
type BannedUser struct {
	Userid   uint64         `json:"userid"`
	Username string         `json:"username"`
	Avatar   sql.NullString `json:"avatar"`
}
type Profile struct {
	Userid                uint64         `json:"userid"`
	Username              string         `json:"username"`
	Avatar                sql.NullString `json:"avatar"`
	QuantitySubscribers   sql.NullInt64  `json:"quantitySubscribers"`
	QuantitySubscriptions sql.NullInt64  `json:"quantitySubscriptions"`
	Me                    bool           `json:"me"`
	CurrentSubscribed     bool           `json:"currentFollow"`
	CurrentBanned         bool           `json:"currentBan"`
}
type ProfileClient struct {
	Userid                uint64 `json:"userid"`
	Username              string `json:"username"`
	Avatar                string `json:"avatar"`
	QuantitySubscribers   int64  `json:"quantitySubscribers"`
	QuantitySubscriptions int64  `json:"quantitySubscriptions"`
	Me                    bool   `json:"me"`
	CurrentSubscribed     bool   `json:"currentFollow"`
	CurrentBanned         bool   `json:"currentBan"`
}
type SessionData struct {
	Token     string `json:"token"`
	Userid    uint64 `json:"userid"`
	Created   string `json:"created"`
	Lastlogin string `json:"lastlogin"`
}

type AppDatabase interface {
	// UpdatePost(Post) (string, error) is equaivalent of updatepost my profile check for the values inside of your component

	GetBannedUsers(userid uint64, offset uint64) ([]BannedUser, error)
	BanUser(banninguserid uint64, banneduserid uint64) error
	UnbanUser(banninguserid uint64, banneduserid uint64) error
	UserSearch(searchedData string, offset uint64) ([]ProfileClient, error)
	GetMyStream(userid uint64, offset uint64) ([]Post, error)
	// then implement data for getting current stream inside
	GetProfile(userid uint64, caller uint64) (uint64, uint64, Profile, error)
	GetProfilePosts(userid uint64, caller uint64, offset uint64) ([]Post, error)

	CreateProfilePost(text string, image string, authorid uint64) (uint64, error)
	UpdateProfilePost(postid uint64, text string, image string, uid uint64) (int64, int64, error)
	DeleteProfilePost(postid uint64, authorid uint64) error
	AddLikePost(idPostLiked uint64, idUserLiked uint64) error
	DeleteLikePost(idPostLiked uint64, idUserLiked uint64) error
	AddDislikePost(idPostDisliked uint64, idUserDislike uint64) error
	DeleteDislikePost(idPostDisliked uint64, idUserDislike uint64) error
	FollowUser(idFolloweduser uint64, idFollowingUser uint64) error
	UnfollowUser(idFolloweduser uint64, idFollowingUser uint64) error
	GetCommentsRelatedToPost(postid uint64, offset uint64, userid uint64) ([]Comment, error)
	CreateCommentRelatedToPost(commentText string, authorid uint64, postid uint64) (uint64, error)
	UpdateCommentRelatedToPost(commentid uint64, postid uint64, authorid uint64, text string) (int64, int64, error)
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
	AuthWithPassword(uid uint64, password string) error

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
	const q = "PRAGMA foreign_keys=ON;"
	_, errForeignKeys := db.Exec(q)
	if errForeignKeys != nil {
		return nil, fmt.Errorf("error creating foreign keys data: %w", errForeignKeys)

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
	var banusersTable string
	errbanusers := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='banusers';`).Scan(&banusersTable)
	if errors.Is(errbanusers, sql.ErrNoRows) {
		banusers := `CREATE TABLE "banusers" (
			"banninguserid"	INTEGER NOT NULL,
			"banneduserid"	INTEGER NOT NULL,
			FOREIGN KEY("banninguserid") REFERENCES "profiles"("userid") ON DELETE CASCADE,
			FOREIGN KEY("banneduserid") REFERENCES "profiles"("userid") ON DELETE CASCADE,
			PRIMARY KEY("banninguserid","banneduserid")
		);`

		_, err = db.Exec(banusers)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	var comment_emotionTable string
	err_comment_emotionTable := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comment_emotion';`).Scan(&comment_emotionTable)
	if errors.Is(err_comment_emotionTable, sql.ErrNoRows) {
		comment_emotion := `CREATE TABLE "comment_emotion" (
			"commentid"	INTEGER,
			"userid"	INTEGER,
			"emotion"	INTEGER,
			FOREIGN KEY("userid") REFERENCES "profiles"("userid") ON DELETE CASCADE,
			PRIMARY KEY("commentid","userid"),
			FOREIGN KEY("commentid") REFERENCES "comments"("id") ON DELETE CASCADE
		);`

		_, err = db.Exec(comment_emotion)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	var commentsTable string
	err_commentsTable := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&commentsTable)
	if errors.Is(err_commentsTable, sql.ErrNoRows) {
		comments := `CREATE TABLE "comments" (
			"id"	INTEGER NOT NULL,
			"text"	TEXT,
			"authorid"	INTEGER,
			"postid"	INTEGER,
			"lastupdate"	TEXT,
			FOREIGN KEY("authorid") REFERENCES "profiles"("userid") ON DELETE CASCADE,
			PRIMARY KEY("id" AUTOINCREMENT),
			FOREIGN KEY("postid") REFERENCES "posts"("postid") ON DELETE CASCADE
		);`

		_, err = db.Exec(comments)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}
	var postEmotionTable string
	err_postEmotionTable := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='post_emotion';`).Scan(&postEmotionTable)
	if errors.Is(err_postEmotionTable, sql.ErrNoRows) {
		post_emotion := `CREATE TABLE "post_emotion" (
			"postid"	INTEGER,
			"userid"	INTEGER,
			"emotion"	INTEGER,
			FOREIGN KEY("userid") REFERENCES "profiles"("userid") ON DELETE CASCADE,
			PRIMARY KEY("postid","userid"),
			FOREIGN KEY("postid") REFERENCES "posts"("postid") ON DELETE CASCADE
		);`
		_, err = db.Exec(post_emotion)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}
	var postsTable string
	err_postsTable := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='posts';`).Scan(&postsTable)
	if errors.Is(err_postsTable, sql.ErrNoRows) {
		posts := `CREATE TABLE "posts" (
			"postid"	INTEGER,
			"text"	TEXT,
			"image"	BLOB,
			"authorid"	INTEGER,
			"lastupdate"	TEXT,
			PRIMARY KEY("postid" AUTOINCREMENT),
			FOREIGN KEY("authorid") REFERENCES "profiles"("userid") ON DELETE CASCADE
		);`
		_, err = db.Exec(posts)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	var profilesTable string
	err_profilesTable := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='profiles';`).Scan(&profilesTable)
	if errors.Is(err_profilesTable, sql.ErrNoRows) {
		profiles := `CREATE TABLE "profiles" (
			"userid"	INTEGER NOT NULL,
			"username"	TEXT NOT NULL UNIQUE,
			"password"	TEXT NOT NULL,
			"avatar"	BLOB,
			PRIMARY KEY("userid" AUTOINCREMENT)
		);`
		_, err = db.Exec(profiles)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	var sessionTable string
	err_sessionTable := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='session';`).Scan(&sessionTable)
	if errors.Is(err_sessionTable, sql.ErrNoRows) {
		session := `CREATE TABLE "session" (
			"token"	TEXT UNIQUE,
			"lastlogin"	TEXT,
			"created"	TEXT,
			"userid"	INTEGER,
			PRIMARY KEY("token"),
			FOREIGN KEY("userid") REFERENCES "profiles"("userid") ON DELETE CASCADE
		);`
		_, err = db.Exec(session)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}
	var subscriptionTable string
	err_subscriptionTable := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='subscriptions';`).Scan(&subscriptionTable)
	if errors.Is(err_subscriptionTable, sql.ErrNoRows) {
		subscription := `CREATE TABLE "subscriptions" (
			"followeduserid"	INTEGER NOT NULL,
			"followeruserid"	INTEGER NOT NULL,
			FOREIGN KEY("followeruserid") REFERENCES "profiles"("userid") ON DELETE CASCADE,
			FOREIGN KEY("followeduserid") REFERENCES "profiles"("userid") ON DELETE CASCADE,
			PRIMARY KEY("followeduserid","followeruserid")
		);`
		_, err = db.Exec(subscription)
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
