package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type API struct {
	DB *Database
}

func NewAPI(db *Database) *API {
	return &API{DB: db}
}

func (api *API) SetupRoutes() *gin.Engine {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// Comic routes
	r.GET("/api/comics", api.GetAllComics)
	r.GET("/api/comics/:endpoint", api.GetComicByEndpoint)
	r.GET("/api/search", api.SearchComics)
	r.GET("/api/comics/type/:type", api.GetComicsByType)

	// Stats
	r.GET("/api/stats", api.GetStats)

	return r
}

func (api *API) GetAllComics(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	comics, err := api.DB.GetComicsPaginated(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	total, _ := api.DB.GetTotalComics()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    comics,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func (api *API) GetComicByEndpoint(c *gin.Context) {
	endpoint := c.Param("endpoint")
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	comic, err := api.DB.GetComicByEndpoint(endpoint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Comic not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    comic,
	})
}

func (api *API) SearchComics(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Query parameter 'q' is required",
		})
		return
	}

	comics, err := api.DB.SearchComics(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    comics,
		"count":   len(comics),
	})
}

func (api *API) GetComicsByType(c *gin.Context) {
	comicType := c.Param("type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	comics, err := api.DB.GetComicsByType(comicType, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    comics,
		"type":    comicType,
		"page":    page,
	})
}

func (api *API) GetStats(c *gin.Context) {
	total, _ := api.DB.GetTotalComics()
	types, _ := api.DB.GetComicsByTypeStats()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"total_comics": total,
			"by_type":      types,
		},
	})
}
