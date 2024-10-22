# POSTStockItems201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StockLocation** | Pointer to [**POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation**](POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation.md) |  | [optional] 
**Sku** | Pointer to [**POSTInStockSubscriptions201ResponseDataRelationshipsSku**](POSTInStockSubscriptions201ResponseDataRelationshipsSku.md) |  | [optional] 
**ReservedStock** | Pointer to [**POSTStockItems201ResponseDataRelationshipsReservedStock**](POSTStockItems201ResponseDataRelationshipsReservedStock.md) |  | [optional] 
**StockReservations** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockReservations**](POSTLineItems201ResponseDataRelationshipsStockReservations.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTStockItems201ResponseDataRelationships

`func NewPOSTStockItems201ResponseDataRelationships() *POSTStockItems201ResponseDataRelationships`

NewPOSTStockItems201ResponseDataRelationships instantiates a new POSTStockItems201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockItems201ResponseDataRelationshipsWithDefaults

`func NewPOSTStockItems201ResponseDataRelationshipsWithDefaults() *POSTStockItems201ResponseDataRelationships`

NewPOSTStockItems201ResponseDataRelationshipsWithDefaults instantiates a new POSTStockItems201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStockLocation

`func (o *POSTStockItems201ResponseDataRelationships) GetStockLocation() POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *POSTStockItems201ResponseDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *POSTStockItems201ResponseDataRelationships) SetStockLocation(v POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *POSTStockItems201ResponseDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetSku

`func (o *POSTStockItems201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *POSTStockItems201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *POSTStockItems201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *POSTStockItems201ResponseDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetReservedStock

`func (o *POSTStockItems201ResponseDataRelationships) GetReservedStock() POSTStockItems201ResponseDataRelationshipsReservedStock`

GetReservedStock returns the ReservedStock field if non-nil, zero value otherwise.

### GetReservedStockOk

`func (o *POSTStockItems201ResponseDataRelationships) GetReservedStockOk() (*POSTStockItems201ResponseDataRelationshipsReservedStock, bool)`

GetReservedStockOk returns a tuple with the ReservedStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStock

`func (o *POSTStockItems201ResponseDataRelationships) SetReservedStock(v POSTStockItems201ResponseDataRelationshipsReservedStock)`

SetReservedStock sets ReservedStock field to given value.

### HasReservedStock

`func (o *POSTStockItems201ResponseDataRelationships) HasReservedStock() bool`

HasReservedStock returns a boolean if a field has been set.

### GetStockReservations

`func (o *POSTStockItems201ResponseDataRelationships) GetStockReservations() POSTLineItems201ResponseDataRelationshipsStockReservations`

GetStockReservations returns the StockReservations field if non-nil, zero value otherwise.

### GetStockReservationsOk

`func (o *POSTStockItems201ResponseDataRelationships) GetStockReservationsOk() (*POSTLineItems201ResponseDataRelationshipsStockReservations, bool)`

GetStockReservationsOk returns a tuple with the StockReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservations

`func (o *POSTStockItems201ResponseDataRelationships) SetStockReservations(v POSTLineItems201ResponseDataRelationshipsStockReservations)`

SetStockReservations sets StockReservations field to given value.

### HasStockReservations

`func (o *POSTStockItems201ResponseDataRelationships) HasStockReservations() bool`

HasStockReservations returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTStockItems201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTStockItems201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTStockItems201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTStockItems201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *POSTStockItems201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTStockItems201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTStockItems201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTStockItems201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


