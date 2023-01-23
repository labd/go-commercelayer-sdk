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

// GETOrganizationOrganizationId200ResponseDataAttributes struct for GETOrganizationOrganizationId200ResponseDataAttributes
type GETOrganizationOrganizationId200ResponseDataAttributes struct {
	// The organization's internal name.
	Name *string `json:"name,omitempty"`
	// The organization's slug name.
	Slug *string `json:"slug,omitempty"`
	// The organization's domain.
	Domain *string `json:"domain,omitempty"`
	// The organization's support phone.
	SupportPhone *string `json:"support_phone,omitempty"`
	// The organization's support email.
	SupportEmail *string `json:"support_email,omitempty"`
	// The URL to the organization's logo.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The URL to the organization's favicon.
	FaviconUrl *string `json:"favicon_url,omitempty"`
	// The organization's primary color.
	PrimaryColor *string `json:"primary_color,omitempty"`
	// The organization's contrast color.
	ContrastColor *string `json:"contrast_color,omitempty"`
	// The organization's Google Tag Manager ID.
	GtmId *string `json:"gtm_id,omitempty"`
	// The organization's Google Tag Manager ID for test.
	GtmIdTest *string `json:"gtm_id_test,omitempty"`
	// Indicates if organization has discount disabled.
	DiscountDisabled *bool `json:"discount_disabled,omitempty"`
	// Indicates if organization has account disabled.
	AccountDisabled *bool `json:"account_disabled,omitempty"`
	// Indicates if organization has acceptance disabled.
	AcceptanceDisabled *bool `json:"acceptance_disabled,omitempty"`
	// The maximum number of active concurrent promotions allowed for your organization.
	MaxConcurrentPromotions *int32 `json:"max_concurrent_promotions,omitempty"`
	// The maximum number of concurrent imports allowed for your organization.
	MaxConcurrentImports *int32 `json:"max_concurrent_imports,omitempty"`
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

// NewGETOrganizationOrganizationId200ResponseDataAttributes instantiates a new GETOrganizationOrganizationId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrganizationOrganizationId200ResponseDataAttributes() *GETOrganizationOrganizationId200ResponseDataAttributes {
	this := GETOrganizationOrganizationId200ResponseDataAttributes{}
	return &this
}

// NewGETOrganizationOrganizationId200ResponseDataAttributesWithDefaults instantiates a new GETOrganizationOrganizationId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrganizationOrganizationId200ResponseDataAttributesWithDefaults() *GETOrganizationOrganizationId200ResponseDataAttributes {
	this := GETOrganizationOrganizationId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSlug(v string) {
	o.Slug = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetDomain(v string) {
	o.Domain = &v
}

// GetSupportPhone returns the SupportPhone field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportPhone() string {
	if o == nil || o.SupportPhone == nil {
		var ret string
		return ret
	}
	return *o.SupportPhone
}

// GetSupportPhoneOk returns a tuple with the SupportPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportPhoneOk() (*string, bool) {
	if o == nil || o.SupportPhone == nil {
		return nil, false
	}
	return o.SupportPhone, true
}

// HasSupportPhone returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasSupportPhone() bool {
	if o != nil && o.SupportPhone != nil {
		return true
	}

	return false
}

// SetSupportPhone gets a reference to the given string and assigns it to the SupportPhone field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSupportPhone(v string) {
	o.SupportPhone = &v
}

// GetSupportEmail returns the SupportEmail field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportEmail() string {
	if o == nil || o.SupportEmail == nil {
		var ret string
		return ret
	}
	return *o.SupportEmail
}

// GetSupportEmailOk returns a tuple with the SupportEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportEmailOk() (*string, bool) {
	if o == nil || o.SupportEmail == nil {
		return nil, false
	}
	return o.SupportEmail, true
}

// HasSupportEmail returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasSupportEmail() bool {
	if o != nil && o.SupportEmail != nil {
		return true
	}

	return false
}

// SetSupportEmail gets a reference to the given string and assigns it to the SupportEmail field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSupportEmail(v string) {
	o.SupportEmail = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetLogoUrl() string {
	if o == nil || o.LogoUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetLogoUrlOk() (*string, bool) {
	if o == nil || o.LogoUrl == nil {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasLogoUrl() bool {
	if o != nil && o.LogoUrl != nil {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetFaviconUrl returns the FaviconUrl field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetFaviconUrl() string {
	if o == nil || o.FaviconUrl == nil {
		var ret string
		return ret
	}
	return *o.FaviconUrl
}

// GetFaviconUrlOk returns a tuple with the FaviconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetFaviconUrlOk() (*string, bool) {
	if o == nil || o.FaviconUrl == nil {
		return nil, false
	}
	return o.FaviconUrl, true
}

// HasFaviconUrl returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasFaviconUrl() bool {
	if o != nil && o.FaviconUrl != nil {
		return true
	}

	return false
}

// SetFaviconUrl gets a reference to the given string and assigns it to the FaviconUrl field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetFaviconUrl(v string) {
	o.FaviconUrl = &v
}

// GetPrimaryColor returns the PrimaryColor field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetPrimaryColor() string {
	if o == nil || o.PrimaryColor == nil {
		var ret string
		return ret
	}
	return *o.PrimaryColor
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetPrimaryColorOk() (*string, bool) {
	if o == nil || o.PrimaryColor == nil {
		return nil, false
	}
	return o.PrimaryColor, true
}

// HasPrimaryColor returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasPrimaryColor() bool {
	if o != nil && o.PrimaryColor != nil {
		return true
	}

	return false
}

// SetPrimaryColor gets a reference to the given string and assigns it to the PrimaryColor field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetPrimaryColor(v string) {
	o.PrimaryColor = &v
}

// GetContrastColor returns the ContrastColor field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetContrastColor() string {
	if o == nil || o.ContrastColor == nil {
		var ret string
		return ret
	}
	return *o.ContrastColor
}

// GetContrastColorOk returns a tuple with the ContrastColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetContrastColorOk() (*string, bool) {
	if o == nil || o.ContrastColor == nil {
		return nil, false
	}
	return o.ContrastColor, true
}

// HasContrastColor returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasContrastColor() bool {
	if o != nil && o.ContrastColor != nil {
		return true
	}

	return false
}

// SetContrastColor gets a reference to the given string and assigns it to the ContrastColor field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetContrastColor(v string) {
	o.ContrastColor = &v
}

// GetGtmId returns the GtmId field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmId() string {
	if o == nil || o.GtmId == nil {
		var ret string
		return ret
	}
	return *o.GtmId
}

// GetGtmIdOk returns a tuple with the GtmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmIdOk() (*string, bool) {
	if o == nil || o.GtmId == nil {
		return nil, false
	}
	return o.GtmId, true
}

// HasGtmId returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasGtmId() bool {
	if o != nil && o.GtmId != nil {
		return true
	}

	return false
}

// SetGtmId gets a reference to the given string and assigns it to the GtmId field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGtmId(v string) {
	o.GtmId = &v
}

// GetGtmIdTest returns the GtmIdTest field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmIdTest() string {
	if o == nil || o.GtmIdTest == nil {
		var ret string
		return ret
	}
	return *o.GtmIdTest
}

// GetGtmIdTestOk returns a tuple with the GtmIdTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmIdTestOk() (*string, bool) {
	if o == nil || o.GtmIdTest == nil {
		return nil, false
	}
	return o.GtmIdTest, true
}

// HasGtmIdTest returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasGtmIdTest() bool {
	if o != nil && o.GtmIdTest != nil {
		return true
	}

	return false
}

// SetGtmIdTest gets a reference to the given string and assigns it to the GtmIdTest field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGtmIdTest(v string) {
	o.GtmIdTest = &v
}

// GetDiscountDisabled returns the DiscountDisabled field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDiscountDisabled() bool {
	if o == nil || o.DiscountDisabled == nil {
		var ret bool
		return ret
	}
	return *o.DiscountDisabled
}

// GetDiscountDisabledOk returns a tuple with the DiscountDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDiscountDisabledOk() (*bool, bool) {
	if o == nil || o.DiscountDisabled == nil {
		return nil, false
	}
	return o.DiscountDisabled, true
}

// HasDiscountDisabled returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasDiscountDisabled() bool {
	if o != nil && o.DiscountDisabled != nil {
		return true
	}

	return false
}

// SetDiscountDisabled gets a reference to the given bool and assigns it to the DiscountDisabled field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetDiscountDisabled(v bool) {
	o.DiscountDisabled = &v
}

// GetAccountDisabled returns the AccountDisabled field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetAccountDisabled() bool {
	if o == nil || o.AccountDisabled == nil {
		var ret bool
		return ret
	}
	return *o.AccountDisabled
}

// GetAccountDisabledOk returns a tuple with the AccountDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetAccountDisabledOk() (*bool, bool) {
	if o == nil || o.AccountDisabled == nil {
		return nil, false
	}
	return o.AccountDisabled, true
}

// HasAccountDisabled returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasAccountDisabled() bool {
	if o != nil && o.AccountDisabled != nil {
		return true
	}

	return false
}

// SetAccountDisabled gets a reference to the given bool and assigns it to the AccountDisabled field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetAccountDisabled(v bool) {
	o.AccountDisabled = &v
}

// GetAcceptanceDisabled returns the AcceptanceDisabled field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetAcceptanceDisabled() bool {
	if o == nil || o.AcceptanceDisabled == nil {
		var ret bool
		return ret
	}
	return *o.AcceptanceDisabled
}

// GetAcceptanceDisabledOk returns a tuple with the AcceptanceDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetAcceptanceDisabledOk() (*bool, bool) {
	if o == nil || o.AcceptanceDisabled == nil {
		return nil, false
	}
	return o.AcceptanceDisabled, true
}

// HasAcceptanceDisabled returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasAcceptanceDisabled() bool {
	if o != nil && o.AcceptanceDisabled != nil {
		return true
	}

	return false
}

// SetAcceptanceDisabled gets a reference to the given bool and assigns it to the AcceptanceDisabled field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetAcceptanceDisabled(v bool) {
	o.AcceptanceDisabled = &v
}

// GetMaxConcurrentPromotions returns the MaxConcurrentPromotions field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMaxConcurrentPromotions() int32 {
	if o == nil || o.MaxConcurrentPromotions == nil {
		var ret int32
		return ret
	}
	return *o.MaxConcurrentPromotions
}

// GetMaxConcurrentPromotionsOk returns a tuple with the MaxConcurrentPromotions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMaxConcurrentPromotionsOk() (*int32, bool) {
	if o == nil || o.MaxConcurrentPromotions == nil {
		return nil, false
	}
	return o.MaxConcurrentPromotions, true
}

// HasMaxConcurrentPromotions returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasMaxConcurrentPromotions() bool {
	if o != nil && o.MaxConcurrentPromotions != nil {
		return true
	}

	return false
}

// SetMaxConcurrentPromotions gets a reference to the given int32 and assigns it to the MaxConcurrentPromotions field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetMaxConcurrentPromotions(v int32) {
	o.MaxConcurrentPromotions = &v
}

// GetMaxConcurrentImports returns the MaxConcurrentImports field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMaxConcurrentImports() int32 {
	if o == nil || o.MaxConcurrentImports == nil {
		var ret int32
		return ret
	}
	return *o.MaxConcurrentImports
}

// GetMaxConcurrentImportsOk returns a tuple with the MaxConcurrentImports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMaxConcurrentImportsOk() (*int32, bool) {
	if o == nil || o.MaxConcurrentImports == nil {
		return nil, false
	}
	return o.MaxConcurrentImports, true
}

// HasMaxConcurrentImports returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasMaxConcurrentImports() bool {
	if o != nil && o.MaxConcurrentImports != nil {
		return true
	}

	return false
}

// SetMaxConcurrentImports gets a reference to the given int32 and assigns it to the MaxConcurrentImports field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetMaxConcurrentImports(v int32) {
	o.MaxConcurrentImports = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETOrganizationOrganizationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.SupportPhone != nil {
		toSerialize["support_phone"] = o.SupportPhone
	}
	if o.SupportEmail != nil {
		toSerialize["support_email"] = o.SupportEmail
	}
	if o.LogoUrl != nil {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if o.FaviconUrl != nil {
		toSerialize["favicon_url"] = o.FaviconUrl
	}
	if o.PrimaryColor != nil {
		toSerialize["primary_color"] = o.PrimaryColor
	}
	if o.ContrastColor != nil {
		toSerialize["contrast_color"] = o.ContrastColor
	}
	if o.GtmId != nil {
		toSerialize["gtm_id"] = o.GtmId
	}
	if o.GtmIdTest != nil {
		toSerialize["gtm_id_test"] = o.GtmIdTest
	}
	if o.DiscountDisabled != nil {
		toSerialize["discount_disabled"] = o.DiscountDisabled
	}
	if o.AccountDisabled != nil {
		toSerialize["account_disabled"] = o.AccountDisabled
	}
	if o.AcceptanceDisabled != nil {
		toSerialize["acceptance_disabled"] = o.AcceptanceDisabled
	}
	if o.MaxConcurrentPromotions != nil {
		toSerialize["max_concurrent_promotions"] = o.MaxConcurrentPromotions
	}
	if o.MaxConcurrentImports != nil {
		toSerialize["max_concurrent_imports"] = o.MaxConcurrentImports
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

type NullableGETOrganizationOrganizationId200ResponseDataAttributes struct {
	value *GETOrganizationOrganizationId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETOrganizationOrganizationId200ResponseDataAttributes) Get() *GETOrganizationOrganizationId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETOrganizationOrganizationId200ResponseDataAttributes) Set(val *GETOrganizationOrganizationId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrganizationOrganizationId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrganizationOrganizationId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrganizationOrganizationId200ResponseDataAttributes(val *GETOrganizationOrganizationId200ResponseDataAttributes) *NullableGETOrganizationOrganizationId200ResponseDataAttributes {
	return &NullableGETOrganizationOrganizationId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETOrganizationOrganizationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrganizationOrganizationId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
