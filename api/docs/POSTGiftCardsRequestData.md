# POSTGiftCardsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTGiftCardsRequestDataAttributes**](POSTGiftCardsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTGiftCardsRequestDataRelationships**](POSTGiftCardsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTGiftCardsRequestData

`func NewPOSTGiftCardsRequestData(type_ interface{}, attributes POSTGiftCardsRequestDataAttributes, ) *POSTGiftCardsRequestData`

NewPOSTGiftCardsRequestData instantiates a new POSTGiftCardsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTGiftCardsRequestDataWithDefaults

`func NewPOSTGiftCardsRequestDataWithDefaults() *POSTGiftCardsRequestData`

NewPOSTGiftCardsRequestDataWithDefaults instantiates a new POSTGiftCardsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTGiftCardsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTGiftCardsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTGiftCardsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTGiftCardsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTGiftCardsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTGiftCardsRequestData) GetAttributes() POSTGiftCardsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTGiftCardsRequestData) GetAttributesOk() (*POSTGiftCardsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTGiftCardsRequestData) SetAttributes(v POSTGiftCardsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTGiftCardsRequestData) GetRelationships() POSTGiftCardsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTGiftCardsRequestData) GetRelationshipsOk() (*POSTGiftCardsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTGiftCardsRequestData) SetRelationships(v POSTGiftCardsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTGiftCardsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


