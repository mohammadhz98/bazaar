package verify

/*
verify used to create a bazaar verify api handler.
it implement handler for verify related APIs such as purchase verify.
*/

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohammadhz98/bazaar/iface"
	response_verify "github.com/mohammadhz98/bazaar/response/verify"
)

var (
	verifyURLFormat string = "https://pardakht.cafebazaar.ir/devapi/v2/api/validate/%s/inapp/%s/purchases/%s/?access_token=%s"
)

// Verify used for bazaar verification APIs
type Verify struct {
	bazaar iface.Bazaar
}

func NewVerify(bazaar iface.Bazaar) (verify *Verify, err error) {
	verify = &Verify{
		bazaar: bazaar,
	}

	return
}

// Purchase will verify a purchase token for specific product
func (v *Verify) Purchase(packageName, productID, purchaseToken string) (resp response_verify.PurchaseVerifyResponse, err error) {
	accessToken, err := v.bazaar.Token().Access()
	if err != nil {
		return
	}

	verifyURL := fmt.Sprintf(verifyURLFormat, packageName, productID, purchaseToken, accessToken)

	res, err := http.Get(verifyURL)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return
	}

	return
}
