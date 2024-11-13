package repository

import (
	"20241113/class/model"
	"database/sql"
	"fmt"
	"math"
)

type Event struct {
	Db *sql.DB
}

func InitEventRepo(db *sql.DB) *Event {
	return &Event{Db: db}
}

func (repo *Event) All(date string, page int, sort string) (int, int, []model.Event, error) {
	const Limit = 6

	startDate := ""
	if date != "" {
		startDate = fmt.Sprintf(" AND events.tour_datetime::date = '%s' ", date)
	}

	var count int
	queryCount := `SELECT COUNT(*) FROM events WHERE events.tour_datetime > NOW() ` + startDate
	err := repo.Db.QueryRow(queryCount).Scan(&count)

	orderBy := ""
	if sort != "" {
		orderBy = "ORDER BY price ASC"
	}

	if sort == "high_to_low" {
		orderBy = "ORDER BY price DESC"
	}

	query := `SELECT events.id, tour_datetime, events.destination_id,destinations.name, destinations.thumbnail, destinations.price, o.sold,o.average
				FROM events
				JOIN destinations ON events.destination_id = destinations.id
				LEFT JOIN (
					SELECT destination_id,COUNT(*) sold,AVG(rating)::decimal(3,2) average FROM orders
					JOIN events ON orders.event_id=events.id
					GROUP BY destination_id
				) o ON events.destination_id=o.destination_id
				WHERE events.tour_datetime > NOW() ` + startDate + orderBy + ` LIMIT 6 OFFSET $1`

	offset := (page - 1) * Limit

	rows, err := repo.Db.Query(query, offset)

	var events []model.Event
	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.Id, &event.TourAt, &event.DestinationId, &event.Destination.Name, &event.Destination.Thumbnail, &event.Destination.Price, &event.Destination.Sold, &event.Destination.Rating); err != nil {
			return count, int(math.Ceil(float64(count) / Limit)), []model.Event{}, err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return count, int(math.Ceil(float64(count) / Limit)), []model.Event{}, err
	}
	return count, int(math.Ceil(float64(count) / Limit)), events, nil

}
