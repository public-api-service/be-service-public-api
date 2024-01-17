package domain

type PublicAPIUseCase interface {
	// GetAllProduct(ctx context.Context, request RequestAdditionalData) (response GetAllProductResponse, err error)
}

type PublicAPIMySQLRepo interface {
	// GetAllProduct(ctx context.Context, request RequestAdditionalData) (response []ProductResponseDTO, err error)
	// DynamicCountTable(ctx context.Context, request RequestAdditionalData, tableName string) (response MetaData, err error)
}

type PublicAPIGRPCRepo interface {
	// GetAllProduct(ctx context.Context, request RequestAdditionalData) (response GetAllProductResponse, err error)
}
