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

// POSTAdyenGateways201ResponseData struct for POSTAdyenGateways201ResponseData
type POSTAdyenGateways201ResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                            `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks             `json:"links,omitempty"`
	Attributes    *POSTAdyenGateways201ResponseDataAttributes        `json:"attributes,omitempty"`
	Relationships *GETAdyenGateways200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewPOSTAdyenGateways201ResponseData instantiates a new POSTAdyenGateways201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdyenGateways201ResponseData() *POSTAdyenGateways201ResponseData {
	this := POSTAdyenGateways201ResponseData{}
	return &this
}

// NewPOSTAdyenGateways201ResponseDataWithDefaults instantiates a new POSTAdyenGateways201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdyenGateways201ResponseDataWithDefaults() *POSTAdyenGateways201ResponseData {
	this := POSTAdyenGateways201ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *POSTAdyenGateways201ResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *POSTAdyenGateways201ResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *POSTAdyenGateways201ResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *POSTAdyenGateways201ResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTAdyenGateways201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *POSTAdyenGateways201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *POSTAdyenGateways201ResponseData) GetAttributes() POSTAdyenGateways201ResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret POSTAdyenGateways201ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseData) GetAttributesOk() (*POSTAdyenGateways201ResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given POSTAdyenGateways201ResponseDataAttributes and assigns it to the Attributes field.
func (o *POSTAdyenGateways201ResponseData) SetAttributes(v POSTAdyenGateways201ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTAdyenGateways201ResponseData) GetRelationships() GETAdyenGateways200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseData) GetRelationshipsOk() (*GETAdyenGateways200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *POSTAdyenGateways201ResponseData) SetRelationships(v GETAdyenGateways200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o POSTAdyenGateways201ResponseData) MarshalJSON() ([]byte, error) {
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

type NullablePOSTAdyenGateways201ResponseData struct {
	value *POSTAdyenGateways201ResponseData
	isSet bool
}

func (v NullablePOSTAdyenGateways201ResponseData) Get() *POSTAdyenGateways201ResponseData {
	return v.value
}

func (v *NullablePOSTAdyenGateways201ResponseData) Set(val *POSTAdyenGateways201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdyenGateways201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdyenGateways201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdyenGateways201ResponseData(val *POSTAdyenGateways201ResponseData) *NullablePOSTAdyenGateways201ResponseData {
	return &NullablePOSTAdyenGateways201ResponseData{value: val, isSet: true}
}

func (v NullablePOSTAdyenGateways201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdyenGateways201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
