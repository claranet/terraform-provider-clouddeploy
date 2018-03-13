package ghost

import (
	"encoding/base64"
)

func StrToB64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
