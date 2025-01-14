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

// checks if the POSTFlexPromotions201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTFlexPromotions201ResponseDataRelationships{}

// POSTFlexPromotions201ResponseDataRelationships struct for POSTFlexPromotions201ResponseDataRelationships
type POSTFlexPromotions201ResponseDataRelationships struct {
	CouponCodesPromotionRule *POSTBuyXPayYPromotions201ResponseDataRelationshipsCouponCodesPromotionRule `json:"coupon_codes_promotion_rule,omitempty"`
	Coupons                  *POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons                  `json:"coupons,omitempty"`
	Attachments              *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments    `json:"attachments,omitempty"`
	Events                   *POSTAddresses201ResponseDataRelationshipsEvents                            `json:"events,omitempty"`
	Tags                     *POSTAddresses201ResponseDataRelationshipsTags                              `json:"tags,omitempty"`
	Versions                 *POSTAddresses201ResponseDataRelationshipsVersions                          `json:"versions,omitempty"`
}

// NewPOSTFlexPromotions201ResponseDataRelationships instantiates a new POSTFlexPromotions201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTFlexPromotions201ResponseDataRelationships() *POSTFlexPromotions201ResponseDataRelationships {
	this := POSTFlexPromotions201ResponseDataRelationships{}
	return &this
}

// NewPOSTFlexPromotions201ResponseDataRelationshipsWithDefaults instantiates a new POSTFlexPromotions201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTFlexPromotions201ResponseDataRelationshipsWithDefaults() *POSTFlexPromotions201ResponseDataRelationships {
	this := POSTFlexPromotions201ResponseDataRelationships{}
	return &this
}

// GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field value if set, zero value otherwise.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetCouponCodesPromotionRule() POSTBuyXPayYPromotions201ResponseDataRelationshipsCouponCodesPromotionRule {
	if o == nil || IsNil(o.CouponCodesPromotionRule) {
		var ret POSTBuyXPayYPromotions201ResponseDataRelationshipsCouponCodesPromotionRule
		return ret
	}
	return *o.CouponCodesPromotionRule
}

// GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetCouponCodesPromotionRuleOk() (*POSTBuyXPayYPromotions201ResponseDataRelationshipsCouponCodesPromotionRule, bool) {
	if o == nil || IsNil(o.CouponCodesPromotionRule) {
		return nil, false
	}
	return o.CouponCodesPromotionRule, true
}

// HasCouponCodesPromotionRule returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) HasCouponCodesPromotionRule() bool {
	if o != nil && !IsNil(o.CouponCodesPromotionRule) {
		return true
	}

	return false
}

// SetCouponCodesPromotionRule gets a reference to the given POSTBuyXPayYPromotions201ResponseDataRelationshipsCouponCodesPromotionRule and assigns it to the CouponCodesPromotionRule field.
func (o *POSTFlexPromotions201ResponseDataRelationships) SetCouponCodesPromotionRule(v POSTBuyXPayYPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) {
	o.CouponCodesPromotionRule = &v
}

// GetCoupons returns the Coupons field value if set, zero value otherwise.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetCoupons() POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons {
	if o == nil || IsNil(o.Coupons) {
		var ret POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons
		return ret
	}
	return *o.Coupons
}

// GetCouponsOk returns a tuple with the Coupons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetCouponsOk() (*POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons, bool) {
	if o == nil || IsNil(o.Coupons) {
		return nil, false
	}
	return o.Coupons, true
}

// HasCoupons returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) HasCoupons() bool {
	if o != nil && !IsNil(o.Coupons) {
		return true
	}

	return false
}

// SetCoupons gets a reference to the given POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons and assigns it to the Coupons field.
func (o *POSTFlexPromotions201ResponseDataRelationships) SetCoupons(v POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons) {
	o.Coupons = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTFlexPromotions201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *POSTFlexPromotions201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetTags() POSTAddresses201ResponseDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret POSTAddresses201ResponseDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetTagsOk() (*POSTAddresses201ResponseDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given POSTAddresses201ResponseDataRelationshipsTags and assigns it to the Tags field.
func (o *POSTFlexPromotions201ResponseDataRelationships) SetTags(v POSTAddresses201ResponseDataRelationshipsTags) {
	o.Tags = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTFlexPromotions201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTFlexPromotions201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTFlexPromotions201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CouponCodesPromotionRule) {
		toSerialize["coupon_codes_promotion_rule"] = o.CouponCodesPromotionRule
	}
	if !IsNil(o.Coupons) {
		toSerialize["coupons"] = o.Coupons
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTFlexPromotions201ResponseDataRelationships struct {
	value *POSTFlexPromotions201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTFlexPromotions201ResponseDataRelationships) Get() *POSTFlexPromotions201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTFlexPromotions201ResponseDataRelationships) Set(val *POSTFlexPromotions201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTFlexPromotions201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTFlexPromotions201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTFlexPromotions201ResponseDataRelationships(val *POSTFlexPromotions201ResponseDataRelationships) *NullablePOSTFlexPromotions201ResponseDataRelationships {
	return &NullablePOSTFlexPromotions201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTFlexPromotions201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTFlexPromotions201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
