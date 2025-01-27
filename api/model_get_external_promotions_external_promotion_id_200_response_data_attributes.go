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

// checks if the GETExternalPromotionsExternalPromotionId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETExternalPromotionsExternalPromotionId200ResponseDataAttributes{}

// GETExternalPromotionsExternalPromotionId200ResponseDataAttributes struct for GETExternalPromotionsExternalPromotionId200ResponseDataAttributes
type GETExternalPromotionsExternalPromotionId200ResponseDataAttributes struct {
	// The promotion's internal name.
	Name interface{} `json:"name,omitempty"`
	// The promotion's type.
	Type interface{} `json:"type,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
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
	// The number of times this promotion has been applied.
	TotalUsageCount interface{} `json:"total_usage_count,omitempty"`
	// Indicates if the promotion is active (enabled and not expired).
	Active interface{} `json:"active,omitempty"`
	// The promotion status. One of 'disabled', 'expired', 'pending', 'active', or 'inactive'.
	Status interface{} `json:"status,omitempty"`
	// Time at which this resource was disabled.
	DisabledAt interface{} `json:"disabled_at,omitempty"`
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
	// The URL to the service that will compute the discount.
	PromotionUrl interface{} `json:"promotion_url,omitempty"`
	// The circuit breaker state, by default it is 'closed'. It can become 'open' once the number of consecutive failures overlaps the specified threshold, in such case no further calls to the failing callback are made.
	CircuitState interface{} `json:"circuit_state,omitempty"`
	// The number of consecutive failures recorded by the circuit breaker associated to this resource, will be reset on first successful call to callback.
	CircuitFailureCount interface{} `json:"circuit_failure_count,omitempty"`
	// The shared secret used to sign the external request payload.
	SharedSecret interface{} `json:"shared_secret,omitempty"`
}

// NewGETExternalPromotionsExternalPromotionId200ResponseDataAttributes instantiates a new GETExternalPromotionsExternalPromotionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPromotionsExternalPromotionId200ResponseDataAttributes() *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes {
	this := GETExternalPromotionsExternalPromotionId200ResponseDataAttributes{}
	return &this
}

// NewGETExternalPromotionsExternalPromotionId200ResponseDataAttributesWithDefaults instantiates a new GETExternalPromotionsExternalPromotionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPromotionsExternalPromotionId200ResponseDataAttributesWithDefaults() *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes {
	this := GETExternalPromotionsExternalPromotionId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetType(v interface{}) {
	o.Type = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetExclusive returns the Exclusive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetExclusive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Exclusive
}

// GetExclusiveOk returns a tuple with the Exclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetExclusiveOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Exclusive) {
		return nil, false
	}
	return &o.Exclusive, true
}

// HasExclusive returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasExclusive() bool {
	if o != nil && IsNil(o.Exclusive) {
		return true
	}

	return false
}

// SetExclusive gets a reference to the given interface{} and assigns it to the Exclusive field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetExclusive(v interface{}) {
	o.Exclusive = v
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetPriority() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetPriorityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return &o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasPriority() bool {
	if o != nil && IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given interface{} and assigns it to the Priority field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetPriority(v interface{}) {
	o.Priority = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return &o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasStartsAt() bool {
	if o != nil && IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given interface{} and assigns it to the StartsAt field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetTotalUsageLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalUsageLimit) {
		return nil, false
	}
	return &o.TotalUsageLimit, true
}

// HasTotalUsageLimit returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasTotalUsageLimit() bool {
	if o != nil && IsNil(o.TotalUsageLimit) {
		return true
	}

	return false
}

// SetTotalUsageLimit gets a reference to the given interface{} and assigns it to the TotalUsageLimit field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetTotalUsageLimit(v interface{}) {
	o.TotalUsageLimit = v
}

// GetTotalUsageCount returns the TotalUsageCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetTotalUsageCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageCount
}

// GetTotalUsageCountOk returns a tuple with the TotalUsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetTotalUsageCountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalUsageCount) {
		return nil, false
	}
	return &o.TotalUsageCount, true
}

// HasTotalUsageCount returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasTotalUsageCount() bool {
	if o != nil && IsNil(o.TotalUsageCount) {
		return true
	}

	return false
}

// SetTotalUsageCount gets a reference to the given interface{} and assigns it to the TotalUsageCount field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetTotalUsageCount(v interface{}) {
	o.TotalUsageCount = v
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetActive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetActiveOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return &o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasActive() bool {
	if o != nil && IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given interface{} and assigns it to the Active field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetActive(v interface{}) {
	o.Active = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasStatus() bool {
	if o != nil && IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetStatus(v interface{}) {
	o.Status = v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetDisabledAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DisabledAt) {
		return nil, false
	}
	return &o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasDisabledAt() bool {
	if o != nil && IsNil(o.DisabledAt) {
		return true
	}

	return false
}

// SetDisabledAt gets a reference to the given interface{} and assigns it to the DisabledAt field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetDisabledAt(v interface{}) {
	o.DisabledAt = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetPromotionUrl returns the PromotionUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetPromotionUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PromotionUrl
}

// GetPromotionUrlOk returns a tuple with the PromotionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetPromotionUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PromotionUrl) {
		return nil, false
	}
	return &o.PromotionUrl, true
}

// HasPromotionUrl returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasPromotionUrl() bool {
	if o != nil && IsNil(o.PromotionUrl) {
		return true
	}

	return false
}

// SetPromotionUrl gets a reference to the given interface{} and assigns it to the PromotionUrl field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetPromotionUrl(v interface{}) {
	o.PromotionUrl = v
}

// GetCircuitState returns the CircuitState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetCircuitState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CircuitState
}

// GetCircuitStateOk returns a tuple with the CircuitState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetCircuitStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CircuitState) {
		return nil, false
	}
	return &o.CircuitState, true
}

// HasCircuitState returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasCircuitState() bool {
	if o != nil && IsNil(o.CircuitState) {
		return true
	}

	return false
}

// SetCircuitState gets a reference to the given interface{} and assigns it to the CircuitState field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetCircuitState(v interface{}) {
	o.CircuitState = v
}

// GetCircuitFailureCount returns the CircuitFailureCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetCircuitFailureCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CircuitFailureCount
}

// GetCircuitFailureCountOk returns a tuple with the CircuitFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetCircuitFailureCountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CircuitFailureCount) {
		return nil, false
	}
	return &o.CircuitFailureCount, true
}

// HasCircuitFailureCount returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasCircuitFailureCount() bool {
	if o != nil && IsNil(o.CircuitFailureCount) {
		return true
	}

	return false
}

// SetCircuitFailureCount gets a reference to the given interface{} and assigns it to the CircuitFailureCount field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetCircuitFailureCount(v interface{}) {
	o.CircuitFailureCount = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetSharedSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetSharedSecretOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return &o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasSharedSecret() bool {
	if o != nil && IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given interface{} and assigns it to the SharedSecret field.
func (o *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetSharedSecret(v interface{}) {
	o.SharedSecret = v
}

func (o GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
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
	if o.TotalUsageCount != nil {
		toSerialize["total_usage_count"] = o.TotalUsageCount
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.DisabledAt != nil {
		toSerialize["disabled_at"] = o.DisabledAt
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
	if o.PromotionUrl != nil {
		toSerialize["promotion_url"] = o.PromotionUrl
	}
	if o.CircuitState != nil {
		toSerialize["circuit_state"] = o.CircuitState
	}
	if o.CircuitFailureCount != nil {
		toSerialize["circuit_failure_count"] = o.CircuitFailureCount
	}
	if o.SharedSecret != nil {
		toSerialize["shared_secret"] = o.SharedSecret
	}
	return toSerialize, nil
}

type NullableGETExternalPromotionsExternalPromotionId200ResponseDataAttributes struct {
	value *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETExternalPromotionsExternalPromotionId200ResponseDataAttributes) Get() *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETExternalPromotionsExternalPromotionId200ResponseDataAttributes) Set(val *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPromotionsExternalPromotionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPromotionsExternalPromotionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPromotionsExternalPromotionId200ResponseDataAttributes(val *GETExternalPromotionsExternalPromotionId200ResponseDataAttributes) *NullableGETExternalPromotionsExternalPromotionId200ResponseDataAttributes {
	return &NullableGETExternalPromotionsExternalPromotionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETExternalPromotionsExternalPromotionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPromotionsExternalPromotionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
