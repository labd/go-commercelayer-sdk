# StockLocationResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**CustomerAddressResponseDataRelationshipsAddress**](CustomerAddressResponseDataRelationshipsAddress.md) |  | [optional] 
**InventoryStockLocations** | Pointer to [**InventoryModelResponseDataRelationshipsInventoryStockLocations**](InventoryModelResponseDataRelationshipsInventoryStockLocations.md) |  | [optional] 
**InventoryReturnLocations** | Pointer to [**InventoryModelResponseDataRelationshipsInventoryReturnLocations**](InventoryModelResponseDataRelationshipsInventoryReturnLocations.md) |  | [optional] 
**StockItems** | Pointer to [**SkuResponseDataRelationshipsStockItems**](SkuResponseDataRelationshipsStockItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**LineItemResponseDataRelationshipsStockTransfers**](LineItemResponseDataRelationshipsStockTransfers.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewStockLocationResponseDataRelationships

`func NewStockLocationResponseDataRelationships() *StockLocationResponseDataRelationships`

NewStockLocationResponseDataRelationships instantiates a new StockLocationResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLocationResponseDataRelationshipsWithDefaults

`func NewStockLocationResponseDataRelationshipsWithDefaults() *StockLocationResponseDataRelationships`

NewStockLocationResponseDataRelationshipsWithDefaults instantiates a new StockLocationResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *StockLocationResponseDataRelationships) GetAddress() CustomerAddressResponseDataRelationshipsAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StockLocationResponseDataRelationships) GetAddressOk() (*CustomerAddressResponseDataRelationshipsAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StockLocationResponseDataRelationships) SetAddress(v CustomerAddressResponseDataRelationshipsAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *StockLocationResponseDataRelationships) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInventoryStockLocations

`func (o *StockLocationResponseDataRelationships) GetInventoryStockLocations() InventoryModelResponseDataRelationshipsInventoryStockLocations`

GetInventoryStockLocations returns the InventoryStockLocations field if non-nil, zero value otherwise.

### GetInventoryStockLocationsOk

`func (o *StockLocationResponseDataRelationships) GetInventoryStockLocationsOk() (*InventoryModelResponseDataRelationshipsInventoryStockLocations, bool)`

GetInventoryStockLocationsOk returns a tuple with the InventoryStockLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryStockLocations

`func (o *StockLocationResponseDataRelationships) SetInventoryStockLocations(v InventoryModelResponseDataRelationshipsInventoryStockLocations)`

SetInventoryStockLocations sets InventoryStockLocations field to given value.

### HasInventoryStockLocations

`func (o *StockLocationResponseDataRelationships) HasInventoryStockLocations() bool`

HasInventoryStockLocations returns a boolean if a field has been set.

### GetInventoryReturnLocations

`func (o *StockLocationResponseDataRelationships) GetInventoryReturnLocations() InventoryModelResponseDataRelationshipsInventoryReturnLocations`

GetInventoryReturnLocations returns the InventoryReturnLocations field if non-nil, zero value otherwise.

### GetInventoryReturnLocationsOk

`func (o *StockLocationResponseDataRelationships) GetInventoryReturnLocationsOk() (*InventoryModelResponseDataRelationshipsInventoryReturnLocations, bool)`

GetInventoryReturnLocationsOk returns a tuple with the InventoryReturnLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReturnLocations

`func (o *StockLocationResponseDataRelationships) SetInventoryReturnLocations(v InventoryModelResponseDataRelationshipsInventoryReturnLocations)`

SetInventoryReturnLocations sets InventoryReturnLocations field to given value.

### HasInventoryReturnLocations

`func (o *StockLocationResponseDataRelationships) HasInventoryReturnLocations() bool`

HasInventoryReturnLocations returns a boolean if a field has been set.

### GetStockItems

`func (o *StockLocationResponseDataRelationships) GetStockItems() SkuResponseDataRelationshipsStockItems`

GetStockItems returns the StockItems field if non-nil, zero value otherwise.

### GetStockItemsOk

`func (o *StockLocationResponseDataRelationships) GetStockItemsOk() (*SkuResponseDataRelationshipsStockItems, bool)`

GetStockItemsOk returns a tuple with the StockItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItems

`func (o *StockLocationResponseDataRelationships) SetStockItems(v SkuResponseDataRelationshipsStockItems)`

SetStockItems sets StockItems field to given value.

### HasStockItems

`func (o *StockLocationResponseDataRelationships) HasStockItems() bool`

HasStockItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *StockLocationResponseDataRelationships) GetStockTransfers() LineItemResponseDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *StockLocationResponseDataRelationships) GetStockTransfersOk() (*LineItemResponseDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *StockLocationResponseDataRelationships) SetStockTransfers(v LineItemResponseDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *StockLocationResponseDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetAttachments

`func (o *StockLocationResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *StockLocationResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *StockLocationResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *StockLocationResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


