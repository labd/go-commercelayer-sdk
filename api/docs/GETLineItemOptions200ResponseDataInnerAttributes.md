# GETLineItemOptions200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the line item option. When blank, it gets populated with the name of the associated SKU option. | [optional] 
**Quantity** | Pointer to **int32** | The line item option&#39;s quantity | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the order&#39;s market. | [optional] 
**UnitAmountCents** | Pointer to **int32** | The unit amount of the line item option, in cents. When you add a line item option to an order, this is automatically populated from associated SKU option&#39;s price. | [optional] 
**UnitAmountFloat** | Pointer to **float32** | The unit amount of the line item option, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedUnitAmount** | Pointer to **string** | The unit amount of the line item option, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**TotalAmountCents** | Pointer to **int32** | The unit amount x quantity, in cents. | [optional] 
**TotalAmountFloat** | Pointer to **float32** | The unit amount x quantity, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedTotalAmount** | Pointer to **string** | The unit amount x quantity, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**DelayHours** | Pointer to **int32** | The shipping delay that the customer can expect when adding this option (hours). Inherited from the associated SKU option. | [optional] 
**DelayDays** | Pointer to **int32** | The shipping delay that the customer can expect when adding this option (days, rounded). | [optional] 
**Options** | Pointer to **map[string]interface{}** | Set of key-value pairs that represent the selected options. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetUnitAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUnitAmountCents() int32`

GetUnitAmountCents returns the UnitAmountCents field if non-nil, zero value otherwise.

### GetUnitAmountCentsOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUnitAmountCentsOk() (*int32, bool)`

GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetUnitAmountCents(v int32)`

SetUnitAmountCents sets UnitAmountCents field to given value.

### HasUnitAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasUnitAmountCents() bool`

HasUnitAmountCents returns a boolean if a field has been set.

### GetUnitAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUnitAmountFloat() float32`

GetUnitAmountFloat returns the UnitAmountFloat field if non-nil, zero value otherwise.

### GetUnitAmountFloatOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUnitAmountFloatOk() (*float32, bool)`

GetUnitAmountFloatOk returns a tuple with the UnitAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetUnitAmountFloat(v float32)`

SetUnitAmountFloat sets UnitAmountFloat field to given value.

### HasUnitAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasUnitAmountFloat() bool`

HasUnitAmountFloat returns a boolean if a field has been set.

### GetFormattedUnitAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetFormattedUnitAmount() string`

GetFormattedUnitAmount returns the FormattedUnitAmount field if non-nil, zero value otherwise.

### GetFormattedUnitAmountOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetFormattedUnitAmountOk() (*string, bool)`

GetFormattedUnitAmountOk returns a tuple with the FormattedUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedUnitAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetFormattedUnitAmount(v string)`

SetFormattedUnitAmount sets FormattedUnitAmount field to given value.

### HasFormattedUnitAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasFormattedUnitAmount() bool`

HasFormattedUnitAmount returns a boolean if a field has been set.

### GetTotalAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetTotalAmountCents() int32`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetTotalAmountCentsOk() (*int32, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetTotalAmountCents(v int32)`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### GetTotalAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetTotalAmountFloat() float32`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetTotalAmountFloatOk() (*float32, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetTotalAmountFloat(v float32)`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### GetFormattedTotalAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetFormattedTotalAmount() string`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetFormattedTotalAmountOk() (*string, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetFormattedTotalAmount(v string)`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### GetDelayHours

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetDelayHours() int32`

GetDelayHours returns the DelayHours field if non-nil, zero value otherwise.

### GetDelayHoursOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetDelayHoursOk() (*int32, bool)`

GetDelayHoursOk returns a tuple with the DelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayHours

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetDelayHours(v int32)`

SetDelayHours sets DelayHours field to given value.

### HasDelayHours

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasDelayHours() bool`

HasDelayHours returns a boolean if a field has been set.

### GetDelayDays

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetDelayDays() int32`

GetDelayDays returns the DelayDays field if non-nil, zero value otherwise.

### GetDelayDaysOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetDelayDaysOk() (*int32, bool)`

GetDelayDaysOk returns a tuple with the DelayDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDays

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetDelayDays(v int32)`

SetDelayDays sets DelayDays field to given value.

### HasDelayDays

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasDelayDays() bool`

HasDelayDays returns a boolean if a field has been set.

### GetOptions

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetId

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETLineItemOptions200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


