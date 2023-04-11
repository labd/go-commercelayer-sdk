# PATCHBundlesBundleIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **interface{}** | The bundle code, that uniquely identifies the bundle within the market. | [optional] 
**Name** | Pointer to **interface{}** | The internal name of the bundle. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Description** | Pointer to **interface{}** | An internal description of the bundle. | [optional] 
**ImageUrl** | Pointer to **interface{}** | The URL of an image that represents the bundle. | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The bundle price amount for the associated market, in cents. | [optional] 
**CompareAtAmountCents** | Pointer to **interface{}** | The compared price amount, in cents. Useful to display a percentage discount. | [optional] 
**ComputePriceAmount** | Pointer to **interface{}** | Send this attribute if you want to compute the price_amount_cents as the sum of the prices of the bundle SKUs for the market. | [optional] 
**ComputeCompareAtAmount** | Pointer to **interface{}** | Send this attribute if you want to compute the compare_at_amount_cents as the sum of the prices of the bundle SKUs for the market. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHBundlesBundleIdRequestDataAttributes

`func NewPATCHBundlesBundleIdRequestDataAttributes() *PATCHBundlesBundleIdRequestDataAttributes`

NewPATCHBundlesBundleIdRequestDataAttributes instantiates a new PATCHBundlesBundleIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHBundlesBundleIdRequestDataAttributesWithDefaults

`func NewPATCHBundlesBundleIdRequestDataAttributesWithDefaults() *PATCHBundlesBundleIdRequestDataAttributes`

NewPATCHBundlesBundleIdRequestDataAttributesWithDefaults instantiates a new PATCHBundlesBundleIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDescription

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetPriceAmountCents

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetCompareAtAmountCents

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetCompareAtAmountCents() interface{}`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetCompareAtAmountCents(v interface{})`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.

### HasCompareAtAmountCents

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasCompareAtAmountCents() bool`

HasCompareAtAmountCents returns a boolean if a field has been set.

### SetCompareAtAmountCentsNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetCompareAtAmountCentsNil(b bool)`

 SetCompareAtAmountCentsNil sets the value for CompareAtAmountCents to be an explicit nil

### UnsetCompareAtAmountCents
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetCompareAtAmountCents()`

UnsetCompareAtAmountCents ensures that no value is present for CompareAtAmountCents, not even an explicit nil
### GetComputePriceAmount

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetComputePriceAmount() interface{}`

GetComputePriceAmount returns the ComputePriceAmount field if non-nil, zero value otherwise.

### GetComputePriceAmountOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetComputePriceAmountOk() (*interface{}, bool)`

GetComputePriceAmountOk returns a tuple with the ComputePriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePriceAmount

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetComputePriceAmount(v interface{})`

SetComputePriceAmount sets ComputePriceAmount field to given value.

### HasComputePriceAmount

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasComputePriceAmount() bool`

HasComputePriceAmount returns a boolean if a field has been set.

### SetComputePriceAmountNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetComputePriceAmountNil(b bool)`

 SetComputePriceAmountNil sets the value for ComputePriceAmount to be an explicit nil

### UnsetComputePriceAmount
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetComputePriceAmount()`

UnsetComputePriceAmount ensures that no value is present for ComputePriceAmount, not even an explicit nil
### GetComputeCompareAtAmount

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetComputeCompareAtAmount() interface{}`

GetComputeCompareAtAmount returns the ComputeCompareAtAmount field if non-nil, zero value otherwise.

### GetComputeCompareAtAmountOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetComputeCompareAtAmountOk() (*interface{}, bool)`

GetComputeCompareAtAmountOk returns a tuple with the ComputeCompareAtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeCompareAtAmount

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetComputeCompareAtAmount(v interface{})`

SetComputeCompareAtAmount sets ComputeCompareAtAmount field to given value.

### HasComputeCompareAtAmount

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasComputeCompareAtAmount() bool`

HasComputeCompareAtAmount returns a boolean if a field has been set.

### SetComputeCompareAtAmountNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetComputeCompareAtAmountNil(b bool)`

 SetComputeCompareAtAmountNil sets the value for ComputeCompareAtAmount to be an explicit nil

### UnsetComputeCompareAtAmount
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetComputeCompareAtAmount()`

UnsetComputeCompareAtAmount ensures that no value is present for ComputeCompareAtAmount, not even an explicit nil
### GetReference

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHBundlesBundleIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHBundlesBundleIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHBundlesBundleIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHBundlesBundleIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


