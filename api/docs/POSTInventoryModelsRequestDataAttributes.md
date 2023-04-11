# POSTInventoryModelsRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The inventory model&#39;s internal name. | 
**Strategy** | Pointer to **interface{}** | The inventory model&#39;s shipping strategy: one between &#39;no_split&#39; (default), &#39;split_shipments&#39;, &#39;ship_from_primary&#39; and &#39;ship_from_first_available_or_primary&#39;. | [optional] 
**StockLocationsCutoff** | Pointer to **interface{}** | The maximum number of stock locations used for inventory computation | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTInventoryModelsRequestDataAttributes

`func NewPOSTInventoryModelsRequestDataAttributes(name interface{}, ) *POSTInventoryModelsRequestDataAttributes`

NewPOSTInventoryModelsRequestDataAttributes instantiates a new POSTInventoryModelsRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTInventoryModelsRequestDataAttributesWithDefaults

`func NewPOSTInventoryModelsRequestDataAttributesWithDefaults() *POSTInventoryModelsRequestDataAttributes`

NewPOSTInventoryModelsRequestDataAttributesWithDefaults instantiates a new POSTInventoryModelsRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTInventoryModelsRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTInventoryModelsRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTInventoryModelsRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTInventoryModelsRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTInventoryModelsRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStrategy

`func (o *POSTInventoryModelsRequestDataAttributes) GetStrategy() interface{}`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *POSTInventoryModelsRequestDataAttributes) GetStrategyOk() (*interface{}, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *POSTInventoryModelsRequestDataAttributes) SetStrategy(v interface{})`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *POSTInventoryModelsRequestDataAttributes) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### SetStrategyNil

`func (o *POSTInventoryModelsRequestDataAttributes) SetStrategyNil(b bool)`

 SetStrategyNil sets the value for Strategy to be an explicit nil

### UnsetStrategy
`func (o *POSTInventoryModelsRequestDataAttributes) UnsetStrategy()`

UnsetStrategy ensures that no value is present for Strategy, not even an explicit nil
### GetStockLocationsCutoff

`func (o *POSTInventoryModelsRequestDataAttributes) GetStockLocationsCutoff() interface{}`

GetStockLocationsCutoff returns the StockLocationsCutoff field if non-nil, zero value otherwise.

### GetStockLocationsCutoffOk

`func (o *POSTInventoryModelsRequestDataAttributes) GetStockLocationsCutoffOk() (*interface{}, bool)`

GetStockLocationsCutoffOk returns a tuple with the StockLocationsCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocationsCutoff

`func (o *POSTInventoryModelsRequestDataAttributes) SetStockLocationsCutoff(v interface{})`

SetStockLocationsCutoff sets StockLocationsCutoff field to given value.

### HasStockLocationsCutoff

`func (o *POSTInventoryModelsRequestDataAttributes) HasStockLocationsCutoff() bool`

HasStockLocationsCutoff returns a boolean if a field has been set.

### SetStockLocationsCutoffNil

`func (o *POSTInventoryModelsRequestDataAttributes) SetStockLocationsCutoffNil(b bool)`

 SetStockLocationsCutoffNil sets the value for StockLocationsCutoff to be an explicit nil

### UnsetStockLocationsCutoff
`func (o *POSTInventoryModelsRequestDataAttributes) UnsetStockLocationsCutoff()`

UnsetStockLocationsCutoff ensures that no value is present for StockLocationsCutoff, not even an explicit nil
### GetReference

`func (o *POSTInventoryModelsRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTInventoryModelsRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTInventoryModelsRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTInventoryModelsRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTInventoryModelsRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTInventoryModelsRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTInventoryModelsRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTInventoryModelsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTInventoryModelsRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTInventoryModelsRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTInventoryModelsRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTInventoryModelsRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTInventoryModelsRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTInventoryModelsRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTInventoryModelsRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTInventoryModelsRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTInventoryModelsRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTInventoryModelsRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


