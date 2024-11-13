package model

type Destination struct {
	Id        int    `json:"-"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Price     string `json:"price"`
}
