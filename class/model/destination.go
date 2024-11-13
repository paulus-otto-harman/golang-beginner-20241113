package model

type Destination struct {
	Id        int      `json:"-"`
	Name      string   `json:"name"`
	Thumbnail string   `json:"thumbnail"`
	Price     string   `json:"price"`
	Sold      *int     `json:"sold"`
	Rating    *float32 `json:"rating"`
}
