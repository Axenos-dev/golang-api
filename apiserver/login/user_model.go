package login

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string
	Password string
}

type UserInDB struct {
	Id            string
	Username      string
	Password_hash string
}

func (user *User) compare_passwords(user_from_db UserInDB) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user_from_db.Password_hash), []byte(user.Password))
	return err == nil
}
