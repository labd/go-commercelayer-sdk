# AddressCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Business** | Pointer to **bool** | Indicates if it&#39;s a business or a personal address | [optional] 
**FirstName** | Pointer to **string** | Address first name (personal) | [optional] 
**LastName** | Pointer to **string** | Address last name (personal) | [optional] 
**Company** | Pointer to **string** | Address company name (business) | [optional] 
**Line1** | **string** | Address line 1, i.e. Street address, PO Box | 
**Line2** | Pointer to **string** | Address line 2, i.e. Apartment, Suite, Building | [optional] 
**City** | **string** | Address city | 
**ZipCode** | Pointer to **string** | ZIP or postal code | [optional] 
**StateCode** | **string** | State, province or region code. | 
**CountryCode** | **string** | The international 2-letter country code as defined by the ISO 3166-1 standard | 
**Phone** | **string** | Phone number (including extension). | 
**Email** | Pointer to **string** | Email address. | [optional] 
**Notes** | Pointer to **string** | A free notes attached to the address. When used as a shipping address, this can be useful to let the customers add specific delivery instructions. | [optional] 
**Lat** | Pointer to **float32** | The address geocoded latitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**Lng** | Pointer to **float32** | The address geocoded longitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**BillingInfo** | Pointer to **string** | Customer&#39;s billing information (i.e. VAT number, codice fiscale) | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewAddressCreateDataAttributes

`func NewAddressCreateDataAttributes(line1 string, city string, stateCode string, countryCode string, phone string, ) *AddressCreateDataAttributes`

NewAddressCreateDataAttributes instantiates a new AddressCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressCreateDataAttributesWithDefaults

`func NewAddressCreateDataAttributesWithDefaults() *AddressCreateDataAttributes`

NewAddressCreateDataAttributesWithDefaults instantiates a new AddressCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusiness

`func (o *AddressCreateDataAttributes) GetBusiness() bool`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *AddressCreateDataAttributes) GetBusinessOk() (*bool, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *AddressCreateDataAttributes) SetBusiness(v bool)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *AddressCreateDataAttributes) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetFirstName

`func (o *AddressCreateDataAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddressCreateDataAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddressCreateDataAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AddressCreateDataAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AddressCreateDataAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddressCreateDataAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddressCreateDataAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AddressCreateDataAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompany

`func (o *AddressCreateDataAttributes) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressCreateDataAttributes) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressCreateDataAttributes) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressCreateDataAttributes) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLine1

`func (o *AddressCreateDataAttributes) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *AddressCreateDataAttributes) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *AddressCreateDataAttributes) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### GetLine2

`func (o *AddressCreateDataAttributes) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *AddressCreateDataAttributes) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *AddressCreateDataAttributes) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *AddressCreateDataAttributes) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetCity

`func (o *AddressCreateDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressCreateDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressCreateDataAttributes) SetCity(v string)`

SetCity sets City field to given value.


### GetZipCode

`func (o *AddressCreateDataAttributes) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *AddressCreateDataAttributes) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *AddressCreateDataAttributes) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *AddressCreateDataAttributes) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetStateCode

`func (o *AddressCreateDataAttributes) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *AddressCreateDataAttributes) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *AddressCreateDataAttributes) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.


### GetCountryCode

`func (o *AddressCreateDataAttributes) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressCreateDataAttributes) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressCreateDataAttributes) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPhone

`func (o *AddressCreateDataAttributes) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressCreateDataAttributes) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressCreateDataAttributes) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetEmail

`func (o *AddressCreateDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressCreateDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressCreateDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressCreateDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNotes

`func (o *AddressCreateDataAttributes) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AddressCreateDataAttributes) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AddressCreateDataAttributes) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AddressCreateDataAttributes) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetLat

`func (o *AddressCreateDataAttributes) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *AddressCreateDataAttributes) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *AddressCreateDataAttributes) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *AddressCreateDataAttributes) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *AddressCreateDataAttributes) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *AddressCreateDataAttributes) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *AddressCreateDataAttributes) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *AddressCreateDataAttributes) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetBillingInfo

`func (o *AddressCreateDataAttributes) GetBillingInfo() string`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *AddressCreateDataAttributes) GetBillingInfoOk() (*string, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *AddressCreateDataAttributes) SetBillingInfo(v string)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *AddressCreateDataAttributes) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetReference

`func (o *AddressCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AddressCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AddressCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AddressCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *AddressCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *AddressCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *AddressCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *AddressCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *AddressCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddressCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddressCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddressCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


