# StockTransferResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to [**InStockSubscriptionResponseDataRelationshipsSku**](InStockSubscriptionResponseDataRelationshipsSku.md) |  | [optional] 
**OriginStockLocation** | Pointer to [**StockTransferResponseDataRelationshipsOriginStockLocation**](StockTransferResponseDataRelationshipsOriginStockLocation.md) |  | [optional] 
**DestinationStockLocation** | Pointer to [**StockTransferResponseDataRelationshipsDestinationStockLocation**](StockTransferResponseDataRelationshipsDestinationStockLocation.md) |  | [optional] 
**Shipment** | Pointer to [**ParcelResponseDataRelationshipsShipment**](ParcelResponseDataRelationshipsShipment.md) |  | [optional] 
**LineItem** | Pointer to [**LineItemOptionResponseDataRelationshipsLineItem**](LineItemOptionResponseDataRelationshipsLineItem.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressResponseDataRelationshipsEvents**](CustomerAddressResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewStockTransferResponseDataRelationships

`func NewStockTransferResponseDataRelationships() *StockTransferResponseDataRelationships`

NewStockTransferResponseDataRelationships instantiates a new StockTransferResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockTransferResponseDataRelationshipsWithDefaults

`func NewStockTransferResponseDataRelationshipsWithDefaults() *StockTransferResponseDataRelationships`

NewStockTransferResponseDataRelationshipsWithDefaults instantiates a new StockTransferResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *StockTransferResponseDataRelationships) GetSku() InStockSubscriptionResponseDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StockTransferResponseDataRelationships) GetSkuOk() (*InStockSubscriptionResponseDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StockTransferResponseDataRelationships) SetSku(v InStockSubscriptionResponseDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *StockTransferResponseDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetOriginStockLocation

`func (o *StockTransferResponseDataRelationships) GetOriginStockLocation() StockTransferResponseDataRelationshipsOriginStockLocation`

GetOriginStockLocation returns the OriginStockLocation field if non-nil, zero value otherwise.

### GetOriginStockLocationOk

`func (o *StockTransferResponseDataRelationships) GetOriginStockLocationOk() (*StockTransferResponseDataRelationshipsOriginStockLocation, bool)`

GetOriginStockLocationOk returns a tuple with the OriginStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStockLocation

`func (o *StockTransferResponseDataRelationships) SetOriginStockLocation(v StockTransferResponseDataRelationshipsOriginStockLocation)`

SetOriginStockLocation sets OriginStockLocation field to given value.

### HasOriginStockLocation

`func (o *StockTransferResponseDataRelationships) HasOriginStockLocation() bool`

HasOriginStockLocation returns a boolean if a field has been set.

### GetDestinationStockLocation

`func (o *StockTransferResponseDataRelationships) GetDestinationStockLocation() StockTransferResponseDataRelationshipsDestinationStockLocation`

GetDestinationStockLocation returns the DestinationStockLocation field if non-nil, zero value otherwise.

### GetDestinationStockLocationOk

`func (o *StockTransferResponseDataRelationships) GetDestinationStockLocationOk() (*StockTransferResponseDataRelationshipsDestinationStockLocation, bool)`

GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationStockLocation

`func (o *StockTransferResponseDataRelationships) SetDestinationStockLocation(v StockTransferResponseDataRelationshipsDestinationStockLocation)`

SetDestinationStockLocation sets DestinationStockLocation field to given value.

### HasDestinationStockLocation

`func (o *StockTransferResponseDataRelationships) HasDestinationStockLocation() bool`

HasDestinationStockLocation returns a boolean if a field has been set.

### GetShipment

`func (o *StockTransferResponseDataRelationships) GetShipment() ParcelResponseDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *StockTransferResponseDataRelationships) GetShipmentOk() (*ParcelResponseDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *StockTransferResponseDataRelationships) SetShipment(v ParcelResponseDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *StockTransferResponseDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetLineItem

`func (o *StockTransferResponseDataRelationships) GetLineItem() LineItemOptionResponseDataRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *StockTransferResponseDataRelationships) GetLineItemOk() (*LineItemOptionResponseDataRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *StockTransferResponseDataRelationships) SetLineItem(v LineItemOptionResponseDataRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *StockTransferResponseDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.

### GetEvents

`func (o *StockTransferResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StockTransferResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StockTransferResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StockTransferResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


