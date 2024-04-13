package handler

import "github.com/gin-gonic/gin"

//go:generate go run github.com/vektra/mockery/v2@v2.42.2 --all
type Handler interface {
	GetAd(c *gin.Context)
	CreateAd(c *gin.Context)
	GetAds(c *gin.Context)
}
