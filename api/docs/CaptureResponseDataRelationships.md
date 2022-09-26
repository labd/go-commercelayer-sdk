# CaptureResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentResponseDataRelationshipsOrder**](AdyenPaymentResponseDataRelationshipsOrder.md) |  | [optional] 
**ReferenceAuthorization** | Pointer to [**CaptureResponseDataRelationshipsReferenceAuthorization**](CaptureResponseDataRelationshipsReferenceAuthorization.md) |  | [optional] 
**Refunds** | Pointer to [**CaptureResponseDataRelationshipsRefunds**](CaptureResponseDataRelationshipsRefunds.md) |  | [optional] 

## Methods

### NewCaptureResponseDataRelationships

`func NewCaptureResponseDataRelationships() *CaptureResponseDataRelationships`

NewCaptureResponseDataRelationships instantiates a new CaptureResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureResponseDataRelationshipsWithDefaults

`func NewCaptureResponseDataRelationshipsWithDefaults() *CaptureResponseDataRelationships`

NewCaptureResponseDataRelationshipsWithDefaults instantiates a new CaptureResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *CaptureResponseDataRelationships) GetOrder() AdyenPaymentResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CaptureResponseDataRelationships) GetOrderOk() (*AdyenPaymentResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CaptureResponseDataRelationships) SetOrder(v AdyenPaymentResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CaptureResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReferenceAuthorization

`func (o *CaptureResponseDataRelationships) GetReferenceAuthorization() CaptureResponseDataRelationshipsReferenceAuthorization`

GetReferenceAuthorization returns the ReferenceAuthorization field if non-nil, zero value otherwise.

### GetReferenceAuthorizationOk

`func (o *CaptureResponseDataRelationships) GetReferenceAuthorizationOk() (*CaptureResponseDataRelationshipsReferenceAuthorization, bool)`

GetReferenceAuthorizationOk returns a tuple with the ReferenceAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceAuthorization

`func (o *CaptureResponseDataRelationships) SetReferenceAuthorization(v CaptureResponseDataRelationshipsReferenceAuthorization)`

SetReferenceAuthorization sets ReferenceAuthorization field to given value.

### HasReferenceAuthorization

`func (o *CaptureResponseDataRelationships) HasReferenceAuthorization() bool`

HasReferenceAuthorization returns a boolean if a field has been set.

### GetRefunds

`func (o *CaptureResponseDataRelationships) GetRefunds() CaptureResponseDataRelationshipsRefunds`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *CaptureResponseDataRelationships) GetRefundsOk() (*CaptureResponseDataRelationshipsRefunds, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *CaptureResponseDataRelationships) SetRefunds(v CaptureResponseDataRelationshipsRefunds)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *CaptureResponseDataRelationships) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


