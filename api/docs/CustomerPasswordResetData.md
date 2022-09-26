# CustomerPasswordResetData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**CustomerPasswordResetDataAttributes**](CustomerPasswordResetDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerPasswordResetDataRelationships**](CustomerPasswordResetDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerPasswordResetData

`func NewCustomerPasswordResetData(type_ string, attributes CustomerPasswordResetDataAttributes, ) *CustomerPasswordResetData`

NewCustomerPasswordResetData instantiates a new CustomerPasswordResetData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPasswordResetDataWithDefaults

`func NewCustomerPasswordResetDataWithDefaults() *CustomerPasswordResetData`

NewCustomerPasswordResetDataWithDefaults instantiates a new CustomerPasswordResetData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerPasswordResetData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerPasswordResetData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerPasswordResetData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CustomerPasswordResetData) GetAttributes() CustomerPasswordResetDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerPasswordResetData) GetAttributesOk() (*CustomerPasswordResetDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerPasswordResetData) SetAttributes(v CustomerPasswordResetDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerPasswordResetData) GetRelationships() CustomerPasswordResetDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerPasswordResetData) GetRelationshipsOk() (*CustomerPasswordResetDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerPasswordResetData) SetRelationships(v CustomerPasswordResetDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerPasswordResetData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


