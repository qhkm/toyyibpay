package toyyibpay

import (
	"fmt"
	"net/url"
	"os"
)

const baseURLDev string = "https://dev.toyyibpay.com/index.php/api/"
const baseURLProd string = "https://toyyibpay.com/index.php/api/"

// CheckErr ...
func CheckErr(err error) error {
	if err != nil {
		// :TODO use errors package to handle error more elegantly
		return err
	}
	return nil
}

// GetAPIPath ...
func (c *Client) GetAPIPath(action string) (string, error) {
	baseURL := c.BaseURL

	if baseURL == "" {
		if getEnv() == "production" {
			baseURL = baseURLProd
		} else {
			baseURL = baseURLDev
		}
	}

	u, err := url.Parse(baseURL)

	if err != nil {
		return "", err
	}
	if u.Host[len(u.Host)-1:] != "/" {
		baseURL = fmt.Sprintf(baseURL + "/")
	}
	return fmt.Sprintf("%s%s", u.String(), action), nil
}

// GetEnv ...
func getEnv() string {
	return os.Getenv("APP_ENV")
}
