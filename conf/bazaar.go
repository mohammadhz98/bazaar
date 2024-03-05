package conf

import conf_token "github.com/mohammadhz98/bazaar/conf/token"

// Bazaar is config needed for bazaar APIs
//
// all of configs must provide from bazaar dashboard
type Bazaar struct {
	Token conf_token.Token
}
