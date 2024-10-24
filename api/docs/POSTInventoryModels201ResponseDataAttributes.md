# POSTInventoryModels201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The inventory model&#39;s internal name. | 
**Strategy** | Pointer to **interface{}** | The inventory model&#39;s shipping strategy: one between &#39;no_split&#39; (default), &#39;split_shipments&#39;, &#39;ship_from_primary&#39; and &#39;ship_from_first_available_or_primary&#39;. | [optional] 
**StockLocationsCutoff** | Pointer to **interface{}** | The maximum number of stock locations used for inventory computation. | [optional] 
**StockReservationCutoff** | Pointer to **interface{}** | The duration in seconds of the generated stock reservations. | [optional] 
**PutStockTransfersOnHold** | Pointer to **interface{}** | Indicates if the the stock transfers must be put on hold automatically with the associated shipment. | [optional] 
**ManualStockDecrement** | Pointer to **interface{}** | Indicates if the the stock will be decremented manually after the order approval. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTInventoryModels201ResponseDataAttributes

`func NewPOSTInventoryModels201ResponseDataAttributes(name interface{}, ) *POSTInventoryModels201ResponseDataAttributes`

NewPOSTInventoryModels201ResponseDataAttributes instantiates a new POSTInventoryModels201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTInventoryModels201ResponseDataAttributesWithDefaults

`func NewPOSTInventoryModels201ResponseDataAttributesWithDefaults() *POSTInventoryModels201ResponseDataAttributes`

NewPOSTInventoryModels201ResponseDataAttributesWithDefaults instantiates a new POSTInventoryModels201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTInventoryModels201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTInventoryModels201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTInventoryModels201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTInventoryModels201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTInventoryModels201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStrategy

`func (o *POSTInventoryModels201ResponseDataAttributes) GetStrategy() interface{}`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *POSTInventoryModels201ResponseDataAttributes) GetStrategyOk() (*interface{}, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *POSTInventoryModels201ResponseDataAttributes) SetStrategy(v interface{})`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *POSTInventoryModels201ResponseDataAttributes) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### SetStrategyNil

`func (o *POSTInventoryModels201ResponseDataAttributes) SetStrategyNil(b bool)`

 SetStrategyNil sets the value for Strategy to be an explicit nil

### UnsetStrategy
`func (o *POSTInventoryModels201ResponseDataAttributes) UnsetStrategy()`

UnsetStrategy ensures that no value is present for Strategy, not even an explicit nil
### GetStockLocationsCutoff

`func (o *POSTInventoryModels201ResponseDataAttributes) GetStockLocationsCutoff() interface{}`

GetStockLocationsCutoff returns the StockLocationsCutoff field if non-nil, zero value otherwise.

### GetStockLocationsCutoffOk

`func (o *POSTInventoryModels201ResponseDataAttributes) GetStockLocationsCutoffOk() (*interface{}, bool)`

GetStockLocationsCutoffOk returns a tuple with the StockLocationsCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocationsCutoff

`func (o *POSTInventoryModels201ResponseDataAttributes) SetStockLocationsCutoff(v interface{})`

SetStockLocationsCutoff sets StockLocationsCutoff field to given value.

### HasStockLocationsCutoff

`func (o *POSTInventoryModels201ResponseDataAttributes) HasStockLocationsCutoff() bool`

HasStockLocationsCutoff returns a boolean if a field has been set.

### SetStockLocationsCutoffNil

`func (o *POSTInventoryModels201ResponseDataAttributes) SetStockLocationsCutoffNil(b bool)`

 SetStockLocationsCutoffNil sets the value for StockLocationsCutoff to be an explicit nil

### UnsetStockLocationsCutoff
`func (o *POSTInventoryModels201ResponseDataAttributes) UnsetStockLocationsCutoff()`

UnsetStockLocationsCutoff ensures that no value is present for StockLocationsCutoff, not even an explicit nil
### GetStockReservationCutoff

`func (o *POSTInventoryModels201ResponseDataAttributes) GetStockReservationCutoff() interface{}`

GetStockReservationCutoff returns the StockReservationCutoff field if non-nil, zero value otherwise.

### GetStockReservationCutoffOk

`func (o *POSTInventoryModels201ResponseDataAttributes) GetStockReservationCutoffOk() (*interface{}, bool)`

GetStockReservationCutoffOk returns a tuple with the StockReservationCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservationCutoff

`func (o *POSTInventoryModels201ResponseDataAttributes) SetStockReservationCutoff(v interface{})`

SetStockReservationCutoff sets StockReservationCutoff field to given value.

### HasStockReservationCutoff

`func (o *POSTInventoryModels201ResponseDataAttributes) HasStockReservationCutoff() bool`

HasStockReservationCutoff returns a boolean if a field has been set.

### SetStockReservationCutoffNil

`func (o *POSTInventoryModels201ResponseDataAttributes) SetStockReservationCutoffNil(b bool)`

 SetStockReservationCutoffNil sets the value for StockReservationCutoff to be an explicit nil

### UnsetStockReservationCutoff
`func (o *POSTInventoryModels201ResponseDataAttributes) UnsetStockReservationCutoff()`

UnsetStockReservationCutoff ensures that no value is present for StockReservationCutoff, not even an explicit nil
### GetPutStockTransfersOnHold

`func (o *POSTInventoryModels201ResponseDataAttributes) GetPutStockTransfersOnHold() interface{}`

GetPutStockTransfersOnHold returns the PutStockTransfersOnHold field if non-nil, zero value otherwise.

### GetPutStockTransfersOnHoldOk

`func (o *POSTInventoryModels201ResponseDataAttributes) GetPutStockTransfersOnHoldOk() (*interface{}, bool)`

GetPutStockTransfersOnHoldOk returns a tuple with the PutStockTransfersOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutStockTransfersOnHold

`func (o *POSTInventoryModels201ResponseDataAttributes) SetPutStockTransfersOnHold(v interface{})`

SetPutStockTransfersOnHold sets PutStockTransfersOnHold field to given value.

### HasPutStockTransfersOnHold

`func (o *POSTInventoryModels201ResponseDataAttributes) HasPutStockTransfersOnHold() bool`

HasPutStockTransfersOnHold returns a boolean if a field has been set.

### SetPutStockTransfersOnHoldNil

`func (o *POSTInventoryModels201ResponseDataAttributes) SetPutStockTransfersOnHoldNil(b bool)`

 SetPutStockTransfersOnHoldNil sets the value for PutStockTransfersOnHold to be an explicit nil

### UnsetPutStockTransfersOnHold
`func (o *POSTInventoryModels201ResponseDataAttributes) UnsetPutStockTransfersOnHold()`

UnsetPutStockTransfersOnHold ensures that no value is present for PutStockTransfersOnHold, not even an explicit nil
### GetManualStockDecrement

`func (o *POSTInventoryModels201ResponseDataAttributes) GetManualStockDecrement() interface{}`

GetManualStockDecrement returns the ManualStockDecrement field if non-nil, zero value otherwise.

### GetManualStockDecrementOk

`func (o *POSTInventoryModels201ResponseDataAttributes) GetManualStockDecrementOk() (*interface{}, bool)`

GetManualStockDecrementOk returns a tuple with the ManualStockDecrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualStockDecrement

`func (o *POSTInventoryModels201ResponseDataAttributes) SetManualStockDecrement(v interface{})`

SetManualStockDecrement sets ManualStockDecrement field to given value.

### HasManualStockDecrement

`func (o *POSTInventoryModels201ResponseDataAttributes) HasManualStockDecrement() bool`

HasManualStockDecrement returns a boolean if a field has been set.

### SetManualStockDecrementNil

`func (o *POSTInventoryModels201ResponseDataAttributes) SetManualStockDecrementNil(b bool)`

 SetManualStockDecrementNil sets the value for ManualStockDecrement to be an explicit nil

### UnsetManualStockDecrement
`func (o *POSTInventoryModels201ResponseDataAttributes) UnsetManualStockDecrement()`

UnsetManualStockDecrement ensures that no value is present for ManualStockDecrement, not even an explicit nil
### GetReference

`func (o *POSTInventoryModels201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTInventoryModels201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTInventoryModels201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTInventoryModels201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTInventoryModels201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTInventoryModels201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTInventoryModels201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTInventoryModels201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTInventoryModels201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTInventoryModels201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTInventoryModels201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTInventoryModels201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTInventoryModels201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTInventoryModels201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTInventoryModels201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTInventoryModels201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTInventoryModels201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTInventoryModels201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


