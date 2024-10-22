# POSTAddresses201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Business** | Pointer to **interface{}** | Indicates if it&#39;s a business or a personal address. | [optional] 
**FirstName** | Pointer to **interface{}** | Address first name (personal). | [optional] 
**LastName** | Pointer to **interface{}** | Address last name (personal). | [optional] 
**Company** | Pointer to **interface{}** | Address company name (business). | [optional] 
**Line1** | **interface{}** | Address line 1, i.e. Street address, PO Box. | 
**Line2** | Pointer to **interface{}** | Address line 2, i.e. Apartment, Suite, Building. | [optional] 
**City** | **interface{}** | Address city. | 
**ZipCode** | Pointer to **interface{}** | ZIP or postal code. | [optional] 
**StateCode** | **interface{}** | State, province or region code. | 
**CountryCode** | **interface{}** | The international 2-letter country code as defined by the ISO 3166-1 standard. | 
**Phone** | **interface{}** | Phone number (including extension). | 
**Email** | Pointer to **interface{}** | Email address. | [optional] 
**Notes** | Pointer to **interface{}** | A free notes attached to the address. When used as a shipping address, this can be useful to let the customers add specific delivery instructions. | [optional] 
**Lat** | Pointer to **interface{}** | The address geocoded latitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**Lng** | Pointer to **interface{}** | The address geocoded longitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**BillingInfo** | Pointer to **interface{}** | Customer&#39;s billing information (i.e. VAT number, codice fiscale). | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTAddresses201ResponseDataAttributes

`func NewPOSTAddresses201ResponseDataAttributes(line1 interface{}, city interface{}, stateCode interface{}, countryCode interface{}, phone interface{}, ) *POSTAddresses201ResponseDataAttributes`

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

`func (o *POSTAddresses201ResponseDataAttributes) GetBusiness() interface{}`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *POSTAddresses201ResponseDataAttributes) GetBusinessOk() (*interface{}, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *POSTAddresses201ResponseDataAttributes) SetBusiness(v interface{})`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *POSTAddresses201ResponseDataAttributes) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### SetBusinessNil

`func (o *POSTAddresses201ResponseDataAttributes) SetBusinessNil(b bool)`

 SetBusinessNil sets the value for Business to be an explicit nil

### UnsetBusiness
`func (o *POSTAddresses201ResponseDataAttributes) UnsetBusiness()`

UnsetBusiness ensures that no value is present for Business, not even an explicit nil
### GetFirstName

`func (o *POSTAddresses201ResponseDataAttributes) GetFirstName() interface{}`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *POSTAddresses201ResponseDataAttributes) GetFirstNameOk() (*interface{}, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *POSTAddresses201ResponseDataAttributes) SetFirstName(v interface{})`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *POSTAddresses201ResponseDataAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *POSTAddresses201ResponseDataAttributes) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *POSTAddresses201ResponseDataAttributes) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *POSTAddresses201ResponseDataAttributes) GetLastName() interface{}`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *POSTAddresses201ResponseDataAttributes) GetLastNameOk() (*interface{}, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *POSTAddresses201ResponseDataAttributes) SetLastName(v interface{})`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *POSTAddresses201ResponseDataAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *POSTAddresses201ResponseDataAttributes) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *POSTAddresses201ResponseDataAttributes) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompany

`func (o *POSTAddresses201ResponseDataAttributes) GetCompany() interface{}`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *POSTAddresses201ResponseDataAttributes) GetCompanyOk() (*interface{}, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *POSTAddresses201ResponseDataAttributes) SetCompany(v interface{})`

SetCompany sets Company field to given value.

### HasCompany

`func (o *POSTAddresses201ResponseDataAttributes) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *POSTAddresses201ResponseDataAttributes) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *POSTAddresses201ResponseDataAttributes) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetLine1

`func (o *POSTAddresses201ResponseDataAttributes) GetLine1() interface{}`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *POSTAddresses201ResponseDataAttributes) GetLine1Ok() (*interface{}, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *POSTAddresses201ResponseDataAttributes) SetLine1(v interface{})`

SetLine1 sets Line1 field to given value.


### SetLine1Nil

`func (o *POSTAddresses201ResponseDataAttributes) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *POSTAddresses201ResponseDataAttributes) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *POSTAddresses201ResponseDataAttributes) GetLine2() interface{}`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *POSTAddresses201ResponseDataAttributes) GetLine2Ok() (*interface{}, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *POSTAddresses201ResponseDataAttributes) SetLine2(v interface{})`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *POSTAddresses201ResponseDataAttributes) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *POSTAddresses201ResponseDataAttributes) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *POSTAddresses201ResponseDataAttributes) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetCity

`func (o *POSTAddresses201ResponseDataAttributes) GetCity() interface{}`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *POSTAddresses201ResponseDataAttributes) GetCityOk() (*interface{}, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *POSTAddresses201ResponseDataAttributes) SetCity(v interface{})`

SetCity sets City field to given value.


### SetCityNil

`func (o *POSTAddresses201ResponseDataAttributes) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *POSTAddresses201ResponseDataAttributes) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetZipCode

`func (o *POSTAddresses201ResponseDataAttributes) GetZipCode() interface{}`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *POSTAddresses201ResponseDataAttributes) GetZipCodeOk() (*interface{}, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *POSTAddresses201ResponseDataAttributes) SetZipCode(v interface{})`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *POSTAddresses201ResponseDataAttributes) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *POSTAddresses201ResponseDataAttributes) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *POSTAddresses201ResponseDataAttributes) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetStateCode

`func (o *POSTAddresses201ResponseDataAttributes) GetStateCode() interface{}`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *POSTAddresses201ResponseDataAttributes) GetStateCodeOk() (*interface{}, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *POSTAddresses201ResponseDataAttributes) SetStateCode(v interface{})`

SetStateCode sets StateCode field to given value.


### SetStateCodeNil

`func (o *POSTAddresses201ResponseDataAttributes) SetStateCodeNil(b bool)`

 SetStateCodeNil sets the value for StateCode to be an explicit nil

### UnsetStateCode
`func (o *POSTAddresses201ResponseDataAttributes) UnsetStateCode()`

UnsetStateCode ensures that no value is present for StateCode, not even an explicit nil
### GetCountryCode

`func (o *POSTAddresses201ResponseDataAttributes) GetCountryCode() interface{}`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *POSTAddresses201ResponseDataAttributes) GetCountryCodeOk() (*interface{}, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *POSTAddresses201ResponseDataAttributes) SetCountryCode(v interface{})`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *POSTAddresses201ResponseDataAttributes) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *POSTAddresses201ResponseDataAttributes) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetPhone

`func (o *POSTAddresses201ResponseDataAttributes) GetPhone() interface{}`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *POSTAddresses201ResponseDataAttributes) GetPhoneOk() (*interface{}, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *POSTAddresses201ResponseDataAttributes) SetPhone(v interface{})`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *POSTAddresses201ResponseDataAttributes) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *POSTAddresses201ResponseDataAttributes) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *POSTAddresses201ResponseDataAttributes) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *POSTAddresses201ResponseDataAttributes) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *POSTAddresses201ResponseDataAttributes) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *POSTAddresses201ResponseDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *POSTAddresses201ResponseDataAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *POSTAddresses201ResponseDataAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetNotes

`func (o *POSTAddresses201ResponseDataAttributes) GetNotes() interface{}`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *POSTAddresses201ResponseDataAttributes) GetNotesOk() (*interface{}, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *POSTAddresses201ResponseDataAttributes) SetNotes(v interface{})`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *POSTAddresses201ResponseDataAttributes) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *POSTAddresses201ResponseDataAttributes) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *POSTAddresses201ResponseDataAttributes) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetLat

`func (o *POSTAddresses201ResponseDataAttributes) GetLat() interface{}`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *POSTAddresses201ResponseDataAttributes) GetLatOk() (*interface{}, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *POSTAddresses201ResponseDataAttributes) SetLat(v interface{})`

SetLat sets Lat field to given value.

### HasLat

`func (o *POSTAddresses201ResponseDataAttributes) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *POSTAddresses201ResponseDataAttributes) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *POSTAddresses201ResponseDataAttributes) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *POSTAddresses201ResponseDataAttributes) GetLng() interface{}`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *POSTAddresses201ResponseDataAttributes) GetLngOk() (*interface{}, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *POSTAddresses201ResponseDataAttributes) SetLng(v interface{})`

SetLng sets Lng field to given value.

### HasLng

`func (o *POSTAddresses201ResponseDataAttributes) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *POSTAddresses201ResponseDataAttributes) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *POSTAddresses201ResponseDataAttributes) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetBillingInfo

`func (o *POSTAddresses201ResponseDataAttributes) GetBillingInfo() interface{}`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *POSTAddresses201ResponseDataAttributes) GetBillingInfoOk() (*interface{}, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *POSTAddresses201ResponseDataAttributes) SetBillingInfo(v interface{})`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *POSTAddresses201ResponseDataAttributes) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### SetBillingInfoNil

`func (o *POSTAddresses201ResponseDataAttributes) SetBillingInfoNil(b bool)`

 SetBillingInfoNil sets the value for BillingInfo to be an explicit nil

### UnsetBillingInfo
`func (o *POSTAddresses201ResponseDataAttributes) UnsetBillingInfo()`

UnsetBillingInfo ensures that no value is present for BillingInfo, not even an explicit nil
### GetReference

`func (o *POSTAddresses201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTAddresses201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTAddresses201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTAddresses201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTAddresses201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTAddresses201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTAddresses201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTAddresses201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTAddresses201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTAddresses201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTAddresses201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTAddresses201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTAddresses201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTAddresses201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTAddresses201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTAddresses201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTAddresses201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTAddresses201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


