package token

type NewAccessResponse struct {
	AccessToken      string `json:"access_token,omitempty"`
	ExpiresIn        uint32 `json:"expires_in,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	Scope            string `json:"scope,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}
