# PATCHInventoryModelsInventoryModelId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The inventory model&#39;s internal name. | [optional] 
**Strategy** | Pointer to **string** | The inventory model&#39;s shipping strategy: one between &#39;no_split&#39; (default), &#39;split_shipments&#39;, &#39;ship_from_primary&#39; and &#39;ship_from_first_available_or_primary&#39;. | [optional] 
**StockLocationsCutoff** | Pointer to **int32** | The maximum number of stock locations used for inventory computation | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHInventoryModelsInventoryModelId200ResponseDataAttributes

`func NewPATCHInventoryModelsInventoryModelId200ResponseDataAttributes() *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes`

NewPATCHInventoryModelsInventoryModelId200ResponseDataAttributes instantiates a new PATCHInventoryModelsInventoryModelId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHInventoryModelsInventoryModelId200ResponseDataAttributesWithDefaults

`func NewPATCHInventoryModelsInventoryModelId200ResponseDataAttributesWithDefaults() *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes`

NewPATCHInventoryModelsInventoryModelId200ResponseDataAttributesWithDefaults instantiates a new PATCHInventoryModelsInventoryModelId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStrategy

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetStockLocationsCutoff

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockLocationsCutoff() int32`

GetStockLocationsCutoff returns the StockLocationsCutoff field if non-nil, zero value otherwise.

### GetStockLocationsCutoffOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockLocationsCutoffOk() (*int32, bool)`

GetStockLocationsCutoffOk returns a tuple with the StockLocationsCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocationsCutoff

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetStockLocationsCutoff(v int32)`

SetStockLocationsCutoff sets StockLocationsCutoff field to given value.

### HasStockLocationsCutoff

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasStockLocationsCutoff() bool`

HasStockLocationsCutoff returns a boolean if a field has been set.

### GetReference

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


