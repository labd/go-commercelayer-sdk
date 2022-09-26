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

// AddressDataAttributes struct for AddressDataAttributes
type AddressDataAttributes struct {
	// Indicates if it's a business or a personal address
	Business *bool `json:"business,omitempty"`
	// Address first name (personal)
	FirstName *string `json:"first_name,omitempty"`
	// Address last name (personal)
	LastName *string `json:"last_name,omitempty"`
	// Address company name (business)
	Company *string `json:"company,omitempty"`
	// Company name (business) of first name and last name (personal)
	FullName *string `json:"full_name,omitempty"`
	// Address line 1, i.e. Street address, PO Box
	Line1 *string `json:"line_1,omitempty"`
	// Address line 2, i.e. Apartment, Suite, Building
	Line2 *string `json:"line_2,omitempty"`
	// Address city
	City *string `json:"city,omitempty"`
	// ZIP or postal code
	ZipCode *string `json:"zip_code,omitempty"`
	// State, province or region code.
	StateCode *string `json:"state_code,omitempty"`
	// The international 2-letter country code as defined by the ISO 3166-1 standard
	CountryCode *string `json:"country_code,omitempty"`
	// Phone number (including extension).
	Phone *string `json:"phone,omitempty"`
	// Compact description of the address location, without the full name
	FullAddress *string `json:"full_address,omitempty"`
	// Compact description of the address location, including the full name
	Name *string `json:"name,omitempty"`
	// Email address.
	Email *string `json:"email,omitempty"`
	// A free notes attached to the address. When used as a shipping address, this can be useful to let the customers add specific delivery instructions.
	Notes *string `json:"notes,omitempty"`
	// The address geocoded latitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order's market.
	Lat *float32 `json:"lat,omitempty"`
	// The address geocoded longitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order's market.
	Lng *float32 `json:"lng,omitempty"`
	// Indicates if the latitude and logitude are present, either geocoded or manually updated
	IsLocalized *bool `json:"is_localized,omitempty"`
	// Indicates if the address has been successfully geocoded
	IsGeocoded *bool `json:"is_geocoded,omitempty"`
	// The geocoder provider name (either Google or Bing)
	ProviderName *string `json:"provider_name,omitempty"`
	// The map url of the geocoded address (if geocoded)
	MapUrl *string `json:"map_url,omitempty"`
	// The static map image url of the geocoded address (if geocoded)
	StaticMapUrl *string `json:"static_map_url,omitempty"`
	// Customer's billing information (i.e. VAT number, codice fiscale)
	BillingInfo *string `json:"billing_info,omitempty"`
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

// NewAddressDataAttributes instantiates a new AddressDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressDataAttributes() *AddressDataAttributes {
	this := AddressDataAttributes{}
	return &this
}

// NewAddressDataAttributesWithDefaults instantiates a new AddressDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressDataAttributesWithDefaults() *AddressDataAttributes {
	this := AddressDataAttributes{}
	return &this
}

// GetBusiness returns the Business field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetBusiness() bool {
	if o == nil || o.Business == nil {
		var ret bool
		return ret
	}
	return *o.Business
}

// GetBusinessOk returns a tuple with the Business field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetBusinessOk() (*bool, bool) {
	if o == nil || o.Business == nil {
		return nil, false
	}
	return o.Business, true
}

// HasBusiness returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasBusiness() bool {
	if o != nil && o.Business != nil {
		return true
	}

	return false
}

// SetBusiness gets a reference to the given bool and assigns it to the Business field.
func (o *AddressDataAttributes) SetBusiness(v bool) {
	o.Business = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *AddressDataAttributes) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *AddressDataAttributes) SetLastName(v string) {
	o.LastName = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *AddressDataAttributes) SetCompany(v string) {
	o.Company = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasFullName() bool {
	if o != nil && o.FullName != nil {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *AddressDataAttributes) SetFullName(v string) {
	o.FullName = &v
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetLine1() string {
	if o == nil || o.Line1 == nil {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetLine1Ok() (*string, bool) {
	if o == nil || o.Line1 == nil {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasLine1() bool {
	if o != nil && o.Line1 != nil {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *AddressDataAttributes) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetLine2() string {
	if o == nil || o.Line2 == nil {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetLine2Ok() (*string, bool) {
	if o == nil || o.Line2 == nil {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasLine2() bool {
	if o != nil && o.Line2 != nil {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *AddressDataAttributes) SetLine2(v string) {
	o.Line2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AddressDataAttributes) SetCity(v string) {
	o.City = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetZipCodeOk() (*string, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *AddressDataAttributes) SetZipCode(v string) {
	o.ZipCode = &v
}

// GetStateCode returns the StateCode field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetStateCode() string {
	if o == nil || o.StateCode == nil {
		var ret string
		return ret
	}
	return *o.StateCode
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetStateCodeOk() (*string, bool) {
	if o == nil || o.StateCode == nil {
		return nil, false
	}
	return o.StateCode, true
}

// HasStateCode returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasStateCode() bool {
	if o != nil && o.StateCode != nil {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given string and assigns it to the StateCode field.
func (o *AddressDataAttributes) SetStateCode(v string) {
	o.StateCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *AddressDataAttributes) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *AddressDataAttributes) SetPhone(v string) {
	o.Phone = &v
}

// GetFullAddress returns the FullAddress field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetFullAddress() string {
	if o == nil || o.FullAddress == nil {
		var ret string
		return ret
	}
	return *o.FullAddress
}

// GetFullAddressOk returns a tuple with the FullAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetFullAddressOk() (*string, bool) {
	if o == nil || o.FullAddress == nil {
		return nil, false
	}
	return o.FullAddress, true
}

// HasFullAddress returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasFullAddress() bool {
	if o != nil && o.FullAddress != nil {
		return true
	}

	return false
}

// SetFullAddress gets a reference to the given string and assigns it to the FullAddress field.
func (o *AddressDataAttributes) SetFullAddress(v string) {
	o.FullAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddressDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AddressDataAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *AddressDataAttributes) SetNotes(v string) {
	o.Notes = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetLat() float32 {
	if o == nil || o.Lat == nil {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetLatOk() (*float32, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *AddressDataAttributes) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetLng() float32 {
	if o == nil || o.Lng == nil {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetLngOk() (*float32, bool) {
	if o == nil || o.Lng == nil {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasLng() bool {
	if o != nil && o.Lng != nil {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *AddressDataAttributes) SetLng(v float32) {
	o.Lng = &v
}

// GetIsLocalized returns the IsLocalized field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetIsLocalized() bool {
	if o == nil || o.IsLocalized == nil {
		var ret bool
		return ret
	}
	return *o.IsLocalized
}

// GetIsLocalizedOk returns a tuple with the IsLocalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetIsLocalizedOk() (*bool, bool) {
	if o == nil || o.IsLocalized == nil {
		return nil, false
	}
	return o.IsLocalized, true
}

// HasIsLocalized returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasIsLocalized() bool {
	if o != nil && o.IsLocalized != nil {
		return true
	}

	return false
}

// SetIsLocalized gets a reference to the given bool and assigns it to the IsLocalized field.
func (o *AddressDataAttributes) SetIsLocalized(v bool) {
	o.IsLocalized = &v
}

// GetIsGeocoded returns the IsGeocoded field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetIsGeocoded() bool {
	if o == nil || o.IsGeocoded == nil {
		var ret bool
		return ret
	}
	return *o.IsGeocoded
}

// GetIsGeocodedOk returns a tuple with the IsGeocoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetIsGeocodedOk() (*bool, bool) {
	if o == nil || o.IsGeocoded == nil {
		return nil, false
	}
	return o.IsGeocoded, true
}

// HasIsGeocoded returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasIsGeocoded() bool {
	if o != nil && o.IsGeocoded != nil {
		return true
	}

	return false
}

// SetIsGeocoded gets a reference to the given bool and assigns it to the IsGeocoded field.
func (o *AddressDataAttributes) SetIsGeocoded(v bool) {
	o.IsGeocoded = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *AddressDataAttributes) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetMapUrl returns the MapUrl field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetMapUrl() string {
	if o == nil || o.MapUrl == nil {
		var ret string
		return ret
	}
	return *o.MapUrl
}

// GetMapUrlOk returns a tuple with the MapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetMapUrlOk() (*string, bool) {
	if o == nil || o.MapUrl == nil {
		return nil, false
	}
	return o.MapUrl, true
}

// HasMapUrl returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasMapUrl() bool {
	if o != nil && o.MapUrl != nil {
		return true
	}

	return false
}

// SetMapUrl gets a reference to the given string and assigns it to the MapUrl field.
func (o *AddressDataAttributes) SetMapUrl(v string) {
	o.MapUrl = &v
}

// GetStaticMapUrl returns the StaticMapUrl field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetStaticMapUrl() string {
	if o == nil || o.StaticMapUrl == nil {
		var ret string
		return ret
	}
	return *o.StaticMapUrl
}

// GetStaticMapUrlOk returns a tuple with the StaticMapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetStaticMapUrlOk() (*string, bool) {
	if o == nil || o.StaticMapUrl == nil {
		return nil, false
	}
	return o.StaticMapUrl, true
}

// HasStaticMapUrl returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasStaticMapUrl() bool {
	if o != nil && o.StaticMapUrl != nil {
		return true
	}

	return false
}

// SetStaticMapUrl gets a reference to the given string and assigns it to the StaticMapUrl field.
func (o *AddressDataAttributes) SetStaticMapUrl(v string) {
	o.StaticMapUrl = &v
}

// GetBillingInfo returns the BillingInfo field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetBillingInfo() string {
	if o == nil || o.BillingInfo == nil {
		var ret string
		return ret
	}
	return *o.BillingInfo
}

// GetBillingInfoOk returns a tuple with the BillingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetBillingInfoOk() (*string, bool) {
	if o == nil || o.BillingInfo == nil {
		return nil, false
	}
	return o.BillingInfo, true
}

// HasBillingInfo returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasBillingInfo() bool {
	if o != nil && o.BillingInfo != nil {
		return true
	}

	return false
}

// SetBillingInfo gets a reference to the given string and assigns it to the BillingInfo field.
func (o *AddressDataAttributes) SetBillingInfo(v string) {
	o.BillingInfo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AddressDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AddressDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AddressDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *AddressDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *AddressDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AddressDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AddressDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AddressDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o AddressDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Business != nil {
		toSerialize["business"] = o.Business
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.FullName != nil {
		toSerialize["full_name"] = o.FullName
	}
	if o.Line1 != nil {
		toSerialize["line_1"] = o.Line1
	}
	if o.Line2 != nil {
		toSerialize["line_2"] = o.Line2
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.ZipCode != nil {
		toSerialize["zip_code"] = o.ZipCode
	}
	if o.StateCode != nil {
		toSerialize["state_code"] = o.StateCode
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.FullAddress != nil {
		toSerialize["full_address"] = o.FullAddress
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Lng != nil {
		toSerialize["lng"] = o.Lng
	}
	if o.IsLocalized != nil {
		toSerialize["is_localized"] = o.IsLocalized
	}
	if o.IsGeocoded != nil {
		toSerialize["is_geocoded"] = o.IsGeocoded
	}
	if o.ProviderName != nil {
		toSerialize["provider_name"] = o.ProviderName
	}
	if o.MapUrl != nil {
		toSerialize["map_url"] = o.MapUrl
	}
	if o.StaticMapUrl != nil {
		toSerialize["static_map_url"] = o.StaticMapUrl
	}
	if o.BillingInfo != nil {
		toSerialize["billing_info"] = o.BillingInfo
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

type NullableAddressDataAttributes struct {
	value *AddressDataAttributes
	isSet bool
}

func (v NullableAddressDataAttributes) Get() *AddressDataAttributes {
	return v.value
}

func (v *NullableAddressDataAttributes) Set(val *AddressDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressDataAttributes(val *AddressDataAttributes) *NullableAddressDataAttributes {
	return &NullableAddressDataAttributes{value: val, isSet: true}
}

func (v NullableAddressDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
