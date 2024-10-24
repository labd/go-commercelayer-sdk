/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETPromotionsPromotionId200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETPromotionsPromotionId200ResponseData{}

// GETPromotionsPromotionId200ResponseData struct for GETPromotionsPromotionId200ResponseData
type GETPromotionsPromotionId200ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                                                `json:"type,omitempty"`
	Links         *POSTAddresses201ResponseDataLinks                                         `json:"links,omitempty"`
	Attributes    *GETFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *POSTExternalPromotions201ResponseDataRelationships                        `json:"relationships,omitempty"`
}

// NewGETPromotionsPromotionId200ResponseData instantiates a new GETPromotionsPromotionId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPromotionsPromotionId200ResponseData() *GETPromotionsPromotionId200ResponseData {
	this := GETPromotionsPromotionId200ResponseData{}
	return &this
}

// NewGETPromotionsPromotionId200ResponseDataWithDefaults instantiates a new GETPromotionsPromotionId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPromotionsPromotionId200ResponseDataWithDefaults() *GETPromotionsPromotionId200ResponseData {
	this := GETPromotionsPromotionId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPromotionsPromotionId200ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPromotionsPromotionId200ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETPromotionsPromotionId200ResponseData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETPromotionsPromotionId200ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPromotionsPromotionId200ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPromotionsPromotionId200ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETPromotionsPromotionId200ResponseData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETPromotionsPromotionId200ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETPromotionsPromotionId200ResponseData) GetLinks() POSTAddresses201ResponseDataLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPromotionsPromotionId200ResponseData) GetLinksOk() (*POSTAddresses201ResponseDataLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETPromotionsPromotionId200ResponseData) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataLinks and assigns it to the Links field.
func (o *GETPromotionsPromotionId200ResponseData) SetLinks(v POSTAddresses201ResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GETPromotionsPromotionId200ResponseData) GetAttributes() GETFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GETFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPromotionsPromotionId200ResponseData) GetAttributesOk() (*GETFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GETPromotionsPromotionId200ResponseData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GETFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes and assigns it to the Attributes field.
func (o *GETPromotionsPromotionId200ResponseData) SetAttributes(v GETFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GETPromotionsPromotionId200ResponseData) GetRelationships() POSTExternalPromotions201ResponseDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTExternalPromotions201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPromotionsPromotionId200ResponseData) GetRelationshipsOk() (*POSTExternalPromotions201ResponseDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GETPromotionsPromotionId200ResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTExternalPromotions201ResponseDataRelationships and assigns it to the Relationships field.
func (o *GETPromotionsPromotionId200ResponseData) SetRelationships(v POSTExternalPromotions201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o GETPromotionsPromotionId200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETPromotionsPromotionId200ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullableGETPromotionsPromotionId200ResponseData struct {
	value *GETPromotionsPromotionId200ResponseData
	isSet bool
}

func (v NullableGETPromotionsPromotionId200ResponseData) Get() *GETPromotionsPromotionId200ResponseData {
	return v.value
}

func (v *NullableGETPromotionsPromotionId200ResponseData) Set(val *GETPromotionsPromotionId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPromotionsPromotionId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPromotionsPromotionId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPromotionsPromotionId200ResponseData(val *GETPromotionsPromotionId200ResponseData) *NullableGETPromotionsPromotionId200ResponseData {
	return &NullableGETPromotionsPromotionId200ResponseData{value: val, isSet: true}
}

func (v NullableGETPromotionsPromotionId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPromotionsPromotionId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}