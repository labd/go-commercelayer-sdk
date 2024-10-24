# ShipmentDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**ShippingCategory** | Pointer to [**ShipmentDataRelationshipsShippingCategory**](ShipmentDataRelationshipsShippingCategory.md) |  | [optional] 
**InventoryStockLocation** | Pointer to [**InventoryModelDataRelationshipsInventoryStockLocations**](InventoryModelDataRelationshipsInventoryStockLocations.md) |  | [optional] 
**StockLocation** | Pointer to [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | [optional] 
**OriginAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**ShippingAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**ShippingMethod** | Pointer to [**DeliveryLeadTimeDataRelationshipsShippingMethod**](DeliveryLeadTimeDataRelationshipsShippingMethod.md) |  | [optional] 
**DeliveryLeadTime** | Pointer to [**ShipmentDataRelationshipsDeliveryLeadTime**](ShipmentDataRelationshipsDeliveryLeadTime.md) |  | [optional] 
**StockLineItems** | Pointer to [**LineItemDataRelationshipsStockLineItems**](LineItemDataRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**LineItemDataRelationshipsStockTransfers**](LineItemDataRelationshipsStockTransfers.md) |  | [optional] 
**LineItems** | Pointer to [**LineItemOptionDataRelationshipsLineItem**](LineItemOptionDataRelationshipsLineItem.md) |  | [optional] 
**AvailableShippingMethods** | Pointer to [**DeliveryLeadTimeDataRelationshipsShippingMethod**](DeliveryLeadTimeDataRelationshipsShippingMethod.md) |  | [optional] 
**CarrierAccounts** | Pointer to [**ShipmentDataRelationshipsCarrierAccounts**](ShipmentDataRelationshipsCarrierAccounts.md) |  | [optional] 
**Parcels** | Pointer to [**PackageDataRelationshipsParcels**](PackageDataRelationshipsParcels.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**AddressDataRelationshipsTags**](AddressDataRelationshipsTags.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewShipmentDataRelationships

`func NewShipmentDataRelationships() *ShipmentDataRelationships`

NewShipmentDataRelationships instantiates a new ShipmentDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentDataRelationshipsWithDefaults

`func NewShipmentDataRelationshipsWithDefaults() *ShipmentDataRelationships`

NewShipmentDataRelationshipsWithDefaults instantiates a new ShipmentDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *ShipmentDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ShipmentDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ShipmentDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ShipmentDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetShippingCategory

`func (o *ShipmentDataRelationships) GetShippingCategory() ShipmentDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *ShipmentDataRelationships) GetShippingCategoryOk() (*ShipmentDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *ShipmentDataRelationships) SetShippingCategory(v ShipmentDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *ShipmentDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetInventoryStockLocation

`func (o *ShipmentDataRelationships) GetInventoryStockLocation() InventoryModelDataRelationshipsInventoryStockLocations`

GetInventoryStockLocation returns the InventoryStockLocation field if non-nil, zero value otherwise.

### GetInventoryStockLocationOk

`func (o *ShipmentDataRelationships) GetInventoryStockLocationOk() (*InventoryModelDataRelationshipsInventoryStockLocations, bool)`

GetInventoryStockLocationOk returns a tuple with the InventoryStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryStockLocation

`func (o *ShipmentDataRelationships) SetInventoryStockLocation(v InventoryModelDataRelationshipsInventoryStockLocations)`

SetInventoryStockLocation sets InventoryStockLocation field to given value.

### HasInventoryStockLocation

`func (o *ShipmentDataRelationships) HasInventoryStockLocation() bool`

HasInventoryStockLocation returns a boolean if a field has been set.

### GetStockLocation

`func (o *ShipmentDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *ShipmentDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *ShipmentDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *ShipmentDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOriginAddress

`func (o *ShipmentDataRelationships) GetOriginAddress() BingGeocoderDataRelationshipsAddresses`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *ShipmentDataRelationships) GetOriginAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *ShipmentDataRelationships) SetOriginAddress(v BingGeocoderDataRelationshipsAddresses)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *ShipmentDataRelationships) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *ShipmentDataRelationships) GetShippingAddress() BingGeocoderDataRelationshipsAddresses`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *ShipmentDataRelationships) GetShippingAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *ShipmentDataRelationships) SetShippingAddress(v BingGeocoderDataRelationshipsAddresses)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *ShipmentDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetShippingMethod

`func (o *ShipmentDataRelationships) GetShippingMethod() DeliveryLeadTimeDataRelationshipsShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *ShipmentDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeDataRelationshipsShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *ShipmentDataRelationships) SetShippingMethod(v DeliveryLeadTimeDataRelationshipsShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *ShipmentDataRelationships) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetDeliveryLeadTime

`func (o *ShipmentDataRelationships) GetDeliveryLeadTime() ShipmentDataRelationshipsDeliveryLeadTime`

GetDeliveryLeadTime returns the DeliveryLeadTime field if non-nil, zero value otherwise.

### GetDeliveryLeadTimeOk

`func (o *ShipmentDataRelationships) GetDeliveryLeadTimeOk() (*ShipmentDataRelationshipsDeliveryLeadTime, bool)`

GetDeliveryLeadTimeOk returns a tuple with the DeliveryLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTime

`func (o *ShipmentDataRelationships) SetDeliveryLeadTime(v ShipmentDataRelationshipsDeliveryLeadTime)`

SetDeliveryLeadTime sets DeliveryLeadTime field to given value.

### HasDeliveryLeadTime

`func (o *ShipmentDataRelationships) HasDeliveryLeadTime() bool`

HasDeliveryLeadTime returns a boolean if a field has been set.

### GetStockLineItems

`func (o *ShipmentDataRelationships) GetStockLineItems() LineItemDataRelationshipsStockLineItems`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *ShipmentDataRelationships) GetStockLineItemsOk() (*LineItemDataRelationshipsStockLineItems, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *ShipmentDataRelationships) SetStockLineItems(v LineItemDataRelationshipsStockLineItems)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *ShipmentDataRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *ShipmentDataRelationships) GetStockTransfers() LineItemDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *ShipmentDataRelationships) GetStockTransfersOk() (*LineItemDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *ShipmentDataRelationships) SetStockTransfers(v LineItemDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *ShipmentDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetLineItems

`func (o *ShipmentDataRelationships) GetLineItems() LineItemOptionDataRelationshipsLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *ShipmentDataRelationships) GetLineItemsOk() (*LineItemOptionDataRelationshipsLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *ShipmentDataRelationships) SetLineItems(v LineItemOptionDataRelationshipsLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *ShipmentDataRelationships) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetAvailableShippingMethods

`func (o *ShipmentDataRelationships) GetAvailableShippingMethods() DeliveryLeadTimeDataRelationshipsShippingMethod`

GetAvailableShippingMethods returns the AvailableShippingMethods field if non-nil, zero value otherwise.

### GetAvailableShippingMethodsOk

`func (o *ShipmentDataRelationships) GetAvailableShippingMethodsOk() (*DeliveryLeadTimeDataRelationshipsShippingMethod, bool)`

GetAvailableShippingMethodsOk returns a tuple with the AvailableShippingMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShippingMethods

`func (o *ShipmentDataRelationships) SetAvailableShippingMethods(v DeliveryLeadTimeDataRelationshipsShippingMethod)`

SetAvailableShippingMethods sets AvailableShippingMethods field to given value.

### HasAvailableShippingMethods

`func (o *ShipmentDataRelationships) HasAvailableShippingMethods() bool`

HasAvailableShippingMethods returns a boolean if a field has been set.

### GetCarrierAccounts

`func (o *ShipmentDataRelationships) GetCarrierAccounts() ShipmentDataRelationshipsCarrierAccounts`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *ShipmentDataRelationships) GetCarrierAccountsOk() (*ShipmentDataRelationshipsCarrierAccounts, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *ShipmentDataRelationships) SetCarrierAccounts(v ShipmentDataRelationshipsCarrierAccounts)`

SetCarrierAccounts sets CarrierAccounts field to given value.

### HasCarrierAccounts

`func (o *ShipmentDataRelationships) HasCarrierAccounts() bool`

HasCarrierAccounts returns a boolean if a field has been set.

### GetParcels

`func (o *ShipmentDataRelationships) GetParcels() PackageDataRelationshipsParcels`

GetParcels returns the Parcels field if non-nil, zero value otherwise.

### GetParcelsOk

`func (o *ShipmentDataRelationships) GetParcelsOk() (*PackageDataRelationshipsParcels, bool)`

GetParcelsOk returns a tuple with the Parcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcels

`func (o *ShipmentDataRelationships) SetParcels(v PackageDataRelationshipsParcels)`

SetParcels sets Parcels field to given value.

### HasParcels

`func (o *ShipmentDataRelationships) HasParcels() bool`

HasParcels returns a boolean if a field has been set.

### GetAttachments

`func (o *ShipmentDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ShipmentDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ShipmentDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ShipmentDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *ShipmentDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ShipmentDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ShipmentDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ShipmentDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *ShipmentDataRelationships) GetTags() AddressDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ShipmentDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ShipmentDataRelationships) SetTags(v AddressDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ShipmentDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *ShipmentDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ShipmentDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ShipmentDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ShipmentDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


