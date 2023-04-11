# PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **interface{}** | The line item quantity. | [optional] 
**Restock** | Pointer to **interface{}** | Send this attribute if you want to restock the line item. | [optional] 
**ReturnReason** | Pointer to **interface{}** | Set of key-value pairs that you can use to add details about return reason. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHReturnLineItemsReturnLineItemIdRequestDataAttributes

`func NewPATCHReturnLineItemsReturnLineItemIdRequestDataAttributes() *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes`

NewPATCHReturnLineItemsReturnLineItemIdRequestDataAttributes instantiates a new PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHReturnLineItemsReturnLineItemIdRequestDataAttributesWithDefaults

`func NewPATCHReturnLineItemsReturnLineItemIdRequestDataAttributesWithDefaults() *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes`

NewPATCHReturnLineItemsReturnLineItemIdRequestDataAttributesWithDefaults instantiates a new PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetRestock

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetRestock() interface{}`

GetRestock returns the Restock field if non-nil, zero value otherwise.

### GetRestockOk

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetRestockOk() (*interface{}, bool)`

GetRestockOk returns a tuple with the Restock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestock

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetRestock(v interface{})`

SetRestock sets Restock field to given value.

### HasRestock

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) HasRestock() bool`

HasRestock returns a boolean if a field has been set.

### SetRestockNil

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetRestockNil(b bool)`

 SetRestockNil sets the value for Restock to be an explicit nil

### UnsetRestock
`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) UnsetRestock()`

UnsetRestock ensures that no value is present for Restock, not even an explicit nil
### GetReturnReason

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetReturnReason() interface{}`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetReturnReasonOk() (*interface{}, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetReturnReason(v interface{})`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### SetReturnReasonNil

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetReturnReasonNil(b bool)`

 SetReturnReasonNil sets the value for ReturnReason to be an explicit nil

### UnsetReturnReason
`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) UnsetReturnReason()`

UnsetReturnReason ensures that no value is present for ReturnReason, not even an explicit nil
### GetReference

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHReturnLineItemsReturnLineItemIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


