package router

import (
	"mota-server/controller"
	"mota-server/repository"
)

func (r *Router) initLogRouter() {
	shortSentenceRepo := repository.NewShortSentenceRepository(r.db)
	ctr := controller.NewLogController(r.db, shortSentenceRepo)

	r.base.GET("/short_sentences", ctr.GetShortSentences)
}
