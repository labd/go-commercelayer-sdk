/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTAddresses201ResponseDataRelationshipsGeocoderLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAddresses201ResponseDataRelationshipsGeocoderLinks{}

// POSTAddresses201ResponseDataRelationshipsGeocoderLinks struct for POSTAddresses201ResponseDataRelationshipsGeocoderLinks
type POSTAddresses201ResponseDataRelationshipsGeocoderLinks struct {
	// URL
	Self interface{} `json:"self,omitempty"`
	// URL
	Related interface{} `json:"related,omitempty"`
}

// NewPOSTAddresses201ResponseDataRelationshipsGeocoderLinks instantiates a new POSTAddresses201ResponseDataRelationshipsGeocoderLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAddresses201ResponseDataRelationshipsGeocoderLinks() *POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	this := POSTAddresses201ResponseDataRelationshipsGeocoderLinks{}
	return &this
}

// NewPOSTAddresses201ResponseDataRelationshipsGeocoderLinksWithDefaults instantiates a new POSTAddresses201ResponseDataRelationshipsGeocoderLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAddresses201ResponseDataRelationshipsGeocoderLinksWithDefaults() *POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	this := POSTAddresses201ResponseDataRelationshipsGeocoderLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderLinks) GetSelf() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderLinks) GetSelfOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return &o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderLinks) HasSelf() bool {
	if o != nil && IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given interface{} and assigns it to the Self field.
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderLinks) SetSelf(v interface{}) {
	o.Self = v
}

// GetRelated returns the Related field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderLinks) GetRelated() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderLinks) GetRelatedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Related) {
		return nil, false
	}
	return &o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderLinks) HasRelated() bool {
	if o != nil && IsNil(o.Related) {
		return true
	}

	return false
}

// SetRelated gets a reference to the given interface{} and assigns it to the Related field.
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderLinks) SetRelated(v interface{}) {
	o.Related = v
}

func (o POSTAddresses201ResponseDataRelationshipsGeocoderLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAddresses201ResponseDataRelationshipsGeocoderLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Related != nil {
		toSerialize["related"] = o.Related
	}
	return toSerialize, nil
}

type NullablePOSTAddresses201ResponseDataRelationshipsGeocoderLinks struct {
	value *POSTAddresses201ResponseDataRelationshipsGeocoderLinks
	isSet bool
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsGeocoderLinks) Get() *POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	return v.value
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsGeocoderLinks) Set(val *POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsGeocoderLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsGeocoderLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAddresses201ResponseDataRelationshipsGeocoderLinks(val *POSTAddresses201ResponseDataRelationshipsGeocoderLinks) *NullablePOSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	return &NullablePOSTAddresses201ResponseDataRelationshipsGeocoderLinks{value: val, isSet: true}
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsGeocoderLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsGeocoderLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
