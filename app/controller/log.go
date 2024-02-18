package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mota-server/app/domain"
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
	log.Info.Println(fmt.Sprintf("%s %s GetShortSentencePlayLogs", c.Request().Method, c.Request().RequestURI))

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

	return c.JSON(http.StatusOK, logs)
}

func (ctr *LogController) CreateShortSentencePlayLogs(c echo.Context) error {
	log.Info.Println(fmt.Sprintf("%s %s CreateShortSentencePlayLogs", c.Request().Method, c.Request().RequestURI))

	model := domain.ShortSentencePlayLog{}

	err := c.Bind(&model)
	if err != nil {
		log.Error.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	id, err := ctr.shortSentencePlayLogRepo.Insert(&model)
	if err != nil {
		log.Error.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, id)
}
