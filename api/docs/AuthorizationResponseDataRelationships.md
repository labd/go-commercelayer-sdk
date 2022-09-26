# AuthorizationResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentResponseDataRelationshipsOrder**](AdyenPaymentResponseDataRelationshipsOrder.md) |  | [optional] 
**Captures** | Pointer to [**AuthorizationResponseDataRelationshipsCaptures**](AuthorizationResponseDataRelationshipsCaptures.md) |  | [optional] 
**Voids** | Pointer to [**AuthorizationResponseDataRelationshipsVoids**](AuthorizationResponseDataRelationshipsVoids.md) |  | [optional] 

## Methods

### NewAuthorizationResponseDataRelationships

`func NewAuthorizationResponseDataRelationships() *AuthorizationResponseDataRelationships`

NewAuthorizationResponseDataRelationships instantiates a new AuthorizationResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationResponseDataRelationshipsWithDefaults

`func NewAuthorizationResponseDataRelationshipsWithDefaults() *AuthorizationResponseDataRelationships`

NewAuthorizationResponseDataRelationshipsWithDefaults instantiates a new AuthorizationResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *AuthorizationResponseDataRelationships) GetOrder() AdyenPaymentResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AuthorizationResponseDataRelationships) GetOrderOk() (*AdyenPaymentResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AuthorizationResponseDataRelationships) SetOrder(v AdyenPaymentResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AuthorizationResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCaptures

`func (o *AuthorizationResponseDataRelationships) GetCaptures() AuthorizationResponseDataRelationshipsCaptures`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *AuthorizationResponseDataRelationships) GetCapturesOk() (*AuthorizationResponseDataRelationshipsCaptures, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *AuthorizationResponseDataRelationships) SetCaptures(v AuthorizationResponseDataRelationshipsCaptures)`

SetCaptures sets Captures field to given value.

### HasCaptures

`func (o *AuthorizationResponseDataRelationships) HasCaptures() bool`

HasCaptures returns a boolean if a field has been set.

### GetVoids

`func (o *AuthorizationResponseDataRelationships) GetVoids() AuthorizationResponseDataRelationshipsVoids`

GetVoids returns the Voids field if non-nil, zero value otherwise.

### GetVoidsOk

`func (o *AuthorizationResponseDataRelationships) GetVoidsOk() (*AuthorizationResponseDataRelationshipsVoids, bool)`

GetVoidsOk returns a tuple with the Voids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoids

`func (o *AuthorizationResponseDataRelationships) SetVoids(v AuthorizationResponseDataRelationshipsVoids)`

SetVoids sets Voids field to given value.

### HasVoids

`func (o *AuthorizationResponseDataRelationships) HasVoids() bool`

HasVoids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


