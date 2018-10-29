package storage

import (
	"github.com/dougbarrett/example-golang-app"
)

type Service interface {
	GetAllEntries() (
		entries []guestbook.Entry,
		err error,
	)
	SaveEntry(
		entry *guestbook.Entry,
	) (
		err error,
	)
	GetEntry(
		entryID uint,
	) (
		entry guestbook.Entry,
		err error,
	)
	DeleteEntry(
		entryID uint,
	) (
		err error,
	)
}
