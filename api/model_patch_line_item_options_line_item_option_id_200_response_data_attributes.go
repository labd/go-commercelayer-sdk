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

// PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes struct for PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes
type PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes struct {
	// The name of the line item option. When blank, it gets populated with the name of the associated SKU option.
	Name *string `json:"name,omitempty"`
	// The line item option's quantity
	Quantity *int32 `json:"quantity,omitempty"`
	// Set of key-value pairs that represent the selected options.
	Options map[string]interface{} `json:"options,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes instantiates a new PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes() *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes {
	this := PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes{}
	return &this
}

// NewPATCHLineItemOptionsLineItemOptionId200ResponseDataAttributesWithDefaults instantiates a new PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHLineItemOptionsLineItemOptionId200ResponseDataAttributesWithDefaults() *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes {
	this := PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
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

type NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes struct {
	value *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) Get() *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) Set(val *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes(val *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes {
	return &NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
