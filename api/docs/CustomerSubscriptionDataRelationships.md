# CustomerSubscriptionDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewCustomerSubscriptionDataRelationships

`func NewCustomerSubscriptionDataRelationships() *CustomerSubscriptionDataRelationships`

NewCustomerSubscriptionDataRelationships instantiates a new CustomerSubscriptionDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSubscriptionDataRelationshipsWithDefaults

`func NewCustomerSubscriptionDataRelationshipsWithDefaults() *CustomerSubscriptionDataRelationships`

NewCustomerSubscriptionDataRelationshipsWithDefaults instantiates a new CustomerSubscriptionDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *CustomerSubscriptionDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerSubscriptionDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerSubscriptionDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CustomerSubscriptionDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetEvents

`func (o *CustomerSubscriptionDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CustomerSubscriptionDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CustomerSubscriptionDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CustomerSubscriptionDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *CustomerSubscriptionDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CustomerSubscriptionDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CustomerSubscriptionDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CustomerSubscriptionDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


