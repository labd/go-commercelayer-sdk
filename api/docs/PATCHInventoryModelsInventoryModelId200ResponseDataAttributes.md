# PATCHInventoryModelsInventoryModelId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The inventory model&#39;s internal name. | [optional] 
**Strategy** | Pointer to **interface{}** | The inventory model&#39;s shipping strategy: one between &#39;no_split&#39; (default), &#39;split_shipments&#39;, &#39;ship_from_primary&#39; and &#39;ship_from_first_available_or_primary&#39;. | [optional] 
**StockLocationsCutoff** | Pointer to **interface{}** | The maximum number of stock locations used for inventory computation. | [optional] 
**StockReservationCutoff** | Pointer to **interface{}** | The duration in seconds of the generated stock reservations. | [optional] 
**PutStockTransfersOnHold** | Pointer to **interface{}** | Indicates if the the stock transfers must be put on hold automatically with the associated shipment. | [optional] 
**ManualStockDecrement** | Pointer to **interface{}** | Indicates if the the stock will be decremented manually after the order approval. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStrategy

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetStrategy() interface{}`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetStrategyOk() (*interface{}, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetStrategy(v interface{})`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### SetStrategyNil

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetStrategyNil(b bool)`

 SetStrategyNil sets the value for Strategy to be an explicit nil

### UnsetStrategy
`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetStrategy()`

UnsetStrategy ensures that no value is present for Strategy, not even an explicit nil
### GetStockLocationsCutoff

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockLocationsCutoff() interface{}`

GetStockLocationsCutoff returns the StockLocationsCutoff field if non-nil, zero value otherwise.

### GetStockLocationsCutoffOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockLocationsCutoffOk() (*interface{}, bool)`

GetStockLocationsCutoffOk returns a tuple with the StockLocationsCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocationsCutoff

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetStockLocationsCutoff(v interface{})`

SetStockLocationsCutoff sets StockLocationsCutoff field to given value.

### HasStockLocationsCutoff

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasStockLocationsCutoff() bool`

HasStockLocationsCutoff returns a boolean if a field has been set.

### SetStockLocationsCutoffNil

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetStockLocationsCutoffNil(b bool)`

 SetStockLocationsCutoffNil sets the value for StockLocationsCutoff to be an explicit nil

### UnsetStockLocationsCutoff
`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetStockLocationsCutoff()`

UnsetStockLocationsCutoff ensures that no value is present for StockLocationsCutoff, not even an explicit nil
### GetStockReservationCutoff

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockReservationCutoff() interface{}`

GetStockReservationCutoff returns the StockReservationCutoff field if non-nil, zero value otherwise.

### GetStockReservationCutoffOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockReservationCutoffOk() (*interface{}, bool)`

GetStockReservationCutoffOk returns a tuple with the StockReservationCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservationCutoff

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetStockReservationCutoff(v interface{})`

SetStockReservationCutoff sets StockReservationCutoff field to given value.

### HasStockReservationCutoff

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasStockReservationCutoff() bool`

HasStockReservationCutoff returns a boolean if a field has been set.

### SetStockReservationCutoffNil

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetStockReservationCutoffNil(b bool)`

 SetStockReservationCutoffNil sets the value for StockReservationCutoff to be an explicit nil

### UnsetStockReservationCutoff
`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetStockReservationCutoff()`

UnsetStockReservationCutoff ensures that no value is present for StockReservationCutoff, not even an explicit nil
### GetPutStockTransfersOnHold

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetPutStockTransfersOnHold() interface{}`

GetPutStockTransfersOnHold returns the PutStockTransfersOnHold field if non-nil, zero value otherwise.

### GetPutStockTransfersOnHoldOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetPutStockTransfersOnHoldOk() (*interface{}, bool)`

GetPutStockTransfersOnHoldOk returns a tuple with the PutStockTransfersOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutStockTransfersOnHold

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetPutStockTransfersOnHold(v interface{})`

SetPutStockTransfersOnHold sets PutStockTransfersOnHold field to given value.

### HasPutStockTransfersOnHold

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasPutStockTransfersOnHold() bool`

HasPutStockTransfersOnHold returns a boolean if a field has been set.

### SetPutStockTransfersOnHoldNil

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetPutStockTransfersOnHoldNil(b bool)`

 SetPutStockTransfersOnHoldNil sets the value for PutStockTransfersOnHold to be an explicit nil

### UnsetPutStockTransfersOnHold
`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetPutStockTransfersOnHold()`

UnsetPutStockTransfersOnHold ensures that no value is present for PutStockTransfersOnHold, not even an explicit nil
### GetManualStockDecrement

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetManualStockDecrement() interface{}`

GetManualStockDecrement returns the ManualStockDecrement field if non-nil, zero value otherwise.

### GetManualStockDecrementOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetManualStockDecrementOk() (*interface{}, bool)`

GetManualStockDecrementOk returns a tuple with the ManualStockDecrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualStockDecrement

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetManualStockDecrement(v interface{})`

SetManualStockDecrement sets ManualStockDecrement field to given value.

### HasManualStockDecrement

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasManualStockDecrement() bool`

HasManualStockDecrement returns a boolean if a field has been set.

### SetManualStockDecrementNil

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetManualStockDecrementNil(b bool)`

 SetManualStockDecrementNil sets the value for ManualStockDecrement to be an explicit nil

### UnsetManualStockDecrement
`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetManualStockDecrement()`

UnsetManualStockDecrement ensures that no value is present for ManualStockDecrement, not even an explicit nil
### GetReference

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


