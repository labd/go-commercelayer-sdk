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

// GiftCardDataAttributes struct for GiftCardDataAttributes
type GiftCardDataAttributes struct {
	// The gift card status, one of 'draft', 'inactive', 'active', or 'redeemed'.
	Status *string `json:"status,omitempty"`
	// The gift card code UUID. If not set, it's automatically generated.
	Code *string `json:"code,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The gift card initial balance, in cents.
	InitialBalanceCents *int32 `json:"initial_balance_cents,omitempty"`
	// The gift card initial balance, float.
	InitialBalanceFloat *float32 `json:"initial_balance_float,omitempty"`
	// The gift card initial balance, formatted.
	FormattedInitialBalance *string `json:"formatted_initial_balance,omitempty"`
	// The gift card balance, in cents.
	BalanceCents *int32 `json:"balance_cents,omitempty"`
	// The gift card balance, float.
	BalanceFloat *float32 `json:"balance_float,omitempty"`
	// The gift card balance, formatted.
	FormattedBalance *string `json:"formatted_balance,omitempty"`
	// The gift card balance max, in cents.
	BalanceMaxCents *string `json:"balance_max_cents,omitempty"`
	// The gift card balance max, float.
	BalanceMaxFloat *float32 `json:"balance_max_float,omitempty"`
	// The gift card balance max, formatted.
	FormattedBalanceMax *string `json:"formatted_balance_max,omitempty"`
	// The gift card balance log. Tracks all the gift card transactions.
	BalanceLog []map[string]interface{} `json:"balance_log,omitempty"`
	// Indicates if the gift card can be used only one.
	SingleUse *bool `json:"single_use,omitempty"`
	// Indicates if the gift card can be recharged.
	Rechargeable *bool `json:"rechargeable,omitempty"`
	// The URL of an image that represents the gift card.
	ImageUrl *string `json:"image_url,omitempty"`
	// Time at which the gift card will expire.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// The email address of the associated recipient. When creating or updating a gift card, this is a shortcut to find or create the associated recipient by email.
	RecipientEmail *string `json:"recipient_email,omitempty"`
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

// NewGiftCardDataAttributes instantiates a new GiftCardDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardDataAttributes() *GiftCardDataAttributes {
	this := GiftCardDataAttributes{}
	return &this
}

// NewGiftCardDataAttributesWithDefaults instantiates a new GiftCardDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardDataAttributesWithDefaults() *GiftCardDataAttributes {
	this := GiftCardDataAttributes{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GiftCardDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GiftCardDataAttributes) SetCode(v string) {
	o.Code = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GiftCardDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetInitialBalanceCents returns the InitialBalanceCents field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetInitialBalanceCents() int32 {
	if o == nil || o.InitialBalanceCents == nil {
		var ret int32
		return ret
	}
	return *o.InitialBalanceCents
}

// GetInitialBalanceCentsOk returns a tuple with the InitialBalanceCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetInitialBalanceCentsOk() (*int32, bool) {
	if o == nil || o.InitialBalanceCents == nil {
		return nil, false
	}
	return o.InitialBalanceCents, true
}

// HasInitialBalanceCents returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasInitialBalanceCents() bool {
	if o != nil && o.InitialBalanceCents != nil {
		return true
	}

	return false
}

// SetInitialBalanceCents gets a reference to the given int32 and assigns it to the InitialBalanceCents field.
func (o *GiftCardDataAttributes) SetInitialBalanceCents(v int32) {
	o.InitialBalanceCents = &v
}

// GetInitialBalanceFloat returns the InitialBalanceFloat field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetInitialBalanceFloat() float32 {
	if o == nil || o.InitialBalanceFloat == nil {
		var ret float32
		return ret
	}
	return *o.InitialBalanceFloat
}

// GetInitialBalanceFloatOk returns a tuple with the InitialBalanceFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetInitialBalanceFloatOk() (*float32, bool) {
	if o == nil || o.InitialBalanceFloat == nil {
		return nil, false
	}
	return o.InitialBalanceFloat, true
}

// HasInitialBalanceFloat returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasInitialBalanceFloat() bool {
	if o != nil && o.InitialBalanceFloat != nil {
		return true
	}

	return false
}

// SetInitialBalanceFloat gets a reference to the given float32 and assigns it to the InitialBalanceFloat field.
func (o *GiftCardDataAttributes) SetInitialBalanceFloat(v float32) {
	o.InitialBalanceFloat = &v
}

// GetFormattedInitialBalance returns the FormattedInitialBalance field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetFormattedInitialBalance() string {
	if o == nil || o.FormattedInitialBalance == nil {
		var ret string
		return ret
	}
	return *o.FormattedInitialBalance
}

// GetFormattedInitialBalanceOk returns a tuple with the FormattedInitialBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetFormattedInitialBalanceOk() (*string, bool) {
	if o == nil || o.FormattedInitialBalance == nil {
		return nil, false
	}
	return o.FormattedInitialBalance, true
}

// HasFormattedInitialBalance returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasFormattedInitialBalance() bool {
	if o != nil && o.FormattedInitialBalance != nil {
		return true
	}

	return false
}

// SetFormattedInitialBalance gets a reference to the given string and assigns it to the FormattedInitialBalance field.
func (o *GiftCardDataAttributes) SetFormattedInitialBalance(v string) {
	o.FormattedInitialBalance = &v
}

// GetBalanceCents returns the BalanceCents field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetBalanceCents() int32 {
	if o == nil || o.BalanceCents == nil {
		var ret int32
		return ret
	}
	return *o.BalanceCents
}

// GetBalanceCentsOk returns a tuple with the BalanceCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetBalanceCentsOk() (*int32, bool) {
	if o == nil || o.BalanceCents == nil {
		return nil, false
	}
	return o.BalanceCents, true
}

// HasBalanceCents returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasBalanceCents() bool {
	if o != nil && o.BalanceCents != nil {
		return true
	}

	return false
}

// SetBalanceCents gets a reference to the given int32 and assigns it to the BalanceCents field.
func (o *GiftCardDataAttributes) SetBalanceCents(v int32) {
	o.BalanceCents = &v
}

// GetBalanceFloat returns the BalanceFloat field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetBalanceFloat() float32 {
	if o == nil || o.BalanceFloat == nil {
		var ret float32
		return ret
	}
	return *o.BalanceFloat
}

// GetBalanceFloatOk returns a tuple with the BalanceFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetBalanceFloatOk() (*float32, bool) {
	if o == nil || o.BalanceFloat == nil {
		return nil, false
	}
	return o.BalanceFloat, true
}

// HasBalanceFloat returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasBalanceFloat() bool {
	if o != nil && o.BalanceFloat != nil {
		return true
	}

	return false
}

// SetBalanceFloat gets a reference to the given float32 and assigns it to the BalanceFloat field.
func (o *GiftCardDataAttributes) SetBalanceFloat(v float32) {
	o.BalanceFloat = &v
}

// GetFormattedBalance returns the FormattedBalance field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetFormattedBalance() string {
	if o == nil || o.FormattedBalance == nil {
		var ret string
		return ret
	}
	return *o.FormattedBalance
}

// GetFormattedBalanceOk returns a tuple with the FormattedBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetFormattedBalanceOk() (*string, bool) {
	if o == nil || o.FormattedBalance == nil {
		return nil, false
	}
	return o.FormattedBalance, true
}

// HasFormattedBalance returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasFormattedBalance() bool {
	if o != nil && o.FormattedBalance != nil {
		return true
	}

	return false
}

// SetFormattedBalance gets a reference to the given string and assigns it to the FormattedBalance field.
func (o *GiftCardDataAttributes) SetFormattedBalance(v string) {
	o.FormattedBalance = &v
}

// GetBalanceMaxCents returns the BalanceMaxCents field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetBalanceMaxCents() string {
	if o == nil || o.BalanceMaxCents == nil {
		var ret string
		return ret
	}
	return *o.BalanceMaxCents
}

// GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetBalanceMaxCentsOk() (*string, bool) {
	if o == nil || o.BalanceMaxCents == nil {
		return nil, false
	}
	return o.BalanceMaxCents, true
}

// HasBalanceMaxCents returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasBalanceMaxCents() bool {
	if o != nil && o.BalanceMaxCents != nil {
		return true
	}

	return false
}

// SetBalanceMaxCents gets a reference to the given string and assigns it to the BalanceMaxCents field.
func (o *GiftCardDataAttributes) SetBalanceMaxCents(v string) {
	o.BalanceMaxCents = &v
}

// GetBalanceMaxFloat returns the BalanceMaxFloat field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetBalanceMaxFloat() float32 {
	if o == nil || o.BalanceMaxFloat == nil {
		var ret float32
		return ret
	}
	return *o.BalanceMaxFloat
}

// GetBalanceMaxFloatOk returns a tuple with the BalanceMaxFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetBalanceMaxFloatOk() (*float32, bool) {
	if o == nil || o.BalanceMaxFloat == nil {
		return nil, false
	}
	return o.BalanceMaxFloat, true
}

// HasBalanceMaxFloat returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasBalanceMaxFloat() bool {
	if o != nil && o.BalanceMaxFloat != nil {
		return true
	}

	return false
}

// SetBalanceMaxFloat gets a reference to the given float32 and assigns it to the BalanceMaxFloat field.
func (o *GiftCardDataAttributes) SetBalanceMaxFloat(v float32) {
	o.BalanceMaxFloat = &v
}

// GetFormattedBalanceMax returns the FormattedBalanceMax field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetFormattedBalanceMax() string {
	if o == nil || o.FormattedBalanceMax == nil {
		var ret string
		return ret
	}
	return *o.FormattedBalanceMax
}

// GetFormattedBalanceMaxOk returns a tuple with the FormattedBalanceMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetFormattedBalanceMaxOk() (*string, bool) {
	if o == nil || o.FormattedBalanceMax == nil {
		return nil, false
	}
	return o.FormattedBalanceMax, true
}

// HasFormattedBalanceMax returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasFormattedBalanceMax() bool {
	if o != nil && o.FormattedBalanceMax != nil {
		return true
	}

	return false
}

// SetFormattedBalanceMax gets a reference to the given string and assigns it to the FormattedBalanceMax field.
func (o *GiftCardDataAttributes) SetFormattedBalanceMax(v string) {
	o.FormattedBalanceMax = &v
}

// GetBalanceLog returns the BalanceLog field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetBalanceLog() []map[string]interface{} {
	if o == nil || o.BalanceLog == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.BalanceLog
}

// GetBalanceLogOk returns a tuple with the BalanceLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetBalanceLogOk() ([]map[string]interface{}, bool) {
	if o == nil || o.BalanceLog == nil {
		return nil, false
	}
	return o.BalanceLog, true
}

// HasBalanceLog returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasBalanceLog() bool {
	if o != nil && o.BalanceLog != nil {
		return true
	}

	return false
}

// SetBalanceLog gets a reference to the given []map[string]interface{} and assigns it to the BalanceLog field.
func (o *GiftCardDataAttributes) SetBalanceLog(v []map[string]interface{}) {
	o.BalanceLog = v
}

// GetSingleUse returns the SingleUse field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetSingleUse() bool {
	if o == nil || o.SingleUse == nil {
		var ret bool
		return ret
	}
	return *o.SingleUse
}

// GetSingleUseOk returns a tuple with the SingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetSingleUseOk() (*bool, bool) {
	if o == nil || o.SingleUse == nil {
		return nil, false
	}
	return o.SingleUse, true
}

// HasSingleUse returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasSingleUse() bool {
	if o != nil && o.SingleUse != nil {
		return true
	}

	return false
}

// SetSingleUse gets a reference to the given bool and assigns it to the SingleUse field.
func (o *GiftCardDataAttributes) SetSingleUse(v bool) {
	o.SingleUse = &v
}

// GetRechargeable returns the Rechargeable field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetRechargeable() bool {
	if o == nil || o.Rechargeable == nil {
		var ret bool
		return ret
	}
	return *o.Rechargeable
}

// GetRechargeableOk returns a tuple with the Rechargeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetRechargeableOk() (*bool, bool) {
	if o == nil || o.Rechargeable == nil {
		return nil, false
	}
	return o.Rechargeable, true
}

// HasRechargeable returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasRechargeable() bool {
	if o != nil && o.Rechargeable != nil {
		return true
	}

	return false
}

// SetRechargeable gets a reference to the given bool and assigns it to the Rechargeable field.
func (o *GiftCardDataAttributes) SetRechargeable(v bool) {
	o.Rechargeable = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *GiftCardDataAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *GiftCardDataAttributes) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetRecipientEmail returns the RecipientEmail field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetRecipientEmail() string {
	if o == nil || o.RecipientEmail == nil {
		var ret string
		return ret
	}
	return *o.RecipientEmail
}

// GetRecipientEmailOk returns a tuple with the RecipientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetRecipientEmailOk() (*string, bool) {
	if o == nil || o.RecipientEmail == nil {
		return nil, false
	}
	return o.RecipientEmail, true
}

// HasRecipientEmail returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasRecipientEmail() bool {
	if o != nil && o.RecipientEmail != nil {
		return true
	}

	return false
}

// SetRecipientEmail gets a reference to the given string and assigns it to the RecipientEmail field.
func (o *GiftCardDataAttributes) SetRecipientEmail(v string) {
	o.RecipientEmail = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GiftCardDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GiftCardDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GiftCardDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GiftCardDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GiftCardDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GiftCardDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GiftCardDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GiftCardDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GiftCardDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.InitialBalanceCents != nil {
		toSerialize["initial_balance_cents"] = o.InitialBalanceCents
	}
	if o.InitialBalanceFloat != nil {
		toSerialize["initial_balance_float"] = o.InitialBalanceFloat
	}
	if o.FormattedInitialBalance != nil {
		toSerialize["formatted_initial_balance"] = o.FormattedInitialBalance
	}
	if o.BalanceCents != nil {
		toSerialize["balance_cents"] = o.BalanceCents
	}
	if o.BalanceFloat != nil {
		toSerialize["balance_float"] = o.BalanceFloat
	}
	if o.FormattedBalance != nil {
		toSerialize["formatted_balance"] = o.FormattedBalance
	}
	if o.BalanceMaxCents != nil {
		toSerialize["balance_max_cents"] = o.BalanceMaxCents
	}
	if o.BalanceMaxFloat != nil {
		toSerialize["balance_max_float"] = o.BalanceMaxFloat
	}
	if o.FormattedBalanceMax != nil {
		toSerialize["formatted_balance_max"] = o.FormattedBalanceMax
	}
	if o.BalanceLog != nil {
		toSerialize["balance_log"] = o.BalanceLog
	}
	if o.SingleUse != nil {
		toSerialize["single_use"] = o.SingleUse
	}
	if o.Rechargeable != nil {
		toSerialize["rechargeable"] = o.Rechargeable
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.RecipientEmail != nil {
		toSerialize["recipient_email"] = o.RecipientEmail
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

type NullableGiftCardDataAttributes struct {
	value *GiftCardDataAttributes
	isSet bool
}

func (v NullableGiftCardDataAttributes) Get() *GiftCardDataAttributes {
	return v.value
}

func (v *NullableGiftCardDataAttributes) Set(val *GiftCardDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardDataAttributes(val *GiftCardDataAttributes) *NullableGiftCardDataAttributes {
	return &NullableGiftCardDataAttributes{value: val, isSet: true}
}

func (v NullableGiftCardDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
