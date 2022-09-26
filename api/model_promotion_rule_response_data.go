/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PromotionRuleResponseData struct for PromotionRuleResponseData
type PromotionRuleResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                            `json:"type,omitempty"`
	Links         *AddressResponseDataLinks                          `json:"links,omitempty"`
	Attributes    *Attributes                                        `json:"attributes,omitempty"`
	Relationships *OrderAmountPromotionRuleResponseDataRelationships `json:"relationships,omitempty"`
}

// NewPromotionRuleResponseData instantiates a new PromotionRuleResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionRuleResponseData() *PromotionRuleResponseData {
	this := PromotionRuleResponseData{}
	return &this
}

// NewPromotionRuleResponseDataWithDefaults instantiates a new PromotionRuleResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionRuleResponseDataWithDefaults() *PromotionRuleResponseData {
	this := PromotionRuleResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PromotionRuleResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRuleResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PromotionRuleResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PromotionRuleResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PromotionRuleResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRuleResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PromotionRuleResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PromotionRuleResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PromotionRuleResponseData) GetLinks() AddressResponseDataLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRuleResponseData) GetLinksOk() (*AddressResponseDataLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PromotionRuleResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataLinks and assigns it to the Links field.
func (o *PromotionRuleResponseData) SetLinks(v AddressResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PromotionRuleResponseData) GetAttributes() Attributes {
	if o == nil || o.Attributes == nil {
		var ret Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRuleResponseData) GetAttributesOk() (*Attributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PromotionRuleResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Attributes and assigns it to the Attributes field.
func (o *PromotionRuleResponseData) SetAttributes(v Attributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PromotionRuleResponseData) GetRelationships() OrderAmountPromotionRuleResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret OrderAmountPromotionRuleResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRuleResponseData) GetRelationshipsOk() (*OrderAmountPromotionRuleResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PromotionRuleResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrderAmountPromotionRuleResponseDataRelationships and assigns it to the Relationships field.
func (o *PromotionRuleResponseData) SetRelationships(v OrderAmountPromotionRuleResponseDataRelationships) {
	o.Relationships = &v
}

func (o PromotionRuleResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullablePromotionRuleResponseData struct {
	value *PromotionRuleResponseData
	isSet bool
}

func (v NullablePromotionRuleResponseData) Get() *PromotionRuleResponseData {
	return v.value
}

func (v *NullablePromotionRuleResponseData) Set(val *PromotionRuleResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionRuleResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionRuleResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionRuleResponseData(val *PromotionRuleResponseData) *NullablePromotionRuleResponseData {
	return &NullablePromotionRuleResponseData{value: val, isSet: true}
}

func (v NullablePromotionRuleResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionRuleResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
