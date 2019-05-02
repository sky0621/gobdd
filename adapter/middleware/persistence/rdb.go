package persistence

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func NewRDBMiddleware(dataSource string) (RDBMiddleware, error) {
	// データベース接続
	dbConn, err := gorm.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	if dbConn == nil {
		return nil, errors.New("can not connect to Cloud SQL")
	}
	dbConn.LogMode(true)
	if err := dbConn.DB().Ping(); err != nil {
		return nil, err
	}

	return &rdbMiddleware{dbConn: dbConn}, nil
}

type RDBMiddleware interface {
	Close() error
}

type rdbMiddleware struct {
	dbConn *gorm.DB
}

func (p *rdbMiddleware) Close() error {
	if p == nil {
		return nil
	}
	if p.dbConn == nil {
		return nil
	}
	return p.dbConn.Close()
}
