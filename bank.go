package toyyibpay

type IBank interface {
	getBank() Bank
	getBankFPX()
}

type Bank struct {
	Id     string `json:"id"`
	Bank   string `json:"bank"`
	Status string `json:"status"`
}

type BankList []Bank

type BankFPX struct {
	Code string `json:"CODE"`
	Name string `json:"NAME"`
}

// BankFPXList ...
type BankFPXList []BankFPX

// GetBankList ...
func (c *Client) GetBankList() (BankList, error) {
	var err error
	bankResponse := BankList{}
	req, err := c.NewRequest("getBank", nil)

	if err != nil {
		return nil, err
	}

	err = c.CallWithJSONResponse(req, &bankResponse)

	if err != nil {
		return nil, err
	}

	return bankResponse, nil
}

func (c *Client) GetFPXCode() (BankFPXList, error) {

	var err error
	bankFPXResponse := BankFPXList{}
	req, err := c.NewRequest("getBankFPX", nil)

	if err != nil {
		return nil, err
	}

	err = c.CallWithJSONResponse(req, &bankFPXResponse)

	if err != nil {
		return nil, err
	}

	return bankFPXResponse, nil
}
