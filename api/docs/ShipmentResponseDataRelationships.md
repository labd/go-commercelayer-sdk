# ShipmentResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentResponseDataRelationshipsOrder**](AdyenPaymentResponseDataRelationshipsOrder.md) |  | [optional] 
**ShippingCategory** | Pointer to [**ShipmentResponseDataRelationshipsShippingCategory**](ShipmentResponseDataRelationshipsShippingCategory.md) |  | [optional] 
**StockLocation** | Pointer to [**DeliveryLeadTimeResponseDataRelationshipsStockLocation**](DeliveryLeadTimeResponseDataRelationshipsStockLocation.md) |  | [optional] 
**OriginAddress** | Pointer to [**ReturnResponseDataRelationshipsOriginAddress**](ReturnResponseDataRelationshipsOriginAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**OrderResponseDataRelationshipsShippingAddress**](OrderResponseDataRelationshipsShippingAddress.md) |  | [optional] 
**ShippingMethod** | Pointer to [**DeliveryLeadTimeResponseDataRelationshipsShippingMethod**](DeliveryLeadTimeResponseDataRelationshipsShippingMethod.md) |  | [optional] 
**DeliveryLeadTime** | Pointer to [**ShipmentResponseDataRelationshipsDeliveryLeadTime**](ShipmentResponseDataRelationshipsDeliveryLeadTime.md) |  | [optional] 
**ShipmentLineItems** | Pointer to [**LineItemResponseDataRelationshipsShipmentLineItems**](LineItemResponseDataRelationshipsShipmentLineItems.md) |  | [optional] 
**StockLineItems** | Pointer to [**LineItemResponseDataRelationshipsStockLineItems**](LineItemResponseDataRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**LineItemResponseDataRelationshipsStockTransfers**](LineItemResponseDataRelationshipsStockTransfers.md) |  | [optional] 
**AvailableShippingMethods** | Pointer to [**ShipmentResponseDataRelationshipsAvailableShippingMethods**](ShipmentResponseDataRelationshipsAvailableShippingMethods.md) |  | [optional] 
**CarrierAccounts** | Pointer to [**ShipmentResponseDataRelationshipsCarrierAccounts**](ShipmentResponseDataRelationshipsCarrierAccounts.md) |  | [optional] 
**Parcels** | Pointer to [**PackageResponseDataRelationshipsParcels**](PackageResponseDataRelationshipsParcels.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressResponseDataRelationshipsEvents**](CustomerAddressResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewShipmentResponseDataRelationships

`func NewShipmentResponseDataRelationships() *ShipmentResponseDataRelationships`

NewShipmentResponseDataRelationships instantiates a new ShipmentResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentResponseDataRelationshipsWithDefaults

`func NewShipmentResponseDataRelationshipsWithDefaults() *ShipmentResponseDataRelationships`

NewShipmentResponseDataRelationshipsWithDefaults instantiates a new ShipmentResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *ShipmentResponseDataRelationships) GetOrder() AdyenPaymentResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ShipmentResponseDataRelationships) GetOrderOk() (*AdyenPaymentResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ShipmentResponseDataRelationships) SetOrder(v AdyenPaymentResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ShipmentResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetShippingCategory

`func (o *ShipmentResponseDataRelationships) GetShippingCategory() ShipmentResponseDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *ShipmentResponseDataRelationships) GetShippingCategoryOk() (*ShipmentResponseDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *ShipmentResponseDataRelationships) SetShippingCategory(v ShipmentResponseDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *ShipmentResponseDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetStockLocation

`func (o *ShipmentResponseDataRelationships) GetStockLocation() DeliveryLeadTimeResponseDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *ShipmentResponseDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeResponseDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *ShipmentResponseDataRelationships) SetStockLocation(v DeliveryLeadTimeResponseDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *ShipmentResponseDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOriginAddress

`func (o *ShipmentResponseDataRelationships) GetOriginAddress() ReturnResponseDataRelationshipsOriginAddress`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *ShipmentResponseDataRelationships) GetOriginAddressOk() (*ReturnResponseDataRelationshipsOriginAddress, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *ShipmentResponseDataRelationships) SetOriginAddress(v ReturnResponseDataRelationshipsOriginAddress)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *ShipmentResponseDataRelationships) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *ShipmentResponseDataRelationships) GetShippingAddress() OrderResponseDataRelationshipsShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *ShipmentResponseDataRelationships) GetShippingAddressOk() (*OrderResponseDataRelationshipsShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *ShipmentResponseDataRelationships) SetShippingAddress(v OrderResponseDataRelationshipsShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *ShipmentResponseDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetShippingMethod

`func (o *ShipmentResponseDataRelationships) GetShippingMethod() DeliveryLeadTimeResponseDataRelationshipsShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *ShipmentResponseDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeResponseDataRelationshipsShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *ShipmentResponseDataRelationships) SetShippingMethod(v DeliveryLeadTimeResponseDataRelationshipsShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *ShipmentResponseDataRelationships) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetDeliveryLeadTime

`func (o *ShipmentResponseDataRelationships) GetDeliveryLeadTime() ShipmentResponseDataRelationshipsDeliveryLeadTime`

GetDeliveryLeadTime returns the DeliveryLeadTime field if non-nil, zero value otherwise.

### GetDeliveryLeadTimeOk

`func (o *ShipmentResponseDataRelationships) GetDeliveryLeadTimeOk() (*ShipmentResponseDataRelationshipsDeliveryLeadTime, bool)`

GetDeliveryLeadTimeOk returns a tuple with the DeliveryLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTime

`func (o *ShipmentResponseDataRelationships) SetDeliveryLeadTime(v ShipmentResponseDataRelationshipsDeliveryLeadTime)`

SetDeliveryLeadTime sets DeliveryLeadTime field to given value.

### HasDeliveryLeadTime

`func (o *ShipmentResponseDataRelationships) HasDeliveryLeadTime() bool`

HasDeliveryLeadTime returns a boolean if a field has been set.

### GetShipmentLineItems

`func (o *ShipmentResponseDataRelationships) GetShipmentLineItems() LineItemResponseDataRelationshipsShipmentLineItems`

GetShipmentLineItems returns the ShipmentLineItems field if non-nil, zero value otherwise.

### GetShipmentLineItemsOk

`func (o *ShipmentResponseDataRelationships) GetShipmentLineItemsOk() (*LineItemResponseDataRelationshipsShipmentLineItems, bool)`

GetShipmentLineItemsOk returns a tuple with the ShipmentLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLineItems

`func (o *ShipmentResponseDataRelationships) SetShipmentLineItems(v LineItemResponseDataRelationshipsShipmentLineItems)`

SetShipmentLineItems sets ShipmentLineItems field to given value.

### HasShipmentLineItems

`func (o *ShipmentResponseDataRelationships) HasShipmentLineItems() bool`

HasShipmentLineItems returns a boolean if a field has been set.

### GetStockLineItems

`func (o *ShipmentResponseDataRelationships) GetStockLineItems() LineItemResponseDataRelationshipsStockLineItems`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *ShipmentResponseDataRelationships) GetStockLineItemsOk() (*LineItemResponseDataRelationshipsStockLineItems, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *ShipmentResponseDataRelationships) SetStockLineItems(v LineItemResponseDataRelationshipsStockLineItems)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *ShipmentResponseDataRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *ShipmentResponseDataRelationships) GetStockTransfers() LineItemResponseDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *ShipmentResponseDataRelationships) GetStockTransfersOk() (*LineItemResponseDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *ShipmentResponseDataRelationships) SetStockTransfers(v LineItemResponseDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *ShipmentResponseDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetAvailableShippingMethods

`func (o *ShipmentResponseDataRelationships) GetAvailableShippingMethods() ShipmentResponseDataRelationshipsAvailableShippingMethods`

GetAvailableShippingMethods returns the AvailableShippingMethods field if non-nil, zero value otherwise.

### GetAvailableShippingMethodsOk

`func (o *ShipmentResponseDataRelationships) GetAvailableShippingMethodsOk() (*ShipmentResponseDataRelationshipsAvailableShippingMethods, bool)`

GetAvailableShippingMethodsOk returns a tuple with the AvailableShippingMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShippingMethods

`func (o *ShipmentResponseDataRelationships) SetAvailableShippingMethods(v ShipmentResponseDataRelationshipsAvailableShippingMethods)`

SetAvailableShippingMethods sets AvailableShippingMethods field to given value.

### HasAvailableShippingMethods

`func (o *ShipmentResponseDataRelationships) HasAvailableShippingMethods() bool`

HasAvailableShippingMethods returns a boolean if a field has been set.

### GetCarrierAccounts

`func (o *ShipmentResponseDataRelationships) GetCarrierAccounts() ShipmentResponseDataRelationshipsCarrierAccounts`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *ShipmentResponseDataRelationships) GetCarrierAccountsOk() (*ShipmentResponseDataRelationshipsCarrierAccounts, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *ShipmentResponseDataRelationships) SetCarrierAccounts(v ShipmentResponseDataRelationshipsCarrierAccounts)`

SetCarrierAccounts sets CarrierAccounts field to given value.

### HasCarrierAccounts

`func (o *ShipmentResponseDataRelationships) HasCarrierAccounts() bool`

HasCarrierAccounts returns a boolean if a field has been set.

### GetParcels

`func (o *ShipmentResponseDataRelationships) GetParcels() PackageResponseDataRelationshipsParcels`

GetParcels returns the Parcels field if non-nil, zero value otherwise.

### GetParcelsOk

`func (o *ShipmentResponseDataRelationships) GetParcelsOk() (*PackageResponseDataRelationshipsParcels, bool)`

GetParcelsOk returns a tuple with the Parcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcels

`func (o *ShipmentResponseDataRelationships) SetParcels(v PackageResponseDataRelationshipsParcels)`

SetParcels sets Parcels field to given value.

### HasParcels

`func (o *ShipmentResponseDataRelationships) HasParcels() bool`

HasParcels returns a boolean if a field has been set.

### GetAttachments

`func (o *ShipmentResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ShipmentResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ShipmentResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ShipmentResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *ShipmentResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ShipmentResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ShipmentResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ShipmentResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


