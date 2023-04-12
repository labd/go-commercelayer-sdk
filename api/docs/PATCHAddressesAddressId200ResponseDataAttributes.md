# PATCHAddressesAddressId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Business** | Pointer to **interface{}** | Indicates if it&#39;s a business or a personal address | [optional] 
**FirstName** | Pointer to **interface{}** | Address first name (personal) | [optional] 
**LastName** | Pointer to **interface{}** | Address last name (personal) | [optional] 
**Company** | Pointer to **interface{}** | Address company name (business) | [optional] 
**Line1** | Pointer to **interface{}** | Address line 1, i.e. Street address, PO Box | [optional] 
**Line2** | Pointer to **interface{}** | Address line 2, i.e. Apartment, Suite, Building | [optional] 
**City** | Pointer to **interface{}** | Address city | [optional] 
**ZipCode** | Pointer to **interface{}** | ZIP or postal code | [optional] 
**StateCode** | Pointer to **interface{}** | State, province or region code. | [optional] 
**CountryCode** | Pointer to **interface{}** | The international 2-letter country code as defined by the ISO 3166-1 standard | [optional] 
**Phone** | Pointer to **interface{}** | Phone number (including extension). | [optional] 
**Email** | Pointer to **interface{}** | Email address. | [optional] 
**Notes** | Pointer to **interface{}** | A free notes attached to the address. When used as a shipping address, this can be useful to let the customers add specific delivery instructions. | [optional] 
**Lat** | Pointer to **interface{}** | The address geocoded latitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**Lng** | Pointer to **interface{}** | The address geocoded longitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**BillingInfo** | Pointer to **interface{}** | Customer&#39;s billing information (i.e. VAT number, codice fiscale) | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHAddressesAddressId200ResponseDataAttributes

`func NewPATCHAddressesAddressId200ResponseDataAttributes() *PATCHAddressesAddressId200ResponseDataAttributes`

NewPATCHAddressesAddressId200ResponseDataAttributes instantiates a new PATCHAddressesAddressId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHAddressesAddressId200ResponseDataAttributesWithDefaults

`func NewPATCHAddressesAddressId200ResponseDataAttributesWithDefaults() *PATCHAddressesAddressId200ResponseDataAttributes`

NewPATCHAddressesAddressId200ResponseDataAttributesWithDefaults instantiates a new PATCHAddressesAddressId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusiness

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetBusiness() interface{}`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetBusinessOk() (*interface{}, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetBusiness(v interface{})`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### SetBusinessNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetBusinessNil(b bool)`

 SetBusinessNil sets the value for Business to be an explicit nil

### UnsetBusiness
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetBusiness()`

UnsetBusiness ensures that no value is present for Business, not even an explicit nil
### GetFirstName

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetFirstName() interface{}`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetFirstNameOk() (*interface{}, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetFirstName(v interface{})`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLastName() interface{}`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLastNameOk() (*interface{}, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLastName(v interface{})`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompany

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCompany() interface{}`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCompanyOk() (*interface{}, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetCompany(v interface{})`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetLine1

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLine1() interface{}`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLine1Ok() (*interface{}, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLine1(v interface{})`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### SetLine1Nil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLine2() interface{}`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLine2Ok() (*interface{}, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLine2(v interface{})`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetCity

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCity() interface{}`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCityOk() (*interface{}, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetCity(v interface{})`

SetCity sets City field to given value.

### HasCity

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetZipCode

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetZipCode() interface{}`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetZipCodeOk() (*interface{}, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetZipCode(v interface{})`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetStateCode

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetStateCode() interface{}`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetStateCodeOk() (*interface{}, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetStateCode(v interface{})`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### SetStateCodeNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetStateCodeNil(b bool)`

 SetStateCodeNil sets the value for StateCode to be an explicit nil

### UnsetStateCode
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetStateCode()`

UnsetStateCode ensures that no value is present for StateCode, not even an explicit nil
### GetCountryCode

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCountryCode() interface{}`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCountryCodeOk() (*interface{}, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetCountryCode(v interface{})`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetPhone

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetPhone() interface{}`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetPhoneOk() (*interface{}, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetPhone(v interface{})`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetNotes

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetNotes() interface{}`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetNotesOk() (*interface{}, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetNotes(v interface{})`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetLat

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLat() interface{}`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLatOk() (*interface{}, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLat(v interface{})`

SetLat sets Lat field to given value.

### HasLat

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLng() interface{}`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLngOk() (*interface{}, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLng(v interface{})`

SetLng sets Lng field to given value.

### HasLng

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetBillingInfo

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetBillingInfo() interface{}`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetBillingInfoOk() (*interface{}, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetBillingInfo(v interface{})`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### SetBillingInfoNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetBillingInfoNil(b bool)`

 SetBillingInfoNil sets the value for BillingInfo to be an explicit nil

### UnsetBillingInfo
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetBillingInfo()`

UnsetBillingInfo ensures that no value is present for BillingInfo, not even an explicit nil
### GetReference

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHAddressesAddressId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


