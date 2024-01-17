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

type PublicAPIUseCase interface {
	GetAllProduct(ctx context.Context, request RequestAdditionalData) (response GetAllProductResponse, err error)
}

type PublicAPIMySQLRepo interface {
	// GetAllClientData(ctx context.Context) (response []ResponseB2BDTO, err error)
}

type PublicAPIGRPCRepo interface {
	// GetAllProduct(ctx context.Context, request RequestAdditionalData) (response GetAllProductResponse, err error)
}

type ProductGRPCRepo interface {
	UpdateListKeyStatusProduct(ctx context.Context, request RequestUpdateKey) (response string, err error)
	GetListKeyProductByProductIDAndLimit(ctx context.Context, request RequestProductIDAndLimit) (response []GetKeyResponse, err error)
	GetProductByID(ctx context.Context, request int64) (response ProductResponseDTO, err error)
	GetAllProduct(ctx context.Context, request RequestAdditionalData) (response GetAllProductResponse, err error)
}
