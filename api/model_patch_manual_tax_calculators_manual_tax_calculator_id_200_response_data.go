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

// checks if the PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData{}

// PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData struct for PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData
type PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                                              `json:"type,omitempty"`
	Links         *POSTAddresses201ResponseDataLinks                                       `json:"links,omitempty"`
	Attributes    *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *POSTManualTaxCalculators201ResponseDataRelationships                    `json:"relationships,omitempty"`
}

// NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData instantiates a new PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData() *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData {
	this := PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData{}
	return &this
}

// NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataWithDefaults instantiates a new PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataWithDefaults() *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData {
	this := PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) GetLinks() POSTAddresses201ResponseDataLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) GetLinksOk() (*POSTAddresses201ResponseDataLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataLinks and assigns it to the Links field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) SetLinks(v POSTAddresses201ResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) GetAttributes() PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) GetAttributesOk() (*PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes and assigns it to the Attributes field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) SetAttributes(v PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) GetRelationships() POSTManualTaxCalculators201ResponseDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTManualTaxCalculators201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) GetRelationshipsOk() (*POSTManualTaxCalculators201ResponseDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTManualTaxCalculators201ResponseDataRelationships and assigns it to the Relationships field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) SetRelationships(v POSTManualTaxCalculators201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData struct {
	value *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData
	isSet bool
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) Get() *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData {
	return v.value
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) Set(val *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData(val *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData {
	return &NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData{value: val, isSet: true}
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
