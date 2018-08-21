package db

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User for lakuh.
type User struct {
	gorm.Model

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create a user.
func (u *User) Create() error {

	if u.Name == "" {
		return errors.New("need a name")
	}

	if u.Email == "" {
		return errors.New("need a valid email")
	}

	t := User{Email: u.Email}

	if err := t.Find(); err == nil {
		return errors.New("email exists")
	}

	if u.Password == "" {
		return errors.New("need a password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return DB.Create(&u).Error
}

// Update a song.
func (u *User) Update() error {

	if err := u.Find(); err != nil {
		return err
	}

	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		u.Password = string(hash)
	}

	return DB.Model(&u).Update(&u).Error
}

// Delete a song.
func (u *User) Delete() error {
	return DB.Delete(&u).Error
}

// Find a single artist that matches a.
func (u *User) Find() error {
	return DB.Where(&u).First(&u).Error
}

// FindAll users that match a.
func (u *User) FindAll(v interface{}) error {
	return DB.Where(&u).Find(&v).Error
}

// Login will return nil if the data is correct.
func (u *User) Login() error {

	e := errors.New("missing valid email and or password")

	if u.Email == "" || u.Password == "" {
		return e
	}

	found := User{
		Email: u.Email,
	}

	if err := found.Find(); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(u.Password)); err != nil {
		return err
	}

	return nil
}
