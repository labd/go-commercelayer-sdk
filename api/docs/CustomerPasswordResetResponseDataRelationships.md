# CustomerPasswordResetResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**CouponRecipientResponseDataRelationshipsCustomer**](CouponRecipientResponseDataRelationshipsCustomer.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressResponseDataRelationshipsEvents**](CustomerAddressResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewCustomerPasswordResetResponseDataRelationships

`func NewCustomerPasswordResetResponseDataRelationships() *CustomerPasswordResetResponseDataRelationships`

NewCustomerPasswordResetResponseDataRelationships instantiates a new CustomerPasswordResetResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPasswordResetResponseDataRelationshipsWithDefaults

`func NewCustomerPasswordResetResponseDataRelationshipsWithDefaults() *CustomerPasswordResetResponseDataRelationships`

NewCustomerPasswordResetResponseDataRelationshipsWithDefaults instantiates a new CustomerPasswordResetResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *CustomerPasswordResetResponseDataRelationships) GetCustomer() CouponRecipientResponseDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerPasswordResetResponseDataRelationships) GetCustomerOk() (*CouponRecipientResponseDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerPasswordResetResponseDataRelationships) SetCustomer(v CouponRecipientResponseDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CustomerPasswordResetResponseDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetEvents

`func (o *CustomerPasswordResetResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CustomerPasswordResetResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CustomerPasswordResetResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CustomerPasswordResetResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


