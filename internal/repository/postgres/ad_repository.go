package postgresRepository

import (
	"advertisement-rest-api-http-service/internal/model"
	"advertisement-rest-api-http-service/internal/repository"
	postgres "advertisement-rest-api-http-service/pkg/postgres"
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"

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
	log.Println("page", page, "sort", sort, "order", order)

	query := ""
	if order == "asc" {
		query = fmt.Sprintf("SELECT * FROM ads ORDER BY %s ASC OFFSET $1 LIMIT 10", sort)
	} else if order == "desc" {
		query = fmt.Sprintf("SELECT * FROM ads ORDER BY %s DESC OFFSET $1 LIMIT 10", sort)
	}

	rows, err := r.db.Query(ctx, query, (page-1)*10)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ad model.Ad
		err := rows.Scan(&ad.ID, &ad.Name, &ad.Description, &ad.Price, &ad.CreatedAt, &ad.UpdatedAt)
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

func (r *adRepositoryPostgres) GetAdByID(ctx context.Context, id_ad string, fields bool) (*model.Ad, error) {
	var ad model.Ad

	if !fields {
		err := r.db.QueryRow(ctx, "SELECT name, price FROM ads WHERE id_ad = $1", id_ad).Scan(&ad.Name, &ad.Price)
		if err != nil {
			return nil, err
		}
	} else {
		var description sql.NullString
		err := r.db.QueryRow(ctx, "SELECT name, description, price FROM ads WHERE id_ad = $1", id_ad).Scan(&ad.Name, &description, &ad.Price)
		if err != nil {
			return nil, err
		}
		ad.Description = description.String
	}

	return &ad, nil
}

func (r *adRepositoryPostgres) CreateAd(ctx context.Context, ad *model.Ad) (string, error) {
	id_ad := GenerateShortID()

	_, err := r.db.Exec(ctx, "INSERT INTO ads (id_ad, name, price, description, created_at) VALUES ($1, $2, $3, $4, $5)", id_ad, ad.Name, ad.Price, ad.Description, ad.CreatedAt)
	if err != nil {
		return "", err
	}

	return id_ad, nil
}

// генерю по подобию тех ID что были в firebase
func GenerateShortID() string {
	// as in firebase
	id := uuid.New()

	uuidBytes := id[:]

	base64Str := base64.RawURLEncoding.EncodeToString(uuidBytes)

	return base64Str[:22]
}
