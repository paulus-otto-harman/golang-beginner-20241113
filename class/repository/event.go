package repository

import (
	"20241113/class/model"
	"database/sql"
	"math"
)

type Event struct {
	Db *sql.DB
}

func InitEventRepo(db *sql.DB) *Event {
	return &Event{Db: db}
}

func (repo *Event) All(date string, page int, sort string) (int, int, []model.Event, error) {
	startDate := "1990-01-01"
	if date != "" {
		startDate = date
	}

	var count int
	queryCount := `SELECT COUNT(*) FROM events WHERE tour_datetime >= $1 AND events.tour_datetime > NOW()`
	err := repo.Db.QueryRow(queryCount, startDate).Scan(&count)

	orderBy := ""
	if sort != "" {
		orderBy = "ORDER BY price ASC"
	}

	if sort == "high_to_low" {
		orderBy = "ORDER BY price DESC"
	}

	query := `SELECT events.id, tour_datetime, destination_id,destinations.name, destinations.thumbnail, destinations.price
				FROM events
				JOIN destinations ON events.destination_id = destinations.id
				WHERE events.tour_datetime >= $2 AND events.tour_datetime > NOW()
				` + orderBy + ` LIMIT 6 OFFSET $1`

	offset := (page - 1) * 6

	rows, err := repo.Db.Query(query, offset, startDate)

	var events []model.Event
	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.Id, &event.TourAt, &event.DestinationId, &event.Destination.Name, &event.Destination.Thumbnail, &event.Destination.Price); err != nil {
			return count, int(math.Ceil(float64(count) / 6)), []model.Event{}, err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return count, int(math.Ceil(float64(count) / 6)), []model.Event{}, err
	}
	return count, int(math.Ceil(float64(count) / 6)), events, nil

}
