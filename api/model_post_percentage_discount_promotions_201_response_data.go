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

// checks if the POSTPercentageDiscountPromotions201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPercentageDiscountPromotions201ResponseData{}

// POSTPercentageDiscountPromotions201ResponseData struct for POSTPercentageDiscountPromotions201ResponseData
type POSTPercentageDiscountPromotions201ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                                `json:"type,omitempty"`
	Links         *POSTAddresses201ResponseDataLinks                         `json:"links,omitempty"`
	Attributes    *POSTPercentageDiscountPromotions201ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *POSTBuyXPayYPromotions201ResponseDataRelationships        `json:"relationships,omitempty"`
}

// NewPOSTPercentageDiscountPromotions201ResponseData instantiates a new POSTPercentageDiscountPromotions201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPercentageDiscountPromotions201ResponseData() *POSTPercentageDiscountPromotions201ResponseData {
	this := POSTPercentageDiscountPromotions201ResponseData{}
	return &this
}

// NewPOSTPercentageDiscountPromotions201ResponseDataWithDefaults instantiates a new POSTPercentageDiscountPromotions201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPercentageDiscountPromotions201ResponseDataWithDefaults() *POSTPercentageDiscountPromotions201ResponseData {
	this := POSTPercentageDiscountPromotions201ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPercentageDiscountPromotions201ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPercentageDiscountPromotions201ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTPercentageDiscountPromotions201ResponseData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTPercentageDiscountPromotions201ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPercentageDiscountPromotions201ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPercentageDiscountPromotions201ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTPercentageDiscountPromotions201ResponseData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTPercentageDiscountPromotions201ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTPercentageDiscountPromotions201ResponseData) GetLinks() POSTAddresses201ResponseDataLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPercentageDiscountPromotions201ResponseData) GetLinksOk() (*POSTAddresses201ResponseDataLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTPercentageDiscountPromotions201ResponseData) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataLinks and assigns it to the Links field.
func (o *POSTPercentageDiscountPromotions201ResponseData) SetLinks(v POSTAddresses201ResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *POSTPercentageDiscountPromotions201ResponseData) GetAttributes() POSTPercentageDiscountPromotions201ResponseDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret POSTPercentageDiscountPromotions201ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPercentageDiscountPromotions201ResponseData) GetAttributesOk() (*POSTPercentageDiscountPromotions201ResponseDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *POSTPercentageDiscountPromotions201ResponseData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given POSTPercentageDiscountPromotions201ResponseDataAttributes and assigns it to the Attributes field.
func (o *POSTPercentageDiscountPromotions201ResponseData) SetAttributes(v POSTPercentageDiscountPromotions201ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTPercentageDiscountPromotions201ResponseData) GetRelationships() POSTBuyXPayYPromotions201ResponseDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTBuyXPayYPromotions201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPercentageDiscountPromotions201ResponseData) GetRelationshipsOk() (*POSTBuyXPayYPromotions201ResponseDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTPercentageDiscountPromotions201ResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTBuyXPayYPromotions201ResponseDataRelationships and assigns it to the Relationships field.
func (o *POSTPercentageDiscountPromotions201ResponseData) SetRelationships(v POSTBuyXPayYPromotions201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o POSTPercentageDiscountPromotions201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPercentageDiscountPromotions201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePOSTPercentageDiscountPromotions201ResponseData struct {
	value *POSTPercentageDiscountPromotions201ResponseData
	isSet bool
}

func (v NullablePOSTPercentageDiscountPromotions201ResponseData) Get() *POSTPercentageDiscountPromotions201ResponseData {
	return v.value
}

func (v *NullablePOSTPercentageDiscountPromotions201ResponseData) Set(val *POSTPercentageDiscountPromotions201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPercentageDiscountPromotions201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPercentageDiscountPromotions201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPercentageDiscountPromotions201ResponseData(val *POSTPercentageDiscountPromotions201ResponseData) *NullablePOSTPercentageDiscountPromotions201ResponseData {
	return &NullablePOSTPercentageDiscountPromotions201ResponseData{value: val, isSet: true}
}

func (v NullablePOSTPercentageDiscountPromotions201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPercentageDiscountPromotions201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
