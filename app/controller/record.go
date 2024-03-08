package controller

import (
	"github.com/labstack/echo/v4"
	"mota-server/app/dto"
	"mota-server/app/entity"
	"mota-server/app/repository"
	"mota-server/log"
	"net/http"
)

type RecordController struct {
	shortSentenceRecordRepo *repository.ShortSentenceRecordRepository
}

func NewRecordController(shortSentenceRecordRepo *repository.ShortSentenceRecordRepository) *RecordController {
	return &RecordController{shortSentenceRecordRepo: shortSentenceRecordRepo}
}

func (ctr *RecordController) GetShortSentenceRecords(c echo.Context) error {
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

	recordEntities, err := ctr.shortSentenceRecordRepo.FindAll(pagination.Limit, pagination.Offset)
	if err != nil {
		log.Error.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	result := make([]dto.ShortSentenceRecordDto, 0)
	for _, recordEntity := range recordEntities {
		shortSentenceRecordDto := dto.ShortSentenceRecordDto{
			ID:         recordEntity.ID,
			IpAddress:  recordEntity.IpAddress,
			UserName:   recordEntity.UserName,
			SentenceID: recordEntity.SentenceID,
			Typing:     recordEntity.Typing,
			Speed:      recordEntity.Speed,
			CreatedAt:  recordEntity.CreatedAt,
		}
		result = append(result, shortSentenceRecordDto)
	}

	return c.JSON(http.StatusOK, result)
}

func (ctr *RecordController) CreateShortSentenceRecords(c echo.Context) error {
	requestDto := dto.ShortSentenceRecordDto{}

	err := c.Bind(&requestDto)
	if err != nil {
		log.Error.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	recordEntity := entity.ShortSentenceRecordEntity{
		ID:         requestDto.ID,
		IpAddress:  requestDto.IpAddress,
		UserName:   requestDto.UserName,
		SentenceID: requestDto.SentenceID,
		Typing:     requestDto.Typing,
		Speed:      requestDto.Speed,
		CreatedAt:  requestDto.CreatedAt,
	}
	recordId, err := ctr.shortSentenceRecordRepo.Insert(&recordEntity)
	if err != nil {
		log.Error.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, recordId)
}
