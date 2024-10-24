# StockItemDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StockLocation** | Pointer to [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | [optional] 
**Sku** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 
**ReservedStock** | Pointer to [**StockItemDataRelationshipsReservedStock**](StockItemDataRelationshipsReservedStock.md) |  | [optional] 
**StockReservations** | Pointer to [**LineItemDataRelationshipsStockReservations**](LineItemDataRelationshipsStockReservations.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewStockItemDataRelationships

`func NewStockItemDataRelationships() *StockItemDataRelationships`

NewStockItemDataRelationships instantiates a new StockItemDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockItemDataRelationshipsWithDefaults

`func NewStockItemDataRelationshipsWithDefaults() *StockItemDataRelationships`

NewStockItemDataRelationshipsWithDefaults instantiates a new StockItemDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStockLocation

`func (o *StockItemDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *StockItemDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *StockItemDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *StockItemDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetSku

`func (o *StockItemDataRelationships) GetSku() BundleDataRelationshipsSkus`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StockItemDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StockItemDataRelationships) SetSku(v BundleDataRelationshipsSkus)`

SetSku sets Sku field to given value.

### HasSku

`func (o *StockItemDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetReservedStock

`func (o *StockItemDataRelationships) GetReservedStock() StockItemDataRelationshipsReservedStock`

GetReservedStock returns the ReservedStock field if non-nil, zero value otherwise.

### GetReservedStockOk

`func (o *StockItemDataRelationships) GetReservedStockOk() (*StockItemDataRelationshipsReservedStock, bool)`

GetReservedStockOk returns a tuple with the ReservedStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStock

`func (o *StockItemDataRelationships) SetReservedStock(v StockItemDataRelationshipsReservedStock)`

SetReservedStock sets ReservedStock field to given value.

### HasReservedStock

`func (o *StockItemDataRelationships) HasReservedStock() bool`

HasReservedStock returns a boolean if a field has been set.

### GetStockReservations

`func (o *StockItemDataRelationships) GetStockReservations() LineItemDataRelationshipsStockReservations`

GetStockReservations returns the StockReservations field if non-nil, zero value otherwise.

### GetStockReservationsOk

`func (o *StockItemDataRelationships) GetStockReservationsOk() (*LineItemDataRelationshipsStockReservations, bool)`

GetStockReservationsOk returns a tuple with the StockReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservations

`func (o *StockItemDataRelationships) SetStockReservations(v LineItemDataRelationshipsStockReservations)`

SetStockReservations sets StockReservations field to given value.

### HasStockReservations

`func (o *StockItemDataRelationships) HasStockReservations() bool`

HasStockReservations returns a boolean if a field has been set.

### GetAttachments

`func (o *StockItemDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *StockItemDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *StockItemDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *StockItemDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *StockItemDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *StockItemDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *StockItemDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *StockItemDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


