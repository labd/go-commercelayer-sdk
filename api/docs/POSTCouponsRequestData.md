# POSTCouponsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTCouponsRequestDataAttributes**](POSTCouponsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTCouponsRequestDataRelationships**](POSTCouponsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTCouponsRequestData

`func NewPOSTCouponsRequestData(type_ interface{}, attributes POSTCouponsRequestDataAttributes, ) *POSTCouponsRequestData`

NewPOSTCouponsRequestData instantiates a new POSTCouponsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCouponsRequestDataWithDefaults

`func NewPOSTCouponsRequestDataWithDefaults() *POSTCouponsRequestData`

NewPOSTCouponsRequestDataWithDefaults instantiates a new POSTCouponsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTCouponsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTCouponsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTCouponsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTCouponsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTCouponsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTCouponsRequestData) GetAttributes() POSTCouponsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTCouponsRequestData) GetAttributesOk() (*POSTCouponsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTCouponsRequestData) SetAttributes(v POSTCouponsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTCouponsRequestData) GetRelationships() POSTCouponsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTCouponsRequestData) GetRelationshipsOk() (*POSTCouponsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTCouponsRequestData) SetRelationships(v POSTCouponsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTCouponsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


