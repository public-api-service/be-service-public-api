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

func (db *mysqlPublicAPIRepository) InsertOriginalTransaction(ctx context.Context, request domain.TransactionDTO) (err error) {
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

func (db *mysqlPublicAPIRepository) IsExistReversalAccount(ctx context.Context, request string) (err error) {
	var count int
	query := `SELECT COUNT(id) FROM transactions WHERE transactionUniqueId = ?`

	err = db.Conn.QueryRowContext(ctx, query, request).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("Data not exist")
		}
	}

	if count > 0 {
		err = errors.New("Duplicate reversal account")
		return err
	}

	return
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

func (db *mysqlPublicAPIRepository) LastTransaction(ctx context.Context) (lastInsertID int64, err error) {
	query := `SELECT MAX(id) FROM transactions`

	err = db.Conn.QueryRowContext(ctx, query).Scan(&lastInsertID)
	if err != nil {
		return 0, nil
	}

	return lastInsertID, nil
}

func (db *mysqlPublicAPIRepository) GetDataDigitalAccountRequest(ctx context.Context, primaryAccountNumber string) (response domain.TransactionDTO, err error) {
	query := `SELECT signature,productCategoryCode,specVersion,primaryAccountNumber,processingCode,transactionAmount,transmissionDateTime,systemTraceAuditNumber,localTransactionTime,localTransactionDate,merchantCategoryCode,pointOfServiceEntryMode,acquiringInstitutionIdentifier,retrievalReferenceNumber,merchantTerminalId,merchantIdentifier,merchantLocation,transactionCurrencyCode,productID,transactionUniqueId,correlatedTransactionUniqueId,balanceAmount,redemptionAccountNumber,ActivationAccountNumber,expiryDate,status FROM transactions WHERE primaryAccountNumber = ?`
	rows, err := db.Conn.QueryContext(ctx, query, primaryAccountNumber)
	if err != nil {
		// Mengembalikan error jika terjadi kesalahan saat menjalankan query
		return response, err
	}
	defer rows.Close()

	err = db.Conn.QueryRowContext(ctx, query, primaryAccountNumber).Scan(
		&response.Signature,
		&response.ProductCategoryCode,
		&response.SpecVersion,
		&response.PrimaryAccountNumber,
		&response.ProcessingCode,
		&response.TransactionAmount,
		&response.TransmissionDateTime,
		&response.SystemTraceAuditNumber,
		&response.LocalTransactionTime,
		&response.LocalTransactionDate,
		&response.MerchantCategoryCode,
		&response.PointOfServiceEntryMode,
		&response.AcquiringInstitutionIdentifier,
		&response.RetrievalReferenceNumber,
		&response.MerchantTerminalId,
		&response.MerchantIdentifier,
		&response.MerchantLocation,
		&response.TransactionCurrencyCode,
		&response.ProductID,
		&response.TransactionUniqueId,
		&response.CorrelatedTransactionUniqueId,
		&response.BalanceAmount,
		&response.RedemptionAccountNumber,
		&response.ActivationAccountNumber,
		&response.ExpiryDate,
		&response.Status,
	)
	if err != nil {
		// Mengembalikan error jika terjadi kesalahan saat menjalankan query
		return response, err
	}

	return response, nil
}
