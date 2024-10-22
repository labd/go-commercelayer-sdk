# OrderSubscriptionUpdateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerPaymentSource** | Pointer to [**OrderSubscriptionUpdateDataRelationshipsCustomerPaymentSource**](OrderSubscriptionUpdateDataRelationshipsCustomerPaymentSource.md) |  | [optional] 
**Tags** | Pointer to [**AddressCreateDataRelationshipsTags**](AddressCreateDataRelationshipsTags.md) |  | [optional] 

## Methods

### NewOrderSubscriptionUpdateDataRelationships

`func NewOrderSubscriptionUpdateDataRelationships() *OrderSubscriptionUpdateDataRelationships`

NewOrderSubscriptionUpdateDataRelationships instantiates a new OrderSubscriptionUpdateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionUpdateDataRelationshipsWithDefaults

`func NewOrderSubscriptionUpdateDataRelationshipsWithDefaults() *OrderSubscriptionUpdateDataRelationships`

NewOrderSubscriptionUpdateDataRelationshipsWithDefaults instantiates a new OrderSubscriptionUpdateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerPaymentSource

`func (o *OrderSubscriptionUpdateDataRelationships) GetCustomerPaymentSource() OrderSubscriptionUpdateDataRelationshipsCustomerPaymentSource`

GetCustomerPaymentSource returns the CustomerPaymentSource field if non-nil, zero value otherwise.

### GetCustomerPaymentSourceOk

`func (o *OrderSubscriptionUpdateDataRelationships) GetCustomerPaymentSourceOk() (*OrderSubscriptionUpdateDataRelationshipsCustomerPaymentSource, bool)`

GetCustomerPaymentSourceOk returns a tuple with the CustomerPaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPaymentSource

`func (o *OrderSubscriptionUpdateDataRelationships) SetCustomerPaymentSource(v OrderSubscriptionUpdateDataRelationshipsCustomerPaymentSource)`

SetCustomerPaymentSource sets CustomerPaymentSource field to given value.

### HasCustomerPaymentSource

`func (o *OrderSubscriptionUpdateDataRelationships) HasCustomerPaymentSource() bool`

HasCustomerPaymentSource returns a boolean if a field has been set.

### GetTags

`func (o *OrderSubscriptionUpdateDataRelationships) GetTags() AddressCreateDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrderSubscriptionUpdateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrderSubscriptionUpdateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrderSubscriptionUpdateDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


