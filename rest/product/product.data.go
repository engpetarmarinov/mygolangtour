package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var productMap = struct {
	sync.RWMutex
	m map[int]Product
}{m: make(map[int]Product)}

func init() {
	log.Println("product.data: loading products...")
	prodMap, err := loadProductMap()
	if err != nil {
		log.Fatal(err)
	}
	productMap.m = prodMap
	log.Println("product.data: products loaded ", len(productMap.m))
}

func loadProductMap() (map[int]Product, error) {
	fileName := "products.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("product.data: file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	productList := make([]Product, 0)
	err = json.Unmarshal(file, &productList)
	if err != nil {
		log.Fatal(err)
	}
	prodMap := make(map[int]Product)
	for i := 0; i < len(productList); i++ {
		prodMap[productList[i].ProductID] = productList[i]
	}
	return prodMap, nil
}

func getProduct(productID int) *Product {
	productMap.RLock()
	defer productMap.RUnlock()

	if product, ok := productMap.m[productID]; ok {
		return &product
	}
	return nil
}

func getProductList() []Product {
	productMap.RLock()
	defer productMap.RUnlock()
	products := make([]Product, 0, len(productMap.m))
	for _, value := range productMap.m {
		products = append(products, value)

	}
	return products
}
func addOrUpdateProduct(product Product) (int, error) {
	//if product id is set update, otherwise add
	addOrUpdateID := -1
	if product.ProductID > 0 {
		oldProduct := getProduct(product.ProductID)

		if oldProduct == nil {
			return 0, fmt.Errorf("product.data: product id [%d] doesn't exist", product.ProductID)
		}

		addOrUpdateID = product.ProductID
	} else {
		addOrUpdateID = getNextProductID()
		product.ProductID = addOrUpdateID
	}
	productMap.Lock()
	productMap.m[addOrUpdateID] = product
	productMap.Unlock()
	return addOrUpdateID, nil
}

func getNextProductID() int {
	productIDs := getProductIds()
	return productIDs[len(productIDs)-1] + 1
}

func getProductIds() []int {
	var productIds []int
	for _, product := range productMap.m {
		productIds = append(productIds, product.ProductID)
	}
	return productIds
}
