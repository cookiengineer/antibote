package constants

import _ "embed"
import "strings"

//go:embed Token.env
var embedded_token []byte

var Token string

func init() {
	Token = strings.TrimSpace(string(embedded_token))
}
