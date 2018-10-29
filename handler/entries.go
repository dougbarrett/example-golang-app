package handler

import (
	"errors"

	"github.com/dougbarrett/example-golang-app"
)

func (h *handler) CreateEntry(
	title string,
	name string,
	email string,
	message string,
) (
	entry guestbook.Entry,
	err error,
) {
	entry.Title = title
	entry.Name = name
	entry.Email = email
	entry.Message = message

	err = h.svc.SaveEntry(&entry)
	return
}

func (h *handler) UpdateEntry(
	id uint,
	title string,
	name string,
	email string,
	message string,
) (
	entry guestbook.Entry,
	err error,
) {
	entry, err = h.svc.GetEntry(id)
	if err != nil {
		return
	}

	entry.Title = title
	entry.Name = name
	entry.Email = email
	entry.Message = message

	err = h.svc.SaveEntry(&entry)
	return
}

func (h *handler) ListEntries() (
	entries []guestbook.Entry,
	err error,
) {
	return h.svc.GetAllEntries()
}

func (h *handler) DeleteEntry(
	entryID uint,
) (
	err error,
) {
	entry, err := h.svc.GetEntry(entryID)

	if entry.ID == 0 || err != nil {
		return errors.New("Entry doesn't exist")
	}

	return h.svc.DeleteEntry(entryID)
}
