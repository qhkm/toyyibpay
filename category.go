package toyyibpay

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
	err = c.CallWithJSONResponse(req, categoryResponse)
	return (*categoryResponse)[0].CategoryCode, err
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
	err = c.CallWithJSONResponse(req, categoryResponse)
	return categoryResponse, err
}
