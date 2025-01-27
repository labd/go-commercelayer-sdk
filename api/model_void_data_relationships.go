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

// checks if the VoidDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoidDataRelationships{}

// VoidDataRelationships struct for VoidDataRelationships
type VoidDataRelationships struct {
	Order                  *AdyenPaymentDataRelationshipsOrder             `json:"order,omitempty"`
	PaymentSource          *AuthorizationDataRelationshipsPaymentSource    `json:"payment_source,omitempty"`
	Attachments            *AuthorizationDataRelationshipsAttachments      `json:"attachments,omitempty"`
	Events                 *AddressDataRelationshipsEvents                 `json:"events,omitempty"`
	Versions               *AddressDataRelationshipsVersions               `json:"versions,omitempty"`
	ReferenceAuthorization *CaptureDataRelationshipsReferenceAuthorization `json:"reference_authorization,omitempty"`
}

// NewVoidDataRelationships instantiates a new VoidDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoidDataRelationships() *VoidDataRelationships {
	this := VoidDataRelationships{}
	return &this
}

// NewVoidDataRelationshipsWithDefaults instantiates a new VoidDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoidDataRelationshipsWithDefaults() *VoidDataRelationships {
	this := VoidDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VoidDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VoidDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *VoidDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *VoidDataRelationships) GetPaymentSource() AuthorizationDataRelationshipsPaymentSource {
	if o == nil || IsNil(o.PaymentSource) {
		var ret AuthorizationDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidDataRelationships) GetPaymentSourceOk() (*AuthorizationDataRelationshipsPaymentSource, bool) {
	if o == nil || IsNil(o.PaymentSource) {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *VoidDataRelationships) HasPaymentSource() bool {
	if o != nil && !IsNil(o.PaymentSource) {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given AuthorizationDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *VoidDataRelationships) SetPaymentSource(v AuthorizationDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *VoidDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *VoidDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *VoidDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *VoidDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *VoidDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *VoidDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *VoidDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *VoidDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *VoidDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

// GetReferenceAuthorization returns the ReferenceAuthorization field value if set, zero value otherwise.
func (o *VoidDataRelationships) GetReferenceAuthorization() CaptureDataRelationshipsReferenceAuthorization {
	if o == nil || IsNil(o.ReferenceAuthorization) {
		var ret CaptureDataRelationshipsReferenceAuthorization
		return ret
	}
	return *o.ReferenceAuthorization
}

// GetReferenceAuthorizationOk returns a tuple with the ReferenceAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidDataRelationships) GetReferenceAuthorizationOk() (*CaptureDataRelationshipsReferenceAuthorization, bool) {
	if o == nil || IsNil(o.ReferenceAuthorization) {
		return nil, false
	}
	return o.ReferenceAuthorization, true
}

// HasReferenceAuthorization returns a boolean if a field has been set.
func (o *VoidDataRelationships) HasReferenceAuthorization() bool {
	if o != nil && !IsNil(o.ReferenceAuthorization) {
		return true
	}

	return false
}

// SetReferenceAuthorization gets a reference to the given CaptureDataRelationshipsReferenceAuthorization and assigns it to the ReferenceAuthorization field.
func (o *VoidDataRelationships) SetReferenceAuthorization(v CaptureDataRelationshipsReferenceAuthorization) {
	o.ReferenceAuthorization = &v
}

func (o VoidDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoidDataRelationships) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableVoidDataRelationships struct {
	value *VoidDataRelationships
	isSet bool
}

func (v NullableVoidDataRelationships) Get() *VoidDataRelationships {
	return v.value
}

func (v *NullableVoidDataRelationships) Set(val *VoidDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableVoidDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableVoidDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoidDataRelationships(val *VoidDataRelationships) *NullableVoidDataRelationships {
	return &NullableVoidDataRelationships{value: val, isSet: true}
}

func (v NullableVoidDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoidDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
