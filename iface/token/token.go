package token

type Token interface {
	Access() (string, error)
}