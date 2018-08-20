package db

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Tag is extra information.
type Tag struct {
	gorm.Model

	Text string
}

// Create a tag.
func (a *Tag) Create() error {
	return errors.New("need code")
}

// Update a tag.
func (a *Tag) Update() error {
	return errors.New("need code")
}

// Delete a tag.
func (a *Tag) Delete() error {
	return errors.New("need code")
}

// Find a single tag that matches a.
func (a *Tag) Find() error {
	return errors.New("need code")
}

// FindAll tags that match a.
func (a *Tag) FindAll(v interface{}) error {
	return errors.New("need code")
}
