# GETAddresses200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Business** | Pointer to **bool** | Indicates if it&#39;s a business or a personal address | [optional] 
**FirstName** | Pointer to **string** | Address first name (personal) | [optional] 
**LastName** | Pointer to **string** | Address last name (personal) | [optional] 
**Company** | Pointer to **string** | Address company name (business) | [optional] 
**FullName** | Pointer to **string** | Company name (business) of first name and last name (personal) | [optional] 
**Line1** | Pointer to **string** | Address line 1, i.e. Street address, PO Box | [optional] 
**Line2** | Pointer to **string** | Address line 2, i.e. Apartment, Suite, Building | [optional] 
**City** | Pointer to **string** | Address city | [optional] 
**ZipCode** | Pointer to **string** | ZIP or postal code | [optional] 
**StateCode** | Pointer to **string** | State, province or region code. | [optional] 
**CountryCode** | Pointer to **string** | The international 2-letter country code as defined by the ISO 3166-1 standard | [optional] 
**Phone** | Pointer to **string** | Phone number (including extension). | [optional] 
**FullAddress** | Pointer to **string** | Compact description of the address location, without the full name | [optional] 
**Name** | Pointer to **string** | Compact description of the address location, including the full name | [optional] 
**Email** | Pointer to **string** | Email address. | [optional] 
**Notes** | Pointer to **string** | A free notes attached to the address. When used as a shipping address, this can be useful to let the customers add specific delivery instructions. | [optional] 
**Lat** | Pointer to **float32** | The address geocoded latitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**Lng** | Pointer to **float32** | The address geocoded longitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order&#39;s market. | [optional] 
**IsLocalized** | Pointer to **bool** | Indicates if the latitude and logitude are present, either geocoded or manually updated | [optional] 
**IsGeocoded** | Pointer to **bool** | Indicates if the address has been successfully geocoded | [optional] 
**ProviderName** | Pointer to **string** | The geocoder provider name (either Google or Bing) | [optional] 
**MapUrl** | Pointer to **string** | The map url of the geocoded address (if geocoded) | [optional] 
**StaticMapUrl** | Pointer to **string** | The static map image url of the geocoded address (if geocoded) | [optional] 
**BillingInfo** | Pointer to **string** | Customer&#39;s billing information (i.e. VAT number, codice fiscale) | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETAddresses200ResponseDataInnerAttributes) GetBusiness() bool`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetBusinessOk() (*bool, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *GETAddresses200ResponseDataInnerAttributes) SetBusiness(v bool)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *GETAddresses200ResponseDataInnerAttributes) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetFirstName

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GETAddresses200ResponseDataInnerAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *GETAddresses200ResponseDataInnerAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *GETAddresses200ResponseDataInnerAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompany

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GETAddresses200ResponseDataInnerAttributes) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetFullName

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GETAddresses200ResponseDataInnerAttributes) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GETAddresses200ResponseDataInnerAttributes) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetLine1

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *GETAddresses200ResponseDataInnerAttributes) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *GETAddresses200ResponseDataInnerAttributes) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetCity

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GETAddresses200ResponseDataInnerAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetZipCode

`func (o *GETAddresses200ResponseDataInnerAttributes) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *GETAddresses200ResponseDataInnerAttributes) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *GETAddresses200ResponseDataInnerAttributes) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetStateCode

`func (o *GETAddresses200ResponseDataInnerAttributes) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *GETAddresses200ResponseDataInnerAttributes) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *GETAddresses200ResponseDataInnerAttributes) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GETAddresses200ResponseDataInnerAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhone

`func (o *GETAddresses200ResponseDataInnerAttributes) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GETAddresses200ResponseDataInnerAttributes) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GETAddresses200ResponseDataInnerAttributes) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetFullAddress

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFullAddress() string`

GetFullAddress returns the FullAddress field if non-nil, zero value otherwise.

### GetFullAddressOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetFullAddressOk() (*string, bool)`

GetFullAddressOk returns a tuple with the FullAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAddress

`func (o *GETAddresses200ResponseDataInnerAttributes) SetFullAddress(v string)`

SetFullAddress sets FullAddress field to given value.

### HasFullAddress

`func (o *GETAddresses200ResponseDataInnerAttributes) HasFullAddress() bool`

HasFullAddress returns a boolean if a field has been set.

### GetName

`func (o *GETAddresses200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETAddresses200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETAddresses200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *GETAddresses200ResponseDataInnerAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GETAddresses200ResponseDataInnerAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GETAddresses200ResponseDataInnerAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNotes

`func (o *GETAddresses200ResponseDataInnerAttributes) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GETAddresses200ResponseDataInnerAttributes) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GETAddresses200ResponseDataInnerAttributes) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetLat

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *GETAddresses200ResponseDataInnerAttributes) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *GETAddresses200ResponseDataInnerAttributes) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *GETAddresses200ResponseDataInnerAttributes) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetIsLocalized

`func (o *GETAddresses200ResponseDataInnerAttributes) GetIsLocalized() bool`

GetIsLocalized returns the IsLocalized field if non-nil, zero value otherwise.

### GetIsLocalizedOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetIsLocalizedOk() (*bool, bool)`

GetIsLocalizedOk returns a tuple with the IsLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalized

`func (o *GETAddresses200ResponseDataInnerAttributes) SetIsLocalized(v bool)`

SetIsLocalized sets IsLocalized field to given value.

### HasIsLocalized

`func (o *GETAddresses200ResponseDataInnerAttributes) HasIsLocalized() bool`

HasIsLocalized returns a boolean if a field has been set.

### GetIsGeocoded

`func (o *GETAddresses200ResponseDataInnerAttributes) GetIsGeocoded() bool`

GetIsGeocoded returns the IsGeocoded field if non-nil, zero value otherwise.

### GetIsGeocodedOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetIsGeocodedOk() (*bool, bool)`

GetIsGeocodedOk returns a tuple with the IsGeocoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGeocoded

`func (o *GETAddresses200ResponseDataInnerAttributes) SetIsGeocoded(v bool)`

SetIsGeocoded sets IsGeocoded field to given value.

### HasIsGeocoded

`func (o *GETAddresses200ResponseDataInnerAttributes) HasIsGeocoded() bool`

HasIsGeocoded returns a boolean if a field has been set.

### GetProviderName

`func (o *GETAddresses200ResponseDataInnerAttributes) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *GETAddresses200ResponseDataInnerAttributes) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *GETAddresses200ResponseDataInnerAttributes) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) GetMapUrl() string`

GetMapUrl returns the MapUrl field if non-nil, zero value otherwise.

### GetMapUrlOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetMapUrlOk() (*string, bool)`

GetMapUrlOk returns a tuple with the MapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) SetMapUrl(v string)`

SetMapUrl sets MapUrl field to given value.

### HasMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) HasMapUrl() bool`

HasMapUrl returns a boolean if a field has been set.

### GetStaticMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) GetStaticMapUrl() string`

GetStaticMapUrl returns the StaticMapUrl field if non-nil, zero value otherwise.

### GetStaticMapUrlOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetStaticMapUrlOk() (*string, bool)`

GetStaticMapUrlOk returns a tuple with the StaticMapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) SetStaticMapUrl(v string)`

SetStaticMapUrl sets StaticMapUrl field to given value.

### HasStaticMapUrl

`func (o *GETAddresses200ResponseDataInnerAttributes) HasStaticMapUrl() bool`

HasStaticMapUrl returns a boolean if a field has been set.

### GetBillingInfo

`func (o *GETAddresses200ResponseDataInnerAttributes) GetBillingInfo() string`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetBillingInfoOk() (*string, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *GETAddresses200ResponseDataInnerAttributes) SetBillingInfo(v string)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *GETAddresses200ResponseDataInnerAttributes) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetId

`func (o *GETAddresses200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETAddresses200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETAddresses200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETAddresses200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETAddresses200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETAddresses200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETAddresses200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETAddresses200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETAddresses200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETAddresses200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETAddresses200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETAddresses200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETAddresses200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETAddresses200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


