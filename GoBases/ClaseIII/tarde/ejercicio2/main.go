package main
type User struct {
	Name, Lastname, Email	string
	Products				[]Product
}

type Product struct {
	Name		string
	Price		float64
	Quantity	int
}

func newProduct(name *string, price *float64) *Product{
	return &Product{Name: *name, Price: *price}
}

func (u *User)addProduct(product *Product, quantity *int) {
	product.Quantity = *quantity
	u.Products = append(u.Products, *product)
}

func (u *User) deleteProduct() {
	u.Products = []Product{}
}

func main() {

}