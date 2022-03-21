package domain

type CustomerReposityStub struct {
	customers []Customer
}

func (s CustomerReposityStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func (s CustomerReposityStub) Find(id int) (*Customer, error) {

	for i, v := range s.customers {
		if v.ID_customer == id {
			c := &s.customers[i]
			return c, nil
		}
	}
	return nil, nil
}

func NewCustomerRepositoryStub() CustomerReposityStub {
	customers := []Customer{
		{ID_customer: 123, Name: "Eder", City: "EDOMEX", Zipcode: "52928"},
		{ID_customer: 456, Name: "Sheila", City: "EDOMEX", Zipcode: "55027"},
		{ID_customer: 789, Name: "Osmar", City: "EDOMEX", Zipcode: "11450"},
	}
	return CustomerReposityStub{customers: customers}
}
