# POSTShipments201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**POSTAdyenPayments201ResponseDataRelationshipsOrder**](POSTAdyenPayments201ResponseDataRelationshipsOrder.md) |  | [optional] 
**ShippingCategory** | Pointer to [**POSTShipments201ResponseDataRelationshipsShippingCategory**](POSTShipments201ResponseDataRelationshipsShippingCategory.md) |  | [optional] 
**InventoryStockLocation** | Pointer to [**POSTShipments201ResponseDataRelationshipsInventoryStockLocation**](POSTShipments201ResponseDataRelationshipsInventoryStockLocation.md) |  | [optional] 
**StockLocation** | Pointer to [**POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation**](POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation.md) |  | [optional] 
**OriginAddress** | Pointer to [**POSTReturns201ResponseDataRelationshipsOriginAddress**](POSTReturns201ResponseDataRelationshipsOriginAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**POSTOrders201ResponseDataRelationshipsShippingAddress**](POSTOrders201ResponseDataRelationshipsShippingAddress.md) |  | [optional] 
**ShippingMethod** | Pointer to [**POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethod**](POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethod.md) |  | [optional] 
**DeliveryLeadTime** | Pointer to [**POSTShipments201ResponseDataRelationshipsDeliveryLeadTime**](POSTShipments201ResponseDataRelationshipsDeliveryLeadTime.md) |  | [optional] 
**StockLineItems** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockLineItems**](POSTLineItems201ResponseDataRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockTransfers**](POSTLineItems201ResponseDataRelationshipsStockTransfers.md) |  | [optional] 
**LineItems** | Pointer to [**POSTOrders201ResponseDataRelationshipsLineItems**](POSTOrders201ResponseDataRelationshipsLineItems.md) |  | [optional] 
**AvailableShippingMethods** | Pointer to [**POSTShipments201ResponseDataRelationshipsAvailableShippingMethods**](POSTShipments201ResponseDataRelationshipsAvailableShippingMethods.md) |  | [optional] 
**CarrierAccounts** | Pointer to [**POSTShipments201ResponseDataRelationshipsCarrierAccounts**](POSTShipments201ResponseDataRelationshipsCarrierAccounts.md) |  | [optional] 
**Parcels** | Pointer to [**POSTPackages201ResponseDataRelationshipsParcels**](POSTPackages201ResponseDataRelationshipsParcels.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**POSTAddresses201ResponseDataRelationshipsEvents**](POSTAddresses201ResponseDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**POSTAddresses201ResponseDataRelationshipsTags**](POSTAddresses201ResponseDataRelationshipsTags.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTShipments201ResponseDataRelationships

`func NewPOSTShipments201ResponseDataRelationships() *POSTShipments201ResponseDataRelationships`

NewPOSTShipments201ResponseDataRelationships instantiates a new POSTShipments201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTShipments201ResponseDataRelationshipsWithDefaults

`func NewPOSTShipments201ResponseDataRelationshipsWithDefaults() *POSTShipments201ResponseDataRelationships`

NewPOSTShipments201ResponseDataRelationshipsWithDefaults instantiates a new POSTShipments201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *POSTShipments201ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *POSTShipments201ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *POSTShipments201ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *POSTShipments201ResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetShippingCategory

`func (o *POSTShipments201ResponseDataRelationships) GetShippingCategory() POSTShipments201ResponseDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *POSTShipments201ResponseDataRelationships) GetShippingCategoryOk() (*POSTShipments201ResponseDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *POSTShipments201ResponseDataRelationships) SetShippingCategory(v POSTShipments201ResponseDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *POSTShipments201ResponseDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetInventoryStockLocation

`func (o *POSTShipments201ResponseDataRelationships) GetInventoryStockLocation() POSTShipments201ResponseDataRelationshipsInventoryStockLocation`

GetInventoryStockLocation returns the InventoryStockLocation field if non-nil, zero value otherwise.

### GetInventoryStockLocationOk

`func (o *POSTShipments201ResponseDataRelationships) GetInventoryStockLocationOk() (*POSTShipments201ResponseDataRelationshipsInventoryStockLocation, bool)`

GetInventoryStockLocationOk returns a tuple with the InventoryStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryStockLocation

`func (o *POSTShipments201ResponseDataRelationships) SetInventoryStockLocation(v POSTShipments201ResponseDataRelationshipsInventoryStockLocation)`

SetInventoryStockLocation sets InventoryStockLocation field to given value.

### HasInventoryStockLocation

`func (o *POSTShipments201ResponseDataRelationships) HasInventoryStockLocation() bool`

HasInventoryStockLocation returns a boolean if a field has been set.

### GetStockLocation

`func (o *POSTShipments201ResponseDataRelationships) GetStockLocation() POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *POSTShipments201ResponseDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *POSTShipments201ResponseDataRelationships) SetStockLocation(v POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *POSTShipments201ResponseDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOriginAddress

`func (o *POSTShipments201ResponseDataRelationships) GetOriginAddress() POSTReturns201ResponseDataRelationshipsOriginAddress`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *POSTShipments201ResponseDataRelationships) GetOriginAddressOk() (*POSTReturns201ResponseDataRelationshipsOriginAddress, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *POSTShipments201ResponseDataRelationships) SetOriginAddress(v POSTReturns201ResponseDataRelationshipsOriginAddress)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *POSTShipments201ResponseDataRelationships) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *POSTShipments201ResponseDataRelationships) GetShippingAddress() POSTOrders201ResponseDataRelationshipsShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *POSTShipments201ResponseDataRelationships) GetShippingAddressOk() (*POSTOrders201ResponseDataRelationshipsShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *POSTShipments201ResponseDataRelationships) SetShippingAddress(v POSTOrders201ResponseDataRelationshipsShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *POSTShipments201ResponseDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetShippingMethod

`func (o *POSTShipments201ResponseDataRelationships) GetShippingMethod() POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *POSTShipments201ResponseDataRelationships) GetShippingMethodOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *POSTShipments201ResponseDataRelationships) SetShippingMethod(v POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *POSTShipments201ResponseDataRelationships) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetDeliveryLeadTime

`func (o *POSTShipments201ResponseDataRelationships) GetDeliveryLeadTime() POSTShipments201ResponseDataRelationshipsDeliveryLeadTime`

GetDeliveryLeadTime returns the DeliveryLeadTime field if non-nil, zero value otherwise.

### GetDeliveryLeadTimeOk

`func (o *POSTShipments201ResponseDataRelationships) GetDeliveryLeadTimeOk() (*POSTShipments201ResponseDataRelationshipsDeliveryLeadTime, bool)`

GetDeliveryLeadTimeOk returns a tuple with the DeliveryLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTime

`func (o *POSTShipments201ResponseDataRelationships) SetDeliveryLeadTime(v POSTShipments201ResponseDataRelationshipsDeliveryLeadTime)`

SetDeliveryLeadTime sets DeliveryLeadTime field to given value.

### HasDeliveryLeadTime

`func (o *POSTShipments201ResponseDataRelationships) HasDeliveryLeadTime() bool`

HasDeliveryLeadTime returns a boolean if a field has been set.

### GetStockLineItems

`func (o *POSTShipments201ResponseDataRelationships) GetStockLineItems() POSTLineItems201ResponseDataRelationshipsStockLineItems`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *POSTShipments201ResponseDataRelationships) GetStockLineItemsOk() (*POSTLineItems201ResponseDataRelationshipsStockLineItems, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *POSTShipments201ResponseDataRelationships) SetStockLineItems(v POSTLineItems201ResponseDataRelationshipsStockLineItems)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *POSTShipments201ResponseDataRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *POSTShipments201ResponseDataRelationships) GetStockTransfers() POSTLineItems201ResponseDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *POSTShipments201ResponseDataRelationships) GetStockTransfersOk() (*POSTLineItems201ResponseDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *POSTShipments201ResponseDataRelationships) SetStockTransfers(v POSTLineItems201ResponseDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *POSTShipments201ResponseDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetLineItems

`func (o *POSTShipments201ResponseDataRelationships) GetLineItems() POSTOrders201ResponseDataRelationshipsLineItems`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *POSTShipments201ResponseDataRelationships) GetLineItemsOk() (*POSTOrders201ResponseDataRelationshipsLineItems, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *POSTShipments201ResponseDataRelationships) SetLineItems(v POSTOrders201ResponseDataRelationshipsLineItems)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *POSTShipments201ResponseDataRelationships) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetAvailableShippingMethods

`func (o *POSTShipments201ResponseDataRelationships) GetAvailableShippingMethods() POSTShipments201ResponseDataRelationshipsAvailableShippingMethods`

GetAvailableShippingMethods returns the AvailableShippingMethods field if non-nil, zero value otherwise.

### GetAvailableShippingMethodsOk

`func (o *POSTShipments201ResponseDataRelationships) GetAvailableShippingMethodsOk() (*POSTShipments201ResponseDataRelationshipsAvailableShippingMethods, bool)`

GetAvailableShippingMethodsOk returns a tuple with the AvailableShippingMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShippingMethods

`func (o *POSTShipments201ResponseDataRelationships) SetAvailableShippingMethods(v POSTShipments201ResponseDataRelationshipsAvailableShippingMethods)`

SetAvailableShippingMethods sets AvailableShippingMethods field to given value.

### HasAvailableShippingMethods

`func (o *POSTShipments201ResponseDataRelationships) HasAvailableShippingMethods() bool`

HasAvailableShippingMethods returns a boolean if a field has been set.

### GetCarrierAccounts

`func (o *POSTShipments201ResponseDataRelationships) GetCarrierAccounts() POSTShipments201ResponseDataRelationshipsCarrierAccounts`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *POSTShipments201ResponseDataRelationships) GetCarrierAccountsOk() (*POSTShipments201ResponseDataRelationshipsCarrierAccounts, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *POSTShipments201ResponseDataRelationships) SetCarrierAccounts(v POSTShipments201ResponseDataRelationshipsCarrierAccounts)`

SetCarrierAccounts sets CarrierAccounts field to given value.

### HasCarrierAccounts

`func (o *POSTShipments201ResponseDataRelationships) HasCarrierAccounts() bool`

HasCarrierAccounts returns a boolean if a field has been set.

### GetParcels

`func (o *POSTShipments201ResponseDataRelationships) GetParcels() POSTPackages201ResponseDataRelationshipsParcels`

GetParcels returns the Parcels field if non-nil, zero value otherwise.

### GetParcelsOk

`func (o *POSTShipments201ResponseDataRelationships) GetParcelsOk() (*POSTPackages201ResponseDataRelationshipsParcels, bool)`

GetParcelsOk returns a tuple with the Parcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcels

`func (o *POSTShipments201ResponseDataRelationships) SetParcels(v POSTPackages201ResponseDataRelationshipsParcels)`

SetParcels sets Parcels field to given value.

### HasParcels

`func (o *POSTShipments201ResponseDataRelationships) HasParcels() bool`

HasParcels returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTShipments201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTShipments201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTShipments201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTShipments201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *POSTShipments201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTShipments201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTShipments201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTShipments201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *POSTShipments201ResponseDataRelationships) GetTags() POSTAddresses201ResponseDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *POSTShipments201ResponseDataRelationships) GetTagsOk() (*POSTAddresses201ResponseDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *POSTShipments201ResponseDataRelationships) SetTags(v POSTAddresses201ResponseDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *POSTShipments201ResponseDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *POSTShipments201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTShipments201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTShipments201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTShipments201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


