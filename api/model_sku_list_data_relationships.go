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

// checks if the SkuListDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuListDataRelationships{}

// SkuListDataRelationships struct for SkuListDataRelationships
type SkuListDataRelationships struct {
	Customer     *CouponRecipientDataRelationshipsCustomer  `json:"customer,omitempty"`
	Skus         *BundleDataRelationshipsSkus               `json:"skus,omitempty"`
	SkuListItems *SkuListDataRelationshipsSkuListItems      `json:"sku_list_items,omitempty"`
	Bundles      *LineItemDataRelationshipsBundle           `json:"bundles,omitempty"`
	Attachments  *AuthorizationDataRelationshipsAttachments `json:"attachments,omitempty"`
	Links        *OrderDataRelationshipsLinks               `json:"links,omitempty"`
	Versions     *AddressDataRelationshipsVersions          `json:"versions,omitempty"`
}

// NewSkuListDataRelationships instantiates a new SkuListDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListDataRelationships() *SkuListDataRelationships {
	this := SkuListDataRelationships{}
	return &this
}

// NewSkuListDataRelationshipsWithDefaults instantiates a new SkuListDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListDataRelationshipsWithDefaults() *SkuListDataRelationships {
	this := SkuListDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *SkuListDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *SkuListDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *SkuListDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *SkuListDataRelationships) GetSkus() BundleDataRelationshipsSkus {
	if o == nil || IsNil(o.Skus) {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListDataRelationships) GetSkusOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || IsNil(o.Skus) {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *SkuListDataRelationships) HasSkus() bool {
	if o != nil && !IsNil(o.Skus) {
		return true
	}

	return false
}

// SetSkus gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Skus field.
func (o *SkuListDataRelationships) SetSkus(v BundleDataRelationshipsSkus) {
	o.Skus = &v
}

// GetSkuListItems returns the SkuListItems field value if set, zero value otherwise.
func (o *SkuListDataRelationships) GetSkuListItems() SkuListDataRelationshipsSkuListItems {
	if o == nil || IsNil(o.SkuListItems) {
		var ret SkuListDataRelationshipsSkuListItems
		return ret
	}
	return *o.SkuListItems
}

// GetSkuListItemsOk returns a tuple with the SkuListItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListDataRelationships) GetSkuListItemsOk() (*SkuListDataRelationshipsSkuListItems, bool) {
	if o == nil || IsNil(o.SkuListItems) {
		return nil, false
	}
	return o.SkuListItems, true
}

// HasSkuListItems returns a boolean if a field has been set.
func (o *SkuListDataRelationships) HasSkuListItems() bool {
	if o != nil && !IsNil(o.SkuListItems) {
		return true
	}

	return false
}

// SetSkuListItems gets a reference to the given SkuListDataRelationshipsSkuListItems and assigns it to the SkuListItems field.
func (o *SkuListDataRelationships) SetSkuListItems(v SkuListDataRelationshipsSkuListItems) {
	o.SkuListItems = &v
}

// GetBundles returns the Bundles field value if set, zero value otherwise.
func (o *SkuListDataRelationships) GetBundles() LineItemDataRelationshipsBundle {
	if o == nil || IsNil(o.Bundles) {
		var ret LineItemDataRelationshipsBundle
		return ret
	}
	return *o.Bundles
}

// GetBundlesOk returns a tuple with the Bundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListDataRelationships) GetBundlesOk() (*LineItemDataRelationshipsBundle, bool) {
	if o == nil || IsNil(o.Bundles) {
		return nil, false
	}
	return o.Bundles, true
}

// HasBundles returns a boolean if a field has been set.
func (o *SkuListDataRelationships) HasBundles() bool {
	if o != nil && !IsNil(o.Bundles) {
		return true
	}

	return false
}

// SetBundles gets a reference to the given LineItemDataRelationshipsBundle and assigns it to the Bundles field.
func (o *SkuListDataRelationships) SetBundles(v LineItemDataRelationshipsBundle) {
	o.Bundles = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *SkuListDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *SkuListDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *SkuListDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SkuListDataRelationships) GetLinks() OrderDataRelationshipsLinks {
	if o == nil || IsNil(o.Links) {
		var ret OrderDataRelationshipsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListDataRelationships) GetLinksOk() (*OrderDataRelationshipsLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SkuListDataRelationships) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrderDataRelationshipsLinks and assigns it to the Links field.
func (o *SkuListDataRelationships) SetLinks(v OrderDataRelationshipsLinks) {
	o.Links = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *SkuListDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *SkuListDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *SkuListDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o SkuListDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuListDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Skus) {
		toSerialize["skus"] = o.Skus
	}
	if !IsNil(o.SkuListItems) {
		toSerialize["sku_list_items"] = o.SkuListItems
	}
	if !IsNil(o.Bundles) {
		toSerialize["bundles"] = o.Bundles
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableSkuListDataRelationships struct {
	value *SkuListDataRelationships
	isSet bool
}

func (v NullableSkuListDataRelationships) Get() *SkuListDataRelationships {
	return v.value
}

func (v *NullableSkuListDataRelationships) Set(val *SkuListDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListDataRelationships(val *SkuListDataRelationships) *NullableSkuListDataRelationships {
	return &NullableSkuListDataRelationships{value: val, isSet: true}
}

func (v NullableSkuListDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
