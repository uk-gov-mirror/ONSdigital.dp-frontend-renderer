package geographyAreaPage

import "github.com/ONSdigital/dp-frontend-models/model"

// Page contains the data re-used on each page as well as the data for the current page
type Page struct {
	model.Page
	Data GeographyAreaPage `json:"data"`
}

// GeographyAreaPage represents the Geography data on the Geography list page
type GeographyAreaPage struct {
	Items []Items `json:"items"`
}

// Items represents the Type of data of the Geography page
type Items struct {
	Label string `json:"label"`
	Text  string `json:"text"`
	ID    string `json:"id"`
}