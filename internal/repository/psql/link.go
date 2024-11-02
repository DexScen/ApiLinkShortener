package psql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DexScen/ApiLinkShortener/internal/domain"
	"github.com/DexScen/ApiLinkShortener/internal/pkg"
)

type Links struct {
	db *sql.DB
}

func NewLinks(db *sql.DB) *Links {
	return &Links{db}
}

func (l *Links) Create(link domain.Link) error {
	_, err := l.db.Exec("INSERT INTO links (id, longLink, shortLink, created) VALUES ($1, $2, $3, $4)",
		link.ID, link.LongLink, link.ShortLink, link.Created)
	return err
}

func (l *Links) GetByShortLink(ctx context.Context, shortLink *domain.Link) error {
	var str string = (*shortLink).ShortLink[len((*shortLink).ShortLink)-5:]
	err := l.db.QueryRow("SELECT id, longLink, shortLink, created FROM links WHERE shortlink = $1", str).
		Scan(&shortLink.ID, &shortLink.LongLink, &shortLink.ShortLink, &shortLink.Created)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.ErrShortLinkNotFound
	}
	(*shortLink).ShortLink = "https://www.shor.ty/" + (*shortLink).ShortLink
	return err
}

func (l *Links) GetByLongLink(ctx context.Context, longLink *domain.Link) error {
	err := l.db.QueryRow("SELECT id, longLink, shortLink, created FROM links WHERE longLink = $1", (*longLink).LongLink).
		Scan(&longLink.ID, &longLink.LongLink, &longLink.ShortLink, &longLink.Created)
	if errors.Is(err, sql.ErrNoRows) {
		var lastLink domain.Link
		var newString string
		lastLink, err = l.GetLast()
		(*longLink).Created = time.Now()
		if err == nil {
			(*longLink).ID = lastLink.ID + 1
			newString, err = pkg.Increment(lastLink.ShortLink)
			if err == nil {
				(*longLink).ShortLink = newString
			}
		} else {
			(*longLink).ID = 0
			(*longLink).ShortLink = "00000"
		}
		err = l.Create(*longLink)
	}
	(*longLink).ShortLink = "https://www.shor.ty/" + (*longLink).ShortLink
	return err
}

func (l *Links) GetLast() (domain.Link, error) {
	var lastLink domain.Link
	err := l.db.QueryRow("SELECT * FROM links ORDER BY id DESC LIMIT 1").
		Scan(&lastLink.ID, &lastLink.LongLink, &lastLink.ShortLink, &lastLink.Created)
	return lastLink, err
}

func (l *Links) Delete(ctx context.Context, time time.Time) error {
	_, err := l.db.Exec("DELETE FROM links WHERE created = $1", time)
	return err
}
