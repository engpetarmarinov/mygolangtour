package product

import (
	"database/sql"
	"errors"
	"github.com/engpetarmarinov/mygolangtour/rest/database"
	"log"
)

func getProduct(productID int) (*Product, error) {
	row := database.DbConn.QueryRow(
		`SELECT ProductId, Manufacturer, ProductName FROM Products WHERE ProductId=@productID`,
		sql.Named("productID", productID))
	product := &Product{}
	err := row.Scan(
		&product.ProductID,
		&product.Manufacturer,
		&product.ProductName,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return product, nil
}

func getProductList() ([]Product, error) {
	results, err := database.DbConn.Query(`SELECT ProductId, Manufacturer, ProductName FROM Products`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()

	products := make([]Product, 0)
	for results.Next() {
		var product Product
		results.Scan(&product.ProductID,
			&product.Manufacturer,
			&product.ProductName)

		products = append(products, product)
	}
	return products, nil
}

func removeProduct(productID int) error {
	_, err := database.DbConn.Exec(`DELETE FROM products where productId = @p1`, productID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func updateProduct(product Product) error {
	if product.ProductID == 0 {
		return errors.New("product has invalid ID")
	}
	_, err := database.DbConn.Exec(`UPDATE products SET 
		Manufacturer=@Manufacturer,
		ProductName=@productName
		WHERE productId=@ProductID`,
		sql.Named("ProductID", product.ProductID),
		sql.Named("Manufacturer", product.Manufacturer),
		sql.Named("ProductName", product.ProductName))
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func insertProduct(product Product) (int, error) {
	rows, err := database.DbConn.Query(
		`INSERT INTO products  
		(Manufacturer, ProductName) 
		VALUES (@Manufacturer, @ProductName);
		SELECT ProductId = convert(bigint, SCOPE_IDENTITY())`,
		sql.Named("Manufacturer", product.Manufacturer),
		sql.Named("ProductName", product.ProductName))
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	var lastInsertId int64
	for rows.Next() {
		err = rows.Scan(&lastInsertId)
	}
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return int(lastInsertId), nil
}
