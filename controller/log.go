package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"mota-server/repository"
	"net/http"
)

type LogController struct {
	db                *gorm.DB
	shortSentenceRepo *repository.ShortSentenceRepository
}

func NewLogController(db *gorm.DB, shortSentenceRepo *repository.ShortSentenceRepository) *LogController {
	return &LogController{db: db, shortSentenceRepo: shortSentenceRepo}
}

func (ctr *LogController) GetShortSentences(c echo.Context) error {
	const (
		DefaultLimit  = 10
		DefaultOffset = 0
	)
	param := PageParam{Limit: DefaultLimit, Offset: DefaultOffset}
	err := c.Bind(&param)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	shortSentences := ctr.shortSentenceRepo.FindAll(param.Limit, param.Offset)
	return c.JSON(http.StatusOK, shortSentences)
}
