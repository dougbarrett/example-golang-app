package server

import (
	"net/http"

	"github.com/dougbarrett/example-golang-app"
	"github.com/labstack/echo"
)

func (s *server) getAllEntries(c echo.Context) error {

	var retData struct {
		Entries []guestbook.Entry
	}
	var err error

	retData.Entries, err = s.handler.ListEntries()

	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "entries_all.html", retData)
}

func (s *server) newEntry(c echo.Context) error {
	return c.Render(http.StatusOK, "entries_new.html", nil)
}

type createEntryData struct {
	guestbook.Entry
}

func (s *server) createEntry(c echo.Context) error {
	e := new(createEntryData)
	if err := c.Bind(e); err != nil {
		return err
	}
	_, err := s.handler.CreateEntry(
		e.Title,
		e.Name,
		e.Email,
		e.Message,
	)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (s *server) deleteEntry(c echo.Context) error {
	e := new(createEntryData)
	if err := c.Bind(e); err != nil {
		return err
	}

	err := s.handler.DeleteEntry(e.ID)

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
