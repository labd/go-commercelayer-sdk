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

// checks if the CustomerCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerCreateData{}

// CustomerCreateData struct for CustomerCreateData
type CustomerCreateData struct {
	// The resource's type
	Type          interface{}                            `json:"type"`
	Attributes    POSTCustomers201ResponseDataAttributes `json:"attributes"`
	Relationships *CustomerCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewCustomerCreateData instantiates a new CustomerCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCreateData(type_ interface{}, attributes POSTCustomers201ResponseDataAttributes) *CustomerCreateData {
	this := CustomerCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerCreateDataWithDefaults instantiates a new CustomerCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCreateDataWithDefaults() *CustomerCreateData {
	this := CustomerCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CustomerCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerCreateData) GetAttributes() POSTCustomers201ResponseDataAttributes {
	if o == nil {
		var ret POSTCustomers201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerCreateData) GetAttributesOk() (*POSTCustomers201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerCreateData) SetAttributes(v POSTCustomers201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerCreateData) GetRelationships() CustomerCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CustomerCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateData) GetRelationshipsOk() (*CustomerCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerCreateDataRelationships and assigns it to the Relationships field.
func (o *CustomerCreateData) SetRelationships(v CustomerCreateDataRelationships) {
	o.Relationships = &v
}

func (o CustomerCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableCustomerCreateData struct {
	value *CustomerCreateData
	isSet bool
}

func (v NullableCustomerCreateData) Get() *CustomerCreateData {
	return v.value
}

func (v *NullableCustomerCreateData) Set(val *CustomerCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCreateData(val *CustomerCreateData) *NullableCustomerCreateData {
	return &NullableCustomerCreateData{value: val, isSet: true}
}

func (v NullableCustomerCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
