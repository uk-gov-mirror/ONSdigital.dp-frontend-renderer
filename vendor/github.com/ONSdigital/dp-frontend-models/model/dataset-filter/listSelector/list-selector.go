package listSelector

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data     ListSelector `json:"data"`
	FilterID string       `json:"job_id"`
}

// ListSelector ...
type ListSelector struct {
	AddFromRange  Link     `json:"add_from_range"`
	SaveAndReturn Link     `json:"save_and_return"`
	Cancel        Link     `json:"cancel"`
	FiltersAmount int      `json:"filters_amount"`
	FiltersAdded  []Filter `json:"filters_added"`
	AddAllInRange Link     `json:"add_all"`
	RemoveAll     Link     `json:"remove_all"`
	RangeData     Range    `json:"range_values"`
}

// Link ...
type Link struct {
	URL   string `json:"url"`
	Label string `json:"label"`
}

// Filter ...
type Filter struct {
	Label     string `json:"label"`
	RemoveURL string `json:"remove_url"`
}

// Range ...
type Range struct {
	URL    string  `json:"url"`
	Values []Value `json:"values"`
}

// Value ...
type Value struct {
	Label      string `json:"label"`
	IsSelected bool   `json:"is_selected"`
}
