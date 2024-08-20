package flags

import (
	"flag"
	"os"
)

var (
	EndPoint string
	DBHost   string
	DBPort   string
	DBUser   string
	DPPass   string
	DPName   string
)

func ParseFlags() {
	flag.StringVar(&EndPoint, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&EndPoint, "h", "3036", "db port")
	flag.StringVar(&EndPoint, "p", "", "db port")
	flag.StringVar(&EndPoint, "u", "root", "user name")
	flag.StringVar(&EndPoint, "pass", "", "db pass")
	flag.StringVar(&EndPoint, "n", "", "bd name")
	flag.Parse()

	if envServerEndPoint := os.Getenv("ADDRESS"); envServerEndPoint != "" {
		EndPoint = envServerEndPoint
	}
	if envServerEndPoint := os.Getenv("DB_HOST"); envServerEndPoint != "" {
		DBHost = envServerEndPoint
	}
	if envServerEndPoint := os.Getenv("DB_PORT"); envServerEndPoint != "" {
		DBPort = envServerEndPoint
	}
	if envServerEndPoint := os.Getenv("DB_USER"); envServerEndPoint != "" {
		DBUser = envServerEndPoint
	}
	if envServerEndPoint := os.Getenv("DB_PASSWORD"); envServerEndPoint != "" {
		DPPass = envServerEndPoint
	}
	if envServerEndPoint := os.Getenv("DB_NAME"); envServerEndPoint != "" {
		DPName = envServerEndPoint
	}

}
