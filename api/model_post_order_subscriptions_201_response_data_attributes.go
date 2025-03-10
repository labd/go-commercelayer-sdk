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

// checks if the POSTOrderSubscriptions201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderSubscriptions201ResponseDataAttributes{}

// POSTOrderSubscriptions201ResponseDataAttributes struct for POSTOrderSubscriptions201ResponseDataAttributes
type POSTOrderSubscriptions201ResponseDataAttributes struct {
	// The frequency of the subscription. Use one of the supported within 'hourly', 'daily', 'weekly', 'monthly', 'two-month', 'three-month', 'four-month', 'six-month', 'yearly', or provide your custom crontab expression (min unit is hour). Must be supported by existing associated subscription_model.
	Frequency interface{} `json:"frequency"`
	// Indicates if the subscription will be activated considering the placed source order as its first run.
	ActivateBySourceOrder interface{} `json:"activate_by_source_order,omitempty"`
	// Indicates if the subscription created orders are automatically placed at the end of the copy.
	PlaceTargetOrder interface{} `json:"place_target_order,omitempty"`
	// Indicates the number of hours the renewal alert will be triggered before the subscription next run. Must be included between 1 and 720 hours.
	RenewalAlertPeriod interface{} `json:"renewal_alert_period,omitempty"`
	// The activation date/time of this subscription.
	StartsAt interface{} `json:"starts_at,omitempty"`
	// The expiration date/time of this subscription (must be after starts_at).
	ExpiresAt interface{} `json:"expires_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTOrderSubscriptions201ResponseDataAttributes instantiates a new POSTOrderSubscriptions201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderSubscriptions201ResponseDataAttributes(frequency interface{}) *POSTOrderSubscriptions201ResponseDataAttributes {
	this := POSTOrderSubscriptions201ResponseDataAttributes{}
	this.Frequency = frequency
	return &this
}

// NewPOSTOrderSubscriptions201ResponseDataAttributesWithDefaults instantiates a new POSTOrderSubscriptions201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderSubscriptions201ResponseDataAttributesWithDefaults() *POSTOrderSubscriptions201ResponseDataAttributes {
	this := POSTOrderSubscriptions201ResponseDataAttributes{}
	return &this
}

// GetFrequency returns the Frequency field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetFrequency() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetFrequencyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetFrequency(v interface{}) {
	o.Frequency = v
}

// GetActivateBySourceOrder returns the ActivateBySourceOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetActivateBySourceOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ActivateBySourceOrder
}

// GetActivateBySourceOrderOk returns a tuple with the ActivateBySourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetActivateBySourceOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ActivateBySourceOrder) {
		return nil, false
	}
	return &o.ActivateBySourceOrder, true
}

// HasActivateBySourceOrder returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasActivateBySourceOrder() bool {
	if o != nil && IsNil(o.ActivateBySourceOrder) {
		return true
	}

	return false
}

// SetActivateBySourceOrder gets a reference to the given interface{} and assigns it to the ActivateBySourceOrder field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetActivateBySourceOrder(v interface{}) {
	o.ActivateBySourceOrder = v
}

// GetPlaceTargetOrder returns the PlaceTargetOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetPlaceTargetOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PlaceTargetOrder
}

// GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetPlaceTargetOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PlaceTargetOrder) {
		return nil, false
	}
	return &o.PlaceTargetOrder, true
}

// HasPlaceTargetOrder returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasPlaceTargetOrder() bool {
	if o != nil && IsNil(o.PlaceTargetOrder) {
		return true
	}

	return false
}

// SetPlaceTargetOrder gets a reference to the given interface{} and assigns it to the PlaceTargetOrder field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetPlaceTargetOrder(v interface{}) {
	o.PlaceTargetOrder = v
}

// GetRenewalAlertPeriod returns the RenewalAlertPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetRenewalAlertPeriod() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RenewalAlertPeriod
}

// GetRenewalAlertPeriodOk returns a tuple with the RenewalAlertPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetRenewalAlertPeriodOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RenewalAlertPeriod) {
		return nil, false
	}
	return &o.RenewalAlertPeriod, true
}

// HasRenewalAlertPeriod returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasRenewalAlertPeriod() bool {
	if o != nil && IsNil(o.RenewalAlertPeriod) {
		return true
	}

	return false
}

// SetRenewalAlertPeriod gets a reference to the given interface{} and assigns it to the RenewalAlertPeriod field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetRenewalAlertPeriod(v interface{}) {
	o.RenewalAlertPeriod = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return &o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasStartsAt() bool {
	if o != nil && IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given interface{} and assigns it to the StartsAt field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTOrderSubscriptions201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderSubscriptions201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	if o.ActivateBySourceOrder != nil {
		toSerialize["activate_by_source_order"] = o.ActivateBySourceOrder
	}
	if o.PlaceTargetOrder != nil {
		toSerialize["place_target_order"] = o.PlaceTargetOrder
	}
	if o.RenewalAlertPeriod != nil {
		toSerialize["renewal_alert_period"] = o.RenewalAlertPeriod
	}
	if o.StartsAt != nil {
		toSerialize["starts_at"] = o.StartsAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
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

type NullablePOSTOrderSubscriptions201ResponseDataAttributes struct {
	value *POSTOrderSubscriptions201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTOrderSubscriptions201ResponseDataAttributes) Get() *POSTOrderSubscriptions201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataAttributes) Set(val *POSTOrderSubscriptions201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderSubscriptions201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderSubscriptions201ResponseDataAttributes(val *POSTOrderSubscriptions201ResponseDataAttributes) *NullablePOSTOrderSubscriptions201ResponseDataAttributes {
	return &NullablePOSTOrderSubscriptions201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTOrderSubscriptions201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
