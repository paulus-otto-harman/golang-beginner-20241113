package model

type DataPage struct {
	Success    bool        `json:"success"`
	Page       int         `json:"page,omitempty"`
	Limit      int         `json:"limit,omitempty"`
	TotalItems int         `json:"total_items,omitempty"`
	TotalPages int         `json:"total_pages,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
