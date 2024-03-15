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
		ProductSalesId:   int64(req.ProductSalesID),
		Qty:              int64(req.QTY),
		TotalPricing:     int64(req.TotalPricing),
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

func (m *grpcRepositoryCustomer) CheckStok(ctx context.Context, req int32) (err error) {
	conn, err := m.Pool.Get(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := grpcCustomer.NewCustomerUseCaseServiceClient(conn)
	_, err = client.CheckStok(ctx, &grpcCustomer.RequestCheckStok{
		ProductId: req,
	})
	if err != nil {
		return err
	}
	return
}

func (m *grpcRepositoryCustomer) PostCheckoutPartner(ctx context.Context, req domain.RequestDataCheckout) (err error) {
	conn, err := m.Pool.Get(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	pricingStr := strconv.FormatFloat(req.Pricing, 'f', -1, 64)
	discountStr := strconv.FormatFloat(req.Discount, 'f', -1, 64)
	taxStr := strconv.FormatFloat(req.Tax, 'f', -1, 64)
	client := grpcCustomer.NewCustomerUseCaseServiceClient(conn)
	_, err = client.PartnerCheckout(ctx, &grpcCustomer.RequestDataCheckout{
		Email:            req.Email,
		Name:             req.Name,
		PhoneNumber:      req.PhoneNumber,
		ProductSalesId:   req.ProductSalesID,
		Qty:              req.QTY,
		TotalPricing:     req.TotalPricing,
		PaymentReference: req.PaymentReference,
		PaymentDomain:    req.PaymentDomain,
		CustomerId:       req.CustomerID,
		ListKey:          req.ListKey,
		Invoice:          req.Invoice,
		TypeDuration:     req.TypeDuration,
		Pricing:          pricingStr,
		Discount:         discountStr,
		Tax:              taxStr,
		Status:           req.Status,
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

func (m *grpcRepositoryCustomer) GetCheckoutBySerialNumber(ctx context.Context, serialNumber string) (response domain.ResponsetDataCheckout, err error) {
	conn, err := m.Pool.Get(ctx)
	if err != nil {
		return response, err
	}
	defer conn.Close()

	client := grpcCustomer.NewCustomerUseCaseServiceClient(conn)
	responseCheckout, err := client.GetCheckoutByKeyNumber(ctx, &grpcCustomer.RequestProductSerialNumber{SerialNumber: serialNumber})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			err = errors.New("not found")
			return
		}
		return response, err
	}

	response = domain.ResponsetDataCheckout{
		ProductSalesID:   responseCheckout.ProductSalesId,
		QTY:              responseCheckout.Qty,
		TotalPricing:     int64(responseCheckout.TotalPricing),
		PaymentReference: responseCheckout.PaymentReference,
		PaymentDomain:    responseCheckout.PaymentDomain,
		CustomerID:       responseCheckout.CustomerId,
		ListKey:          responseCheckout.ListKey,
		Invoice:          responseCheckout.Invoice,
		TypeDuration:     responseCheckout.TypeDuration,
		Pricing:          response.Pricing,
		Discount:         response.Discount,
		Tax:              response.Tax,
	}
	return
}
