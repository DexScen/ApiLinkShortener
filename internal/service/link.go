package service

import (
	"context"
	"github.com/DexScen/ApiLinkShortener/internal/domain"
	"time"
)

type LinksRepository interface {
	GetByShortLink(ctx context.Context, shortLink *domain.Link) error
	GetByLongLink(ctx context.Context, longLink *domain.Link) error
	Delete(ctx context.Context, time time.Time) error
}

type Links struct {
	repo LinksRepository
}

func NewLinks(repo LinksRepository) *Links {
	return &Links{
		repo: repo,
	}
}

func (l *Links) GetByShortLink(ctx context.Context, shortLink *domain.Link) error {
	return l.repo.GetByShortLink(ctx, shortLink)
}

func (l *Links) GetByLongLink(ctx context.Context, longLink *domain.Link) error {
	return l.repo.GetByLongLink(ctx, longLink)
}

func (l *Links) Delete(ctx context.Context, time time.Time) error {
	return l.repo.Delete(ctx, time)
}
