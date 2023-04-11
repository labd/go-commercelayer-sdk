# GETOrderSubscriptionItems200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **interface{}** | The subscription item quantity. | [optional] 
**UnitAmountCents** | Pointer to **interface{}** | The unit amount of the subscription item, in cents. | [optional] 
**UnitAmountFloat** | Pointer to **interface{}** | The unit amount of the subscription item, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedUnitAmount** | Pointer to **interface{}** | The unit amount of the subscription item, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**TotalAmountCents** | Pointer to **interface{}** | Calculated as unit amount x quantity amount, in cents. | [optional] 
**TotalAmountFloat** | Pointer to **interface{}** | Calculated as unit amount x quantity amount, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedTotalAmount** | Pointer to **interface{}** | Calculated as unit amount x quantity amount, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETOrderSubscriptionItems200ResponseDataInnerAttributes

`func NewGETOrderSubscriptionItems200ResponseDataInnerAttributes() *GETOrderSubscriptionItems200ResponseDataInnerAttributes`

NewGETOrderSubscriptionItems200ResponseDataInnerAttributes instantiates a new GETOrderSubscriptionItems200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrderSubscriptionItems200ResponseDataInnerAttributesWithDefaults

`func NewGETOrderSubscriptionItems200ResponseDataInnerAttributesWithDefaults() *GETOrderSubscriptionItems200ResponseDataInnerAttributes`

NewGETOrderSubscriptionItems200ResponseDataInnerAttributesWithDefaults instantiates a new GETOrderSubscriptionItems200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetUnitAmountCents

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetUnitAmountCents() interface{}`

GetUnitAmountCents returns the UnitAmountCents field if non-nil, zero value otherwise.

### GetUnitAmountCentsOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetUnitAmountCentsOk() (*interface{}, bool)`

GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountCents

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetUnitAmountCents(v interface{})`

SetUnitAmountCents sets UnitAmountCents field to given value.

### HasUnitAmountCents

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasUnitAmountCents() bool`

HasUnitAmountCents returns a boolean if a field has been set.

### SetUnitAmountCentsNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetUnitAmountCentsNil(b bool)`

 SetUnitAmountCentsNil sets the value for UnitAmountCents to be an explicit nil

### UnsetUnitAmountCents
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetUnitAmountCents()`

UnsetUnitAmountCents ensures that no value is present for UnitAmountCents, not even an explicit nil
### GetUnitAmountFloat

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetUnitAmountFloat() interface{}`

GetUnitAmountFloat returns the UnitAmountFloat field if non-nil, zero value otherwise.

### GetUnitAmountFloatOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetUnitAmountFloatOk() (*interface{}, bool)`

GetUnitAmountFloatOk returns a tuple with the UnitAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountFloat

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetUnitAmountFloat(v interface{})`

SetUnitAmountFloat sets UnitAmountFloat field to given value.

### HasUnitAmountFloat

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasUnitAmountFloat() bool`

HasUnitAmountFloat returns a boolean if a field has been set.

### SetUnitAmountFloatNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetUnitAmountFloatNil(b bool)`

 SetUnitAmountFloatNil sets the value for UnitAmountFloat to be an explicit nil

### UnsetUnitAmountFloat
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetUnitAmountFloat()`

UnsetUnitAmountFloat ensures that no value is present for UnitAmountFloat, not even an explicit nil
### GetFormattedUnitAmount

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetFormattedUnitAmount() interface{}`

GetFormattedUnitAmount returns the FormattedUnitAmount field if non-nil, zero value otherwise.

### GetFormattedUnitAmountOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetFormattedUnitAmountOk() (*interface{}, bool)`

GetFormattedUnitAmountOk returns a tuple with the FormattedUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedUnitAmount

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetFormattedUnitAmount(v interface{})`

SetFormattedUnitAmount sets FormattedUnitAmount field to given value.

### HasFormattedUnitAmount

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasFormattedUnitAmount() bool`

HasFormattedUnitAmount returns a boolean if a field has been set.

### SetFormattedUnitAmountNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetFormattedUnitAmountNil(b bool)`

 SetFormattedUnitAmountNil sets the value for FormattedUnitAmount to be an explicit nil

### UnsetFormattedUnitAmount
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetFormattedUnitAmount()`

UnsetFormattedUnitAmount ensures that no value is present for FormattedUnitAmount, not even an explicit nil
### GetTotalAmountCents

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetTotalAmountCents() interface{}`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetTotalAmountCentsOk() (*interface{}, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetTotalAmountCents(v interface{})`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### SetTotalAmountCentsNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetTotalAmountCentsNil(b bool)`

 SetTotalAmountCentsNil sets the value for TotalAmountCents to be an explicit nil

### UnsetTotalAmountCents
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetTotalAmountCents()`

UnsetTotalAmountCents ensures that no value is present for TotalAmountCents, not even an explicit nil
### GetTotalAmountFloat

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetTotalAmountFloat() interface{}`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetTotalAmountFloatOk() (*interface{}, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetTotalAmountFloat(v interface{})`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### SetTotalAmountFloatNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetTotalAmountFloatNil(b bool)`

 SetTotalAmountFloatNil sets the value for TotalAmountFloat to be an explicit nil

### UnsetTotalAmountFloat
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetTotalAmountFloat()`

UnsetTotalAmountFloat ensures that no value is present for TotalAmountFloat, not even an explicit nil
### GetFormattedTotalAmount

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetFormattedTotalAmount() interface{}`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetFormattedTotalAmountOk() (*interface{}, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetFormattedTotalAmount(v interface{})`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### SetFormattedTotalAmountNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetFormattedTotalAmountNil(b bool)`

 SetFormattedTotalAmountNil sets the value for FormattedTotalAmount to be an explicit nil

### UnsetFormattedTotalAmount
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetFormattedTotalAmount()`

UnsetFormattedTotalAmount ensures that no value is present for FormattedTotalAmount, not even an explicit nil
### GetCreatedAt

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETOrderSubscriptionItems200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


