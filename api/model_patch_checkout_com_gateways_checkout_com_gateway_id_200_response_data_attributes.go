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

// PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes struct for PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes
type PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name *string `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The gateway secret key.
	SecretKey *string `json:"secret_key,omitempty"`
	// The gateway public key.
	PublicKey *string `json:"public_key,omitempty"`
}

// NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes instantiates a new PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes() *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	this := PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes{}
	return &this
}

// NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributesWithDefaults() *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	this := PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
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
	if o.SecretKey != nil {
		toSerialize["secret_key"] = o.SecretKey
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes struct {
	value *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) Get() *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) Set(val *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes(val *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	return &NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
