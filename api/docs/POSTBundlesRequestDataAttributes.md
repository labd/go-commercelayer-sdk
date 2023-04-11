# POSTBundlesRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **interface{}** | The bundle code, that uniquely identifies the bundle within the market. | 
**Name** | **interface{}** | The internal name of the bundle. | 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Description** | Pointer to **interface{}** | An internal description of the bundle. | [optional] 
**ImageUrl** | Pointer to **interface{}** | The URL of an image that represents the bundle. | [optional] 
**PriceAmountCents** | **interface{}** | The bundle price amount for the associated market, in cents. | 
**CompareAtAmountCents** | **interface{}** | The compared price amount, in cents. Useful to display a percentage discount. | 
**ComputePriceAmount** | Pointer to **interface{}** | Send this attribute if you want to compute the price_amount_cents as the sum of the prices of the bundle SKUs for the market. | [optional] 
**ComputeCompareAtAmount** | Pointer to **interface{}** | Send this attribute if you want to compute the compare_at_amount_cents as the sum of the prices of the bundle SKUs for the market. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTBundlesRequestDataAttributes

`func NewPOSTBundlesRequestDataAttributes(code interface{}, name interface{}, priceAmountCents interface{}, compareAtAmountCents interface{}, ) *POSTBundlesRequestDataAttributes`

NewPOSTBundlesRequestDataAttributes instantiates a new POSTBundlesRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTBundlesRequestDataAttributesWithDefaults

`func NewPOSTBundlesRequestDataAttributesWithDefaults() *POSTBundlesRequestDataAttributes`

NewPOSTBundlesRequestDataAttributesWithDefaults instantiates a new POSTBundlesRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *POSTBundlesRequestDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *POSTBundlesRequestDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *POSTBundlesRequestDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *POSTBundlesRequestDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *POSTBundlesRequestDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *POSTBundlesRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTBundlesRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTBundlesRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTBundlesRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTBundlesRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *POSTBundlesRequestDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *POSTBundlesRequestDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *POSTBundlesRequestDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *POSTBundlesRequestDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *POSTBundlesRequestDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *POSTBundlesRequestDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDescription

`func (o *POSTBundlesRequestDataAttributes) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *POSTBundlesRequestDataAttributes) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *POSTBundlesRequestDataAttributes) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *POSTBundlesRequestDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *POSTBundlesRequestDataAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *POSTBundlesRequestDataAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *POSTBundlesRequestDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *POSTBundlesRequestDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *POSTBundlesRequestDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *POSTBundlesRequestDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *POSTBundlesRequestDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *POSTBundlesRequestDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetPriceAmountCents

`func (o *POSTBundlesRequestDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *POSTBundlesRequestDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *POSTBundlesRequestDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.


### SetPriceAmountCentsNil

`func (o *POSTBundlesRequestDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *POSTBundlesRequestDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetCompareAtAmountCents

`func (o *POSTBundlesRequestDataAttributes) GetCompareAtAmountCents() interface{}`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *POSTBundlesRequestDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *POSTBundlesRequestDataAttributes) SetCompareAtAmountCents(v interface{})`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.


### SetCompareAtAmountCentsNil

`func (o *POSTBundlesRequestDataAttributes) SetCompareAtAmountCentsNil(b bool)`

 SetCompareAtAmountCentsNil sets the value for CompareAtAmountCents to be an explicit nil

### UnsetCompareAtAmountCents
`func (o *POSTBundlesRequestDataAttributes) UnsetCompareAtAmountCents()`

UnsetCompareAtAmountCents ensures that no value is present for CompareAtAmountCents, not even an explicit nil
### GetComputePriceAmount

`func (o *POSTBundlesRequestDataAttributes) GetComputePriceAmount() interface{}`

GetComputePriceAmount returns the ComputePriceAmount field if non-nil, zero value otherwise.

### GetComputePriceAmountOk

`func (o *POSTBundlesRequestDataAttributes) GetComputePriceAmountOk() (*interface{}, bool)`

GetComputePriceAmountOk returns a tuple with the ComputePriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePriceAmount

`func (o *POSTBundlesRequestDataAttributes) SetComputePriceAmount(v interface{})`

SetComputePriceAmount sets ComputePriceAmount field to given value.

### HasComputePriceAmount

`func (o *POSTBundlesRequestDataAttributes) HasComputePriceAmount() bool`

HasComputePriceAmount returns a boolean if a field has been set.

### SetComputePriceAmountNil

`func (o *POSTBundlesRequestDataAttributes) SetComputePriceAmountNil(b bool)`

 SetComputePriceAmountNil sets the value for ComputePriceAmount to be an explicit nil

### UnsetComputePriceAmount
`func (o *POSTBundlesRequestDataAttributes) UnsetComputePriceAmount()`

UnsetComputePriceAmount ensures that no value is present for ComputePriceAmount, not even an explicit nil
### GetComputeCompareAtAmount

`func (o *POSTBundlesRequestDataAttributes) GetComputeCompareAtAmount() interface{}`

GetComputeCompareAtAmount returns the ComputeCompareAtAmount field if non-nil, zero value otherwise.

### GetComputeCompareAtAmountOk

`func (o *POSTBundlesRequestDataAttributes) GetComputeCompareAtAmountOk() (*interface{}, bool)`

GetComputeCompareAtAmountOk returns a tuple with the ComputeCompareAtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeCompareAtAmount

`func (o *POSTBundlesRequestDataAttributes) SetComputeCompareAtAmount(v interface{})`

SetComputeCompareAtAmount sets ComputeCompareAtAmount field to given value.

### HasComputeCompareAtAmount

`func (o *POSTBundlesRequestDataAttributes) HasComputeCompareAtAmount() bool`

HasComputeCompareAtAmount returns a boolean if a field has been set.

### SetComputeCompareAtAmountNil

`func (o *POSTBundlesRequestDataAttributes) SetComputeCompareAtAmountNil(b bool)`

 SetComputeCompareAtAmountNil sets the value for ComputeCompareAtAmount to be an explicit nil

### UnsetComputeCompareAtAmount
`func (o *POSTBundlesRequestDataAttributes) UnsetComputeCompareAtAmount()`

UnsetComputeCompareAtAmount ensures that no value is present for ComputeCompareAtAmount, not even an explicit nil
### GetReference

`func (o *POSTBundlesRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTBundlesRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTBundlesRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTBundlesRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTBundlesRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTBundlesRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTBundlesRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTBundlesRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTBundlesRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTBundlesRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTBundlesRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTBundlesRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTBundlesRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTBundlesRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTBundlesRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTBundlesRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTBundlesRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTBundlesRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


