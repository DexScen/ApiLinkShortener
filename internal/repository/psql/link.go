package psql

import (
	"context"
	"database/sql"
	"github.com/DexScen/ApiLinkShortener/internal/domain"
	"time"
)

type Links struct {
	db *sql.DB
}

func NewLinks(db *sql.DB) *Links {
	return &Links{db}
}

func (l *Links) Create(ctx context.Context, link domain.Link) error {
	_, err := l.db.Exec("INSERT INTO links (id, longLink, shortLink, created) values $1 $2 $3 $4",
		link.ID, link.LongLink, link.ShortLink, link.Created)
	return err
}

func (l *Links) GetByShortLink(ctx context.Context, shortLink string) (domain.Link, error) {
	var link domain.Link
	err := l.db.QueryRow("SELECT id, longLink, shortLink, created FROM links where shortlink = $3", shortLink).
		Scan(&link.ID, &link.LongLink, &link.ShortLink, &link.Created)
	if err == sql.ErrNoRows {
		return link, domain.ErrShortLinkNotFound
	}
	return link, err
}

func (l *Links) GetByLongLink(ctx context.Context, longLink string) (domain.Link, bool) {
	var link domain.Link
	err := l.db.QueryRow("SELECT id, longLink, shortLink, created FROM links where longLink = $2", longLink).
		Scan(&link.ID, &link.LongLink, &link.ShortLink, &link.Created)
	if err == sql.ErrNoRows {
		return link, false
	}
	return link, true
}

func (l *Links) Delete(ctx context.Context, time time.Time) error {
	_, err := l.db.Exec("DELETE FROM links WHERE created = $4", time)
	return err
}
