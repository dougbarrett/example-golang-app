package handler

import (
	"github.com/dougbarrett/example-golang-app"
	"github.com/dougbarrett/example-golang-app/storage"
)

type handler struct {
	svc storage.Service
}

type Service interface {
	CreateEntry(
		title string,
		name string,
		email string,
		message string,
	) (
		entry guestbook.Entry,
		err error,
	)
	UpdateEntry(
		id uint,
		title string,
		name string,
		email string,
		message string,
	) (
		entry guestbook.Entry,
		err error,
	)
	ListEntries() (
		entries []guestbook.Entry,
		err error,
	)
	DeleteEntry(
		entryID uint,
	) (
		err error,
	)
}

func NewFrontend(
	svc storage.Service,
) (
	Service,
	error,
) {
	h := handler{}
	h.svc = svc
	return &h, nil
}
