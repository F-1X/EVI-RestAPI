package firestoreRepository

import (
	"advertisement-rest-api-http-service/internal/model"
	"advertisement-rest-api-http-service/internal/repository"
	firebase "advertisement-rest-api-http-service/pkg/firebase"
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type adRepository struct {
	client firebase.DB
}

func NewAdRepository(client firebase.DB) repository.AdRepository {

	return &adRepository{client: client}
}

func (r *adRepository) GetAds(ctx context.Context, page int, sort string, order string) ([]*model.Ad, error) {
	var ads []*model.Ad
	var query firestore.Query

	if sort == "" {
		sort = "created_at"
	}

	orderBy := firestore.Desc
	if order == "asd" {
		orderBy = firestore.Asc
	}

	query = r.client.Collection("ads").OrderBy(sort, orderBy).Offset((page - 1) * 10).Limit(10)



	iter := query.Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		log.Println("doc", doc)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var ad model.Ad
		if err := doc.DataTo(&ad); err != nil {
			return nil, err
		}

		ads = append(ads, &ad)
	}
	log.Println(ads)

	return ads, nil
}

func (r *adRepository) GetAdByID(ctx context.Context, id string, fields bool) (*model.Ad, error) {

	doc, err := r.client.Collection("ads").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var ad model.Ad

	err = doc.DataTo(&ad)
	if err != nil {
		return nil, err
	}

	if !fields {
		ad.Description = ""
	}

	return &ad, nil
}

func (r *adRepository) CreateAd(ctx context.Context, ad *model.Ad) (string, error) {

	ref, _, err := r.client.Collection("ads").Add(ctx, ad)
	if err != nil {
		return "", err
	}

	return ref.ID, nil
}
