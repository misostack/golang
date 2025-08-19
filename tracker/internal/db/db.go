package db

import (
	"fmt"
	"gogym/tracker/internal/shared"
)

type DBConnection struct {
	WriterDsn string
	ReaderDsn string
}

func NewDBConnection(dsn string) *DBConnection {
	config := shared.NewConfig()
	writerDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	readerDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	return &DBConnection{
		WriterDsn: writerDsn,
		ReaderDsn: readerDsn,
	}
}
