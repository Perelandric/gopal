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

Refund struct

******************************/

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *Refund) JSONEncode(encoder *gJson.Encoder) bool {
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

	if len(self.Description) != 0 {
		var d interface{} = self.Description

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Description).Kind() == reflect.Struct {
				d = &self.Description
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("description", d, first, true) && first
		}
	}

	if len(self.Reason) != 0 {
		var d interface{} = self.Reason

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Reason).Kind() == reflect.Struct {
				d = &self.Reason
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("reason", d, first, true) && first
		}
	}

	if len(self.SaleId) != 0 {
		var d interface{} = self.SaleId

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.SaleId).Kind() == reflect.Struct {
				d = &self.SaleId
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("sale_id", d, first, true) && first
		}
	}

	if len(self.CaptureId) != 0 {
		var d interface{} = self.CaptureId

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.CaptureId).Kind() == reflect.Struct {
				d = &self.CaptureId
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("capture_id", d, first, true) && first
		}
	}

	encoder.WriteRawByte('}')

	return true || !first
}

func (self *Refund) MarshalJSON() ([]byte, error) {
	var encoder gJson.Encoder
	self.JSONEncode(&encoder)
	return encoder.Bytes(), nil
}

func (self *Refund) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	// First unmarshal using the default unmarshaler. The temp type is so that
	// this method is not called recursively.
	type temp *Refund
	if err := json.Unmarshal(j, temp(self)); err != nil {
		return err
	}

	return nil
}