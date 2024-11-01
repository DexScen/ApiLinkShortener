package domain

import (
	"errors"
	"time"
)

var (
	ErrShortLinkNotFound = errors.New("short link not present in db")
	ErrOverflow          = errors.New("short link amount exceeded, overflow")
)

type Link struct {
	ID        int64     `json:"id"`
	LongLink  string    `json:"longlink"`
	ShortLink string    `json:"shortlink"`
	Created   time.Time `json:"created"`
}
