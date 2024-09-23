package api

import (
	"dynamic/database"
	"dynamic/types"
	"fmt"
)

type ApiClient struct {
	Datas database.DynamodbClient
}

func NewApiClient(datass database.DynamodbClient) ApiClient {
	return ApiClient{
		Datas: datass,
	}
}

func (api *ApiClient) RegisterClient(body types.UserDetails) error {
	if body.Username == "" || body.Password == "" {
		return fmt.Errorf("please enter the details correctly")
	}

	result, err := api.Datas.DoesexistorNot(body.Username)
	if err != nil {
		return fmt.Errorf("we are getting the %w", err)
	}
	if result {
		return fmt.Errorf("the User already exist")
	}
	e := api.Datas.InsertDetails(body)
	if e != nil {
		return fmt.Errorf("the following error is observed %w", e)
	}
	return nil
}
