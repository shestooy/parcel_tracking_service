package app

import (
	"database/sql"
	m "parcel_tracking_service/internal/model"
	s "parcel_tracking_service/internal/storage"
)

func Start() error {
	db, err := sql.Open("sqlite", "internal/db/tracker.db")
	if err != nil {
		return err
	}
	defer db.Close()
	store := s.NewParcelStore(db)
	service := m.NewParcelService(store)

	client := 1
	address := "Псков, д. Пушкина, ул. Колотушкина, д. 5"
	p, err := service.Register(client, address)
	if err != nil {
		return err
	}

	newAddress := "Саратов, д. Верхние Зори, ул. Козлова, д. 25"
	err = service.ChangeAddress(p.Number, newAddress)
	if err != nil {
		return err
	}

	err = service.NextStatus(p.Number)
	if err != nil {
		return err
	}

	err = service.PrintClientParcels(client)
	if err != nil {
		return err
	}

	err = service.Delete(p.Number)
	if err != nil {
		return err
	}

	err = service.PrintClientParcels(client)
	if err != nil {
		return err
	}

	p, err = service.Register(client, address)
	if err != nil {
		return err
	}

	err = service.Delete(p.Number)
	if err != nil {
		return err
	}

	err = service.PrintClientParcels(client)
	if err != nil {
		return err
	}
	return nil
}
