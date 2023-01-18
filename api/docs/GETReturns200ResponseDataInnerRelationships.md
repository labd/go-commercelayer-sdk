# GETReturns200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**GETAdyenPayments200ResponseDataInnerRelationshipsOrder**](GETAdyenPayments200ResponseDataInnerRelationshipsOrder.md) |  | [optional] 
**Customer** | Pointer to [**GETCouponRecipients200ResponseDataInnerRelationshipsCustomer**](GETCouponRecipients200ResponseDataInnerRelationshipsCustomer.md) |  | [optional] 
**StockLocation** | Pointer to [**GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation**](GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation.md) |  | [optional] 
**OriginAddress** | Pointer to [**GETReturns200ResponseDataInnerRelationshipsOriginAddress**](GETReturns200ResponseDataInnerRelationshipsOriginAddress.md) |  | [optional] 
**DestinationAddress** | Pointer to [**GETReturns200ResponseDataInnerRelationshipsDestinationAddress**](GETReturns200ResponseDataInnerRelationshipsDestinationAddress.md) |  | [optional] 
**ReturnLineItems** | Pointer to [**GETReturns200ResponseDataInnerRelationshipsReturnLineItems**](GETReturns200ResponseDataInnerRelationshipsReturnLineItems.md) |  | [optional] 
**Attachments** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments**](GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**GETCleanups200ResponseDataInnerRelationshipsEvents**](GETCleanups200ResponseDataInnerRelationshipsEvents.md) |  | [optional] 

## Methods

### NewGETReturns200ResponseDataInnerRelationships

`func NewGETReturns200ResponseDataInnerRelationships() *GETReturns200ResponseDataInnerRelationships`

NewGETReturns200ResponseDataInnerRelationships instantiates a new GETReturns200ResponseDataInnerRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETReturns200ResponseDataInnerRelationshipsWithDefaults

`func NewGETReturns200ResponseDataInnerRelationshipsWithDefaults() *GETReturns200ResponseDataInnerRelationships`

NewGETReturns200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETReturns200ResponseDataInnerRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *GETReturns200ResponseDataInnerRelationships) GetOrder() GETAdyenPayments200ResponseDataInnerRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetOrderOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GETReturns200ResponseDataInnerRelationships) SetOrder(v GETAdyenPayments200ResponseDataInnerRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GETReturns200ResponseDataInnerRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCustomer

`func (o *GETReturns200ResponseDataInnerRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GETReturns200ResponseDataInnerRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *GETReturns200ResponseDataInnerRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetStockLocation

`func (o *GETReturns200ResponseDataInnerRelationships) GetStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *GETReturns200ResponseDataInnerRelationships) SetStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *GETReturns200ResponseDataInnerRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOriginAddress

`func (o *GETReturns200ResponseDataInnerRelationships) GetOriginAddress() GETReturns200ResponseDataInnerRelationshipsOriginAddress`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetOriginAddressOk() (*GETReturns200ResponseDataInnerRelationshipsOriginAddress, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *GETReturns200ResponseDataInnerRelationships) SetOriginAddress(v GETReturns200ResponseDataInnerRelationshipsOriginAddress)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *GETReturns200ResponseDataInnerRelationships) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *GETReturns200ResponseDataInnerRelationships) GetDestinationAddress() GETReturns200ResponseDataInnerRelationshipsDestinationAddress`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetDestinationAddressOk() (*GETReturns200ResponseDataInnerRelationshipsDestinationAddress, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *GETReturns200ResponseDataInnerRelationships) SetDestinationAddress(v GETReturns200ResponseDataInnerRelationshipsDestinationAddress)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *GETReturns200ResponseDataInnerRelationships) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetReturnLineItems

`func (o *GETReturns200ResponseDataInnerRelationships) GetReturnLineItems() GETReturns200ResponseDataInnerRelationshipsReturnLineItems`

GetReturnLineItems returns the ReturnLineItems field if non-nil, zero value otherwise.

### GetReturnLineItemsOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetReturnLineItemsOk() (*GETReturns200ResponseDataInnerRelationshipsReturnLineItems, bool)`

GetReturnLineItemsOk returns a tuple with the ReturnLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLineItems

`func (o *GETReturns200ResponseDataInnerRelationships) SetReturnLineItems(v GETReturns200ResponseDataInnerRelationshipsReturnLineItems)`

SetReturnLineItems sets ReturnLineItems field to given value.

### HasReturnLineItems

`func (o *GETReturns200ResponseDataInnerRelationships) HasReturnLineItems() bool`

HasReturnLineItems returns a boolean if a field has been set.

### GetAttachments

`func (o *GETReturns200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETReturns200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETReturns200ResponseDataInnerRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GETReturns200ResponseDataInnerRelationships) GetEvents() GETCleanups200ResponseDataInnerRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetEventsOk() (*GETCleanups200ResponseDataInnerRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETReturns200ResponseDataInnerRelationships) SetEvents(v GETCleanups200ResponseDataInnerRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETReturns200ResponseDataInnerRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


