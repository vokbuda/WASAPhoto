package database

import "fmt"

/*
these are components which must be in profile user in database::::
avatar: "binary string",
userid: 1,
username: 'User1',
quantitySubscribers: '2.9k',
quantitySubscriptions: '30m'


*/

func (db *appdbimpl) Session(username string, password string, image string) error {

	_, err := db.c.Exec(`insert into profiles(username, image, password) values (?, ?, ?)`,
		username, image, password)
	fmt.Println("Check data inside of your component")

	if err != nil {
		return err
	}

	return nil
}
