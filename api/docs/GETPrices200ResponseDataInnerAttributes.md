# GETPrices200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard, inherited from the associated price list. | [optional] 
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. When creating a price, either a valid sku_code or a SKU relationship must be present. | [optional] 
**AmountCents** | Pointer to **interface{}** | The SKU price amount for the associated price list, in cents. | [optional] 
**AmountFloat** | Pointer to **interface{}** | The SKU price amount for the associated price list, float. | [optional] 
**FormattedAmount** | Pointer to **interface{}** | The SKU price amount for the associated price list, formatted. | [optional] 
**CompareAtAmountCents** | Pointer to **interface{}** | The compared price amount, in cents. Useful to display a percentage discount. | [optional] 
**CompareAtAmountFloat** | Pointer to **interface{}** | The compared price amount, float. | [optional] 
**FormattedCompareAtAmount** | Pointer to **interface{}** | The compared price amount, formatted. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETPrices200ResponseDataInnerAttributes

`func NewGETPrices200ResponseDataInnerAttributes() *GETPrices200ResponseDataInnerAttributes`

NewGETPrices200ResponseDataInnerAttributes instantiates a new GETPrices200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPrices200ResponseDataInnerAttributesWithDefaults

`func NewGETPrices200ResponseDataInnerAttributesWithDefaults() *GETPrices200ResponseDataInnerAttributes`

NewGETPrices200ResponseDataInnerAttributesWithDefaults instantiates a new GETPrices200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *GETPrices200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETPrices200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETPrices200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetSkuCode

`func (o *GETPrices200ResponseDataInnerAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *GETPrices200ResponseDataInnerAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *GETPrices200ResponseDataInnerAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetAmountCents

`func (o *GETPrices200ResponseDataInnerAttributes) GetAmountCents() interface{}`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetAmountCentsOk() (*interface{}, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GETPrices200ResponseDataInnerAttributes) SetAmountCents(v interface{})`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GETPrices200ResponseDataInnerAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountFloat

`func (o *GETPrices200ResponseDataInnerAttributes) GetAmountFloat() interface{}`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetAmountFloatOk() (*interface{}, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *GETPrices200ResponseDataInnerAttributes) SetAmountFloat(v interface{})`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *GETPrices200ResponseDataInnerAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### SetAmountFloatNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetAmountFloatNil(b bool)`

 SetAmountFloatNil sets the value for AmountFloat to be an explicit nil

### UnsetAmountFloat
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetAmountFloat()`

UnsetAmountFloat ensures that no value is present for AmountFloat, not even an explicit nil
### GetFormattedAmount

`func (o *GETPrices200ResponseDataInnerAttributes) GetFormattedAmount() interface{}`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetFormattedAmountOk() (*interface{}, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *GETPrices200ResponseDataInnerAttributes) SetFormattedAmount(v interface{})`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *GETPrices200ResponseDataInnerAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### SetFormattedAmountNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetFormattedAmountNil(b bool)`

 SetFormattedAmountNil sets the value for FormattedAmount to be an explicit nil

### UnsetFormattedAmount
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetFormattedAmount()`

UnsetFormattedAmount ensures that no value is present for FormattedAmount, not even an explicit nil
### GetCompareAtAmountCents

`func (o *GETPrices200ResponseDataInnerAttributes) GetCompareAtAmountCents() interface{}`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *GETPrices200ResponseDataInnerAttributes) SetCompareAtAmountCents(v interface{})`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.

### HasCompareAtAmountCents

`func (o *GETPrices200ResponseDataInnerAttributes) HasCompareAtAmountCents() bool`

HasCompareAtAmountCents returns a boolean if a field has been set.

### SetCompareAtAmountCentsNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetCompareAtAmountCentsNil(b bool)`

 SetCompareAtAmountCentsNil sets the value for CompareAtAmountCents to be an explicit nil

### UnsetCompareAtAmountCents
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetCompareAtAmountCents()`

UnsetCompareAtAmountCents ensures that no value is present for CompareAtAmountCents, not even an explicit nil
### GetCompareAtAmountFloat

`func (o *GETPrices200ResponseDataInnerAttributes) GetCompareAtAmountFloat() interface{}`

GetCompareAtAmountFloat returns the CompareAtAmountFloat field if non-nil, zero value otherwise.

### GetCompareAtAmountFloatOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetCompareAtAmountFloatOk() (*interface{}, bool)`

GetCompareAtAmountFloatOk returns a tuple with the CompareAtAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountFloat

`func (o *GETPrices200ResponseDataInnerAttributes) SetCompareAtAmountFloat(v interface{})`

SetCompareAtAmountFloat sets CompareAtAmountFloat field to given value.

### HasCompareAtAmountFloat

`func (o *GETPrices200ResponseDataInnerAttributes) HasCompareAtAmountFloat() bool`

HasCompareAtAmountFloat returns a boolean if a field has been set.

### SetCompareAtAmountFloatNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetCompareAtAmountFloatNil(b bool)`

 SetCompareAtAmountFloatNil sets the value for CompareAtAmountFloat to be an explicit nil

### UnsetCompareAtAmountFloat
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetCompareAtAmountFloat()`

UnsetCompareAtAmountFloat ensures that no value is present for CompareAtAmountFloat, not even an explicit nil
### GetFormattedCompareAtAmount

`func (o *GETPrices200ResponseDataInnerAttributes) GetFormattedCompareAtAmount() interface{}`

GetFormattedCompareAtAmount returns the FormattedCompareAtAmount field if non-nil, zero value otherwise.

### GetFormattedCompareAtAmountOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetFormattedCompareAtAmountOk() (*interface{}, bool)`

GetFormattedCompareAtAmountOk returns a tuple with the FormattedCompareAtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCompareAtAmount

`func (o *GETPrices200ResponseDataInnerAttributes) SetFormattedCompareAtAmount(v interface{})`

SetFormattedCompareAtAmount sets FormattedCompareAtAmount field to given value.

### HasFormattedCompareAtAmount

`func (o *GETPrices200ResponseDataInnerAttributes) HasFormattedCompareAtAmount() bool`

HasFormattedCompareAtAmount returns a boolean if a field has been set.

### SetFormattedCompareAtAmountNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetFormattedCompareAtAmountNil(b bool)`

 SetFormattedCompareAtAmountNil sets the value for FormattedCompareAtAmount to be an explicit nil

### UnsetFormattedCompareAtAmount
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetFormattedCompareAtAmount()`

UnsetFormattedCompareAtAmount ensures that no value is present for FormattedCompareAtAmount, not even an explicit nil
### GetCreatedAt

`func (o *GETPrices200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPrices200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPrices200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPrices200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPrices200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPrices200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPrices200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPrices200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPrices200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPrices200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPrices200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPrices200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETPrices200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPrices200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPrices200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPrices200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPrices200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPrices200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


