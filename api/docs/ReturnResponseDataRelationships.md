# ReturnResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentResponseDataRelationshipsOrder**](AdyenPaymentResponseDataRelationshipsOrder.md) |  | [optional] 
**Customer** | Pointer to [**CouponRecipientResponseDataRelationshipsCustomer**](CouponRecipientResponseDataRelationshipsCustomer.md) |  | [optional] 
**StockLocation** | Pointer to [**DeliveryLeadTimeResponseDataRelationshipsStockLocation**](DeliveryLeadTimeResponseDataRelationshipsStockLocation.md) |  | [optional] 
**OriginAddress** | Pointer to [**ReturnResponseDataRelationshipsOriginAddress**](ReturnResponseDataRelationshipsOriginAddress.md) |  | [optional] 
**DestinationAddress** | Pointer to [**ReturnResponseDataRelationshipsDestinationAddress**](ReturnResponseDataRelationshipsDestinationAddress.md) |  | [optional] 
**ReturnLineItems** | Pointer to [**ReturnResponseDataRelationshipsReturnLineItems**](ReturnResponseDataRelationshipsReturnLineItems.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressResponseDataRelationshipsEvents**](CustomerAddressResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewReturnResponseDataRelationships

`func NewReturnResponseDataRelationships() *ReturnResponseDataRelationships`

NewReturnResponseDataRelationships instantiates a new ReturnResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnResponseDataRelationshipsWithDefaults

`func NewReturnResponseDataRelationshipsWithDefaults() *ReturnResponseDataRelationships`

NewReturnResponseDataRelationshipsWithDefaults instantiates a new ReturnResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *ReturnResponseDataRelationships) GetOrder() AdyenPaymentResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ReturnResponseDataRelationships) GetOrderOk() (*AdyenPaymentResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ReturnResponseDataRelationships) SetOrder(v AdyenPaymentResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ReturnResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCustomer

`func (o *ReturnResponseDataRelationships) GetCustomer() CouponRecipientResponseDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *ReturnResponseDataRelationships) GetCustomerOk() (*CouponRecipientResponseDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *ReturnResponseDataRelationships) SetCustomer(v CouponRecipientResponseDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *ReturnResponseDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetStockLocation

`func (o *ReturnResponseDataRelationships) GetStockLocation() DeliveryLeadTimeResponseDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *ReturnResponseDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeResponseDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *ReturnResponseDataRelationships) SetStockLocation(v DeliveryLeadTimeResponseDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *ReturnResponseDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOriginAddress

`func (o *ReturnResponseDataRelationships) GetOriginAddress() ReturnResponseDataRelationshipsOriginAddress`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *ReturnResponseDataRelationships) GetOriginAddressOk() (*ReturnResponseDataRelationshipsOriginAddress, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *ReturnResponseDataRelationships) SetOriginAddress(v ReturnResponseDataRelationshipsOriginAddress)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *ReturnResponseDataRelationships) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *ReturnResponseDataRelationships) GetDestinationAddress() ReturnResponseDataRelationshipsDestinationAddress`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *ReturnResponseDataRelationships) GetDestinationAddressOk() (*ReturnResponseDataRelationshipsDestinationAddress, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *ReturnResponseDataRelationships) SetDestinationAddress(v ReturnResponseDataRelationshipsDestinationAddress)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *ReturnResponseDataRelationships) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetReturnLineItems

`func (o *ReturnResponseDataRelationships) GetReturnLineItems() ReturnResponseDataRelationshipsReturnLineItems`

GetReturnLineItems returns the ReturnLineItems field if non-nil, zero value otherwise.

### GetReturnLineItemsOk

`func (o *ReturnResponseDataRelationships) GetReturnLineItemsOk() (*ReturnResponseDataRelationshipsReturnLineItems, bool)`

GetReturnLineItemsOk returns a tuple with the ReturnLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLineItems

`func (o *ReturnResponseDataRelationships) SetReturnLineItems(v ReturnResponseDataRelationshipsReturnLineItems)`

SetReturnLineItems sets ReturnLineItems field to given value.

### HasReturnLineItems

`func (o *ReturnResponseDataRelationships) HasReturnLineItems() bool`

HasReturnLineItems returns a boolean if a field has been set.

### GetAttachments

`func (o *ReturnResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ReturnResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ReturnResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ReturnResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *ReturnResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ReturnResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ReturnResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ReturnResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


