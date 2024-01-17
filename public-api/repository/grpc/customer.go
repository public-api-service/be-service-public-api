package grpc

import (
	"be-service-public-api/domain"
	grpcCustomer "be-service-public-api/public-api/repository/grpc/customer"
	"context"
	"errors"
	"strconv"

	grpcpool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcRepositoryCustomer struct {
	Pool *grpcpool.Pool
}

func NewGRPCCustomerRepository(Pool *grpcpool.Pool) domain.CustomerGRPCRepo {
	return &grpcRepositoryCustomer{Pool}
}

func (m *grpcRepositoryCustomer) PostCheckout(ctx context.Context, req domain.RequestDataCheckout) (err error) {
	conn, err := m.Pool.Get(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	pricingStr := strconv.FormatFloat(req.Pricing, 'f', -1, 64)
	discountStr := strconv.FormatFloat(req.Discount, 'f', -1, 64)
	taxStr := strconv.FormatFloat(req.Tax, 'f', -1, 64)
	client := grpcCustomer.NewCustomerUseCaseServiceClient(conn)
	_, err = client.PostCheckout(ctx, &grpcCustomer.RequestDataCheckout{
		Email:            req.Email,
		Name:             req.Name,
		PhoneNumber:      req.PhoneNumber,
		ProductSalesId:   int32(req.ProductSalesID),
		Qty:              int32(req.QTY),
		TotalPricing:     int32(req.TotalPricing),
		PaymentReference: req.PaymentReference,
		PaymentDomain:    req.PaymentDomain,
		CustomerId:       req.CustomerID,
		ListKey:          req.ListKey,
		Invoice:          req.Invoice,
		TypeDuration:     req.TypeDuration,
		Pricing:          pricingStr,
		Discount:         discountStr,
		Tax:              taxStr,
	})

	if err != nil {
		if status.Code(err) == codes.NotFound {
			err = errors.New("not found")
			return
		}
		return err
	}

	return
}
