# SubscriptionModelData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes**](GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**SubscriptionModelDataRelationships**](SubscriptionModelDataRelationships.md) |  | [optional] 

## Methods

### NewSubscriptionModelData

`func NewSubscriptionModelData(type_ interface{}, attributes GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes, ) *SubscriptionModelData`

NewSubscriptionModelData instantiates a new SubscriptionModelData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionModelDataWithDefaults

`func NewSubscriptionModelDataWithDefaults() *SubscriptionModelData`

NewSubscriptionModelDataWithDefaults instantiates a new SubscriptionModelData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubscriptionModelData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionModelData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionModelData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SubscriptionModelData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SubscriptionModelData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *SubscriptionModelData) GetAttributes() GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubscriptionModelData) GetAttributesOk() (*GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubscriptionModelData) SetAttributes(v GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SubscriptionModelData) GetRelationships() SubscriptionModelDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SubscriptionModelData) GetRelationshipsOk() (*SubscriptionModelDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SubscriptionModelData) SetRelationships(v SubscriptionModelDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SubscriptionModelData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


