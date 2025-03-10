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

// checks if the GETVoidsVoidId200ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETVoidsVoidId200ResponseDataRelationships{}

// GETVoidsVoidId200ResponseDataRelationships struct for GETVoidsVoidId200ResponseDataRelationships
type GETVoidsVoidId200ResponseDataRelationships struct {
	Order                  *POSTAdyenPayments201ResponseDataRelationshipsOrder                        `json:"order,omitempty"`
	PaymentSource          *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource `json:"payment_source,omitempty"`
	Attachments            *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments   `json:"attachments,omitempty"`
	Events                 *POSTAddresses201ResponseDataRelationshipsEvents                           `json:"events,omitempty"`
	Versions               *POSTAddresses201ResponseDataRelationshipsVersions                         `json:"versions,omitempty"`
	ReferenceAuthorization *GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization    `json:"reference_authorization,omitempty"`
}

// NewGETVoidsVoidId200ResponseDataRelationships instantiates a new GETVoidsVoidId200ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETVoidsVoidId200ResponseDataRelationships() *GETVoidsVoidId200ResponseDataRelationships {
	this := GETVoidsVoidId200ResponseDataRelationships{}
	return &this
}

// NewGETVoidsVoidId200ResponseDataRelationshipsWithDefaults instantiates a new GETVoidsVoidId200ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETVoidsVoidId200ResponseDataRelationshipsWithDefaults() *GETVoidsVoidId200ResponseDataRelationships {
	this := GETVoidsVoidId200ResponseDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret POSTAdyenPayments201ResponseDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given POSTAdyenPayments201ResponseDataRelationshipsOrder and assigns it to the Order field.
func (o *GETVoidsVoidId200ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder) {
	o.Order = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetPaymentSource() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource {
	if o == nil || IsNil(o.PaymentSource) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetPaymentSourceOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource, bool) {
	if o == nil || IsNil(o.PaymentSource) {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) HasPaymentSource() bool {
	if o != nil && !IsNil(o.PaymentSource) {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *GETVoidsVoidId200ResponseDataRelationships) SetPaymentSource(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETVoidsVoidId200ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *GETVoidsVoidId200ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *GETVoidsVoidId200ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetReferenceAuthorization returns the ReferenceAuthorization field value if set, zero value otherwise.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetReferenceAuthorization() GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization {
	if o == nil || IsNil(o.ReferenceAuthorization) {
		var ret GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization
		return ret
	}
	return *o.ReferenceAuthorization
}

// GetReferenceAuthorizationOk returns a tuple with the ReferenceAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) GetReferenceAuthorizationOk() (*GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization, bool) {
	if o == nil || IsNil(o.ReferenceAuthorization) {
		return nil, false
	}
	return o.ReferenceAuthorization, true
}

// HasReferenceAuthorization returns a boolean if a field has been set.
func (o *GETVoidsVoidId200ResponseDataRelationships) HasReferenceAuthorization() bool {
	if o != nil && !IsNil(o.ReferenceAuthorization) {
		return true
	}

	return false
}

// SetReferenceAuthorization gets a reference to the given GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization and assigns it to the ReferenceAuthorization field.
func (o *GETVoidsVoidId200ResponseDataRelationships) SetReferenceAuthorization(v GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization) {
	o.ReferenceAuthorization = &v
}

func (o GETVoidsVoidId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETVoidsVoidId200ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
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

type NullableGETVoidsVoidId200ResponseDataRelationships struct {
	value *GETVoidsVoidId200ResponseDataRelationships
	isSet bool
}

func (v NullableGETVoidsVoidId200ResponseDataRelationships) Get() *GETVoidsVoidId200ResponseDataRelationships {
	return v.value
}

func (v *NullableGETVoidsVoidId200ResponseDataRelationships) Set(val *GETVoidsVoidId200ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETVoidsVoidId200ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETVoidsVoidId200ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETVoidsVoidId200ResponseDataRelationships(val *GETVoidsVoidId200ResponseDataRelationships) *NullableGETVoidsVoidId200ResponseDataRelationships {
	return &NullableGETVoidsVoidId200ResponseDataRelationships{value: val, isSet: true}
}

func (v NullableGETVoidsVoidId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETVoidsVoidId200ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
