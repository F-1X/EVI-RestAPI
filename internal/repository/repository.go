package repository

import (
	"advertisement-rest-api-http-service/internal/model"
	"context"
)
//go:generate go run github.com/vektra/mockery/v2@v2.42.2 --all
type AdRepository interface {
	GetAds(ctx context.Context, page int, sort string, order string) ([]*model.Ad, error)
	GetAdByID(ctx context.Context, id string, fields bool) (*model.Ad, error)
	CreateAd(ctx context.Context, ad *model.Ad) (string, error)
}
