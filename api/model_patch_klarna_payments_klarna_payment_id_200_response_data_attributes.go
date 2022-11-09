/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes struct for PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes
type PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes struct {
	// The token returned by a successful client authorization, mandatory to place the order.
	AuthToken *string `json:"auth_token,omitempty"`
	// Send this attribute if you want to update the payment session with fresh order data.
	Update *bool `json:"_update,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes instantiates a new PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes() *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes {
	this := PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes{}
	return &this
}

// NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributesWithDefaults() *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes {
	this := PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes{}
	return &this
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetAuthToken() string {
	if o == nil || o.AuthToken == nil {
		var ret string
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetAuthTokenOk() (*string, bool) {
	if o == nil || o.AuthToken == nil {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasAuthToken() bool {
	if o != nil && o.AuthToken != nil {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given string and assigns it to the AuthToken field.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetAuthToken(v string) {
	o.AuthToken = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetUpdate() bool {
	if o == nil || o.Update == nil {
		var ret bool
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetUpdateOk() (*bool, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given bool and assigns it to the Update field.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetUpdate(v bool) {
	o.Update = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthToken != nil {
		toSerialize["auth_token"] = o.AuthToken
	}
	if o.Update != nil {
		toSerialize["_update"] = o.Update
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
	return json.Marshal(toSerialize)
}

type NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes struct {
	value *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) Get() *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) Set(val *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes(val *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) *NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes {
	return &NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
