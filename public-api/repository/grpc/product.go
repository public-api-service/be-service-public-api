package grpc

import (
	"be-service-public-api/domain"
	grpcProduct "be-service-public-api/public-api/repository/grpc/product"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	grpcpool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcRepositoryProduct struct {
	Pool *grpcpool.Pool
}

func NewGRPCProductRepository(Pool *grpcpool.Pool) domain.ProductGRPCRepo {
	return &grpcRepositoryProduct{Pool}
}

func (m *grpcRepositoryProduct) GetListKeyProductByProductIDAndLimit(ctx context.Context, request domain.RequestProductIDAndLimit) (response []domain.GetKeyResponse, err error) {
	conn, err := m.Pool.Get(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := grpcProduct.NewProductServiceClient(conn)
	productGRPC, err := client.GetListKeyProductByProductIDAndLimit(ctx, &grpcProduct.ListKeyProductByProductIDAndLimitServiceRequest{
		IdProduct: request.ProductID,
		Limit:     request.Limit,
	})

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	if productGRPC.Data == nil {
		err = errors.New("Data not found")
		return
	}

	var productResponse []domain.GetKeyResponse
	for _, grpcResp := range productGRPC.Data {
		responseItem := domain.GetKeyResponse{
			ID:         grpcResp.Id,
			ProductID:  grpcResp.IdProduct,
			NumberKeys: grpcResp.NumberKeys,
			Status:     grpcResp.Status,
		}

		productResponse = append(productResponse, responseItem)
	}

	return productResponse, nil
}

func (m *grpcRepositoryProduct) UpdateListKeyStatusProduct(ctx context.Context, request domain.RequestUpdateKey) (response string, err error) {
	conn, err := m.Pool.Get(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := grpcProduct.NewProductServiceClient(conn)
	_, err = client.UpdateListKeyStatusProduct(ctx, &grpcProduct.UpdateListKeyStatusProductServiceRequest{
		ListID: request.ProductID,
		Status: "Purchased",
	})

	if err != nil {
		return "", err
	}
	return
}

func (m *grpcRepositoryProduct) UpdatedStatusDynamicByKeyNumber(ctx context.Context, request domain.RequestUpdateKey) (response string, err error) {
	conn, err := m.Pool.Get(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := grpcProduct.NewProductServiceClient(conn)
	_, err = client.UpdatedStatusDynamicByKeyNumber(ctx, &grpcProduct.UpdateListKeyStatusProductServiceRequest{
		ListID: request.ProductID,
		Status: request.Status,
	})

	if err != nil {
		return "", err
	}
	return
}

func (m *grpcRepositoryProduct) GetProductByID(ctx context.Context, request int64) (response domain.ProductResponseDTO, err error) {
	conn, err := m.Pool.Get(ctx)
	if err != nil {
		return domain.ProductResponseDTO{}, err
	}
	defer conn.Close()

	client := grpcProduct.NewProductServiceClient(conn)
	res, err := client.GetProductByID(ctx, &grpcProduct.ProductIDRequestServiceRequest{
		Id: request,
	})

	if err != nil {
		return domain.ProductResponseDTO{}, err
	}

	idStr := strconv.Itoa(int(res.Id))
	price, err := strconv.ParseFloat(res.Price, 64)
	if err != nil {
		log.Println("Error converting string to float64:", err)
		return
	}

	var discount int
	if res.Discount != "" {
		discountStr, err := strconv.Atoi(res.Discount)
		if err != nil {
			return domain.ProductResponseDTO{}, err
		}
		discount = discountStr
	} else {
		discount = 0
	}

	tax, err := strconv.ParseFloat(res.Tax, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}

	finalPrice, err := strconv.ParseFloat(res.FinalPrice, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}
	response = domain.ProductResponseDTO{
		ID:          idStr,
		ProductID:   res.ProductId,
		Name:        res.Name,
		SKU:         res.Sku,
		Tipe:        res.Tipe,
		Description: res.Description,
		Stok:        int(res.Stok),
		Duration:    res.Duration,
		Price:       price,
		Discount:    &discount,
		Tax:         tax,
		FinalPrice:  finalPrice,
		DtmCrt:      res.DtmCrt,
		DtmUpd:      res.DtmUp,
	}
	return
}

func (m *grpcRepositoryProduct) GetAllProduct(ctx context.Context, request domain.RequestAdditionalData) (response domain.GetAllProductResponse, err error) {
	conn, err := m.Pool.Get(ctx)
	if err != nil {
		return response, err
	}
	defer conn.Close()

	ID := 0
	Page := 0
	Limit := 0
	FK1 := 0
	var FK2, FK3, NameSearch, Order, Sort string

	if request.ID != nil {
		ID = *request.ID
	}

	if request.Page != nil {
		Page = *request.Page
	}

	if request.Limit != nil {
		Limit = *request.Limit
	}

	if request.FK1 != nil {
		FK1 = *request.FK1
	}

	if request.FK2 == nil {
		FK2 = ""
	} else {
		FK2 = *request.FK2
	}

	if request.NameSearch == nil {
		NameSearch = ""
	} else {
		NameSearch = *request.NameSearch
	}

	if request.FK3 == nil {
		FK3 = ""
	} else {
		FK3 = *request.FK3
	}

	if request.Order == nil {
		Order = ""
	} else {
		Order = *request.Order
	}

	client := grpcProduct.NewProductServiceClient(conn)
	res, err := client.GetAllProduct(ctx, &grpcProduct.RequestAdditionalData{
		Id:         int32(ID),
		Page:       int32(Page),
		Limit:      int32(Limit),
		Fk_1:       int32(FK1),
		Fk_2:       FK2,
		Fk_3:       FK3,
		NameSearch: NameSearch,
		Order:      Order,
		Sort:       Sort,
	})

	if err != nil {
		return response, err
	}

	var resData []domain.AllProductResponseDTO

	for _, i := range res.Data {
		var resDataDetailProduct []domain.DetailProduct
		for _, v := range i.DetailProduct {
			var discount int
			if v.Discount != 0 {
				discount = int(v.Discount)
			} else {
				discount = 0
			}
			resDataDetailProduct = append(resDataDetailProduct, domain.DetailProduct{
				ID:         v.Id,
				Stok:       int(v.Stok),
				Duration:   v.Duration,
				Price:      v.Price,
				FinalPrice: v.FinalPrice,
				Tax:        int(v.Tax),
				Discount:   &discount,
			})
		}
		resData = append(resData,
			domain.AllProductResponseDTO{
				ID:            i.Id,
				Name:          i.Name,
				SKU:           i.Sku,
				Tipe:          i.Tipe,
				Description:   i.Desc,
				DetailProduct: resDataDetailProduct,
				DtmCrt:        i.DtmCrt,
				DtmUpd:        i.DtmUpd,
			})
	}

	response = domain.GetAllProductResponse{
		MetaData: domain.MetaData{
			TotalData: uint(res.MetaData.TotalData),
			TotalPage: uint(res.MetaData.TotalPage),
			Page:      uint(res.MetaData.Page),
			Limit:     uint(res.MetaData.Limit),
			Sort:      res.MetaData.Sort,
			Order:     res.MetaData.Order,
		},
		Data: resData,
	}
	return
}
