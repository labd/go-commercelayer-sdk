# AddressDataAttributes

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

### NewAddressDataAttributes

`func NewAddressDataAttributes() *AddressDataAttributes`

NewAddressDataAttributes instantiates a new AddressDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressDataAttributesWithDefaults

`func NewAddressDataAttributesWithDefaults() *AddressDataAttributes`

NewAddressDataAttributesWithDefaults instantiates a new AddressDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusiness

`func (o *AddressDataAttributes) GetBusiness() bool`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *AddressDataAttributes) GetBusinessOk() (*bool, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *AddressDataAttributes) SetBusiness(v bool)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *AddressDataAttributes) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetFirstName

`func (o *AddressDataAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddressDataAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddressDataAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AddressDataAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AddressDataAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddressDataAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddressDataAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AddressDataAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompany

`func (o *AddressDataAttributes) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressDataAttributes) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressDataAttributes) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressDataAttributes) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetFullName

`func (o *AddressDataAttributes) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AddressDataAttributes) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AddressDataAttributes) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AddressDataAttributes) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetLine1

`func (o *AddressDataAttributes) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *AddressDataAttributes) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *AddressDataAttributes) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *AddressDataAttributes) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *AddressDataAttributes) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *AddressDataAttributes) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *AddressDataAttributes) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *AddressDataAttributes) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetCity

`func (o *AddressDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressDataAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressDataAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetZipCode

`func (o *AddressDataAttributes) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *AddressDataAttributes) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *AddressDataAttributes) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *AddressDataAttributes) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetStateCode

`func (o *AddressDataAttributes) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *AddressDataAttributes) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *AddressDataAttributes) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *AddressDataAttributes) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressDataAttributes) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressDataAttributes) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressDataAttributes) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AddressDataAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhone

`func (o *AddressDataAttributes) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressDataAttributes) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressDataAttributes) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressDataAttributes) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetFullAddress

`func (o *AddressDataAttributes) GetFullAddress() string`

GetFullAddress returns the FullAddress field if non-nil, zero value otherwise.

### GetFullAddressOk

`func (o *AddressDataAttributes) GetFullAddressOk() (*string, bool)`

GetFullAddressOk returns a tuple with the FullAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAddress

`func (o *AddressDataAttributes) SetFullAddress(v string)`

SetFullAddress sets FullAddress field to given value.

### HasFullAddress

`func (o *AddressDataAttributes) HasFullAddress() bool`

HasFullAddress returns a boolean if a field has been set.

### GetName

`func (o *AddressDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *AddressDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNotes

`func (o *AddressDataAttributes) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AddressDataAttributes) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AddressDataAttributes) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AddressDataAttributes) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetLat

`func (o *AddressDataAttributes) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *AddressDataAttributes) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *AddressDataAttributes) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *AddressDataAttributes) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *AddressDataAttributes) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *AddressDataAttributes) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *AddressDataAttributes) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *AddressDataAttributes) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetIsLocalized

`func (o *AddressDataAttributes) GetIsLocalized() bool`

GetIsLocalized returns the IsLocalized field if non-nil, zero value otherwise.

### GetIsLocalizedOk

`func (o *AddressDataAttributes) GetIsLocalizedOk() (*bool, bool)`

GetIsLocalizedOk returns a tuple with the IsLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalized

`func (o *AddressDataAttributes) SetIsLocalized(v bool)`

SetIsLocalized sets IsLocalized field to given value.

### HasIsLocalized

`func (o *AddressDataAttributes) HasIsLocalized() bool`

HasIsLocalized returns a boolean if a field has been set.

### GetIsGeocoded

`func (o *AddressDataAttributes) GetIsGeocoded() bool`

GetIsGeocoded returns the IsGeocoded field if non-nil, zero value otherwise.

### GetIsGeocodedOk

`func (o *AddressDataAttributes) GetIsGeocodedOk() (*bool, bool)`

GetIsGeocodedOk returns a tuple with the IsGeocoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGeocoded

`func (o *AddressDataAttributes) SetIsGeocoded(v bool)`

SetIsGeocoded sets IsGeocoded field to given value.

### HasIsGeocoded

`func (o *AddressDataAttributes) HasIsGeocoded() bool`

HasIsGeocoded returns a boolean if a field has been set.

### GetProviderName

`func (o *AddressDataAttributes) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddressDataAttributes) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddressDataAttributes) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *AddressDataAttributes) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetMapUrl

`func (o *AddressDataAttributes) GetMapUrl() string`

GetMapUrl returns the MapUrl field if non-nil, zero value otherwise.

### GetMapUrlOk

`func (o *AddressDataAttributes) GetMapUrlOk() (*string, bool)`

GetMapUrlOk returns a tuple with the MapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapUrl

`func (o *AddressDataAttributes) SetMapUrl(v string)`

SetMapUrl sets MapUrl field to given value.

### HasMapUrl

`func (o *AddressDataAttributes) HasMapUrl() bool`

HasMapUrl returns a boolean if a field has been set.

### GetStaticMapUrl

`func (o *AddressDataAttributes) GetStaticMapUrl() string`

GetStaticMapUrl returns the StaticMapUrl field if non-nil, zero value otherwise.

### GetStaticMapUrlOk

`func (o *AddressDataAttributes) GetStaticMapUrlOk() (*string, bool)`

GetStaticMapUrlOk returns a tuple with the StaticMapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMapUrl

`func (o *AddressDataAttributes) SetStaticMapUrl(v string)`

SetStaticMapUrl sets StaticMapUrl field to given value.

### HasStaticMapUrl

`func (o *AddressDataAttributes) HasStaticMapUrl() bool`

HasStaticMapUrl returns a boolean if a field has been set.

### GetBillingInfo

`func (o *AddressDataAttributes) GetBillingInfo() string`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *AddressDataAttributes) GetBillingInfoOk() (*string, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *AddressDataAttributes) SetBillingInfo(v string)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *AddressDataAttributes) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetId

`func (o *AddressDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddressDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AddressDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AddressDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AddressDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AddressDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AddressDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AddressDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AddressDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AddressDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *AddressDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AddressDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AddressDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AddressDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *AddressDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *AddressDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *AddressDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *AddressDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *AddressDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddressDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddressDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddressDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


