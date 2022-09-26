# VoidResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentResponseDataRelationshipsOrder**](AdyenPaymentResponseDataRelationshipsOrder.md) |  | [optional] 
**ReferenceAuthorization** | Pointer to [**CaptureResponseDataRelationshipsReferenceAuthorization**](CaptureResponseDataRelationshipsReferenceAuthorization.md) |  | [optional] 

## Methods

### NewVoidResponseDataRelationships

`func NewVoidResponseDataRelationships() *VoidResponseDataRelationships`

NewVoidResponseDataRelationships instantiates a new VoidResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidResponseDataRelationshipsWithDefaults

`func NewVoidResponseDataRelationshipsWithDefaults() *VoidResponseDataRelationships`

NewVoidResponseDataRelationshipsWithDefaults instantiates a new VoidResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *VoidResponseDataRelationships) GetOrder() AdyenPaymentResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VoidResponseDataRelationships) GetOrderOk() (*AdyenPaymentResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VoidResponseDataRelationships) SetOrder(v AdyenPaymentResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VoidResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReferenceAuthorization

`func (o *VoidResponseDataRelationships) GetReferenceAuthorization() CaptureResponseDataRelationshipsReferenceAuthorization`

GetReferenceAuthorization returns the ReferenceAuthorization field if non-nil, zero value otherwise.

### GetReferenceAuthorizationOk

`func (o *VoidResponseDataRelationships) GetReferenceAuthorizationOk() (*CaptureResponseDataRelationshipsReferenceAuthorization, bool)`

GetReferenceAuthorizationOk returns a tuple with the ReferenceAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceAuthorization

`func (o *VoidResponseDataRelationships) SetReferenceAuthorization(v CaptureResponseDataRelationshipsReferenceAuthorization)`

SetReferenceAuthorization sets ReferenceAuthorization field to given value.

### HasReferenceAuthorization

`func (o *VoidResponseDataRelationships) HasReferenceAuthorization() bool`

HasReferenceAuthorization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


