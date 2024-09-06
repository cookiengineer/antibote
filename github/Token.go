package github

import _ "embed"

//go:embed Token.env
var embedded_token []string

var Token string

func init() {
	Token = string(embedded_token)
}
