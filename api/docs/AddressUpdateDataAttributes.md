# AddressUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Business** | Pointer to **bool** | Indicates if it&#39;s a business or a personal address | [optional] 
**FirstName** | Pointer to **string** | Address first name (personal) | [optional] 
**LastName** | Pointer to **string** | Address last name (personal) | [optional] 
**Company** | Pointer to **string** | Address company name (business) | [optional] 
**Line1** | Pointer to **string** | Address line 1, i.e. Street address, PO Box | [optional] 
**Line2** | Pointer to **string** | Address line 2, i.e. Apartment, Suite, Building | [optional] 
**City** | Pointer to **string** | Address city | [optional] 
**ZipCode** | Pointer to **string** | ZIP or postal code | [optional] 
**StateCode** | Pointer to **string** | State, province or region code. | [optional] 
**CountryCode** | Pointer to **string** | The international 2-letter country code as defined by the ISO 3166-1 standard | [optional] 
**Phone** | Pointer to **string** | Phone number (including extension). | [optional] 
**Email** | Pointer to **string** | Email address. | [optional] 
**Notes** | Pointer to **string** | A free notes attached to the address. When used as a shipping address, this can be useful to let the customers add specific delivery instructions. | [optional] 
**Lat** | Pointer to **float32** | The address geocoded latitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**Lng** | Pointer to **float32** | The address geocoded longitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**BillingInfo** | Pointer to **string** | Customer&#39;s billing information (i.e. VAT number, codice fiscale) | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewAddressUpdateDataAttributes

`func NewAddressUpdateDataAttributes() *AddressUpdateDataAttributes`

NewAddressUpdateDataAttributes instantiates a new AddressUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressUpdateDataAttributesWithDefaults

`func NewAddressUpdateDataAttributesWithDefaults() *AddressUpdateDataAttributes`

NewAddressUpdateDataAttributesWithDefaults instantiates a new AddressUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusiness

`func (o *AddressUpdateDataAttributes) GetBusiness() bool`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *AddressUpdateDataAttributes) GetBusinessOk() (*bool, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *AddressUpdateDataAttributes) SetBusiness(v bool)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *AddressUpdateDataAttributes) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetFirstName

`func (o *AddressUpdateDataAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddressUpdateDataAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddressUpdateDataAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AddressUpdateDataAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AddressUpdateDataAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddressUpdateDataAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddressUpdateDataAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AddressUpdateDataAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompany

`func (o *AddressUpdateDataAttributes) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressUpdateDataAttributes) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressUpdateDataAttributes) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressUpdateDataAttributes) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLine1

`func (o *AddressUpdateDataAttributes) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *AddressUpdateDataAttributes) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *AddressUpdateDataAttributes) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *AddressUpdateDataAttributes) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *AddressUpdateDataAttributes) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *AddressUpdateDataAttributes) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *AddressUpdateDataAttributes) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *AddressUpdateDataAttributes) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetCity

`func (o *AddressUpdateDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressUpdateDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressUpdateDataAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressUpdateDataAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetZipCode

`func (o *AddressUpdateDataAttributes) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *AddressUpdateDataAttributes) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *AddressUpdateDataAttributes) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *AddressUpdateDataAttributes) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetStateCode

`func (o *AddressUpdateDataAttributes) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *AddressUpdateDataAttributes) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *AddressUpdateDataAttributes) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *AddressUpdateDataAttributes) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressUpdateDataAttributes) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressUpdateDataAttributes) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressUpdateDataAttributes) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AddressUpdateDataAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhone

`func (o *AddressUpdateDataAttributes) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressUpdateDataAttributes) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressUpdateDataAttributes) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressUpdateDataAttributes) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *AddressUpdateDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressUpdateDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressUpdateDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressUpdateDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNotes

`func (o *AddressUpdateDataAttributes) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AddressUpdateDataAttributes) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AddressUpdateDataAttributes) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AddressUpdateDataAttributes) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetLat

`func (o *AddressUpdateDataAttributes) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *AddressUpdateDataAttributes) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *AddressUpdateDataAttributes) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *AddressUpdateDataAttributes) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *AddressUpdateDataAttributes) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *AddressUpdateDataAttributes) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *AddressUpdateDataAttributes) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *AddressUpdateDataAttributes) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetBillingInfo

`func (o *AddressUpdateDataAttributes) GetBillingInfo() string`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *AddressUpdateDataAttributes) GetBillingInfoOk() (*string, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *AddressUpdateDataAttributes) SetBillingInfo(v string)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *AddressUpdateDataAttributes) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetReference

`func (o *AddressUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AddressUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AddressUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AddressUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *AddressUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *AddressUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *AddressUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *AddressUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *AddressUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddressUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddressUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddressUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


