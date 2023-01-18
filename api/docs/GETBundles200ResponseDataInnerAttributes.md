# GETBundles200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The bundle code, that uniquely identifies the bundle within the market. | [optional] 
**Name** | Pointer to **string** | The internal name of the bundle. | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Description** | Pointer to **string** | An internal description of the bundle. | [optional] 
**ImageUrl** | Pointer to **string** | The URL of an image that represents the bundle. | [optional] 
**DoNotShip** | Pointer to **bool** | Indicates if the bundle doesn&#39;t generate shipments (all sku_list&#39;s SKUs must be do_not_ship). | [optional] 
**DoNotTrack** | Pointer to **bool** | Indicates if the bundle doesn&#39;t track the stock inventory (all sku_list&#39;s SKUs must be do_not_track). | [optional] 
**PriceAmountCents** | Pointer to **int32** | The bundle price amount for the associated market, in cents. | [optional] 
**PriceAmountFloat** | Pointer to **float32** | The bundle price amount for the associated market, float. | [optional] 
**FormattedPriceAmount** | Pointer to **string** | The bundle price amount for the associated market, formatted. | [optional] 
**CompareAtAmountCents** | Pointer to **int32** | The compared price amount, in cents. Useful to display a percentage discount. | [optional] 
**CompareAtAmountFloat** | Pointer to **float32** | The compared price amount, float. | [optional] 
**FormattedCompareAtAmount** | Pointer to **string** | The compared price amount, formatted. | [optional] 
**SkusCount** | Pointer to **int32** | The total number of SKUs in the bundle. | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETBundles200ResponseDataInnerAttributes

`func NewGETBundles200ResponseDataInnerAttributes() *GETBundles200ResponseDataInnerAttributes`

NewGETBundles200ResponseDataInnerAttributes instantiates a new GETBundles200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETBundles200ResponseDataInnerAttributesWithDefaults

`func NewGETBundles200ResponseDataInnerAttributesWithDefaults() *GETBundles200ResponseDataInnerAttributes`

NewGETBundles200ResponseDataInnerAttributesWithDefaults instantiates a new GETBundles200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GETBundles200ResponseDataInnerAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GETBundles200ResponseDataInnerAttributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GETBundles200ResponseDataInnerAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GETBundles200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETBundles200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETBundles200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GETBundles200ResponseDataInnerAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETBundles200ResponseDataInnerAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETBundles200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDescription

`func (o *GETBundles200ResponseDataInnerAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GETBundles200ResponseDataInnerAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GETBundles200ResponseDataInnerAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageUrl

`func (o *GETBundles200ResponseDataInnerAttributes) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GETBundles200ResponseDataInnerAttributes) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GETBundles200ResponseDataInnerAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetDoNotShip

`func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotShip() bool`

GetDoNotShip returns the DoNotShip field if non-nil, zero value otherwise.

### GetDoNotShipOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotShipOk() (*bool, bool)`

GetDoNotShipOk returns a tuple with the DoNotShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotShip

`func (o *GETBundles200ResponseDataInnerAttributes) SetDoNotShip(v bool)`

SetDoNotShip sets DoNotShip field to given value.

### HasDoNotShip

`func (o *GETBundles200ResponseDataInnerAttributes) HasDoNotShip() bool`

HasDoNotShip returns a boolean if a field has been set.

### GetDoNotTrack

`func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotTrack() bool`

GetDoNotTrack returns the DoNotTrack field if non-nil, zero value otherwise.

### GetDoNotTrackOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotTrackOk() (*bool, bool)`

GetDoNotTrackOk returns a tuple with the DoNotTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrack

`func (o *GETBundles200ResponseDataInnerAttributes) SetDoNotTrack(v bool)`

SetDoNotTrack sets DoNotTrack field to given value.

### HasDoNotTrack

`func (o *GETBundles200ResponseDataInnerAttributes) HasDoNotTrack() bool`

HasDoNotTrack returns a boolean if a field has been set.

### GetPriceAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountCents() int32`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountCentsOk() (*int32, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) SetPriceAmountCents(v int32)`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### GetPriceAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountFloat() float32`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountFloatOk() (*float32, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) SetPriceAmountFloat(v float32)`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### GetFormattedPriceAmount

`func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedPriceAmount() string`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedPriceAmountOk() (*string, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *GETBundles200ResponseDataInnerAttributes) SetFormattedPriceAmount(v string)`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *GETBundles200ResponseDataInnerAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### GetCompareAtAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountCents() int32`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountCentsOk() (*int32, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) SetCompareAtAmountCents(v int32)`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.

### HasCompareAtAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) HasCompareAtAmountCents() bool`

HasCompareAtAmountCents returns a boolean if a field has been set.

### GetCompareAtAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountFloat() float32`

GetCompareAtAmountFloat returns the CompareAtAmountFloat field if non-nil, zero value otherwise.

### GetCompareAtAmountFloatOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountFloatOk() (*float32, bool)`

GetCompareAtAmountFloatOk returns a tuple with the CompareAtAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) SetCompareAtAmountFloat(v float32)`

SetCompareAtAmountFloat sets CompareAtAmountFloat field to given value.

### HasCompareAtAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) HasCompareAtAmountFloat() bool`

HasCompareAtAmountFloat returns a boolean if a field has been set.

### GetFormattedCompareAtAmount

`func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedCompareAtAmount() string`

GetFormattedCompareAtAmount returns the FormattedCompareAtAmount field if non-nil, zero value otherwise.

### GetFormattedCompareAtAmountOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedCompareAtAmountOk() (*string, bool)`

GetFormattedCompareAtAmountOk returns a tuple with the FormattedCompareAtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCompareAtAmount

`func (o *GETBundles200ResponseDataInnerAttributes) SetFormattedCompareAtAmount(v string)`

SetFormattedCompareAtAmount sets FormattedCompareAtAmount field to given value.

### HasFormattedCompareAtAmount

`func (o *GETBundles200ResponseDataInnerAttributes) HasFormattedCompareAtAmount() bool`

HasFormattedCompareAtAmount returns a boolean if a field has been set.

### GetSkusCount

`func (o *GETBundles200ResponseDataInnerAttributes) GetSkusCount() int32`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetSkusCountOk() (*int32, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *GETBundles200ResponseDataInnerAttributes) SetSkusCount(v int32)`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *GETBundles200ResponseDataInnerAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETBundles200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETBundles200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETBundles200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETBundles200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETBundles200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETBundles200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETBundles200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETBundles200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETBundles200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


