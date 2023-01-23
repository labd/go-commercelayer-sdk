/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETShipments200ResponseDataInnerAttributes struct for GETShipments200ResponseDataInnerAttributes
type GETShipments200ResponseDataInnerAttributes struct {
	// Unique identifier for the shipment
	Number *string `json:"number,omitempty"`
	// The shipment status, one of 'draft', 'upcoming', 'cancelled', 'on_hold', 'picking', 'packing', 'ready_to_ship', or 'shipped'
	Status *string `json:"status,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the associated order.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The cost of this shipment from the selected carrier account, in cents.
	CostAmountCents *int32 `json:"cost_amount_cents,omitempty"`
	// The cost of this shipment from the selected carrier account, float.
	CostAmountFloat *float32 `json:"cost_amount_float,omitempty"`
	// The cost of this shipment from the selected carrier account, formatted.
	FormattedCostAmount *string `json:"formatted_cost_amount,omitempty"`
	// The total number of SKUs in the shipment's line items. This can be useful to display a preview of the shipment content.
	SkusCount *int32 `json:"skus_count,omitempty"`
	// The selected purchase rate from the available shipping rates.
	SelectedRateId *string `json:"selected_rate_id,omitempty"`
	// The available shipping rates.
	Rates []map[string]interface{} `json:"rates,omitempty"`
	// The shipping rate purchase error code, if any.
	PurchaseErrorCode *string `json:"purchase_error_code,omitempty"`
	// The shipping rate purchase error message, if any.
	PurchaseErrorMessage *string `json:"purchase_error_message,omitempty"`
	// Any errors collected when fetching shipping rates.
	GetRatesErrors []map[string]interface{} `json:"get_rates_errors,omitempty"`
	// Time at which the getting of the shipping rates started.
	GetRatesStartedAt *string `json:"get_rates_started_at,omitempty"`
	// Time at which the getting of the shipping rates completed.
	GetRatesCompletedAt *string `json:"get_rates_completed_at,omitempty"`
	// Time at which the purchasing of the shipping rate started.
	PurchaseStartedAt *string `json:"purchase_started_at,omitempty"`
	// Time at which the purchasing of the shipping rate completed.
	PurchaseCompletedAt *string `json:"purchase_completed_at,omitempty"`
	// Time at which the purchasing of the shipping rate failed.
	PurchaseFailedAt *string `json:"purchase_failed_at,omitempty"`
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

// NewGETShipments200ResponseDataInnerAttributes instantiates a new GETShipments200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShipments200ResponseDataInnerAttributes() *GETShipments200ResponseDataInnerAttributes {
	this := GETShipments200ResponseDataInnerAttributes{}
	return &this
}

// NewGETShipments200ResponseDataInnerAttributesWithDefaults instantiates a new GETShipments200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShipments200ResponseDataInnerAttributesWithDefaults() *GETShipments200ResponseDataInnerAttributes {
	this := GETShipments200ResponseDataInnerAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *GETShipments200ResponseDataInnerAttributes) SetNumber(v string) {
	o.Number = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GETShipments200ResponseDataInnerAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GETShipments200ResponseDataInnerAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCostAmountCents returns the CostAmountCents field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountCents() int32 {
	if o == nil || o.CostAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.CostAmountCents
}

// GetCostAmountCentsOk returns a tuple with the CostAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountCentsOk() (*int32, bool) {
	if o == nil || o.CostAmountCents == nil {
		return nil, false
	}
	return o.CostAmountCents, true
}

// HasCostAmountCents returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasCostAmountCents() bool {
	if o != nil && o.CostAmountCents != nil {
		return true
	}

	return false
}

// SetCostAmountCents gets a reference to the given int32 and assigns it to the CostAmountCents field.
func (o *GETShipments200ResponseDataInnerAttributes) SetCostAmountCents(v int32) {
	o.CostAmountCents = &v
}

// GetCostAmountFloat returns the CostAmountFloat field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountFloat() float32 {
	if o == nil || o.CostAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.CostAmountFloat
}

// GetCostAmountFloatOk returns a tuple with the CostAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountFloatOk() (*float32, bool) {
	if o == nil || o.CostAmountFloat == nil {
		return nil, false
	}
	return o.CostAmountFloat, true
}

// HasCostAmountFloat returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasCostAmountFloat() bool {
	if o != nil && o.CostAmountFloat != nil {
		return true
	}

	return false
}

// SetCostAmountFloat gets a reference to the given float32 and assigns it to the CostAmountFloat field.
func (o *GETShipments200ResponseDataInnerAttributes) SetCostAmountFloat(v float32) {
	o.CostAmountFloat = &v
}

// GetFormattedCostAmount returns the FormattedCostAmount field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetFormattedCostAmount() string {
	if o == nil || o.FormattedCostAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedCostAmount
}

// GetFormattedCostAmountOk returns a tuple with the FormattedCostAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetFormattedCostAmountOk() (*string, bool) {
	if o == nil || o.FormattedCostAmount == nil {
		return nil, false
	}
	return o.FormattedCostAmount, true
}

// HasFormattedCostAmount returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasFormattedCostAmount() bool {
	if o != nil && o.FormattedCostAmount != nil {
		return true
	}

	return false
}

// SetFormattedCostAmount gets a reference to the given string and assigns it to the FormattedCostAmount field.
func (o *GETShipments200ResponseDataInnerAttributes) SetFormattedCostAmount(v string) {
	o.FormattedCostAmount = &v
}

// GetSkusCount returns the SkusCount field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetSkusCount() int32 {
	if o == nil || o.SkusCount == nil {
		var ret int32
		return ret
	}
	return *o.SkusCount
}

// GetSkusCountOk returns a tuple with the SkusCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetSkusCountOk() (*int32, bool) {
	if o == nil || o.SkusCount == nil {
		return nil, false
	}
	return o.SkusCount, true
}

// HasSkusCount returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasSkusCount() bool {
	if o != nil && o.SkusCount != nil {
		return true
	}

	return false
}

// SetSkusCount gets a reference to the given int32 and assigns it to the SkusCount field.
func (o *GETShipments200ResponseDataInnerAttributes) SetSkusCount(v int32) {
	o.SkusCount = &v
}

// GetSelectedRateId returns the SelectedRateId field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetSelectedRateId() string {
	if o == nil || o.SelectedRateId == nil {
		var ret string
		return ret
	}
	return *o.SelectedRateId
}

// GetSelectedRateIdOk returns a tuple with the SelectedRateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetSelectedRateIdOk() (*string, bool) {
	if o == nil || o.SelectedRateId == nil {
		return nil, false
	}
	return o.SelectedRateId, true
}

// HasSelectedRateId returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasSelectedRateId() bool {
	if o != nil && o.SelectedRateId != nil {
		return true
	}

	return false
}

// SetSelectedRateId gets a reference to the given string and assigns it to the SelectedRateId field.
func (o *GETShipments200ResponseDataInnerAttributes) SetSelectedRateId(v string) {
	o.SelectedRateId = &v
}

// GetRates returns the Rates field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetRates() []map[string]interface{} {
	if o == nil || o.Rates == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetRatesOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Rates == nil {
		return nil, false
	}
	return o.Rates, true
}

// HasRates returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasRates() bool {
	if o != nil && o.Rates != nil {
		return true
	}

	return false
}

// SetRates gets a reference to the given []map[string]interface{} and assigns it to the Rates field.
func (o *GETShipments200ResponseDataInnerAttributes) SetRates(v []map[string]interface{}) {
	o.Rates = v
}

// GetPurchaseErrorCode returns the PurchaseErrorCode field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorCode() string {
	if o == nil || o.PurchaseErrorCode == nil {
		var ret string
		return ret
	}
	return *o.PurchaseErrorCode
}

// GetPurchaseErrorCodeOk returns a tuple with the PurchaseErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorCodeOk() (*string, bool) {
	if o == nil || o.PurchaseErrorCode == nil {
		return nil, false
	}
	return o.PurchaseErrorCode, true
}

// HasPurchaseErrorCode returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseErrorCode() bool {
	if o != nil && o.PurchaseErrorCode != nil {
		return true
	}

	return false
}

// SetPurchaseErrorCode gets a reference to the given string and assigns it to the PurchaseErrorCode field.
func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseErrorCode(v string) {
	o.PurchaseErrorCode = &v
}

// GetPurchaseErrorMessage returns the PurchaseErrorMessage field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorMessage() string {
	if o == nil || o.PurchaseErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.PurchaseErrorMessage
}

// GetPurchaseErrorMessageOk returns a tuple with the PurchaseErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorMessageOk() (*string, bool) {
	if o == nil || o.PurchaseErrorMessage == nil {
		return nil, false
	}
	return o.PurchaseErrorMessage, true
}

// HasPurchaseErrorMessage returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseErrorMessage() bool {
	if o != nil && o.PurchaseErrorMessage != nil {
		return true
	}

	return false
}

// SetPurchaseErrorMessage gets a reference to the given string and assigns it to the PurchaseErrorMessage field.
func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseErrorMessage(v string) {
	o.PurchaseErrorMessage = &v
}

// GetGetRatesErrors returns the GetRatesErrors field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesErrors() []map[string]interface{} {
	if o == nil || o.GetRatesErrors == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.GetRatesErrors
}

// GetGetRatesErrorsOk returns a tuple with the GetRatesErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesErrorsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.GetRatesErrors == nil {
		return nil, false
	}
	return o.GetRatesErrors, true
}

// HasGetRatesErrors returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasGetRatesErrors() bool {
	if o != nil && o.GetRatesErrors != nil {
		return true
	}

	return false
}

// SetGetRatesErrors gets a reference to the given []map[string]interface{} and assigns it to the GetRatesErrors field.
func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesErrors(v []map[string]interface{}) {
	o.GetRatesErrors = v
}

// GetGetRatesStartedAt returns the GetRatesStartedAt field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesStartedAt() string {
	if o == nil || o.GetRatesStartedAt == nil {
		var ret string
		return ret
	}
	return *o.GetRatesStartedAt
}

// GetGetRatesStartedAtOk returns a tuple with the GetRatesStartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesStartedAtOk() (*string, bool) {
	if o == nil || o.GetRatesStartedAt == nil {
		return nil, false
	}
	return o.GetRatesStartedAt, true
}

// HasGetRatesStartedAt returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasGetRatesStartedAt() bool {
	if o != nil && o.GetRatesStartedAt != nil {
		return true
	}

	return false
}

// SetGetRatesStartedAt gets a reference to the given string and assigns it to the GetRatesStartedAt field.
func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesStartedAt(v string) {
	o.GetRatesStartedAt = &v
}

// GetGetRatesCompletedAt returns the GetRatesCompletedAt field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesCompletedAt() string {
	if o == nil || o.GetRatesCompletedAt == nil {
		var ret string
		return ret
	}
	return *o.GetRatesCompletedAt
}

// GetGetRatesCompletedAtOk returns a tuple with the GetRatesCompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesCompletedAtOk() (*string, bool) {
	if o == nil || o.GetRatesCompletedAt == nil {
		return nil, false
	}
	return o.GetRatesCompletedAt, true
}

// HasGetRatesCompletedAt returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasGetRatesCompletedAt() bool {
	if o != nil && o.GetRatesCompletedAt != nil {
		return true
	}

	return false
}

// SetGetRatesCompletedAt gets a reference to the given string and assigns it to the GetRatesCompletedAt field.
func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesCompletedAt(v string) {
	o.GetRatesCompletedAt = &v
}

// GetPurchaseStartedAt returns the PurchaseStartedAt field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseStartedAt() string {
	if o == nil || o.PurchaseStartedAt == nil {
		var ret string
		return ret
	}
	return *o.PurchaseStartedAt
}

// GetPurchaseStartedAtOk returns a tuple with the PurchaseStartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseStartedAtOk() (*string, bool) {
	if o == nil || o.PurchaseStartedAt == nil {
		return nil, false
	}
	return o.PurchaseStartedAt, true
}

// HasPurchaseStartedAt returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseStartedAt() bool {
	if o != nil && o.PurchaseStartedAt != nil {
		return true
	}

	return false
}

// SetPurchaseStartedAt gets a reference to the given string and assigns it to the PurchaseStartedAt field.
func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseStartedAt(v string) {
	o.PurchaseStartedAt = &v
}

// GetPurchaseCompletedAt returns the PurchaseCompletedAt field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseCompletedAt() string {
	if o == nil || o.PurchaseCompletedAt == nil {
		var ret string
		return ret
	}
	return *o.PurchaseCompletedAt
}

// GetPurchaseCompletedAtOk returns a tuple with the PurchaseCompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseCompletedAtOk() (*string, bool) {
	if o == nil || o.PurchaseCompletedAt == nil {
		return nil, false
	}
	return o.PurchaseCompletedAt, true
}

// HasPurchaseCompletedAt returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseCompletedAt() bool {
	if o != nil && o.PurchaseCompletedAt != nil {
		return true
	}

	return false
}

// SetPurchaseCompletedAt gets a reference to the given string and assigns it to the PurchaseCompletedAt field.
func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseCompletedAt(v string) {
	o.PurchaseCompletedAt = &v
}

// GetPurchaseFailedAt returns the PurchaseFailedAt field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseFailedAt() string {
	if o == nil || o.PurchaseFailedAt == nil {
		var ret string
		return ret
	}
	return *o.PurchaseFailedAt
}

// GetPurchaseFailedAtOk returns a tuple with the PurchaseFailedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseFailedAtOk() (*string, bool) {
	if o == nil || o.PurchaseFailedAt == nil {
		return nil, false
	}
	return o.PurchaseFailedAt, true
}

// HasPurchaseFailedAt returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseFailedAt() bool {
	if o != nil && o.PurchaseFailedAt != nil {
		return true
	}

	return false
}

// SetPurchaseFailedAt gets a reference to the given string and assigns it to the PurchaseFailedAt field.
func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseFailedAt(v string) {
	o.PurchaseFailedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETShipments200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETShipments200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETShipments200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETShipments200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETShipments200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETShipments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.CostAmountCents != nil {
		toSerialize["cost_amount_cents"] = o.CostAmountCents
	}
	if o.CostAmountFloat != nil {
		toSerialize["cost_amount_float"] = o.CostAmountFloat
	}
	if o.FormattedCostAmount != nil {
		toSerialize["formatted_cost_amount"] = o.FormattedCostAmount
	}
	if o.SkusCount != nil {
		toSerialize["skus_count"] = o.SkusCount
	}
	if o.SelectedRateId != nil {
		toSerialize["selected_rate_id"] = o.SelectedRateId
	}
	if o.Rates != nil {
		toSerialize["rates"] = o.Rates
	}
	if o.PurchaseErrorCode != nil {
		toSerialize["purchase_error_code"] = o.PurchaseErrorCode
	}
	if o.PurchaseErrorMessage != nil {
		toSerialize["purchase_error_message"] = o.PurchaseErrorMessage
	}
	if o.GetRatesErrors != nil {
		toSerialize["get_rates_errors"] = o.GetRatesErrors
	}
	if o.GetRatesStartedAt != nil {
		toSerialize["get_rates_started_at"] = o.GetRatesStartedAt
	}
	if o.GetRatesCompletedAt != nil {
		toSerialize["get_rates_completed_at"] = o.GetRatesCompletedAt
	}
	if o.PurchaseStartedAt != nil {
		toSerialize["purchase_started_at"] = o.PurchaseStartedAt
	}
	if o.PurchaseCompletedAt != nil {
		toSerialize["purchase_completed_at"] = o.PurchaseCompletedAt
	}
	if o.PurchaseFailedAt != nil {
		toSerialize["purchase_failed_at"] = o.PurchaseFailedAt
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

type NullableGETShipments200ResponseDataInnerAttributes struct {
	value *GETShipments200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETShipments200ResponseDataInnerAttributes) Get() *GETShipments200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETShipments200ResponseDataInnerAttributes) Set(val *GETShipments200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShipments200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShipments200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShipments200ResponseDataInnerAttributes(val *GETShipments200ResponseDataInnerAttributes) *NullableGETShipments200ResponseDataInnerAttributes {
	return &NullableGETShipments200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETShipments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShipments200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
