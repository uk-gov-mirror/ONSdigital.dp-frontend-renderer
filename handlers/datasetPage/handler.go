package datasetPage

import (
	"github.com/ONSdigital/dp-frontend-models/model/datasetPage"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"net/http"
)

//Handler builds the template for dataset page
func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page datasetPage.Page

		render.Handler(w, req, &page, &page.Page, "datasetPage", nil, cfg)
	}
}