package service

import (
	"advertisement-rest-api-http-service/internal/model"
	"advertisement-rest-api-http-service/internal/repository"
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.2 --all
type AdServicer interface {
	GetAdByID(ctx context.Context, id string, includeDescription bool) (*model.Ad, error)
	GetAds(ctx context.Context, page int, sort string, order string) ([]*model.Ad, error)
	CreateAd(ctx context.Context, ad *model.Ad) (string, error)
}

type adService struct {
	repo repository.AdRepository
}

func NewAdService(repo repository.AdRepository) AdServicer {
	return &adService{repo: repo}
}

func (s *adService) GetAdByID(ctx context.Context, id string, includeDescription bool) (*model.Ad, error) {
	return s.repo.GetAdByID(ctx, id, includeDescription)

}

func (s *adService) GetAds(ctx context.Context, page int, sort string, order string) ([]*model.Ad, error) {
	return s.repo.GetAds(ctx, page, sort, order)
}

func (s *adService) CreateAd(ctx context.Context, ad *model.Ad) (string, error) {
	return s.repo.CreateAd(ctx, ad)

}
