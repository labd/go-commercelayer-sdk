# CustomerSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETCustomerSubscriptionsCustomerSubscriptionId200ResponseDataAttributes**](GETCustomerSubscriptionsCustomerSubscriptionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerSubscriptionDataRelationships**](CustomerSubscriptionDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerSubscriptionData

`func NewCustomerSubscriptionData(type_ interface{}, attributes GETCustomerSubscriptionsCustomerSubscriptionId200ResponseDataAttributes, ) *CustomerSubscriptionData`

NewCustomerSubscriptionData instantiates a new CustomerSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSubscriptionDataWithDefaults

`func NewCustomerSubscriptionDataWithDefaults() *CustomerSubscriptionData`

NewCustomerSubscriptionDataWithDefaults instantiates a new CustomerSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerSubscriptionData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerSubscriptionData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerSubscriptionData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CustomerSubscriptionData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomerSubscriptionData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CustomerSubscriptionData) GetAttributes() GETCustomerSubscriptionsCustomerSubscriptionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerSubscriptionData) GetAttributesOk() (*GETCustomerSubscriptionsCustomerSubscriptionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerSubscriptionData) SetAttributes(v GETCustomerSubscriptionsCustomerSubscriptionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerSubscriptionData) GetRelationships() CustomerSubscriptionDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerSubscriptionData) GetRelationshipsOk() (*CustomerSubscriptionDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerSubscriptionData) SetRelationships(v CustomerSubscriptionDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerSubscriptionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


