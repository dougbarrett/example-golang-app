package mysql

import (
	"github.com/dougbarrett/example-golang-app"
)

func (m *mysql) GetAllEntries() (
	entries []guestbook.Entry,
	err error,
) {
	err = m.db.Find(&entries).Error
	return
}

func (m *mysql) SaveEntry(
	entry *guestbook.Entry,
) (
	err error,
) {
	err = m.db.Save(entry).Error
	return
}

func (m *mysql) GetEntry(
	entryID uint,
) (
	entry guestbook.Entry,
	err error,
) {
	err = m.db.Find(&entry, "id = ?", entryID).Error
	return
}

func (m *mysql) DeleteEntry(
	entryID uint,
) (
	err error,
) {
	return m.db.Delete(guestbook.Entry{}, "id = ?", entryID).Error
}
