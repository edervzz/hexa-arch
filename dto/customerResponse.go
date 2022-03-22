package dto

type CustomerResponse struct {
	ID       int    `json:"id" xml:"id" db:"id"`
	Fullname string `json:"fullName" xml:"name"`
	City     string `json:"city" xml:"city"`
	Zipcode  string `json:"zipCode" xml:"zipcode"`
}
