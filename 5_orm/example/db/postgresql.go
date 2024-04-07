package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type PostgresSQL struct {
	host    string
	db      string
	user    string
	passwd  string
	port    string
	tz      string
	gormCgf *gorm.Config
}

func NewPostgresSQL(gormCgf *gorm.Config) *PostgresSQL {
	return &PostgresSQL{host: os.Getenv("DB_HOST"), db: os.Getenv("DB_NAME"), user: os.Getenv("DB_USER"), passwd: os.Getenv("DB_PASSWD"), port: os.Getenv("DB_PORT"), tz: os.Getenv("DB_TZ"), gormCgf: gormCgf}
}

func (p *PostgresSQL) Open() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", p.host, p.user, p.passwd, p.db, p.port, p.tz)), p.gormCgf)
}
