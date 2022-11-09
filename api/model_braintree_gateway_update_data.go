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

// BraintreeGatewayUpdateData struct for BraintreeGatewayUpdateData
type BraintreeGatewayUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                            `json:"id"`
	Attributes    PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes `json:"attributes"`
	Relationships *BraintreeGatewayCreateDataRelationships                          `json:"relationships,omitempty"`
}

// NewBraintreeGatewayUpdateData instantiates a new BraintreeGatewayUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreeGatewayUpdateData(type_ string, id string, attributes PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) *BraintreeGatewayUpdateData {
	this := BraintreeGatewayUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewBraintreeGatewayUpdateDataWithDefaults instantiates a new BraintreeGatewayUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreeGatewayUpdateDataWithDefaults() *BraintreeGatewayUpdateData {
	this := BraintreeGatewayUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *BraintreeGatewayUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BraintreeGatewayUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BraintreeGatewayUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BraintreeGatewayUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *BraintreeGatewayUpdateData) GetAttributes() PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayUpdateData) GetAttributesOk() (*PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BraintreeGatewayUpdateData) SetAttributes(v PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BraintreeGatewayUpdateData) GetRelationships() BraintreeGatewayCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret BraintreeGatewayCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayUpdateData) GetRelationshipsOk() (*BraintreeGatewayCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BraintreeGatewayUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BraintreeGatewayCreateDataRelationships and assigns it to the Relationships field.
func (o *BraintreeGatewayUpdateData) SetRelationships(v BraintreeGatewayCreateDataRelationships) {
	o.Relationships = &v
}

func (o BraintreeGatewayUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableBraintreeGatewayUpdateData struct {
	value *BraintreeGatewayUpdateData
	isSet bool
}

func (v NullableBraintreeGatewayUpdateData) Get() *BraintreeGatewayUpdateData {
	return v.value
}

func (v *NullableBraintreeGatewayUpdateData) Set(val *BraintreeGatewayUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreeGatewayUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreeGatewayUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreeGatewayUpdateData(val *BraintreeGatewayUpdateData) *NullableBraintreeGatewayUpdateData {
	return &NullableBraintreeGatewayUpdateData{value: val, isSet: true}
}

func (v NullableBraintreeGatewayUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreeGatewayUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
