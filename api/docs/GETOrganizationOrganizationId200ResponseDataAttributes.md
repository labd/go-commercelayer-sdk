# GETOrganizationOrganizationId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The organization&#39;s internal name. | [optional] 
**Slug** | Pointer to **string** | The organization&#39;s slug name. | [optional] 
**Domain** | Pointer to **string** | The organization&#39;s domain. | [optional] 
**SupportPhone** | Pointer to **string** | The organization&#39;s support phone. | [optional] 
**SupportEmail** | Pointer to **string** | The organization&#39;s support email. | [optional] 
**LogoUrl** | Pointer to **string** | The URL to the organization&#39;s logo. | [optional] 
**FaviconUrl** | Pointer to **string** | The URL to the organization&#39;s favicon. | [optional] 
**PrimaryColor** | Pointer to **string** | The organization&#39;s primary color. | [optional] 
**ContrastColor** | Pointer to **string** | The organization&#39;s contrast color. | [optional] 
**GtmId** | Pointer to **string** | The organization&#39;s Google Tag Manager ID. | [optional] 
**GtmIdTest** | Pointer to **string** | The organization&#39;s Google Tag Manager ID for test. | [optional] 
**DiscountDisabled** | Pointer to **bool** | Indicates if organization has discount disabled. | [optional] 
**AccountDisabled** | Pointer to **bool** | Indicates if organization has account disabled. | [optional] 
**AcceptanceDisabled** | Pointer to **bool** | Indicates if organization has acceptance disabled. | [optional] 
**MaxConcurrentPromotions** | Pointer to **int32** | The maximum number of active concurrent promotions allowed for your organization. | [optional] 
**MaxConcurrentImports** | Pointer to **int32** | The maximum number of concurrent imports allowed for your organization. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETOrganizationOrganizationId200ResponseDataAttributes

`func NewGETOrganizationOrganizationId200ResponseDataAttributes() *GETOrganizationOrganizationId200ResponseDataAttributes`

NewGETOrganizationOrganizationId200ResponseDataAttributes instantiates a new GETOrganizationOrganizationId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrganizationOrganizationId200ResponseDataAttributesWithDefaults

`func NewGETOrganizationOrganizationId200ResponseDataAttributesWithDefaults() *GETOrganizationOrganizationId200ResponseDataAttributes`

NewGETOrganizationOrganizationId200ResponseDataAttributesWithDefaults instantiates a new GETOrganizationOrganizationId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDomain

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSupportPhone

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportPhone() string`

GetSupportPhone returns the SupportPhone field if non-nil, zero value otherwise.

### GetSupportPhoneOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportPhoneOk() (*string, bool)`

GetSupportPhoneOk returns a tuple with the SupportPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhone

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSupportPhone(v string)`

SetSupportPhone sets SupportPhone field to given value.

### HasSupportPhone

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasSupportPhone() bool`

HasSupportPhone returns a boolean if a field has been set.

### GetSupportEmail

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### GetLogoUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetFaviconUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetFaviconUrl() string`

GetFaviconUrl returns the FaviconUrl field if non-nil, zero value otherwise.

### GetFaviconUrlOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetFaviconUrlOk() (*string, bool)`

GetFaviconUrlOk returns a tuple with the FaviconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaviconUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetFaviconUrl(v string)`

SetFaviconUrl sets FaviconUrl field to given value.

### HasFaviconUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasFaviconUrl() bool`

HasFaviconUrl returns a boolean if a field has been set.

### GetPrimaryColor

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetPrimaryColor() string`

GetPrimaryColor returns the PrimaryColor field if non-nil, zero value otherwise.

### GetPrimaryColorOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetPrimaryColorOk() (*string, bool)`

GetPrimaryColorOk returns a tuple with the PrimaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColor

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetPrimaryColor(v string)`

SetPrimaryColor sets PrimaryColor field to given value.

### HasPrimaryColor

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasPrimaryColor() bool`

HasPrimaryColor returns a boolean if a field has been set.

### GetContrastColor

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetContrastColor() string`

GetContrastColor returns the ContrastColor field if non-nil, zero value otherwise.

### GetContrastColorOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetContrastColorOk() (*string, bool)`

GetContrastColorOk returns a tuple with the ContrastColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContrastColor

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetContrastColor(v string)`

SetContrastColor sets ContrastColor field to given value.

### HasContrastColor

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasContrastColor() bool`

HasContrastColor returns a boolean if a field has been set.

### GetGtmId

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmId() string`

GetGtmId returns the GtmId field if non-nil, zero value otherwise.

### GetGtmIdOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmIdOk() (*string, bool)`

GetGtmIdOk returns a tuple with the GtmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtmId

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGtmId(v string)`

SetGtmId sets GtmId field to given value.

### HasGtmId

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasGtmId() bool`

HasGtmId returns a boolean if a field has been set.

### GetGtmIdTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmIdTest() string`

GetGtmIdTest returns the GtmIdTest field if non-nil, zero value otherwise.

### GetGtmIdTestOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmIdTestOk() (*string, bool)`

GetGtmIdTestOk returns a tuple with the GtmIdTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtmIdTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGtmIdTest(v string)`

SetGtmIdTest sets GtmIdTest field to given value.

### HasGtmIdTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasGtmIdTest() bool`

HasGtmIdTest returns a boolean if a field has been set.

### GetDiscountDisabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDiscountDisabled() bool`

GetDiscountDisabled returns the DiscountDisabled field if non-nil, zero value otherwise.

### GetDiscountDisabledOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDiscountDisabledOk() (*bool, bool)`

GetDiscountDisabledOk returns a tuple with the DiscountDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDisabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetDiscountDisabled(v bool)`

SetDiscountDisabled sets DiscountDisabled field to given value.

### HasDiscountDisabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasDiscountDisabled() bool`

HasDiscountDisabled returns a boolean if a field has been set.

### GetAccountDisabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetAccountDisabled() bool`

GetAccountDisabled returns the AccountDisabled field if non-nil, zero value otherwise.

### GetAccountDisabledOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetAccountDisabledOk() (*bool, bool)`

GetAccountDisabledOk returns a tuple with the AccountDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDisabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetAccountDisabled(v bool)`

SetAccountDisabled sets AccountDisabled field to given value.

### HasAccountDisabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasAccountDisabled() bool`

HasAccountDisabled returns a boolean if a field has been set.

### GetAcceptanceDisabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetAcceptanceDisabled() bool`

GetAcceptanceDisabled returns the AcceptanceDisabled field if non-nil, zero value otherwise.

### GetAcceptanceDisabledOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetAcceptanceDisabledOk() (*bool, bool)`

GetAcceptanceDisabledOk returns a tuple with the AcceptanceDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceDisabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetAcceptanceDisabled(v bool)`

SetAcceptanceDisabled sets AcceptanceDisabled field to given value.

### HasAcceptanceDisabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasAcceptanceDisabled() bool`

HasAcceptanceDisabled returns a boolean if a field has been set.

### GetMaxConcurrentPromotions

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMaxConcurrentPromotions() int32`

GetMaxConcurrentPromotions returns the MaxConcurrentPromotions field if non-nil, zero value otherwise.

### GetMaxConcurrentPromotionsOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMaxConcurrentPromotionsOk() (*int32, bool)`

GetMaxConcurrentPromotionsOk returns a tuple with the MaxConcurrentPromotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentPromotions

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetMaxConcurrentPromotions(v int32)`

SetMaxConcurrentPromotions sets MaxConcurrentPromotions field to given value.

### HasMaxConcurrentPromotions

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasMaxConcurrentPromotions() bool`

HasMaxConcurrentPromotions returns a boolean if a field has been set.

### GetMaxConcurrentImports

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMaxConcurrentImports() int32`

GetMaxConcurrentImports returns the MaxConcurrentImports field if non-nil, zero value otherwise.

### GetMaxConcurrentImportsOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMaxConcurrentImportsOk() (*int32, bool)`

GetMaxConcurrentImportsOk returns a tuple with the MaxConcurrentImports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentImports

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetMaxConcurrentImports(v int32)`

SetMaxConcurrentImports sets MaxConcurrentImports field to given value.

### HasMaxConcurrentImports

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasMaxConcurrentImports() bool`

HasMaxConcurrentImports returns a boolean if a field has been set.

### GetId

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


