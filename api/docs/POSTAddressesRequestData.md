# POSTAddressesRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTAddressesRequestDataAttributes**](POSTAddressesRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTAddressesRequestDataRelationships**](POSTAddressesRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTAddressesRequestData

`func NewPOSTAddressesRequestData(type_ interface{}, attributes POSTAddressesRequestDataAttributes, ) *POSTAddressesRequestData`

NewPOSTAddressesRequestData instantiates a new POSTAddressesRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTAddressesRequestDataWithDefaults

`func NewPOSTAddressesRequestDataWithDefaults() *POSTAddressesRequestData`

NewPOSTAddressesRequestDataWithDefaults instantiates a new POSTAddressesRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTAddressesRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTAddressesRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTAddressesRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTAddressesRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTAddressesRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTAddressesRequestData) GetAttributes() POSTAddressesRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTAddressesRequestData) GetAttributesOk() (*POSTAddressesRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTAddressesRequestData) SetAttributes(v POSTAddressesRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTAddressesRequestData) GetRelationships() POSTAddressesRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTAddressesRequestData) GetRelationshipsOk() (*POSTAddressesRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTAddressesRequestData) SetRelationships(v POSTAddressesRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTAddressesRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


