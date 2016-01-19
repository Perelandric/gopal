package gopal

import (
	"log"
	"strconv"
)

/*****************************

sortByEnum

******************************/

type sortByEnum struct{ value_1bdb4312qifp9 uint8 }

var sortBy = struct {
	CreateTime sortByEnum
	UpdateTime sortByEnum
}{
	CreateTime: sortByEnum{value_1bdb4312qifp9: 1},
	UpdateTime: sortByEnum{value_1bdb4312qifp9: 2},
}

// Used to iterate in range loops
var sortByValues = [...]sortByEnum{
	sortBy.CreateTime, sortBy.UpdateTime,
}

// Get the integer value of the enum variant
func (self sortByEnum) Value() uint8 {
	return self.value_1bdb4312qifp9
}

func (self sortByEnum) IntValue() int {
	return int(self.value_1bdb4312qifp9)
}

// Get the string representation of the enum variant
func (self sortByEnum) String() string {
	switch self.value_1bdb4312qifp9 {
	case 1:
		return "create_time"
	case 2:
		return "update_time"
	}

	return ""
}

// Get the string description of the enum variant
func (self sortByEnum) Description() string {
	switch self.value_1bdb4312qifp9 {
	case 1:
		return "create_time"
	case 2:
		return "update_time"
	}
	return ""
}

func (self sortByEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(self.String())), nil
}

func (self *sortByEnum) UnmarshalJSON(b []byte) error {
	var s, err = strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	if len(s) == 0 {
		return nil
	}

	switch s {
	case "create_time":
		self.value_1bdb4312qifp9 = 1
		return nil
	case "update_time":
		self.value_1bdb4312qifp9 = 2
		return nil
	default:
		log.Printf("Unexpected value: %q while unmarshaling sortByEnum\n", s)
	}

	return nil
}

/*****************************

sortOrderEnum

******************************/

type sortOrderEnum struct{ value_1bsicaugk4iip uint8 }

var sortOrder = struct {
	DESC sortOrderEnum
	ASC  sortOrderEnum
}{
	DESC: sortOrderEnum{value_1bsicaugk4iip: 1},
	ASC:  sortOrderEnum{value_1bsicaugk4iip: 2},
}

// Used to iterate in range loops
var sortOrderValues = [...]sortOrderEnum{
	sortOrder.DESC, sortOrder.ASC,
}

// Get the integer value of the enum variant
func (self sortOrderEnum) Value() uint8 {
	return self.value_1bsicaugk4iip
}

func (self sortOrderEnum) IntValue() int {
	return int(self.value_1bsicaugk4iip)
}

// Get the string representation of the enum variant
func (self sortOrderEnum) String() string {
	switch self.value_1bsicaugk4iip {
	case 1:
		return "DESC"
	case 2:
		return "ASC"
	}

	return ""
}

// Get the string description of the enum variant
func (self sortOrderEnum) Description() string {
	switch self.value_1bsicaugk4iip {
	case 1:
		return "DESC"
	case 2:
		return "ASC"
	}
	return ""
}

func (self sortOrderEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *sortOrderEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_1bsicaugk4iip = uint8(n)
	return nil
}

/*****************************

relTypeEnum

******************************/

type relTypeEnum struct{ value_1alutenzzdh1v uint8 }

var relType = struct {
	self          relTypeEnum
	parentPayment relTypeEnum
	execute       relTypeEnum
	refund        relTypeEnum
	approvalUrl   relTypeEnum
	suspend       relTypeEnum
	reActivate    relTypeEnum
	cancel        relTypeEnum
	void          relTypeEnum
	authorization relTypeEnum
	capture       relTypeEnum
	reauthorize   relTypeEnum
	order         relTypeEnum
	item          relTypeEnum
	batch         relTypeEnum
	delete        relTypeEnum
	patch         relTypeEnum
	first         relTypeEnum
	last          relTypeEnum
	update        relTypeEnum
	resend        relTypeEnum
	next          relTypeEnum
	previous      relTypeEnum
	start         relTypeEnum
	nextPage      relTypeEnum
	previousPage  relTypeEnum
}{
	self:          relTypeEnum{value_1alutenzzdh1v: 1},
	parentPayment: relTypeEnum{value_1alutenzzdh1v: 2},
	execute:       relTypeEnum{value_1alutenzzdh1v: 3},
	refund:        relTypeEnum{value_1alutenzzdh1v: 4},
	approvalUrl:   relTypeEnum{value_1alutenzzdh1v: 5},
	suspend:       relTypeEnum{value_1alutenzzdh1v: 6},
	reActivate:    relTypeEnum{value_1alutenzzdh1v: 7},
	cancel:        relTypeEnum{value_1alutenzzdh1v: 8},
	void:          relTypeEnum{value_1alutenzzdh1v: 9},
	authorization: relTypeEnum{value_1alutenzzdh1v: 10},
	capture:       relTypeEnum{value_1alutenzzdh1v: 11},
	reauthorize:   relTypeEnum{value_1alutenzzdh1v: 12},
	order:         relTypeEnum{value_1alutenzzdh1v: 13},
	item:          relTypeEnum{value_1alutenzzdh1v: 14},
	batch:         relTypeEnum{value_1alutenzzdh1v: 15},
	delete:        relTypeEnum{value_1alutenzzdh1v: 16},
	patch:         relTypeEnum{value_1alutenzzdh1v: 17},
	first:         relTypeEnum{value_1alutenzzdh1v: 18},
	last:          relTypeEnum{value_1alutenzzdh1v: 19},
	update:        relTypeEnum{value_1alutenzzdh1v: 20},
	resend:        relTypeEnum{value_1alutenzzdh1v: 21},
	next:          relTypeEnum{value_1alutenzzdh1v: 22},
	previous:      relTypeEnum{value_1alutenzzdh1v: 23},
	start:         relTypeEnum{value_1alutenzzdh1v: 24},
	nextPage:      relTypeEnum{value_1alutenzzdh1v: 25},
	previousPage:  relTypeEnum{value_1alutenzzdh1v: 26},
}

// Used to iterate in range loops
var relTypeValues = [...]relTypeEnum{
	relType.self, relType.parentPayment, relType.execute, relType.refund, relType.approvalUrl, relType.suspend, relType.reActivate, relType.cancel, relType.void, relType.authorization, relType.capture, relType.reauthorize, relType.order, relType.item, relType.batch, relType.delete, relType.patch, relType.first, relType.last, relType.update, relType.resend, relType.next, relType.previous, relType.start, relType.nextPage, relType.previousPage,
}

// Get the integer value of the enum variant
func (self relTypeEnum) Value() uint8 {
	return self.value_1alutenzzdh1v
}

func (self relTypeEnum) IntValue() int {
	return int(self.value_1alutenzzdh1v)
}

// Get the string representation of the enum variant
func (self relTypeEnum) String() string {
	switch self.value_1alutenzzdh1v {
	case 1:
		return "self"
	case 2:
		return "parent_payment"
	case 3:
		return "execute"
	case 4:
		return "refund"
	case 5:
		return "approval_url"
	case 6:
		return "suspend"
	case 7:
		return "re_activate"
	case 8:
		return "cancel"
	case 9:
		return "void"
	case 10:
		return "authorization"
	case 11:
		return "capture"
	case 12:
		return "reauthorize"
	case 13:
		return "order"
	case 14:
		return "item"
	case 15:
		return "batch"
	case 16:
		return "delete"
	case 17:
		return "patch"
	case 18:
		return "first"
	case 19:
		return "last"
	case 20:
		return "update"
	case 21:
		return "resend"
	case 22:
		return "next"
	case 23:
		return "previous"
	case 24:
		return "start"
	case 25:
		return "next_page"
	case 26:
		return "previous_page"
	}

	return ""
}

// Get the string description of the enum variant
func (self relTypeEnum) Description() string {
	switch self.value_1alutenzzdh1v {
	case 1:
		return "self"
	case 2:
		return "parent_payment"
	case 3:
		return "execute"
	case 4:
		return "refund"
	case 5:
		return "approval_url"
	case 6:
		return "suspend"
	case 7:
		return "re_activate"
	case 8:
		return "cancel"
	case 9:
		return "void"
	case 10:
		return "authorization"
	case 11:
		return "capture"
	case 12:
		return "reauthorize"
	case 13:
		return "order"
	case 14:
		return "item"
	case 15:
		return "batch"
	case 16:
		return "delete"
	case 17:
		return "patch"
	case 18:
		return "first"
	case 19:
		return "last"
	case 20:
		return "update"
	case 21:
		return "resend"
	case 22:
		return "next"
	case 23:
		return "previous"
	case 24:
		return "start"
	case 25:
		return "next_page"
	case 26:
		return "previous_page"
	}
	return ""
}

func (self relTypeEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *relTypeEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_1alutenzzdh1v = uint8(n)
	return nil
}

/*****************************

methodEnum

******************************/

type methodEnum struct{ value_omo42qsnx47d uint8 }

var method = struct {
	get      methodEnum
	post     methodEnum
	redirect methodEnum
	delete   methodEnum
	patch    methodEnum
}{
	get:      methodEnum{value_omo42qsnx47d: 1},
	post:     methodEnum{value_omo42qsnx47d: 2},
	redirect: methodEnum{value_omo42qsnx47d: 3},
	delete:   methodEnum{value_omo42qsnx47d: 4},
	patch:    methodEnum{value_omo42qsnx47d: 5},
}

// Used to iterate in range loops
var methodValues = [...]methodEnum{
	method.get, method.post, method.redirect, method.delete, method.patch,
}

// Get the integer value of the enum variant
func (self methodEnum) Value() uint8 {
	return self.value_omo42qsnx47d
}

func (self methodEnum) IntValue() int {
	return int(self.value_omo42qsnx47d)
}

// Get the string representation of the enum variant
func (self methodEnum) String() string {
	switch self.value_omo42qsnx47d {
	case 1:
		return "GET"
	case 2:
		return "POST"
	case 3:
		return "REDIRECT"
	case 4:
		return "DELETE"
	case 5:
		return "PATCH"
	}

	return ""
}

// Get the string description of the enum variant
func (self methodEnum) Description() string {
	switch self.value_omo42qsnx47d {
	case 1:
		return "GET"
	case 2:
		return "POST"
	case 3:
		return "REDIRECT"
	case 4:
		return "DELETE"
	case 5:
		return "PATCH"
	}
	return ""
}

func (self methodEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *methodEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_omo42qsnx47d = uint8(n)
	return nil
}

/*****************************

payerStatusEnum

******************************/

type payerStatusEnum struct{ value_1gh7wwhh83zx2 uint8 }

var payerStatus = struct {
	verified   payerStatusEnum
	unverified payerStatusEnum
}{
	verified:   payerStatusEnum{value_1gh7wwhh83zx2: 1},
	unverified: payerStatusEnum{value_1gh7wwhh83zx2: 2},
}

// Used to iterate in range loops
var payerStatusValues = [...]payerStatusEnum{
	payerStatus.verified, payerStatus.unverified,
}

// Get the integer value of the enum variant
func (self payerStatusEnum) Value() uint8 {
	return self.value_1gh7wwhh83zx2
}

func (self payerStatusEnum) IntValue() int {
	return int(self.value_1gh7wwhh83zx2)
}

// Get the string representation of the enum variant
func (self payerStatusEnum) String() string {
	switch self.value_1gh7wwhh83zx2 {
	case 1:
		return "VERIFIED"
	case 2:
		return "UNVERIFIED"
	}

	return ""
}

// Get the string description of the enum variant
func (self payerStatusEnum) Description() string {
	switch self.value_1gh7wwhh83zx2 {
	case 1:
		return "VERIFIED"
	case 2:
		return "UNVERIFIED"
	}
	return ""
}

func (self payerStatusEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *payerStatusEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_1gh7wwhh83zx2 = uint8(n)
	return nil
}

/*****************************

intentEnum

******************************/

type intentEnum struct{ value_1bdpyb5hld7lv uint8 }

var intent = struct {
	sale      intentEnum
	authorize intentEnum
	order     intentEnum
}{
	sale:      intentEnum{value_1bdpyb5hld7lv: 1},
	authorize: intentEnum{value_1bdpyb5hld7lv: 2},
	order:     intentEnum{value_1bdpyb5hld7lv: 3},
}

// Used to iterate in range loops
var intentValues = [...]intentEnum{
	intent.sale, intent.authorize, intent.order,
}

// Get the integer value of the enum variant
func (self intentEnum) Value() uint8 {
	return self.value_1bdpyb5hld7lv
}

func (self intentEnum) IntValue() int {
	return int(self.value_1bdpyb5hld7lv)
}

// Get the string representation of the enum variant
func (self intentEnum) String() string {
	switch self.value_1bdpyb5hld7lv {
	case 1:
		return "sale"
	case 2:
		return "authorize"
	case 3:
		return "order"
	}

	return ""
}

// Get the string description of the enum variant
func (self intentEnum) Description() string {
	switch self.value_1bdpyb5hld7lv {
	case 1:
		return "sale"
	case 2:
		return "authorize"
	case 3:
		return "order"
	}
	return ""
}

func (self intentEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *intentEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_1bdpyb5hld7lv = uint8(n)
	return nil
}

/*****************************

normStatusEnum

******************************/

type normStatusEnum struct{ value_z3e0jx0heqx1 uint8 }

var normStatus = struct {
	unknown                   normStatusEnum
	unnormalizedUserPreferred normStatusEnum
	normalized                normStatusEnum
	unnormalized              normStatusEnum
}{
	unknown:                   normStatusEnum{value_z3e0jx0heqx1: 1},
	unnormalizedUserPreferred: normStatusEnum{value_z3e0jx0heqx1: 2},
	normalized:                normStatusEnum{value_z3e0jx0heqx1: 3},
	unnormalized:              normStatusEnum{value_z3e0jx0heqx1: 4},
}

// Used to iterate in range loops
var normStatusValues = [...]normStatusEnum{
	normStatus.unknown, normStatus.unnormalizedUserPreferred, normStatus.normalized, normStatus.unnormalized,
}

// Get the integer value of the enum variant
func (self normStatusEnum) Value() uint8 {
	return self.value_z3e0jx0heqx1
}

func (self normStatusEnum) IntValue() int {
	return int(self.value_z3e0jx0heqx1)
}

// Get the string representation of the enum variant
func (self normStatusEnum) String() string {
	switch self.value_z3e0jx0heqx1 {
	case 1:
		return "UNKNOWN"
	case 2:
		return "UNNORMALIZED_USER_PREFERRED"
	case 3:
		return "NORMALIZED"
	case 4:
		return "UNNORMALIZED"
	}

	return ""
}

// Get the string description of the enum variant
func (self normStatusEnum) Description() string {
	switch self.value_z3e0jx0heqx1 {
	case 1:
		return "UNKNOWN"
	case 2:
		return "UNNORMALIZED_USER_PREFERRED"
	case 3:
		return "NORMALIZED"
	case 4:
		return "UNNORMALIZED"
	}
	return ""
}

func (self normStatusEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *normStatusEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_z3e0jx0heqx1 = uint8(n)
	return nil
}

/*****************************

addressStatusEnum

******************************/

type addressStatusEnum struct{ value_1b5598h1o0rwp uint8 }

var addressStatus = struct {
	confirmed   addressStatusEnum
	unconfirmed addressStatusEnum
}{
	confirmed:   addressStatusEnum{value_1b5598h1o0rwp: 1},
	unconfirmed: addressStatusEnum{value_1b5598h1o0rwp: 2},
}

// Used to iterate in range loops
var addressStatusValues = [...]addressStatusEnum{
	addressStatus.confirmed, addressStatus.unconfirmed,
}

// Get the integer value of the enum variant
func (self addressStatusEnum) Value() uint8 {
	return self.value_1b5598h1o0rwp
}

func (self addressStatusEnum) IntValue() int {
	return int(self.value_1b5598h1o0rwp)
}

// Get the string representation of the enum variant
func (self addressStatusEnum) String() string {
	switch self.value_1b5598h1o0rwp {
	case 1:
		return "CONFIRMED"
	case 2:
		return "UNCONFIRMED"
	}

	return ""
}

// Get the string description of the enum variant
func (self addressStatusEnum) Description() string {
	switch self.value_1b5598h1o0rwp {
	case 1:
		return "CONFIRMED"
	case 2:
		return "UNCONFIRMED"
	}
	return ""
}

func (self addressStatusEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *addressStatusEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_1b5598h1o0rwp = uint8(n)
	return nil
}

/*****************************

AddressTypeEnum

******************************/

type AddressTypeEnum struct{ value_1hikwwg1aao4l uint8 }

var AddressType = struct {
	Residential AddressTypeEnum
	Business    AddressTypeEnum
	Mailbox     AddressTypeEnum
}{
	Residential: AddressTypeEnum{value_1hikwwg1aao4l: 1},
	Business:    AddressTypeEnum{value_1hikwwg1aao4l: 2},
	Mailbox:     AddressTypeEnum{value_1hikwwg1aao4l: 3},
}

// Used to iterate in range loops
var AddressTypeValues = [...]AddressTypeEnum{
	AddressType.Residential, AddressType.Business, AddressType.Mailbox,
}

// Get the integer value of the enum variant
func (self AddressTypeEnum) Value() uint8 {
	return self.value_1hikwwg1aao4l
}

func (self AddressTypeEnum) IntValue() int {
	return int(self.value_1hikwwg1aao4l)
}

// Get the string representation of the enum variant
func (self AddressTypeEnum) String() string {
	switch self.value_1hikwwg1aao4l {
	case 1:
		return "residential"
	case 2:
		return "business"
	case 3:
		return "mailbox"
	}

	return ""
}

// Get the string description of the enum variant
func (self AddressTypeEnum) Description() string {
	switch self.value_1hikwwg1aao4l {
	case 1:
		return "residential"
	case 2:
		return "business"
	case 3:
		return "mailbox"
	}
	return ""
}

func (self AddressTypeEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *AddressTypeEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_1hikwwg1aao4l = uint8(n)
	return nil
}

/*****************************

PaymentMethodEnum

******************************/

type PaymentMethodEnum struct{ value_12tvrulfmfchw uint8 }

var PaymentMethod = struct {
	CreditCard PaymentMethodEnum
	PayPal     PaymentMethodEnum
}{
	CreditCard: PaymentMethodEnum{value_12tvrulfmfchw: 1},
	PayPal:     PaymentMethodEnum{value_12tvrulfmfchw: 2},
}

// Used to iterate in range loops
var PaymentMethodValues = [...]PaymentMethodEnum{
	PaymentMethod.CreditCard, PaymentMethod.PayPal,
}

// Get the integer value of the enum variant
func (self PaymentMethodEnum) Value() uint8 {
	return self.value_12tvrulfmfchw
}

func (self PaymentMethodEnum) IntValue() int {
	return int(self.value_12tvrulfmfchw)
}

// Get the string representation of the enum variant
func (self PaymentMethodEnum) String() string {
	switch self.value_12tvrulfmfchw {
	case 1:
		return "credit_card"
	case 2:
		return "paypal"
	}

	return ""
}

// Get the string description of the enum variant
func (self PaymentMethodEnum) Description() string {
	switch self.value_12tvrulfmfchw {
	case 1:
		return "credit_card"
	case 2:
		return "paypal"
	}
	return ""
}

func (self PaymentMethodEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *PaymentMethodEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_12tvrulfmfchw = uint8(n)
	return nil
}

/*****************************

TaxIdTypeEnum

******************************/

type TaxIdTypeEnum struct{ value_hcc5w6wugby9 uint8 }

var TaxIdType = struct {
	BrCpf  TaxIdTypeEnum
	BrCnpj TaxIdTypeEnum
}{
	BrCpf:  TaxIdTypeEnum{value_hcc5w6wugby9: 1},
	BrCnpj: TaxIdTypeEnum{value_hcc5w6wugby9: 2},
}

// Used to iterate in range loops
var TaxIdTypeValues = [...]TaxIdTypeEnum{
	TaxIdType.BrCpf, TaxIdType.BrCnpj,
}

// Get the integer value of the enum variant
func (self TaxIdTypeEnum) Value() uint8 {
	return self.value_hcc5w6wugby9
}

func (self TaxIdTypeEnum) IntValue() int {
	return int(self.value_hcc5w6wugby9)
}

// Get the string representation of the enum variant
func (self TaxIdTypeEnum) String() string {
	switch self.value_hcc5w6wugby9 {
	case 1:
		return "BR_CPF"
	case 2:
		return "BR_CNPJ"
	}

	return ""
}

// Get the string description of the enum variant
func (self TaxIdTypeEnum) Description() string {
	switch self.value_hcc5w6wugby9 {
	case 1:
		return "BR_CPF"
	case 2:
		return "BR_CNPJ"
	}
	return ""
}

func (self TaxIdTypeEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *TaxIdTypeEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_hcc5w6wugby9 = uint8(n)
	return nil
}

/*****************************

CreditCardTypeEnum

******************************/

type CreditCardTypeEnum struct{ value_apa8ddtp76st uint8 }

var CreditCardType = struct {
	Visa       CreditCardTypeEnum
	MasterCard CreditCardTypeEnum
	Discover   CreditCardTypeEnum
	Amex       CreditCardTypeEnum
}{
	Visa:       CreditCardTypeEnum{value_apa8ddtp76st: 1},
	MasterCard: CreditCardTypeEnum{value_apa8ddtp76st: 2},
	Discover:   CreditCardTypeEnum{value_apa8ddtp76st: 3},
	Amex:       CreditCardTypeEnum{value_apa8ddtp76st: 4},
}

// Used to iterate in range loops
var CreditCardTypeValues = [...]CreditCardTypeEnum{
	CreditCardType.Visa, CreditCardType.MasterCard, CreditCardType.Discover, CreditCardType.Amex,
}

// Get the integer value of the enum variant
func (self CreditCardTypeEnum) Value() uint8 {
	return self.value_apa8ddtp76st
}

func (self CreditCardTypeEnum) IntValue() int {
	return int(self.value_apa8ddtp76st)
}

// Get the string representation of the enum variant
func (self CreditCardTypeEnum) String() string {
	switch self.value_apa8ddtp76st {
	case 1:
		return "visa"
	case 2:
		return "mastercard"
	case 3:
		return "discover"
	case 4:
		return "amex"
	}

	return ""
}

// Get the string description of the enum variant
func (self CreditCardTypeEnum) Description() string {
	switch self.value_apa8ddtp76st {
	case 1:
		return "visa"
	case 2:
		return "mastercard"
	case 3:
		return "discover"
	case 4:
		return "amex"
	}
	return ""
}

func (self CreditCardTypeEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *CreditCardTypeEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_apa8ddtp76st = uint8(n)
	return nil
}

/*****************************

stateEnum

******************************/

type stateEnum struct{ value_xmb4rx7sals7 uint8 }

var state = struct {
	Created           stateEnum
	Approved          stateEnum
	Canceled          stateEnum
	InProgress        stateEnum
	Failed            stateEnum
	Pending           stateEnum
	Completed         stateEnum
	Refunded          stateEnum
	PartiallyRefunded stateEnum
	Expired           stateEnum
	Ok                stateEnum
	Authorized        stateEnum
	Captured          stateEnum
	PartiallyCaptured stateEnum
	Voided            stateEnum
}{
	Created:           stateEnum{value_xmb4rx7sals7: 1},
	Approved:          stateEnum{value_xmb4rx7sals7: 2},
	Canceled:          stateEnum{value_xmb4rx7sals7: 3},
	InProgress:        stateEnum{value_xmb4rx7sals7: 4},
	Failed:            stateEnum{value_xmb4rx7sals7: 5},
	Pending:           stateEnum{value_xmb4rx7sals7: 6},
	Completed:         stateEnum{value_xmb4rx7sals7: 7},
	Refunded:          stateEnum{value_xmb4rx7sals7: 8},
	PartiallyRefunded: stateEnum{value_xmb4rx7sals7: 9},
	Expired:           stateEnum{value_xmb4rx7sals7: 10},
	Ok:                stateEnum{value_xmb4rx7sals7: 11},
	Authorized:        stateEnum{value_xmb4rx7sals7: 12},
	Captured:          stateEnum{value_xmb4rx7sals7: 13},
	PartiallyCaptured: stateEnum{value_xmb4rx7sals7: 14},
	Voided:            stateEnum{value_xmb4rx7sals7: 15},
}

// Used to iterate in range loops
var stateValues = [...]stateEnum{
	state.Created, state.Approved, state.Canceled, state.InProgress, state.Failed, state.Pending, state.Completed, state.Refunded, state.PartiallyRefunded, state.Expired, state.Ok, state.Authorized, state.Captured, state.PartiallyCaptured, state.Voided,
}

// Get the integer value of the enum variant
func (self stateEnum) Value() uint8 {
	return self.value_xmb4rx7sals7
}

func (self stateEnum) IntValue() int {
	return int(self.value_xmb4rx7sals7)
}

// Get the string representation of the enum variant
func (self stateEnum) String() string {
	switch self.value_xmb4rx7sals7 {
	case 1:
		return "created"
	case 2:
		return "approved"
	case 3:
		return "canceled"
	case 4:
		return "in_progress"
	case 5:
		return "failed"
	case 6:
		return "pending"
	case 7:
		return "completed"
	case 8:
		return "refunded"
	case 9:
		return "partially_refunded"
	case 10:
		return "expired"
	case 11:
		return "ok"
	case 12:
		return "authorized"
	case 13:
		return "captured"
	case 14:
		return "partially_captured"
	case 15:
		return "voided"
	}

	return ""
}

// Get the string description of the enum variant
func (self stateEnum) Description() string {
	switch self.value_xmb4rx7sals7 {
	case 1:
		return "created"
	case 2:
		return "approved"
	case 3:
		return "canceled"
	case 4:
		return "in_progress"
	case 5:
		return "failed"
	case 6:
		return "pending"
	case 7:
		return "completed"
	case 8:
		return "refunded"
	case 9:
		return "partially_refunded"
	case 10:
		return "expired"
	case 11:
		return "ok"
	case 12:
		return "authorized"
	case 13:
		return "captured"
	case 14:
		return "partially_captured"
	case 15:
		return "voided"
	}
	return ""
}

func (self stateEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *stateEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_xmb4rx7sals7 = uint8(n)
	return nil
}

/*****************************

reasonCodeEnum

******************************/

type reasonCodeEnum struct{ value_1idhpk8vx7oji uint8 }

var reasonCode = struct {
	chargeback                              reasonCodeEnum
	guarantee                               reasonCodeEnum
	buyerComplaint                          reasonCodeEnum
	refundCode                              reasonCodeEnum
	unconfirmedShippingAddress              reasonCodeEnum
	echeckCode                              reasonCodeEnum
	internationalWithdrawal                 reasonCodeEnum
	receivingPreferenceMandatesManualAction reasonCodeEnum
	paymentReview                           reasonCodeEnum
	regulatoryReview                        reasonCodeEnum
	unilateral                              reasonCodeEnum
	verificationRequired                    reasonCodeEnum
}{
	chargeback:                              reasonCodeEnum{value_1idhpk8vx7oji: 1},
	guarantee:                               reasonCodeEnum{value_1idhpk8vx7oji: 2},
	buyerComplaint:                          reasonCodeEnum{value_1idhpk8vx7oji: 3},
	refundCode:                              reasonCodeEnum{value_1idhpk8vx7oji: 4},
	unconfirmedShippingAddress:              reasonCodeEnum{value_1idhpk8vx7oji: 5},
	echeckCode:                              reasonCodeEnum{value_1idhpk8vx7oji: 6},
	internationalWithdrawal:                 reasonCodeEnum{value_1idhpk8vx7oji: 7},
	receivingPreferenceMandatesManualAction: reasonCodeEnum{value_1idhpk8vx7oji: 8},
	paymentReview:                           reasonCodeEnum{value_1idhpk8vx7oji: 9},
	regulatoryReview:                        reasonCodeEnum{value_1idhpk8vx7oji: 10},
	unilateral:                              reasonCodeEnum{value_1idhpk8vx7oji: 11},
	verificationRequired:                    reasonCodeEnum{value_1idhpk8vx7oji: 12},
}

// Used to iterate in range loops
var reasonCodeValues = [...]reasonCodeEnum{
	reasonCode.chargeback, reasonCode.guarantee, reasonCode.buyerComplaint, reasonCode.refundCode, reasonCode.unconfirmedShippingAddress, reasonCode.echeckCode, reasonCode.internationalWithdrawal, reasonCode.receivingPreferenceMandatesManualAction, reasonCode.paymentReview, reasonCode.regulatoryReview, reasonCode.unilateral, reasonCode.verificationRequired,
}

// Get the integer value of the enum variant
func (self reasonCodeEnum) Value() uint8 {
	return self.value_1idhpk8vx7oji
}

func (self reasonCodeEnum) IntValue() int {
	return int(self.value_1idhpk8vx7oji)
}

// Get the string representation of the enum variant
func (self reasonCodeEnum) String() string {
	switch self.value_1idhpk8vx7oji {
	case 1:
		return "CHARGEBACK"
	case 2:
		return "GUARANTEE"
	case 3:
		return "BUYER_COMPLAINT"
	case 4:
		return "REFUND"
	case 5:
		return "UNCONFIRMED_SHIPPING_ADDRESS"
	case 6:
		return "ECHECK"
	case 7:
		return "INTERNATIONAL_WITHDRAWAL"
	case 8:
		return "RECEIVING_PREFERENCE_MANDATES_MANUAL_ACTION"
	case 9:
		return "PAYMENT_REVIEW"
	case 10:
		return "REGULATORY_REVIEW"
	case 11:
		return "UNILATERAL"
	case 12:
		return "VERIFICATION_REQUIRED"
	}

	return ""
}

// Get the string description of the enum variant
func (self reasonCodeEnum) Description() string {
	switch self.value_1idhpk8vx7oji {
	case 1:
		return "CHARGEBACK"
	case 2:
		return "GUARANTEE"
	case 3:
		return "BUYER_COMPLAINT"
	case 4:
		return "REFUND"
	case 5:
		return "UNCONFIRMED_SHIPPING_ADDRESS"
	case 6:
		return "ECHECK"
	case 7:
		return "INTERNATIONAL_WITHDRAWAL"
	case 8:
		return "RECEIVING_PREFERENCE_MANDATES_MANUAL_ACTION"
	case 9:
		return "PAYMENT_REVIEW"
	case 10:
		return "REGULATORY_REVIEW"
	case 11:
		return "UNILATERAL"
	case 12:
		return "VERIFICATION_REQUIRED"
	}
	return ""
}

func (self reasonCodeEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *reasonCodeEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_1idhpk8vx7oji = uint8(n)
	return nil
}

/*****************************

protectionEligEnum

******************************/

type protectionEligEnum struct{ value_1vqbkups5p19a uint8 }

var protectionElig = struct {
	eligible          protectionEligEnum
	partiallyEligible protectionEligEnum
	ineligibile       protectionEligEnum
}{
	eligible:          protectionEligEnum{value_1vqbkups5p19a: 1},
	partiallyEligible: protectionEligEnum{value_1vqbkups5p19a: 2},
	ineligibile:       protectionEligEnum{value_1vqbkups5p19a: 3},
}

// Used to iterate in range loops
var protectionEligValues = [...]protectionEligEnum{
	protectionElig.eligible, protectionElig.partiallyEligible, protectionElig.ineligibile,
}

// Get the integer value of the enum variant
func (self protectionEligEnum) Value() uint8 {
	return self.value_1vqbkups5p19a
}

func (self protectionEligEnum) IntValue() int {
	return int(self.value_1vqbkups5p19a)
}

// Get the string representation of the enum variant
func (self protectionEligEnum) String() string {
	switch self.value_1vqbkups5p19a {
	case 1:
		return "ELIGIBLE"
	case 2:
		return "PARTIALLY_ELIGIBLE"
	case 3:
		return "INELIGIBLE"
	}

	return ""
}

// Get the string description of the enum variant
func (self protectionEligEnum) Description() string {
	switch self.value_1vqbkups5p19a {
	case 1:
		return "Merchant is protected by PayPal's Seller Protection Policy for Unauthorized. Payments and Item Not Received."
	case 2:
		return "Merchant is protected by PayPal's Seller Protection Policy for Item Not Received or Unauthorized Payments. Refer to protection_eligibility_type for specifics."
	case 3:
		return "Merchant is not protected under the Seller Protection Policy."
	}
	return ""
}

func (self protectionEligEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *protectionEligEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_1vqbkups5p19a = uint8(n)
	return nil
}

/*****************************

protectionEligTypeEnum

******************************/

type protectionEligTypeEnum struct{ value_16l0qe767r26c uint8 }

var protectionEligType = struct {
	itemNotReceivedEligible     protectionEligTypeEnum
	unauthorizedPaymentEligible protectionEligTypeEnum
}{
	itemNotReceivedEligible:     protectionEligTypeEnum{value_16l0qe767r26c: 1},
	unauthorizedPaymentEligible: protectionEligTypeEnum{value_16l0qe767r26c: 2},
}

// Used to iterate in range loops
var protectionEligTypeValues = [...]protectionEligTypeEnum{
	protectionEligType.itemNotReceivedEligible, protectionEligType.unauthorizedPaymentEligible,
}

// Get the integer value of the enum variant
func (self protectionEligTypeEnum) Value() uint8 {
	return self.value_16l0qe767r26c
}

func (self protectionEligTypeEnum) IntValue() int {
	return int(self.value_16l0qe767r26c)
}

// Get the string representation of the enum variant
func (self protectionEligTypeEnum) String() string {
	switch self.value_16l0qe767r26c {
	case 1:
		return "ITEM_NOT_RECEIVED_ELIGIBLE"
	case 2:
		return "UNAUTHORIZED_PAYMENT_ELIGIBLE"
	}

	return ""
}

// Get the string description of the enum variant
func (self protectionEligTypeEnum) Description() string {
	switch self.value_16l0qe767r26c {
	case 1:
		return "Sellers are protected against claims for items not received."
	case 2:
		return "Sellers are protected against claims for unauthorized payments."
	}
	return ""
}

func (self protectionEligTypeEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *protectionEligTypeEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_16l0qe767r26c = uint8(n)
	return nil
}

/*****************************

paymentModeEnum

******************************/

type paymentModeEnum struct{ value_1ohaudjwuc0ex uint8 }

var paymentMode = struct {
	instantTransfer    paymentModeEnum
	manualBankTransfer paymentModeEnum
	delayedTransfer    paymentModeEnum
	echeck             paymentModeEnum
}{
	instantTransfer:    paymentModeEnum{value_1ohaudjwuc0ex: 1},
	manualBankTransfer: paymentModeEnum{value_1ohaudjwuc0ex: 2},
	delayedTransfer:    paymentModeEnum{value_1ohaudjwuc0ex: 3},
	echeck:             paymentModeEnum{value_1ohaudjwuc0ex: 4},
}

// Used to iterate in range loops
var paymentModeValues = [...]paymentModeEnum{
	paymentMode.instantTransfer, paymentMode.manualBankTransfer, paymentMode.delayedTransfer, paymentMode.echeck,
}

// Get the integer value of the enum variant
func (self paymentModeEnum) Value() uint8 {
	return self.value_1ohaudjwuc0ex
}

func (self paymentModeEnum) IntValue() int {
	return int(self.value_1ohaudjwuc0ex)
}

// Get the string representation of the enum variant
func (self paymentModeEnum) String() string {
	switch self.value_1ohaudjwuc0ex {
	case 1:
		return "INSTANT_TRANSFER"
	case 2:
		return "MANUAL_BANK_TRANSFER"
	case 3:
		return "DELAYED_TRANSFER"
	case 4:
		return "ECHECK"
	}

	return ""
}

// Get the string description of the enum variant
func (self paymentModeEnum) Description() string {
	switch self.value_1ohaudjwuc0ex {
	case 1:
		return "INSTANT_TRANSFER"
	case 2:
		return "MANUAL_BANK_TRANSFER"
	case 3:
		return "DELAYED_TRANSFER"
	case 4:
		return "ECHECK"
	}
	return ""
}

func (self paymentModeEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *paymentModeEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_1ohaudjwuc0ex = uint8(n)
	return nil
}

/*****************************

pendingReasonEnum

******************************/

type pendingReasonEnum struct{ value_1oeqp39gm9qz uint8 }

var pendingReason = struct {
	payerShippingUnconfirmed pendingReasonEnum
	multiCurrency            pendingReasonEnum
	riskReview               pendingReasonEnum
	regulatoryReview         pendingReasonEnum
	verificationRequired     pendingReasonEnum
	orderPending             pendingReasonEnum
	otherPending             pendingReasonEnum
}{
	payerShippingUnconfirmed: pendingReasonEnum{value_1oeqp39gm9qz: 1},
	multiCurrency:            pendingReasonEnum{value_1oeqp39gm9qz: 2},
	riskReview:               pendingReasonEnum{value_1oeqp39gm9qz: 3},
	regulatoryReview:         pendingReasonEnum{value_1oeqp39gm9qz: 4},
	verificationRequired:     pendingReasonEnum{value_1oeqp39gm9qz: 5},
	orderPending:             pendingReasonEnum{value_1oeqp39gm9qz: 6},
	otherPending:             pendingReasonEnum{value_1oeqp39gm9qz: 7},
}

// Used to iterate in range loops
var pendingReasonValues = [...]pendingReasonEnum{
	pendingReason.payerShippingUnconfirmed, pendingReason.multiCurrency, pendingReason.riskReview, pendingReason.regulatoryReview, pendingReason.verificationRequired, pendingReason.orderPending, pendingReason.otherPending,
}

// Get the integer value of the enum variant
func (self pendingReasonEnum) Value() uint8 {
	return self.value_1oeqp39gm9qz
}

func (self pendingReasonEnum) IntValue() int {
	return int(self.value_1oeqp39gm9qz)
}

// Get the string representation of the enum variant
func (self pendingReasonEnum) String() string {
	switch self.value_1oeqp39gm9qz {
	case 1:
		return "PAYER-SHIPPING-UNCONFIRMED"
	case 2:
		return "MULTI-CURRENCY"
	case 3:
		return "RISK-REVIEW"
	case 4:
		return "REGULATORY-REVIEW"
	case 5:
		return "VERIFICATION-REQUIRED"
	case 6:
		return "ORDER"
	case 7:
		return "OTHER"
	}

	return ""
}

// Get the string description of the enum variant
func (self pendingReasonEnum) Description() string {
	switch self.value_1oeqp39gm9qz {
	case 1:
		return "PAYER-SHIPPING-UNCONFIRMED"
	case 2:
		return "MULTI-CURRENCY"
	case 3:
		return "RISK-REVIEW"
	case 4:
		return "REGULATORY-REVIEW"
	case 5:
		return "VERIFICATION-REQUIRED"
	case 6:
		return "ORDER"
	case 7:
		return "OTHER"
	}
	return ""
}

func (self pendingReasonEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *pendingReasonEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_1oeqp39gm9qz = uint8(n)
	return nil
}

/*****************************

CurrencyTypeEnum

******************************/

type CurrencyTypeEnum struct{ value_1eheyr0zxnr4f uint8 }

var CurrencyType = struct {
	AUD CurrencyTypeEnum
	BRL CurrencyTypeEnum
	CAD CurrencyTypeEnum
	CZK CurrencyTypeEnum
	DKK CurrencyTypeEnum
	EUR CurrencyTypeEnum
	HKD CurrencyTypeEnum
	HUF CurrencyTypeEnum
	ILS CurrencyTypeEnum
	JPY CurrencyTypeEnum
	MYR CurrencyTypeEnum
	MXN CurrencyTypeEnum
	TWD CurrencyTypeEnum
	NZD CurrencyTypeEnum
	NOK CurrencyTypeEnum
	PHP CurrencyTypeEnum
	PLN CurrencyTypeEnum
	GBP CurrencyTypeEnum
	SGD CurrencyTypeEnum
	SEK CurrencyTypeEnum
	CHF CurrencyTypeEnum
	THB CurrencyTypeEnum
	TRY CurrencyTypeEnum
	USD CurrencyTypeEnum
}{
	AUD: CurrencyTypeEnum{value_1eheyr0zxnr4f: 1},
	BRL: CurrencyTypeEnum{value_1eheyr0zxnr4f: 2},
	CAD: CurrencyTypeEnum{value_1eheyr0zxnr4f: 3},
	CZK: CurrencyTypeEnum{value_1eheyr0zxnr4f: 4},
	DKK: CurrencyTypeEnum{value_1eheyr0zxnr4f: 5},
	EUR: CurrencyTypeEnum{value_1eheyr0zxnr4f: 6},
	HKD: CurrencyTypeEnum{value_1eheyr0zxnr4f: 7},
	HUF: CurrencyTypeEnum{value_1eheyr0zxnr4f: 8},
	ILS: CurrencyTypeEnum{value_1eheyr0zxnr4f: 9},
	JPY: CurrencyTypeEnum{value_1eheyr0zxnr4f: 10},
	MYR: CurrencyTypeEnum{value_1eheyr0zxnr4f: 11},
	MXN: CurrencyTypeEnum{value_1eheyr0zxnr4f: 12},
	TWD: CurrencyTypeEnum{value_1eheyr0zxnr4f: 13},
	NZD: CurrencyTypeEnum{value_1eheyr0zxnr4f: 14},
	NOK: CurrencyTypeEnum{value_1eheyr0zxnr4f: 15},
	PHP: CurrencyTypeEnum{value_1eheyr0zxnr4f: 16},
	PLN: CurrencyTypeEnum{value_1eheyr0zxnr4f: 17},
	GBP: CurrencyTypeEnum{value_1eheyr0zxnr4f: 18},
	SGD: CurrencyTypeEnum{value_1eheyr0zxnr4f: 19},
	SEK: CurrencyTypeEnum{value_1eheyr0zxnr4f: 20},
	CHF: CurrencyTypeEnum{value_1eheyr0zxnr4f: 21},
	THB: CurrencyTypeEnum{value_1eheyr0zxnr4f: 22},
	TRY: CurrencyTypeEnum{value_1eheyr0zxnr4f: 23},
	USD: CurrencyTypeEnum{value_1eheyr0zxnr4f: 24},
}

// Used to iterate in range loops
var CurrencyTypeValues = [...]CurrencyTypeEnum{
	CurrencyType.AUD, CurrencyType.BRL, CurrencyType.CAD, CurrencyType.CZK, CurrencyType.DKK, CurrencyType.EUR, CurrencyType.HKD, CurrencyType.HUF, CurrencyType.ILS, CurrencyType.JPY, CurrencyType.MYR, CurrencyType.MXN, CurrencyType.TWD, CurrencyType.NZD, CurrencyType.NOK, CurrencyType.PHP, CurrencyType.PLN, CurrencyType.GBP, CurrencyType.SGD, CurrencyType.SEK, CurrencyType.CHF, CurrencyType.THB, CurrencyType.TRY, CurrencyType.USD,
}

// Get the integer value of the enum variant
func (self CurrencyTypeEnum) Value() uint8 {
	return self.value_1eheyr0zxnr4f
}

func (self CurrencyTypeEnum) IntValue() int {
	return int(self.value_1eheyr0zxnr4f)
}

// Get the string representation of the enum variant
func (self CurrencyTypeEnum) String() string {
	switch self.value_1eheyr0zxnr4f {
	case 1:
		return "AUD"
	case 2:
		return "BRL"
	case 3:
		return "CAD"
	case 4:
		return "CZK"
	case 5:
		return "DKK"
	case 6:
		return "EUR"
	case 7:
		return "HKD"
	case 8:
		return "HUF"
	case 9:
		return "ILS"
	case 10:
		return "JPY"
	case 11:
		return "MYR"
	case 12:
		return "MXN"
	case 13:
		return "TWD"
	case 14:
		return "NZD"
	case 15:
		return "NOK"
	case 16:
		return "PHP"
	case 17:
		return "PLN"
	case 18:
		return "GBP"
	case 19:
		return "SGD"
	case 20:
		return "SEK"
	case 21:
		return "CHF"
	case 22:
		return "THB"
	case 23:
		return "TRY"
	case 24:
		return "USD"
	}

	return ""
}

// Get the string description of the enum variant
func (self CurrencyTypeEnum) Description() string {
	switch self.value_1eheyr0zxnr4f {
	case 1:
		return "Australian dollar"
	case 2:
		return "Brazilian real**"
	case 3:
		return "Canadian dollar"
	case 4:
		return "Czech koruna"
	case 5:
		return "Danish krone"
	case 6:
		return "Euro"
	case 7:
		return "Hong Kong dollar"
	case 8:
		return "Hungarian forint"
	case 9:
		return "Israeli new shekel"
	case 10:
		return "Japanese yen*"
	case 11:
		return "Malaysian ringgit**"
	case 12:
		return "Mexican peso"
	case 13:
		return "New Taiwan dollar*"
	case 14:
		return "New Zealand dollar"
	case 15:
		return "Norwegian krone"
	case 16:
		return "Philippine peso"
	case 17:
		return "Polish złoty"
	case 18:
		return "Pound sterling"
	case 19:
		return "Singapore dollar"
	case 20:
		return "Swedish krona"
	case 21:
		return "Swiss franc"
	case 22:
		return "Thai baht"
	case 23:
		return "Turkish lira**"
	case 24:
		return "United States dollar"
	}
	return ""
}

func (self CurrencyTypeEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *CurrencyTypeEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_1eheyr0zxnr4f = uint8(n)
	return nil
}

/*****************************

CountryCodeEnum

******************************/

type CountryCodeEnum struct{ value_30bx5a8om6xs uint8 }

var CountryCode = struct {
	AL CountryCodeEnum
	DZ CountryCodeEnum
	AD CountryCodeEnum
	AO CountryCodeEnum
	AI CountryCodeEnum
	AG CountryCodeEnum
	AR CountryCodeEnum
	AM CountryCodeEnum
	AW CountryCodeEnum
	AU CountryCodeEnum
	AT CountryCodeEnum
	AZ CountryCodeEnum
	BS CountryCodeEnum
	BH CountryCodeEnum
	BB CountryCodeEnum
	BE CountryCodeEnum
	BZ CountryCodeEnum
	BJ CountryCodeEnum
	BM CountryCodeEnum
	BT CountryCodeEnum
	BO CountryCodeEnum
	BA CountryCodeEnum
	BW CountryCodeEnum
	BR CountryCodeEnum
	BN CountryCodeEnum
	BG CountryCodeEnum
	BF CountryCodeEnum
	BI CountryCodeEnum
	KH CountryCodeEnum
	CA CountryCodeEnum
	CV CountryCodeEnum
	KY CountryCodeEnum
	TD CountryCodeEnum
	CL CountryCodeEnum
	CN CountryCodeEnum
	C2 CountryCodeEnum
	CO CountryCodeEnum
	KM CountryCodeEnum
	CD CountryCodeEnum
	CG CountryCodeEnum
	CK CountryCodeEnum
	CR CountryCodeEnum
	HR CountryCodeEnum
	CY CountryCodeEnum
	CZ CountryCodeEnum
	DK CountryCodeEnum
	DJ CountryCodeEnum
	DM CountryCodeEnum
	DO CountryCodeEnum
	EC CountryCodeEnum
	EG CountryCodeEnum
	SV CountryCodeEnum
	ER CountryCodeEnum
	EE CountryCodeEnum
	ET CountryCodeEnum
	FK CountryCodeEnum
	FJ CountryCodeEnum
	FI CountryCodeEnum
	FR CountryCodeEnum
	GF CountryCodeEnum
	PF CountryCodeEnum
	GA CountryCodeEnum
	GM CountryCodeEnum
	GE CountryCodeEnum
	DE CountryCodeEnum
	GI CountryCodeEnum
	GR CountryCodeEnum
	GL CountryCodeEnum
	GD CountryCodeEnum
	GP CountryCodeEnum
	GU CountryCodeEnum
	GT CountryCodeEnum
	GN CountryCodeEnum
	GW CountryCodeEnum
	GY CountryCodeEnum
	VA CountryCodeEnum
	HN CountryCodeEnum
	HK CountryCodeEnum
	HU CountryCodeEnum
	IS CountryCodeEnum
	IN CountryCodeEnum
	ID CountryCodeEnum
	IE CountryCodeEnum
	IL CountryCodeEnum
	IT CountryCodeEnum
	JM CountryCodeEnum
	JP CountryCodeEnum
	JO CountryCodeEnum
	KZ CountryCodeEnum
	KE CountryCodeEnum
	KI CountryCodeEnum
	KR CountryCodeEnum
	KW CountryCodeEnum
	KG CountryCodeEnum
	LA CountryCodeEnum
	LV CountryCodeEnum
	LS CountryCodeEnum
	LI CountryCodeEnum
	LT CountryCodeEnum
	LU CountryCodeEnum
	MG CountryCodeEnum
	MW CountryCodeEnum
	MY CountryCodeEnum
	MV CountryCodeEnum
	ML CountryCodeEnum
	MT CountryCodeEnum
	MH CountryCodeEnum
	MQ CountryCodeEnum
	MR CountryCodeEnum
	MU CountryCodeEnum
	YT CountryCodeEnum
	MX CountryCodeEnum
	FM CountryCodeEnum
	MN CountryCodeEnum
	MS CountryCodeEnum
	MA CountryCodeEnum
	MZ CountryCodeEnum
	NA CountryCodeEnum
	NR CountryCodeEnum
	NP CountryCodeEnum
	NL CountryCodeEnum
	AN CountryCodeEnum
	NC CountryCodeEnum
	NZ CountryCodeEnum
	NI CountryCodeEnum
	NE CountryCodeEnum
	NU CountryCodeEnum
	NF CountryCodeEnum
	NO CountryCodeEnum
	OM CountryCodeEnum
	PW CountryCodeEnum
	PA CountryCodeEnum
	PG CountryCodeEnum
	PE CountryCodeEnum
	PH CountryCodeEnum
	PN CountryCodeEnum
	PL CountryCodeEnum
	PT CountryCodeEnum
	QA CountryCodeEnum
	RE CountryCodeEnum
	RO CountryCodeEnum
	RU CountryCodeEnum
	RW CountryCodeEnum
	SH CountryCodeEnum
	KN CountryCodeEnum
	LC CountryCodeEnum
	PM CountryCodeEnum
	VC CountryCodeEnum
	WS CountryCodeEnum
	SM CountryCodeEnum
	ST CountryCodeEnum
	SA CountryCodeEnum
	SN CountryCodeEnum
	RS CountryCodeEnum
	SC CountryCodeEnum
	SL CountryCodeEnum
	SG CountryCodeEnum
	SK CountryCodeEnum
	SI CountryCodeEnum
	SB CountryCodeEnum
	SO CountryCodeEnum
	ZA CountryCodeEnum
	ES CountryCodeEnum
	LK CountryCodeEnum
	SR CountryCodeEnum
	SJ CountryCodeEnum
	SZ CountryCodeEnum
	SE CountryCodeEnum
	CH CountryCodeEnum
	TW CountryCodeEnum
	TJ CountryCodeEnum
	TZ CountryCodeEnum
	TH CountryCodeEnum
	TG CountryCodeEnum
	TO CountryCodeEnum
	TT CountryCodeEnum
	TN CountryCodeEnum
	TR CountryCodeEnum
	TM CountryCodeEnum
	TC CountryCodeEnum
	TV CountryCodeEnum
	UG CountryCodeEnum
	UA CountryCodeEnum
	AE CountryCodeEnum
	GB CountryCodeEnum
	US CountryCodeEnum
	UY CountryCodeEnum
	VU CountryCodeEnum
	VE CountryCodeEnum
	VN CountryCodeEnum
	VG CountryCodeEnum
	WF CountryCodeEnum
	YE CountryCodeEnum
	ZM CountryCodeEnum
}{
	AL: CountryCodeEnum{value_30bx5a8om6xs: 1},
	DZ: CountryCodeEnum{value_30bx5a8om6xs: 2},
	AD: CountryCodeEnum{value_30bx5a8om6xs: 3},
	AO: CountryCodeEnum{value_30bx5a8om6xs: 4},
	AI: CountryCodeEnum{value_30bx5a8om6xs: 5},
	AG: CountryCodeEnum{value_30bx5a8om6xs: 6},
	AR: CountryCodeEnum{value_30bx5a8om6xs: 7},
	AM: CountryCodeEnum{value_30bx5a8om6xs: 8},
	AW: CountryCodeEnum{value_30bx5a8om6xs: 9},
	AU: CountryCodeEnum{value_30bx5a8om6xs: 10},
	AT: CountryCodeEnum{value_30bx5a8om6xs: 11},
	AZ: CountryCodeEnum{value_30bx5a8om6xs: 12},
	BS: CountryCodeEnum{value_30bx5a8om6xs: 13},
	BH: CountryCodeEnum{value_30bx5a8om6xs: 14},
	BB: CountryCodeEnum{value_30bx5a8om6xs: 15},
	BE: CountryCodeEnum{value_30bx5a8om6xs: 16},
	BZ: CountryCodeEnum{value_30bx5a8om6xs: 17},
	BJ: CountryCodeEnum{value_30bx5a8om6xs: 18},
	BM: CountryCodeEnum{value_30bx5a8om6xs: 19},
	BT: CountryCodeEnum{value_30bx5a8om6xs: 20},
	BO: CountryCodeEnum{value_30bx5a8om6xs: 21},
	BA: CountryCodeEnum{value_30bx5a8om6xs: 22},
	BW: CountryCodeEnum{value_30bx5a8om6xs: 23},
	BR: CountryCodeEnum{value_30bx5a8om6xs: 24},
	BN: CountryCodeEnum{value_30bx5a8om6xs: 25},
	BG: CountryCodeEnum{value_30bx5a8om6xs: 26},
	BF: CountryCodeEnum{value_30bx5a8om6xs: 27},
	BI: CountryCodeEnum{value_30bx5a8om6xs: 28},
	KH: CountryCodeEnum{value_30bx5a8om6xs: 29},
	CA: CountryCodeEnum{value_30bx5a8om6xs: 30},
	CV: CountryCodeEnum{value_30bx5a8om6xs: 31},
	KY: CountryCodeEnum{value_30bx5a8om6xs: 32},
	TD: CountryCodeEnum{value_30bx5a8om6xs: 33},
	CL: CountryCodeEnum{value_30bx5a8om6xs: 34},
	CN: CountryCodeEnum{value_30bx5a8om6xs: 35},
	C2: CountryCodeEnum{value_30bx5a8om6xs: 36},
	CO: CountryCodeEnum{value_30bx5a8om6xs: 37},
	KM: CountryCodeEnum{value_30bx5a8om6xs: 38},
	CD: CountryCodeEnum{value_30bx5a8om6xs: 39},
	CG: CountryCodeEnum{value_30bx5a8om6xs: 40},
	CK: CountryCodeEnum{value_30bx5a8om6xs: 41},
	CR: CountryCodeEnum{value_30bx5a8om6xs: 42},
	HR: CountryCodeEnum{value_30bx5a8om6xs: 43},
	CY: CountryCodeEnum{value_30bx5a8om6xs: 44},
	CZ: CountryCodeEnum{value_30bx5a8om6xs: 45},
	DK: CountryCodeEnum{value_30bx5a8om6xs: 46},
	DJ: CountryCodeEnum{value_30bx5a8om6xs: 47},
	DM: CountryCodeEnum{value_30bx5a8om6xs: 48},
	DO: CountryCodeEnum{value_30bx5a8om6xs: 49},
	EC: CountryCodeEnum{value_30bx5a8om6xs: 50},
	EG: CountryCodeEnum{value_30bx5a8om6xs: 51},
	SV: CountryCodeEnum{value_30bx5a8om6xs: 52},
	ER: CountryCodeEnum{value_30bx5a8om6xs: 53},
	EE: CountryCodeEnum{value_30bx5a8om6xs: 54},
	ET: CountryCodeEnum{value_30bx5a8om6xs: 55},
	FK: CountryCodeEnum{value_30bx5a8om6xs: 56},
	FJ: CountryCodeEnum{value_30bx5a8om6xs: 57},
	FI: CountryCodeEnum{value_30bx5a8om6xs: 58},
	FR: CountryCodeEnum{value_30bx5a8om6xs: 59},
	GF: CountryCodeEnum{value_30bx5a8om6xs: 60},
	PF: CountryCodeEnum{value_30bx5a8om6xs: 61},
	GA: CountryCodeEnum{value_30bx5a8om6xs: 62},
	GM: CountryCodeEnum{value_30bx5a8om6xs: 63},
	GE: CountryCodeEnum{value_30bx5a8om6xs: 64},
	DE: CountryCodeEnum{value_30bx5a8om6xs: 65},
	GI: CountryCodeEnum{value_30bx5a8om6xs: 66},
	GR: CountryCodeEnum{value_30bx5a8om6xs: 67},
	GL: CountryCodeEnum{value_30bx5a8om6xs: 68},
	GD: CountryCodeEnum{value_30bx5a8om6xs: 69},
	GP: CountryCodeEnum{value_30bx5a8om6xs: 70},
	GU: CountryCodeEnum{value_30bx5a8om6xs: 71},
	GT: CountryCodeEnum{value_30bx5a8om6xs: 72},
	GN: CountryCodeEnum{value_30bx5a8om6xs: 73},
	GW: CountryCodeEnum{value_30bx5a8om6xs: 74},
	GY: CountryCodeEnum{value_30bx5a8om6xs: 75},
	VA: CountryCodeEnum{value_30bx5a8om6xs: 76},
	HN: CountryCodeEnum{value_30bx5a8om6xs: 77},
	HK: CountryCodeEnum{value_30bx5a8om6xs: 78},
	HU: CountryCodeEnum{value_30bx5a8om6xs: 79},
	IS: CountryCodeEnum{value_30bx5a8om6xs: 80},
	IN: CountryCodeEnum{value_30bx5a8om6xs: 81},
	ID: CountryCodeEnum{value_30bx5a8om6xs: 82},
	IE: CountryCodeEnum{value_30bx5a8om6xs: 83},
	IL: CountryCodeEnum{value_30bx5a8om6xs: 84},
	IT: CountryCodeEnum{value_30bx5a8om6xs: 85},
	JM: CountryCodeEnum{value_30bx5a8om6xs: 86},
	JP: CountryCodeEnum{value_30bx5a8om6xs: 87},
	JO: CountryCodeEnum{value_30bx5a8om6xs: 88},
	KZ: CountryCodeEnum{value_30bx5a8om6xs: 89},
	KE: CountryCodeEnum{value_30bx5a8om6xs: 90},
	KI: CountryCodeEnum{value_30bx5a8om6xs: 91},
	KR: CountryCodeEnum{value_30bx5a8om6xs: 92},
	KW: CountryCodeEnum{value_30bx5a8om6xs: 93},
	KG: CountryCodeEnum{value_30bx5a8om6xs: 94},
	LA: CountryCodeEnum{value_30bx5a8om6xs: 95},
	LV: CountryCodeEnum{value_30bx5a8om6xs: 96},
	LS: CountryCodeEnum{value_30bx5a8om6xs: 97},
	LI: CountryCodeEnum{value_30bx5a8om6xs: 98},
	LT: CountryCodeEnum{value_30bx5a8om6xs: 99},
	LU: CountryCodeEnum{value_30bx5a8om6xs: 100},
	MG: CountryCodeEnum{value_30bx5a8om6xs: 101},
	MW: CountryCodeEnum{value_30bx5a8om6xs: 102},
	MY: CountryCodeEnum{value_30bx5a8om6xs: 103},
	MV: CountryCodeEnum{value_30bx5a8om6xs: 104},
	ML: CountryCodeEnum{value_30bx5a8om6xs: 105},
	MT: CountryCodeEnum{value_30bx5a8om6xs: 106},
	MH: CountryCodeEnum{value_30bx5a8om6xs: 107},
	MQ: CountryCodeEnum{value_30bx5a8om6xs: 108},
	MR: CountryCodeEnum{value_30bx5a8om6xs: 109},
	MU: CountryCodeEnum{value_30bx5a8om6xs: 110},
	YT: CountryCodeEnum{value_30bx5a8om6xs: 111},
	MX: CountryCodeEnum{value_30bx5a8om6xs: 112},
	FM: CountryCodeEnum{value_30bx5a8om6xs: 113},
	MN: CountryCodeEnum{value_30bx5a8om6xs: 114},
	MS: CountryCodeEnum{value_30bx5a8om6xs: 115},
	MA: CountryCodeEnum{value_30bx5a8om6xs: 116},
	MZ: CountryCodeEnum{value_30bx5a8om6xs: 117},
	NA: CountryCodeEnum{value_30bx5a8om6xs: 118},
	NR: CountryCodeEnum{value_30bx5a8om6xs: 119},
	NP: CountryCodeEnum{value_30bx5a8om6xs: 120},
	NL: CountryCodeEnum{value_30bx5a8om6xs: 121},
	AN: CountryCodeEnum{value_30bx5a8om6xs: 122},
	NC: CountryCodeEnum{value_30bx5a8om6xs: 123},
	NZ: CountryCodeEnum{value_30bx5a8om6xs: 124},
	NI: CountryCodeEnum{value_30bx5a8om6xs: 125},
	NE: CountryCodeEnum{value_30bx5a8om6xs: 126},
	NU: CountryCodeEnum{value_30bx5a8om6xs: 127},
	NF: CountryCodeEnum{value_30bx5a8om6xs: 128},
	NO: CountryCodeEnum{value_30bx5a8om6xs: 129},
	OM: CountryCodeEnum{value_30bx5a8om6xs: 130},
	PW: CountryCodeEnum{value_30bx5a8om6xs: 131},
	PA: CountryCodeEnum{value_30bx5a8om6xs: 132},
	PG: CountryCodeEnum{value_30bx5a8om6xs: 133},
	PE: CountryCodeEnum{value_30bx5a8om6xs: 134},
	PH: CountryCodeEnum{value_30bx5a8om6xs: 135},
	PN: CountryCodeEnum{value_30bx5a8om6xs: 136},
	PL: CountryCodeEnum{value_30bx5a8om6xs: 137},
	PT: CountryCodeEnum{value_30bx5a8om6xs: 138},
	QA: CountryCodeEnum{value_30bx5a8om6xs: 139},
	RE: CountryCodeEnum{value_30bx5a8om6xs: 140},
	RO: CountryCodeEnum{value_30bx5a8om6xs: 141},
	RU: CountryCodeEnum{value_30bx5a8om6xs: 142},
	RW: CountryCodeEnum{value_30bx5a8om6xs: 143},
	SH: CountryCodeEnum{value_30bx5a8om6xs: 144},
	KN: CountryCodeEnum{value_30bx5a8om6xs: 145},
	LC: CountryCodeEnum{value_30bx5a8om6xs: 146},
	PM: CountryCodeEnum{value_30bx5a8om6xs: 147},
	VC: CountryCodeEnum{value_30bx5a8om6xs: 148},
	WS: CountryCodeEnum{value_30bx5a8om6xs: 149},
	SM: CountryCodeEnum{value_30bx5a8om6xs: 150},
	ST: CountryCodeEnum{value_30bx5a8om6xs: 151},
	SA: CountryCodeEnum{value_30bx5a8om6xs: 152},
	SN: CountryCodeEnum{value_30bx5a8om6xs: 153},
	RS: CountryCodeEnum{value_30bx5a8om6xs: 154},
	SC: CountryCodeEnum{value_30bx5a8om6xs: 155},
	SL: CountryCodeEnum{value_30bx5a8om6xs: 156},
	SG: CountryCodeEnum{value_30bx5a8om6xs: 157},
	SK: CountryCodeEnum{value_30bx5a8om6xs: 158},
	SI: CountryCodeEnum{value_30bx5a8om6xs: 159},
	SB: CountryCodeEnum{value_30bx5a8om6xs: 160},
	SO: CountryCodeEnum{value_30bx5a8om6xs: 161},
	ZA: CountryCodeEnum{value_30bx5a8om6xs: 162},
	ES: CountryCodeEnum{value_30bx5a8om6xs: 163},
	LK: CountryCodeEnum{value_30bx5a8om6xs: 164},
	SR: CountryCodeEnum{value_30bx5a8om6xs: 165},
	SJ: CountryCodeEnum{value_30bx5a8om6xs: 166},
	SZ: CountryCodeEnum{value_30bx5a8om6xs: 167},
	SE: CountryCodeEnum{value_30bx5a8om6xs: 168},
	CH: CountryCodeEnum{value_30bx5a8om6xs: 169},
	TW: CountryCodeEnum{value_30bx5a8om6xs: 170},
	TJ: CountryCodeEnum{value_30bx5a8om6xs: 171},
	TZ: CountryCodeEnum{value_30bx5a8om6xs: 172},
	TH: CountryCodeEnum{value_30bx5a8om6xs: 173},
	TG: CountryCodeEnum{value_30bx5a8om6xs: 174},
	TO: CountryCodeEnum{value_30bx5a8om6xs: 175},
	TT: CountryCodeEnum{value_30bx5a8om6xs: 176},
	TN: CountryCodeEnum{value_30bx5a8om6xs: 177},
	TR: CountryCodeEnum{value_30bx5a8om6xs: 178},
	TM: CountryCodeEnum{value_30bx5a8om6xs: 179},
	TC: CountryCodeEnum{value_30bx5a8om6xs: 180},
	TV: CountryCodeEnum{value_30bx5a8om6xs: 181},
	UG: CountryCodeEnum{value_30bx5a8om6xs: 182},
	UA: CountryCodeEnum{value_30bx5a8om6xs: 183},
	AE: CountryCodeEnum{value_30bx5a8om6xs: 184},
	GB: CountryCodeEnum{value_30bx5a8om6xs: 185},
	US: CountryCodeEnum{value_30bx5a8om6xs: 186},
	UY: CountryCodeEnum{value_30bx5a8om6xs: 187},
	VU: CountryCodeEnum{value_30bx5a8om6xs: 188},
	VE: CountryCodeEnum{value_30bx5a8om6xs: 189},
	VN: CountryCodeEnum{value_30bx5a8om6xs: 190},
	VG: CountryCodeEnum{value_30bx5a8om6xs: 191},
	WF: CountryCodeEnum{value_30bx5a8om6xs: 192},
	YE: CountryCodeEnum{value_30bx5a8om6xs: 193},
	ZM: CountryCodeEnum{value_30bx5a8om6xs: 194},
}

// Used to iterate in range loops
var CountryCodeValues = [...]CountryCodeEnum{
	CountryCode.AL, CountryCode.DZ, CountryCode.AD, CountryCode.AO, CountryCode.AI, CountryCode.AG, CountryCode.AR, CountryCode.AM, CountryCode.AW, CountryCode.AU, CountryCode.AT, CountryCode.AZ, CountryCode.BS, CountryCode.BH, CountryCode.BB, CountryCode.BE, CountryCode.BZ, CountryCode.BJ, CountryCode.BM, CountryCode.BT, CountryCode.BO, CountryCode.BA, CountryCode.BW, CountryCode.BR, CountryCode.BN, CountryCode.BG, CountryCode.BF, CountryCode.BI, CountryCode.KH, CountryCode.CA, CountryCode.CV, CountryCode.KY, CountryCode.TD, CountryCode.CL, CountryCode.CN, CountryCode.C2, CountryCode.CO, CountryCode.KM, CountryCode.CD, CountryCode.CG, CountryCode.CK, CountryCode.CR, CountryCode.HR, CountryCode.CY, CountryCode.CZ, CountryCode.DK, CountryCode.DJ, CountryCode.DM, CountryCode.DO, CountryCode.EC, CountryCode.EG, CountryCode.SV, CountryCode.ER, CountryCode.EE, CountryCode.ET, CountryCode.FK, CountryCode.FJ, CountryCode.FI, CountryCode.FR, CountryCode.GF, CountryCode.PF, CountryCode.GA, CountryCode.GM, CountryCode.GE, CountryCode.DE, CountryCode.GI, CountryCode.GR, CountryCode.GL, CountryCode.GD, CountryCode.GP, CountryCode.GU, CountryCode.GT, CountryCode.GN, CountryCode.GW, CountryCode.GY, CountryCode.VA, CountryCode.HN, CountryCode.HK, CountryCode.HU, CountryCode.IS, CountryCode.IN, CountryCode.ID, CountryCode.IE, CountryCode.IL, CountryCode.IT, CountryCode.JM, CountryCode.JP, CountryCode.JO, CountryCode.KZ, CountryCode.KE, CountryCode.KI, CountryCode.KR, CountryCode.KW, CountryCode.KG, CountryCode.LA, CountryCode.LV, CountryCode.LS, CountryCode.LI, CountryCode.LT, CountryCode.LU, CountryCode.MG, CountryCode.MW, CountryCode.MY, CountryCode.MV, CountryCode.ML, CountryCode.MT, CountryCode.MH, CountryCode.MQ, CountryCode.MR, CountryCode.MU, CountryCode.YT, CountryCode.MX, CountryCode.FM, CountryCode.MN, CountryCode.MS, CountryCode.MA, CountryCode.MZ, CountryCode.NA, CountryCode.NR, CountryCode.NP, CountryCode.NL, CountryCode.AN, CountryCode.NC, CountryCode.NZ, CountryCode.NI, CountryCode.NE, CountryCode.NU, CountryCode.NF, CountryCode.NO, CountryCode.OM, CountryCode.PW, CountryCode.PA, CountryCode.PG, CountryCode.PE, CountryCode.PH, CountryCode.PN, CountryCode.PL, CountryCode.PT, CountryCode.QA, CountryCode.RE, CountryCode.RO, CountryCode.RU, CountryCode.RW, CountryCode.SH, CountryCode.KN, CountryCode.LC, CountryCode.PM, CountryCode.VC, CountryCode.WS, CountryCode.SM, CountryCode.ST, CountryCode.SA, CountryCode.SN, CountryCode.RS, CountryCode.SC, CountryCode.SL, CountryCode.SG, CountryCode.SK, CountryCode.SI, CountryCode.SB, CountryCode.SO, CountryCode.ZA, CountryCode.ES, CountryCode.LK, CountryCode.SR, CountryCode.SJ, CountryCode.SZ, CountryCode.SE, CountryCode.CH, CountryCode.TW, CountryCode.TJ, CountryCode.TZ, CountryCode.TH, CountryCode.TG, CountryCode.TO, CountryCode.TT, CountryCode.TN, CountryCode.TR, CountryCode.TM, CountryCode.TC, CountryCode.TV, CountryCode.UG, CountryCode.UA, CountryCode.AE, CountryCode.GB, CountryCode.US, CountryCode.UY, CountryCode.VU, CountryCode.VE, CountryCode.VN, CountryCode.VG, CountryCode.WF, CountryCode.YE, CountryCode.ZM,
}

// Get the integer value of the enum variant
func (self CountryCodeEnum) Value() uint8 {
	return self.value_30bx5a8om6xs
}

func (self CountryCodeEnum) IntValue() int {
	return int(self.value_30bx5a8om6xs)
}

// Get the string representation of the enum variant
func (self CountryCodeEnum) String() string {
	switch self.value_30bx5a8om6xs {
	case 1:
		return "AL"
	case 2:
		return "DZ"
	case 3:
		return "AD"
	case 4:
		return "AO"
	case 5:
		return "AI"
	case 6:
		return "AG"
	case 7:
		return "AR"
	case 8:
		return "AM"
	case 9:
		return "AW"
	case 10:
		return "AU"
	case 11:
		return "AT"
	case 12:
		return "AZ"
	case 13:
		return "BS"
	case 14:
		return "BH"
	case 15:
		return "BB"
	case 16:
		return "BE"
	case 17:
		return "BZ"
	case 18:
		return "BJ"
	case 19:
		return "BM"
	case 20:
		return "BT"
	case 21:
		return "BO"
	case 22:
		return "BA"
	case 23:
		return "BW"
	case 24:
		return "BR"
	case 25:
		return "BN"
	case 26:
		return "BG"
	case 27:
		return "BF"
	case 28:
		return "BI"
	case 29:
		return "KH"
	case 30:
		return "CA"
	case 31:
		return "CV"
	case 32:
		return "KY"
	case 33:
		return "TD"
	case 34:
		return "CL"
	case 35:
		return "CN"
	case 36:
		return "C2"
	case 37:
		return "CO"
	case 38:
		return "KM"
	case 39:
		return "CD"
	case 40:
		return "CG"
	case 41:
		return "CK"
	case 42:
		return "CR"
	case 43:
		return "HR"
	case 44:
		return "CY"
	case 45:
		return "CZ"
	case 46:
		return "DK"
	case 47:
		return "DJ"
	case 48:
		return "DM"
	case 49:
		return "DO"
	case 50:
		return "EC"
	case 51:
		return "EG"
	case 52:
		return "SV"
	case 53:
		return "ER"
	case 54:
		return "EE"
	case 55:
		return "ET"
	case 56:
		return "FK"
	case 57:
		return "FJ"
	case 58:
		return "FI"
	case 59:
		return "FR"
	case 60:
		return "GF"
	case 61:
		return "PF"
	case 62:
		return "GA"
	case 63:
		return "GM"
	case 64:
		return "GE"
	case 65:
		return "DE"
	case 66:
		return "GI"
	case 67:
		return "GR"
	case 68:
		return "GL"
	case 69:
		return "GD"
	case 70:
		return "GP"
	case 71:
		return "GU"
	case 72:
		return "GT"
	case 73:
		return "GN"
	case 74:
		return "GW"
	case 75:
		return "GY"
	case 76:
		return "VA"
	case 77:
		return "HN"
	case 78:
		return "HK"
	case 79:
		return "HU"
	case 80:
		return "IS"
	case 81:
		return "IN"
	case 82:
		return "ID"
	case 83:
		return "IE"
	case 84:
		return "IL"
	case 85:
		return "IT"
	case 86:
		return "JM"
	case 87:
		return "JP"
	case 88:
		return "JO"
	case 89:
		return "KZ"
	case 90:
		return "KE"
	case 91:
		return "KI"
	case 92:
		return "KR"
	case 93:
		return "KW"
	case 94:
		return "KG"
	case 95:
		return "LA"
	case 96:
		return "LV"
	case 97:
		return "LS"
	case 98:
		return "LI"
	case 99:
		return "LT"
	case 100:
		return "LU"
	case 101:
		return "MG"
	case 102:
		return "MW"
	case 103:
		return "MY"
	case 104:
		return "MV"
	case 105:
		return "ML"
	case 106:
		return "MT"
	case 107:
		return "MH"
	case 108:
		return "MQ"
	case 109:
		return "MR"
	case 110:
		return "MU"
	case 111:
		return "YT"
	case 112:
		return "MX"
	case 113:
		return "FM"
	case 114:
		return "MN"
	case 115:
		return "MS"
	case 116:
		return "MA"
	case 117:
		return "MZ"
	case 118:
		return "NA"
	case 119:
		return "NR"
	case 120:
		return "NP"
	case 121:
		return "NL"
	case 122:
		return "AN"
	case 123:
		return "NC"
	case 124:
		return "NZ"
	case 125:
		return "NI"
	case 126:
		return "NE"
	case 127:
		return "NU"
	case 128:
		return "NF"
	case 129:
		return "NO"
	case 130:
		return "OM"
	case 131:
		return "PW"
	case 132:
		return "PA"
	case 133:
		return "PG"
	case 134:
		return "PE"
	case 135:
		return "PH"
	case 136:
		return "PN"
	case 137:
		return "PL"
	case 138:
		return "PT"
	case 139:
		return "QA"
	case 140:
		return "RE"
	case 141:
		return "RO"
	case 142:
		return "RU"
	case 143:
		return "RW"
	case 144:
		return "SH"
	case 145:
		return "KN"
	case 146:
		return "LC"
	case 147:
		return "PM"
	case 148:
		return "VC"
	case 149:
		return "WS"
	case 150:
		return "SM"
	case 151:
		return "ST"
	case 152:
		return "SA"
	case 153:
		return "SN"
	case 154:
		return "RS"
	case 155:
		return "SC"
	case 156:
		return "SL"
	case 157:
		return "SG"
	case 158:
		return "SK"
	case 159:
		return "SI"
	case 160:
		return "SB"
	case 161:
		return "SO"
	case 162:
		return "ZA"
	case 163:
		return "ES"
	case 164:
		return "LK"
	case 165:
		return "SR"
	case 166:
		return "SJ"
	case 167:
		return "SZ"
	case 168:
		return "SE"
	case 169:
		return "CH"
	case 170:
		return "TW"
	case 171:
		return "TJ"
	case 172:
		return "TZ"
	case 173:
		return "TH"
	case 174:
		return "TG"
	case 175:
		return "TO"
	case 176:
		return "TT"
	case 177:
		return "TN"
	case 178:
		return "TR"
	case 179:
		return "TM"
	case 180:
		return "TC"
	case 181:
		return "TV"
	case 182:
		return "UG"
	case 183:
		return "UA"
	case 184:
		return "AE"
	case 185:
		return "GB"
	case 186:
		return "US"
	case 187:
		return "UY"
	case 188:
		return "VU"
	case 189:
		return "VE"
	case 190:
		return "VN"
	case 191:
		return "VG"
	case 192:
		return "WF"
	case 193:
		return "YE"
	case 194:
		return "ZM"
	}

	return ""
}

// Get the string description of the enum variant
func (self CountryCodeEnum) Description() string {
	switch self.value_30bx5a8om6xs {
	case 1:
		return "ALBANIA"
	case 2:
		return "ALGERIA"
	case 3:
		return "ANDORRA"
	case 4:
		return "ANGOLA"
	case 5:
		return "ANGUILLA"
	case 6:
		return "ANTIGUA AND BARBUDA"
	case 7:
		return "ARGENTINA"
	case 8:
		return "ARMENIA"
	case 9:
		return "ARUBA"
	case 10:
		return "AUSTRALIA"
	case 11:
		return "AUSTRIA"
	case 12:
		return "AZERBAIJAN"
	case 13:
		return "BAHAMAS"
	case 14:
		return "BAHRAIN"
	case 15:
		return "BARBADOS"
	case 16:
		return "BELGIUM"
	case 17:
		return "BELIZE"
	case 18:
		return "BENIN"
	case 19:
		return "BERMUDA"
	case 20:
		return "BHUTAN"
	case 21:
		return "BOLIVIA"
	case 22:
		return "BOSNIA-HERZEGOVINA"
	case 23:
		return "BOTSWANA"
	case 24:
		return "BRAZIL"
	case 25:
		return "BRUNEI DARUSSALAM"
	case 26:
		return "BULGARIA"
	case 27:
		return "BURKINA FASO"
	case 28:
		return "BURUNDI"
	case 29:
		return "CAMBODIA"
	case 30:
		return "CANADA"
	case 31:
		return "CAPE VERDE"
	case 32:
		return "CAYMAN ISLANDS"
	case 33:
		return "CHAD"
	case 34:
		return "CHILE"
	case 35:
		return "CHINA (For domestic Chinese bank transactions only)"
	case 36:
		return "CHINA (For CUP, bank card and cross-border transactions)"
	case 37:
		return "COLOMBIA"
	case 38:
		return "COMOROS"
	case 39:
		return "DEMOCRATIC REPUBLIC OF CONGO"
	case 40:
		return "CONGO"
	case 41:
		return "COOK ISLANDS"
	case 42:
		return "COSTA RICA"
	case 43:
		return "CROATIA"
	case 44:
		return "CYPRUS"
	case 45:
		return "CZECH REPUBLIC"
	case 46:
		return "DENMARK"
	case 47:
		return "DJIBOUTI"
	case 48:
		return "DOMINICA"
	case 49:
		return "DOMINICAN REPUBLIC"
	case 50:
		return "ECUADOR"
	case 51:
		return "EGYPT"
	case 52:
		return "EL SALVADOR"
	case 53:
		return "ERITERIA"
	case 54:
		return "ESTONIA"
	case 55:
		return "ETHIOPIA"
	case 56:
		return "FALKLAND ISLANDS (MALVINAS)"
	case 57:
		return "FIJI"
	case 58:
		return "FINLAND"
	case 59:
		return "FRANCE"
	case 60:
		return "FRENCH GUIANA"
	case 61:
		return "FRENCH POLYNESIA"
	case 62:
		return "GABON"
	case 63:
		return "GAMBIA"
	case 64:
		return "GEORGIA"
	case 65:
		return "GERMANY"
	case 66:
		return "GIBRALTAR"
	case 67:
		return "GREECE"
	case 68:
		return "GREENLAND"
	case 69:
		return "GRENADA"
	case 70:
		return "GUADELOUPE"
	case 71:
		return "GUAM"
	case 72:
		return "GUATEMALA"
	case 73:
		return "GUINEA"
	case 74:
		return "GUINEA BISSAU"
	case 75:
		return "GUYANA"
	case 76:
		return "HOLY SEE (VATICAN CITY STATE)"
	case 77:
		return "HONDURAS"
	case 78:
		return "HONG KONG"
	case 79:
		return "HUNGARY"
	case 80:
		return "ICELAND"
	case 81:
		return "INDIA"
	case 82:
		return "INDONESIA"
	case 83:
		return "IRELAND"
	case 84:
		return "ISRAEL"
	case 85:
		return "ITALY"
	case 86:
		return "JAMAICA"
	case 87:
		return "JAPAN"
	case 88:
		return "JORDAN"
	case 89:
		return "KAZAKHSTAN"
	case 90:
		return "KENYA"
	case 91:
		return "KIRIBATI"
	case 92:
		return "KOREA, REPUBLIC OF (or SOUTH KOREA)"
	case 93:
		return "KUWAIT"
	case 94:
		return "KYRGYZSTAN"
	case 95:
		return "LAOS"
	case 96:
		return "LATVIA"
	case 97:
		return "LESOTHO"
	case 98:
		return "LIECHTENSTEIN"
	case 99:
		return "LITHUANIA"
	case 100:
		return "LUXEMBOURG"
	case 101:
		return "MADAGASCAR"
	case 102:
		return "MALAWI"
	case 103:
		return "MALAYSIA"
	case 104:
		return "MALDIVES"
	case 105:
		return "MALI"
	case 106:
		return "MALTA"
	case 107:
		return "MARSHALL ISLANDS"
	case 108:
		return "MARTINIQUE"
	case 109:
		return "MAURITANIA"
	case 110:
		return "MAURITIUS"
	case 111:
		return "MAYOTTE"
	case 112:
		return "MEXICO"
	case 113:
		return "MICRONESIA, FEDERATED STATES OF"
	case 114:
		return "MONGOLIA"
	case 115:
		return "MONTSERRAT"
	case 116:
		return "MOROCCO"
	case 117:
		return "MOZAMBIQUE"
	case 118:
		return "NAMIBIA"
	case 119:
		return "NAURU"
	case 120:
		return "NEPAL"
	case 121:
		return "NETHERLANDS"
	case 122:
		return "NETHERLANDS ANTILLES"
	case 123:
		return "NEW CALEDONIA"
	case 124:
		return "NEW ZEALAND"
	case 125:
		return "NICARAGUA"
	case 126:
		return "NIGER"
	case 127:
		return "NIUE"
	case 128:
		return "NORFOLK ISLAND"
	case 129:
		return "NORWAY"
	case 130:
		return "OMAN"
	case 131:
		return "PALAU"
	case 132:
		return "PANAMA"
	case 133:
		return "PAPUA NEW GUINEA"
	case 134:
		return "PERU"
	case 135:
		return "PHILIPPINES"
	case 136:
		return "PITCAIRN"
	case 137:
		return "POLAND"
	case 138:
		return "PORTUGAL"
	case 139:
		return "QATAR"
	case 140:
		return "REUNION"
	case 141:
		return "ROMANIA"
	case 142:
		return "RUSSIAN FEDERATION"
	case 143:
		return "RWANDA"
	case 144:
		return "SAINT HELENA"
	case 145:
		return "SAINT KITTS AND NEVIS"
	case 146:
		return "SAINT LUCIA"
	case 147:
		return "SAINT PIERRE AND MIQUELON"
	case 148:
		return "SAINT VINCENT AND THE GRENADINES"
	case 149:
		return "SAMOA"
	case 150:
		return "SAN MARINO"
	case 151:
		return "SAO TOME AND PRINCIPE"
	case 152:
		return "SAUDI ARABIA"
	case 153:
		return "SENEGAL"
	case 154:
		return "SERBIA"
	case 155:
		return "SEYCHELLES"
	case 156:
		return "SIERRA LEONE"
	case 157:
		return "SINGAPORE"
	case 158:
		return "SLOVAKIA"
	case 159:
		return "SLOVENIA"
	case 160:
		return "SOLOMON ISLANDS"
	case 161:
		return "SOMALIA"
	case 162:
		return "SOUTH AFRICA"
	case 163:
		return "SPAIN"
	case 164:
		return "SRI LANKA"
	case 165:
		return "SURINAME"
	case 166:
		return "SVALBARD AND JAN MAYEN"
	case 167:
		return "SWAZILAND"
	case 168:
		return "SWEDEN"
	case 169:
		return "SWITZERLAND"
	case 170:
		return "TAIWAN, PROVINCE OF CHINA"
	case 171:
		return "TAJIKISTAN"
	case 172:
		return "TANZANIA, UNITED REPUBLIC OF"
	case 173:
		return "THAILAND"
	case 174:
		return "TOGO"
	case 175:
		return "TONGA"
	case 176:
		return "TRINIDAD AND TOBAGO"
	case 177:
		return "TUNISIA"
	case 178:
		return "TURKEY"
	case 179:
		return "TURKMENISTAN"
	case 180:
		return "TURKS AND CAICOS ISLANDS"
	case 181:
		return "TUVALU"
	case 182:
		return "UGANDA"
	case 183:
		return "UKRAINE"
	case 184:
		return "UNITED ARAB EMIRATES"
	case 185:
		return "UNITED KINGDOM"
	case 186:
		return "UNITED STATES"
	case 187:
		return "URUGUAY"
	case 188:
		return "VANUATU"
	case 189:
		return "VENEZUELA"
	case 190:
		return "VIETNAM"
	case 191:
		return "VIRGIN ISLANDS, BRITISH"
	case 192:
		return "WALLIS AND FUTUNA"
	case 193:
		return "YEMEN"
	case 194:
		return "ZAMBIA"
	}
	return ""
}

func (self CountryCodeEnum) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(self.IntValue())), nil
}

func (self *CountryCodeEnum) UnmarshalJSON(b []byte) error {
	var n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	self.value_30bx5a8om6xs = uint8(n)
	return nil
}
