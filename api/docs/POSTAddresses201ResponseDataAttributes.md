# POSTAddresses201ResponseDataAttributes

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

### NewPOSTAddresses201ResponseDataAttributes

`func NewPOSTAddresses201ResponseDataAttributes(line1 string, city string, stateCode string, countryCode string, phone string, ) *POSTAddresses201ResponseDataAttributes`

NewPOSTAddresses201ResponseDataAttributes instantiates a new POSTAddresses201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTAddresses201ResponseDataAttributesWithDefaults

`func NewPOSTAddresses201ResponseDataAttributesWithDefaults() *POSTAddresses201ResponseDataAttributes`

NewPOSTAddresses201ResponseDataAttributesWithDefaults instantiates a new POSTAddresses201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusiness

`func (o *POSTAddresses201ResponseDataAttributes) GetBusiness() bool`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *POSTAddresses201ResponseDataAttributes) GetBusinessOk() (*bool, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *POSTAddresses201ResponseDataAttributes) SetBusiness(v bool)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *POSTAddresses201ResponseDataAttributes) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetFirstName

`func (o *POSTAddresses201ResponseDataAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *POSTAddresses201ResponseDataAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *POSTAddresses201ResponseDataAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *POSTAddresses201ResponseDataAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *POSTAddresses201ResponseDataAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *POSTAddresses201ResponseDataAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *POSTAddresses201ResponseDataAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *POSTAddresses201ResponseDataAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompany

`func (o *POSTAddresses201ResponseDataAttributes) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *POSTAddresses201ResponseDataAttributes) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *POSTAddresses201ResponseDataAttributes) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *POSTAddresses201ResponseDataAttributes) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLine1

`func (o *POSTAddresses201ResponseDataAttributes) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *POSTAddresses201ResponseDataAttributes) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *POSTAddresses201ResponseDataAttributes) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### GetLine2

`func (o *POSTAddresses201ResponseDataAttributes) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *POSTAddresses201ResponseDataAttributes) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *POSTAddresses201ResponseDataAttributes) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *POSTAddresses201ResponseDataAttributes) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetCity

`func (o *POSTAddresses201ResponseDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *POSTAddresses201ResponseDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *POSTAddresses201ResponseDataAttributes) SetCity(v string)`

SetCity sets City field to given value.


### GetZipCode

`func (o *POSTAddresses201ResponseDataAttributes) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *POSTAddresses201ResponseDataAttributes) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *POSTAddresses201ResponseDataAttributes) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *POSTAddresses201ResponseDataAttributes) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetStateCode

`func (o *POSTAddresses201ResponseDataAttributes) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *POSTAddresses201ResponseDataAttributes) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *POSTAddresses201ResponseDataAttributes) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.


### GetCountryCode

`func (o *POSTAddresses201ResponseDataAttributes) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *POSTAddresses201ResponseDataAttributes) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *POSTAddresses201ResponseDataAttributes) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPhone

`func (o *POSTAddresses201ResponseDataAttributes) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *POSTAddresses201ResponseDataAttributes) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *POSTAddresses201ResponseDataAttributes) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetEmail

`func (o *POSTAddresses201ResponseDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *POSTAddresses201ResponseDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *POSTAddresses201ResponseDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *POSTAddresses201ResponseDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNotes

`func (o *POSTAddresses201ResponseDataAttributes) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *POSTAddresses201ResponseDataAttributes) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *POSTAddresses201ResponseDataAttributes) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *POSTAddresses201ResponseDataAttributes) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetLat

`func (o *POSTAddresses201ResponseDataAttributes) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *POSTAddresses201ResponseDataAttributes) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *POSTAddresses201ResponseDataAttributes) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *POSTAddresses201ResponseDataAttributes) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *POSTAddresses201ResponseDataAttributes) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *POSTAddresses201ResponseDataAttributes) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *POSTAddresses201ResponseDataAttributes) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *POSTAddresses201ResponseDataAttributes) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetBillingInfo

`func (o *POSTAddresses201ResponseDataAttributes) GetBillingInfo() string`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *POSTAddresses201ResponseDataAttributes) GetBillingInfoOk() (*string, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *POSTAddresses201ResponseDataAttributes) SetBillingInfo(v string)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *POSTAddresses201ResponseDataAttributes) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetReference

`func (o *POSTAddresses201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTAddresses201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTAddresses201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTAddresses201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTAddresses201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTAddresses201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTAddresses201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTAddresses201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTAddresses201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTAddresses201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTAddresses201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTAddresses201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


