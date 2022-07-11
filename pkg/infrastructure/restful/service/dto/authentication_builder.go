package dto

import "github.com/mitchellh/mapstructure"

type LoginResponse struct {
	Token     string `json:"token"`
	Username  string `json:"username"`
	Expired   string `json:"expired"`
	JenisUser string `json:"jenis_user"`
}
type AuthenticationBuilder struct {
}

func (b *AuthenticationBuilder) LoginResponse(in interface{}) (interface{}, error) {
	var out *LoginResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
