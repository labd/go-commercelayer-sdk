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

// checks if the PATCHInventoryReturnLocationsInventoryReturnLocationId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHInventoryReturnLocationsInventoryReturnLocationId200Response{}

// PATCHInventoryReturnLocationsInventoryReturnLocationId200Response struct for PATCHInventoryReturnLocationsInventoryReturnLocationId200Response
type PATCHInventoryReturnLocationsInventoryReturnLocationId200Response struct {
	Data *PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseData `json:"data,omitempty"`
}

// NewPATCHInventoryReturnLocationsInventoryReturnLocationId200Response instantiates a new PATCHInventoryReturnLocationsInventoryReturnLocationId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHInventoryReturnLocationsInventoryReturnLocationId200Response() *PATCHInventoryReturnLocationsInventoryReturnLocationId200Response {
	this := PATCHInventoryReturnLocationsInventoryReturnLocationId200Response{}
	return &this
}

// NewPATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseWithDefaults instantiates a new PATCHInventoryReturnLocationsInventoryReturnLocationId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseWithDefaults() *PATCHInventoryReturnLocationsInventoryReturnLocationId200Response {
	this := PATCHInventoryReturnLocationsInventoryReturnLocationId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHInventoryReturnLocationsInventoryReturnLocationId200Response) GetData() PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHInventoryReturnLocationsInventoryReturnLocationId200Response) GetDataOk() (*PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHInventoryReturnLocationsInventoryReturnLocationId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseData and assigns it to the Data field.
func (o *PATCHInventoryReturnLocationsInventoryReturnLocationId200Response) SetData(v PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseData) {
	o.Data = &v
}

func (o PATCHInventoryReturnLocationsInventoryReturnLocationId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHInventoryReturnLocationsInventoryReturnLocationId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHInventoryReturnLocationsInventoryReturnLocationId200Response struct {
	value *PATCHInventoryReturnLocationsInventoryReturnLocationId200Response
	isSet bool
}

func (v NullablePATCHInventoryReturnLocationsInventoryReturnLocationId200Response) Get() *PATCHInventoryReturnLocationsInventoryReturnLocationId200Response {
	return v.value
}

func (v *NullablePATCHInventoryReturnLocationsInventoryReturnLocationId200Response) Set(val *PATCHInventoryReturnLocationsInventoryReturnLocationId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHInventoryReturnLocationsInventoryReturnLocationId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHInventoryReturnLocationsInventoryReturnLocationId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHInventoryReturnLocationsInventoryReturnLocationId200Response(val *PATCHInventoryReturnLocationsInventoryReturnLocationId200Response) *NullablePATCHInventoryReturnLocationsInventoryReturnLocationId200Response {
	return &NullablePATCHInventoryReturnLocationsInventoryReturnLocationId200Response{value: val, isSet: true}
}

func (v NullablePATCHInventoryReturnLocationsInventoryReturnLocationId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHInventoryReturnLocationsInventoryReturnLocationId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
