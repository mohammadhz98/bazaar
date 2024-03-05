package bazaar
/*
bazaar used to create a complete bazaar api handler.
it not implement any handler and just contains them.
it contains all handlers of dependent parts and also contains configs of them.
*/

import (
	"github.com/mohammadhz98/bazaar/concrete/token"
	"github.com/mohammadhz98/bazaar/concrete/verify"
	"github.com/mohammadhz98/bazaar/conf"
	iface_token "github.com/mohammadhz98/bazaar/iface/token"
	iface_verify "github.com/mohammadhz98/bazaar/iface/verify"
)

// Bazaar handles all APIs and is the main part of package
// users create and use this struct in their app
// all other part of package is accessible by this.
type Bazaar struct {
	conf   conf.Bazaar
	token  iface_token.Token
	verify iface_verify.Verify
}

// NewBazaar create a bazaar api handler
func NewBazaar(cfg conf.Bazaar) (bazaar *Bazaar, err error) {
	bazaar = &Bazaar{}
	
	// initialize config
	bazaar.conf = cfg

	// initialize token api handler
	if bazaar.token, err = token.NewToken(bazaar); err != nil {
		return
	}

	// initialize verify api handler
	if bazaar.verify, err = verify.NewVerify(bazaar); err != nil {
		return
	}

	return
}

// Conf return bazaar handler configs
func (bazaar *Bazaar) Conf() conf.Bazaar {
	return bazaar.conf
}

// Token return handler for token APIs
func (bazaar *Bazaar) Token() iface_token.Token {
	return bazaar.token
}

// Verify return handler for verify APIs
func (bazaar *Bazaar) Verify() iface_verify.Verify {
	return bazaar.verify
}
