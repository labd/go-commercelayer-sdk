# POSTBundles201ResponseDataAttributes

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

### NewPOSTBundles201ResponseDataAttributes

`func NewPOSTBundles201ResponseDataAttributes(code interface{}, name interface{}, priceAmountCents interface{}, compareAtAmountCents interface{}, ) *POSTBundles201ResponseDataAttributes`

NewPOSTBundles201ResponseDataAttributes instantiates a new POSTBundles201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTBundles201ResponseDataAttributesWithDefaults

`func NewPOSTBundles201ResponseDataAttributesWithDefaults() *POSTBundles201ResponseDataAttributes`

NewPOSTBundles201ResponseDataAttributesWithDefaults instantiates a new POSTBundles201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *POSTBundles201ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *POSTBundles201ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *POSTBundles201ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *POSTBundles201ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *POSTBundles201ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *POSTBundles201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTBundles201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTBundles201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTBundles201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTBundles201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *POSTBundles201ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *POSTBundles201ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *POSTBundles201ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *POSTBundles201ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *POSTBundles201ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *POSTBundles201ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDescription

`func (o *POSTBundles201ResponseDataAttributes) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *POSTBundles201ResponseDataAttributes) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *POSTBundles201ResponseDataAttributes) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *POSTBundles201ResponseDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *POSTBundles201ResponseDataAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *POSTBundles201ResponseDataAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *POSTBundles201ResponseDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *POSTBundles201ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *POSTBundles201ResponseDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *POSTBundles201ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *POSTBundles201ResponseDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *POSTBundles201ResponseDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetPriceAmountCents

`func (o *POSTBundles201ResponseDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *POSTBundles201ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *POSTBundles201ResponseDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.


### SetPriceAmountCentsNil

`func (o *POSTBundles201ResponseDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *POSTBundles201ResponseDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetCompareAtAmountCents

`func (o *POSTBundles201ResponseDataAttributes) GetCompareAtAmountCents() interface{}`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *POSTBundles201ResponseDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *POSTBundles201ResponseDataAttributes) SetCompareAtAmountCents(v interface{})`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.


### SetCompareAtAmountCentsNil

`func (o *POSTBundles201ResponseDataAttributes) SetCompareAtAmountCentsNil(b bool)`

 SetCompareAtAmountCentsNil sets the value for CompareAtAmountCents to be an explicit nil

### UnsetCompareAtAmountCents
`func (o *POSTBundles201ResponseDataAttributes) UnsetCompareAtAmountCents()`

UnsetCompareAtAmountCents ensures that no value is present for CompareAtAmountCents, not even an explicit nil
### GetComputePriceAmount

`func (o *POSTBundles201ResponseDataAttributes) GetComputePriceAmount() interface{}`

GetComputePriceAmount returns the ComputePriceAmount field if non-nil, zero value otherwise.

### GetComputePriceAmountOk

`func (o *POSTBundles201ResponseDataAttributes) GetComputePriceAmountOk() (*interface{}, bool)`

GetComputePriceAmountOk returns a tuple with the ComputePriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePriceAmount

`func (o *POSTBundles201ResponseDataAttributes) SetComputePriceAmount(v interface{})`

SetComputePriceAmount sets ComputePriceAmount field to given value.

### HasComputePriceAmount

`func (o *POSTBundles201ResponseDataAttributes) HasComputePriceAmount() bool`

HasComputePriceAmount returns a boolean if a field has been set.

### SetComputePriceAmountNil

`func (o *POSTBundles201ResponseDataAttributes) SetComputePriceAmountNil(b bool)`

 SetComputePriceAmountNil sets the value for ComputePriceAmount to be an explicit nil

### UnsetComputePriceAmount
`func (o *POSTBundles201ResponseDataAttributes) UnsetComputePriceAmount()`

UnsetComputePriceAmount ensures that no value is present for ComputePriceAmount, not even an explicit nil
### GetComputeCompareAtAmount

`func (o *POSTBundles201ResponseDataAttributes) GetComputeCompareAtAmount() interface{}`

GetComputeCompareAtAmount returns the ComputeCompareAtAmount field if non-nil, zero value otherwise.

### GetComputeCompareAtAmountOk

`func (o *POSTBundles201ResponseDataAttributes) GetComputeCompareAtAmountOk() (*interface{}, bool)`

GetComputeCompareAtAmountOk returns a tuple with the ComputeCompareAtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeCompareAtAmount

`func (o *POSTBundles201ResponseDataAttributes) SetComputeCompareAtAmount(v interface{})`

SetComputeCompareAtAmount sets ComputeCompareAtAmount field to given value.

### HasComputeCompareAtAmount

`func (o *POSTBundles201ResponseDataAttributes) HasComputeCompareAtAmount() bool`

HasComputeCompareAtAmount returns a boolean if a field has been set.

### SetComputeCompareAtAmountNil

`func (o *POSTBundles201ResponseDataAttributes) SetComputeCompareAtAmountNil(b bool)`

 SetComputeCompareAtAmountNil sets the value for ComputeCompareAtAmount to be an explicit nil

### UnsetComputeCompareAtAmount
`func (o *POSTBundles201ResponseDataAttributes) UnsetComputeCompareAtAmount()`

UnsetComputeCompareAtAmount ensures that no value is present for ComputeCompareAtAmount, not even an explicit nil
### GetReference

`func (o *POSTBundles201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTBundles201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTBundles201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTBundles201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTBundles201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTBundles201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTBundles201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTBundles201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTBundles201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTBundles201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTBundles201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTBundles201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTBundles201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTBundles201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTBundles201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTBundles201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTBundles201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTBundles201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


