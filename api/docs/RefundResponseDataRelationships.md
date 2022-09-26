# RefundResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentResponseDataRelationshipsOrder**](AdyenPaymentResponseDataRelationshipsOrder.md) |  | [optional] 
**ReferenceCapture** | Pointer to [**RefundResponseDataRelationshipsReferenceCapture**](RefundResponseDataRelationshipsReferenceCapture.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressResponseDataRelationshipsEvents**](CustomerAddressResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewRefundResponseDataRelationships

`func NewRefundResponseDataRelationships() *RefundResponseDataRelationships`

NewRefundResponseDataRelationships instantiates a new RefundResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundResponseDataRelationshipsWithDefaults

`func NewRefundResponseDataRelationshipsWithDefaults() *RefundResponseDataRelationships`

NewRefundResponseDataRelationshipsWithDefaults instantiates a new RefundResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *RefundResponseDataRelationships) GetOrder() AdyenPaymentResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RefundResponseDataRelationships) GetOrderOk() (*AdyenPaymentResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RefundResponseDataRelationships) SetOrder(v AdyenPaymentResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RefundResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReferenceCapture

`func (o *RefundResponseDataRelationships) GetReferenceCapture() RefundResponseDataRelationshipsReferenceCapture`

GetReferenceCapture returns the ReferenceCapture field if non-nil, zero value otherwise.

### GetReferenceCaptureOk

`func (o *RefundResponseDataRelationships) GetReferenceCaptureOk() (*RefundResponseDataRelationshipsReferenceCapture, bool)`

GetReferenceCaptureOk returns a tuple with the ReferenceCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCapture

`func (o *RefundResponseDataRelationships) SetReferenceCapture(v RefundResponseDataRelationshipsReferenceCapture)`

SetReferenceCapture sets ReferenceCapture field to given value.

### HasReferenceCapture

`func (o *RefundResponseDataRelationships) HasReferenceCapture() bool`

HasReferenceCapture returns a boolean if a field has been set.

### GetEvents

`func (o *RefundResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RefundResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RefundResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RefundResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


