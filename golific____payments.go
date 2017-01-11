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

CreditCardItem struct

******************************/

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *CreditCardItem) JSONEncode(encoder *gJson.Encoder) bool {
	if self == nil {
		return encoder.EncodeNull(false)
	}

	encoder.WriteRawByte('{')
	var first = true

	if true {
		var d interface{} = self.Currency

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.Currency).Kind() == reflect.Struct {
				d = &self.Currency
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("currency", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.Quantity

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Quantity).Kind() == reflect.Struct {
				d = &self.Quantity
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("quantity", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.Name

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Name).Kind() == reflect.Struct {
				d = &self.Name
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("name", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.Price

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Price).Kind() == reflect.Struct {
				d = &self.Price
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("price", d, first, false) && first
		}
	}

	if len(self.Sku) != 0 {
		var d interface{} = self.Sku

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Sku).Kind() == reflect.Struct {
				d = &self.Sku
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("sku", d, first, true) && first
		}
	}

	if len(self.Url) != 0 {
		var d interface{} = self.Url

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Url).Kind() == reflect.Struct {
				d = &self.Url
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("url", d, first, true) && first
		}
	}

	encoder.WriteRawByte('}')

	return true || !first
}

func (self *CreditCardItem) MarshalJSON() ([]byte, error) {
	var encoder gJson.Encoder
	self.JSONEncode(&encoder)
	return encoder.Bytes(), nil
}

func (self *CreditCardItem) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	// First unmarshal using the default unmarshaler. The temp type is so that
	// this method is not called recursively.
	type temp *CreditCardItem
	if err := json.Unmarshal(j, temp(self)); err != nil {
		return err
	}

	return nil
}

/*****************************

PaypalItem struct

******************************/

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *PaypalItem) JSONEncode(encoder *gJson.Encoder) bool {
	if self == nil {
		return encoder.EncodeNull(false)
	}

	encoder.WriteRawByte('{')
	var first = true

	if true {
		var d interface{} = self.Currency

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.Currency).Kind() == reflect.Struct {
				d = &self.Currency
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("currency", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.Quantity

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Quantity).Kind() == reflect.Struct {
				d = &self.Quantity
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("quantity", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.Name

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Name).Kind() == reflect.Struct {
				d = &self.Name
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("name", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.Price

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Price).Kind() == reflect.Struct {
				d = &self.Price
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("price", d, first, false) && first
		}
	}

	if len(self.Sku) != 0 {
		var d interface{} = self.Sku

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Sku).Kind() == reflect.Struct {
				d = &self.Sku
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("sku", d, first, true) && first
		}
	}

	if len(self.Url) != 0 {
		var d interface{} = self.Url

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Url).Kind() == reflect.Struct {
				d = &self.Url
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("url", d, first, true) && first
		}
	}

	if self.Tax != 0 {
		var d interface{} = self.Tax

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Tax).Kind() == reflect.Struct {
				d = &self.Tax
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("tax", d, first, true) && first
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

	encoder.WriteRawByte('}')

	return true || !first
}

func (self *PaypalItem) MarshalJSON() ([]byte, error) {
	var encoder gJson.Encoder
	self.JSONEncode(&encoder)
	return encoder.Bytes(), nil
}

func (self *PaypalItem) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	// First unmarshal using the default unmarshaler. The temp type is so that
	// this method is not called recursively.
	type temp *PaypalItem
	if err := json.Unmarshal(j, temp(self)); err != nil {
		return err
	}

	return nil
}

/*****************************

_shared struct

******************************/

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *_shared) JSONEncode(encoder *gJson.Encoder) bool {
	if self == nil {
		return encoder.EncodeNull(false)
	}

	encoder.WriteRawByte('{')
	var first = true

	if je, ok := interface{}(self.connection).(gJson.JSONEncodable); ok {
		first = !encoder.EmbedEncodedStruct(je, first) && first
	} else {
		first = !encoder.EmbedMarshaledStruct(self.connection, first) && first
	}

	if len(self.Id) != 0 {
		var d interface{} = self.Id

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Id).Kind() == reflect.Struct {
				d = &self.Id
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("id", d, first, true) && first
		}
	}

	if z, ok := interface{}(self.CreateTime).(gJson.Zeroable); !ok || !z.IsZero() {
		var d interface{} = self.CreateTime

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.CreateTime).Kind() == reflect.Struct {
				d = &self.CreateTime
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("create_time", d, first, true) && first
		}
	}

	if z, ok := interface{}(self.UpdateTime).(gJson.Zeroable); !ok || !z.IsZero() {
		var d interface{} = self.UpdateTime

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.UpdateTime).Kind() == reflect.Struct {
				d = &self.UpdateTime
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("update_time", d, first, true) && first
		}
	}

	if z, ok := interface{}(self.State).(gJson.Zeroable); !ok || !z.IsZero() {
		var d interface{} = self.State

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.State).Kind() == reflect.Struct {
				d = &self.State
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("state", d, first, true) && first
		}
	}

	if len(self.ParentPayment) != 0 {
		var d interface{} = self.ParentPayment

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.ParentPayment).Kind() == reflect.Struct {
				d = &self.ParentPayment
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("parent_payment", d, first, true) && first
		}
	}

	if z, ok := interface{}(self.Links).(gJson.Zeroable); !ok || !z.IsZero() {
		var d interface{} = self.Links

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.Links).Kind() == reflect.Struct {
				d = &self.Links
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("links", d, first, true) && first
		}
	}

	if je, ok := interface{}(self.identity_error).(gJson.JSONEncodable); ok {
		first = !encoder.EmbedEncodedStruct(je, first) && first
	} else {
		first = !encoder.EmbedMarshaledStruct(self.identity_error, first) && first
	}

	encoder.WriteRawByte('}')

	return true || !first
}

func (self *_shared) MarshalJSON() ([]byte, error) {
	var encoder gJson.Encoder
	self.JSONEncode(&encoder)
	return encoder.Bytes(), nil
}

func (self *_shared) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	// First unmarshal using the default unmarshaler. The temp type is so that
	// this method is not called recursively.
	type temp *_shared
	if err := json.Unmarshal(j, temp(self)); err != nil {
		return err
	}

	return nil
}

/*****************************

amount struct

******************************/

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *amount) JSONEncode(encoder *gJson.Encoder) bool {
	if self == nil {
		return encoder.EncodeNull(false)
	}

	encoder.WriteRawByte('{')
	var first = true

	if true {
		var d interface{} = self.Currency

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.Currency).Kind() == reflect.Struct {
				d = &self.Currency
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("currency", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.Total

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Total).Kind() == reflect.Struct {
				d = &self.Total
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("total", d, first, false) && first
		}
	}

	if z, ok := interface{}(self.Details).(gJson.Zeroable); !ok || !z.IsZero() {
		var d interface{} = self.Details

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.Details).Kind() == reflect.Struct {
				d = &self.Details
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("details", d, first, true) && first
		}
	}

	encoder.WriteRawByte('}')

	return true || !first
}

func (self *amount) MarshalJSON() ([]byte, error) {
	var encoder gJson.Encoder
	self.JSONEncode(&encoder)
	return encoder.Bytes(), nil
}

func (self *amount) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	// First unmarshal using the default unmarshaler. The temp type is so that
	// this method is not called recursively.
	type temp *amount
	if err := json.Unmarshal(j, temp(self)); err != nil {
		return err
	}

	return nil
}

/*****************************

details struct

******************************/

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *details) JSONEncode(encoder *gJson.Encoder) bool {
	if self == nil {
		return encoder.EncodeNull(false)
	}

	encoder.WriteRawByte('{')
	var first = true

	if self.Subtotal != 0 {
		var d interface{} = self.Subtotal

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Subtotal).Kind() == reflect.Struct {
				d = &self.Subtotal
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("subtotal", d, first, true) && first
		}
	}

	if self.Tax != 0 {
		var d interface{} = self.Tax

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Tax).Kind() == reflect.Struct {
				d = &self.Tax
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("tax", d, first, true) && first
		}
	}

	if self.Shipping != 0 {
		var d interface{} = self.Shipping

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Shipping).Kind() == reflect.Struct {
				d = &self.Shipping
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("shipping", d, first, true) && first
		}
	}

	if self.HandlingFee != 0 {
		var d interface{} = self.HandlingFee

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.HandlingFee).Kind() == reflect.Struct {
				d = &self.HandlingFee
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("handling_fee", d, first, true) && first
		}
	}

	if self.Insurance != 0 {
		var d interface{} = self.Insurance

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Insurance).Kind() == reflect.Struct {
				d = &self.Insurance
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("insurance", d, first, true) && first
		}
	}

	if self.ShippingDiscount != 0 {
		var d interface{} = self.ShippingDiscount

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.ShippingDiscount).Kind() == reflect.Struct {
				d = &self.ShippingDiscount
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("shipping_discount", d, first, true) && first
		}
	}

	encoder.WriteRawByte('}')

	return true || !first
}

func (self *details) MarshalJSON() ([]byte, error) {
	var encoder gJson.Encoder
	self.JSONEncode(&encoder)
	return encoder.Bytes(), nil
}

func (self *details) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	// First unmarshal using the default unmarshaler. The temp type is so that
	// this method is not called recursively.
	type temp *details
	if err := json.Unmarshal(j, temp(self)); err != nil {
		return err
	}

	return nil
}

/*****************************

link struct

******************************/

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *link) JSONEncode(encoder *gJson.Encoder) bool {
	if self == nil {
		return encoder.EncodeNull(false)
	}

	encoder.WriteRawByte('{')
	var first = true

	if len(self.Href) != 0 {
		var d interface{} = self.Href

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Href).Kind() == reflect.Struct {
				d = &self.Href
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("href", d, first, true) && first
		}
	}

	if z, ok := interface{}(self.Rel).(gJson.Zeroable); !ok || !z.IsZero() {
		var d interface{} = self.Rel

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.Rel).Kind() == reflect.Struct {
				d = &self.Rel
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("rel", d, first, true) && first
		}
	}

	if len(self.Method) != 0 {
		var d interface{} = self.Method

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Method).Kind() == reflect.Struct {
				d = &self.Method
			}
		}

		var doEncode = true
		if true { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("method", d, first, true) && first
		}
	}

	encoder.WriteRawByte('}')

	return true || !first
}

func (self *link) MarshalJSON() ([]byte, error) {
	var encoder gJson.Encoder
	self.JSONEncode(&encoder)
	return encoder.Bytes(), nil
}

func (self *link) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	// First unmarshal using the default unmarshaler. The temp type is so that
	// this method is not called recursively.
	type temp *link
	if err := json.Unmarshal(j, temp(self)); err != nil {
		return err
	}

	return nil
}

/*****************************

currency struct

******************************/

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *currency) JSONEncode(encoder *gJson.Encoder) bool {
	if self == nil {
		return encoder.EncodeNull(false)
	}

	encoder.WriteRawByte('{')
	var first = true

	if true {
		var d interface{} = self.Currency

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Currency).Kind() == reflect.Struct {
				d = &self.Currency
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("currency", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.Value

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Value).Kind() == reflect.Struct {
				d = &self.Value
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("value", d, first, false) && first
		}
	}

	encoder.WriteRawByte('}')

	return true || !first
}

func (self *currency) MarshalJSON() ([]byte, error) {
	var encoder gJson.Encoder
	self.JSONEncode(&encoder)
	return encoder.Bytes(), nil
}

func (self *currency) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	// First unmarshal using the default unmarshaler. The temp type is so that
	// this method is not called recursively.
	type temp *currency
	if err := json.Unmarshal(j, temp(self)); err != nil {
		return err
	}

	return nil
}

/*****************************

fmfDetails struct

******************************/

// JSONEncode implements part of Golific's JSONEncodable interface.
func (self *fmfDetails) JSONEncode(encoder *gJson.Encoder) bool {
	if self == nil {
		return encoder.EncodeNull(false)
	}

	encoder.WriteRawByte('{')
	var first = true

	if true {
		var d interface{} = self.FilterType

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.FilterType).Kind() == reflect.Struct {
				d = &self.FilterType
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("filter_type", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.FilterID

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if true && reflect.ValueOf(self.FilterID).Kind() == reflect.Struct {
				d = &self.FilterID
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("filter_id", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.Name

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Name).Kind() == reflect.Struct {
				d = &self.Name
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("name", d, first, false) && first
		}
	}

	if true {
		var d interface{} = self.Description

		if _, ok := d.(gJson.JSONEncodable); !ok {
			if false && reflect.ValueOf(self.Description).Kind() == reflect.Struct {
				d = &self.Description
			}
		}

		var doEncode = true
		if false { // has omitempty?
			if zer, okCanZero := d.(gJson.Zeroable); okCanZero {
				doEncode = !zer.IsZero()
			}
		}

		if doEncode {
			first = !encoder.EncodeKeyVal("description", d, first, false) && first
		}
	}

	encoder.WriteRawByte('}')

	return true || !first
}

func (self *fmfDetails) MarshalJSON() ([]byte, error) {
	var encoder gJson.Encoder
	self.JSONEncode(&encoder)
	return encoder.Bytes(), nil
}

func (self *fmfDetails) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	// First unmarshal using the default unmarshaler. The temp type is so that
	// this method is not called recursively.
	type temp *fmfDetails
	if err := json.Unmarshal(j, temp(self)); err != nil {
		return err
	}

	return nil
}
