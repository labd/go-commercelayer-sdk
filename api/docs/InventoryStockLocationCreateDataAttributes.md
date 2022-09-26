# InventoryStockLocationCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | **int32** | The stock location priority within the associated invetory model. | 
**OnHold** | Pointer to **bool** | Indicates if the shipment should be put on hold if fulfilled from the associated stock location. This is useful to manage use cases like back-orders, pre-orders or personalized orders that need to be customized before being fulfilled. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewInventoryStockLocationCreateDataAttributes

`func NewInventoryStockLocationCreateDataAttributes(priority int32, ) *InventoryStockLocationCreateDataAttributes`

NewInventoryStockLocationCreateDataAttributes instantiates a new InventoryStockLocationCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryStockLocationCreateDataAttributesWithDefaults

`func NewInventoryStockLocationCreateDataAttributesWithDefaults() *InventoryStockLocationCreateDataAttributes`

NewInventoryStockLocationCreateDataAttributesWithDefaults instantiates a new InventoryStockLocationCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *InventoryStockLocationCreateDataAttributes) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InventoryStockLocationCreateDataAttributes) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InventoryStockLocationCreateDataAttributes) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetOnHold

`func (o *InventoryStockLocationCreateDataAttributes) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *InventoryStockLocationCreateDataAttributes) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *InventoryStockLocationCreateDataAttributes) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *InventoryStockLocationCreateDataAttributes) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetReference

`func (o *InventoryStockLocationCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *InventoryStockLocationCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *InventoryStockLocationCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *InventoryStockLocationCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *InventoryStockLocationCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *InventoryStockLocationCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *InventoryStockLocationCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *InventoryStockLocationCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *InventoryStockLocationCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InventoryStockLocationCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InventoryStockLocationCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InventoryStockLocationCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


