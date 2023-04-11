# GETSkuOptions200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The SKU option&#39;s internal name | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Description** | Pointer to **interface{}** | An internal description of the SKU option. | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The price of this shipping method, in cents. | [optional] 
**PriceAmountFloat** | Pointer to **interface{}** | The price of this shipping method, float. | [optional] 
**FormattedPriceAmount** | Pointer to **interface{}** | The price of this shipping method, formatted. | [optional] 
**DelayHours** | Pointer to **interface{}** | The delay time (in hours) that should be added to the delivery lead time when this option is purchased. | [optional] 
**DelayDays** | Pointer to **interface{}** | The delay time, in days (rounded) | [optional] 
**SkuCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the SKU codes. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETSkuOptions200ResponseDataInnerAttributes

`func NewGETSkuOptions200ResponseDataInnerAttributes() *GETSkuOptions200ResponseDataInnerAttributes`

NewGETSkuOptions200ResponseDataInnerAttributes instantiates a new GETSkuOptions200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETSkuOptions200ResponseDataInnerAttributesWithDefaults

`func NewGETSkuOptions200ResponseDataInnerAttributesWithDefaults() *GETSkuOptions200ResponseDataInnerAttributes`

NewGETSkuOptions200ResponseDataInnerAttributesWithDefaults instantiates a new GETSkuOptions200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDescription

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriceAmountCents

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetPriceAmountFloat

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetPriceAmountFloat() interface{}`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetPriceAmountFloatOk() (*interface{}, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetPriceAmountFloat(v interface{})`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### SetPriceAmountFloatNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetPriceAmountFloatNil(b bool)`

 SetPriceAmountFloatNil sets the value for PriceAmountFloat to be an explicit nil

### UnsetPriceAmountFloat
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetPriceAmountFloat()`

UnsetPriceAmountFloat ensures that no value is present for PriceAmountFloat, not even an explicit nil
### GetFormattedPriceAmount

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetFormattedPriceAmount() interface{}`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetFormattedPriceAmountOk() (*interface{}, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetFormattedPriceAmount(v interface{})`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### SetFormattedPriceAmountNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetFormattedPriceAmountNil(b bool)`

 SetFormattedPriceAmountNil sets the value for FormattedPriceAmount to be an explicit nil

### UnsetFormattedPriceAmount
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetFormattedPriceAmount()`

UnsetFormattedPriceAmount ensures that no value is present for FormattedPriceAmount, not even an explicit nil
### GetDelayHours

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDelayHours() interface{}`

GetDelayHours returns the DelayHours field if non-nil, zero value otherwise.

### GetDelayHoursOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDelayHoursOk() (*interface{}, bool)`

GetDelayHoursOk returns a tuple with the DelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayHours

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetDelayHours(v interface{})`

SetDelayHours sets DelayHours field to given value.

### HasDelayHours

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasDelayHours() bool`

HasDelayHours returns a boolean if a field has been set.

### SetDelayHoursNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetDelayHoursNil(b bool)`

 SetDelayHoursNil sets the value for DelayHours to be an explicit nil

### UnsetDelayHours
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetDelayHours()`

UnsetDelayHours ensures that no value is present for DelayHours, not even an explicit nil
### GetDelayDays

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDelayDays() interface{}`

GetDelayDays returns the DelayDays field if non-nil, zero value otherwise.

### GetDelayDaysOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDelayDaysOk() (*interface{}, bool)`

GetDelayDaysOk returns a tuple with the DelayDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDays

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetDelayDays(v interface{})`

SetDelayDays sets DelayDays field to given value.

### HasDelayDays

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasDelayDays() bool`

HasDelayDays returns a boolean if a field has been set.

### SetDelayDaysNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetDelayDaysNil(b bool)`

 SetDelayDaysNil sets the value for DelayDays to be an explicit nil

### UnsetDelayDays
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetDelayDays()`

UnsetDelayDays ensures that no value is present for DelayDays, not even an explicit nil
### GetSkuCodeRegex

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetSkuCodeRegex() interface{}`

GetSkuCodeRegex returns the SkuCodeRegex field if non-nil, zero value otherwise.

### GetSkuCodeRegexOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetSkuCodeRegexOk() (*interface{}, bool)`

GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCodeRegex

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetSkuCodeRegex(v interface{})`

SetSkuCodeRegex sets SkuCodeRegex field to given value.

### HasSkuCodeRegex

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasSkuCodeRegex() bool`

HasSkuCodeRegex returns a boolean if a field has been set.

### SetSkuCodeRegexNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetSkuCodeRegexNil(b bool)`

 SetSkuCodeRegexNil sets the value for SkuCodeRegex to be an explicit nil

### UnsetSkuCodeRegex
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetSkuCodeRegex()`

UnsetSkuCodeRegex ensures that no value is present for SkuCodeRegex, not even an explicit nil
### GetCreatedAt

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETSkuOptions200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETSkuOptions200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETSkuOptions200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETSkuOptions200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


