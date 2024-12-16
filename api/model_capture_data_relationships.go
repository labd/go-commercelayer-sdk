/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CaptureDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptureDataRelationships{}

// CaptureDataRelationships struct for CaptureDataRelationships
type CaptureDataRelationships struct {
	Order                  *AdyenPaymentDataRelationshipsOrder             `json:"order,omitempty"`
	PaymentSource          *AuthorizationDataRelationshipsPaymentSource    `json:"payment_source,omitempty"`
	Attachments            *AuthorizationDataRelationshipsAttachments      `json:"attachments,omitempty"`
	Events                 *AddressDataRelationshipsEvents                 `json:"events,omitempty"`
	Versions               *AddressDataRelationshipsVersions               `json:"versions,omitempty"`
	ReferenceAuthorization *CaptureDataRelationshipsReferenceAuthorization `json:"reference_authorization,omitempty"`
	Refunds                *CaptureDataRelationshipsRefunds                `json:"refunds,omitempty"`
	Return                 *CaptureDataRelationshipsReturn                 `json:"return,omitempty"`
}

// NewCaptureDataRelationships instantiates a new CaptureDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureDataRelationships() *CaptureDataRelationships {
	this := CaptureDataRelationships{}
	return &this
}

// NewCaptureDataRelationshipsWithDefaults instantiates a new CaptureDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureDataRelationshipsWithDefaults() *CaptureDataRelationships {
	this := CaptureDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *CaptureDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *CaptureDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *CaptureDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *CaptureDataRelationships) GetPaymentSource() AuthorizationDataRelationshipsPaymentSource {
	if o == nil || IsNil(o.PaymentSource) {
		var ret AuthorizationDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureDataRelationships) GetPaymentSourceOk() (*AuthorizationDataRelationshipsPaymentSource, bool) {
	if o == nil || IsNil(o.PaymentSource) {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *CaptureDataRelationships) HasPaymentSource() bool {
	if o != nil && !IsNil(o.PaymentSource) {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given AuthorizationDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *CaptureDataRelationships) SetPaymentSource(v AuthorizationDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *CaptureDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CaptureDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *CaptureDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CaptureDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *CaptureDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *CaptureDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *CaptureDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *CaptureDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *CaptureDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

// GetReferenceAuthorization returns the ReferenceAuthorization field value if set, zero value otherwise.
func (o *CaptureDataRelationships) GetReferenceAuthorization() CaptureDataRelationshipsReferenceAuthorization {
	if o == nil || IsNil(o.ReferenceAuthorization) {
		var ret CaptureDataRelationshipsReferenceAuthorization
		return ret
	}
	return *o.ReferenceAuthorization
}

// GetReferenceAuthorizationOk returns a tuple with the ReferenceAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureDataRelationships) GetReferenceAuthorizationOk() (*CaptureDataRelationshipsReferenceAuthorization, bool) {
	if o == nil || IsNil(o.ReferenceAuthorization) {
		return nil, false
	}
	return o.ReferenceAuthorization, true
}

// HasReferenceAuthorization returns a boolean if a field has been set.
func (o *CaptureDataRelationships) HasReferenceAuthorization() bool {
	if o != nil && !IsNil(o.ReferenceAuthorization) {
		return true
	}

	return false
}

// SetReferenceAuthorization gets a reference to the given CaptureDataRelationshipsReferenceAuthorization and assigns it to the ReferenceAuthorization field.
func (o *CaptureDataRelationships) SetReferenceAuthorization(v CaptureDataRelationshipsReferenceAuthorization) {
	o.ReferenceAuthorization = &v
}

// GetRefunds returns the Refunds field value if set, zero value otherwise.
func (o *CaptureDataRelationships) GetRefunds() CaptureDataRelationshipsRefunds {
	if o == nil || IsNil(o.Refunds) {
		var ret CaptureDataRelationshipsRefunds
		return ret
	}
	return *o.Refunds
}

// GetRefundsOk returns a tuple with the Refunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureDataRelationships) GetRefundsOk() (*CaptureDataRelationshipsRefunds, bool) {
	if o == nil || IsNil(o.Refunds) {
		return nil, false
	}
	return o.Refunds, true
}

// HasRefunds returns a boolean if a field has been set.
func (o *CaptureDataRelationships) HasRefunds() bool {
	if o != nil && !IsNil(o.Refunds) {
		return true
	}

	return false
}

// SetRefunds gets a reference to the given CaptureDataRelationshipsRefunds and assigns it to the Refunds field.
func (o *CaptureDataRelationships) SetRefunds(v CaptureDataRelationshipsRefunds) {
	o.Refunds = &v
}

// GetReturn returns the Return field value if set, zero value otherwise.
func (o *CaptureDataRelationships) GetReturn() CaptureDataRelationshipsReturn {
	if o == nil || IsNil(o.Return) {
		var ret CaptureDataRelationshipsReturn
		return ret
	}
	return *o.Return
}

// GetReturnOk returns a tuple with the Return field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureDataRelationships) GetReturnOk() (*CaptureDataRelationshipsReturn, bool) {
	if o == nil || IsNil(o.Return) {
		return nil, false
	}
	return o.Return, true
}

// HasReturn returns a boolean if a field has been set.
func (o *CaptureDataRelationships) HasReturn() bool {
	if o != nil && !IsNil(o.Return) {
		return true
	}

	return false
}

// SetReturn gets a reference to the given CaptureDataRelationshipsReturn and assigns it to the Return field.
func (o *CaptureDataRelationships) SetReturn(v CaptureDataRelationshipsReturn) {
	o.Return = &v
}

func (o CaptureDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.PaymentSource) {
		toSerialize["payment_source"] = o.PaymentSource
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.ReferenceAuthorization) {
		toSerialize["reference_authorization"] = o.ReferenceAuthorization
	}
	if !IsNil(o.Refunds) {
		toSerialize["refunds"] = o.Refunds
	}
	if !IsNil(o.Return) {
		toSerialize["return"] = o.Return
	}
	return toSerialize, nil
}

type NullableCaptureDataRelationships struct {
	value *CaptureDataRelationships
	isSet bool
}

func (v NullableCaptureDataRelationships) Get() *CaptureDataRelationships {
	return v.value
}

func (v *NullableCaptureDataRelationships) Set(val *CaptureDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureDataRelationships(val *CaptureDataRelationships) *NullableCaptureDataRelationships {
	return &NullableCaptureDataRelationships{value: val, isSet: true}
}

func (v NullableCaptureDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
