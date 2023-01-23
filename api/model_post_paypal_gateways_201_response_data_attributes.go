/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTPaypalGateways201ResponseDataAttributes struct for POSTPaypalGateways201ResponseDataAttributes
type POSTPaypalGateways201ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name string `json:"name"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The gateway client ID.
	ClientId string `json:"client_id"`
	// The gateway client secret.
	ClientSecret string `json:"client_secret"`
}

// NewPOSTPaypalGateways201ResponseDataAttributes instantiates a new POSTPaypalGateways201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPaypalGateways201ResponseDataAttributes(name string, clientId string, clientSecret string) *POSTPaypalGateways201ResponseDataAttributes {
	this := POSTPaypalGateways201ResponseDataAttributes{}
	this.Name = name
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewPOSTPaypalGateways201ResponseDataAttributesWithDefaults instantiates a new POSTPaypalGateways201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPaypalGateways201ResponseDataAttributesWithDefaults() *POSTPaypalGateways201ResponseDataAttributes {
	this := POSTPaypalGateways201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *POSTPaypalGateways201ResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTPaypalGateways201ResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTPaypalGateways201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTPaypalGateways201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTPaypalGateways201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTPaypalGateways201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTPaypalGateways201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTPaypalGateways201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTPaypalGateways201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTPaypalGateways201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTPaypalGateways201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetClientId returns the ClientId field value
func (o *POSTPaypalGateways201ResponseDataAttributes) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201ResponseDataAttributes) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *POSTPaypalGateways201ResponseDataAttributes) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *POSTPaypalGateways201ResponseDataAttributes) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201ResponseDataAttributes) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *POSTPaypalGateways201ResponseDataAttributes) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o POSTPaypalGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
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
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["client_secret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTPaypalGateways201ResponseDataAttributes struct {
	value *POSTPaypalGateways201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTPaypalGateways201ResponseDataAttributes) Get() *POSTPaypalGateways201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTPaypalGateways201ResponseDataAttributes) Set(val *POSTPaypalGateways201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPaypalGateways201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPaypalGateways201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPaypalGateways201ResponseDataAttributes(val *POSTPaypalGateways201ResponseDataAttributes) *NullablePOSTPaypalGateways201ResponseDataAttributes {
	return &NullablePOSTPaypalGateways201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTPaypalGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPaypalGateways201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
