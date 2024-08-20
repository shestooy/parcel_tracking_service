package app

import (
	"net/http"
	f "parcel_tracking_service/internal/flags"
	r "parcel_tracking_service/internal/httpserver"
)

func Start() error {
	f.ParseFlags()
	return http.ListenAndServe(f.EndPoint, r.GetRouter())
}
