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

// checks if the GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes{}

// GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes struct for GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes
type GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes struct {
	// The delivery lead minimum time (in hours) when shipping from the associated stock location with the associated shipping method.
	MinHours interface{} `json:"min_hours,omitempty"`
	// The delivery lead maximun time (in hours) when shipping from the associated stock location with the associated shipping method.
	MaxHours interface{} `json:"max_hours,omitempty"`
	// The delivery lead minimum time, in days (rounded).
	MinDays interface{} `json:"min_days,omitempty"`
	// The delivery lead maximun time, in days (rounded).
	MaxDays interface{} `json:"max_days,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes instantiates a new GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes() *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes {
	this := GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes{}
	return &this
}

// NewGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributesWithDefaults instantiates a new GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributesWithDefaults() *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes {
	this := GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes{}
	return &this
}

// GetMinHours returns the MinHours field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetMinHours() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MinHours
}

// GetMinHoursOk returns a tuple with the MinHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetMinHoursOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MinHours) {
		return nil, false
	}
	return &o.MinHours, true
}

// HasMinHours returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) HasMinHours() bool {
	if o != nil && IsNil(o.MinHours) {
		return true
	}

	return false
}

// SetMinHours gets a reference to the given interface{} and assigns it to the MinHours field.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) SetMinHours(v interface{}) {
	o.MinHours = v
}

// GetMaxHours returns the MaxHours field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetMaxHours() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MaxHours
}

// GetMaxHoursOk returns a tuple with the MaxHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetMaxHoursOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MaxHours) {
		return nil, false
	}
	return &o.MaxHours, true
}

// HasMaxHours returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) HasMaxHours() bool {
	if o != nil && IsNil(o.MaxHours) {
		return true
	}

	return false
}

// SetMaxHours gets a reference to the given interface{} and assigns it to the MaxHours field.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) SetMaxHours(v interface{}) {
	o.MaxHours = v
}

// GetMinDays returns the MinDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetMinDays() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MinDays
}

// GetMinDaysOk returns a tuple with the MinDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetMinDaysOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MinDays) {
		return nil, false
	}
	return &o.MinDays, true
}

// HasMinDays returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) HasMinDays() bool {
	if o != nil && IsNil(o.MinDays) {
		return true
	}

	return false
}

// SetMinDays gets a reference to the given interface{} and assigns it to the MinDays field.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) SetMinDays(v interface{}) {
	o.MinDays = v
}

// GetMaxDays returns the MaxDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetMaxDays() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MaxDays
}

// GetMaxDaysOk returns a tuple with the MaxDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetMaxDaysOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MaxDays) {
		return nil, false
	}
	return &o.MaxDays, true
}

// HasMaxDays returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) HasMaxDays() bool {
	if o != nil && IsNil(o.MaxDays) {
		return true
	}

	return false
}

// SetMaxDays gets a reference to the given interface{} and assigns it to the MaxDays field.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) SetMaxDays(v interface{}) {
	o.MaxDays = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MinHours != nil {
		toSerialize["min_hours"] = o.MinHours
	}
	if o.MaxHours != nil {
		toSerialize["max_hours"] = o.MaxHours
	}
	if o.MinDays != nil {
		toSerialize["min_days"] = o.MinDays
	}
	if o.MaxDays != nil {
		toSerialize["max_days"] = o.MaxDays
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes struct {
	value *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) Get() *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) Set(val *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes(val *GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) *NullableGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes {
	return &NullableGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
