# POSTReturns201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**POSTAdyenPayments201ResponseDataRelationshipsOrder**](POSTAdyenPayments201ResponseDataRelationshipsOrder.md) |  | [optional] 
**Customer** | Pointer to [**POSTCouponRecipients201ResponseDataRelationshipsCustomer**](POSTCouponRecipients201ResponseDataRelationshipsCustomer.md) |  | [optional] 
**StockLocation** | Pointer to [**POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation**](POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation.md) |  | [optional] 
**OriginAddress** | Pointer to [**POSTReturns201ResponseDataRelationshipsOriginAddress**](POSTReturns201ResponseDataRelationshipsOriginAddress.md) |  | [optional] 
**DestinationAddress** | Pointer to [**POSTReturns201ResponseDataRelationshipsDestinationAddress**](POSTReturns201ResponseDataRelationshipsDestinationAddress.md) |  | [optional] 
**ReturnLineItems** | Pointer to [**POSTReturns201ResponseDataRelationshipsReturnLineItems**](POSTReturns201ResponseDataRelationshipsReturnLineItems.md) |  | [optional] 
**Attachments** | Pointer to [**POSTAvalaraAccounts201ResponseDataRelationshipsAttachments**](POSTAvalaraAccounts201ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewPOSTReturns201ResponseDataRelationships

`func NewPOSTReturns201ResponseDataRelationships() *POSTReturns201ResponseDataRelationships`

NewPOSTReturns201ResponseDataRelationships instantiates a new POSTReturns201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTReturns201ResponseDataRelationshipsWithDefaults

`func NewPOSTReturns201ResponseDataRelationshipsWithDefaults() *POSTReturns201ResponseDataRelationships`

NewPOSTReturns201ResponseDataRelationshipsWithDefaults instantiates a new POSTReturns201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *POSTReturns201ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *POSTReturns201ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *POSTReturns201ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *POSTReturns201ResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCustomer

`func (o *POSTReturns201ResponseDataRelationships) GetCustomer() POSTCouponRecipients201ResponseDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *POSTReturns201ResponseDataRelationships) GetCustomerOk() (*POSTCouponRecipients201ResponseDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *POSTReturns201ResponseDataRelationships) SetCustomer(v POSTCouponRecipients201ResponseDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *POSTReturns201ResponseDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetStockLocation

`func (o *POSTReturns201ResponseDataRelationships) GetStockLocation() POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *POSTReturns201ResponseDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *POSTReturns201ResponseDataRelationships) SetStockLocation(v POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *POSTReturns201ResponseDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOriginAddress

`func (o *POSTReturns201ResponseDataRelationships) GetOriginAddress() POSTReturns201ResponseDataRelationshipsOriginAddress`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *POSTReturns201ResponseDataRelationships) GetOriginAddressOk() (*POSTReturns201ResponseDataRelationshipsOriginAddress, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *POSTReturns201ResponseDataRelationships) SetOriginAddress(v POSTReturns201ResponseDataRelationshipsOriginAddress)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *POSTReturns201ResponseDataRelationships) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *POSTReturns201ResponseDataRelationships) GetDestinationAddress() POSTReturns201ResponseDataRelationshipsDestinationAddress`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *POSTReturns201ResponseDataRelationships) GetDestinationAddressOk() (*POSTReturns201ResponseDataRelationshipsDestinationAddress, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *POSTReturns201ResponseDataRelationships) SetDestinationAddress(v POSTReturns201ResponseDataRelationshipsDestinationAddress)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *POSTReturns201ResponseDataRelationships) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetReturnLineItems

`func (o *POSTReturns201ResponseDataRelationships) GetReturnLineItems() POSTReturns201ResponseDataRelationshipsReturnLineItems`

GetReturnLineItems returns the ReturnLineItems field if non-nil, zero value otherwise.

### GetReturnLineItemsOk

`func (o *POSTReturns201ResponseDataRelationships) GetReturnLineItemsOk() (*POSTReturns201ResponseDataRelationshipsReturnLineItems, bool)`

GetReturnLineItemsOk returns a tuple with the ReturnLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLineItems

`func (o *POSTReturns201ResponseDataRelationships) SetReturnLineItems(v POSTReturns201ResponseDataRelationshipsReturnLineItems)`

SetReturnLineItems sets ReturnLineItems field to given value.

### HasReturnLineItems

`func (o *POSTReturns201ResponseDataRelationships) HasReturnLineItems() bool`

HasReturnLineItems returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTReturns201ResponseDataRelationships) GetAttachments() POSTAvalaraAccounts201ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTReturns201ResponseDataRelationships) GetAttachmentsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTReturns201ResponseDataRelationships) SetAttachments(v POSTAvalaraAccounts201ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTReturns201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *POSTReturns201ResponseDataRelationships) GetEvents() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTReturns201ResponseDataRelationships) GetEventsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTReturns201ResponseDataRelationships) SetEvents(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTReturns201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


