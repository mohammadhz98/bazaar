package verify

import response_verify "github.com/mohammadhz98/bazaar/response/verify"

type Verify interface {
	Purchase(packageName, productID, purchaseToken string) (response_verify.PurchaseVerifyResponse, error)
}
