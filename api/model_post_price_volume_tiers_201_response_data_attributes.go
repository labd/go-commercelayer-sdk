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

// checks if the POSTPriceVolumeTiers201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPriceVolumeTiers201ResponseDataAttributes{}

// POSTPriceVolumeTiers201ResponseDataAttributes struct for POSTPriceVolumeTiers201ResponseDataAttributes
type POSTPriceVolumeTiers201ResponseDataAttributes struct {
	// The price tier's name.
	Name interface{} `json:"name"`
	// The tier upper limit, expressed as the line item quantity. When 'null' it means infinity (useful to have an always matching tier).
	UpTo interface{} `json:"up_to,omitempty"`
	// The price of this price tier, in cents.
	PriceAmountCents interface{} `json:"price_amount_cents"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTPriceVolumeTiers201ResponseDataAttributes instantiates a new POSTPriceVolumeTiers201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPriceVolumeTiers201ResponseDataAttributes(name interface{}, priceAmountCents interface{}) *POSTPriceVolumeTiers201ResponseDataAttributes {
	this := POSTPriceVolumeTiers201ResponseDataAttributes{}
	this.Name = name
	this.PriceAmountCents = priceAmountCents
	return &this
}

// NewPOSTPriceVolumeTiers201ResponseDataAttributesWithDefaults instantiates a new POSTPriceVolumeTiers201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPriceVolumeTiers201ResponseDataAttributesWithDefaults() *POSTPriceVolumeTiers201ResponseDataAttributes {
	this := POSTPriceVolumeTiers201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetUpTo returns the UpTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetUpTo() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpTo
}

// GetUpToOk returns a tuple with the UpTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetUpToOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpTo) {
		return nil, false
	}
	return &o.UpTo, true
}

// HasUpTo returns a boolean if a field has been set.
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) HasUpTo() bool {
	if o != nil && IsNil(o.UpTo) {
		return true
	}

	return false
}

// SetUpTo gets a reference to the given interface{} and assigns it to the UpTo field.
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) SetUpTo(v interface{}) {
	o.UpTo = v
}

// GetPriceAmountCents returns the PriceAmountCents field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetPriceAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PriceAmountCents) {
		return nil, false
	}
	return &o.PriceAmountCents, true
}

// SetPriceAmountCents sets field value
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) SetPriceAmountCents(v interface{}) {
	o.PriceAmountCents = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTPriceVolumeTiers201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTPriceVolumeTiers201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPriceVolumeTiers201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UpTo != nil {
		toSerialize["up_to"] = o.UpTo
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
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

type NullablePOSTPriceVolumeTiers201ResponseDataAttributes struct {
	value *POSTPriceVolumeTiers201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTPriceVolumeTiers201ResponseDataAttributes) Get() *POSTPriceVolumeTiers201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTPriceVolumeTiers201ResponseDataAttributes) Set(val *POSTPriceVolumeTiers201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPriceVolumeTiers201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPriceVolumeTiers201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPriceVolumeTiers201ResponseDataAttributes(val *POSTPriceVolumeTiers201ResponseDataAttributes) *NullablePOSTPriceVolumeTiers201ResponseDataAttributes {
	return &NullablePOSTPriceVolumeTiers201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTPriceVolumeTiers201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPriceVolumeTiers201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
