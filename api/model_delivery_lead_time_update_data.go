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

// DeliveryLeadTimeUpdateData struct for DeliveryLeadTimeUpdateData
type DeliveryLeadTimeUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                            `json:"id"`
	Attributes    PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes `json:"attributes"`
	Relationships *DeliveryLeadTimeUpdateDataRelationships                          `json:"relationships,omitempty"`
}

// NewDeliveryLeadTimeUpdateData instantiates a new DeliveryLeadTimeUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeUpdateData(type_ string, id string, attributes PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) *DeliveryLeadTimeUpdateData {
	this := DeliveryLeadTimeUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewDeliveryLeadTimeUpdateDataWithDefaults instantiates a new DeliveryLeadTimeUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeUpdateDataWithDefaults() *DeliveryLeadTimeUpdateData {
	this := DeliveryLeadTimeUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *DeliveryLeadTimeUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeliveryLeadTimeUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *DeliveryLeadTimeUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeliveryLeadTimeUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *DeliveryLeadTimeUpdateData) GetAttributes() PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeUpdateData) GetAttributesOk() (*PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *DeliveryLeadTimeUpdateData) SetAttributes(v PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *DeliveryLeadTimeUpdateData) GetRelationships() DeliveryLeadTimeUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret DeliveryLeadTimeUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeUpdateData) GetRelationshipsOk() (*DeliveryLeadTimeUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *DeliveryLeadTimeUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given DeliveryLeadTimeUpdateDataRelationships and assigns it to the Relationships field.
func (o *DeliveryLeadTimeUpdateData) SetRelationships(v DeliveryLeadTimeUpdateDataRelationships) {
	o.Relationships = &v
}

func (o DeliveryLeadTimeUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTimeUpdateData struct {
	value *DeliveryLeadTimeUpdateData
	isSet bool
}

func (v NullableDeliveryLeadTimeUpdateData) Get() *DeliveryLeadTimeUpdateData {
	return v.value
}

func (v *NullableDeliveryLeadTimeUpdateData) Set(val *DeliveryLeadTimeUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeUpdateData(val *DeliveryLeadTimeUpdateData) *NullableDeliveryLeadTimeUpdateData {
	return &NullableDeliveryLeadTimeUpdateData{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
