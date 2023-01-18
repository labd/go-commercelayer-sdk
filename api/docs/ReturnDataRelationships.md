# ReturnDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**StockLocation** | Pointer to [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | [optional] 
**OriginAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**DestinationAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**ReturnLineItems** | Pointer to [**ReturnDataRelationshipsReturnLineItems**](ReturnDataRelationshipsReturnLineItems.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountDataRelationshipsAttachments**](AvalaraAccountDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**CleanupDataRelationshipsEvents**](CleanupDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewReturnDataRelationships

`func NewReturnDataRelationships() *ReturnDataRelationships`

NewReturnDataRelationships instantiates a new ReturnDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnDataRelationshipsWithDefaults

`func NewReturnDataRelationshipsWithDefaults() *ReturnDataRelationships`

NewReturnDataRelationshipsWithDefaults instantiates a new ReturnDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *ReturnDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ReturnDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ReturnDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ReturnDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCustomer

`func (o *ReturnDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *ReturnDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *ReturnDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *ReturnDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetStockLocation

`func (o *ReturnDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *ReturnDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *ReturnDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *ReturnDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOriginAddress

`func (o *ReturnDataRelationships) GetOriginAddress() BingGeocoderDataRelationshipsAddresses`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *ReturnDataRelationships) GetOriginAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *ReturnDataRelationships) SetOriginAddress(v BingGeocoderDataRelationshipsAddresses)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *ReturnDataRelationships) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *ReturnDataRelationships) GetDestinationAddress() BingGeocoderDataRelationshipsAddresses`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *ReturnDataRelationships) GetDestinationAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *ReturnDataRelationships) SetDestinationAddress(v BingGeocoderDataRelationshipsAddresses)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *ReturnDataRelationships) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetReturnLineItems

`func (o *ReturnDataRelationships) GetReturnLineItems() ReturnDataRelationshipsReturnLineItems`

GetReturnLineItems returns the ReturnLineItems field if non-nil, zero value otherwise.

### GetReturnLineItemsOk

`func (o *ReturnDataRelationships) GetReturnLineItemsOk() (*ReturnDataRelationshipsReturnLineItems, bool)`

GetReturnLineItemsOk returns a tuple with the ReturnLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLineItems

`func (o *ReturnDataRelationships) SetReturnLineItems(v ReturnDataRelationshipsReturnLineItems)`

SetReturnLineItems sets ReturnLineItems field to given value.

### HasReturnLineItems

`func (o *ReturnDataRelationships) HasReturnLineItems() bool`

HasReturnLineItems returns a boolean if a field has been set.

### GetAttachments

`func (o *ReturnDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ReturnDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ReturnDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ReturnDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *ReturnDataRelationships) GetEvents() CleanupDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ReturnDataRelationships) GetEventsOk() (*CleanupDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ReturnDataRelationships) SetEvents(v CleanupDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ReturnDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


