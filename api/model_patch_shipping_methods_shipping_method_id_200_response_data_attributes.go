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

// checks if the PATCHShippingMethodsShippingMethodId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHShippingMethodsShippingMethodId200ResponseDataAttributes{}

// PATCHShippingMethodsShippingMethodId200ResponseDataAttributes struct for PATCHShippingMethodsShippingMethodId200ResponseDataAttributes
type PATCHShippingMethodsShippingMethodId200ResponseDataAttributes struct {
	// The shipping method's name.
	Name interface{} `json:"name,omitempty"`
	// The shipping method's scheme. One of 'flat', 'weight_tiered', or 'external'.
	Scheme interface{} `json:"scheme,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// The URL used to overwrite prices by an external source.
	ExternalPricesUrl interface{} `json:"external_prices_url,omitempty"`
	// The price of this shipping method, in cents.
	PriceAmountCents interface{} `json:"price_amount_cents,omitempty"`
	// Apply free shipping if the order amount is over this value, in cents.
	FreeOverAmountCents interface{} `json:"free_over_amount_cents,omitempty"`
	// Send this attribute if you want to compare the free over amount with order's subtotal (excluding discounts, if any).
	UseSubtotal interface{} `json:"use_subtotal,omitempty"`
	// The minimum weight for which this shipping method is available.
	MinWeight interface{} `json:"min_weight,omitempty"`
	// The maximum weight for which this shipping method is available.
	MaxWeight interface{} `json:"max_weight,omitempty"`
	// The unit of weight. One of 'gr', 'oz', or 'lb'.
	UnitOfWeight interface{} `json:"unit_of_weight,omitempty"`
	// Send this attribute if you want to mark this resource as disabled.
	Disable interface{} `json:"_disable,omitempty"`
	// Send this attribute if you want to mark this resource as enabled.
	Enable interface{} `json:"_enable,omitempty"`
	// Send this attribute if you want to reset the circuit breaker associated to this resource to 'closed' state and zero failures count. Cannot be passed by sales channels.
	ResetCircuit interface{} `json:"_reset_circuit,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHShippingMethodsShippingMethodId200ResponseDataAttributes instantiates a new PATCHShippingMethodsShippingMethodId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHShippingMethodsShippingMethodId200ResponseDataAttributes() *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes {
	this := PATCHShippingMethodsShippingMethodId200ResponseDataAttributes{}
	return &this
}

// NewPATCHShippingMethodsShippingMethodId200ResponseDataAttributesWithDefaults instantiates a new PATCHShippingMethodsShippingMethodId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHShippingMethodsShippingMethodId200ResponseDataAttributesWithDefaults() *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes {
	this := PATCHShippingMethodsShippingMethodId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetScheme returns the Scheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetScheme() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetSchemeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return &o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasScheme() bool {
	if o != nil && IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given interface{} and assigns it to the Scheme field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetScheme(v interface{}) {
	o.Scheme = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetExternalPricesUrl returns the ExternalPricesUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetExternalPricesUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExternalPricesUrl
}

// GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetExternalPricesUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExternalPricesUrl) {
		return nil, false
	}
	return &o.ExternalPricesUrl, true
}

// HasExternalPricesUrl returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasExternalPricesUrl() bool {
	if o != nil && IsNil(o.ExternalPricesUrl) {
		return true
	}

	return false
}

// SetExternalPricesUrl gets a reference to the given interface{} and assigns it to the ExternalPricesUrl field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetExternalPricesUrl(v interface{}) {
	o.ExternalPricesUrl = v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PriceAmountCents) {
		return nil, false
	}
	return &o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasPriceAmountCents() bool {
	if o != nil && IsNil(o.PriceAmountCents) {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given interface{} and assigns it to the PriceAmountCents field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountCents(v interface{}) {
	o.PriceAmountCents = v
}

// GetFreeOverAmountCents returns the FreeOverAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetFreeOverAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FreeOverAmountCents
}

// GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetFreeOverAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FreeOverAmountCents) {
		return nil, false
	}
	return &o.FreeOverAmountCents, true
}

// HasFreeOverAmountCents returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasFreeOverAmountCents() bool {
	if o != nil && IsNil(o.FreeOverAmountCents) {
		return true
	}

	return false
}

// SetFreeOverAmountCents gets a reference to the given interface{} and assigns it to the FreeOverAmountCents field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetFreeOverAmountCents(v interface{}) {
	o.FreeOverAmountCents = v
}

// GetUseSubtotal returns the UseSubtotal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetUseSubtotal() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UseSubtotal
}

// GetUseSubtotalOk returns a tuple with the UseSubtotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetUseSubtotalOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UseSubtotal) {
		return nil, false
	}
	return &o.UseSubtotal, true
}

// HasUseSubtotal returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasUseSubtotal() bool {
	if o != nil && IsNil(o.UseSubtotal) {
		return true
	}

	return false
}

// SetUseSubtotal gets a reference to the given interface{} and assigns it to the UseSubtotal field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetUseSubtotal(v interface{}) {
	o.UseSubtotal = v
}

// GetMinWeight returns the MinWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMinWeight() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MinWeight
}

// GetMinWeightOk returns a tuple with the MinWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMinWeightOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MinWeight) {
		return nil, false
	}
	return &o.MinWeight, true
}

// HasMinWeight returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasMinWeight() bool {
	if o != nil && IsNil(o.MinWeight) {
		return true
	}

	return false
}

// SetMinWeight gets a reference to the given interface{} and assigns it to the MinWeight field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMinWeight(v interface{}) {
	o.MinWeight = v
}

// GetMaxWeight returns the MaxWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMaxWeight() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MaxWeight
}

// GetMaxWeightOk returns a tuple with the MaxWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMaxWeightOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MaxWeight) {
		return nil, false
	}
	return &o.MaxWeight, true
}

// HasMaxWeight returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasMaxWeight() bool {
	if o != nil && IsNil(o.MaxWeight) {
		return true
	}

	return false
}

// SetMaxWeight gets a reference to the given interface{} and assigns it to the MaxWeight field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMaxWeight(v interface{}) {
	o.MaxWeight = v
}

// GetUnitOfWeight returns the UnitOfWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetUnitOfWeight() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UnitOfWeight
}

// GetUnitOfWeightOk returns a tuple with the UnitOfWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetUnitOfWeightOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UnitOfWeight) {
		return nil, false
	}
	return &o.UnitOfWeight, true
}

// HasUnitOfWeight returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasUnitOfWeight() bool {
	if o != nil && IsNil(o.UnitOfWeight) {
		return true
	}

	return false
}

// SetUnitOfWeight gets a reference to the given interface{} and assigns it to the UnitOfWeight field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetUnitOfWeight(v interface{}) {
	o.UnitOfWeight = v
}

// GetDisable returns the Disable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetDisable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetDisableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return &o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasDisable() bool {
	if o != nil && IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given interface{} and assigns it to the Disable field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetDisable(v interface{}) {
	o.Disable = v
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetEnable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetEnableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return &o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasEnable() bool {
	if o != nil && IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given interface{} and assigns it to the Enable field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetEnable(v interface{}) {
	o.Enable = v
}

// GetResetCircuit returns the ResetCircuit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetResetCircuit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResetCircuit
}

// GetResetCircuitOk returns a tuple with the ResetCircuit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetResetCircuitOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResetCircuit) {
		return nil, false
	}
	return &o.ResetCircuit, true
}

// HasResetCircuit returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasResetCircuit() bool {
	if o != nil && IsNil(o.ResetCircuit) {
		return true
	}

	return false
}

// SetResetCircuit gets a reference to the given interface{} and assigns it to the ResetCircuit field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetResetCircuit(v interface{}) {
	o.ResetCircuit = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Scheme != nil {
		toSerialize["scheme"] = o.Scheme
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.ExternalPricesUrl != nil {
		toSerialize["external_prices_url"] = o.ExternalPricesUrl
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if o.FreeOverAmountCents != nil {
		toSerialize["free_over_amount_cents"] = o.FreeOverAmountCents
	}
	if o.UseSubtotal != nil {
		toSerialize["use_subtotal"] = o.UseSubtotal
	}
	if o.MinWeight != nil {
		toSerialize["min_weight"] = o.MinWeight
	}
	if o.MaxWeight != nil {
		toSerialize["max_weight"] = o.MaxWeight
	}
	if o.UnitOfWeight != nil {
		toSerialize["unit_of_weight"] = o.UnitOfWeight
	}
	if o.Disable != nil {
		toSerialize["_disable"] = o.Disable
	}
	if o.Enable != nil {
		toSerialize["_enable"] = o.Enable
	}
	if o.ResetCircuit != nil {
		toSerialize["_reset_circuit"] = o.ResetCircuit
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

type NullablePATCHShippingMethodsShippingMethodId200ResponseDataAttributes struct {
	value *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHShippingMethodsShippingMethodId200ResponseDataAttributes) Get() *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHShippingMethodsShippingMethodId200ResponseDataAttributes) Set(val *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHShippingMethodsShippingMethodId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHShippingMethodsShippingMethodId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHShippingMethodsShippingMethodId200ResponseDataAttributes(val *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) *NullablePATCHShippingMethodsShippingMethodId200ResponseDataAttributes {
	return &NullablePATCHShippingMethodsShippingMethodId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHShippingMethodsShippingMethodId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
