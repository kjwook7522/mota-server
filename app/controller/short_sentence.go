package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mota-server/app/dto"
	"mota-server/app/repository"
	"mota-server/log"
	"net/http"
)

type ShortSentenceController struct {
	shortSentenceRepo *repository.ShortSentenceRepository
}

func NewShortSentenceController(shortSentenceRepo *repository.ShortSentenceRepository) *ShortSentenceController {
	return &ShortSentenceController{shortSentenceRepo: shortSentenceRepo}
}

func (ctr *ShortSentenceController) GetShortSentences(c echo.Context) error {
	const (
		DefaultLimit  = 10
		DefaultOffset = 0
	)
	log.Info.Println(fmt.Sprintf("%s %s GetShortSentences", c.Request().Method, c.Request().RequestURI))

	pagination := Pagination{Limit: DefaultLimit, Offset: DefaultOffset}
	err := c.Bind(&pagination)
	if err != nil {
		log.Error.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	sentences, err := ctr.shortSentenceRepo.FindAll(pagination.Limit, pagination.Offset)
	if err != nil {
		log.Error.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	result := make([]dto.ShortSentenceDto, 0)
	for _, sentence := range sentences {
		shortSentenceDto := dto.ShortSentenceDto{
			ID:       sentence.ID,
			Sentence: sentence.Sentence,
		}
		result = append(result, shortSentenceDto)
	}

	return c.JSON(http.StatusOK, result)
}
