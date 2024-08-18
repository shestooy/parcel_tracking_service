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
	res, err := s.db.Exec("INSERT INTO parcel (client,status,address,created_at) VALUES (:c, :s, :a, :cr)",
		sql.Named("c", p.Client),
		sql.Named("s", p.Status),
		sql.Named("a", p.Address),
		sql.Named("cr", p.CreatedAt))
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
	row := s.db.QueryRow("SELECT number, client, status,address, created_at FROM parcel WHERE number = :n",
		sql.Named("n", number))

	p := Parcel{}
	err := row.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)

	if err != nil {
		return p, err
	}

	return p, nil
}

func (s *ParcelStore) GetByClient(client int) ([]Parcel, error) {
	row, err := s.db.Query("SELECT number, client, status,address, created_at FROM parcel WHERE client = :c",
		sql.Named("c", client))

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

	return res, nil
}

func (s *ParcelStore) SetStatus(number int64, status string) error {
	_, err := s.db.Exec("UPDATE parcel SET status = :s WHERE number = :n",
		sql.Named("s", status),
		sql.Named("n", number))
	if err != nil {
		return err
	}

	return nil
}

func (s *ParcelStore) SetAddress(number int64, address string) error {
	row := s.db.QueryRow("SELECT status FROM parcel WHERE number = :n",
		sql.Named("n", number))

	var status string
	err := row.Scan(&status)

	if err != nil {
		return err
	}

	if status != "registered" {
		return nil
	}

	_, err = s.db.Exec("UPDATE parcel SET address = :a WHERE number = :n",
		sql.Named("a", address),
		sql.Named("n", number))

	if err != nil {
		return err
	}
	return nil
}

func (s *ParcelStore) Delete(number int64) error {
	row := s.db.QueryRow("SELECT status FROM parcel WHERE number = :n",
		sql.Named("n", number))

	var status string
	err := row.Scan(&status)

	if err != nil {
		return err
	}
	if status != "registered" {
		return nil
	}

	_, err = s.db.Exec("DELETE FROM parcel WHERE number = :n",
		sql.Named("n", number))

	if err != nil {
		return err
	}

	return nil
}
