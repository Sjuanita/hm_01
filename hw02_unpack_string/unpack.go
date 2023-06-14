package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/example/stringutil"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var res strings.Builder
	str = stringutil.Reverse(str)
	num := 0

	for _, val := range str {
		switch {
		case val >= 48 && val <= 57:
			if num > 0 || num == -1 {
				return "", ErrInvalidString
			}

			if val == 48 {
				num = -1
			} else {
				num, _ = strconv.Atoi(string(val))
			}
		default:
			if num > 0 {
				fmt.Fprintf(&res, "%s", strings.Repeat(string(val), num))
			} else if num != -1 {
				fmt.Fprintf(&res, "%s", string(val))
			}
			num = 0
		}
	}

	if num != 0 {
		return "", ErrInvalidString
	}

	return stringutil.Reverse(res.String()), nil
}
