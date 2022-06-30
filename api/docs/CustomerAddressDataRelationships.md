# CustomerAddressDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**Address** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressDataRelationshipsEvents**](CustomerAddressDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewCustomerAddressDataRelationships

`func NewCustomerAddressDataRelationships() *CustomerAddressDataRelationships`

NewCustomerAddressDataRelationships instantiates a new CustomerAddressDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddressDataRelationshipsWithDefaults

`func NewCustomerAddressDataRelationshipsWithDefaults() *CustomerAddressDataRelationships`

NewCustomerAddressDataRelationshipsWithDefaults instantiates a new CustomerAddressDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *CustomerAddressDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerAddressDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerAddressDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CustomerAddressDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetAddress

`func (o *CustomerAddressDataRelationships) GetAddress() BingGeocoderDataRelationshipsAddresses`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerAddressDataRelationships) GetAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerAddressDataRelationships) SetAddress(v BingGeocoderDataRelationshipsAddresses)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustomerAddressDataRelationships) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEvents

`func (o *CustomerAddressDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CustomerAddressDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CustomerAddressDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CustomerAddressDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


