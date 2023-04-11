# POSTExternalPromotionsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTExternalPromotionsRequestDataAttributes**](POSTExternalPromotionsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTExternalPromotionsRequestDataRelationships**](POSTExternalPromotionsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTExternalPromotionsRequestData

`func NewPOSTExternalPromotionsRequestData(type_ interface{}, attributes POSTExternalPromotionsRequestDataAttributes, ) *POSTExternalPromotionsRequestData`

NewPOSTExternalPromotionsRequestData instantiates a new POSTExternalPromotionsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTExternalPromotionsRequestDataWithDefaults

`func NewPOSTExternalPromotionsRequestDataWithDefaults() *POSTExternalPromotionsRequestData`

NewPOSTExternalPromotionsRequestDataWithDefaults instantiates a new POSTExternalPromotionsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTExternalPromotionsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTExternalPromotionsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTExternalPromotionsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTExternalPromotionsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTExternalPromotionsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTExternalPromotionsRequestData) GetAttributes() POSTExternalPromotionsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTExternalPromotionsRequestData) GetAttributesOk() (*POSTExternalPromotionsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTExternalPromotionsRequestData) SetAttributes(v POSTExternalPromotionsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTExternalPromotionsRequestData) GetRelationships() POSTExternalPromotionsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTExternalPromotionsRequestData) GetRelationshipsOk() (*POSTExternalPromotionsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTExternalPromotionsRequestData) SetRelationships(v POSTExternalPromotionsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTExternalPromotionsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


