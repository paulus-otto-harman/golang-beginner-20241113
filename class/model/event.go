package model

type Event struct {
	Id            int         `json:"id"`
	TourAt        string      `json:"tour_at"`
	DestinationId int         `json:"destination_id"`
	Destination   Destination `json:"destination"`
}
