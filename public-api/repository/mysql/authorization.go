package mysql

import (
	"be-service-public-api/domain"
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
)

type mysqlAuthorizationRepository struct {
	Conn *sql.DB
}

func NewMySQLAuthorizationRepository(Conn *sql.DB) domain.AuthorizationMySQLRepo {
	return &mysqlAuthorizationRepository{Conn}
}

func (db *mysqlAuthorizationRepository) GetAllClientData(ctx context.Context) (response []domain.ResponseB2BDTO, err error) {
	query := `SELECT id, client_id, secret_key, role, domain, dtm_crt, dtm_upd FROM oauth`
	rows, err := db.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var listB2B []domain.ResponseB2BDTO
	for rows.Next() {
		var b2bList domain.ResponseB2BDTO
		if err := rows.Scan(
			&b2bList.ID,
			&b2bList.ClientID,
			&b2bList.ClientSecret,
			&b2bList.Role,
			&b2bList.Domain,
			&b2bList.DtmCrt,
			&b2bList.DtmUpd,
		); err != nil {
			log.Error(err)
			return nil, err
		}

		listB2B = append(listB2B, b2bList)
	}

	return listB2B, nil
}
