# CustomerPasswordResetUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**CustomerPasswordResetUpdateDataAttributes**](CustomerPasswordResetUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCustomerPasswordResetUpdateData

`func NewCustomerPasswordResetUpdateData(type_ string, id string, attributes CustomerPasswordResetUpdateDataAttributes, ) *CustomerPasswordResetUpdateData`

NewCustomerPasswordResetUpdateData instantiates a new CustomerPasswordResetUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPasswordResetUpdateDataWithDefaults

`func NewCustomerPasswordResetUpdateDataWithDefaults() *CustomerPasswordResetUpdateData`

NewCustomerPasswordResetUpdateDataWithDefaults instantiates a new CustomerPasswordResetUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerPasswordResetUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerPasswordResetUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerPasswordResetUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CustomerPasswordResetUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerPasswordResetUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerPasswordResetUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CustomerPasswordResetUpdateData) GetAttributes() CustomerPasswordResetUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerPasswordResetUpdateData) GetAttributesOk() (*CustomerPasswordResetUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerPasswordResetUpdateData) SetAttributes(v CustomerPasswordResetUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerPasswordResetUpdateData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerPasswordResetUpdateData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerPasswordResetUpdateData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerPasswordResetUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


