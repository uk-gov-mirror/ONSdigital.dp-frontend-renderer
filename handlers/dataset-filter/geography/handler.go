package geography

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/geography"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page geography.Page

	render.Handler(w, req, &page, &page.Page, "dataset-filter/geography", nil)
}