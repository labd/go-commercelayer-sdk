# GETReturns200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**Customer** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**StockLocation** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**OriginAddress** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**DestinationAddress** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**ReturnLineItems** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Attachments** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Events** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 

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

`func (o *GETReturns200ResponseDataInnerRelationships) GetOrder() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetOrderOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GETReturns200ResponseDataInnerRelationships) SetOrder(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GETReturns200ResponseDataInnerRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCustomer

`func (o *GETReturns200ResponseDataInnerRelationships) GetCustomer() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetCustomerOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GETReturns200ResponseDataInnerRelationships) SetCustomer(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *GETReturns200ResponseDataInnerRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetStockLocation

`func (o *GETReturns200ResponseDataInnerRelationships) GetStockLocation() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetStockLocationOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *GETReturns200ResponseDataInnerRelationships) SetStockLocation(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *GETReturns200ResponseDataInnerRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOriginAddress

`func (o *GETReturns200ResponseDataInnerRelationships) GetOriginAddress() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetOriginAddressOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *GETReturns200ResponseDataInnerRelationships) SetOriginAddress(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *GETReturns200ResponseDataInnerRelationships) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *GETReturns200ResponseDataInnerRelationships) GetDestinationAddress() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetDestinationAddressOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *GETReturns200ResponseDataInnerRelationships) SetDestinationAddress(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *GETReturns200ResponseDataInnerRelationships) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetReturnLineItems

`func (o *GETReturns200ResponseDataInnerRelationships) GetReturnLineItems() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetReturnLineItems returns the ReturnLineItems field if non-nil, zero value otherwise.

### GetReturnLineItemsOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetReturnLineItemsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetReturnLineItemsOk returns a tuple with the ReturnLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLineItems

`func (o *GETReturns200ResponseDataInnerRelationships) SetReturnLineItems(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetReturnLineItems sets ReturnLineItems field to given value.

### HasReturnLineItems

`func (o *GETReturns200ResponseDataInnerRelationships) HasReturnLineItems() bool`

HasReturnLineItems returns a boolean if a field has been set.

### GetAttachments

`func (o *GETReturns200ResponseDataInnerRelationships) GetAttachments() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETReturns200ResponseDataInnerRelationships) SetAttachments(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETReturns200ResponseDataInnerRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GETReturns200ResponseDataInnerRelationships) GetEvents() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETReturns200ResponseDataInnerRelationships) GetEventsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETReturns200ResponseDataInnerRelationships) SetEvents(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETReturns200ResponseDataInnerRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


