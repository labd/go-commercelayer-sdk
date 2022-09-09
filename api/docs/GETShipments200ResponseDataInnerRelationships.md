# GETShipments200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**ShippingCategory** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**StockLocation** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**OriginAddress** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**ShippingAddress** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**ShippingMethod** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**DeliveryLeadTime** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**ShipmentLineItems** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**StockLineItems** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**StockTransfers** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**AvailableShippingMethods** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**CarrierAccounts** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Parcels** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Attachments** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Events** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 

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

`func (o *GETShipments200ResponseDataInnerRelationships) GetOrder() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetOrderOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GETShipments200ResponseDataInnerRelationships) SetOrder(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GETShipments200ResponseDataInnerRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetShippingCategory

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingCategory() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingCategoryOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *GETShipments200ResponseDataInnerRelationships) SetShippingCategory(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *GETShipments200ResponseDataInnerRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetStockLocation

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockLocation() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockLocationOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *GETShipments200ResponseDataInnerRelationships) SetStockLocation(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *GETShipments200ResponseDataInnerRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOriginAddress

`func (o *GETShipments200ResponseDataInnerRelationships) GetOriginAddress() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetOriginAddressOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *GETShipments200ResponseDataInnerRelationships) SetOriginAddress(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *GETShipments200ResponseDataInnerRelationships) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingAddress() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingAddressOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *GETShipments200ResponseDataInnerRelationships) SetShippingAddress(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *GETShipments200ResponseDataInnerRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetShippingMethod

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingMethod() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetShippingMethodOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *GETShipments200ResponseDataInnerRelationships) SetShippingMethod(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *GETShipments200ResponseDataInnerRelationships) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetDeliveryLeadTime

`func (o *GETShipments200ResponseDataInnerRelationships) GetDeliveryLeadTime() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetDeliveryLeadTime returns the DeliveryLeadTime field if non-nil, zero value otherwise.

### GetDeliveryLeadTimeOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetDeliveryLeadTimeOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetDeliveryLeadTimeOk returns a tuple with the DeliveryLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTime

`func (o *GETShipments200ResponseDataInnerRelationships) SetDeliveryLeadTime(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetDeliveryLeadTime sets DeliveryLeadTime field to given value.

### HasDeliveryLeadTime

`func (o *GETShipments200ResponseDataInnerRelationships) HasDeliveryLeadTime() bool`

HasDeliveryLeadTime returns a boolean if a field has been set.

### GetShipmentLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) GetShipmentLineItems() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetShipmentLineItems returns the ShipmentLineItems field if non-nil, zero value otherwise.

### GetShipmentLineItemsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetShipmentLineItemsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetShipmentLineItemsOk returns a tuple with the ShipmentLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) SetShipmentLineItems(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetShipmentLineItems sets ShipmentLineItems field to given value.

### HasShipmentLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) HasShipmentLineItems() bool`

HasShipmentLineItems returns a boolean if a field has been set.

### GetStockLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockLineItems() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockLineItemsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) SetStockLineItems(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *GETShipments200ResponseDataInnerRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockTransfers() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetStockTransfersOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *GETShipments200ResponseDataInnerRelationships) SetStockTransfers(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *GETShipments200ResponseDataInnerRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetAvailableShippingMethods

`func (o *GETShipments200ResponseDataInnerRelationships) GetAvailableShippingMethods() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetAvailableShippingMethods returns the AvailableShippingMethods field if non-nil, zero value otherwise.

### GetAvailableShippingMethodsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetAvailableShippingMethodsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetAvailableShippingMethodsOk returns a tuple with the AvailableShippingMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShippingMethods

`func (o *GETShipments200ResponseDataInnerRelationships) SetAvailableShippingMethods(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetAvailableShippingMethods sets AvailableShippingMethods field to given value.

### HasAvailableShippingMethods

`func (o *GETShipments200ResponseDataInnerRelationships) HasAvailableShippingMethods() bool`

HasAvailableShippingMethods returns a boolean if a field has been set.

### GetCarrierAccounts

`func (o *GETShipments200ResponseDataInnerRelationships) GetCarrierAccounts() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetCarrierAccountsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *GETShipments200ResponseDataInnerRelationships) SetCarrierAccounts(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetCarrierAccounts sets CarrierAccounts field to given value.

### HasCarrierAccounts

`func (o *GETShipments200ResponseDataInnerRelationships) HasCarrierAccounts() bool`

HasCarrierAccounts returns a boolean if a field has been set.

### GetParcels

`func (o *GETShipments200ResponseDataInnerRelationships) GetParcels() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetParcels returns the Parcels field if non-nil, zero value otherwise.

### GetParcelsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetParcelsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetParcelsOk returns a tuple with the Parcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcels

`func (o *GETShipments200ResponseDataInnerRelationships) SetParcels(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetParcels sets Parcels field to given value.

### HasParcels

`func (o *GETShipments200ResponseDataInnerRelationships) HasParcels() bool`

HasParcels returns a boolean if a field has been set.

### GetAttachments

`func (o *GETShipments200ResponseDataInnerRelationships) GetAttachments() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETShipments200ResponseDataInnerRelationships) SetAttachments(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETShipments200ResponseDataInnerRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GETShipments200ResponseDataInnerRelationships) GetEvents() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETShipments200ResponseDataInnerRelationships) GetEventsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETShipments200ResponseDataInnerRelationships) SetEvents(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETShipments200ResponseDataInnerRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


