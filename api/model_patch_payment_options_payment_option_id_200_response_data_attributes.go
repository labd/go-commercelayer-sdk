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

// checks if the PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes{}

// PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes struct for PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes
type PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes struct {
	// The payment option's name. Wehn blank is inherited by payment source type.
	Name interface{} `json:"name,omitempty"`
	// The payment options data to be added to the payment source payload. Check payment specific API for more details.
	Data interface{} `json:"data,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes instantiates a new PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes() *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes {
	this := PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes{}
	return &this
}

// NewPATCHPaymentOptionsPaymentOptionId200ResponseDataAttributesWithDefaults instantiates a new PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPaymentOptionsPaymentOptionId200ResponseDataAttributesWithDefaults() *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes {
	this := PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetData(v interface{}) {
	o.Data = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
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

type NullablePATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes struct {
	value *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) Get() *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) Set(val *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes(val *PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) *NullablePATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes {
	return &NullablePATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
