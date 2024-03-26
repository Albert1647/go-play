package models

import (
	"natthan.com/go-play/db"
)

type Registration struct {
	ID      int64
	EventId int64 `binding:"required"`
	UserId  int64 `binding:"required"`
}

func GetAllRegistration() ([]Registration, error) {
	query := "SELECT * FROM registrations"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registrations []Registration

	for rows.Next() {
		var regis Registration
		err := rows.Scan(&regis.ID, &regis.EventId, &regis.UserId)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, regis)
	}

	return registrations, nil
}
