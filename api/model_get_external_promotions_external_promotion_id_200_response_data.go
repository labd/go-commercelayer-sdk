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

// checks if the GETExternalPromotionsExternalPromotionId200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETExternalPromotionsExternalPromotionId200ResponseData{}

// GETExternalPromotionsExternalPromotionId200ResponseData struct for GETExternalPromotionsExternalPromotionId200ResponseData
type GETExternalPromotionsExternalPromotionId200ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                                        `json:"type,omitempty"`
	Links         *POSTAddresses201ResponseDataLinks                                 `json:"links,omitempty"`
	Attributes    *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *POSTBuyXPayYPromotions201ResponseDataRelationships                `json:"relationships,omitempty"`
}

// NewGETExternalPromotionsExternalPromotionId200ResponseData instantiates a new GETExternalPromotionsExternalPromotionId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPromotionsExternalPromotionId200ResponseData() *GETExternalPromotionsExternalPromotionId200ResponseData {
	this := GETExternalPromotionsExternalPromotionId200ResponseData{}
	return &this
}

// NewGETExternalPromotionsExternalPromotionId200ResponseDataWithDefaults instantiates a new GETExternalPromotionsExternalPromotionId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPromotionsExternalPromotionId200ResponseDataWithDefaults() *GETExternalPromotionsExternalPromotionId200ResponseData {
	this := GETExternalPromotionsExternalPromotionId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) GetLinks() POSTAddresses201ResponseDataLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) GetLinksOk() (*POSTAddresses201ResponseDataLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataLinks and assigns it to the Links field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) SetLinks(v POSTAddresses201ResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) GetAttributes() GETExternalPromotionsExternalPromotionId200ResponseDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GETExternalPromotionsExternalPromotionId200ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) GetAttributesOk() (*GETExternalPromotionsExternalPromotionId200ResponseDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GETExternalPromotionsExternalPromotionId200ResponseDataAttributes and assigns it to the Attributes field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) SetAttributes(v GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) GetRelationships() POSTBuyXPayYPromotions201ResponseDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTBuyXPayYPromotions201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) GetRelationshipsOk() (*POSTBuyXPayYPromotions201ResponseDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTBuyXPayYPromotions201ResponseDataRelationships and assigns it to the Relationships field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseData) SetRelationships(v POSTBuyXPayYPromotions201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o GETExternalPromotionsExternalPromotionId200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETExternalPromotionsExternalPromotionId200ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullableGETExternalPromotionsExternalPromotionId200ResponseData struct {
	value *GETExternalPromotionsExternalPromotionId200ResponseData
	isSet bool
}

func (v NullableGETExternalPromotionsExternalPromotionId200ResponseData) Get() *GETExternalPromotionsExternalPromotionId200ResponseData {
	return v.value
}

func (v *NullableGETExternalPromotionsExternalPromotionId200ResponseData) Set(val *GETExternalPromotionsExternalPromotionId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPromotionsExternalPromotionId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPromotionsExternalPromotionId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPromotionsExternalPromotionId200ResponseData(val *GETExternalPromotionsExternalPromotionId200ResponseData) *NullableGETExternalPromotionsExternalPromotionId200ResponseData {
	return &NullableGETExternalPromotionsExternalPromotionId200ResponseData{value: val, isSet: true}
}

func (v NullableGETExternalPromotionsExternalPromotionId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPromotionsExternalPromotionId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
