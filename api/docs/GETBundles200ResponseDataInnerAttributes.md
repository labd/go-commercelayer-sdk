# GETBundles200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **interface{}** | The bundle code, that uniquely identifies the bundle within the market. | [optional] 
**Name** | Pointer to **interface{}** | The internal name of the bundle. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Description** | Pointer to **interface{}** | An internal description of the bundle. | [optional] 
**ImageUrl** | Pointer to **interface{}** | The URL of an image that represents the bundle. | [optional] 
**DoNotShip** | Pointer to **interface{}** | Indicates if the bundle doesn&#39;t generate shipments (all sku_list&#39;s SKUs must be do_not_ship). | [optional] 
**DoNotTrack** | Pointer to **interface{}** | Indicates if the bundle doesn&#39;t track the stock inventory (all sku_list&#39;s SKUs must be do_not_track). | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The bundle price amount for the associated market, in cents. | [optional] 
**PriceAmountFloat** | Pointer to **interface{}** | The bundle price amount for the associated market, float. | [optional] 
**FormattedPriceAmount** | Pointer to **interface{}** | The bundle price amount for the associated market, formatted. | [optional] 
**CompareAtAmountCents** | Pointer to **interface{}** | The compared price amount, in cents. Useful to display a percentage discount. | [optional] 
**CompareAtAmountFloat** | Pointer to **interface{}** | The compared price amount, float. | [optional] 
**FormattedCompareAtAmount** | Pointer to **interface{}** | The compared price amount, formatted. | [optional] 
**SkusCount** | Pointer to **interface{}** | The total number of SKUs in the bundle. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETBundles200ResponseDataInnerAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GETBundles200ResponseDataInnerAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *GETBundles200ResponseDataInnerAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *GETBundles200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETBundles200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETBundles200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *GETBundles200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETBundles200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETBundles200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDescription

`func (o *GETBundles200ResponseDataInnerAttributes) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GETBundles200ResponseDataInnerAttributes) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GETBundles200ResponseDataInnerAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *GETBundles200ResponseDataInnerAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GETBundles200ResponseDataInnerAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GETBundles200ResponseDataInnerAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetDoNotShip

`func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotShip() interface{}`

GetDoNotShip returns the DoNotShip field if non-nil, zero value otherwise.

### GetDoNotShipOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotShipOk() (*interface{}, bool)`

GetDoNotShipOk returns a tuple with the DoNotShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotShip

`func (o *GETBundles200ResponseDataInnerAttributes) SetDoNotShip(v interface{})`

SetDoNotShip sets DoNotShip field to given value.

### HasDoNotShip

`func (o *GETBundles200ResponseDataInnerAttributes) HasDoNotShip() bool`

HasDoNotShip returns a boolean if a field has been set.

### SetDoNotShipNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetDoNotShipNil(b bool)`

 SetDoNotShipNil sets the value for DoNotShip to be an explicit nil

### UnsetDoNotShip
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetDoNotShip()`

UnsetDoNotShip ensures that no value is present for DoNotShip, not even an explicit nil
### GetDoNotTrack

`func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotTrack() interface{}`

GetDoNotTrack returns the DoNotTrack field if non-nil, zero value otherwise.

### GetDoNotTrackOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotTrackOk() (*interface{}, bool)`

GetDoNotTrackOk returns a tuple with the DoNotTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrack

`func (o *GETBundles200ResponseDataInnerAttributes) SetDoNotTrack(v interface{})`

SetDoNotTrack sets DoNotTrack field to given value.

### HasDoNotTrack

`func (o *GETBundles200ResponseDataInnerAttributes) HasDoNotTrack() bool`

HasDoNotTrack returns a boolean if a field has been set.

### SetDoNotTrackNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetDoNotTrackNil(b bool)`

 SetDoNotTrackNil sets the value for DoNotTrack to be an explicit nil

### UnsetDoNotTrack
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetDoNotTrack()`

UnsetDoNotTrack ensures that no value is present for DoNotTrack, not even an explicit nil
### GetPriceAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetPriceAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountFloat() interface{}`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountFloatOk() (*interface{}, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) SetPriceAmountFloat(v interface{})`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### SetPriceAmountFloatNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetPriceAmountFloatNil(b bool)`

 SetPriceAmountFloatNil sets the value for PriceAmountFloat to be an explicit nil

### UnsetPriceAmountFloat
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetPriceAmountFloat()`

UnsetPriceAmountFloat ensures that no value is present for PriceAmountFloat, not even an explicit nil
### GetFormattedPriceAmount

`func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedPriceAmount() interface{}`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedPriceAmountOk() (*interface{}, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *GETBundles200ResponseDataInnerAttributes) SetFormattedPriceAmount(v interface{})`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *GETBundles200ResponseDataInnerAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### SetFormattedPriceAmountNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetFormattedPriceAmountNil(b bool)`

 SetFormattedPriceAmountNil sets the value for FormattedPriceAmount to be an explicit nil

### UnsetFormattedPriceAmount
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetFormattedPriceAmount()`

UnsetFormattedPriceAmount ensures that no value is present for FormattedPriceAmount, not even an explicit nil
### GetCompareAtAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountCents() interface{}`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) SetCompareAtAmountCents(v interface{})`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.

### HasCompareAtAmountCents

`func (o *GETBundles200ResponseDataInnerAttributes) HasCompareAtAmountCents() bool`

HasCompareAtAmountCents returns a boolean if a field has been set.

### SetCompareAtAmountCentsNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetCompareAtAmountCentsNil(b bool)`

 SetCompareAtAmountCentsNil sets the value for CompareAtAmountCents to be an explicit nil

### UnsetCompareAtAmountCents
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetCompareAtAmountCents()`

UnsetCompareAtAmountCents ensures that no value is present for CompareAtAmountCents, not even an explicit nil
### GetCompareAtAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountFloat() interface{}`

GetCompareAtAmountFloat returns the CompareAtAmountFloat field if non-nil, zero value otherwise.

### GetCompareAtAmountFloatOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountFloatOk() (*interface{}, bool)`

GetCompareAtAmountFloatOk returns a tuple with the CompareAtAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) SetCompareAtAmountFloat(v interface{})`

SetCompareAtAmountFloat sets CompareAtAmountFloat field to given value.

### HasCompareAtAmountFloat

`func (o *GETBundles200ResponseDataInnerAttributes) HasCompareAtAmountFloat() bool`

HasCompareAtAmountFloat returns a boolean if a field has been set.

### SetCompareAtAmountFloatNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetCompareAtAmountFloatNil(b bool)`

 SetCompareAtAmountFloatNil sets the value for CompareAtAmountFloat to be an explicit nil

### UnsetCompareAtAmountFloat
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetCompareAtAmountFloat()`

UnsetCompareAtAmountFloat ensures that no value is present for CompareAtAmountFloat, not even an explicit nil
### GetFormattedCompareAtAmount

`func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedCompareAtAmount() interface{}`

GetFormattedCompareAtAmount returns the FormattedCompareAtAmount field if non-nil, zero value otherwise.

### GetFormattedCompareAtAmountOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedCompareAtAmountOk() (*interface{}, bool)`

GetFormattedCompareAtAmountOk returns a tuple with the FormattedCompareAtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCompareAtAmount

`func (o *GETBundles200ResponseDataInnerAttributes) SetFormattedCompareAtAmount(v interface{})`

SetFormattedCompareAtAmount sets FormattedCompareAtAmount field to given value.

### HasFormattedCompareAtAmount

`func (o *GETBundles200ResponseDataInnerAttributes) HasFormattedCompareAtAmount() bool`

HasFormattedCompareAtAmount returns a boolean if a field has been set.

### SetFormattedCompareAtAmountNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetFormattedCompareAtAmountNil(b bool)`

 SetFormattedCompareAtAmountNil sets the value for FormattedCompareAtAmount to be an explicit nil

### UnsetFormattedCompareAtAmount
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetFormattedCompareAtAmount()`

UnsetFormattedCompareAtAmount ensures that no value is present for FormattedCompareAtAmount, not even an explicit nil
### GetSkusCount

`func (o *GETBundles200ResponseDataInnerAttributes) GetSkusCount() interface{}`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetSkusCountOk() (*interface{}, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *GETBundles200ResponseDataInnerAttributes) SetSkusCount(v interface{})`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *GETBundles200ResponseDataInnerAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### SetSkusCountNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetSkusCountNil(b bool)`

 SetSkusCountNil sets the value for SkusCount to be an explicit nil

### UnsetSkusCount
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetSkusCount()`

UnsetSkusCount ensures that no value is present for SkusCount, not even an explicit nil
### GetCreatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETBundles200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETBundles200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETBundles200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETBundles200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETBundles200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETBundles200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETBundles200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETBundles200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETBundles200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETBundles200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETBundles200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETBundles200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETBundles200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


