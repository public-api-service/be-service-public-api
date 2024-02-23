package mysql

import (
	"be-service-public-api/domain"
	"context"
	"database/sql"
	"errors"
)

type mysqlPublicAPIRepository struct {
	Conn *sql.DB
}

func NewMySQLPublicAPIRepository(Conn *sql.DB) domain.PublicAPIMySQLRepo {
	return &mysqlPublicAPIRepository{Conn}
}

func (db *mysqlPublicAPIRepository) InsertOriginalTransaction(ctx context.Context, request domain.TransactionRequest) (err error) {
	// Prepare the SQL query
	query := `
		INSERT INTO transactions (
			signature,
			productCategoryCode,
			specVersion,
			primaryAccountNumber,
			processingCode,
			transactionAmount,
			transmissionDateTime,
			systemTraceAuditNumber,
			localTransactionTime,
			localTransactionDate,
			merchantCategoryCode,
			pointOfServiceEntryMode,
			acquiringInstitutionIdentifier,
			retrievalReferenceNumber,
			merchantTerminalId,
			merchantIdentifier,
			merchantLocation,
			transactionCurrencyCode,
			productID,
			transactionUniqueId,
			correlatedTransactionUniqueId,
			balanceAmount,
			redemptionAccountNumber,
			ActivationAccountNumber,
			expiryDate,
			status
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	// Execute the SQL query with the request data
	_, err = db.Conn.ExecContext(
		ctx,
		query,
		request.Signature,
		request.ProductCategoryCode,
		request.SpecVersion,
		request.PrimaryAccountNumber,
		request.ProcessingCode,
		request.TransactionAmount,
		request.TransmissionDateTime,
		request.SystemTraceAuditNumber,
		request.LocalTransactionTime,
		request.LocalTransactionDate,
		request.MerchantCategoryCode,
		request.PointOfServiceEntryMode,
		request.AcquiringInstitutionIdentifier,
		request.RetrievalReferenceNumber,
		request.MerchantTerminalId,
		request.MerchantIdentifier,
		request.MerchantLocation,
		request.TransactionCurrencyCode,
		request.ProductID,
		request.TransactionUniqueId,
		request.CorrelatedTransactionUniqueId,
		request.BalanceAmount,
		request.RedemptionAccountNumber,
		request.ActivationAccountNumber,
		request.ExpiryDate,
		request.Status,
	)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (db *mysqlPublicAPIRepository) GetDataMerchantExist(ctx context.Context, merchantID string) (err error) {
	var count int
	query := `SELECT COUNT(id) FROM oauth WHERE client_id = ?`

	err = db.Conn.QueryRowContext(ctx, query, merchantID).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("Data not exist")
			return
		}
		return
	}

	if count == 0 {
		err = errors.New("Merchant not exist")
		return err
	}

	return
}
