# ReturnDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**StockLocation** | Pointer to [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | [optional] 
**OriginAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**DestinationAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**ReferenceCapture** | Pointer to [**AuthorizationDataRelationshipsCaptures**](AuthorizationDataRelationshipsCaptures.md) |  | [optional] 
**ReferenceRefund** | Pointer to [**CaptureDataRelationshipsRefunds**](CaptureDataRelationshipsRefunds.md) |  | [optional] 
**ReturnLineItems** | Pointer to [**LineItemDataRelationshipsReturnLineItems**](LineItemDataRelationshipsReturnLineItems.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**ResourceErrors** | Pointer to [**OrderDataRelationshipsResourceErrors**](OrderDataRelationshipsResourceErrors.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**AddressDataRelationshipsTags**](AddressDataRelationshipsTags.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

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

### GetReferenceCapture

`func (o *ReturnDataRelationships) GetReferenceCapture() AuthorizationDataRelationshipsCaptures`

GetReferenceCapture returns the ReferenceCapture field if non-nil, zero value otherwise.

### GetReferenceCaptureOk

`func (o *ReturnDataRelationships) GetReferenceCaptureOk() (*AuthorizationDataRelationshipsCaptures, bool)`

GetReferenceCaptureOk returns a tuple with the ReferenceCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCapture

`func (o *ReturnDataRelationships) SetReferenceCapture(v AuthorizationDataRelationshipsCaptures)`

SetReferenceCapture sets ReferenceCapture field to given value.

### HasReferenceCapture

`func (o *ReturnDataRelationships) HasReferenceCapture() bool`

HasReferenceCapture returns a boolean if a field has been set.

### GetReferenceRefund

`func (o *ReturnDataRelationships) GetReferenceRefund() CaptureDataRelationshipsRefunds`

GetReferenceRefund returns the ReferenceRefund field if non-nil, zero value otherwise.

### GetReferenceRefundOk

`func (o *ReturnDataRelationships) GetReferenceRefundOk() (*CaptureDataRelationshipsRefunds, bool)`

GetReferenceRefundOk returns a tuple with the ReferenceRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceRefund

`func (o *ReturnDataRelationships) SetReferenceRefund(v CaptureDataRelationshipsRefunds)`

SetReferenceRefund sets ReferenceRefund field to given value.

### HasReferenceRefund

`func (o *ReturnDataRelationships) HasReferenceRefund() bool`

HasReferenceRefund returns a boolean if a field has been set.

### GetReturnLineItems

`func (o *ReturnDataRelationships) GetReturnLineItems() LineItemDataRelationshipsReturnLineItems`

GetReturnLineItems returns the ReturnLineItems field if non-nil, zero value otherwise.

### GetReturnLineItemsOk

`func (o *ReturnDataRelationships) GetReturnLineItemsOk() (*LineItemDataRelationshipsReturnLineItems, bool)`

GetReturnLineItemsOk returns a tuple with the ReturnLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLineItems

`func (o *ReturnDataRelationships) SetReturnLineItems(v LineItemDataRelationshipsReturnLineItems)`

SetReturnLineItems sets ReturnLineItems field to given value.

### HasReturnLineItems

`func (o *ReturnDataRelationships) HasReturnLineItems() bool`

HasReturnLineItems returns a boolean if a field has been set.

### GetAttachments

`func (o *ReturnDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ReturnDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ReturnDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ReturnDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetResourceErrors

`func (o *ReturnDataRelationships) GetResourceErrors() OrderDataRelationshipsResourceErrors`

GetResourceErrors returns the ResourceErrors field if non-nil, zero value otherwise.

### GetResourceErrorsOk

`func (o *ReturnDataRelationships) GetResourceErrorsOk() (*OrderDataRelationshipsResourceErrors, bool)`

GetResourceErrorsOk returns a tuple with the ResourceErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceErrors

`func (o *ReturnDataRelationships) SetResourceErrors(v OrderDataRelationshipsResourceErrors)`

SetResourceErrors sets ResourceErrors field to given value.

### HasResourceErrors

`func (o *ReturnDataRelationships) HasResourceErrors() bool`

HasResourceErrors returns a boolean if a field has been set.

### GetEvents

`func (o *ReturnDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ReturnDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ReturnDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ReturnDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *ReturnDataRelationships) GetTags() AddressDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReturnDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReturnDataRelationships) SetTags(v AddressDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReturnDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *ReturnDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ReturnDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ReturnDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ReturnDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


