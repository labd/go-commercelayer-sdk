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

// checks if the GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes{}

// GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes struct for GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes
type GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes struct {
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
	// The discount fixed amount to be applied, in cents.
	FixedAmountCents interface{} `json:"fixed_amount_cents,omitempty"`
	// The discount fixed amount to be applied, float.
	FixedAmountFloat interface{} `json:"fixed_amount_float,omitempty"`
	// The discount fixed amount to be applied, formatted.
	FormattedFixedAmount interface{} `json:"formatted_fixed_amount,omitempty"`
}

// NewGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes instantiates a new GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes() *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes {
	this := GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes{}
	return &this
}

// NewGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributesWithDefaults instantiates a new GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributesWithDefaults() *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes {
	this := GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetType(v interface{}) {
	o.Type = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetExclusive returns the Exclusive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetExclusive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Exclusive
}

// GetExclusiveOk returns a tuple with the Exclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetExclusiveOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Exclusive) {
		return nil, false
	}
	return &o.Exclusive, true
}

// HasExclusive returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasExclusive() bool {
	if o != nil && IsNil(o.Exclusive) {
		return true
	}

	return false
}

// SetExclusive gets a reference to the given interface{} and assigns it to the Exclusive field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetExclusive(v interface{}) {
	o.Exclusive = v
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetPriority() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetPriorityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return &o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasPriority() bool {
	if o != nil && IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given interface{} and assigns it to the Priority field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetPriority(v interface{}) {
	o.Priority = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return &o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasStartsAt() bool {
	if o != nil && IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given interface{} and assigns it to the StartsAt field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetTotalUsageLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalUsageLimit) {
		return nil, false
	}
	return &o.TotalUsageLimit, true
}

// HasTotalUsageLimit returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasTotalUsageLimit() bool {
	if o != nil && IsNil(o.TotalUsageLimit) {
		return true
	}

	return false
}

// SetTotalUsageLimit gets a reference to the given interface{} and assigns it to the TotalUsageLimit field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetTotalUsageLimit(v interface{}) {
	o.TotalUsageLimit = v
}

// GetTotalUsageCount returns the TotalUsageCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetTotalUsageCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageCount
}

// GetTotalUsageCountOk returns a tuple with the TotalUsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetTotalUsageCountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalUsageCount) {
		return nil, false
	}
	return &o.TotalUsageCount, true
}

// HasTotalUsageCount returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasTotalUsageCount() bool {
	if o != nil && IsNil(o.TotalUsageCount) {
		return true
	}

	return false
}

// SetTotalUsageCount gets a reference to the given interface{} and assigns it to the TotalUsageCount field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetTotalUsageCount(v interface{}) {
	o.TotalUsageCount = v
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetActive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetActiveOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return &o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasActive() bool {
	if o != nil && IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given interface{} and assigns it to the Active field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetActive(v interface{}) {
	o.Active = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasStatus() bool {
	if o != nil && IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetStatus(v interface{}) {
	o.Status = v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetDisabledAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DisabledAt) {
		return nil, false
	}
	return &o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasDisabledAt() bool {
	if o != nil && IsNil(o.DisabledAt) {
		return true
	}

	return false
}

// SetDisabledAt gets a reference to the given interface{} and assigns it to the DisabledAt field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetDisabledAt(v interface{}) {
	o.DisabledAt = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetFixedAmountCents returns the FixedAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetFixedAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FixedAmountCents
}

// GetFixedAmountCentsOk returns a tuple with the FixedAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetFixedAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FixedAmountCents) {
		return nil, false
	}
	return &o.FixedAmountCents, true
}

// HasFixedAmountCents returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasFixedAmountCents() bool {
	if o != nil && IsNil(o.FixedAmountCents) {
		return true
	}

	return false
}

// SetFixedAmountCents gets a reference to the given interface{} and assigns it to the FixedAmountCents field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetFixedAmountCents(v interface{}) {
	o.FixedAmountCents = v
}

// GetFixedAmountFloat returns the FixedAmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetFixedAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FixedAmountFloat
}

// GetFixedAmountFloatOk returns a tuple with the FixedAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetFixedAmountFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FixedAmountFloat) {
		return nil, false
	}
	return &o.FixedAmountFloat, true
}

// HasFixedAmountFloat returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasFixedAmountFloat() bool {
	if o != nil && IsNil(o.FixedAmountFloat) {
		return true
	}

	return false
}

// SetFixedAmountFloat gets a reference to the given interface{} and assigns it to the FixedAmountFloat field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetFixedAmountFloat(v interface{}) {
	o.FixedAmountFloat = v
}

// GetFormattedFixedAmount returns the FormattedFixedAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetFormattedFixedAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedFixedAmount
}

// GetFormattedFixedAmountOk returns a tuple with the FormattedFixedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetFormattedFixedAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedFixedAmount) {
		return nil, false
	}
	return &o.FormattedFixedAmount, true
}

// HasFormattedFixedAmount returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasFormattedFixedAmount() bool {
	if o != nil && IsNil(o.FormattedFixedAmount) {
		return true
	}

	return false
}

// SetFormattedFixedAmount gets a reference to the given interface{} and assigns it to the FormattedFixedAmount field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetFormattedFixedAmount(v interface{}) {
	o.FormattedFixedAmount = v
}

func (o GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.FixedAmountCents != nil {
		toSerialize["fixed_amount_cents"] = o.FixedAmountCents
	}
	if o.FixedAmountFloat != nil {
		toSerialize["fixed_amount_float"] = o.FixedAmountFloat
	}
	if o.FormattedFixedAmount != nil {
		toSerialize["formatted_fixed_amount"] = o.FormattedFixedAmount
	}
	return toSerialize, nil
}

type NullableGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes struct {
	value *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) Get() *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) Set(val *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes(val *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) *NullableGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes {
	return &NullableGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
