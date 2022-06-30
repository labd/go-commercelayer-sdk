# CustomerAddressUpdateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**Address** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 

## Methods

### NewCustomerAddressUpdateDataRelationships

`func NewCustomerAddressUpdateDataRelationships() *CustomerAddressUpdateDataRelationships`

NewCustomerAddressUpdateDataRelationships instantiates a new CustomerAddressUpdateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddressUpdateDataRelationshipsWithDefaults

`func NewCustomerAddressUpdateDataRelationshipsWithDefaults() *CustomerAddressUpdateDataRelationships`

NewCustomerAddressUpdateDataRelationshipsWithDefaults instantiates a new CustomerAddressUpdateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *CustomerAddressUpdateDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerAddressUpdateDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerAddressUpdateDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CustomerAddressUpdateDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetAddress

`func (o *CustomerAddressUpdateDataRelationships) GetAddress() BingGeocoderDataRelationshipsAddresses`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerAddressUpdateDataRelationships) GetAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerAddressUpdateDataRelationships) SetAddress(v BingGeocoderDataRelationshipsAddresses)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustomerAddressUpdateDataRelationships) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


