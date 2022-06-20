# ReturnLineItemUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **int32** | The line item quantity. | [optional] 
**Restock** | Pointer to **bool** | Send this attribute if you want to restock the line item. | [optional] 
**ReturnReason** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can use to add details about return reason. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewReturnLineItemUpdateDataAttributes

`func NewReturnLineItemUpdateDataAttributes() *ReturnLineItemUpdateDataAttributes`

NewReturnLineItemUpdateDataAttributes instantiates a new ReturnLineItemUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnLineItemUpdateDataAttributesWithDefaults

`func NewReturnLineItemUpdateDataAttributesWithDefaults() *ReturnLineItemUpdateDataAttributes`

NewReturnLineItemUpdateDataAttributesWithDefaults instantiates a new ReturnLineItemUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *ReturnLineItemUpdateDataAttributes) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReturnLineItemUpdateDataAttributes) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReturnLineItemUpdateDataAttributes) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReturnLineItemUpdateDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRestock

`func (o *ReturnLineItemUpdateDataAttributes) GetRestock() bool`

GetRestock returns the Restock field if non-nil, zero value otherwise.

### GetRestockOk

`func (o *ReturnLineItemUpdateDataAttributes) GetRestockOk() (*bool, bool)`

GetRestockOk returns a tuple with the Restock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestock

`func (o *ReturnLineItemUpdateDataAttributes) SetRestock(v bool)`

SetRestock sets Restock field to given value.

### HasRestock

`func (o *ReturnLineItemUpdateDataAttributes) HasRestock() bool`

HasRestock returns a boolean if a field has been set.

### GetReturnReason

`func (o *ReturnLineItemUpdateDataAttributes) GetReturnReason() map[string]interface{}`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *ReturnLineItemUpdateDataAttributes) GetReturnReasonOk() (*map[string]interface{}, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *ReturnLineItemUpdateDataAttributes) SetReturnReason(v map[string]interface{})`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *ReturnLineItemUpdateDataAttributes) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetReference

`func (o *ReturnLineItemUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ReturnLineItemUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ReturnLineItemUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ReturnLineItemUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ReturnLineItemUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ReturnLineItemUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ReturnLineItemUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ReturnLineItemUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ReturnLineItemUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReturnLineItemUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReturnLineItemUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ReturnLineItemUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


