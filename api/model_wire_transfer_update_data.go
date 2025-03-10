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

// checks if the WireTransferUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WireTransferUpdateData{}

// WireTransferUpdateData struct for WireTransferUpdateData
type WireTransferUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                                       `json:"id"`
	Attributes    PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentUpdateDataRelationships                                              `json:"relationships,omitempty"`
}

// NewWireTransferUpdateData instantiates a new WireTransferUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireTransferUpdateData(type_ interface{}, id interface{}, attributes PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes) *WireTransferUpdateData {
	this := WireTransferUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewWireTransferUpdateDataWithDefaults instantiates a new WireTransferUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireTransferUpdateDataWithDefaults() *WireTransferUpdateData {
	this := WireTransferUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *WireTransferUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireTransferUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WireTransferUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *WireTransferUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireTransferUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WireTransferUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *WireTransferUpdateData) GetAttributes() PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WireTransferUpdateData) GetAttributesOk() (*PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *WireTransferUpdateData) SetAttributes(v PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WireTransferUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AdyenPaymentUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireTransferUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WireTransferUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentUpdateDataRelationships and assigns it to the Relationships field.
func (o *WireTransferUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships) {
	o.Relationships = &v
}

func (o WireTransferUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WireTransferUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableWireTransferUpdateData struct {
	value *WireTransferUpdateData
	isSet bool
}

func (v NullableWireTransferUpdateData) Get() *WireTransferUpdateData {
	return v.value
}

func (v *NullableWireTransferUpdateData) Set(val *WireTransferUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableWireTransferUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableWireTransferUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireTransferUpdateData(val *WireTransferUpdateData) *NullableWireTransferUpdateData {
	return &NullableWireTransferUpdateData{value: val, isSet: true}
}

func (v NullableWireTransferUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireTransferUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
