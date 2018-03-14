package ghost

import (
	"encoding/base64"
	"fmt"
	"regexp"
)

func StrToB64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func B64ToStr(data string) string {
	str, _ := base64.StdEncoding.DecodeString(data)

	return string(str)
}
<<<<<<< 49067eb1fd7fc52fb01240f0939e715a6f7e73bf

func MatchesRegexp(exp string) func(v interface{}, k string) (ws []string, errors []error) {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)
		if !regexp.MustCompile(exp).MatchString(value) {
			errors = append(errors, fmt.Errorf("%q must match %s", k, exp))
		}
		return
	}
}
=======
>>>>>>> (TER-239) Update app read feature: read modules and features
