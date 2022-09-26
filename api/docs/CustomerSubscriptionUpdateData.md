# CustomerSubscriptionUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**AdyenPaymentCreateDataAttributes**](AdyenPaymentCreateDataAttributes.md) |  | 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCustomerSubscriptionUpdateData

`func NewCustomerSubscriptionUpdateData(type_ string, id string, attributes AdyenPaymentCreateDataAttributes, ) *CustomerSubscriptionUpdateData`

NewCustomerSubscriptionUpdateData instantiates a new CustomerSubscriptionUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSubscriptionUpdateDataWithDefaults

`func NewCustomerSubscriptionUpdateDataWithDefaults() *CustomerSubscriptionUpdateData`

NewCustomerSubscriptionUpdateDataWithDefaults instantiates a new CustomerSubscriptionUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerSubscriptionUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerSubscriptionUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerSubscriptionUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CustomerSubscriptionUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerSubscriptionUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerSubscriptionUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CustomerSubscriptionUpdateData) GetAttributes() AdyenPaymentCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerSubscriptionUpdateData) GetAttributesOk() (*AdyenPaymentCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerSubscriptionUpdateData) SetAttributes(v AdyenPaymentCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerSubscriptionUpdateData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerSubscriptionUpdateData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerSubscriptionUpdateData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerSubscriptionUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


