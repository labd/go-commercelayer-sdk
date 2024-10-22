# GETBundlesBundleId200ResponseDataAttributes

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
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETBundlesBundleId200ResponseDataAttributes

`func NewGETBundlesBundleId200ResponseDataAttributes() *GETBundlesBundleId200ResponseDataAttributes`

NewGETBundlesBundleId200ResponseDataAttributes instantiates a new GETBundlesBundleId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETBundlesBundleId200ResponseDataAttributesWithDefaults

`func NewGETBundlesBundleId200ResponseDataAttributesWithDefaults() *GETBundlesBundleId200ResponseDataAttributes`

NewGETBundlesBundleId200ResponseDataAttributesWithDefaults instantiates a new GETBundlesBundleId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDescription

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetDoNotShip

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetDoNotShip() interface{}`

GetDoNotShip returns the DoNotShip field if non-nil, zero value otherwise.

### GetDoNotShipOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetDoNotShipOk() (*interface{}, bool)`

GetDoNotShipOk returns a tuple with the DoNotShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotShip

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetDoNotShip(v interface{})`

SetDoNotShip sets DoNotShip field to given value.

### HasDoNotShip

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasDoNotShip() bool`

HasDoNotShip returns a boolean if a field has been set.

### SetDoNotShipNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetDoNotShipNil(b bool)`

 SetDoNotShipNil sets the value for DoNotShip to be an explicit nil

### UnsetDoNotShip
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetDoNotShip()`

UnsetDoNotShip ensures that no value is present for DoNotShip, not even an explicit nil
### GetDoNotTrack

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetDoNotTrack() interface{}`

GetDoNotTrack returns the DoNotTrack field if non-nil, zero value otherwise.

### GetDoNotTrackOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetDoNotTrackOk() (*interface{}, bool)`

GetDoNotTrackOk returns a tuple with the DoNotTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrack

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetDoNotTrack(v interface{})`

SetDoNotTrack sets DoNotTrack field to given value.

### HasDoNotTrack

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasDoNotTrack() bool`

HasDoNotTrack returns a boolean if a field has been set.

### SetDoNotTrackNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetDoNotTrackNil(b bool)`

 SetDoNotTrackNil sets the value for DoNotTrack to be an explicit nil

### UnsetDoNotTrack
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetDoNotTrack()`

UnsetDoNotTrack ensures that no value is present for DoNotTrack, not even an explicit nil
### GetPriceAmountCents

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetPriceAmountFloat

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetPriceAmountFloat() interface{}`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetPriceAmountFloatOk() (*interface{}, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetPriceAmountFloat(v interface{})`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### SetPriceAmountFloatNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetPriceAmountFloatNil(b bool)`

 SetPriceAmountFloatNil sets the value for PriceAmountFloat to be an explicit nil

### UnsetPriceAmountFloat
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetPriceAmountFloat()`

UnsetPriceAmountFloat ensures that no value is present for PriceAmountFloat, not even an explicit nil
### GetFormattedPriceAmount

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetFormattedPriceAmount() interface{}`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetFormattedPriceAmountOk() (*interface{}, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetFormattedPriceAmount(v interface{})`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### SetFormattedPriceAmountNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetFormattedPriceAmountNil(b bool)`

 SetFormattedPriceAmountNil sets the value for FormattedPriceAmount to be an explicit nil

### UnsetFormattedPriceAmount
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetFormattedPriceAmount()`

UnsetFormattedPriceAmount ensures that no value is present for FormattedPriceAmount, not even an explicit nil
### GetCompareAtAmountCents

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetCompareAtAmountCents() interface{}`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetCompareAtAmountCents(v interface{})`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.

### HasCompareAtAmountCents

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasCompareAtAmountCents() bool`

HasCompareAtAmountCents returns a boolean if a field has been set.

### SetCompareAtAmountCentsNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetCompareAtAmountCentsNil(b bool)`

 SetCompareAtAmountCentsNil sets the value for CompareAtAmountCents to be an explicit nil

### UnsetCompareAtAmountCents
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetCompareAtAmountCents()`

UnsetCompareAtAmountCents ensures that no value is present for CompareAtAmountCents, not even an explicit nil
### GetCompareAtAmountFloat

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetCompareAtAmountFloat() interface{}`

GetCompareAtAmountFloat returns the CompareAtAmountFloat field if non-nil, zero value otherwise.

### GetCompareAtAmountFloatOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetCompareAtAmountFloatOk() (*interface{}, bool)`

GetCompareAtAmountFloatOk returns a tuple with the CompareAtAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountFloat

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetCompareAtAmountFloat(v interface{})`

SetCompareAtAmountFloat sets CompareAtAmountFloat field to given value.

### HasCompareAtAmountFloat

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasCompareAtAmountFloat() bool`

HasCompareAtAmountFloat returns a boolean if a field has been set.

### SetCompareAtAmountFloatNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetCompareAtAmountFloatNil(b bool)`

 SetCompareAtAmountFloatNil sets the value for CompareAtAmountFloat to be an explicit nil

### UnsetCompareAtAmountFloat
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetCompareAtAmountFloat()`

UnsetCompareAtAmountFloat ensures that no value is present for CompareAtAmountFloat, not even an explicit nil
### GetFormattedCompareAtAmount

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetFormattedCompareAtAmount() interface{}`

GetFormattedCompareAtAmount returns the FormattedCompareAtAmount field if non-nil, zero value otherwise.

### GetFormattedCompareAtAmountOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetFormattedCompareAtAmountOk() (*interface{}, bool)`

GetFormattedCompareAtAmountOk returns a tuple with the FormattedCompareAtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCompareAtAmount

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetFormattedCompareAtAmount(v interface{})`

SetFormattedCompareAtAmount sets FormattedCompareAtAmount field to given value.

### HasFormattedCompareAtAmount

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasFormattedCompareAtAmount() bool`

HasFormattedCompareAtAmount returns a boolean if a field has been set.

### SetFormattedCompareAtAmountNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetFormattedCompareAtAmountNil(b bool)`

 SetFormattedCompareAtAmountNil sets the value for FormattedCompareAtAmount to be an explicit nil

### UnsetFormattedCompareAtAmount
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetFormattedCompareAtAmount()`

UnsetFormattedCompareAtAmount ensures that no value is present for FormattedCompareAtAmount, not even an explicit nil
### GetSkusCount

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetSkusCount() interface{}`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetSkusCountOk() (*interface{}, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetSkusCount(v interface{})`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### SetSkusCountNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetSkusCountNil(b bool)`

 SetSkusCountNil sets the value for SkusCount to be an explicit nil

### UnsetSkusCount
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetSkusCount()`

UnsetSkusCount ensures that no value is present for SkusCount, not even an explicit nil
### GetCreatedAt

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETBundlesBundleId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETBundlesBundleId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETBundlesBundleId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETBundlesBundleId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


