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

// checks if the POSTMarkets201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarkets201ResponseDataAttributes{}

// POSTMarkets201ResponseDataAttributes struct for POSTMarkets201ResponseDataAttributes
type POSTMarkets201ResponseDataAttributes struct {
	// The market's internal name.
	Name interface{} `json:"name"`
	// A string that you can use to identify the market (must be unique within the environment).
	Code interface{} `json:"code,omitempty"`
	// The Facebook Pixed ID.
	FacebookPixelId interface{} `json:"facebook_pixel_id,omitempty"`
	// The checkout URL for this market.
	CheckoutUrl interface{} `json:"checkout_url,omitempty"`
	// The URL used to overwrite prices by an external source.
	ExternalPricesUrl interface{} `json:"external_prices_url,omitempty"`
	// The URL used to validate orders by an external source.
	ExternalOrderValidationUrl interface{} `json:"external_order_validation_url,omitempty"`
	// When specified indicates the maximum number of shipping line items with cost that will be added to an order.
	ShippingCostCutoff interface{} `json:"shipping_cost_cutoff,omitempty"`
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

// NewPOSTMarkets201ResponseDataAttributes instantiates a new POSTMarkets201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarkets201ResponseDataAttributes(name interface{}) *POSTMarkets201ResponseDataAttributes {
	this := POSTMarkets201ResponseDataAttributes{}
	this.Name = name
	return &this
}

// NewPOSTMarkets201ResponseDataAttributesWithDefaults instantiates a new POSTMarkets201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarkets201ResponseDataAttributesWithDefaults() *POSTMarkets201ResponseDataAttributes {
	this := POSTMarkets201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTMarkets201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *POSTMarkets201ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetFacebookPixelId returns the FacebookPixelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataAttributes) GetFacebookPixelId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FacebookPixelId
}

// GetFacebookPixelIdOk returns a tuple with the FacebookPixelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetFacebookPixelIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FacebookPixelId) {
		return nil, false
	}
	return &o.FacebookPixelId, true
}

// HasFacebookPixelId returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataAttributes) HasFacebookPixelId() bool {
	if o != nil && IsNil(o.FacebookPixelId) {
		return true
	}

	return false
}

// SetFacebookPixelId gets a reference to the given interface{} and assigns it to the FacebookPixelId field.
func (o *POSTMarkets201ResponseDataAttributes) SetFacebookPixelId(v interface{}) {
	o.FacebookPixelId = v
}

// GetCheckoutUrl returns the CheckoutUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataAttributes) GetCheckoutUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CheckoutUrl
}

// GetCheckoutUrlOk returns a tuple with the CheckoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetCheckoutUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CheckoutUrl) {
		return nil, false
	}
	return &o.CheckoutUrl, true
}

// HasCheckoutUrl returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataAttributes) HasCheckoutUrl() bool {
	if o != nil && IsNil(o.CheckoutUrl) {
		return true
	}

	return false
}

// SetCheckoutUrl gets a reference to the given interface{} and assigns it to the CheckoutUrl field.
func (o *POSTMarkets201ResponseDataAttributes) SetCheckoutUrl(v interface{}) {
	o.CheckoutUrl = v
}

// GetExternalPricesUrl returns the ExternalPricesUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataAttributes) GetExternalPricesUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExternalPricesUrl
}

// GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetExternalPricesUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExternalPricesUrl) {
		return nil, false
	}
	return &o.ExternalPricesUrl, true
}

// HasExternalPricesUrl returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataAttributes) HasExternalPricesUrl() bool {
	if o != nil && IsNil(o.ExternalPricesUrl) {
		return true
	}

	return false
}

// SetExternalPricesUrl gets a reference to the given interface{} and assigns it to the ExternalPricesUrl field.
func (o *POSTMarkets201ResponseDataAttributes) SetExternalPricesUrl(v interface{}) {
	o.ExternalPricesUrl = v
}

// GetExternalOrderValidationUrl returns the ExternalOrderValidationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataAttributes) GetExternalOrderValidationUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExternalOrderValidationUrl
}

// GetExternalOrderValidationUrlOk returns a tuple with the ExternalOrderValidationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetExternalOrderValidationUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExternalOrderValidationUrl) {
		return nil, false
	}
	return &o.ExternalOrderValidationUrl, true
}

// HasExternalOrderValidationUrl returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataAttributes) HasExternalOrderValidationUrl() bool {
	if o != nil && IsNil(o.ExternalOrderValidationUrl) {
		return true
	}

	return false
}

// SetExternalOrderValidationUrl gets a reference to the given interface{} and assigns it to the ExternalOrderValidationUrl field.
func (o *POSTMarkets201ResponseDataAttributes) SetExternalOrderValidationUrl(v interface{}) {
	o.ExternalOrderValidationUrl = v
}

// GetShippingCostCutoff returns the ShippingCostCutoff field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataAttributes) GetShippingCostCutoff() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ShippingCostCutoff
}

// GetShippingCostCutoffOk returns a tuple with the ShippingCostCutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetShippingCostCutoffOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ShippingCostCutoff) {
		return nil, false
	}
	return &o.ShippingCostCutoff, true
}

// HasShippingCostCutoff returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataAttributes) HasShippingCostCutoff() bool {
	if o != nil && IsNil(o.ShippingCostCutoff) {
		return true
	}

	return false
}

// SetShippingCostCutoff gets a reference to the given interface{} and assigns it to the ShippingCostCutoff field.
func (o *POSTMarkets201ResponseDataAttributes) SetShippingCostCutoff(v interface{}) {
	o.ShippingCostCutoff = v
}

// GetDisable returns the Disable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataAttributes) GetDisable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetDisableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return &o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataAttributes) HasDisable() bool {
	if o != nil && IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given interface{} and assigns it to the Disable field.
func (o *POSTMarkets201ResponseDataAttributes) SetDisable(v interface{}) {
	o.Disable = v
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataAttributes) GetEnable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetEnableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return &o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataAttributes) HasEnable() bool {
	if o != nil && IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given interface{} and assigns it to the Enable field.
func (o *POSTMarkets201ResponseDataAttributes) SetEnable(v interface{}) {
	o.Enable = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTMarkets201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTMarkets201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTMarkets201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTMarkets201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarkets201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.FacebookPixelId != nil {
		toSerialize["facebook_pixel_id"] = o.FacebookPixelId
	}
	if o.CheckoutUrl != nil {
		toSerialize["checkout_url"] = o.CheckoutUrl
	}
	if o.ExternalPricesUrl != nil {
		toSerialize["external_prices_url"] = o.ExternalPricesUrl
	}
	if o.ExternalOrderValidationUrl != nil {
		toSerialize["external_order_validation_url"] = o.ExternalOrderValidationUrl
	}
	if o.ShippingCostCutoff != nil {
		toSerialize["shipping_cost_cutoff"] = o.ShippingCostCutoff
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

type NullablePOSTMarkets201ResponseDataAttributes struct {
	value *POSTMarkets201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTMarkets201ResponseDataAttributes) Get() *POSTMarkets201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTMarkets201ResponseDataAttributes) Set(val *POSTMarkets201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarkets201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarkets201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarkets201ResponseDataAttributes(val *POSTMarkets201ResponseDataAttributes) *NullablePOSTMarkets201ResponseDataAttributes {
	return &NullablePOSTMarkets201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTMarkets201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarkets201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
