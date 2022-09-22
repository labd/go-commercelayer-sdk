# GETStockLocations200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**GETCustomerAddresses200ResponseDataInnerRelationshipsAddress**](GETCustomerAddresses200ResponseDataInnerRelationshipsAddress.md) |  | [optional] 
**InventoryStockLocations** | Pointer to [**GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations**](GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations.md) |  | [optional] 
**InventoryReturnLocations** | Pointer to [**GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations**](GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations.md) |  | [optional] 
**StockItems** | Pointer to [**GETSkus200ResponseDataInnerRelationshipsStockItems**](GETSkus200ResponseDataInnerRelationshipsStockItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**GETLineItems200ResponseDataInnerRelationshipsStockTransfers**](GETLineItems200ResponseDataInnerRelationshipsStockTransfers.md) |  | [optional] 
**Attachments** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments**](GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewGETStockLocations200ResponseDataInnerRelationships

`func NewGETStockLocations200ResponseDataInnerRelationships() *GETStockLocations200ResponseDataInnerRelationships`

NewGETStockLocations200ResponseDataInnerRelationships instantiates a new GETStockLocations200ResponseDataInnerRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETStockLocations200ResponseDataInnerRelationshipsWithDefaults

`func NewGETStockLocations200ResponseDataInnerRelationshipsWithDefaults() *GETStockLocations200ResponseDataInnerRelationships`

NewGETStockLocations200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETStockLocations200ResponseDataInnerRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetAddress() GETCustomerAddresses200ResponseDataInnerRelationshipsAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetAddressOk() (*GETCustomerAddresses200ResponseDataInnerRelationshipsAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GETStockLocations200ResponseDataInnerRelationships) SetAddress(v GETCustomerAddresses200ResponseDataInnerRelationshipsAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GETStockLocations200ResponseDataInnerRelationships) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInventoryStockLocations

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetInventoryStockLocations() GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations`

GetInventoryStockLocations returns the InventoryStockLocations field if non-nil, zero value otherwise.

### GetInventoryStockLocationsOk

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetInventoryStockLocationsOk() (*GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations, bool)`

GetInventoryStockLocationsOk returns a tuple with the InventoryStockLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryStockLocations

`func (o *GETStockLocations200ResponseDataInnerRelationships) SetInventoryStockLocations(v GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations)`

SetInventoryStockLocations sets InventoryStockLocations field to given value.

### HasInventoryStockLocations

`func (o *GETStockLocations200ResponseDataInnerRelationships) HasInventoryStockLocations() bool`

HasInventoryStockLocations returns a boolean if a field has been set.

### GetInventoryReturnLocations

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetInventoryReturnLocations() GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations`

GetInventoryReturnLocations returns the InventoryReturnLocations field if non-nil, zero value otherwise.

### GetInventoryReturnLocationsOk

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetInventoryReturnLocationsOk() (*GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations, bool)`

GetInventoryReturnLocationsOk returns a tuple with the InventoryReturnLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReturnLocations

`func (o *GETStockLocations200ResponseDataInnerRelationships) SetInventoryReturnLocations(v GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations)`

SetInventoryReturnLocations sets InventoryReturnLocations field to given value.

### HasInventoryReturnLocations

`func (o *GETStockLocations200ResponseDataInnerRelationships) HasInventoryReturnLocations() bool`

HasInventoryReturnLocations returns a boolean if a field has been set.

### GetStockItems

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetStockItems() GETSkus200ResponseDataInnerRelationshipsStockItems`

GetStockItems returns the StockItems field if non-nil, zero value otherwise.

### GetStockItemsOk

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetStockItemsOk() (*GETSkus200ResponseDataInnerRelationshipsStockItems, bool)`

GetStockItemsOk returns a tuple with the StockItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItems

`func (o *GETStockLocations200ResponseDataInnerRelationships) SetStockItems(v GETSkus200ResponseDataInnerRelationshipsStockItems)`

SetStockItems sets StockItems field to given value.

### HasStockItems

`func (o *GETStockLocations200ResponseDataInnerRelationships) HasStockItems() bool`

HasStockItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetStockTransfers() GETLineItems200ResponseDataInnerRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetStockTransfersOk() (*GETLineItems200ResponseDataInnerRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *GETStockLocations200ResponseDataInnerRelationships) SetStockTransfers(v GETLineItems200ResponseDataInnerRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *GETStockLocations200ResponseDataInnerRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetAttachments

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETStockLocations200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETStockLocations200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETStockLocations200ResponseDataInnerRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


