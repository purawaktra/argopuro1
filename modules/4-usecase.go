package modules

import (
	"errors"
	"github.com/purawaktra/argopuro1-go/entities"
)

type Argopuro1Usecase struct {
	repo Argopuro1Repo
}

func (uc Argopuro1Usecase) SelectAccountById(accountId int, offset int) ([]Accounts, error) {
	// create check input on account id and offset
	if accountId < 1 {
		return nil, errors.New("accountId can not be nil, negative or zero")
	}
	if offset < 0 {
		return nil, errors.New("offset can not be negative")
	}

	// convert input to entity
	account := entities.Accounts{AccountId: uint(accountId)}

	// call repo for the account id
	accounts, err := uc.repo.SelectAccountById(account, uint(offset))

	// check for error on call usecase
	if err != nil {
		return nil, err
	}

	// convert entity to dto
	results := make([]Accounts, 0)
	for _, account := range accounts {
		// get city data
		// convert input to entity
		city := entities.Cities{CityId: account.City}
		cities, err := uc.repo.SelectCityById(city, 0)

		// check for error on call usecase
		if err != nil || cities == nil {
			cities = []entities.Cities{{Name: ""}}
		}

		// get province data
		// convert input to entity
		province := entities.Provinces{ProvinceId: account.Province}
		provinces, err := uc.repo.SelectProvinceById(province, 0)

		// check for error on call usecase
		if err != nil || provinces == nil {
			provinces = []entities.Provinces{{Name: ""}}
		}

		results = append(results, Accounts{
			AccountId:    account.AccountId,
			FirstName:    account.FirstName,
			LastName:     account.LastName,
			Address:      account.Address,
			City:         cities[0].Name,
			Province:     provinces[0].Name,
			Zipcode:      account.Zipcode,
			EmailAddress: account.EmailAddress,
			PhoneNumber:  account.PhoneNumber,
		})
	}
	// create return
	return results, nil
}

func (uc Argopuro1Usecase) SelectAccountByEmail(emailAddress string, offset int) ([]Accounts, error) {
	// create check input on account email address and offset
	if emailAddress == "" {
		return nil, errors.New("emailAddress can not be empty")
	}
	if offset < 0 {
		return nil, errors.New("offset can not be negative")
	}

	// convert input to entity
	account := entities.Accounts{EmailAddress: emailAddress}

	// call repo for the account email address
	accounts, err := uc.repo.SelectAccountByEmail(account, uint(offset))

	// check for error on call usecase
	if err != nil {
		return nil, err
	}

	// convert entity to dto
	result := make([]Accounts, 0)
	for _, account := range accounts {
		// get city data
		// convert input to entity
		city := entities.Cities{CityId: account.City}
		cities, err := uc.repo.SelectCityById(city, 0)

		// check for error on call usecase
		if err != nil || cities == nil {
			cities = []entities.Cities{{Name: ""}}
		}

		// get province data
		// convert input to entity
		province := entities.Provinces{ProvinceId: account.Province}
		provinces, err := uc.repo.SelectProvinceById(province, 0)

		// check for error on call usecase
		if err != nil || provinces == nil {
			provinces = []entities.Provinces{{Name: ""}}
		}

		result = append(result, Accounts{
			AccountId:    account.AccountId,
			FirstName:    account.FirstName,
			LastName:     account.LastName,
			Address:      account.Address,
			City:         cities[0].Name,
			Province:     provinces[0].Name,
			Zipcode:      account.Zipcode,
			EmailAddress: account.EmailAddress,
			PhoneNumber:  account.PhoneNumber,
		})
	}
	// create return
	return result, nil
}

func (uc Argopuro1Usecase) InsertSingleAccount(body Accounts) (Accounts, error) {
	// create check input on account email
	if body.EmailAddress == "" {
		return Accounts{}, errors.New("emailAddress can not be empty")
	}

	// convert city input to entity
	city := entities.Cities{Name: body.City}

	// convert city name to id
	cities, err := uc.repo.SelectCityByName(city, 0)

	// check for error on call usecase
	if err != nil || cities == nil {
		cities = []entities.Cities{{Name: ""}}
	}

	// convert province input to entity
	province := entities.Provinces{Name: body.Province}

	// convert province name to id
	provinces, err := uc.repo.SelectProvinceById(province, 0)

	// check for error on call usecase
	if err != nil || provinces == nil {
		provinces = []entities.Provinces{{Name: ""}}
	}

	// convert input to entity
	account := entities.Accounts{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Address:      body.Address,
		City:         cities[0].CityId,
		Province:     provinces[0].ProvinceId,
		Zipcode:      body.Zipcode,
		EmailAddress: body.EmailAddress,
		PhoneNumber:  body.PhoneNumber,
	}

	// call repo for the insert account
	account, err = uc.repo.InsertSingleAccount(account)

	// check for error on call usecase
	if err != nil {
		return Accounts{}, err
	}

	// convert entity to dto
	result := Accounts{
		AccountId:    account.AccountId,
		FirstName:    account.FirstName,
		LastName:     account.LastName,
		Address:      account.Address,
		City:         cities[0].Name,
		Province:     provinces[0].Name,
		Zipcode:      account.Zipcode,
		EmailAddress: account.EmailAddress,
		PhoneNumber:  account.PhoneNumber,
	}

	// create return
	return result, nil
}

func (uc Argopuro1Usecase) UpdateSingleAccountById(body Accounts) (Accounts, error) {
	// create check input on account id and email address
	if body.AccountId < 1 {
		return Accounts{}, errors.New("accountId can not be nil, negative or zero")
	}

	if body.EmailAddress == "" {
		return Accounts{}, errors.New("emailAddress can not be empty")
	}

	// convert city input to entity
	city := entities.Cities{Name: body.City}

	// convert city name to id
	cities, err := uc.repo.SelectCityByName(city, 0)

	// check for error on call usecase
	if err != nil || cities == nil {
		cities = []entities.Cities{{Name: ""}}
	}

	// convert province input to entity
	province := entities.Provinces{Name: body.Province}

	// convert province name to id
	provinces, err := uc.repo.SelectProvinceById(province, 0)

	// check for error on call usecase
	if err != nil || provinces == nil {
		provinces = []entities.Provinces{{Name: ""}}
	}

	// convert input to entity
	account := entities.Accounts{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Address:      body.Address,
		City:         cities[0].CityId,
		Province:     provinces[0].ProvinceId,
		Zipcode:      body.Zipcode,
		EmailAddress: body.EmailAddress,
		PhoneNumber:  body.PhoneNumber,
	}

	// call repo for the insert account
	account, err = uc.repo.UpdateSingleAccountById(account)

	// check for error on call usecase
	if err != nil {
		return Accounts{}, err
	}

	// convert entity to dto
	result := Accounts{
		AccountId:    account.AccountId,
		FirstName:    account.FirstName,
		LastName:     account.LastName,
		Address:      account.Address,
		City:         cities[0].Name,
		Province:     cities[0].Name,
		Zipcode:      account.Zipcode,
		EmailAddress: account.EmailAddress,
		PhoneNumber:  account.PhoneNumber,
	}

	// create return
	return result, nil
}

func (uc Argopuro1Usecase) DeleteSingleAccountById(body Accounts) (Accounts, error) {
	// create check input on account id
	if body.AccountId < 1 {
		return Accounts{}, errors.New("accountId can not be nil, negative or zero")
	}

	// convert city input to entity
	city := entities.Cities{Name: body.City}

	// convert city name to id
	cities, err := uc.repo.SelectCityByName(city, 0)

	// check for error on call usecase
	if err != nil || cities == nil {
		cities = []entities.Cities{{Name: ""}}
	}

	// convert province input to entity
	province := entities.Provinces{Name: body.Province}

	// convert province name to id
	provinces, err := uc.repo.SelectProvinceById(province, 0)

	// check for error on call usecase
	if err != nil || provinces == nil {
		provinces = []entities.Provinces{{Name: ""}}
	}

	// convert input to entity
	account := entities.Accounts{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Address:      body.Address,
		City:         cities[0].CityId,
		Province:     provinces[0].ProvinceId,
		Zipcode:      body.Zipcode,
		EmailAddress: body.EmailAddress,
		PhoneNumber:  body.PhoneNumber,
	}

	// call repo for the insert account
	account, err = uc.repo.UpdateSingleAccountById(account)

	// check for error on call usecase
	if err != nil {
		return Accounts{}, err
	}

	// convert entity to dto
	result := Accounts{
		AccountId:    account.AccountId,
		FirstName:    account.FirstName,
		LastName:     account.LastName,
		Address:      account.Address,
		City:         cities[0].Name,
		Province:     provinces[0].Name,
		Zipcode:      account.Zipcode,
		EmailAddress: account.EmailAddress,
		PhoneNumber:  account.PhoneNumber,
	}

	// create return
	return result, nil
}
