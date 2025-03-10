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

// checks if the PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes{}

// PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes struct for PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes
type PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes struct {
	// The promotion's internal name.
	Name interface{} `json:"name,omitempty"`
	// Indicates if the promotion will be applied exclusively, based on its priority score.
	Exclusive interface{} `json:"exclusive,omitempty"`
	// The priority assigned to the promotion (lower means higher priority).
	Priority interface{} `json:"priority,omitempty"`
	// The activation date/time of this promotion.
	StartsAt interface{} `json:"starts_at,omitempty"`
	// The expiration date/time of this promotion (must be after starts_at).
	ExpiresAt interface{} `json:"expires_at,omitempty"`
	// The total number of times this promotion can be applied. When 'null' it means promotion can be applied infinite times.
	TotalUsageLimit interface{} `json:"total_usage_limit,omitempty"`
	// The discount rule to be applied.
	Rules interface{} `json:"rules,omitempty"`
	// Send this attribute if you want to mark this resource as disabled.
	Disable interface{} `json:"_disable,omitempty"`
	// Send this attribute if you want to mark this resource as enabled.
	Enable interface{} `json:"_enable,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes instantiates a new PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes() *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes {
	this := PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes{}
	return &this
}

// NewPATCHFlexPromotionsFlexPromotionId200ResponseDataAttributesWithDefaults instantiates a new PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHFlexPromotionsFlexPromotionId200ResponseDataAttributesWithDefaults() *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes {
	this := PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetExclusive returns the Exclusive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetExclusive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Exclusive
}

// GetExclusiveOk returns a tuple with the Exclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetExclusiveOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Exclusive) {
		return nil, false
	}
	return &o.Exclusive, true
}

// HasExclusive returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasExclusive() bool {
	if o != nil && IsNil(o.Exclusive) {
		return true
	}

	return false
}

// SetExclusive gets a reference to the given interface{} and assigns it to the Exclusive field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetExclusive(v interface{}) {
	o.Exclusive = v
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetPriority() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetPriorityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return &o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasPriority() bool {
	if o != nil && IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given interface{} and assigns it to the Priority field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetPriority(v interface{}) {
	o.Priority = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return &o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasStartsAt() bool {
	if o != nil && IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given interface{} and assigns it to the StartsAt field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetTotalUsageLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalUsageLimit) {
		return nil, false
	}
	return &o.TotalUsageLimit, true
}

// HasTotalUsageLimit returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasTotalUsageLimit() bool {
	if o != nil && IsNil(o.TotalUsageLimit) {
		return true
	}

	return false
}

// SetTotalUsageLimit gets a reference to the given interface{} and assigns it to the TotalUsageLimit field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetTotalUsageLimit(v interface{}) {
	o.TotalUsageLimit = v
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetRules() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetRulesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasRules() bool {
	if o != nil && IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given interface{} and assigns it to the Rules field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetRules(v interface{}) {
	o.Rules = v
}

// GetDisable returns the Disable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetDisable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetDisableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return &o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasDisable() bool {
	if o != nil && IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given interface{} and assigns it to the Disable field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetDisable(v interface{}) {
	o.Disable = v
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetEnable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetEnableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return &o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasEnable() bool {
	if o != nil && IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given interface{} and assigns it to the Enable field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetEnable(v interface{}) {
	o.Enable = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Exclusive != nil {
		toSerialize["exclusive"] = o.Exclusive
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.StartsAt != nil {
		toSerialize["starts_at"] = o.StartsAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.TotalUsageLimit != nil {
		toSerialize["total_usage_limit"] = o.TotalUsageLimit
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Disable != nil {
		toSerialize["_disable"] = o.Disable
	}
	if o.Enable != nil {
		toSerialize["_enable"] = o.Enable
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

type NullablePATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes struct {
	value *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) Get() *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) Set(val *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes(val *PATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) *NullablePATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes {
	return &NullablePATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
