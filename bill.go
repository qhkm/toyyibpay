package toyyibpay

import (
	"fmt"
	"log"
)

type (
	// Bill is the resource representing Toyyibpay event
	Bill struct {
		UserSecretKey           string
		CategoryCode            string
		BillName                string
		BillDescription         string
		BillPriceSetting        int
		BillPayorInfo           int
		BillAmount              int
		BillReturnURL           string
		BillCallbackURL         string
		BillExternalReferenceNo string
		BillTo                  string
		BillEmail               string
		BillPhone               string
		BillSplitPayment        int
		BillSplitPaymentArgs    string
		// billMultiPayment        int
		BillPaymentChannel   string
		BillDisplayMerchant  int
		BillContentEmail     string
		BillChargeToCustomer int
	}

	// CreateBillParams is the struct sent to create bill
	CreateBillParams struct {
		UserSecretKey           string `form:"userSecretKey"`
		CategoryCode            string `form:"categoryCode"`
		BillName                string `form:"billName"`
		BillDescription         string `form:"billDescription"`
		BillPriceSetting        int    `form:"billPriceSetting"`
		BillPayorInfo           int    `form:"billPayorInfo"`
		BillAmount              int    `form:"billAmount"`
		BillReturnURL           string `form:"billReturnUrl"`
		BillCallbackURL         string `form:"billCallbackUrl"`
		BillExternalReferenceNo string `form:"billExternalReferenceNo"`
		BillTo                  string `form:"billTo"`
		BillEmail               string `form:"billEmail"`
		BillPhone               string `form:"billPhone"`
		BillSplitPayment        int    `form:"billSplitPayment"`
		BillSplitPaymentArgs    string `form:"billSplitPaymentArgs"`
		billMultiPayment        int    `form:"userSecretKey. omitempty"`
		BillPaymentChannel      string `form:"billPaymentChannel"`
		BillDisplayMerchant     int    `form:"billDisplayMerchant"`
		BillContentEmail        string `form:"billContentEmail"`
		BillChargeToCustomer    int    `form:"billChargeToCustomer"`
	}

	// RunBillParams is the params need to be sent before sent to run the bill
	RunBillParams struct {
		UserSecretKey         string `form:"userSecretKey"`
		BillCode              string `form:"billCode"`
		BillDescription       string `form:"billDescription"`
		BillPaymentAmount     int    `form:"billPaymentAmount"`
		BillPaymentPayorName  string `form:"billPaymentPayorName"`
		BillPaymentPayorPhone string `form:"billPaymentPayorPhone"`
		BillPaymentPayorEmail string `form:"billPaymentPayorEmail"`
		BillBankID            string `form:"billBankId,omitempty"`
	}

	// CreateBillTransactions ...
	CreateBillTransactions struct {
		BillCode          string `form:"billCode"`
		BillPaymentStatus int    `form:"billpaymentStatus"`
	}

	// BillTransactionsResponse is the response struct from calling
	BillTransactionsResponse struct {
		BillName             string `json:"billName"`
		BillDescription      string `json:"billDescription"`
		BillTo               string `json:"billTo"`
		BillEmail            string `json:"billEmail"`
		BillPhone            string `json:"billPhone"`
		BillStatus           string `json:"billStatus"`
		BillPermalink        string `json:"billPermalink"`
		CategoryCode         string `json:"categoryCode"`
		CategoryName         string `json:"categoryName"`
		Username             string `json:"userName"`
		BillPaymentStatus    string `json:"billPaymentStatus"`
		BillPaymentAmount    string `json:"billPaymentAmount"`
		BillPaymentInvoiceNo string `json:"billPaymentInvoiceNo"`
	}

	// TransactionResp ...
	TransactionResp struct {
		BillName string `json:"billName"`
	}

	// BillResponse is the API response from toyyibpay createBill
	// CreateBillResponse struct {
	// 	BillResponse
	// }

	BillResponse struct {
		BillCode string
	}

	APIBillResponse []BillResponse

	RunBillResponse struct {
		Body *string
	}
)

// APIResponse ...
type APIResponse struct {
	BillCode string
}

// ResponseStruct ...
type ResponseStruct map[string]interface{}

// CreateSingleBill will create a single bill and return the unique code id
func (c *Client) CreateSingleBill(billParams CreateBillParams) (string, error) {

	var err error
	billResponse := &APIBillResponse{}
	billParams.UserSecretKey = c.UserSecretKey
	req, err := c.NewRequest("createBill", billParams)

	if err != nil {
		return "", err
	}

	err = c.CallWithJSONResponse(req, billResponse)

	if err != nil {
		return "", err
	}
	// billResponse

	return (*billResponse)[0].BillCode, nil
}

// GetTransactions ...
func (c *Client) GetTransactions(transactionParams interface{}) ([]BillTransactionsResponse, error) {
	var err error
	billResponse := []BillTransactionsResponse{}
	req, err := c.NewRequest("getBillTransactions", transactionParams)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = c.CallWithJSONResponse(req, &billResponse)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("billResponse ", billResponse)

	return billResponse, nil
}

// RunBill ...
func (c *Client) RunBill(params RunBillParams) (string, error) {

	var err error
	resp := &RunBillResponse{}
	params.UserSecretKey = c.UserSecretKey
	req, err := c.NewRequest("runBill", params)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	err = c.CallWithHTMLResponse(req, resp)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// fmt.Printf("resp: %s", resp.Body)
	return *resp.Body, nil
}

// CreateMulti ...
func (c *Client) CreateMulti() {

}

// GetAll ...
func (c *Client) GetAll() {

}
