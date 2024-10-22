# POSTStockLineItems201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**Quantity** | **interface{}** | The line item quantity. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTStockLineItems201ResponseDataAttributes

`func NewPOSTStockLineItems201ResponseDataAttributes(quantity interface{}, ) *POSTStockLineItems201ResponseDataAttributes`

NewPOSTStockLineItems201ResponseDataAttributes instantiates a new POSTStockLineItems201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockLineItems201ResponseDataAttributesWithDefaults

`func NewPOSTStockLineItems201ResponseDataAttributesWithDefaults() *POSTStockLineItems201ResponseDataAttributes`

NewPOSTStockLineItems201ResponseDataAttributesWithDefaults instantiates a new POSTStockLineItems201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *POSTStockLineItems201ResponseDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *POSTStockLineItems201ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *POSTStockLineItems201ResponseDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *POSTStockLineItems201ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *POSTStockLineItems201ResponseDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *POSTStockLineItems201ResponseDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetQuantity

`func (o *POSTStockLineItems201ResponseDataAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *POSTStockLineItems201ResponseDataAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *POSTStockLineItems201ResponseDataAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.


### SetQuantityNil

`func (o *POSTStockLineItems201ResponseDataAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *POSTStockLineItems201ResponseDataAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetReference

`func (o *POSTStockLineItems201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTStockLineItems201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTStockLineItems201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTStockLineItems201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTStockLineItems201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTStockLineItems201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTStockLineItems201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTStockLineItems201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTStockLineItems201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTStockLineItems201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTStockLineItems201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTStockLineItems201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTStockLineItems201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTStockLineItems201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTStockLineItems201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTStockLineItems201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTStockLineItems201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTStockLineItems201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


