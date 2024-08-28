package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	f "parcel_tracking_service/internal/flags"
	r "parcel_tracking_service/internal/httpserver"
	m "parcel_tracking_service/internal/model"
	s "parcel_tracking_service/internal/storage"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var service m.ParcelService

func Start() error {
	f.ParseFlags()

	time.Sleep(4 * time.Second) // чтобы контейнер БД успел развернуться
	log.Println("Start server on " + f.EndPoint)
	db, err := DBConnect()
	if err != nil {
		return err
	}
	service = m.NewParcelService(s.NewParcelStore(db))

	err = http.ListenAndServe(f.EndPoint, r.GetRouter())
	if err != nil {
		return err
	}
	if err = service.Close(); err != nil {
		return err
	}
	return nil
}

func DBConnect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		f.DBUser, f.DPPass, f.DBHost, f.DBPort, f.DPName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
