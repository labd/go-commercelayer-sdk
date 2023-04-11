# GETAddresses200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Business** | Pointer to **interface{}** | Indicates if it&#39;s a business or a personal address | [optional] 
**FirstName** | Pointer to **interface{}** | Address first name (personal) | [optional] 
**LastName** | Pointer to **interface{}** | Address last name (personal) | [optional] 
**Company** | Pointer to **interface{}** | Address company name (business) | [optional] 
**FullName** | Pointer to **interface{}** | Company name (business) of first name and last name (personal) | [optional] 
**Line1** | Pointer to **interface{}** | Address line 1, i.e. Street address, PO Box | [optional] 
**Line2** | Pointer to **interface{}** | Address line 2, i.e. Apartment, Suite, Building | [optional] 
**City** | Pointer to **interface{}** | Address city | [optional] 
**ZipCode** | Pointer to **interface{}** | ZIP or postal code | [optional] 
**StateCode** | Pointer to **interface{}** | State, province or region code. | [optional] 
**CountryCode** | Pointer to **interface{}** | The international 2-letter country code as defined by the ISO 3166-1 standard | [optional] 
**Phone** | Pointer to **interface{}** | Phone number (including extension). | [optional] 
**FullAddress** | Pointer to **interface{}** | Compact description of the address location, without the full name | [optional] 
**Name** | Pointer to **interface{}** | Compact description of the address location, including the full name | [optional] 
**Email** | Pointer to **interface{}** | Email address. | [optional] 
**Notes** | Pointer to **interface{}** | A free notes attached to the address. When used as a shipping address, this can be useful to let the customers add specific delivery instructions. | [optional] 
**Lat** | Pointer to **interface{}** | The address geocoded latitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**Lng** | Pointer to **interface{}** | The address geocoded longitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**IsLocalized** | Pointer to **interface{}** | Indicates if the latitude and logitude are present, either geocoded or manually updated | [optional] 
**IsGeocoded** | Pointer to **interface{}** | Indicates if the address has been successfully geocoded | [optional] 
**ProviderName** | Pointer to **interface{}** | The geocoder provider name (either Google or Bing) | [optional] 
**MapUrl** | Pointer to **interface{}** | The map url of the geocoded address (if geocoded) | [optional] 
**StaticMapUrl** | Pointer to **interface{}** | The static map image url of the geocoded address (if geocoded) | [optional] 
**BillingInfo** | Pointer to **interface{}** | Customer&#39;s billing information (i.e. VAT number, codice fiscale) | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETAddresses200ResponseDataInnerAttributes

`func NewGETAddresses200ResponseDataInnerAttributes() *GETAddresses200ResponseDataInnerAttributes`

NewGETAddresses200ResponseDataInnerAttributes instantiates a new GETAddresses200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETAddresses200ResponseDataInnerAttributesWithDefaults

`func NewGETAddresses200ResponseDataInnerAttributesWithDefaults() *GETAddresses200ResponseDataInnerAttributes`

NewGETAddresses200ResponseDataInnerAttributesWithDefaults instantiates a new GETAddresses200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusiness

`func (o *GETAddresses200ResponseDataInnerAttributes) GetBusiness() interface{}`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetBusinessOk() (*interface{}, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *GETAddresses200ResponseDataInnerAttributes) SetBusiness(v interface{})`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *GETAddresses200ResponseDataInnerAttributes) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### SetBusinessNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetBusinessNil(b bool)`

 SetBusinessNil sets the value for Business to be an explicit nil

### UnsetBusiness
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetBusiness()`

UnsetBusiness ensures that no value is present for Business, not even an explicit nil
### GetFirstName

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFirstName() interface{}`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFirstNameOk() (*interface{}, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GETAddresses200ResponseDataInnerAttributes) SetFirstName(v interface{})`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *GETAddresses200ResponseDataInnerAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLastName() interface{}`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLastNameOk() (*interface{}, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLastName(v interface{})`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *GETAddresses200ResponseDataInnerAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompany

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCompany() interface{}`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCompanyOk() (*interface{}, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCompany(v interface{})`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GETAddresses200ResponseDataInnerAttributes) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetFullName

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFullName() interface{}`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFullNameOk() (*interface{}, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GETAddresses200ResponseDataInnerAttributes) SetFullName(v interface{})`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GETAddresses200ResponseDataInnerAttributes) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetLine1

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLine1() interface{}`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLine1Ok() (*interface{}, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLine1(v interface{})`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *GETAddresses200ResponseDataInnerAttributes) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### SetLine1Nil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLine2() interface{}`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLine2Ok() (*interface{}, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLine2(v interface{})`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *GETAddresses200ResponseDataInnerAttributes) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetCity

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCity() interface{}`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCityOk() (*interface{}, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCity(v interface{})`

SetCity sets City field to given value.

### HasCity

`func (o *GETAddresses200ResponseDataInnerAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetZipCode

`func (o *GETAddresses200ResponseDataInnerAttributes) GetZipCode() interface{}`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetZipCodeOk() (*interface{}, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *GETAddresses200ResponseDataInnerAttributes) SetZipCode(v interface{})`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *GETAddresses200ResponseDataInnerAttributes) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetStateCode

`func (o *GETAddresses200ResponseDataInnerAttributes) GetStateCode() interface{}`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetStateCodeOk() (*interface{}, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *GETAddresses200ResponseDataInnerAttributes) SetStateCode(v interface{})`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *GETAddresses200ResponseDataInnerAttributes) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### SetStateCodeNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetStateCodeNil(b bool)`

 SetStateCodeNil sets the value for StateCode to be an explicit nil

### UnsetStateCode
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetStateCode()`

UnsetStateCode ensures that no value is present for StateCode, not even an explicit nil
### GetCountryCode

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCountryCode() interface{}`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCountryCodeOk() (*interface{}, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCountryCode(v interface{})`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GETAddresses200ResponseDataInnerAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetPhone

`func (o *GETAddresses200ResponseDataInnerAttributes) GetPhone() interface{}`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetPhoneOk() (*interface{}, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GETAddresses200ResponseDataInnerAttributes) SetPhone(v interface{})`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GETAddresses200ResponseDataInnerAttributes) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetFullAddress

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFullAddress() interface{}`

GetFullAddress returns the FullAddress field if non-nil, zero value otherwise.

### GetFullAddressOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFullAddressOk() (*interface{}, bool)`

GetFullAddressOk returns a tuple with the FullAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAddress

`func (o *GETAddresses200ResponseDataInnerAttributes) SetFullAddress(v interface{})`

SetFullAddress sets FullAddress field to given value.

### HasFullAddress

`func (o *GETAddresses200ResponseDataInnerAttributes) HasFullAddress() bool`

HasFullAddress returns a boolean if a field has been set.

### SetFullAddressNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetFullAddressNil(b bool)`

 SetFullAddressNil sets the value for FullAddress to be an explicit nil

### UnsetFullAddress
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetFullAddress()`

UnsetFullAddress ensures that no value is present for FullAddress, not even an explicit nil
### GetName

`func (o *GETAddresses200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETAddresses200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETAddresses200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *GETAddresses200ResponseDataInnerAttributes) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GETAddresses200ResponseDataInnerAttributes) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GETAddresses200ResponseDataInnerAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetNotes

`func (o *GETAddresses200ResponseDataInnerAttributes) GetNotes() interface{}`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetNotesOk() (*interface{}, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GETAddresses200ResponseDataInnerAttributes) SetNotes(v interface{})`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GETAddresses200ResponseDataInnerAttributes) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetLat

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLat() interface{}`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLatOk() (*interface{}, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLat(v interface{})`

SetLat sets Lat field to given value.

### HasLat

`func (o *GETAddresses200ResponseDataInnerAttributes) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLng() interface{}`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLngOk() (*interface{}, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLng(v interface{})`

SetLng sets Lng field to given value.

### HasLng

`func (o *GETAddresses200ResponseDataInnerAttributes) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetIsLocalized

`func (o *GETAddresses200ResponseDataInnerAttributes) GetIsLocalized() interface{}`

GetIsLocalized returns the IsLocalized field if non-nil, zero value otherwise.

### GetIsLocalizedOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetIsLocalizedOk() (*interface{}, bool)`

GetIsLocalizedOk returns a tuple with the IsLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalized

`func (o *GETAddresses200ResponseDataInnerAttributes) SetIsLocalized(v interface{})`

SetIsLocalized sets IsLocalized field to given value.

### HasIsLocalized

`func (o *GETAddresses200ResponseDataInnerAttributes) HasIsLocalized() bool`

HasIsLocalized returns a boolean if a field has been set.

### SetIsLocalizedNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetIsLocalizedNil(b bool)`

 SetIsLocalizedNil sets the value for IsLocalized to be an explicit nil

### UnsetIsLocalized
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetIsLocalized()`

UnsetIsLocalized ensures that no value is present for IsLocalized, not even an explicit nil
### GetIsGeocoded

`func (o *GETAddresses200ResponseDataInnerAttributes) GetIsGeocoded() interface{}`

GetIsGeocoded returns the IsGeocoded field if non-nil, zero value otherwise.

### GetIsGeocodedOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetIsGeocodedOk() (*interface{}, bool)`

GetIsGeocodedOk returns a tuple with the IsGeocoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGeocoded

`func (o *GETAddresses200ResponseDataInnerAttributes) SetIsGeocoded(v interface{})`

SetIsGeocoded sets IsGeocoded field to given value.

### HasIsGeocoded

`func (o *GETAddresses200ResponseDataInnerAttributes) HasIsGeocoded() bool`

HasIsGeocoded returns a boolean if a field has been set.

### SetIsGeocodedNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetIsGeocodedNil(b bool)`

 SetIsGeocodedNil sets the value for IsGeocoded to be an explicit nil

### UnsetIsGeocoded
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetIsGeocoded()`

UnsetIsGeocoded ensures that no value is present for IsGeocoded, not even an explicit nil
### GetProviderName

`func (o *GETAddresses200ResponseDataInnerAttributes) GetProviderName() interface{}`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetProviderNameOk() (*interface{}, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *GETAddresses200ResponseDataInnerAttributes) SetProviderName(v interface{})`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *GETAddresses200ResponseDataInnerAttributes) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) GetMapUrl() interface{}`

GetMapUrl returns the MapUrl field if non-nil, zero value otherwise.

### GetMapUrlOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetMapUrlOk() (*interface{}, bool)`

GetMapUrlOk returns a tuple with the MapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) SetMapUrl(v interface{})`

SetMapUrl sets MapUrl field to given value.

### HasMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) HasMapUrl() bool`

HasMapUrl returns a boolean if a field has been set.

### SetMapUrlNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetMapUrlNil(b bool)`

 SetMapUrlNil sets the value for MapUrl to be an explicit nil

### UnsetMapUrl
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetMapUrl()`

UnsetMapUrl ensures that no value is present for MapUrl, not even an explicit nil
### GetStaticMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) GetStaticMapUrl() interface{}`

GetStaticMapUrl returns the StaticMapUrl field if non-nil, zero value otherwise.

### GetStaticMapUrlOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetStaticMapUrlOk() (*interface{}, bool)`

GetStaticMapUrlOk returns a tuple with the StaticMapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) SetStaticMapUrl(v interface{})`

SetStaticMapUrl sets StaticMapUrl field to given value.

### HasStaticMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) HasStaticMapUrl() bool`

HasStaticMapUrl returns a boolean if a field has been set.

### SetStaticMapUrlNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetStaticMapUrlNil(b bool)`

 SetStaticMapUrlNil sets the value for StaticMapUrl to be an explicit nil

### UnsetStaticMapUrl
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetStaticMapUrl()`

UnsetStaticMapUrl ensures that no value is present for StaticMapUrl, not even an explicit nil
### GetBillingInfo

`func (o *GETAddresses200ResponseDataInnerAttributes) GetBillingInfo() interface{}`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetBillingInfoOk() (*interface{}, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *GETAddresses200ResponseDataInnerAttributes) SetBillingInfo(v interface{})`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *GETAddresses200ResponseDataInnerAttributes) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### SetBillingInfoNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetBillingInfoNil(b bool)`

 SetBillingInfoNil sets the value for BillingInfo to be an explicit nil

### UnsetBillingInfo
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetBillingInfo()`

UnsetBillingInfo ensures that no value is present for BillingInfo, not even an explicit nil
### GetCreatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETAddresses200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETAddresses200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETAddresses200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETAddresses200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETAddresses200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETAddresses200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETAddresses200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETAddresses200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETAddresses200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETAddresses200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETAddresses200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


