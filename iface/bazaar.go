package iface

import (
	"github.com/mohammadhz98/bazaar/conf"
	"github.com/mohammadhz98/bazaar/iface/token"
	"github.com/mohammadhz98/bazaar/iface/verify"
)

type Bazaar interface {
	Conf() conf.Bazaar
	Token() token.Token
	Verify() verify.Verify
}
