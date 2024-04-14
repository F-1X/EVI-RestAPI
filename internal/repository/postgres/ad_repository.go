package postgresRepository

import (
	"advertisement-rest-api-http-service/internal/model"
	"advertisement-rest-api-http-service/internal/repository"
	postgres "advertisement-rest-api-http-service/pkg/postgres"
	"context"
	"encoding/base64"

	"github.com/google/uuid"
)

type adRepositoryPostgres struct {
	db postgres.PostgresDB
}

func NewAdRepositoryPostgres(db postgres.PostgresDB) repository.AdRepository {
	return &adRepositoryPostgres{db: db}
}

func (r *adRepositoryPostgres) GetAds(ctx context.Context, page int, sort string, order string) ([]*model.Ad, error) {
	var ads []*model.Ad

	rows, err := r.db.Query(ctx, "SELECT * FROM ads ORDER BY $1 $2 OFFSET $3 LIMIT 10", sort, order, (page-1)*10)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ad model.Ad
		err := rows.Scan(&ad.ID, &ad.Name, &ad.Description, &ad.CreatedAt)
		if err != nil {
			return nil, err
		}
		ads = append(ads, &ad)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ads, nil
}

func (r *adRepositoryPostgres) GetAdByID(ctx context.Context, id string, fields bool) (*model.Ad, error) {
	var ad model.Ad

	err := r.db.QueryRow(ctx, "SELECT * FROM ads WHERE id = $1", id).Scan(&ad.ID, &ad.Name, &ad.Description, &ad.CreatedAt)
	if err != nil {
		return nil, err
	}

	if !fields {
		ad.Description = ""
	}

	return &ad, nil
}

func (r *adRepositoryPostgres) CreateAd(ctx context.Context, ad *model.Ad) (string, error) {
	id := GenerateShortID()

	_, err := r.db.Exec(ctx, "INSERT INTO ads (id, name, description, created_at) VALUES ($1, $2, $3, $4)", id, ad.Name, ad.Description, ad.CreatedAt)
	if err != nil {
		return "", err
	}

	return id, nil
}

func GenerateShortID() string {
	// as in firebase
	id := uuid.New()

	uuidBytes := id[:]

	base64Str := base64.RawURLEncoding.EncodeToString(uuidBytes)

	return base64Str[:22]
}
