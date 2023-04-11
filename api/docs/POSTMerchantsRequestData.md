# POSTMerchantsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTMerchantsRequestDataAttributes**](POSTMerchantsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTMerchantsRequestDataRelationships**](POSTMerchantsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTMerchantsRequestData

`func NewPOSTMerchantsRequestData(type_ interface{}, attributes POSTMerchantsRequestDataAttributes, ) *POSTMerchantsRequestData`

NewPOSTMerchantsRequestData instantiates a new POSTMerchantsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTMerchantsRequestDataWithDefaults

`func NewPOSTMerchantsRequestDataWithDefaults() *POSTMerchantsRequestData`

NewPOSTMerchantsRequestDataWithDefaults instantiates a new POSTMerchantsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTMerchantsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTMerchantsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTMerchantsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTMerchantsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTMerchantsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTMerchantsRequestData) GetAttributes() POSTMerchantsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTMerchantsRequestData) GetAttributesOk() (*POSTMerchantsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTMerchantsRequestData) SetAttributes(v POSTMerchantsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTMerchantsRequestData) GetRelationships() POSTMerchantsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTMerchantsRequestData) GetRelationshipsOk() (*POSTMerchantsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTMerchantsRequestData) SetRelationships(v POSTMerchantsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTMerchantsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


