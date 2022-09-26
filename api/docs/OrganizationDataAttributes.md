# OrganizationDataAttributes

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

### NewOrganizationDataAttributes

`func NewOrganizationDataAttributes() *OrganizationDataAttributes`

NewOrganizationDataAttributes instantiates a new OrganizationDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDataAttributesWithDefaults

`func NewOrganizationDataAttributesWithDefaults() *OrganizationDataAttributes`

NewOrganizationDataAttributesWithDefaults instantiates a new OrganizationDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *OrganizationDataAttributes) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OrganizationDataAttributes) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OrganizationDataAttributes) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *OrganizationDataAttributes) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDomain

`func (o *OrganizationDataAttributes) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OrganizationDataAttributes) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OrganizationDataAttributes) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *OrganizationDataAttributes) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSupportPhone

`func (o *OrganizationDataAttributes) GetSupportPhone() string`

GetSupportPhone returns the SupportPhone field if non-nil, zero value otherwise.

### GetSupportPhoneOk

`func (o *OrganizationDataAttributes) GetSupportPhoneOk() (*string, bool)`

GetSupportPhoneOk returns a tuple with the SupportPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhone

`func (o *OrganizationDataAttributes) SetSupportPhone(v string)`

SetSupportPhone sets SupportPhone field to given value.

### HasSupportPhone

`func (o *OrganizationDataAttributes) HasSupportPhone() bool`

HasSupportPhone returns a boolean if a field has been set.

### GetSupportEmail

`func (o *OrganizationDataAttributes) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *OrganizationDataAttributes) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *OrganizationDataAttributes) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *OrganizationDataAttributes) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### GetLogoUrl

`func (o *OrganizationDataAttributes) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *OrganizationDataAttributes) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *OrganizationDataAttributes) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *OrganizationDataAttributes) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetFaviconUrl

`func (o *OrganizationDataAttributes) GetFaviconUrl() string`

GetFaviconUrl returns the FaviconUrl field if non-nil, zero value otherwise.

### GetFaviconUrlOk

`func (o *OrganizationDataAttributes) GetFaviconUrlOk() (*string, bool)`

GetFaviconUrlOk returns a tuple with the FaviconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaviconUrl

`func (o *OrganizationDataAttributes) SetFaviconUrl(v string)`

SetFaviconUrl sets FaviconUrl field to given value.

### HasFaviconUrl

`func (o *OrganizationDataAttributes) HasFaviconUrl() bool`

HasFaviconUrl returns a boolean if a field has been set.

### GetPrimaryColor

`func (o *OrganizationDataAttributes) GetPrimaryColor() string`

GetPrimaryColor returns the PrimaryColor field if non-nil, zero value otherwise.

### GetPrimaryColorOk

`func (o *OrganizationDataAttributes) GetPrimaryColorOk() (*string, bool)`

GetPrimaryColorOk returns a tuple with the PrimaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColor

`func (o *OrganizationDataAttributes) SetPrimaryColor(v string)`

SetPrimaryColor sets PrimaryColor field to given value.

### HasPrimaryColor

`func (o *OrganizationDataAttributes) HasPrimaryColor() bool`

HasPrimaryColor returns a boolean if a field has been set.

### GetContrastColor

`func (o *OrganizationDataAttributes) GetContrastColor() string`

GetContrastColor returns the ContrastColor field if non-nil, zero value otherwise.

### GetContrastColorOk

`func (o *OrganizationDataAttributes) GetContrastColorOk() (*string, bool)`

GetContrastColorOk returns a tuple with the ContrastColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContrastColor

`func (o *OrganizationDataAttributes) SetContrastColor(v string)`

SetContrastColor sets ContrastColor field to given value.

### HasContrastColor

`func (o *OrganizationDataAttributes) HasContrastColor() bool`

HasContrastColor returns a boolean if a field has been set.

### GetGtmId

`func (o *OrganizationDataAttributes) GetGtmId() string`

GetGtmId returns the GtmId field if non-nil, zero value otherwise.

### GetGtmIdOk

`func (o *OrganizationDataAttributes) GetGtmIdOk() (*string, bool)`

GetGtmIdOk returns a tuple with the GtmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtmId

`func (o *OrganizationDataAttributes) SetGtmId(v string)`

SetGtmId sets GtmId field to given value.

### HasGtmId

`func (o *OrganizationDataAttributes) HasGtmId() bool`

HasGtmId returns a boolean if a field has been set.

### GetGtmIdTest

`func (o *OrganizationDataAttributes) GetGtmIdTest() string`

GetGtmIdTest returns the GtmIdTest field if non-nil, zero value otherwise.

### GetGtmIdTestOk

`func (o *OrganizationDataAttributes) GetGtmIdTestOk() (*string, bool)`

GetGtmIdTestOk returns a tuple with the GtmIdTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtmIdTest

`func (o *OrganizationDataAttributes) SetGtmIdTest(v string)`

SetGtmIdTest sets GtmIdTest field to given value.

### HasGtmIdTest

`func (o *OrganizationDataAttributes) HasGtmIdTest() bool`

HasGtmIdTest returns a boolean if a field has been set.

### GetDiscountDisabled

`func (o *OrganizationDataAttributes) GetDiscountDisabled() bool`

GetDiscountDisabled returns the DiscountDisabled field if non-nil, zero value otherwise.

### GetDiscountDisabledOk

`func (o *OrganizationDataAttributes) GetDiscountDisabledOk() (*bool, bool)`

GetDiscountDisabledOk returns a tuple with the DiscountDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDisabled

`func (o *OrganizationDataAttributes) SetDiscountDisabled(v bool)`

SetDiscountDisabled sets DiscountDisabled field to given value.

### HasDiscountDisabled

`func (o *OrganizationDataAttributes) HasDiscountDisabled() bool`

HasDiscountDisabled returns a boolean if a field has been set.

### GetAccountDisabled

`func (o *OrganizationDataAttributes) GetAccountDisabled() bool`

GetAccountDisabled returns the AccountDisabled field if non-nil, zero value otherwise.

### GetAccountDisabledOk

`func (o *OrganizationDataAttributes) GetAccountDisabledOk() (*bool, bool)`

GetAccountDisabledOk returns a tuple with the AccountDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDisabled

`func (o *OrganizationDataAttributes) SetAccountDisabled(v bool)`

SetAccountDisabled sets AccountDisabled field to given value.

### HasAccountDisabled

`func (o *OrganizationDataAttributes) HasAccountDisabled() bool`

HasAccountDisabled returns a boolean if a field has been set.

### GetAcceptanceDisabled

`func (o *OrganizationDataAttributes) GetAcceptanceDisabled() bool`

GetAcceptanceDisabled returns the AcceptanceDisabled field if non-nil, zero value otherwise.

### GetAcceptanceDisabledOk

`func (o *OrganizationDataAttributes) GetAcceptanceDisabledOk() (*bool, bool)`

GetAcceptanceDisabledOk returns a tuple with the AcceptanceDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceDisabled

`func (o *OrganizationDataAttributes) SetAcceptanceDisabled(v bool)`

SetAcceptanceDisabled sets AcceptanceDisabled field to given value.

### HasAcceptanceDisabled

`func (o *OrganizationDataAttributes) HasAcceptanceDisabled() bool`

HasAcceptanceDisabled returns a boolean if a field has been set.

### GetMaxConcurrentPromotions

`func (o *OrganizationDataAttributes) GetMaxConcurrentPromotions() int32`

GetMaxConcurrentPromotions returns the MaxConcurrentPromotions field if non-nil, zero value otherwise.

### GetMaxConcurrentPromotionsOk

`func (o *OrganizationDataAttributes) GetMaxConcurrentPromotionsOk() (*int32, bool)`

GetMaxConcurrentPromotionsOk returns a tuple with the MaxConcurrentPromotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentPromotions

`func (o *OrganizationDataAttributes) SetMaxConcurrentPromotions(v int32)`

SetMaxConcurrentPromotions sets MaxConcurrentPromotions field to given value.

### HasMaxConcurrentPromotions

`func (o *OrganizationDataAttributes) HasMaxConcurrentPromotions() bool`

HasMaxConcurrentPromotions returns a boolean if a field has been set.

### GetMaxConcurrentImports

`func (o *OrganizationDataAttributes) GetMaxConcurrentImports() int32`

GetMaxConcurrentImports returns the MaxConcurrentImports field if non-nil, zero value otherwise.

### GetMaxConcurrentImportsOk

`func (o *OrganizationDataAttributes) GetMaxConcurrentImportsOk() (*int32, bool)`

GetMaxConcurrentImportsOk returns a tuple with the MaxConcurrentImports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentImports

`func (o *OrganizationDataAttributes) SetMaxConcurrentImports(v int32)`

SetMaxConcurrentImports sets MaxConcurrentImports field to given value.

### HasMaxConcurrentImports

`func (o *OrganizationDataAttributes) HasMaxConcurrentImports() bool`

HasMaxConcurrentImports returns a boolean if a field has been set.

### GetId

`func (o *OrganizationDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganizationDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrganizationDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *OrganizationDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *OrganizationDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *OrganizationDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *OrganizationDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *OrganizationDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *OrganizationDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *OrganizationDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *OrganizationDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *OrganizationDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrganizationDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrganizationDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrganizationDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


