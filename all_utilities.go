package gopal

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const errRequired = "%q is a required field"
const errTooManyChars = "%q permits %d characters; found %d"
const logTooManyChars = errTooManyChars + "; truncating to fit\n"

func checkStr(name string, s *string, leng int, required, doErr bool) error {
	var trimmed = strings.TrimSpace(*s)
	*s = trimmed
	if required && len(trimmed) == 0 {
		return fmt.Errorf(errRequired, name)
	}
	if len(trimmed) > leng {
		if doErr {
			return fmt.Errorf(errTooManyChars, name, leng, len(trimmed))
		} else {
			log.Printf(logTooManyChars, name, leng, len(trimmed))
			*s = trimmed[0:leng]
		}
	}
	return nil
}

func roundTwoDecimalPlaces(nn float64) float64 {
	nn, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", nn), 64)
	return nn
}

const zeroWarning = "WARNING: %q was given a 0 (zero) quantity\n"

// Validate positive float, no more than 7 digits to the left of the decimal
func checkFloat7_10(name string, n *float64, warnZero bool) error {
	var nn = *n
	if nn < 0 {
		return fmt.Errorf("%q must not be negative number", name)
	}
	if nn == 0 && warnZero {
		log.Printf(zeroWarning, name)
	}
	if nn > 9999999.99 {
		return fmt.Errorf(
			"%q permits no more than 7 digits to the left of the decimal", name)
	}

	*n = roundTwoDecimalPlaces(nn)
	return nil
}

// Validate positive int, no more than 10 digits
func checkInt10(name string, n int64, warnZero bool) error {
	if n < 0 {
		return fmt.Errorf("%q must not be negative number", name)
	}
	if n == 0 && warnZero {
		log.Printf(zeroWarning, name)
	}
	if n > 9999999999 {
		return fmt.Errorf("%q permits no more than 10 digits", name)
	}
	return nil
}
