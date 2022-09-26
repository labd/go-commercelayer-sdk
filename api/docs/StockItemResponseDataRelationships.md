# StockItemResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StockLocation** | Pointer to [**DeliveryLeadTimeResponseDataRelationshipsStockLocation**](DeliveryLeadTimeResponseDataRelationshipsStockLocation.md) |  | [optional] 
**Sku** | Pointer to [**InStockSubscriptionResponseDataRelationshipsSku**](InStockSubscriptionResponseDataRelationshipsSku.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewStockItemResponseDataRelationships

`func NewStockItemResponseDataRelationships() *StockItemResponseDataRelationships`

NewStockItemResponseDataRelationships instantiates a new StockItemResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockItemResponseDataRelationshipsWithDefaults

`func NewStockItemResponseDataRelationshipsWithDefaults() *StockItemResponseDataRelationships`

NewStockItemResponseDataRelationshipsWithDefaults instantiates a new StockItemResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStockLocation

`func (o *StockItemResponseDataRelationships) GetStockLocation() DeliveryLeadTimeResponseDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *StockItemResponseDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeResponseDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *StockItemResponseDataRelationships) SetStockLocation(v DeliveryLeadTimeResponseDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *StockItemResponseDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetSku

`func (o *StockItemResponseDataRelationships) GetSku() InStockSubscriptionResponseDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StockItemResponseDataRelationships) GetSkuOk() (*InStockSubscriptionResponseDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StockItemResponseDataRelationships) SetSku(v InStockSubscriptionResponseDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *StockItemResponseDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetAttachments

`func (o *StockItemResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *StockItemResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *StockItemResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *StockItemResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


