/****************************************************************************
	This file was generated by Golific.

	Do not edit this file. If you do, your changes will be overwritten the next
	time 'generate' is invoked.
******************************************************************************/

package gopal

import (
	"Golific/gJson"
	"encoding/json"
	"reflect"
)

/*****************************

Capture struct

******************************/

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *Capture) JSONEncode(encoder *gJson.Encoder) bool {
	if self == nil {
		return encoder.EncodeNull(false)
	}

	encoder.WriteRawByte('{')
	var first = true

	if je, ok := interface{}(self._shared).(gJson.JSONEncodable); ok {
		first = !encoder.EmbedEncodedStruct(je, first) && first
	} else {
		first = !encoder.EmbedMarshaledStruct(self._shared, first) && first
	}

	if true {
		var d interface{} = self.Amount

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.Amount).Kind() == reflect.Struct {
				d = &self.Amount
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("amount", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.TransactionFee

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.TransactionFee).Kind() == reflect.Struct {
				d = &self.TransactionFee
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("transaction_fee", d, first, false) && first
		}
	}

	if !self.IsFinalCapture {
		var d interface{} = self.IsFinalCapture

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.IsFinalCapture).Kind() == reflect.Struct {
				d = &self.IsFinalCapture
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
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
	var encoder gJson.Encoder
	self.JSONEncode(&encoder)
	return encoder.Bytes(), nil
}

func (self *Capture) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	// First unmarshal using the default unmarshaler. The temp type is so that
	// this method is not called recursively.
	type temp *Capture
	if err := json.Unmarshal(j, temp(self)); err != nil {
		return err
	}

	return nil
}
