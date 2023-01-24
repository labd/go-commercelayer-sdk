# RefundDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**ReferenceCapture** | Pointer to [**AuthorizationDataRelationshipsCaptures**](AuthorizationDataRelationshipsCaptures.md) |  | [optional] 
**Events** | Pointer to [**CleanupDataRelationshipsEvents**](CleanupDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewRefundDataRelationships

`func NewRefundDataRelationships() *RefundDataRelationships`

NewRefundDataRelationships instantiates a new RefundDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundDataRelationshipsWithDefaults

`func NewRefundDataRelationshipsWithDefaults() *RefundDataRelationships`

NewRefundDataRelationshipsWithDefaults instantiates a new RefundDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *RefundDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RefundDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RefundDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RefundDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReferenceCapture

`func (o *RefundDataRelationships) GetReferenceCapture() AuthorizationDataRelationshipsCaptures`

GetReferenceCapture returns the ReferenceCapture field if non-nil, zero value otherwise.

### GetReferenceCaptureOk

`func (o *RefundDataRelationships) GetReferenceCaptureOk() (*AuthorizationDataRelationshipsCaptures, bool)`

GetReferenceCaptureOk returns a tuple with the ReferenceCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCapture

`func (o *RefundDataRelationships) SetReferenceCapture(v AuthorizationDataRelationshipsCaptures)`

SetReferenceCapture sets ReferenceCapture field to given value.

### HasReferenceCapture

`func (o *RefundDataRelationships) HasReferenceCapture() bool`

HasReferenceCapture returns a boolean if a field has been set.

### GetEvents

`func (o *RefundDataRelationships) GetEvents() CleanupDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RefundDataRelationships) GetEventsOk() (*CleanupDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RefundDataRelationships) SetEvents(v CleanupDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RefundDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


