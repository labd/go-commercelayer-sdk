# GETShipmentsShipmentId200ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**POSTAdyenPayments201ResponseDataRelationshipsOrder**](POSTAdyenPayments201ResponseDataRelationshipsOrder.md) |  | [optional] 
**ShippingCategory** | Pointer to [**GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory**](GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory.md) |  | [optional] 
**StockLocation** | Pointer to [**POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation**](POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation.md) |  | [optional] 
**OriginAddress** | Pointer to [**POSTReturns201ResponseDataRelationshipsOriginAddress**](POSTReturns201ResponseDataRelationshipsOriginAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**POSTOrders201ResponseDataRelationshipsShippingAddress**](POSTOrders201ResponseDataRelationshipsShippingAddress.md) |  | [optional] 
**ShippingMethod** | Pointer to [**POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethod**](POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethod.md) |  | [optional] 
**DeliveryLeadTime** | Pointer to [**GETShipmentsShipmentId200ResponseDataRelationshipsDeliveryLeadTime**](GETShipmentsShipmentId200ResponseDataRelationshipsDeliveryLeadTime.md) |  | [optional] 
**ShipmentLineItems** | Pointer to [**POSTLineItems201ResponseDataRelationshipsShipmentLineItems**](POSTLineItems201ResponseDataRelationshipsShipmentLineItems.md) |  | [optional] 
**StockLineItems** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockLineItems**](POSTLineItems201ResponseDataRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockTransfers**](POSTLineItems201ResponseDataRelationshipsStockTransfers.md) |  | [optional] 
**AvailableShippingMethods** | Pointer to [**GETShipmentsShipmentId200ResponseDataRelationshipsAvailableShippingMethods**](GETShipmentsShipmentId200ResponseDataRelationshipsAvailableShippingMethods.md) |  | [optional] 
**CarrierAccounts** | Pointer to [**GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts**](GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts.md) |  | [optional] 
**Parcels** | Pointer to [**POSTPackages201ResponseDataRelationshipsParcels**](POSTPackages201ResponseDataRelationshipsParcels.md) |  | [optional] 
**Attachments** | Pointer to [**POSTAvalaraAccounts201ResponseDataRelationshipsAttachments**](POSTAvalaraAccounts201ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewGETShipmentsShipmentId200ResponseDataRelationships

`func NewGETShipmentsShipmentId200ResponseDataRelationships() *GETShipmentsShipmentId200ResponseDataRelationships`

NewGETShipmentsShipmentId200ResponseDataRelationships instantiates a new GETShipmentsShipmentId200ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETShipmentsShipmentId200ResponseDataRelationshipsWithDefaults

`func NewGETShipmentsShipmentId200ResponseDataRelationshipsWithDefaults() *GETShipmentsShipmentId200ResponseDataRelationships`

NewGETShipmentsShipmentId200ResponseDataRelationshipsWithDefaults instantiates a new GETShipmentsShipmentId200ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetShippingCategory

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetShippingCategory() GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetShippingCategoryOk() (*GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetShippingCategory(v GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetStockLocation

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetStockLocation() POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetStockLocation(v POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOriginAddress

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetOriginAddress() POSTReturns201ResponseDataRelationshipsOriginAddress`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetOriginAddressOk() (*POSTReturns201ResponseDataRelationshipsOriginAddress, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetOriginAddress(v POSTReturns201ResponseDataRelationshipsOriginAddress)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetShippingAddress() POSTOrders201ResponseDataRelationshipsShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetShippingAddressOk() (*POSTOrders201ResponseDataRelationshipsShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetShippingAddress(v POSTOrders201ResponseDataRelationshipsShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetShippingMethod

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetShippingMethod() POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetShippingMethodOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetShippingMethod(v POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetDeliveryLeadTime

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetDeliveryLeadTime() GETShipmentsShipmentId200ResponseDataRelationshipsDeliveryLeadTime`

GetDeliveryLeadTime returns the DeliveryLeadTime field if non-nil, zero value otherwise.

### GetDeliveryLeadTimeOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetDeliveryLeadTimeOk() (*GETShipmentsShipmentId200ResponseDataRelationshipsDeliveryLeadTime, bool)`

GetDeliveryLeadTimeOk returns a tuple with the DeliveryLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTime

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetDeliveryLeadTime(v GETShipmentsShipmentId200ResponseDataRelationshipsDeliveryLeadTime)`

SetDeliveryLeadTime sets DeliveryLeadTime field to given value.

### HasDeliveryLeadTime

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasDeliveryLeadTime() bool`

HasDeliveryLeadTime returns a boolean if a field has been set.

### GetShipmentLineItems

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetShipmentLineItems() POSTLineItems201ResponseDataRelationshipsShipmentLineItems`

GetShipmentLineItems returns the ShipmentLineItems field if non-nil, zero value otherwise.

### GetShipmentLineItemsOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetShipmentLineItemsOk() (*POSTLineItems201ResponseDataRelationshipsShipmentLineItems, bool)`

GetShipmentLineItemsOk returns a tuple with the ShipmentLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLineItems

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetShipmentLineItems(v POSTLineItems201ResponseDataRelationshipsShipmentLineItems)`

SetShipmentLineItems sets ShipmentLineItems field to given value.

### HasShipmentLineItems

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasShipmentLineItems() bool`

HasShipmentLineItems returns a boolean if a field has been set.

### GetStockLineItems

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetStockLineItems() POSTLineItems201ResponseDataRelationshipsStockLineItems`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetStockLineItemsOk() (*POSTLineItems201ResponseDataRelationshipsStockLineItems, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetStockLineItems(v POSTLineItems201ResponseDataRelationshipsStockLineItems)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetStockTransfers() POSTLineItems201ResponseDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetStockTransfersOk() (*POSTLineItems201ResponseDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetStockTransfers(v POSTLineItems201ResponseDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetAvailableShippingMethods

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetAvailableShippingMethods() GETShipmentsShipmentId200ResponseDataRelationshipsAvailableShippingMethods`

GetAvailableShippingMethods returns the AvailableShippingMethods field if non-nil, zero value otherwise.

### GetAvailableShippingMethodsOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetAvailableShippingMethodsOk() (*GETShipmentsShipmentId200ResponseDataRelationshipsAvailableShippingMethods, bool)`

GetAvailableShippingMethodsOk returns a tuple with the AvailableShippingMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShippingMethods

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetAvailableShippingMethods(v GETShipmentsShipmentId200ResponseDataRelationshipsAvailableShippingMethods)`

SetAvailableShippingMethods sets AvailableShippingMethods field to given value.

### HasAvailableShippingMethods

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasAvailableShippingMethods() bool`

HasAvailableShippingMethods returns a boolean if a field has been set.

### GetCarrierAccounts

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetCarrierAccounts() GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetCarrierAccountsOk() (*GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetCarrierAccounts(v GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts)`

SetCarrierAccounts sets CarrierAccounts field to given value.

### HasCarrierAccounts

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasCarrierAccounts() bool`

HasCarrierAccounts returns a boolean if a field has been set.

### GetParcels

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetParcels() POSTPackages201ResponseDataRelationshipsParcels`

GetParcels returns the Parcels field if non-nil, zero value otherwise.

### GetParcelsOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetParcelsOk() (*POSTPackages201ResponseDataRelationshipsParcels, bool)`

GetParcelsOk returns a tuple with the Parcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcels

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetParcels(v POSTPackages201ResponseDataRelationshipsParcels)`

SetParcels sets Parcels field to given value.

### HasParcels

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasParcels() bool`

HasParcels returns a boolean if a field has been set.

### GetAttachments

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetAttachments() POSTAvalaraAccounts201ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetAttachmentsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetAttachments(v POSTAvalaraAccounts201ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetEvents() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) GetEventsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) SetEvents(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETShipmentsShipmentId200ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


