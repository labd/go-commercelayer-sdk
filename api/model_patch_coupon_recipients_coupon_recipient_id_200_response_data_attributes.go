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

// checks if the PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes{}

// PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes struct for PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes
type PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes struct {
	// The recipient email address.
	Email interface{} `json:"email,omitempty"`
	// The recipient first name.
	FirstName interface{} `json:"first_name,omitempty"`
	// The recipient last name.
	LastName interface{} `json:"last_name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes instantiates a new PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes() *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	this := PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes{}
	return &this
}

// NewPATCHCouponRecipientsCouponRecipientId200ResponseDataAttributesWithDefaults instantiates a new PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCouponRecipientsCouponRecipientId200ResponseDataAttributesWithDefaults() *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	this := PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return &o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasEmail() bool {
	if o != nil && IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given interface{} and assigns it to the Email field.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetEmail(v interface{}) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetFirstName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetFirstNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return &o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasFirstName() bool {
	if o != nil && IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given interface{} and assigns it to the FirstName field.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetFirstName(v interface{}) {
	o.FirstName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetLastName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetLastNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return &o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasLastName() bool {
	if o != nil && IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given interface{} and assigns it to the LastName field.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetLastName(v interface{}) {
	o.LastName = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes struct {
	value *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) Get() *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) Set(val *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes(val *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) *NullablePATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	return &NullablePATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
