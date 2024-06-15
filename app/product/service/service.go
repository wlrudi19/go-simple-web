package service

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/wlrudi19/go-simple-web/app/product/model"
	"github.com/wlrudi19/go-simple-web/app/product/repository"
)

type ProductLogic interface {
	CreateProductLogic(ctx context.Context, req model.CreateProductRequest) error
	FindProductLogic(ctx context.Context, id int) (model.FindProductResponse, error)
	FindProductAllLogic(ctx context.Context) ([]model.Product, error)
	DeleteProductLogic(ctx context.Context, id int) error
	UpdateProductLogic(ctx context.Context, id int, fields model.UpdateProductRequest) error

	//user
	OrderLogic(ctx context.Context, param model.Order) error
	FindOrderByIdLogic(ctx context.Context, userId int) ([]model.Order, error)
	FindOrderConditionLogic(ctx context.Context, userId int, param model.Order) ([]model.Order, error)
	OrderSummaryLogic(ctx context.Context, userId int) (model.OrderSummary, error)
	BulkUpdateOrderLogic(ctx context.Context, params []model.BulkUpdateOrder) error
	CreateOrderHistoryLogic(ctx context.Context, param model.OrderHistory) error
	FindOrderHistoryByIdLogic(ctx context.Context, userId int) ([]model.OrderHistory, error)
}

type productlogic struct {
	ProductRepository repository.ProductRepository
}

func NewProductLogic(productRepository repository.ProductRepository) ProductLogic {
	return &productlogic{
		ProductRepository: productRepository,
	}
}

func (l *productlogic) CreateProductLogic(ctx context.Context, req model.CreateProductRequest) error {
	log.Printf("[LOGIC] create new product: %s", req.Name)

	tx, err := l.ProductRepository.WithTransaction()

	if err != nil {
		log.Printf("[LOGIC] failed to create product :%v", err)
		return err
	}

	product := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Amount:      req.Amount,
		Stok:        req.Stok,
	}

	err = l.ProductRepository.CreateProduct(ctx, tx, product)

	if err != nil {
		log.Printf("[LOGIC] failed to create product :%v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	log.Printf("[LOGIC] created product success with id: %d", product.Id)
	return nil
}

func (l *productlogic) FindProductLogic(ctx context.Context, id int) (model.FindProductResponse, error) {
	log.Printf("[LOGIC] find product with id: %d", id)

	var product model.FindProductResponse

	product, err := l.ProductRepository.FindProduct(ctx, id)

	if err != nil {
		log.Printf("[LOGIC] failed to find product :%v", err)
		return product, err
	}

	log.Printf("[LOGIC] product find successfulyy, id: %d", id)
	return product, nil
}

func (l *productlogic) FindProductAllLogic(ctx context.Context) ([]model.Product, error) {
	log.Printf("[LOGIC] find all products")

	var products []model.Product

	products, err := l.ProductRepository.FindProductAll(ctx)
	if err != nil {
		log.Printf("[LOGIC] failed to find product s:%v", err)
		return products, err
	}

	log.Printf("[LOGIC] products find successfulyy")
	return products, nil
}

func (l *productlogic) DeleteProductLogic(ctx context.Context, id int) error {
	log.Printf("[LOGIC] delete product with id: %d", id)

	tx, err := l.ProductRepository.WithTransaction()

	if err != nil {
		log.Printf("[LOGIC] failed to delete product :%v", err)
		return err
	}

	err = l.ProductRepository.DeleteProduct(ctx, tx, id)

	if err != nil {
		log.Printf("[LOGIC] failed to delete product :%v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	log.Printf("[LOGIC] deleted product success with id: %d", id)
	return nil
}

func (l *productlogic) UpdateProductLogic(ctx context.Context, id int, fields model.UpdateProductRequest) error {
	log.Printf("[LOGIC] update product with id: %d", id)

	tx, err := l.ProductRepository.WithTransaction()
	if err != nil {
		log.Printf("[LOGIC] failed to update product :%v", err)
		return err
	}

	err = l.ProductRepository.UpdateProduct(ctx, tx, id, fields)
	if err != nil {
		log.Printf("[LOGIC] failed to update product :%v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	log.Printf("[LOGIC] update product success with id: %d", id)
	return nil
}

func (l *productlogic) OrderLogic(ctx context.Context, param model.Order) error {
	log.Printf("[LOGIC] order with param: %v", param)

	tx, err := l.ProductRepository.WithTransaction()
	if err != nil {
		log.Printf("[LOGIC] failed to order :%v", err)
		return err
	}

	err = l.ProductRepository.CreateOrder(ctx, tx, param)
	if err != nil {
		log.Printf("[LOGIC] failed to create order :%v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	log.Printf("[LOGIC] update product success")
	return nil
}

func (l *productlogic) CreateOrderHistoryLogic(ctx context.Context, param model.OrderHistory) error {
	log.Printf("[LOGIC] order with param: %v", param)

	tx, err := l.ProductRepository.WithTransaction()
	if err != nil {
		log.Printf("[LOGIC] failed to order :%v", err)
		return err
	}

	err = l.ProductRepository.CreateOrderHistory(ctx, tx, param)
	if err != nil {
		log.Printf("[LOGIC] failed to create order history :%v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	log.Printf("[LOGIC] update product success")
	return nil
}

func (l *productlogic) FindOrderByIdLogic(ctx context.Context, userId int) ([]model.Order, error) {
	log.Printf("[LOGIC] find order with id: %v", userId)

	var orders []model.Order

	orders, err := l.ProductRepository.FindOrderById(ctx, userId)
	if err != nil {
		log.Printf("[LOGIC] failed to find order s:%v", err)
		return orders, err
	}

	log.Printf("[LOGIC] order find successfully")
	return orders, nil
}

func (l *productlogic) FindOrderHistoryByIdLogic(ctx context.Context, userId int) ([]model.OrderHistory, error) {
	log.Printf("[LOGIC] find order history with id: %v", userId)

	var history []model.OrderHistory

	history, err := l.ProductRepository.FindOrderHistoryById(ctx, userId)
	if err != nil {
		log.Printf("[LOGIC] failed to find order history:%v", err)
		return history, err
	}

	for k, v := range history {
		for _, val := range v.CollectOrder {
			param := model.Order{
				Id: val,
			}
			order, err := l.FindOrderConditionLogic(ctx, userId, param)
			if err != nil {
				log.Printf("[LOGIC] failed to find order:%v", err)
				return history, err
			}

			history[k].Amount = order[0].Amount
		}
	}

	log.Printf("[LOGIC] order history find successfully")
	return history, nil
}

func (l *productlogic) FindOrderConditionLogic(ctx context.Context, userId int, param model.Order) ([]model.Order, error) {
	log.Printf("[LOGIC] find order with id: %v %v", userId, param)

	var result []model.Order

	orders, err := l.ProductRepository.FindOrderById(ctx, userId)
	if err != nil {
		log.Printf("[LOGIC] failed to find order s:%v", err)
		return orders, err
	}

	summaryMap := make(map[int]model.Order)
	for _, order := range orders {
		if order.Status == param.Status {
			if existing, found := summaryMap[order.ProductID]; found {
				existing.Total += order.Total
				existing.Amount = l.sumAmounts(existing.Amount, order.Amount)
				existing.CollectId = append(existing.CollectId, order.Id)
				summaryMap[order.ProductID] = existing
			} else {
				summaryMap[order.ProductID] = model.Order{
					UserID:    userId,
					ProductID: order.ProductID,
					Total:     order.Total,
					Amount:    order.Amount,
					Status:    order.Status,
					CollectId: []int{order.Id},
				}
			}
		}
	}

	for _, summary := range summaryMap {
		product, err := l.ProductRepository.FindProduct(ctx, summary.ProductID)
		if err != nil {
			log.Printf("[LOGIC] failed to find product:%v", err)
			return orders, err
		}
		summary.ProductName = product.Name
		result = append(result, summary)
	}

	log.Printf("[LOGIC] find order cart successfully")
	return result, nil
}

func (l *productlogic) sumAmounts(amount1, amount2 string) string {
	a1, _ := strconv.ParseFloat(amount1, 64)
	a2, _ := strconv.ParseFloat(amount2, 64)

	return fmt.Sprintf("%.2f", a1+a2)
}

func (l *productlogic) OrderSummaryLogic(ctx context.Context, userId int) (model.OrderSummary, error) {
	log.Printf("[LOGIC] order summary with user id: %v", userId)
	var result model.OrderSummary
	var filter model.Order

	filter.Status = "PENDING"
	orders, err := l.FindOrderConditionLogic(ctx, userId, filter)
	if err != nil {
		log.Printf("[LOGIC] failed to find order s:%v", err)
		return result, err
	}

	var totalAmount float64
	var kupon int

	for _, order := range orders {
		amtFloat, err := strconv.ParseFloat(order.Amount, 64)
		if err != nil {
			log.Printf("[LOGIC] failed to parse amount: %v", err)
			return result, err
		}
		totalAmount += amtFloat

		if amtFloat > 50000 {
			kupon += 1
		}
	}

	kupon += int(totalAmount) / 100000

	result = model.OrderSummary{
		Data:       orders,
		Kupon:      kupon,
		TotalBayar: totalAmount,
	}

	return result, nil
}

func (l *productlogic) UpdateOrderLogic(ctx context.Context, params model.Order) error {
	log.Printf("[LOGIC] update order with param: %v", params)

	tx, err := l.ProductRepository.WithTransaction()
	if err != nil {
		log.Printf("[LOGIC] failed to order product :%v", err)
		return err
	}

	err = l.ProductRepository.UpdateOrder(ctx, tx, params)
	if err != nil {
		log.Printf("[LOGIC] failed to order product :%v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	log.Printf("[LOGIC] update order success")
	return nil
}

func (l *productlogic) BulkUpdateOrderLogic(ctx context.Context, params []model.BulkUpdateOrder) error {
	log.Printf("[LOGIC] bulk update order with param: %v", params)

	var mapCollect []int
	for k, val := range params {
		for _, v := range val.CollectId {
			var order model.Order
			order = model.Order{
				Id:     v,
				Status: params[k].Status,
			}
			err := l.UpdateOrderLogic(ctx, order)
			if err != nil {
				log.Printf("[LOGIC] failed to order product :%v", err)
				return err
			}
		}

		if val.Status == "PAID" {
			mapCollect = append(mapCollect, val.CollectId...)
		}

		if val.ProductUpdate {
			err := l.ReduceStokLogic(ctx, val.ProductID, val.Total)
			if err != nil {
				log.Printf("[LOGIC] failed to update product :%v", err)
				return err
			}
		}
	}

	if len(mapCollect) > 0 {
		var oh model.OrderHistory
		oh = model.OrderHistory{
			Status:       "PAID",
			CollectOrder: mapCollect,
		}
		err := l.CreateOrderHistoryLogic(ctx, oh)
		if err != nil {
			log.Printf("[LOGIC] failed to order order history :%v", err)
			return err
		}
	}

	log.Printf("[LOGIC] bulk update order success")
	return nil
}

func (l *productlogic) ReduceStokLogic(ctx context.Context, id int, quantity int) error {
	log.Printf("[LOGIC] update product with id: %d", id)

	tx, err := l.ProductRepository.WithTransaction()
	if err != nil {
		log.Printf("[LOGIC] failed to update product :%v", err)
		return err
	}

	err = l.ProductRepository.ReduceStok(ctx, tx, id, quantity)
	if err != nil {
		log.Printf("[LOGIC] failed to update product :%v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	log.Printf("[LOGIC] update product success with id: %d", id)
	return nil
}
