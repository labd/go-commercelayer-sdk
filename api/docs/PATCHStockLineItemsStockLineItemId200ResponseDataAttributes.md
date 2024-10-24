# PATCHStockLineItemsStockLineItemId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**Quantity** | Pointer to **interface{}** | The line item quantity. | [optional] 
**ReserveStock** | Pointer to **interface{}** | Send this attribute if you want to automatically reserve the stock for this stock line item. Can be done only when fulfillment is in progress. Cannot be passed by sales channels. | [optional] 
**ReleaseStock** | Pointer to **interface{}** | Send this attribute if you want to automatically destroy the stock reservation for this stock line item. Can be done only when fulfillment is in progress. Cannot be passed by sales channels. | [optional] 
**DecrementStock** | Pointer to **interface{}** | Send this attribute if you want to automatically decrement and release the stock this stock line item. Can be done only when fulfillment is in progress. Cannot be passed by sales channels. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHStockLineItemsStockLineItemId200ResponseDataAttributes

`func NewPATCHStockLineItemsStockLineItemId200ResponseDataAttributes() *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes`

NewPATCHStockLineItemsStockLineItemId200ResponseDataAttributes instantiates a new PATCHStockLineItemsStockLineItemId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHStockLineItemsStockLineItemId200ResponseDataAttributesWithDefaults

`func NewPATCHStockLineItemsStockLineItemId200ResponseDataAttributesWithDefaults() *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes`

NewPATCHStockLineItemsStockLineItemId200ResponseDataAttributesWithDefaults instantiates a new PATCHStockLineItemsStockLineItemId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetQuantity

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetReserveStock

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetReserveStock() interface{}`

GetReserveStock returns the ReserveStock field if non-nil, zero value otherwise.

### GetReserveStockOk

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetReserveStockOk() (*interface{}, bool)`

GetReserveStockOk returns a tuple with the ReserveStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveStock

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetReserveStock(v interface{})`

SetReserveStock sets ReserveStock field to given value.

### HasReserveStock

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) HasReserveStock() bool`

HasReserveStock returns a boolean if a field has been set.

### SetReserveStockNil

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetReserveStockNil(b bool)`

 SetReserveStockNil sets the value for ReserveStock to be an explicit nil

### UnsetReserveStock
`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) UnsetReserveStock()`

UnsetReserveStock ensures that no value is present for ReserveStock, not even an explicit nil
### GetReleaseStock

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetReleaseStock() interface{}`

GetReleaseStock returns the ReleaseStock field if non-nil, zero value otherwise.

### GetReleaseStockOk

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetReleaseStockOk() (*interface{}, bool)`

GetReleaseStockOk returns a tuple with the ReleaseStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStock

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetReleaseStock(v interface{})`

SetReleaseStock sets ReleaseStock field to given value.

### HasReleaseStock

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) HasReleaseStock() bool`

HasReleaseStock returns a boolean if a field has been set.

### SetReleaseStockNil

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetReleaseStockNil(b bool)`

 SetReleaseStockNil sets the value for ReleaseStock to be an explicit nil

### UnsetReleaseStock
`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) UnsetReleaseStock()`

UnsetReleaseStock ensures that no value is present for ReleaseStock, not even an explicit nil
### GetDecrementStock

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetDecrementStock() interface{}`

GetDecrementStock returns the DecrementStock field if non-nil, zero value otherwise.

### GetDecrementStockOk

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetDecrementStockOk() (*interface{}, bool)`

GetDecrementStockOk returns a tuple with the DecrementStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecrementStock

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetDecrementStock(v interface{})`

SetDecrementStock sets DecrementStock field to given value.

### HasDecrementStock

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) HasDecrementStock() bool`

HasDecrementStock returns a boolean if a field has been set.

### SetDecrementStockNil

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetDecrementStockNil(b bool)`

 SetDecrementStockNil sets the value for DecrementStock to be an explicit nil

### UnsetDecrementStock
`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) UnsetDecrementStock()`

UnsetDecrementStock ensures that no value is present for DecrementStock, not even an explicit nil
### GetReference

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHStockLineItemsStockLineItemId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


