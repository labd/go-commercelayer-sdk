# CaptureDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**ReferenceAuthorization** | Pointer to [**CaptureDataRelationshipsReferenceAuthorization**](CaptureDataRelationshipsReferenceAuthorization.md) |  | [optional] 
**Refunds** | Pointer to [**CaptureDataRelationshipsRefunds**](CaptureDataRelationshipsRefunds.md) |  | [optional] 
**Events** | Pointer to [**AuthorizationDataRelationshipsEvents**](AuthorizationDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewCaptureDataRelationships

`func NewCaptureDataRelationships() *CaptureDataRelationships`

NewCaptureDataRelationships instantiates a new CaptureDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureDataRelationshipsWithDefaults

`func NewCaptureDataRelationshipsWithDefaults() *CaptureDataRelationships`

NewCaptureDataRelationshipsWithDefaults instantiates a new CaptureDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *CaptureDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CaptureDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CaptureDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CaptureDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReferenceAuthorization

`func (o *CaptureDataRelationships) GetReferenceAuthorization() CaptureDataRelationshipsReferenceAuthorization`

GetReferenceAuthorization returns the ReferenceAuthorization field if non-nil, zero value otherwise.

### GetReferenceAuthorizationOk

`func (o *CaptureDataRelationships) GetReferenceAuthorizationOk() (*CaptureDataRelationshipsReferenceAuthorization, bool)`

GetReferenceAuthorizationOk returns a tuple with the ReferenceAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceAuthorization

`func (o *CaptureDataRelationships) SetReferenceAuthorization(v CaptureDataRelationshipsReferenceAuthorization)`

SetReferenceAuthorization sets ReferenceAuthorization field to given value.

### HasReferenceAuthorization

`func (o *CaptureDataRelationships) HasReferenceAuthorization() bool`

HasReferenceAuthorization returns a boolean if a field has been set.

### GetRefunds

`func (o *CaptureDataRelationships) GetRefunds() CaptureDataRelationshipsRefunds`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *CaptureDataRelationships) GetRefundsOk() (*CaptureDataRelationshipsRefunds, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *CaptureDataRelationships) SetRefunds(v CaptureDataRelationshipsRefunds)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *CaptureDataRelationships) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetEvents

`func (o *CaptureDataRelationships) GetEvents() AuthorizationDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CaptureDataRelationships) GetEventsOk() (*AuthorizationDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CaptureDataRelationships) SetEvents(v AuthorizationDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CaptureDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


