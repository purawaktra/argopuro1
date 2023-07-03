package entities

type Accounts struct {
	AccountId    string `json:"account-id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Zipcode      string `json:"zipcode"`
	EmailAddress string `json:"email_address"`
	PhoneNumber  string `json:"phone_number"`
}
