package service

import (
	"context"
	"github.com/DexScen/ApiLinkShortener/internal/domain"
	"time"
)

type LinksRepository interface {
	Create(ctx context.Context, link domain.Link) error
	GetByShortLink(ctx context.Context, shortLink string) (domain.Link, error)
	GetByLongLink(ctx context.Context, longLink string) (domain.Link, bool)
}

type Links struct {
	repo LinksRepository
}

func NewLinks(repo LinksRepository) *Links {
	return &Links{
		repo: repo,
	}
}

func (l *Links) Create(ctx context.Context, link domain.Link) error {
	if link.Created.IsZero() {
		link.Created = time.Now()
	}

	return l.repo.Create(ctx, link)
}

func (l *Links) GetByShortLink(ctx context.Context, shortLink string) (domain.Link, error) {
	return l.repo.GetByShortLink(ctx, shortLink)
}

func (l *Links) GetByLongLink(ctx context.Context, longLink string) (domain.Link, bool) {
	return l.repo.GetByLongLink(ctx, longLink)
}
