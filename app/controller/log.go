package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"mota-server/app/dto"
	"mota-server/app/repository"
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

	result := make([]dto.ShortSentenceDto, 0)
	for _, shortSentence := range shortSentences {
		shortSentenceDto := dto.ShortSentenceDto{
			ID:       shortSentence.ID,
			Sentence: shortSentence.Sentence,
		}
		result = append(result, shortSentenceDto)
	}

	return c.JSON(http.StatusOK, result)
}
