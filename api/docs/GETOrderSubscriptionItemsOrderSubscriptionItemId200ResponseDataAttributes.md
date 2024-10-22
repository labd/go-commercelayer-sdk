# GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**BundleCode** | Pointer to **interface{}** | The code of the associated bundle. | [optional] 
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
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes

`func NewGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes() *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes`

NewGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes instantiates a new GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributesWithDefaults

`func NewGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributesWithDefaults() *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes`

NewGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributesWithDefaults instantiates a new GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetBundleCode

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetBundleCode() interface{}`

GetBundleCode returns the BundleCode field if non-nil, zero value otherwise.

### GetBundleCodeOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetBundleCodeOk() (*interface{}, bool)`

GetBundleCodeOk returns a tuple with the BundleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleCode

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetBundleCode(v interface{})`

SetBundleCode sets BundleCode field to given value.

### HasBundleCode

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasBundleCode() bool`

HasBundleCode returns a boolean if a field has been set.

### SetBundleCodeNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetBundleCodeNil(b bool)`

 SetBundleCodeNil sets the value for BundleCode to be an explicit nil

### UnsetBundleCode
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetBundleCode()`

UnsetBundleCode ensures that no value is present for BundleCode, not even an explicit nil
### GetQuantity

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetUnitAmountCents

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUnitAmountCents() interface{}`

GetUnitAmountCents returns the UnitAmountCents field if non-nil, zero value otherwise.

### GetUnitAmountCentsOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUnitAmountCentsOk() (*interface{}, bool)`

GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountCents

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetUnitAmountCents(v interface{})`

SetUnitAmountCents sets UnitAmountCents field to given value.

### HasUnitAmountCents

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasUnitAmountCents() bool`

HasUnitAmountCents returns a boolean if a field has been set.

### SetUnitAmountCentsNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetUnitAmountCentsNil(b bool)`

 SetUnitAmountCentsNil sets the value for UnitAmountCents to be an explicit nil

### UnsetUnitAmountCents
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetUnitAmountCents()`

UnsetUnitAmountCents ensures that no value is present for UnitAmountCents, not even an explicit nil
### GetUnitAmountFloat

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUnitAmountFloat() interface{}`

GetUnitAmountFloat returns the UnitAmountFloat field if non-nil, zero value otherwise.

### GetUnitAmountFloatOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUnitAmountFloatOk() (*interface{}, bool)`

GetUnitAmountFloatOk returns a tuple with the UnitAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountFloat

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetUnitAmountFloat(v interface{})`

SetUnitAmountFloat sets UnitAmountFloat field to given value.

### HasUnitAmountFloat

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasUnitAmountFloat() bool`

HasUnitAmountFloat returns a boolean if a field has been set.

### SetUnitAmountFloatNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetUnitAmountFloatNil(b bool)`

 SetUnitAmountFloatNil sets the value for UnitAmountFloat to be an explicit nil

### UnsetUnitAmountFloat
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetUnitAmountFloat()`

UnsetUnitAmountFloat ensures that no value is present for UnitAmountFloat, not even an explicit nil
### GetFormattedUnitAmount

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetFormattedUnitAmount() interface{}`

GetFormattedUnitAmount returns the FormattedUnitAmount field if non-nil, zero value otherwise.

### GetFormattedUnitAmountOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetFormattedUnitAmountOk() (*interface{}, bool)`

GetFormattedUnitAmountOk returns a tuple with the FormattedUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedUnitAmount

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetFormattedUnitAmount(v interface{})`

SetFormattedUnitAmount sets FormattedUnitAmount field to given value.

### HasFormattedUnitAmount

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasFormattedUnitAmount() bool`

HasFormattedUnitAmount returns a boolean if a field has been set.

### SetFormattedUnitAmountNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetFormattedUnitAmountNil(b bool)`

 SetFormattedUnitAmountNil sets the value for FormattedUnitAmount to be an explicit nil

### UnsetFormattedUnitAmount
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetFormattedUnitAmount()`

UnsetFormattedUnitAmount ensures that no value is present for FormattedUnitAmount, not even an explicit nil
### GetTotalAmountCents

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetTotalAmountCents() interface{}`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetTotalAmountCentsOk() (*interface{}, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetTotalAmountCents(v interface{})`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### SetTotalAmountCentsNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetTotalAmountCentsNil(b bool)`

 SetTotalAmountCentsNil sets the value for TotalAmountCents to be an explicit nil

### UnsetTotalAmountCents
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetTotalAmountCents()`

UnsetTotalAmountCents ensures that no value is present for TotalAmountCents, not even an explicit nil
### GetTotalAmountFloat

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetTotalAmountFloat() interface{}`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetTotalAmountFloatOk() (*interface{}, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetTotalAmountFloat(v interface{})`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### SetTotalAmountFloatNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetTotalAmountFloatNil(b bool)`

 SetTotalAmountFloatNil sets the value for TotalAmountFloat to be an explicit nil

### UnsetTotalAmountFloat
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetTotalAmountFloat()`

UnsetTotalAmountFloat ensures that no value is present for TotalAmountFloat, not even an explicit nil
### GetFormattedTotalAmount

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetFormattedTotalAmount() interface{}`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetFormattedTotalAmountOk() (*interface{}, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetFormattedTotalAmount(v interface{})`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### SetFormattedTotalAmountNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetFormattedTotalAmountNil(b bool)`

 SetFormattedTotalAmountNil sets the value for FormattedTotalAmount to be an explicit nil

### UnsetFormattedTotalAmount
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetFormattedTotalAmount()`

UnsetFormattedTotalAmount ensures that no value is present for FormattedTotalAmount, not even an explicit nil
### GetCreatedAt

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


