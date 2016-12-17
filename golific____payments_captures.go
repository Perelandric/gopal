/****************************************************************************
	This file was generated by Golific.

	Do not edit this file. If you do, your changes will be overwritten the next
	time 'generate' is invoked.
******************************************************************************/

package gopal

import (
	"Golific/gJson"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

/*****************************

Capture struct

******************************/

// State values are: pending, completed, refunded, partially_refunded
type Capture struct {
	private private_1v66jzy7wsqve
	_shared
	TransactionFee currency `json:"transaction_fee"`
	IsFinalCapture bool     `json:"is_final_capture,omitempty"`
}
type private_1v66jzy7wsqve struct {
	Amount amount `json:"amount"`
}

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *private_1v66jzy7wsqve) JSONEncode(encoder *gJson.Encoder) bool {
	var first = true

	if true {
		var d interface{}

		if true && reflect.ValueOf(self.Amount).Kind() == reflect.Struct {
			d = &self.Amount
		} else {
			d = self.Amount
		}

		var doEncode = true
		if false { // has omitempty?
			if eli, okCanElide := d.(gJson.Elidable); okCanElide {
				doEncode = !eli.CanElide()

			} else if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("amount", d, first, false) && first
		}
	}

	return !first
}

type json_1v66jzy7wsqve struct {
	*private_1v66jzy7wsqve

	_shared
	TransactionFee currency `json:"transaction_fee"`
	IsFinalCapture bool     `json:"is_final_capture,omitempty"`
}

func (self *Capture) Amount() amount {
	return self.private.Amount
}

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *Capture) JSONEncode(encoder *gJson.Encoder) bool {
	if self == nil {
		return encoder.EncodeNull(false)
	}

	encoder.WriteRawByte('{')

	// Encodes only the fields of the struct, without curly braces
	var first = !self.private.JSONEncode(encoder)

	if je, ok := interface{}(self._shared).(gJson.JSONEncodable); ok {
		first = !encoder.EmbedEncodedStruct(je, first) && first
	} else {
		first = !encoder.EmbedMarshaledStruct(self._shared, first) && first
	}

	if true {
		var d interface{}

		if true && reflect.ValueOf(self.TransactionFee).Kind() == reflect.Struct {
			d = &self.TransactionFee
		} else {
			d = self.TransactionFee
		}

		var doEncode = true
		if false { // has omitempty?
			if eli, okCanElide := d.(gJson.Elidable); okCanElide {
				doEncode = !eli.CanElide()

			} else if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("transaction_fee", d, first, false) && first
		}
	}

	if !self.IsFinalCapture {
		var d interface{}

		if false && reflect.ValueOf(self.IsFinalCapture).Kind() == reflect.Struct {
			d = &self.IsFinalCapture
		} else {
			d = self.IsFinalCapture
		}

		var doEncode = true
		if true { // has omitempty?
			if eli, okCanElide := d.(gJson.Elidable); okCanElide {
				doEncode = !eli.CanElide()

			} else if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("is_final_capture", d, first, true) && first
		}
	}

	encoder.WriteRawByte('}')

	return true || !first
}

func (self *Capture) MarshalJSON() ([]byte, error) {
	return json.Marshal(json_1v66jzy7wsqve{

		&self.private,
		self._shared,
		self.TransactionFee,
		self.IsFinalCapture,
	})
}

func (self *Capture) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	// For every property found, perform a separate UnmarshalJSON operation. This
	// prevents overwrite of values in 'self' where properties are absent.
	m := make(map[string]json.RawMessage)

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// JSON key comparisons are case-insensitive
	for k, v := range m {
		m[strings.ToLower(k)] = v
	}

	var data json.RawMessage
	var ok bool
	if data, ok = m["amount"]; ok {
		var temp struct {
			Amount amount `json:"amount"`
		}
		data = append(append([]byte("{ \"amount\":"), data...), '}')

		if err = json.Unmarshal(data, &temp); err != nil {
			return fmt.Errorf(
				"Field: %s, Error: %s", "amount", err.Error(),
			)
		}

		self.private.Amount = temp.Amount
	}

	if data, ok = m["transaction_fee"]; ok {
		var temp struct {
			TransactionFee currency `json:"transaction_fee"`
		}
		data = append(append([]byte("{ \"transaction_fee\":"), data...), '}')

		if err = json.Unmarshal(data, &temp); err != nil {
			return fmt.Errorf(
				"Field: %s, Error: %s", "transaction_fee", err.Error(),
			)
		}

		self.TransactionFee = temp.TransactionFee
	}

	if data, ok = m["is_final_capture"]; ok {
		var temp struct {
			IsFinalCapture bool `json:"is_final_capture,omitempty"`
		}
		data = append(append([]byte("{ \"is_final_capture\":"), data...), '}')

		if err = json.Unmarshal(data, &temp); err != nil {
			return fmt.Errorf(
				"Field: %s, Error: %s", "is_final_capture", err.Error(),
			)
		}

		self.IsFinalCapture = temp.IsFinalCapture
	}
	return nil
}
