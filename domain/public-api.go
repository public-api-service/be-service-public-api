package domain

import "context"

type RequestUpdateKey struct {
	ProductID string `json:"product_id"`
	Status    string `json:"status"`
}

type RequestProductIDAndLimit struct {
	ProductID string `json:"product_id"`
	Limit     string `json:"limit"`
}

type GetKeyResponse struct {
	ID         int64  `json:"id"`
	ProductID  int64  `json:"product_id"`
	NumberKeys string `json:"number_keys"`
	Status     string `json:"status"`
}

type ProductResponseDTO struct {
	ID          string  `json:"id,omitempty"`
	ProductID   int64   `json:"product_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	SKU         string  `json:"sku,omitempty"`
	Tipe        string  `json:"tipe,omitempty"`
	Description string  `json:"desc,omitempty"`
	Stok        int     `json:"stok,omitempty"`
	Duration    string  `json:"duration,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Discount    *int    `json:"discount,omitempty"`
	Tax         float64 `json:"tax,omitempty"`
	FinalPrice  float64 `json:"final_price,omitempty"`
	DtmCrt      string  `json:"dtm_crt,omitempty"`
	DtmUpd      string  `json:"dtm_upd,omitempty"`
}

type DetailProduct struct {
	ID         int64   `json:"id"`
	Stok       int     `json:"stok"`
	Duration   string  `json:"duration"`
	Price      float64 `json:"price"`
	Discount   *int    `json:"discount"`
	Tax        int     `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type RequestAdditionalData struct {
	Page         *int    `json:"page"`
	Limit        *int    `json:"limit"`
	ID           *int    `json:"id"`
	FK1          *int    `json:"fk_1"`
	FK2          *string `json:"fk_2"`
	FK3          *string `json:"fk_3"`
	NameSearch   *string `json:"name_search"`
	Order        *string `json:"order"`
	Sort         *string `json:"sort"`
	CustomColumn string  `json:"custom_column"`
}

type GetAllProductResponse struct {
	MetaData MetaData                `json:"meta_data"`
	Data     []AllProductResponseDTO `json:"data"`
}

type MetaData struct {
	TotalData uint   `json:"total_data"`
	TotalPage uint   `json:"total_page"`
	Page      uint   `json:"page"`
	Limit     uint   `json:"limit"`
	Sort      string `json:"sort"`
	Order     string `json:"order"`
}
type AllProductResponseDTO struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	SKU           string          `json:"sku"`
	Tipe          string          `json:"tipe"`
	Description   string          `json:"desc"`
	DetailProduct []DetailProduct `json:"detail_product"`
	DtmCrt        string          `json:"dtm_crt,omitempty"`
	DtmUpd        string          `json:"dtm_upd"`
}

type RequestDataCheckout struct {
	Email            string  `json:"email"`
	Name             string  `json:"name"`
	PhoneNumber      string  `json:"phone_number"`
	ProductSalesID   int64   `json:"product_sales_id"`
	QTY              int64   `json:"qty" form:"qty"`
	TotalPricing     int64   `json:"total_pricing"`
	PaymentReference string  `json:"payment_reference"`
	PaymentDomain    string  `json:"payment_domain"`
	CustomerID       int64   `json:"customer_id"`
	ListKey          string  `json:"list_key"`
	Invoice          string  `json:"invoice"`
	TypeDuration     string  `json:"type_duration"`
	Pricing          float64 `json:"pricing"`
	Discount         float64 `json:"discount"`
	Tax              float64 `json:"tax"`
	Status           string  `json:"status"`
	References       string  `json:"references"`
	IP               string  `json:"ip"`
}

type ResponsetDataCheckout struct {
	Email            string  `json:"email"`
	Name             string  `json:"name"`
	PhoneNumber      string  `json:"phone_number"`
	ProductSalesID   int64   `json:"product_sales_id"`
	QTY              int64   `json:"qty" form:"qty"`
	TotalPricing     int64   `json:"total_pricing"`
	PaymentReference string  `json:"payment_reference"`
	PaymentDomain    string  `json:"payment_domain"`
	CustomerID       int64   `json:"customer_id"`
	ListKey          string  `json:"list_key"`
	Invoice          string  `json:"invoice"`
	TypeDuration     string  `json:"type_duration"`
	Pricing          float64 `json:"pricing"`
	Discount         float64 `json:"discount"`
	Tax              float64 `json:"tax"`
	Status           string  `json:"status"`
	IP               string  `json:"ip"`
}

type RequestDataCustomer struct {
	Email       string `json:"email" form:"email"`
	Name        string `json:"name" form:"name"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
}

type JsonRequest struct {
	Header      Header      `json:"header"`
	Transaction Transaction `json:"transaction"`
}

type Header struct {
	Signature string     `json:"signature"`
	Details   DetailInfo `json:"details"`
}

type DetailInfo struct {
	ProductCategoryCode string `json:"productCategoryCode"`
	SpecVersion         string `json:"specVersion"`
}

type Transaction struct {
	PrimaryAccountNumber     string           `json:"primaryAccountNumber"`
	ProcessingCode           string           `json:"processingCode"`
	TransactionAmount        string           `json:"transactionAmount"`
	TransmissioDateTime      string           `json:"transmissionDateTime"`
	SystemTraceAuditNumber   string           `json:"systemTraceAuditNumber"`
	LocalTransactionTime     string           `json:"localTransactionTime"`
	LocalTransactionDate     string           `json:"localTransactionDate"`
	MerchantCategoryCode     string           `json:"merchantCategoryCode"`
	PointOfServiceEntryMode  string           `json:"pointOfServiceEntryMode"`
	AcquiringInstitutionID   string           `json:"acquiringInstitutionIdentifier"`
	RetrievalReferenceNumber string           `json:"retrievalReferenceNumber"`
	MerchantTerminalID       string           `json:"merchantTerminalId"`
	MerchantID               string           `json:"merchantIdentifier"`
	MerchantLocation         string           `json:"merchantLocation"`
	TransactionCurrencyCode  string           `json:"transactionCurrencyCode"`
	AdditionalTxnFields      AdditionalFields `json:"additionalTxnFields"`
}

type AdditionalFields struct {
	ProductId                     string `json:"productId"`
	TransactionUniqueId           string `json:"transactionUniqueId"`
	CorrelatedTransactionUniqueId string `json:"correlatedTransactionUniqueId"`
	BalanceAmount                 string `json:"balanceAmount"`
	RedemptionPin                 string `json:"redemptionPin"`
	RedemptionAccountNumber       string `json:"redemptionAccountNumber"`
	ActivationAccountNumber       string `json:"activationAccountNumber"`
	ExpiryDate                    string `json:"expiryDate"`
}

type TransactionDTO struct {
	ID                             int64  `json:"id"`
	Signature                      string `json:"signature"`
	ProductCategoryCode            string `json:"productCategoryCode"`
	SpecVersion                    string `json:"specVersion"`
	PrimaryAccountNumber           string `json:"primaryAccountNumber"`
	ProcessingCode                 string `json:"processingCode"`
	TransactionAmount              string `json:"transactionAmount"`
	TransmissionDateTime           string `json:"transmissionDateTime"`
	SystemTraceAuditNumber         string `json:"systemTraceAuditNumber"`
	LocalTransactionTime           string `json:"localTransactionTime"`
	LocalTransactionDate           string `json:"localTransactionDate"`
	MerchantCategoryCode           string `json:"merchantCategoryCode"`
	PointOfServiceEntryMode        string `json:"pointOfServiceEntryMode"`
	AcquiringInstitutionIdentifier string `json:"acquiringInstitutionIdentifier"`
	RetrievalReferenceNumber       string `json:"retrievalReferenceNumber"`
	MerchantTerminalId             string `json:"merchantTerminalId"`
	MerchantIdentifier             string `json:"merchantIdentifier"`
	MerchantLocation               string `json:"merchantLocation"`
	TransactionCurrencyCode        string `json:"transactionCurrencyCode"`
	ProductID                      string `json:"productId"`
	TransactionUniqueId            string `json:"transactionUniqueId"`
	CorrelatedTransactionUniqueId  string `json:"correlatedTransactionUniqueId"`
	BalanceAmount                  string `json:"balanceAmount"`
	RedemptionAccountNumber        string `json:"redemptionAccountNumber"`
	ActivationAccountNumber        string `json:"activationAccountNumber"`
	ExpiryDate                     string `json:"expiryDate"`
	Status                         string `json:"status"`
	IP                             string `json:"ip"`
}

type RequestMarshal struct {
	Header struct {
		Signature string `json:"signature"`
		Details   struct {
			ProductCategoryCode string `json:"productCategoryCode"`
			SpecVersion         string `json:"specVersion"`
		} `json:"details"`
	} `json:"header"`
	Transaction struct {
		PrimaryAccountNumber           string `json:"primaryAccountNumber"`
		ProcessingCode                 string `json:"processingCode"`
		TransactionAmount              string `json:"transactionAmount"`
		TransmissionDateTime           string `json:"transmissionDateTime"`
		SystemTraceAuditNumber         string `json:"systemTraceAuditNumber"`
		LocalTransactionTime           string `json:"localTransactionTime"`
		LocalTransactionDate           string `json:"localTransactionDate"`
		MerchantCategoryCode           string `json:"merchantCategoryCode"`
		PointOfServiceEntryMode        string `json:"pointOfServiceEntryMode"`
		AcquiringInstitutionIdentifier string `json:"acquiringInstitutionIdentifier"`
		RetrievalReferenceNumber       string `json:"retrievalReferenceNumber"`
		MerchantTerminalID             string `json:"merchantTerminalId"`
		MerchantIdentifier             string `json:"merchantIdentifier"`
		MerchantLocation               string `json:"merchantLocation"`
		TransactionCurrencyCode        string `json:"transactionCurrencyCode"`
		AdditionalTxnFields            struct {
			ProductID                     string `json:"productId"`
			TransactionUniqueID           string `json:"transactionUniqueId"`
			CorrelatedTransactionUniqueID string `json:"correlatedTransactionUniqueId"`
		} `json:"additionalTxnFields"`
	} `json:"transaction"`
}

type DigitalAccountReverseParam struct {
	ProcessingCode                 string `json:"processingCode"`
	TransactionAmount              string `json:"transactionAmount"`
	LocalTransactionTime           string `json:"localTransactionTime"`
	LocalTransactionDate           string `json:"localTransactionDate"`
	AcquiringInstitutionIdentifier string `json:"acquiringInstitutionIdentifier"`
	RetrievalReferenceNumber       string `json:"retrievalReferenceNumber"`
	MerchantTerminalID             string `json:"merchantTerminalId"`
	MerchantIdentifier             string `json:"merchantIdentifier"`
}

type LogRequest struct {
	SystemTraceAuditNumber string                 `json:"systemTraceAuditNumber"`
	Request                map[string]interface{} `json:"request"`
	Response               map[string]interface{} `json:"response"`
	Section                string                 `json:"section"`
	Status                 string                 `json:"status"`
}

type PublicAPIUseCase interface {
	PostCheckout(ctx context.Context, request RequestDataCheckout) (err error)
	GetAllProduct(ctx context.Context, request RequestAdditionalData) (response GetAllProductResponse, err error)
	GetProduct(ctx context.Context, request int) (response ProductResponseDTO, err error)
	CheckStok(ctx context.Context, id int32) (err error)
	GetDataMerchantExist(ctx context.Context, merchantID string) (err error)
	AccountRequest(ctx context.Context, request TransactionDTO) (response AdditionalFields, err error)
	AccountReverse(ctx context.Context, request TransactionDTO) (response AdditionalFields, err error)
	InsertLog(ctx context.Context, request LogRequest) (err error)
}

type PublicAPIMySQLRepo interface {
	InsertOriginalTransaction(ctx context.Context, request TransactionDTO) (err error)
	GetDataMerchantExist(ctx context.Context, merchantID string) (err error)
	IsExistReversalAccount(ctx context.Context, request string) (err error)
	GetDataDigitalAccountRequest(ctx context.Context, primaryAccountNumber string) (response TransactionDTO, err error)
	GetDataDigitalAccountRequestByParam(ctx context.Context, request DigitalAccountReverseParam) (response TransactionDTO, err error)
	LastTransaction(ctx context.Context) (lastID int64, err error)
	InsertLog(ctx context.Context, request LogRequest) (err error)
}

type PublicAPIGRPCRepo interface {
	// GetAllProduct(ctx context.Context, request RequestAdditionalData) (response GetAllProductResponse, err error)
}

type ProductGRPCRepo interface {
	UpdateListKeyStatusProduct(ctx context.Context, request RequestUpdateKey) (response string, err error)
	UpdatedStatusDynamicByKeyNumber(ctx context.Context, request RequestUpdateKey) (response string, err error)
	GetListKeyProductByProductIDAndLimit(ctx context.Context, request RequestProductIDAndLimit) (response []GetKeyResponse, err error)
	GetProductByID(ctx context.Context, request int64) (response ProductResponseDTO, err error)
	GetAllProduct(ctx context.Context, request RequestAdditionalData) (response GetAllProductResponse, err error)
}

type CustomerGRPCRepo interface {
	PostCheckoutPartner(ctx context.Context, request RequestDataCheckout) (err error)
	PostCheckout(ctx context.Context, request RequestDataCheckout) (err error)
	CheckStok(ctx context.Context, id int32) (err error)
	GetCheckoutBySerialNumber(ctx context.Context, serialNumber string) (response ResponsetDataCheckout, err error)

	// PostCustomer(ctx context.Context, request RequestDataCustomer) (err error)
}
