package v1

import (
	"advertisement-rest-api-http-service/internal/handler"
	"advertisement-rest-api-http-service/internal/model"
	"advertisement-rest-api-http-service/internal/service"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type adHandler struct {
	service service.AdServicer
}

func NewAdHandler(service service.AdServicer) handler.Handler {
	return &adHandler{service: service}
}

// GetAd retrieves an advertisement by ID.
//
//	@Summary		Get Ad
//	@Description	Retrieve an advertisement by ID
//	@Tags			Advertisements
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of the advertisement to retrieve"
//	@Success		200	{object}	object	"{"id": "ffZ0Wo9KIh29G24iNE1t"}"
//	@Failure		500	{object}	object	"{"error": "Internal	server	error"}"
//	@Router			/ad/{id} [get]
func (h *adHandler) GetAd(c *gin.Context) {
	id := c.Param("id")

	if _, ok := c.Request.URL.Query()["fields"]; ok {
		ad, err := h.service.GetAdByID(c.Request.Context(), id, true)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, ad)
		return
	}

	ad, err := h.service.GetAdByID(c.Request.Context(), id, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ad)
}

// CreateAd creates a new advertisement.
//
//	@Summary		Create Ad
//	@Description	Create a new advertisement
//	@Tags			Advertisements
//	@Accept			json
//	@Produce		json
//	@Param			ad	body		model.Ad	true	"Data of the new advertisement"
//	@Success		201	{object}	object		"{"id": "ffZ0Wo9KIh29G24iNE1t"}"
//	@Failure		400	{object}	object		"{"error": "Bad			request"}"
//	@Failure		500	{object}	object		"{"error": "Internal	server	error"}"
//	@Router			/ad [post]
func (h *adHandler) CreateAd(c *gin.Context) {

	var ad model.Ad
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	location, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		panic(err)
	}

	ad.CreatedAt = time.Now().In(location)
	ad.UpdatedAt = time.Now().In(location)

	log.Println(ad)
	id, err := h.service.CreateAd(c.Request.Context(), &ad)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetAds retrieves a list of advertisements.
//
//	@Summary		Get Ads
//	@Description	Retrieve a list of advertisements
//	@Tags			Advertisements
//	@Produce		json
//	@Param			page	query		int		false	"Page number"
//	@Param			sort	query		string	false	"Sorting field (price, created_at, updated_at)"
//	@Param			order	query		string	false	"Sorting order (asc, desc)"
//	@Success		200		{array}		object  
//	@Failure		400		{object}	object	false "{"error": "Bad			request"}" 
//	@Failure		500		{object}	object	"{"error": "Internal	server	error"}"
//	@Failure		400		{object}	object	"{"error": "invalid page parameter"}"


// @Router			/ads [get]
func (h *adHandler) GetAds(c *gin.Context) {
	page := 1
	order := "desc"
	sort := "price"
	var err error
	if _, ok := c.Request.URL.Query()["page"]; ok {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
			return
		}
		if !validatePage(page) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
			return
		}
	}

	if orderStr, ok := c.Request.URL.Query()["order"]; ok {
		if len(orderStr) > 0 && (orderStr[0] != "asc" && orderStr[0] != "desc") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order parameter"})
			return
		}
		order = orderStr[0]
	}

	if sortStr, ok := c.Request.URL.Query()["sort"]; ok {
		if len(sortStr) > 0 && (sortStr[0] != "price" && sortStr[0] != "created_at" && sortStr[0] != "updated_at") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sort parameter"})
			return
		}
		sort = sortStr[0]
	}

	ads, err := h.service.GetAds(c.Request.Context(), page, sort, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ads)
}

// better to place in validate package and test it inplace
func validatePage(page int) bool {
	return page >= 0
}
