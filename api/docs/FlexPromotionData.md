# FlexPromotionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETFlexPromotionsFlexPromotionId200ResponseDataAttributes**](GETFlexPromotionsFlexPromotionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**FlexPromotionDataRelationships**](FlexPromotionDataRelationships.md) |  | [optional] 

## Methods

### NewFlexPromotionData

`func NewFlexPromotionData(type_ interface{}, attributes GETFlexPromotionsFlexPromotionId200ResponseDataAttributes, ) *FlexPromotionData`

NewFlexPromotionData instantiates a new FlexPromotionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexPromotionDataWithDefaults

`func NewFlexPromotionDataWithDefaults() *FlexPromotionData`

NewFlexPromotionDataWithDefaults instantiates a new FlexPromotionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlexPromotionData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlexPromotionData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlexPromotionData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *FlexPromotionData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FlexPromotionData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *FlexPromotionData) GetAttributes() GETFlexPromotionsFlexPromotionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FlexPromotionData) GetAttributesOk() (*GETFlexPromotionsFlexPromotionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FlexPromotionData) SetAttributes(v GETFlexPromotionsFlexPromotionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FlexPromotionData) GetRelationships() FlexPromotionDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FlexPromotionData) GetRelationshipsOk() (*FlexPromotionDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FlexPromotionData) SetRelationships(v FlexPromotionDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FlexPromotionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


