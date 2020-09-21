package product

type Product struct {
	ProductID    int    `json:"productId"`
	Manufacturer string `json:"manufacturer"`
	ProductName  string `json:"productName"`
}
