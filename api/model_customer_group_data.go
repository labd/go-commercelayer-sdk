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

// checks if the CustomerGroupData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerGroupData{}

// CustomerGroupData struct for CustomerGroupData
type CustomerGroupData struct {
	// The resource's type
	Type          interface{}                                               `json:"type"`
	Attributes    GETCustomerGroupsCustomerGroupId200ResponseDataAttributes `json:"attributes"`
	Relationships *CustomerGroupDataRelationships                           `json:"relationships,omitempty"`
}

// NewCustomerGroupData instantiates a new CustomerGroupData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerGroupData(type_ interface{}, attributes GETCustomerGroupsCustomerGroupId200ResponseDataAttributes) *CustomerGroupData {
	this := CustomerGroupData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerGroupDataWithDefaults instantiates a new CustomerGroupData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerGroupDataWithDefaults() *CustomerGroupData {
	this := CustomerGroupData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CustomerGroupData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerGroupData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerGroupData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerGroupData) GetAttributes() GETCustomerGroupsCustomerGroupId200ResponseDataAttributes {
	if o == nil {
		var ret GETCustomerGroupsCustomerGroupId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerGroupData) GetAttributesOk() (*GETCustomerGroupsCustomerGroupId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerGroupData) SetAttributes(v GETCustomerGroupsCustomerGroupId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerGroupData) GetRelationships() CustomerGroupDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CustomerGroupDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupData) GetRelationshipsOk() (*CustomerGroupDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerGroupData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerGroupDataRelationships and assigns it to the Relationships field.
func (o *CustomerGroupData) SetRelationships(v CustomerGroupDataRelationships) {
	o.Relationships = &v
}

func (o CustomerGroupData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerGroupData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableCustomerGroupData struct {
	value *CustomerGroupData
	isSet bool
}

func (v NullableCustomerGroupData) Get() *CustomerGroupData {
	return v.value
}

func (v *NullableCustomerGroupData) Set(val *CustomerGroupData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerGroupData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerGroupData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerGroupData(val *CustomerGroupData) *NullableCustomerGroupData {
	return &NullableCustomerGroupData{value: val, isSet: true}
}

func (v NullableCustomerGroupData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerGroupData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
