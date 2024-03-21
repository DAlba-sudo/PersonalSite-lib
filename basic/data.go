package basic

import "gorm.io/gorm"

type BasicContentType struct {
	gorm.Model

	Name    string
	Summary string
}

type BasicContentBlurb struct {
	gorm.Model

	Title          string
	BasicContentID int
}

type BasicContent struct {
	gorm.Model

	Title              string
	Desc               string
	Body               string
	BasicContentTypeID int
	BasicContentType   BasicContentType
}
