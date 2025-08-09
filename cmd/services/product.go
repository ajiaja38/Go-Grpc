package services

import (
	"context"
	"go-grpc/helpers"
	productPb "go-grpc/pb"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
	DB *gorm.DB
}

func (p *ProductService) GetProducts(ctx context.Context, pageParam *productPb.Page) (*productPb.Products, error) {
	var page int64 = 1
	var limit int64 = 5

	if pageParam.GetPage() != 0 || pageParam.GetLimit() != 0 {
		page = pageParam.GetPage()
		limit = pageParam.GetLimit()
	}

	var pagination productPb.Pagination
	var products []*productPb.Product

	sql := p.DB.Table("products as p").
		Joins("LEFT JOIN categories as c ON c.id = p.category_id").
		Select("p.id, p.name, p.price, p.stock, c.id as category_id, c.name as category_name")

	offset, limit := helpers.Pagination(sql, page, limit, &pagination)

	rows, err := sql.Offset(int(offset)).Limit(int(limit)).Rows()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var product productPb.Product
		var category productPb.Category

		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &category.Id, &category.Name); err != nil {
			log.Fatalf("Failed to scan the data: %v", err.Error())
		}

		product.Category = &category

		products = append(products, &product)
	}

	response := &productPb.Products{
		Pagination: &pagination,
		Data:       products,
	}

	return response, nil
}
