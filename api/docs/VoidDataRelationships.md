# VoidDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**ReferenceAuthorization** | Pointer to [**CaptureDataRelationshipsReferenceAuthorization**](CaptureDataRelationshipsReferenceAuthorization.md) |  | [optional] 
**Events** | Pointer to [**AuthorizationDataRelationshipsEvents**](AuthorizationDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewVoidDataRelationships

`func NewVoidDataRelationships() *VoidDataRelationships`

NewVoidDataRelationships instantiates a new VoidDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidDataRelationshipsWithDefaults

`func NewVoidDataRelationshipsWithDefaults() *VoidDataRelationships`

NewVoidDataRelationshipsWithDefaults instantiates a new VoidDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *VoidDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VoidDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VoidDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VoidDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReferenceAuthorization

`func (o *VoidDataRelationships) GetReferenceAuthorization() CaptureDataRelationshipsReferenceAuthorization`

GetReferenceAuthorization returns the ReferenceAuthorization field if non-nil, zero value otherwise.

### GetReferenceAuthorizationOk

`func (o *VoidDataRelationships) GetReferenceAuthorizationOk() (*CaptureDataRelationshipsReferenceAuthorization, bool)`

GetReferenceAuthorizationOk returns a tuple with the ReferenceAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceAuthorization

`func (o *VoidDataRelationships) SetReferenceAuthorization(v CaptureDataRelationshipsReferenceAuthorization)`

SetReferenceAuthorization sets ReferenceAuthorization field to given value.

### HasReferenceAuthorization

`func (o *VoidDataRelationships) HasReferenceAuthorization() bool`

HasReferenceAuthorization returns a boolean if a field has been set.

### GetEvents

`func (o *VoidDataRelationships) GetEvents() AuthorizationDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *VoidDataRelationships) GetEventsOk() (*AuthorizationDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *VoidDataRelationships) SetEvents(v AuthorizationDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *VoidDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


