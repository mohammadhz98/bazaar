package bazaar

import (
	"github.com/mohammadhz98/bazaar/concrete/token"
	"github.com/mohammadhz98/bazaar/concrete/verify"
	"github.com/mohammadhz98/bazaar/conf"
	iface_token "github.com/mohammadhz98/bazaar/iface/token"
	iface_verify "github.com/mohammadhz98/bazaar/iface/verify"
)

type Bazaar struct {
	conf   conf.Bazaar
	token  iface_token.Token
	verify iface_verify.Verify
}

func NewBazaar(cfg conf.Bazaar) (bazaar *Bazaar, err error) {
	bazaar = &Bazaar{}
	bazaar.conf = cfg

	if bazaar.token, err = token.NewToken(bazaar); err != nil {
		return
	}

	if bazaar.verify, err = verify.NewVerify(bazaar); err != nil {
		return
	}

	return
}

func (bazaar *Bazaar) Conf() conf.Bazaar {
	return bazaar.conf
}

func (bazaar *Bazaar) Token() iface_token.Token {
	return bazaar.token
}

func (bazaar *Bazaar) Verify() iface_verify.Verify {
	return bazaar.verify
}
