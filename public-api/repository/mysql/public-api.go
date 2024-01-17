package mysql

import (
	"be-service-public-api/domain"
	"database/sql"
)

type mysqlPublicAPIRepository struct {
	Conn *sql.DB
}

func NewMySQLPublicAPIRepository(Conn *sql.DB) domain.PublicAPIMySQLRepo {
	return &mysqlPublicAPIRepository{Conn}
}
