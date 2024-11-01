package service

import (
	"context"

	"github.com/DexScen/ApiLinkShortener/internal/domain"
)

type LinksRepository interface {
	Create(ctx context.Context, link domain.Link) error
	GetByShortLink(ctx context.Context, shortLink string) (domain.Link, error)
	GetByLongLink(ctx context.Context, longLink string) (domain.Link, error)
}

type Links struct {
	repo LinksRepository
}

func NewLinks(repo LinksRepository) *Links {
	return &Links{
		repo: repo,
	}
}

func (l *Links) GetByShortLink(ctx context.Context, shortLink string) (domain.Link, error) {
	return l.repo.GetByShortLink(ctx, shortLink)
}

func (l *Links) GetByLongLink(ctx context.Context, longLink string) (domain.Link, error) {
	return l.repo.GetByLongLink(ctx, longLink)
}
