/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ShippingMethodTierDataAttributes struct for ShippingMethodTierDataAttributes
type ShippingMethodTierDataAttributes struct {
	// The shipping method tier's name
	Name *string `json:"name,omitempty"`
	// The tier upper limit. When 'null' it means infinity (useful to have an always matching tier).
	UpTo *float32 `json:"up_to,omitempty"`
	// The price of this shipping method tier, in cents.
	PriceAmountCents *int32 `json:"price_amount_cents,omitempty"`
	// The price of this shipping method tier, float.
	PriceAmountFloat *float32 `json:"price_amount_float,omitempty"`
	// The price of this shipping method tier, formatted.
	FormattedPriceAmount *string `json:"formatted_price_amount,omitempty"`
	// Unique identifier for the resource (hash).
	Id *string `json:"id,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewShippingMethodTierDataAttributes instantiates a new ShippingMethodTierDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodTierDataAttributes() *ShippingMethodTierDataAttributes {
	this := ShippingMethodTierDataAttributes{}
	return &this
}

// NewShippingMethodTierDataAttributesWithDefaults instantiates a new ShippingMethodTierDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodTierDataAttributesWithDefaults() *ShippingMethodTierDataAttributes {
	this := ShippingMethodTierDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShippingMethodTierDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShippingMethodTierDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShippingMethodTierDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetUpTo returns the UpTo field value if set, zero value otherwise.
func (o *ShippingMethodTierDataAttributes) GetUpTo() float32 {
	if o == nil || o.UpTo == nil {
		var ret float32
		return ret
	}
	return *o.UpTo
}

// GetUpToOk returns a tuple with the UpTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierDataAttributes) GetUpToOk() (*float32, bool) {
	if o == nil || o.UpTo == nil {
		return nil, false
	}
	return o.UpTo, true
}

// HasUpTo returns a boolean if a field has been set.
func (o *ShippingMethodTierDataAttributes) HasUpTo() bool {
	if o != nil && o.UpTo != nil {
		return true
	}

	return false
}

// SetUpTo gets a reference to the given float32 and assigns it to the UpTo field.
func (o *ShippingMethodTierDataAttributes) SetUpTo(v float32) {
	o.UpTo = &v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise.
func (o *ShippingMethodTierDataAttributes) GetPriceAmountCents() int32 {
	if o == nil || o.PriceAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierDataAttributes) GetPriceAmountCentsOk() (*int32, bool) {
	if o == nil || o.PriceAmountCents == nil {
		return nil, false
	}
	return o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *ShippingMethodTierDataAttributes) HasPriceAmountCents() bool {
	if o != nil && o.PriceAmountCents != nil {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given int32 and assigns it to the PriceAmountCents field.
func (o *ShippingMethodTierDataAttributes) SetPriceAmountCents(v int32) {
	o.PriceAmountCents = &v
}

// GetPriceAmountFloat returns the PriceAmountFloat field value if set, zero value otherwise.
func (o *ShippingMethodTierDataAttributes) GetPriceAmountFloat() float32 {
	if o == nil || o.PriceAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.PriceAmountFloat
}

// GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierDataAttributes) GetPriceAmountFloatOk() (*float32, bool) {
	if o == nil || o.PriceAmountFloat == nil {
		return nil, false
	}
	return o.PriceAmountFloat, true
}

// HasPriceAmountFloat returns a boolean if a field has been set.
func (o *ShippingMethodTierDataAttributes) HasPriceAmountFloat() bool {
	if o != nil && o.PriceAmountFloat != nil {
		return true
	}

	return false
}

// SetPriceAmountFloat gets a reference to the given float32 and assigns it to the PriceAmountFloat field.
func (o *ShippingMethodTierDataAttributes) SetPriceAmountFloat(v float32) {
	o.PriceAmountFloat = &v
}

// GetFormattedPriceAmount returns the FormattedPriceAmount field value if set, zero value otherwise.
func (o *ShippingMethodTierDataAttributes) GetFormattedPriceAmount() string {
	if o == nil || o.FormattedPriceAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedPriceAmount
}

// GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierDataAttributes) GetFormattedPriceAmountOk() (*string, bool) {
	if o == nil || o.FormattedPriceAmount == nil {
		return nil, false
	}
	return o.FormattedPriceAmount, true
}

// HasFormattedPriceAmount returns a boolean if a field has been set.
func (o *ShippingMethodTierDataAttributes) HasFormattedPriceAmount() bool {
	if o != nil && o.FormattedPriceAmount != nil {
		return true
	}

	return false
}

// SetFormattedPriceAmount gets a reference to the given string and assigns it to the FormattedPriceAmount field.
func (o *ShippingMethodTierDataAttributes) SetFormattedPriceAmount(v string) {
	o.FormattedPriceAmount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShippingMethodTierDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShippingMethodTierDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ShippingMethodTierDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ShippingMethodTierDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ShippingMethodTierDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ShippingMethodTierDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ShippingMethodTierDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ShippingMethodTierDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ShippingMethodTierDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ShippingMethodTierDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ShippingMethodTierDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ShippingMethodTierDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *ShippingMethodTierDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *ShippingMethodTierDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *ShippingMethodTierDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ShippingMethodTierDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ShippingMethodTierDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ShippingMethodTierDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ShippingMethodTierDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UpTo != nil {
		toSerialize["up_to"] = o.UpTo
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if o.PriceAmountFloat != nil {
		toSerialize["price_amount_float"] = o.PriceAmountFloat
	}
	if o.FormattedPriceAmount != nil {
		toSerialize["formatted_price_amount"] = o.FormattedPriceAmount
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	return json.Marshal(toSerialize)
}

type NullableShippingMethodTierDataAttributes struct {
	value *ShippingMethodTierDataAttributes
	isSet bool
}

func (v NullableShippingMethodTierDataAttributes) Get() *ShippingMethodTierDataAttributes {
	return v.value
}

func (v *NullableShippingMethodTierDataAttributes) Set(val *ShippingMethodTierDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodTierDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodTierDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodTierDataAttributes(val *ShippingMethodTierDataAttributes) *NullableShippingMethodTierDataAttributes {
	return &NullableShippingMethodTierDataAttributes{value: val, isSet: true}
}

func (v NullableShippingMethodTierDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodTierDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
