package toyyibpay

import (
	"fmt"
	"log"
)

type Category struct {
}

type CategoryParam struct {
	UserSecretKey       string
	CategoryName        string `json:"catname"`
	CategoryDescription string `json:"catdescription"`
}

//CategoryResponse ...
type CategoryResponse struct {
	CategoryCode string `json:"CategoryCode"`
}

// CategoryResponseList is categoryResponse list
type CategoryResponseList []CategoryResponse

// CreateCategory will create a single bill and return the unique code id
func (c *Client) CreateCategory(billParams CategoryParam) (string, error) {

	var err error
	categoryResponse := &CategoryResponseList{}
	billParams.UserSecretKey = c.UserSecretKey
	req, err := c.NewRequest("createCategory", billParams)

	if err != nil {
		return "", err
	}

	err = c.CallWithJSONResponse(req, categoryResponse)

	if err != nil {
		return "", err
	}
	// billResponse

	return (*categoryResponse)[0].CategoryCode, nil
}

// APIGetCategoryResponse ...
type APIGetCategoryResponse struct {
	CategoryName        string `json:"catname"`
	CategoryDescription string `json:"catdescription"`
	CategoryStatus      string `json:"categoryStatus"`
}

// APICategoryResponseList ...
type APICategoryResponseList []CategoryResponseList

// GetCategory ...
func (c *Client) GetCategory(transactionParams interface{}) (*APICategoryResponseList, error) {
	var err error
	categoryResponse := &APICategoryResponseList{}
	req, err := c.NewRequest("getCategory", transactionParams)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = c.CallWithJSONResponse(req, categoryResponse)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("categoryResponse ", categoryResponse)

	return categoryResponse, nil
}
