# GETInventoryModelsInventoryModelId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The inventory model&#39;s internal name. | [optional] 
**Strategy** | Pointer to **interface{}** | The inventory model&#39;s shipping strategy: one between &#39;no_split&#39; (default), &#39;split_shipments&#39;, &#39;ship_from_primary&#39; and &#39;ship_from_first_available_or_primary&#39;. | [optional] 
**StockLocationsCutoff** | Pointer to **interface{}** | The maximum number of stock locations used for inventory computation | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETInventoryModelsInventoryModelId200ResponseDataAttributes

`func NewGETInventoryModelsInventoryModelId200ResponseDataAttributes() *GETInventoryModelsInventoryModelId200ResponseDataAttributes`

NewGETInventoryModelsInventoryModelId200ResponseDataAttributes instantiates a new GETInventoryModelsInventoryModelId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETInventoryModelsInventoryModelId200ResponseDataAttributesWithDefaults

`func NewGETInventoryModelsInventoryModelId200ResponseDataAttributesWithDefaults() *GETInventoryModelsInventoryModelId200ResponseDataAttributes`

NewGETInventoryModelsInventoryModelId200ResponseDataAttributesWithDefaults instantiates a new GETInventoryModelsInventoryModelId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStrategy

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetStrategy() interface{}`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetStrategyOk() (*interface{}, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetStrategy(v interface{})`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### SetStrategyNil

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetStrategyNil(b bool)`

 SetStrategyNil sets the value for Strategy to be an explicit nil

### UnsetStrategy
`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetStrategy()`

UnsetStrategy ensures that no value is present for Strategy, not even an explicit nil
### GetStockLocationsCutoff

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockLocationsCutoff() interface{}`

GetStockLocationsCutoff returns the StockLocationsCutoff field if non-nil, zero value otherwise.

### GetStockLocationsCutoffOk

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockLocationsCutoffOk() (*interface{}, bool)`

GetStockLocationsCutoffOk returns a tuple with the StockLocationsCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocationsCutoff

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetStockLocationsCutoff(v interface{})`

SetStockLocationsCutoff sets StockLocationsCutoff field to given value.

### HasStockLocationsCutoff

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasStockLocationsCutoff() bool`

HasStockLocationsCutoff returns a boolean if a field has been set.

### SetStockLocationsCutoffNil

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetStockLocationsCutoffNil(b bool)`

 SetStockLocationsCutoffNil sets the value for StockLocationsCutoff to be an explicit nil

### UnsetStockLocationsCutoff
`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetStockLocationsCutoff()`

UnsetStockLocationsCutoff ensures that no value is present for StockLocationsCutoff, not even an explicit nil
### GetCreatedAt

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


