# AuthorizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "authorizations"]
**Attributes** | [**AuthorizationDataAttributes**](AuthorizationDataAttributes.md) |  | 
**Relationships** | Pointer to [**AuthorizationDataRelationships**](AuthorizationDataRelationships.md) |  | [optional] 

## Methods

### NewAuthorizationData

`func NewAuthorizationData(type_ string, attributes AuthorizationDataAttributes, ) *AuthorizationData`

NewAuthorizationData instantiates a new AuthorizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationDataWithDefaults

`func NewAuthorizationDataWithDefaults() *AuthorizationData`

NewAuthorizationDataWithDefaults instantiates a new AuthorizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizationData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizationData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizationData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AuthorizationData) GetAttributes() AuthorizationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AuthorizationData) GetAttributesOk() (*AuthorizationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AuthorizationData) SetAttributes(v AuthorizationDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AuthorizationData) GetRelationships() AuthorizationDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AuthorizationData) GetRelationshipsOk() (*AuthorizationDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AuthorizationData) SetRelationships(v AuthorizationDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AuthorizationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


