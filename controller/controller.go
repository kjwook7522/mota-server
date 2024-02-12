package controller

type PageParam struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}
