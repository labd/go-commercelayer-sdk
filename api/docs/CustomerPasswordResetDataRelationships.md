# CustomerPasswordResetDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**Events** | Pointer to [**CleanupDataRelationshipsEvents**](CleanupDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewCustomerPasswordResetDataRelationships

`func NewCustomerPasswordResetDataRelationships() *CustomerPasswordResetDataRelationships`

NewCustomerPasswordResetDataRelationships instantiates a new CustomerPasswordResetDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPasswordResetDataRelationshipsWithDefaults

`func NewCustomerPasswordResetDataRelationshipsWithDefaults() *CustomerPasswordResetDataRelationships`

NewCustomerPasswordResetDataRelationshipsWithDefaults instantiates a new CustomerPasswordResetDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *CustomerPasswordResetDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerPasswordResetDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerPasswordResetDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CustomerPasswordResetDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetEvents

`func (o *CustomerPasswordResetDataRelationships) GetEvents() CleanupDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CustomerPasswordResetDataRelationships) GetEventsOk() (*CleanupDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CustomerPasswordResetDataRelationships) SetEvents(v CleanupDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CustomerPasswordResetDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


