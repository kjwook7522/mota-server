package controller

import (
	"github.com/labstack/echo/v4"
	"mota-server/app/domain"
	"mota-server/app/dto"
	"mota-server/app/repository"
	"mota-server/log"
	"net/http"
)

type LogController struct {
	shortSentencePlayLogRepo *repository.ShortSentencePlayLogRepository
}

func NewLogController(shortSentencePlayLogRepo *repository.ShortSentencePlayLogRepository) *LogController {
	return &LogController{shortSentencePlayLogRepo: shortSentencePlayLogRepo}
}

func (ctr *LogController) GetShortSentencePlayLogs(c echo.Context) error {
	const (
		DefaultLimit  = 10
		DefaultOffset = 0
	)

	pagination := Pagination{Limit: DefaultLimit, Offset: DefaultOffset}
	err := c.Bind(&pagination)
	if err != nil {
		log.Error.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	logs, err := ctr.shortSentencePlayLogRepo.FindAll(pagination.Limit, pagination.Offset)
	if err != nil {
		log.Error.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	result := make([]dto.ShortSentencePlayLogDto, 0)
	for _, l := range logs {
		shortSentencePlayLogDto := dto.ShortSentencePlayLogDto{
			ID:         l.ID,
			IpAddress:  l.IpAddress,
			UserName:   l.UserName,
			SentenceID: l.SentenceID,
			Typing:     l.Typing,
			Speed:      l.Speed,
			CreatedAt:  l.CreatedAt,
		}
		result = append(result, shortSentencePlayLogDto)
	}

	return c.JSON(http.StatusOK, result)
}

func (ctr *LogController) CreateShortSentencePlayLogs(c echo.Context) error {
	requestDto := dto.ShortSentencePlayLogDto{}

	err := c.Bind(&requestDto)
	if err != nil {
		log.Error.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	model := domain.ShortSentencePlayLog{
		ID:         requestDto.ID,
		IpAddress:  requestDto.IpAddress,
		UserName:   requestDto.UserName,
		SentenceID: requestDto.SentenceID,
		Typing:     requestDto.Typing,
		Speed:      requestDto.Speed,
		CreatedAt:  requestDto.CreatedAt,
	}
	id, err := ctr.shortSentencePlayLogRepo.Insert(&model)
	if err != nil {
		log.Error.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, id)
}
