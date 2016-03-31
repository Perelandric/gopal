/****************************************************************************
	This file was generated by Golific.

	Do not edit this file. If you do, your changes will be overwritten the next
	time 'generate' is invoked.
******************************************************************************/

package gopal

import (
	"encoding/json"
)

/*****************************

_shared struct

******************************/

type _shared struct {
	private private_bqmm7d5xpizm
	*connection
	*identity_error
}

type private_bqmm7d5xpizm struct {
	Id            string    `json:"id,omitempty"`
	CreateTime    dateTime  `json:"create_time,omitempty"`
	UpdateTime    dateTime  `json:"update_time,omitempty"`
	State         StateEnum `json:"state,omitempty"`
	ParentPayment string    `json:"parent_payment,omitempty"`
	Links         links     `json:"links,omitempty"`
}

type json_bqmm7d5xpizm struct {
	*private_bqmm7d5xpizm
	*connection
	*identity_error
}

func (self *_shared) Id() string {
	return self.private.Id
}

func (self *_shared) CreateTime() dateTime {
	return self.private.CreateTime
}

func (self *_shared) UpdateTime() dateTime {
	return self.private.UpdateTime
}

func (self *_shared) State() StateEnum {
	return self.private.State
}

func (self *_shared) ParentPayment() string {
	return self.private.ParentPayment
}

func (self *_shared) MarshalJSON() ([]byte, error) {
	return json.Marshal(json_bqmm7d5xpizm{
		&self.private,
		self.connection,
		self.identity_error,
	})
}

func (self *_shared) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	m := make(map[string]json.RawMessage)

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// For every property found, perform a separate UnmarshalJSON operation. This
	// prevents overwrite of values in 'self' where properties are absent.
	for key, rawMsg := range m {
		// The anon structs in each case are needed for field tags

		switch key {
		case "id":

			var x struct {
				F string `json:"id,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Id = x.F
		case "create_time":

			var x struct {
				F dateTime `json:"create_time,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.CreateTime = x.F
		case "update_time":

			var x struct {
				F dateTime `json:"update_time,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.UpdateTime = x.F
		case "state":

			var x struct {
				F StateEnum `json:"state,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.State = x.F
		case "parent_payment":

			var x struct {
				F string `json:"parent_payment,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ParentPayment = x.F
		case "links":

			var x struct {
				F links `json:"links,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Links = x.F
		default:
			// Ignoring unknown property
		}
	}
	return nil
}

/*****************************

Authorization struct

******************************/

// State items are: pending, authorized, captured, partially_captured, expired,
// voided
type Authorization struct {
	private private_ph2ovsxgja60
	_shared
}

type private_ph2ovsxgja60 struct {
	Amount             amount                 `json:"amount"`
	BillingAgreementId string                 `json:"billing_agreement_id"`
	PaymentMode        paymentModeEnum        `json:"payment_mode"`
	ReasonCode         reasonCodeEnum         `json:"reason_code"`
	ValidUntil         dateTime               `json:"valid_until"`
	ClearingTime       string                 `json:"clearing_time"`
	ProtectionElig     protectionEligEnum     `json:"protection_eligibility"`
	ProtectionEligType protectionEligTypeEnum `json:"protection_eligibility_type"`
	FmfDetails         fmfDetails             `json:"fmf_details"`
}

type json_ph2ovsxgja60 struct {
	*private_ph2ovsxgja60
	_shared
}

func (self *Authorization) Amount() amount {
	return self.private.Amount
}

func (self *Authorization) BillingAgreementId() string {
	return self.private.BillingAgreementId
}

func (self *Authorization) PaymentMode() paymentModeEnum {
	return self.private.PaymentMode
}

func (self *Authorization) ReasonCode() reasonCodeEnum {
	return self.private.ReasonCode
}

func (self *Authorization) ValidUntil() dateTime {
	return self.private.ValidUntil
}

func (self *Authorization) ClearingTime() string {
	return self.private.ClearingTime
}

func (self *Authorization) ProtectionElig() protectionEligEnum {
	return self.private.ProtectionElig
}

func (self *Authorization) ProtectionEligType() protectionEligTypeEnum {
	return self.private.ProtectionEligType
}

func (self *Authorization) FmfDetails() fmfDetails {
	return self.private.FmfDetails
}

func (self *Authorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(json_ph2ovsxgja60{
		&self.private,
		self._shared,
	})
}

func (self *Authorization) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	m := make(map[string]json.RawMessage)

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// For every property found, perform a separate UnmarshalJSON operation. This
	// prevents overwrite of values in 'self' where properties are absent.
	for key, rawMsg := range m {
		// The anon structs in each case are needed for field tags

		switch key {
		case "amount":

			var x struct {
				F amount `json:"amount"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Amount = x.F
		case "billing_agreement_id":

			var x struct {
				F string `json:"billing_agreement_id"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.BillingAgreementId = x.F
		case "payment_mode":

			var x struct {
				F paymentModeEnum `json:"payment_mode"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.PaymentMode = x.F
		case "reason_code":

			var x struct {
				F reasonCodeEnum `json:"reason_code"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ReasonCode = x.F
		case "valid_until":

			var x struct {
				F dateTime `json:"valid_until"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ValidUntil = x.F
		case "clearing_time":

			var x struct {
				F string `json:"clearing_time"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ClearingTime = x.F
		case "protection_eligibility":

			var x struct {
				F protectionEligEnum `json:"protection_eligibility"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ProtectionElig = x.F
		case "protection_eligibility_type":

			var x struct {
				F protectionEligTypeEnum `json:"protection_eligibility_type"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ProtectionEligType = x.F
		case "fmf_details":

			var x struct {
				F fmfDetails `json:"fmf_details"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.FmfDetails = x.F
		default:
			// Ignoring unknown property
		}
	}
	return nil
}

/*****************************

Capture struct

******************************/
func NewCapture() *Capture {
	return &Capture{
		private: private_14luxywpo2i7z{},
	}
}

// State values are: pending, completed, refunded, partially_refunded
type Capture struct {
	private private_14luxywpo2i7z
	_shared
	TransactionFee currency `json:"transaction_fee"`
	IsFinalCapture bool     `json:"is_final_capture,omitempty"`
}

type private_14luxywpo2i7z struct {
	Amount amount `json:"amount"`
}

type json_14luxywpo2i7z struct {
	*private_14luxywpo2i7z
	_shared
	TransactionFee currency `json:"transaction_fee"`
	IsFinalCapture bool     `json:"is_final_capture,omitempty"`
}

func (self *Capture) Amount() amount {
	return self.private.Amount
}

func (self *Capture) MarshalJSON() ([]byte, error) {
	return json.Marshal(json_14luxywpo2i7z{
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

	m := make(map[string]json.RawMessage)

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// For every property found, perform a separate UnmarshalJSON operation. This
	// prevents overwrite of values in 'self' where properties are absent.
	for key, rawMsg := range m {
		// The anon structs in each case are needed for field tags

		switch key {
		case "amount":

			var x struct {
				F amount `json:"amount"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Amount = x.F
		case "transaction_fee":

			var x struct {
				F currency `json:"transaction_fee"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.TransactionFee = x.F

		case "is_final_capture":

			var x struct {
				F bool `json:"is_final_capture,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.IsFinalCapture = x.F

		default:
			// Ignoring unknown property
		}
	}
	return nil
}

/*****************************

Sale struct

******************************/
func NewSale() *Sale {
	return &Sale{
		private: private_gw12g16gjwvf{},
	}
}

// State values are: pending; completed; refunded; partially_refunded
// TODO: PendingReason appears in the old docs under the general Sale object description
// but not under the lower "sale object" definition. The new docs have it
// marked as [DEPRECATED] in one area, but not another.
type Sale struct {
	private private_gw12g16gjwvf
	_shared
	Description      string   `json:"description,omitempty"`
	TransactionFee   currency `json:"transaction_fee"`
	ReceivableAmount currency `json:"receivable_amount"`
}

type private_gw12g16gjwvf struct {
	Amount                    amount                 `json:"amount"`
	PendingReason             pendingReasonEnum      `json:"pending_reason"`
	PaymentMode               paymentModeEnum        `json:"payment_mode"`
	ExchangeRate              string                 `json:"exchange_rate"`
	FmfDetails                fmfDetails             `json:"fmf_details"`
	ReceiptId                 string                 `json:"receipt_id"`
	ReasonCode                reasonCodeEnum         `json:"reason_code"`
	ProtectionEligibility     protectionEligEnum     `json:"protection_eligibility"`
	ProtectionEligibilityType protectionEligTypeEnum `json:"protection_eligibility_type"`
	ClearingTime              string                 `json:"clearing_time"`
	BillingAgreementId        string                 `json:"billing_agreement_id"`
}

type json_gw12g16gjwvf struct {
	*private_gw12g16gjwvf
	_shared
	Description      string   `json:"description,omitempty"`
	TransactionFee   currency `json:"transaction_fee"`
	ReceivableAmount currency `json:"receivable_amount"`
}

func (self *Sale) Amount() amount {
	return self.private.Amount
}

func (self *Sale) PendingReason() pendingReasonEnum {
	return self.private.PendingReason
}

func (self *Sale) PaymentMode() paymentModeEnum {
	return self.private.PaymentMode
}

func (self *Sale) ExchangeRate() string {
	return self.private.ExchangeRate
}

func (self *Sale) FmfDetails() fmfDetails {
	return self.private.FmfDetails
}

func (self *Sale) ReceiptId() string {
	return self.private.ReceiptId
}

func (self *Sale) ReasonCode() reasonCodeEnum {
	return self.private.ReasonCode
}

func (self *Sale) ProtectionEligibility() protectionEligEnum {
	return self.private.ProtectionEligibility
}

func (self *Sale) ProtectionEligibilityType() protectionEligTypeEnum {
	return self.private.ProtectionEligibilityType
}

func (self *Sale) ClearingTime() string {
	return self.private.ClearingTime
}

func (self *Sale) BillingAgreementId() string {
	return self.private.BillingAgreementId
}

func (self *Sale) MarshalJSON() ([]byte, error) {
	return json.Marshal(json_gw12g16gjwvf{
		&self.private,
		self._shared,
		self.Description,
		self.TransactionFee,
		self.ReceivableAmount,
	})
}

func (self *Sale) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	m := make(map[string]json.RawMessage)

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// For every property found, perform a separate UnmarshalJSON operation. This
	// prevents overwrite of values in 'self' where properties are absent.
	for key, rawMsg := range m {
		// The anon structs in each case are needed for field tags

		switch key {
		case "amount":

			var x struct {
				F amount `json:"amount"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Amount = x.F
		case "description":

			var x struct {
				F string `json:"description,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.Description = x.F

		case "transaction_fee":

			var x struct {
				F currency `json:"transaction_fee"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.TransactionFee = x.F

		case "receivable_amount":

			var x struct {
				F currency `json:"receivable_amount"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.ReceivableAmount = x.F

		case "pending_reason":

			var x struct {
				F pendingReasonEnum `json:"pending_reason"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.PendingReason = x.F
		case "payment_mode":

			var x struct {
				F paymentModeEnum `json:"payment_mode"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.PaymentMode = x.F
		case "exchange_rate":

			var x struct {
				F string `json:"exchange_rate"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ExchangeRate = x.F
		case "fmf_details":

			var x struct {
				F fmfDetails `json:"fmf_details"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.FmfDetails = x.F
		case "receipt_id":

			var x struct {
				F string `json:"receipt_id"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ReceiptId = x.F
		case "reason_code":

			var x struct {
				F reasonCodeEnum `json:"reason_code"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ReasonCode = x.F
		case "protection_eligibility":

			var x struct {
				F protectionEligEnum `json:"protection_eligibility"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ProtectionEligibility = x.F
		case "protection_eligibility_type":

			var x struct {
				F protectionEligTypeEnum `json:"protection_eligibility_type"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ProtectionEligibilityType = x.F
		case "clearing_time":

			var x struct {
				F string `json:"clearing_time"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.ClearingTime = x.F
		case "billing_agreement_id":

			var x struct {
				F string `json:"billing_agreement_id"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.BillingAgreementId = x.F
		default:
			// Ignoring unknown property
		}
	}
	return nil
}

/*****************************

Refund struct

******************************/
func NewRefund() *Refund {
	return &Refund{
		private: private_ne29r7kd8wog{},
	}
}

// State items are: pending; completed; failed
type Refund struct {
	private private_ne29r7kd8wog
	_shared
	Description string `json:"description,omitempty"`
	Reason      string `json:"reason,omitempty"`
}

type private_ne29r7kd8wog struct {
	Amount    amount `json:"amount"`
	SaleId    string `json:"sale_id,omitempty"`
	CaptureId string `json:"capture_id,omitempty"`
}

type json_ne29r7kd8wog struct {
	*private_ne29r7kd8wog
	_shared
	Description string `json:"description,omitempty"`
	Reason      string `json:"reason,omitempty"`
}

func (self *Refund) Amount() amount {
	return self.private.Amount
}

func (self *Refund) SaleId() string {
	return self.private.SaleId
}

func (self *Refund) CaptureId() string {
	return self.private.CaptureId
}

func (self *Refund) MarshalJSON() ([]byte, error) {
	return json.Marshal(json_ne29r7kd8wog{
		&self.private,
		self._shared,
		self.Description,
		self.Reason,
	})
}

func (self *Refund) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	m := make(map[string]json.RawMessage)

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// For every property found, perform a separate UnmarshalJSON operation. This
	// prevents overwrite of values in 'self' where properties are absent.
	for key, rawMsg := range m {
		// The anon structs in each case are needed for field tags

		switch key {
		case "amount":

			var x struct {
				F amount `json:"amount"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Amount = x.F
		case "description":

			var x struct {
				F string `json:"description,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.Description = x.F

		case "reason":

			var x struct {
				F string `json:"reason,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.Reason = x.F

		case "sale_id":

			var x struct {
				F string `json:"sale_id,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.SaleId = x.F
		case "capture_id":

			var x struct {
				F string `json:"capture_id,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.CaptureId = x.F
		default:
			// Ignoring unknown property
		}
	}
	return nil
}

/*****************************

amount struct

******************************/
func NewAmount() *amount {
	return &amount{
		private: private_p6lzdx9ziub7{},
	}
}

// Amount Object
// A`Transaction` object also may have an `ItemList`, which has dollar amounts.
// These amounts are used to calculate the `Total` field of the `Amount` object
//
// All other uses of `Amount` do have `shipping`, `shipping_discount` and
// `subtotal` to calculate the `Total`.
type amount struct {
	private private_p6lzdx9ziub7
	Details *details `json:"details,omitempty"`
}

type private_p6lzdx9ziub7 struct {
	Currency CurrencyTypeEnum `json:"currency"`
	Total    float64          `json:"total,string"`
}

type json_p6lzdx9ziub7 struct {
	*private_p6lzdx9ziub7
	Details *details `json:"details,omitempty"`
}

func (self *amount) Currency() CurrencyTypeEnum {
	return self.private.Currency
}

func (self *amount) Total() float64 {
	return self.private.Total
}

func (self *amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(json_p6lzdx9ziub7{
		&self.private,
		self.Details,
	})
}

func (self *amount) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	m := make(map[string]json.RawMessage)

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// For every property found, perform a separate UnmarshalJSON operation. This
	// prevents overwrite of values in 'self' where properties are absent.
	for key, rawMsg := range m {
		// The anon structs in each case are needed for field tags

		switch key {
		case "currency":

			var x struct {
				F CurrencyTypeEnum `json:"currency"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Currency = x.F
		case "total":

			var x struct {
				F float64 `json:"total,string"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Total = x.F
		case "details":

			var x struct {
				F *details `json:"details,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.Details = x.F

		default:
			// Ignoring unknown property
		}
	}
	return nil
}

/*****************************

details struct

******************************/
func NewDetails() *details {
	return &details{
		private: private_1m1bkd6tss1pk{},
	}
}

type details struct {
	private private_1m1bkd6tss1pk
	// Amount charged for shipping. 10 chars max, with support for 2 decimal places
	Shipping float64 `json:"shipping,string,omitempty"`
	// Amount being charged for handling fee. When `payment_method` is `paypal`
	HandlingFee float64 `json:"handling_fee,string,omitempty"`
	// Amount being charged for insurance fee. When `payment_method` is `paypal`
	Insurance float64 `json:"insurance,string,omitempty"`
	// Amount being discounted for shipping fee. When `payment_method` is `paypal`
	ShippingDiscount float64 `json:"shipping_discount,string,omitempty"`
}

type private_1m1bkd6tss1pk struct {
	// Amount of the subtotal of the items. REQUIRED if line items are specified.
	// 10 chars max, with support for 2 decimal places
	Subtotal float64 `json:"subtotal,string,omitempty"`
	// Amount charged for tax. 10 chars max, with support for 2 decimal places
	Tax float64 `json:"tax,string,omitempty"`
}

type json_1m1bkd6tss1pk struct {
	*private_1m1bkd6tss1pk
	Shipping         float64 `json:"shipping,string,omitempty"`
	HandlingFee      float64 `json:"handling_fee,string,omitempty"`
	Insurance        float64 `json:"insurance,string,omitempty"`
	ShippingDiscount float64 `json:"shipping_discount,string,omitempty"`
}

func (self *details) Subtotal() float64 {
	return self.private.Subtotal
}

func (self *details) Tax() float64 {
	return self.private.Tax
}

func (self *details) MarshalJSON() ([]byte, error) {
	return json.Marshal(json_1m1bkd6tss1pk{
		&self.private,
		self.Shipping,
		self.HandlingFee,
		self.Insurance,
		self.ShippingDiscount,
	})
}

func (self *details) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	m := make(map[string]json.RawMessage)

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// For every property found, perform a separate UnmarshalJSON operation. This
	// prevents overwrite of values in 'self' where properties are absent.
	for key, rawMsg := range m {
		// The anon structs in each case are needed for field tags

		switch key {
		case "subtotal":

			var x struct {
				F float64 `json:"subtotal,string,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Subtotal = x.F
		case "tax":

			var x struct {
				F float64 `json:"tax,string,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Tax = x.F
		case "shipping":

			var x struct {
				F float64 `json:"shipping,string,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.Shipping = x.F

		case "handling_fee":

			var x struct {
				F float64 `json:"handling_fee,string,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.HandlingFee = x.F

		case "insurance":

			var x struct {
				F float64 `json:"insurance,string,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.Insurance = x.F

		case "shipping_discount":

			var x struct {
				F float64 `json:"shipping_discount,string,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.ShippingDiscount = x.F

		default:
			// Ignoring unknown property
		}
	}
	return nil
}

/*****************************

link struct

******************************/
func NewLink() *link {
	return &link{
		private: private_hu9r4k8h8ws4{},
	}
}

type link struct {
	private private_hu9r4k8h8ws4
}

type private_hu9r4k8h8ws4 struct {
	Href   string      `json:"href,omitempty"`
	Rel    relTypeEnum `json:"rel,omitempty"`
	Method string      `json:"method,omitempty"`
}

type json_hu9r4k8h8ws4 struct {
	*private_hu9r4k8h8ws4
}

func (self *link) MarshalJSON() ([]byte, error) {
	return json.Marshal(json_hu9r4k8h8ws4{
		&self.private,
	})
}

func (self *link) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	m := make(map[string]json.RawMessage)

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// For every property found, perform a separate UnmarshalJSON operation. This
	// prevents overwrite of values in 'self' where properties are absent.
	for key, rawMsg := range m {
		// The anon structs in each case are needed for field tags

		switch key {
		case "href":

			var x struct {
				F string `json:"href,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Href = x.F
		case "rel":

			var x struct {
				F relTypeEnum `json:"rel,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Rel = x.F
		case "method":

			var x struct {
				F string `json:"method,omitempty"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Method = x.F
		default:
			// Ignoring unknown property
		}
	}
	return nil
}

/*****************************

currency struct

******************************/
func NewCurrency() *currency {
	return &currency{
		private: private_d1qbnu0nzezr{},
	}
}

// Base object for all financial value related fields (balance, payment due, etc.)
type currency struct {
	private  private_d1qbnu0nzezr
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type private_d1qbnu0nzezr struct{}

type json_d1qbnu0nzezr struct {
	*private_d1qbnu0nzezr
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

func (self *currency) MarshalJSON() ([]byte, error) {
	return json.Marshal(json_d1qbnu0nzezr{
		&self.private,
		self.Currency,
		self.Value,
	})
}

func (self *currency) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	m := make(map[string]json.RawMessage)

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// For every property found, perform a separate UnmarshalJSON operation. This
	// prevents overwrite of values in 'self' where properties are absent.
	for key, rawMsg := range m {
		// The anon structs in each case are needed for field tags

		switch key {
		case "currency":

			var x struct {
				F string `json:"currency"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.Currency = x.F

		case "value":

			var x struct {
				F string `json:"value"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.Value = x.F

		default:
			// Ignoring unknown property
		}
	}
	return nil
}

/*****************************

fmfDetails struct

******************************/
func NewFmfDetails() *fmfDetails {
	return &fmfDetails{
		private: private_1pbe00qed41oy{},
	}
}

// This object represents Fraud Management Filter (FMF) details for a payment.
type fmfDetails struct {
	private private_1pbe00qed41oy
}

type private_1pbe00qed41oy struct {
	FilterType  fmfFilterEnum `json:"filter_type"`
	FilterID    filterIdEnum  `json:"filter_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
}

type json_1pbe00qed41oy struct {
	*private_1pbe00qed41oy
}

func (self *fmfDetails) FilterType() fmfFilterEnum {
	return self.private.FilterType
}

func (self *fmfDetails) FilterID() filterIdEnum {
	return self.private.FilterID
}

func (self *fmfDetails) Name() string {
	return self.private.Name
}

func (self *fmfDetails) Description() string {
	return self.private.Description
}

func (self *fmfDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(json_1pbe00qed41oy{
		&self.private,
	})
}

func (self *fmfDetails) UnmarshalJSON(j []byte) error {
	if len(j) == 4 && string(j) == "null" {
		return nil
	}

	m := make(map[string]json.RawMessage)

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// For every property found, perform a separate UnmarshalJSON operation. This
	// prevents overwrite of values in 'self' where properties are absent.
	for key, rawMsg := range m {
		// The anon structs in each case are needed for field tags

		switch key {
		case "filter_type":

			var x struct {
				F fmfFilterEnum `json:"filter_type"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.FilterType = x.F
		case "filter_id":

			var x struct {
				F filterIdEnum `json:"filter_id"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.FilterID = x.F
		case "name":

			var x struct {
				F string `json:"name"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Name = x.F
		case "description":

			var x struct {
				F string `json:"description"`
			}

			var msgForStruct = append(append(append(append(
				[]byte("{\""), key...), "\":"...), rawMsg...), '}')

			if err = json.Unmarshal(msgForStruct, &x); err != nil {
				return err
			}

			self.private.Description = x.F
		default:
			// Ignoring unknown property
		}
	}
	return nil
}
