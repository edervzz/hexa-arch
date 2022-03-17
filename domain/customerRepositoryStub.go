package domain

type CustomerReposityStub struct {
	customers []Customer
}

func (s CustomerReposityStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func (s CustomerReposityStub) Find(id int) (Customer, error) {

	for i, v := range s.customers {
		if v.ID == id {
			return s.customers[i], nil
		}
	}
	return Customer{}, nil
}

func NewCustomerRepositoryStub() CustomerReposityStub {
	customers := []Customer{
		{ID: 123, Name: "Eder", City: "EDOMEX", Zipcode: "52928"},
		{ID: 456, Name: "Sheila", City: "EDOMEX", Zipcode: "55027"},
		{ID: 789, Name: "Osmar", City: "EDOMEX", Zipcode: "11450"},
	}
	return CustomerReposityStub{customers: customers}
}
