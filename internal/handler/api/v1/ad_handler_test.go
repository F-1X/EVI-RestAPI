package v1_test

import (
	v1 "advertisement-rest-api-http-service/internal/handler/api/v1"
	"advertisement-rest-api-http-service/internal/model"
	service "advertisement-rest-api-http-service/internal/service/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getAdsResponseBody(ads ...*model.Ad) string {

	respBody, err := json.Marshal(ads)
	if err != nil {
		return ""
	}
	return string(respBody)
}

func TestGetAds(t *testing.T) {

	gin.SetMode(gin.TestMode)

	mockResp1 := &model.Ad{
		Name:        "1",
		Description: "1",
		Price:       100,
	}

	mockResp2 := &model.Ad{
		Name:        "2",
		Description: "2",
		Price:       1000,
	}

	mockResp3 := &model.Ad{
		Name:        "3",
		Description: "3",
		Price:       10000,
	}

	type query struct {
		page  string
		sort  string
		order string
	}

	testCases := []struct {
		name           string
		description    string
		query          query
		input          []*model.Ad
		expectedOutput string
		expectedCode   int
	}{
		{
			"no query",
			"success",
			query{},
			[]*model.Ad{mockResp1, mockResp2, mockResp3},
			getAdsResponseBody(mockResp1, mockResp2, mockResp3),
			200,
		},
		{
			"bad page type",
			"bad request",
			query{
				page: "text",
			},
			[]*model.Ad{mockResp1, mockResp2, mockResp3},
			`{"error":"invalid page parameter"}`,
			400,
		},

		{
			"good page 1",
			"page >0",
			query{
				page: "1",
			},
			[]*model.Ad{},
			"[]",
			200,
		},

		{
			"bad page 2",
			"bad request",
			query{
				page: "-1",
			},
			[]*model.Ad{},
			`{"error":"invalid page parameter"}`,
			400,
		},
		{
			"good order 1",
			"validate order",
			query{
				order: "asc",
			},
			[]*model.Ad{},
			"[]",
			200,
		},
		{
			"good order 2",
			"validate order",
			query{
				order: "asc",
			},
			[]*model.Ad{},
			`[]`,
			200,
		},
		{
			"bad order 1",
			"bad request",
			query{
				order: "any",
			},
			[]*model.Ad{},
			`{"error":"invalid order parameter"}`,
			400,
		},
		{
			"",
			"OK",
			query{
				page:  "1",
				order: "asc",
				sort:  "price",
			},
			[]*model.Ad{mockResp3, mockResp2, mockResp1},
			`[{"name":"3","description":"3","price":10000},{"name":"2","description":"2","price":1000},{"name":"1","description":"1","price":100}]`,
			200,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(service.AdServicer)

			mockService.On("GetAds", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tt.input, nil)

			rr := httptest.NewRecorder()
			router := gin.Default()
			api := v1.NewAdHandler(
				mockService,
			)

			BASE_URL := "/api/v1/ads"
			router.GET(BASE_URL, api.GetAds)

			u, err := url.Parse(BASE_URL)
			if err != nil {
				t.Fatalf("failed to parse URL: %v", err)
			}
			q := u.Query()
			if tt.query.page != "" {
				q.Set("page", tt.query.page)
			}
			if tt.query.sort != "" {
				q.Set("sort", tt.query.sort)
			}
			if tt.query.order != "" {
				q.Set("order", tt.query.order)
			}
			u.RawQuery = q.Encode()

			request, err := http.NewRequest(http.MethodGet, u.String(), nil)
			assert.NoError(t, err)

			router.ServeHTTP(rr, request)
			assert.Equal(t, tt.expectedCode, rr.Code, "status code: got (%d) want (%d), testtase: (%s), url: (%s)", rr.Code, tt.expectedCode, tt.name, BASE_URL)
			assert.Equal(t, tt.expectedOutput, rr.Body.String(), "output got: (%s) want (%s), testtase: (%s), url: (%s)", rr.Body.String(), tt.expectedOutput, tt.name, BASE_URL)
		})

	}

}
