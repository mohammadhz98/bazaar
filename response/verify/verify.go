package verify

type PurchaseVerifyResponse struct {
	ConsumptionState uint8  `json:"consumptionState,omitempty"`
	PurchaseState    uint8  `json:"purchaseState,omitempty"`
	Kind             string `json:"kind,omitempty"`
	DeveloperPayload string `json:"developerPayload,omitempty"`
	PurchaseTime     int64  `json:"purchaseTime,omitempty"` // purchase time in millisecond
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}
