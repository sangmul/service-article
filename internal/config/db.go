package config

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() (*sql.DB, error) {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// ambil env
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	tlsName := os.Getenv("DB_TLS")

	// 🔐 load CA cert
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile("certs/ca.pem")
	if err != nil {
		log.Fatal(err)
	}

	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append CA cert")
	}

	mysql.RegisterTLSConfig(tlsName, &tls.Config{
		RootCAs: rootCertPool,
	})

	// build DSN
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=%s",
		user, pass, host, port, dbname, tlsName,
	)

	return sql.Open("mysql", dsn)
}
