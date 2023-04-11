# PATCHInventoryModelsInventoryModelIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The inventory model&#39;s internal name. | [optional] 
**Strategy** | Pointer to **interface{}** | The inventory model&#39;s shipping strategy: one between &#39;no_split&#39; (default), &#39;split_shipments&#39;, &#39;ship_from_primary&#39; and &#39;ship_from_first_available_or_primary&#39;. | [optional] 
**StockLocationsCutoff** | Pointer to **interface{}** | The maximum number of stock locations used for inventory computation | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHInventoryModelsInventoryModelIdRequestDataAttributes

`func NewPATCHInventoryModelsInventoryModelIdRequestDataAttributes() *PATCHInventoryModelsInventoryModelIdRequestDataAttributes`

NewPATCHInventoryModelsInventoryModelIdRequestDataAttributes instantiates a new PATCHInventoryModelsInventoryModelIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHInventoryModelsInventoryModelIdRequestDataAttributesWithDefaults

`func NewPATCHInventoryModelsInventoryModelIdRequestDataAttributesWithDefaults() *PATCHInventoryModelsInventoryModelIdRequestDataAttributes`

NewPATCHInventoryModelsInventoryModelIdRequestDataAttributesWithDefaults instantiates a new PATCHInventoryModelsInventoryModelIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStrategy

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetStrategy() interface{}`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetStrategyOk() (*interface{}, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetStrategy(v interface{})`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### SetStrategyNil

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetStrategyNil(b bool)`

 SetStrategyNil sets the value for Strategy to be an explicit nil

### UnsetStrategy
`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) UnsetStrategy()`

UnsetStrategy ensures that no value is present for Strategy, not even an explicit nil
### GetStockLocationsCutoff

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetStockLocationsCutoff() interface{}`

GetStockLocationsCutoff returns the StockLocationsCutoff field if non-nil, zero value otherwise.

### GetStockLocationsCutoffOk

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetStockLocationsCutoffOk() (*interface{}, bool)`

GetStockLocationsCutoffOk returns a tuple with the StockLocationsCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocationsCutoff

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetStockLocationsCutoff(v interface{})`

SetStockLocationsCutoff sets StockLocationsCutoff field to given value.

### HasStockLocationsCutoff

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) HasStockLocationsCutoff() bool`

HasStockLocationsCutoff returns a boolean if a field has been set.

### SetStockLocationsCutoffNil

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetStockLocationsCutoffNil(b bool)`

 SetStockLocationsCutoffNil sets the value for StockLocationsCutoff to be an explicit nil

### UnsetStockLocationsCutoff
`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) UnsetStockLocationsCutoff()`

UnsetStockLocationsCutoff ensures that no value is present for StockLocationsCutoff, not even an explicit nil
### GetReference

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHInventoryModelsInventoryModelIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


