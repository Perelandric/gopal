package gopal

import (
	"fmt"
	"strconv"
)

func make10CharAmount(amt float64) (string, error) {
	if amt < 0 {
		return "", fmt.Errorf("Amount may not be a negative number. Found: %f\n",
			amt)
	}

	var t string
	if (amt - float64(int(amt))) < .005 { // If no fraction, drop the decimal
		t = strconv.FormatInt(int64(amt), 10)
	} else {
		t = fmt.Sprintf("%.2f", amt)
	}

	if len(t) > 10 {
		return "", fmt.Errorf("Amount Total allows 10 chars max. Found: %q\n", t)
	}
	return t, nil
}
