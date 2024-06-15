package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/wlrudi19/go-simple-web/app/product/model"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, tx *sql.Tx, product model.Product) error
	FindProduct(ctx context.Context, id int) (model.FindProductResponse, error)
	FindProductAll(ctx context.Context) ([]model.Product, error)
	DeleteProduct(ctx context.Context, tx *sql.Tx, id int) error

	//user
	UpdateProduct(ctx context.Context, tx *sql.Tx, id int, fields model.UpdateProductRequest) error
	CreateOrder(ctx context.Context, tx *sql.Tx, order model.Order) error
	UpdateOrder(ctx context.Context, tx *sql.Tx, fields model.Order) error
	FindOrderById(ctx context.Context, userId int) ([]model.Order, error)
	ReduceStok(ctx context.Context, tx *sql.Tx, id int, quantity int) error
	CreateOrderHistory(ctx context.Context, tx *sql.Tx, order model.OrderHistory) error
	FindOrderHistoryById(ctx context.Context, userId int) ([]model.OrderHistory, error)
	WithTransaction() (*sql.Tx, error)
}

type productrepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productrepository{
		db: db,
	}
}

func (pr *productrepository) intsToString(numbers []int) string {
	str := make([]string, len(numbers))
	for i, num := range numbers {
		str[i] = fmt.Sprint(num)
	}
	return strings.Join(str, ",")
}

func (pr *productrepository) stringToIntSlice(str string) ([]int, error) {
	if str == "" {
		return nil, nil
	}

	strSlice := strings.Split(str, ",")
	intSlice := make([]int, len(strSlice))

	for i, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		intSlice[i] = num
	}

	return intSlice, nil
}

func (pr *productrepository) WithTransaction() (*sql.Tx, error) {
	tx, err := pr.db.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (pr *productrepository) CreateProduct(ctx context.Context, tx *sql.Tx, product model.Product) error {
	log.Printf("[QUERY] creating product: %s", product.Name)

	var id int
	sql := "insert into products (name,description,amount,stok) values ($1, $2, $3, $4) RETURNING id"
	err := tx.QueryRowContext(ctx, sql, product.Name, product.Description, product.Amount, product.Stok).Scan(&id)
	if err != nil {
		log.Printf("[QUERY] failed insert into database :%v", err)
		return err
	}

	product.Id = int(id)
	return nil
}

func (pr *productrepository) CreateOrder(ctx context.Context, tx *sql.Tx, order model.Order) error {
	log.Printf("[QUERY] creating order: %v", order)

	var id int
	sql := "insert into orders (product_id,user_id,amount,total,status) values ($1, $2, $3, $4, $5) RETURNING id"
	err := tx.QueryRowContext(ctx, sql, order.ProductID, order.UserID, order.Amount, order.Total, order.Status).Scan(&id)
	if err != nil {
		log.Printf("[QUERY] failed insert into database :%v", err)
		return err
	}

	order.Id = int(id)
	return nil
}

func (pr *productrepository) CreateOrderHistory(ctx context.Context, tx *sql.Tx, order model.OrderHistory) error {
	log.Printf("[QUERY] creating order history: %v", order)

	var id int
	coStr := pr.intsToString(order.CollectOrder)
	sql := "insert into orders_history (status,collect_order, user_id) values ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, sql, order.Status, coStr, order.UserID).Scan(&id)
	if err != nil {
		log.Printf("[QUERY] failed insert into database :%v", err)
		return err
	}

	order.Id = int(id)
	return nil
}

func (pr *productrepository) FindProduct(ctx context.Context, id int) (model.FindProductResponse, error) {
	log.Printf("[QUERY] finding product with id: %d", id)

	var product model.FindProductResponse

	sql := "select name, description, amount, stok from products p where deleted_on isnull and id = $1"
	err := pr.db.QueryRowContext(ctx, sql, id).Scan(
		&product.Name,
		&product.Description,
		&product.Amount,
		&product.Stok,
	)

	if err != nil {
		log.Printf("[QUERY] failed to finding product, %v", err)
		return product, err
	}

	return product, nil
}

func (pr *productrepository) FindProductAll(ctx context.Context) ([]model.Product, error) {
	log.Printf("[QUERY] find all products")

	sql := "select id, name, description, amount, stok from products where deleted_on isnull"
	rows, err := pr.db.QueryContext(ctx, sql)

	if err != nil {
		log.Printf("[QUERY]] failed to finding products, %v", err)
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Amount,
			&product.Stok,
		)

		if err != nil {
			log.Fatalf("[QUERY] failed to finding product row: %v", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (pr *productrepository) FindOrderById(ctx context.Context, userId int) ([]model.Order, error) {
	log.Printf("[QUERY] find order")

	sql := "select id, product_id, user_id, amount, total, status from orders where user_id = $1 and deleted_on isnull"
	rows, err := pr.db.QueryContext(ctx, sql, userId)
	if err != nil {
		log.Printf("[QUERY]] failed to finding products, %v", err)
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		err := rows.Scan(
			&order.Id,
			&order.ProductID,
			&order.UserID,
			&order.Amount,
			&order.Total,
			&order.Status,
		)

		if err != nil {
			log.Fatalf("[QUERY] failed to finding order row: %v", err)
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (pr *productrepository) FindOrderHistoryById(ctx context.Context, userId int) ([]model.OrderHistory, error) {
	log.Printf("[QUERY] find order history")

	sql := "select id, status, collect_order from orders_history where user_id = $1 and deleted_on isnull order by created_on desc"
	rows, err := pr.db.QueryContext(ctx, sql, userId)
	if err != nil {
		log.Printf("[QUERY]] failed to finding products, %v", err)
		return nil, err
	}
	defer rows.Close()

	var orders []model.OrderHistory
	for rows.Next() {
		var order model.OrderHistory
		var collectOrderStr string
		err := rows.Scan(
			&order.Id,
			&order.Status,
			&collectOrderStr,
		)
		if err != nil {
			log.Fatalf("[QUERY] failed to finding order row: %v", err)
			return nil, err
		}

		order.CollectOrder, err = pr.stringToIntSlice(collectOrderStr)
		if err != nil {
			log.Fatalf("[QUERY] failed to convert CollectOrder to []int: %v", err)
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (pr *productrepository) DeleteProduct(ctx context.Context, tx *sql.Tx, id int) error {
	log.Printf("[QUERY] deleting product with id: %d", id)

	var deletedOn sql.NullTime

	checkSQL := "SELECT deleted_on FROM products WHERE id = $1"
	err := tx.QueryRowContext(ctx, checkSQL, id).Scan(&deletedOn)

	if err != nil {
		log.Printf("[QUERY] failed to deleting product: %v", err)
		return err
	}

	if deletedOn.Valid {
		err = errors.New("product has been deleted before")
		log.Printf("[QUERY] failed to deleting product: %v", err)
		return err
	}

	sql := "update products set deleted_on = now() where id = $1"
	_, err = tx.ExecContext(ctx, sql, id)

	if err != nil {
		log.Printf("[QUERY] failed deleting product, %v", err)
		return err
	}

	return nil
}

func (pr *productrepository) UpdateProduct(ctx context.Context, tx *sql.Tx, id int, fields model.UpdateProductRequest) error {
	log.Printf("[QUERY] updating product with id: %d", id)

	updateBuilder := squirrel.Update("products").
		Where(squirrel.Eq{"id": id})

	if fields.Name != nil {
		updateBuilder = updateBuilder.Set("name", *fields.Name)
	}
	if fields.Description != nil {
		updateBuilder = updateBuilder.Set("description", *fields.Description)
	}
	if fields.Amount != nil {
		updateBuilder = updateBuilder.Set("amount", *fields.Amount)
	}
	if fields.Stok != nil {
		updateBuilder = updateBuilder.Set("stok", *fields.Stok)
	}

	query, args, err := updateBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, args...)

	if err != nil {
		log.Printf("[QUERY] failed to update product, %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[QUERY] failed to update product, %v", err)
		return err
	}

	if rowsAffected == 0 {
		err := errors.New("sql: no rows in result set")
		log.Printf("[QUERY] product not found, %v", err)
		return err
	}

	return nil
}

func (pr *productrepository) UpdateOrder(ctx context.Context, tx *sql.Tx, fields model.Order) error {
	log.Printf("[QUERY] updating order with id: %d", fields.Id)

	updateBuilder := squirrel.Update("orders").
		Where(squirrel.Eq{"id": fields.Id})

	if fields.ProductID > 0 {
		updateBuilder = updateBuilder.Set("product_id", &fields.ProductID)
	}
	if fields.UserID > 0 {
		updateBuilder = updateBuilder.Set("user_id", &fields.UserID)
	}
	if fields.Amount != "" {
		updateBuilder = updateBuilder.Set("amount", &fields.Amount)
	}
	if fields.Total > 0 {
		updateBuilder = updateBuilder.Set("total", &fields.Total)
	}
	if fields.Status != "" {
		updateBuilder = updateBuilder.Set("status", &fields.Status)
	}

	query, args, err := updateBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		log.Printf("[QUERY] failed to update orders, %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[QUERY] failed to update orders, %v", err)
		return err
	}

	if rowsAffected == 0 {
		err := errors.New("sql: no rows in result set")
		log.Printf("[QUERY] product not found, %v", err)
		return err
	}

	return nil
}

func (pr *productrepository) ReduceStok(ctx context.Context, tx *sql.Tx, id int, quantity int) error {
	log.Printf("[QUERY] updating order with id: %d", id)

	updateBuilder := squirrel.Update("products").
		Set("stok", squirrel.Expr("stok - ?", quantity)).
		Where(squirrel.And{
			squirrel.Eq{"id": id},
			squirrel.Expr("stok >= ?", quantity),
		})

	query, args, err := updateBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		log.Printf("[QUERY] failed to update orders, %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[QUERY] failed to update orders, %v", err)
		return err
	}

	if rowsAffected == 0 {
		err := errors.New("sql: no rows in result set")
		log.Printf("[QUERY] product not found, %v", err)
		return err
	}

	return nil
}
