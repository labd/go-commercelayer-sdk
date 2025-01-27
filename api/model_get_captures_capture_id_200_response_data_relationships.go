/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETCapturesCaptureId200ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCapturesCaptureId200ResponseDataRelationships{}

// GETCapturesCaptureId200ResponseDataRelationships struct for GETCapturesCaptureId200ResponseDataRelationships
type GETCapturesCaptureId200ResponseDataRelationships struct {
	Order                  *POSTAdyenPayments201ResponseDataRelationshipsOrder                        `json:"order,omitempty"`
	PaymentSource          *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource `json:"payment_source,omitempty"`
	Attachments            *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments   `json:"attachments,omitempty"`
	Events                 *POSTAddresses201ResponseDataRelationshipsEvents                           `json:"events,omitempty"`
	Versions               *POSTAddresses201ResponseDataRelationshipsVersions                         `json:"versions,omitempty"`
	ReferenceAuthorization *GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization    `json:"reference_authorization,omitempty"`
	Refunds                *GETCapturesCaptureId200ResponseDataRelationshipsRefunds                   `json:"refunds,omitempty"`
	Return                 *GETCapturesCaptureId200ResponseDataRelationshipsReturn                    `json:"return,omitempty"`
}

// NewGETCapturesCaptureId200ResponseDataRelationships instantiates a new GETCapturesCaptureId200ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCapturesCaptureId200ResponseDataRelationships() *GETCapturesCaptureId200ResponseDataRelationships {
	this := GETCapturesCaptureId200ResponseDataRelationships{}
	return &this
}

// NewGETCapturesCaptureId200ResponseDataRelationshipsWithDefaults instantiates a new GETCapturesCaptureId200ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCapturesCaptureId200ResponseDataRelationshipsWithDefaults() *GETCapturesCaptureId200ResponseDataRelationships {
	this := GETCapturesCaptureId200ResponseDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret POSTAdyenPayments201ResponseDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given POSTAdyenPayments201ResponseDataRelationshipsOrder and assigns it to the Order field.
func (o *GETCapturesCaptureId200ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder) {
	o.Order = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetPaymentSource() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource {
	if o == nil || IsNil(o.PaymentSource) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetPaymentSourceOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource, bool) {
	if o == nil || IsNil(o.PaymentSource) {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) HasPaymentSource() bool {
	if o != nil && !IsNil(o.PaymentSource) {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *GETCapturesCaptureId200ResponseDataRelationships) SetPaymentSource(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETCapturesCaptureId200ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *GETCapturesCaptureId200ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *GETCapturesCaptureId200ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetReferenceAuthorization returns the ReferenceAuthorization field value if set, zero value otherwise.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetReferenceAuthorization() GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization {
	if o == nil || IsNil(o.ReferenceAuthorization) {
		var ret GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization
		return ret
	}
	return *o.ReferenceAuthorization
}

// GetReferenceAuthorizationOk returns a tuple with the ReferenceAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetReferenceAuthorizationOk() (*GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization, bool) {
	if o == nil || IsNil(o.ReferenceAuthorization) {
		return nil, false
	}
	return o.ReferenceAuthorization, true
}

// HasReferenceAuthorization returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) HasReferenceAuthorization() bool {
	if o != nil && !IsNil(o.ReferenceAuthorization) {
		return true
	}

	return false
}

// SetReferenceAuthorization gets a reference to the given GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization and assigns it to the ReferenceAuthorization field.
func (o *GETCapturesCaptureId200ResponseDataRelationships) SetReferenceAuthorization(v GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization) {
	o.ReferenceAuthorization = &v
}

// GetRefunds returns the Refunds field value if set, zero value otherwise.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetRefunds() GETCapturesCaptureId200ResponseDataRelationshipsRefunds {
	if o == nil || IsNil(o.Refunds) {
		var ret GETCapturesCaptureId200ResponseDataRelationshipsRefunds
		return ret
	}
	return *o.Refunds
}

// GetRefundsOk returns a tuple with the Refunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetRefundsOk() (*GETCapturesCaptureId200ResponseDataRelationshipsRefunds, bool) {
	if o == nil || IsNil(o.Refunds) {
		return nil, false
	}
	return o.Refunds, true
}

// HasRefunds returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) HasRefunds() bool {
	if o != nil && !IsNil(o.Refunds) {
		return true
	}

	return false
}

// SetRefunds gets a reference to the given GETCapturesCaptureId200ResponseDataRelationshipsRefunds and assigns it to the Refunds field.
func (o *GETCapturesCaptureId200ResponseDataRelationships) SetRefunds(v GETCapturesCaptureId200ResponseDataRelationshipsRefunds) {
	o.Refunds = &v
}

// GetReturn returns the Return field value if set, zero value otherwise.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetReturn() GETCapturesCaptureId200ResponseDataRelationshipsReturn {
	if o == nil || IsNil(o.Return) {
		var ret GETCapturesCaptureId200ResponseDataRelationshipsReturn
		return ret
	}
	return *o.Return
}

// GetReturnOk returns a tuple with the Return field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) GetReturnOk() (*GETCapturesCaptureId200ResponseDataRelationshipsReturn, bool) {
	if o == nil || IsNil(o.Return) {
		return nil, false
	}
	return o.Return, true
}

// HasReturn returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationships) HasReturn() bool {
	if o != nil && !IsNil(o.Return) {
		return true
	}

	return false
}

// SetReturn gets a reference to the given GETCapturesCaptureId200ResponseDataRelationshipsReturn and assigns it to the Return field.
func (o *GETCapturesCaptureId200ResponseDataRelationships) SetReturn(v GETCapturesCaptureId200ResponseDataRelationshipsReturn) {
	o.Return = &v
}

func (o GETCapturesCaptureId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCapturesCaptureId200ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
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

type NullableGETCapturesCaptureId200ResponseDataRelationships struct {
	value *GETCapturesCaptureId200ResponseDataRelationships
	isSet bool
}

func (v NullableGETCapturesCaptureId200ResponseDataRelationships) Get() *GETCapturesCaptureId200ResponseDataRelationships {
	return v.value
}

func (v *NullableGETCapturesCaptureId200ResponseDataRelationships) Set(val *GETCapturesCaptureId200ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCapturesCaptureId200ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCapturesCaptureId200ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCapturesCaptureId200ResponseDataRelationships(val *GETCapturesCaptureId200ResponseDataRelationships) *NullableGETCapturesCaptureId200ResponseDataRelationships {
	return &NullableGETCapturesCaptureId200ResponseDataRelationships{value: val, isSet: true}
}

func (v NullableGETCapturesCaptureId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCapturesCaptureId200ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
