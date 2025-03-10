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

// checks if the FreeShippingPromotionDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreeShippingPromotionDataRelationships{}

// FreeShippingPromotionDataRelationships struct for FreeShippingPromotionDataRelationships
type FreeShippingPromotionDataRelationships struct {
	Market                   *AvalaraAccountDataRelationshipsMarkets                     `json:"market,omitempty"`
	PromotionRules           *BuyXPayYPromotionDataRelationshipsPromotionRules           `json:"promotion_rules,omitempty"`
	OrderAmountPromotionRule *BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRule `json:"order_amount_promotion_rule,omitempty"`
	SkuListPromotionRule     *BuyXPayYPromotionDataRelationshipsSkuListPromotionRule     `json:"sku_list_promotion_rule,omitempty"`
	CouponCodesPromotionRule *BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule `json:"coupon_codes_promotion_rule,omitempty"`
	CustomPromotionRule      *BuyXPayYPromotionDataRelationshipsCustomPromotionRule      `json:"custom_promotion_rule,omitempty"`
	SkuList                  *BundleDataRelationshipsSkuList                             `json:"sku_list,omitempty"`
	Coupons                  *BuyXPayYPromotionDataRelationshipsCoupons                  `json:"coupons,omitempty"`
	Attachments              *AuthorizationDataRelationshipsAttachments                  `json:"attachments,omitempty"`
	Events                   *AddressDataRelationshipsEvents                             `json:"events,omitempty"`
	Tags                     *AddressDataRelationshipsTags                               `json:"tags,omitempty"`
	Versions                 *AddressDataRelationshipsVersions                           `json:"versions,omitempty"`
}

// NewFreeShippingPromotionDataRelationships instantiates a new FreeShippingPromotionDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeShippingPromotionDataRelationships() *FreeShippingPromotionDataRelationships {
	this := FreeShippingPromotionDataRelationships{}
	return &this
}

// NewFreeShippingPromotionDataRelationshipsWithDefaults instantiates a new FreeShippingPromotionDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeShippingPromotionDataRelationshipsWithDefaults() *FreeShippingPromotionDataRelationships {
	this := FreeShippingPromotionDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || IsNil(o.Market) {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *FreeShippingPromotionDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetPromotionRules returns the PromotionRules field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetPromotionRules() BuyXPayYPromotionDataRelationshipsPromotionRules {
	if o == nil || IsNil(o.PromotionRules) {
		var ret BuyXPayYPromotionDataRelationshipsPromotionRules
		return ret
	}
	return *o.PromotionRules
}

// GetPromotionRulesOk returns a tuple with the PromotionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetPromotionRulesOk() (*BuyXPayYPromotionDataRelationshipsPromotionRules, bool) {
	if o == nil || IsNil(o.PromotionRules) {
		return nil, false
	}
	return o.PromotionRules, true
}

// HasPromotionRules returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasPromotionRules() bool {
	if o != nil && !IsNil(o.PromotionRules) {
		return true
	}

	return false
}

// SetPromotionRules gets a reference to the given BuyXPayYPromotionDataRelationshipsPromotionRules and assigns it to the PromotionRules field.
func (o *FreeShippingPromotionDataRelationships) SetPromotionRules(v BuyXPayYPromotionDataRelationshipsPromotionRules) {
	o.PromotionRules = &v
}

// GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetOrderAmountPromotionRule() BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRule {
	if o == nil || IsNil(o.OrderAmountPromotionRule) {
		var ret BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRule
		return ret
	}
	return *o.OrderAmountPromotionRule
}

// GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetOrderAmountPromotionRuleOk() (*BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRule, bool) {
	if o == nil || IsNil(o.OrderAmountPromotionRule) {
		return nil, false
	}
	return o.OrderAmountPromotionRule, true
}

// HasOrderAmountPromotionRule returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasOrderAmountPromotionRule() bool {
	if o != nil && !IsNil(o.OrderAmountPromotionRule) {
		return true
	}

	return false
}

// SetOrderAmountPromotionRule gets a reference to the given BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRule and assigns it to the OrderAmountPromotionRule field.
func (o *FreeShippingPromotionDataRelationships) SetOrderAmountPromotionRule(v BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRule) {
	o.OrderAmountPromotionRule = &v
}

// GetSkuListPromotionRule returns the SkuListPromotionRule field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetSkuListPromotionRule() BuyXPayYPromotionDataRelationshipsSkuListPromotionRule {
	if o == nil || IsNil(o.SkuListPromotionRule) {
		var ret BuyXPayYPromotionDataRelationshipsSkuListPromotionRule
		return ret
	}
	return *o.SkuListPromotionRule
}

// GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetSkuListPromotionRuleOk() (*BuyXPayYPromotionDataRelationshipsSkuListPromotionRule, bool) {
	if o == nil || IsNil(o.SkuListPromotionRule) {
		return nil, false
	}
	return o.SkuListPromotionRule, true
}

// HasSkuListPromotionRule returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasSkuListPromotionRule() bool {
	if o != nil && !IsNil(o.SkuListPromotionRule) {
		return true
	}

	return false
}

// SetSkuListPromotionRule gets a reference to the given BuyXPayYPromotionDataRelationshipsSkuListPromotionRule and assigns it to the SkuListPromotionRule field.
func (o *FreeShippingPromotionDataRelationships) SetSkuListPromotionRule(v BuyXPayYPromotionDataRelationshipsSkuListPromotionRule) {
	o.SkuListPromotionRule = &v
}

// GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetCouponCodesPromotionRule() BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule {
	if o == nil || IsNil(o.CouponCodesPromotionRule) {
		var ret BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule
		return ret
	}
	return *o.CouponCodesPromotionRule
}

// GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetCouponCodesPromotionRuleOk() (*BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule, bool) {
	if o == nil || IsNil(o.CouponCodesPromotionRule) {
		return nil, false
	}
	return o.CouponCodesPromotionRule, true
}

// HasCouponCodesPromotionRule returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasCouponCodesPromotionRule() bool {
	if o != nil && !IsNil(o.CouponCodesPromotionRule) {
		return true
	}

	return false
}

// SetCouponCodesPromotionRule gets a reference to the given BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule and assigns it to the CouponCodesPromotionRule field.
func (o *FreeShippingPromotionDataRelationships) SetCouponCodesPromotionRule(v BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule) {
	o.CouponCodesPromotionRule = &v
}

// GetCustomPromotionRule returns the CustomPromotionRule field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetCustomPromotionRule() BuyXPayYPromotionDataRelationshipsCustomPromotionRule {
	if o == nil || IsNil(o.CustomPromotionRule) {
		var ret BuyXPayYPromotionDataRelationshipsCustomPromotionRule
		return ret
	}
	return *o.CustomPromotionRule
}

// GetCustomPromotionRuleOk returns a tuple with the CustomPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetCustomPromotionRuleOk() (*BuyXPayYPromotionDataRelationshipsCustomPromotionRule, bool) {
	if o == nil || IsNil(o.CustomPromotionRule) {
		return nil, false
	}
	return o.CustomPromotionRule, true
}

// HasCustomPromotionRule returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasCustomPromotionRule() bool {
	if o != nil && !IsNil(o.CustomPromotionRule) {
		return true
	}

	return false
}

// SetCustomPromotionRule gets a reference to the given BuyXPayYPromotionDataRelationshipsCustomPromotionRule and assigns it to the CustomPromotionRule field.
func (o *FreeShippingPromotionDataRelationships) SetCustomPromotionRule(v BuyXPayYPromotionDataRelationshipsCustomPromotionRule) {
	o.CustomPromotionRule = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetSkuList() BundleDataRelationshipsSkuList {
	if o == nil || IsNil(o.SkuList) {
		var ret BundleDataRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetSkuListOk() (*BundleDataRelationshipsSkuList, bool) {
	if o == nil || IsNil(o.SkuList) {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasSkuList() bool {
	if o != nil && !IsNil(o.SkuList) {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given BundleDataRelationshipsSkuList and assigns it to the SkuList field.
func (o *FreeShippingPromotionDataRelationships) SetSkuList(v BundleDataRelationshipsSkuList) {
	o.SkuList = &v
}

// GetCoupons returns the Coupons field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetCoupons() BuyXPayYPromotionDataRelationshipsCoupons {
	if o == nil || IsNil(o.Coupons) {
		var ret BuyXPayYPromotionDataRelationshipsCoupons
		return ret
	}
	return *o.Coupons
}

// GetCouponsOk returns a tuple with the Coupons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetCouponsOk() (*BuyXPayYPromotionDataRelationshipsCoupons, bool) {
	if o == nil || IsNil(o.Coupons) {
		return nil, false
	}
	return o.Coupons, true
}

// HasCoupons returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasCoupons() bool {
	if o != nil && !IsNil(o.Coupons) {
		return true
	}

	return false
}

// SetCoupons gets a reference to the given BuyXPayYPromotionDataRelationshipsCoupons and assigns it to the Coupons field.
func (o *FreeShippingPromotionDataRelationships) SetCoupons(v BuyXPayYPromotionDataRelationshipsCoupons) {
	o.Coupons = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *FreeShippingPromotionDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *FreeShippingPromotionDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetTags() AddressDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressDataRelationshipsTags and assigns it to the Tags field.
func (o *FreeShippingPromotionDataRelationships) SetTags(v AddressDataRelationshipsTags) {
	o.Tags = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *FreeShippingPromotionDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *FreeShippingPromotionDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *FreeShippingPromotionDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o FreeShippingPromotionDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreeShippingPromotionDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.PromotionRules) {
		toSerialize["promotion_rules"] = o.PromotionRules
	}
	if !IsNil(o.OrderAmountPromotionRule) {
		toSerialize["order_amount_promotion_rule"] = o.OrderAmountPromotionRule
	}
	if !IsNil(o.SkuListPromotionRule) {
		toSerialize["sku_list_promotion_rule"] = o.SkuListPromotionRule
	}
	if !IsNil(o.CouponCodesPromotionRule) {
		toSerialize["coupon_codes_promotion_rule"] = o.CouponCodesPromotionRule
	}
	if !IsNil(o.CustomPromotionRule) {
		toSerialize["custom_promotion_rule"] = o.CustomPromotionRule
	}
	if !IsNil(o.SkuList) {
		toSerialize["sku_list"] = o.SkuList
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

type NullableFreeShippingPromotionDataRelationships struct {
	value *FreeShippingPromotionDataRelationships
	isSet bool
}

func (v NullableFreeShippingPromotionDataRelationships) Get() *FreeShippingPromotionDataRelationships {
	return v.value
}

func (v *NullableFreeShippingPromotionDataRelationships) Set(val *FreeShippingPromotionDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeShippingPromotionDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeShippingPromotionDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeShippingPromotionDataRelationships(val *FreeShippingPromotionDataRelationships) *NullableFreeShippingPromotionDataRelationships {
	return &NullableFreeShippingPromotionDataRelationships{value: val, isSet: true}
}

func (v NullableFreeShippingPromotionDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeShippingPromotionDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
