package domain

type Customer struct {
	ID_customer int    `json:"id" xml:"id" db:"id"`
	Name        string `json:"fullname" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
}

type ICustomerRepository interface {
	FindAll() ([]Customer, error)
	Find(id int) (*Customer, error)
}
