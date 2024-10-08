package storage

import (
	"database/sql"
)

const (
	ParcelStatusRegistered = "registered"
	ParcelStatusSent       = "sent"
	ParcelStatusDelivered  = "delivered"
)

type Parcel struct {
	Number    int64
	Client    int
	Status    string
	Address   string
	CreatedAt string
}

type ParcelStore struct {
	db *sql.DB
}

func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

func (s *ParcelStore) Add(p Parcel) (int64, error) {
	res, err := s.db.Exec("INSERT INTO parcel (client,status,address,created_at) VALUES (:client, :status, :address, :created_at)",
		sql.Named("client", p.Client),
		sql.Named("status", p.Status),
		sql.Named("address", p.Address),
		sql.Named("created_at", p.CreatedAt))
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *ParcelStore) Get(number int64) (Parcel, error) {
	row := s.db.QueryRow("SELECT number, client, status,address, created_at FROM parcel WHERE number = :number",
		sql.Named("number", number))

	p := Parcel{}
	err := row.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)

	if err != nil {
		return p, err
	}

	return p, nil
}

func (s *ParcelStore) GetByClient(client int) ([]Parcel, error) {
	row, err := s.db.Query("SELECT number, client, status,address, created_at FROM parcel WHERE client = :client",
		sql.Named("client", client))

	if err != nil {
		return nil, err
	}
	defer row.Close()

	var res []Parcel
	for row.Next() {
		p := Parcel{}
		err = row.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)

		if err != nil {
			return nil, err
		}

		res = append(res, p)
	}
	err = row.Err()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ParcelStore) SetStatus(number int64, status string) error {
	_, err := s.db.Exec("UPDATE parcel SET status = :status WHERE number = :number",
		sql.Named("status", status),
		sql.Named("number", number))
	if err != nil {
		return err
	}

	return nil
}

func (s *ParcelStore) SetAddress(number int64, address string) error {

	_, err := s.db.Exec("UPDATE parcel SET address = :address WHERE number = :number AND status = :status",
		sql.Named("address", address),
		sql.Named("number", number),
		sql.Named("status", ParcelStatusRegistered))

	if err != nil {
		return err
	}
	return nil
}

func (s *ParcelStore) Delete(number int64) error {
	_, err := s.db.Exec("DELETE FROM parcel WHERE number = :number AND status = :status",
		sql.Named("number", number),
		sql.Named("status", ParcelStatusRegistered))

	if err != nil {
		return err
	}

	return nil
}
