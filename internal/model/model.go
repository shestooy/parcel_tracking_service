package model

import (
	"fmt"
	m "parcel_tracking_service/internal/storage"
	"time"
)

type ParcelService struct {
	store m.ParcelStore
}

func NewParcelService(store m.ParcelStore) ParcelService {
	return ParcelService{store: store}
}

func (s ParcelService) Register(client int, address string) (m.Parcel, error) {
	parcel := m.Parcel{
		Client:    client,
		Status:    m.ParcelStatusRegistered,
		Address:   address,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	id, err := s.store.Add(parcel)
	if err != nil {
		return parcel, err
	}

	parcel.Number = id

	fmt.Printf("Новая посылка № %d на адрес %s от клиента с идентификатором %d зарегистрирована %s\n",
		parcel.Number, parcel.Address, parcel.Client, parcel.CreatedAt)

	return parcel, nil
}

func (s ParcelService) PrintClientParcels(client int) error {
	parcels, err := s.store.GetByClient(client)
	if err != nil {
		return err
	}

	fmt.Printf("Посылки клиента %d:\n", client)
	for _, parcel := range parcels {
		fmt.Printf("Посылка № %d на адрес %s от клиента с идентификатором %d зарегистрирована %s, статус %s\n",
			parcel.Number, parcel.Address, parcel.Client, parcel.CreatedAt, parcel.Status)
	}
	fmt.Println()

	return nil
}

func (s ParcelService) NextStatus(number int64) error {
	parcel, err := s.store.Get(number)
	if err != nil {
		return err
	}

	var nextStatus string
	switch parcel.Status {
	case m.ParcelStatusRegistered:
		nextStatus = m.ParcelStatusSent
	case m.ParcelStatusSent:
		nextStatus = m.ParcelStatusDelivered
	case m.ParcelStatusDelivered:
		return nil
	}

	fmt.Printf("У посылки № %d новый статус: %s\n", number, nextStatus)

	return s.store.SetStatus(number, nextStatus)
}

func (s ParcelService) ChangeAddress(number int64, address string) error {
	return s.store.SetAddress(number, address)
}

func (s ParcelService) Delete(number int64) error {
	return s.store.Delete(number)
}

func (s ParcelService) Close() error {
	return s.store.Close()
}
