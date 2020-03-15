package main

import "log"

// User is a in-memory user model.
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// NewMemDb creates and returns a new instance of MemDb.
func NewMemDb() *MemDb {
	return &MemDb{
		users:    []*User{},
		kvsUsers: map[string]int{},
	}
}

// MemDb is a memory database.
type MemDb struct {
	users    []*User
	kvsUsers map[string]int
}

// AddUser stores user in in-memory database.
func (db *MemDb) AddUser(u *User) *User {
	log.Println("AddUser", u)

	ind, ok := db.kvsUsers[u.Username]
	if !ok {
		db.kvsUsers[u.Username] = len(db.users)
		db.users = append(db.users, u)
	} else {
		db.users[ind] = u
	}

	return db.users[ind]
}

// GetUser returns user in in-memory database.
func (db *MemDb) GetUser(username string) *User {
	log.Println("GetUser", username)

	ind, ok := db.kvsUsers[username]
	if ok {
		return db.users[ind]
	}

	return nil
}

// DeleteUser deletes user in in-memory database.
func (db *MemDb) DeleteUser(username string) {
	log.Println("DeleteUser", username)

	if ind, ok := db.kvsUsers[username]; ok {
		db.users[ind] = db.users[len(db.users)-1]
		db.users = db.users[:len(db.users)-1]

		delete(db.kvsUsers, username)
	}
}

// ListUsers returns all users in in-memory database.
func (db *MemDb) ListUsers() []*User {
	log.Println("ListUsers")

	return db.users
}
