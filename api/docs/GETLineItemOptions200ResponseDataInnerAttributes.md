# GETLineItemOptions200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The name of the line item option. When blank, it gets populated with the name of the associated SKU option. | [optional] 
**Quantity** | Pointer to **interface{}** | The line item option&#39;s quantity | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the order&#39;s market. | [optional] 
**UnitAmountCents** | Pointer to **interface{}** | The unit amount of the line item option, in cents. When you add a line item option to an order, this is automatically populated from associated SKU option&#39;s price. | [optional] 
**UnitAmountFloat** | Pointer to **interface{}** | The unit amount of the line item option, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedUnitAmount** | Pointer to **interface{}** | The unit amount of the line item option, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**TotalAmountCents** | Pointer to **interface{}** | The unit amount x quantity, in cents. | [optional] 
**TotalAmountFloat** | Pointer to **interface{}** | The unit amount x quantity, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedTotalAmount** | Pointer to **interface{}** | The unit amount x quantity, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**DelayHours** | Pointer to **interface{}** | The shipping delay that the customer can expect when adding this option (hours). Inherited from the associated SKU option. | [optional] 
**DelayDays** | Pointer to **interface{}** | The shipping delay that the customer can expect when adding this option (days, rounded). | [optional] 
**Options** | Pointer to **interface{}** | Set of key-value pairs that represent the selected options. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETLineItemOptions200ResponseDataInnerAttributes

`func NewGETLineItemOptions200ResponseDataInnerAttributes() *GETLineItemOptions200ResponseDataInnerAttributes`

NewGETLineItemOptions200ResponseDataInnerAttributes instantiates a new GETLineItemOptions200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETLineItemOptions200ResponseDataInnerAttributesWithDefaults

`func NewGETLineItemOptions200ResponseDataInnerAttributesWithDefaults() *GETLineItemOptions200ResponseDataInnerAttributes`

NewGETLineItemOptions200ResponseDataInnerAttributesWithDefaults instantiates a new GETLineItemOptions200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuantity

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetCurrencyCode

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetUnitAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUnitAmountCents() interface{}`

GetUnitAmountCents returns the UnitAmountCents field if non-nil, zero value otherwise.

### GetUnitAmountCentsOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUnitAmountCentsOk() (*interface{}, bool)`

GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetUnitAmountCents(v interface{})`

SetUnitAmountCents sets UnitAmountCents field to given value.

### HasUnitAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasUnitAmountCents() bool`

HasUnitAmountCents returns a boolean if a field has been set.

### SetUnitAmountCentsNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetUnitAmountCentsNil(b bool)`

 SetUnitAmountCentsNil sets the value for UnitAmountCents to be an explicit nil

### UnsetUnitAmountCents
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetUnitAmountCents()`

UnsetUnitAmountCents ensures that no value is present for UnitAmountCents, not even an explicit nil
### GetUnitAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUnitAmountFloat() interface{}`

GetUnitAmountFloat returns the UnitAmountFloat field if non-nil, zero value otherwise.

### GetUnitAmountFloatOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUnitAmountFloatOk() (*interface{}, bool)`

GetUnitAmountFloatOk returns a tuple with the UnitAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetUnitAmountFloat(v interface{})`

SetUnitAmountFloat sets UnitAmountFloat field to given value.

### HasUnitAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasUnitAmountFloat() bool`

HasUnitAmountFloat returns a boolean if a field has been set.

### SetUnitAmountFloatNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetUnitAmountFloatNil(b bool)`

 SetUnitAmountFloatNil sets the value for UnitAmountFloat to be an explicit nil

### UnsetUnitAmountFloat
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetUnitAmountFloat()`

UnsetUnitAmountFloat ensures that no value is present for UnitAmountFloat, not even an explicit nil
### GetFormattedUnitAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetFormattedUnitAmount() interface{}`

GetFormattedUnitAmount returns the FormattedUnitAmount field if non-nil, zero value otherwise.

### GetFormattedUnitAmountOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetFormattedUnitAmountOk() (*interface{}, bool)`

GetFormattedUnitAmountOk returns a tuple with the FormattedUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedUnitAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetFormattedUnitAmount(v interface{})`

SetFormattedUnitAmount sets FormattedUnitAmount field to given value.

### HasFormattedUnitAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasFormattedUnitAmount() bool`

HasFormattedUnitAmount returns a boolean if a field has been set.

### SetFormattedUnitAmountNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetFormattedUnitAmountNil(b bool)`

 SetFormattedUnitAmountNil sets the value for FormattedUnitAmount to be an explicit nil

### UnsetFormattedUnitAmount
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetFormattedUnitAmount()`

UnsetFormattedUnitAmount ensures that no value is present for FormattedUnitAmount, not even an explicit nil
### GetTotalAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetTotalAmountCents() interface{}`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetTotalAmountCentsOk() (*interface{}, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetTotalAmountCents(v interface{})`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### SetTotalAmountCentsNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetTotalAmountCentsNil(b bool)`

 SetTotalAmountCentsNil sets the value for TotalAmountCents to be an explicit nil

### UnsetTotalAmountCents
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetTotalAmountCents()`

UnsetTotalAmountCents ensures that no value is present for TotalAmountCents, not even an explicit nil
### GetTotalAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetTotalAmountFloat() interface{}`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetTotalAmountFloatOk() (*interface{}, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetTotalAmountFloat(v interface{})`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### SetTotalAmountFloatNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetTotalAmountFloatNil(b bool)`

 SetTotalAmountFloatNil sets the value for TotalAmountFloat to be an explicit nil

### UnsetTotalAmountFloat
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetTotalAmountFloat()`

UnsetTotalAmountFloat ensures that no value is present for TotalAmountFloat, not even an explicit nil
### GetFormattedTotalAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetFormattedTotalAmount() interface{}`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetFormattedTotalAmountOk() (*interface{}, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetFormattedTotalAmount(v interface{})`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### SetFormattedTotalAmountNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetFormattedTotalAmountNil(b bool)`

 SetFormattedTotalAmountNil sets the value for FormattedTotalAmount to be an explicit nil

### UnsetFormattedTotalAmount
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetFormattedTotalAmount()`

UnsetFormattedTotalAmount ensures that no value is present for FormattedTotalAmount, not even an explicit nil
### GetDelayHours

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetDelayHours() interface{}`

GetDelayHours returns the DelayHours field if non-nil, zero value otherwise.

### GetDelayHoursOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetDelayHoursOk() (*interface{}, bool)`

GetDelayHoursOk returns a tuple with the DelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayHours

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetDelayHours(v interface{})`

SetDelayHours sets DelayHours field to given value.

### HasDelayHours

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasDelayHours() bool`

HasDelayHours returns a boolean if a field has been set.

### SetDelayHoursNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetDelayHoursNil(b bool)`

 SetDelayHoursNil sets the value for DelayHours to be an explicit nil

### UnsetDelayHours
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetDelayHours()`

UnsetDelayHours ensures that no value is present for DelayHours, not even an explicit nil
### GetDelayDays

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetDelayDays() interface{}`

GetDelayDays returns the DelayDays field if non-nil, zero value otherwise.

### GetDelayDaysOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetDelayDaysOk() (*interface{}, bool)`

GetDelayDaysOk returns a tuple with the DelayDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDays

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetDelayDays(v interface{})`

SetDelayDays sets DelayDays field to given value.

### HasDelayDays

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasDelayDays() bool`

HasDelayDays returns a boolean if a field has been set.

### SetDelayDaysNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetDelayDaysNil(b bool)`

 SetDelayDaysNil sets the value for DelayDays to be an explicit nil

### UnsetDelayDays
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetDelayDays()`

UnsetDelayDays ensures that no value is present for DelayDays, not even an explicit nil
### GetOptions

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetCreatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETLineItemOptions200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


