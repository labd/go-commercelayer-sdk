# FlexPromotionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTFlexPromotions201ResponseDataAttributes**](POSTFlexPromotions201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**FlexPromotionCreateDataRelationships**](FlexPromotionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewFlexPromotionCreateData

`func NewFlexPromotionCreateData(type_ interface{}, attributes POSTFlexPromotions201ResponseDataAttributes, ) *FlexPromotionCreateData`

NewFlexPromotionCreateData instantiates a new FlexPromotionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexPromotionCreateDataWithDefaults

`func NewFlexPromotionCreateDataWithDefaults() *FlexPromotionCreateData`

NewFlexPromotionCreateDataWithDefaults instantiates a new FlexPromotionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlexPromotionCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlexPromotionCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlexPromotionCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *FlexPromotionCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FlexPromotionCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *FlexPromotionCreateData) GetAttributes() POSTFlexPromotions201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FlexPromotionCreateData) GetAttributesOk() (*POSTFlexPromotions201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FlexPromotionCreateData) SetAttributes(v POSTFlexPromotions201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FlexPromotionCreateData) GetRelationships() FlexPromotionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FlexPromotionCreateData) GetRelationshipsOk() (*FlexPromotionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FlexPromotionCreateData) SetRelationships(v FlexPromotionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FlexPromotionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


