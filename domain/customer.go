package domain

type Customer struct {
	ID      int    `json:"id" xml:"id"`
	Name    string `json:"fullname" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	Find(id int) (Customer, error)
}
