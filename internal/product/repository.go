package product

import (
	"database/sql"
	"github.com/CaioMtho/unstock/internal/config"
)

func rowsToProducts(rows *sql.Rows) ([]Product, error){
	var products []Product
	for rows.Next() {
		var p Product
        err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.MinStock, &p.IsActive)
        if err != nil {
            return nil, err
        }
        products = append(products, p)
    }

    return products, nil
}

func GetAllProducts() ([]Product, error) {
    rows, err := config.DB.Query("SELECT * FROM products WHERE is_active = 1")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    products, err := rowsToProducts(rows)
	return products, err
}

func GetProductById(id int) (Product, error) {
    var p Product
    stmt := "SELECT id, name, price, stock, min_stock, is_active FROM products WHERE id = ? AND is_active = 1"
    err := config.DB.QueryRow(stmt, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.MinStock, &p.IsActive)
    if err != nil {
        return Product{}, err
    }
    return p, nil
}

func InsertProduct (p Product) error {
	stmt := `INSERT INTO products (name, price, stock, min_stock, is_active) VALUES (?, ?, ?, ?, 1)`
	_, err := config.DB.Exec(stmt, p.Name, p.Price, p.Stock, p.MinStock)
	return err
}


func GetLowStockProducts() ([]Product, error){
	rows, err := config.DB.Query("SELECT * FROM products WHERE MinStock > Stock AND is_active = 1 ORDER BY Stock")
	if err != nil {
		return nil, err
	}

	products, err := rowsToProducts(rows)

    return products, err
}

func DeleteProductById(id int) error {
	stmt := "UPDATE products SET is_active = 0 WHERE id = ? AND is_active = 1"
	_, err := config.DB.Exec(stmt, id);
	return err
}

func UpdateProduct(p Product) error {
	stmt := `UPDATE products SET name = ?, price = ?, min_stock = ? WHERE id = ? AND is_active = 1`
	_, err := config.DB.Exec(stmt, p.Name, p.Price, p.MinStock, p.ID)
	return err
}

func UpdateStock(id int, value int) error {
	stmtQuery := `SELECT stock FROM products WHERE id = ? AND is_active = 1`
	var currentStock int
	err := config.DB.QueryRow(stmtQuery, id).Scan(&currentStock)
	if err != nil {
		return err
	}
	
	newStock := currentStock + value
	stmtUpdate := `UPDATE products SET stock = ? WHERE id = ?`
	_, err = config.DB.Exec(stmtUpdate, newStock, id)

	if err != nil {
		return err
	}

	updatedProduct, err := GetProductById(id)
	if err == nil {
		StockUpdateChannel <- updatedProduct
	}

	return err
}