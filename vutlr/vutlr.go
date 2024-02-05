package vutlr

type Vutlr struct {
	APIKey string
}

func New() *Vutlr {
	return &Vutlr{}
}

func (v *Vutlr) SetAPIKey(key string) {
	v.APIKey = key
}
