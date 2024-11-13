package model

type Event struct {
	Id            int         `json:"id"`
	TourAt        string      `json:"tour_at"`
	DestinationId int         `json:"-"`
	Destination   Destination `json:"destination"`
}
