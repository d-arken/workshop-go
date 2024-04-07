package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	file    string
	gormCfg *gorm.Config
}

func NewSQLiteInMemory() *SQLite {
	return &SQLite{file: "file::memory:?cache=shared", gormCfg: &gorm.Config{}}
}

func (s *SQLite) Open() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(s.file), s.gormCfg)
}
