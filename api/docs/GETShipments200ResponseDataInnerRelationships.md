# GETShipments200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**GETAdyenPayments200ResponseDataInnerRelationshipsOrder**](GETAdyenPayments200ResponseDataInnerRelationshipsOrder.md) |  | [optional] 
**ShippingCategory** | Pointer to [**GETShipments200ResponseDataInnerRelationshipsShippingCategory**](GETShipments200ResponseDataInnerRelationshipsShippingCategory.md) |  | [optional] 
**StockLocation** | Pointer to [**GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation**](GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation.md) |  | [optional] 
**OriginAddress** | Pointer to [**GETReturns200ResponseDataInnerRelationshipsOriginAddress**](GETReturns200ResponseDataInnerRelationshipsOriginAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsShippingAddress**](GETOrders200ResponseDataInnerRelationshipsShippingAddress.md) |  | [optional] 
**ShippingMethod** | Pointer to [**GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod**](GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod.md) |  | [optional] 
**DeliveryLeadTime** | Pointer to [**GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime**](GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime.md) |  | [optional] 
**ShipmentLineItems** | Pointer to [**GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems**](GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems.md) |  | [optional] 
**StockLineItems** | Pointer to [**GETLineItems200ResponseDataInnerRelationshipsStockLineItems**](GETLineItems200ResponseDataInnerRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**GETLineItems200ResponseDataInnerRelationshipsStockTransfers**](GETLineItems200ResponseDataInnerRelationshipsStockTransfers.md) |  | [optional] 
**AvailableShippingMethods** | Pointer to [**GETShipments200ResponseDataInnerRelationshipsAvailableShippingMethods**](GETShipments200ResponseDataInnerRelationshipsAvailableShippingMethods.md) |  | [optional] 
**CarrierAccounts** | Pointer to [**GETShipments200ResponseDataInnerRelationshipsCarrierAccounts**](GETShipments200ResponseDataInnerRelationshipsCarrierAccounts.md) |  | [optional] 
**Parcels** | Pointer to [**GETPackages200ResponseDataInnerRelationshipsParcels**](GETPackages200ResponseDataInnerRelationshipsParcels.md) |  | [optional] 
**Attachments** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments**](GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**GETAuthorizations200ResponseDataInnerRelationshipsEvents**](GETAuthorizations200ResponseDataInnerRelationshipsEvents.md) |  | [optional] 

## Methods

### NewGETShipments200ResponseDataInnerRelationships

`func NewGETShipments200ResponseDataInnerRelationships() *GETShipments200ResponseDataInnerRelationships`

NewGETShipments200ResponseDataInnerRelationships instantiates a new GETShipments200ResponseDataInnerRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETShipments200ResponseDataInnerRelationshipsWithDefaults

`func NewGETShipments200ResponseDataInnerRelationshipsWithDefaults() *GETShipments200ResponseDataInnerRelationships`

NewGETShipments200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETShipments200ResponseDataInnerRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *GETShipments200ResponseDataInnerRelationships) GetOrder() GETAdyenPayments200ResponseDataInnerRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetOrderOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GETShipments200ResponseDataInnerRelationships) SetOrder(v GETAdyenPayments200ResponseDataInnerRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GETShipments200ResponseDataInnerRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetShippingCategory

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingCategory() GETShipments200ResponseDataInnerRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingCategoryOk() (*GETShipments200ResponseDataInnerRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *GETShipments200ResponseDataInnerRelationships) SetShippingCategory(v GETShipments200ResponseDataInnerRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *GETShipments200ResponseDataInnerRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetStockLocation

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *GETShipments200ResponseDataInnerRelationships) SetStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *GETShipments200ResponseDataInnerRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOriginAddress

`func (o *GETShipments200ResponseDataInnerRelationships) GetOriginAddress() GETReturns200ResponseDataInnerRelationshipsOriginAddress`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetOriginAddressOk() (*GETReturns200ResponseDataInnerRelationshipsOriginAddress, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *GETShipments200ResponseDataInnerRelationships) SetOriginAddress(v GETReturns200ResponseDataInnerRelationshipsOriginAddress)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *GETShipments200ResponseDataInnerRelationships) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingAddress() GETOrders200ResponseDataInnerRelationshipsShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingAddressOk() (*GETOrders200ResponseDataInnerRelationshipsShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *GETShipments200ResponseDataInnerRelationships) SetShippingAddress(v GETOrders200ResponseDataInnerRelationshipsShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *GETShipments200ResponseDataInnerRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetShippingMethod

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingMethod() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingMethodOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *GETShipments200ResponseDataInnerRelationships) SetShippingMethod(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *GETShipments200ResponseDataInnerRelationships) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetDeliveryLeadTime

`func (o *GETShipments200ResponseDataInnerRelationships) GetDeliveryLeadTime() GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime`

GetDeliveryLeadTime returns the DeliveryLeadTime field if non-nil, zero value otherwise.

### GetDeliveryLeadTimeOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetDeliveryLeadTimeOk() (*GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime, bool)`

GetDeliveryLeadTimeOk returns a tuple with the DeliveryLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTime

`func (o *GETShipments200ResponseDataInnerRelationships) SetDeliveryLeadTime(v GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime)`

SetDeliveryLeadTime sets DeliveryLeadTime field to given value.

### HasDeliveryLeadTime

`func (o *GETShipments200ResponseDataInnerRelationships) HasDeliveryLeadTime() bool`

HasDeliveryLeadTime returns a boolean if a field has been set.

### GetShipmentLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) GetShipmentLineItems() GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems`

GetShipmentLineItems returns the ShipmentLineItems field if non-nil, zero value otherwise.

### GetShipmentLineItemsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetShipmentLineItemsOk() (*GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems, bool)`

GetShipmentLineItemsOk returns a tuple with the ShipmentLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) SetShipmentLineItems(v GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems)`

SetShipmentLineItems sets ShipmentLineItems field to given value.

### HasShipmentLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) HasShipmentLineItems() bool`

HasShipmentLineItems returns a boolean if a field has been set.

### GetStockLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockLineItems() GETLineItems200ResponseDataInnerRelationshipsStockLineItems`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockLineItemsOk() (*GETLineItems200ResponseDataInnerRelationshipsStockLineItems, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) SetStockLineItems(v GETLineItems200ResponseDataInnerRelationshipsStockLineItems)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockTransfers() GETLineItems200ResponseDataInnerRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockTransfersOk() (*GETLineItems200ResponseDataInnerRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *GETShipments200ResponseDataInnerRelationships) SetStockTransfers(v GETLineItems200ResponseDataInnerRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *GETShipments200ResponseDataInnerRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetAvailableShippingMethods

`func (o *GETShipments200ResponseDataInnerRelationships) GetAvailableShippingMethods() GETShipments200ResponseDataInnerRelationshipsAvailableShippingMethods`

GetAvailableShippingMethods returns the AvailableShippingMethods field if non-nil, zero value otherwise.

### GetAvailableShippingMethodsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetAvailableShippingMethodsOk() (*GETShipments200ResponseDataInnerRelationshipsAvailableShippingMethods, bool)`

GetAvailableShippingMethodsOk returns a tuple with the AvailableShippingMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShippingMethods

`func (o *GETShipments200ResponseDataInnerRelationships) SetAvailableShippingMethods(v GETShipments200ResponseDataInnerRelationshipsAvailableShippingMethods)`

SetAvailableShippingMethods sets AvailableShippingMethods field to given value.

### HasAvailableShippingMethods

`func (o *GETShipments200ResponseDataInnerRelationships) HasAvailableShippingMethods() bool`

HasAvailableShippingMethods returns a boolean if a field has been set.

### GetCarrierAccounts

`func (o *GETShipments200ResponseDataInnerRelationships) GetCarrierAccounts() GETShipments200ResponseDataInnerRelationshipsCarrierAccounts`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetCarrierAccountsOk() (*GETShipments200ResponseDataInnerRelationshipsCarrierAccounts, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *GETShipments200ResponseDataInnerRelationships) SetCarrierAccounts(v GETShipments200ResponseDataInnerRelationshipsCarrierAccounts)`

SetCarrierAccounts sets CarrierAccounts field to given value.

### HasCarrierAccounts

`func (o *GETShipments200ResponseDataInnerRelationships) HasCarrierAccounts() bool`

HasCarrierAccounts returns a boolean if a field has been set.

### GetParcels

`func (o *GETShipments200ResponseDataInnerRelationships) GetParcels() GETPackages200ResponseDataInnerRelationshipsParcels`

GetParcels returns the Parcels field if non-nil, zero value otherwise.

### GetParcelsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetParcelsOk() (*GETPackages200ResponseDataInnerRelationshipsParcels, bool)`

GetParcelsOk returns a tuple with the Parcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcels

`func (o *GETShipments200ResponseDataInnerRelationships) SetParcels(v GETPackages200ResponseDataInnerRelationshipsParcels)`

SetParcels sets Parcels field to given value.

### HasParcels

`func (o *GETShipments200ResponseDataInnerRelationships) HasParcels() bool`

HasParcels returns a boolean if a field has been set.

### GetAttachments

`func (o *GETShipments200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETShipments200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETShipments200ResponseDataInnerRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GETShipments200ResponseDataInnerRelationships) GetEvents() GETAuthorizations200ResponseDataInnerRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetEventsOk() (*GETAuthorizations200ResponseDataInnerRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETShipments200ResponseDataInnerRelationships) SetEvents(v GETAuthorizations200ResponseDataInnerRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETShipments200ResponseDataInnerRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


