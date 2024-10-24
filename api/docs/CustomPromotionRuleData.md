# CustomPromotionRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes**](GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomPromotionRuleDataRelationships**](CustomPromotionRuleDataRelationships.md) |  | [optional] 

## Methods

### NewCustomPromotionRuleData

`func NewCustomPromotionRuleData(type_ interface{}, attributes GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes, ) *CustomPromotionRuleData`

NewCustomPromotionRuleData instantiates a new CustomPromotionRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPromotionRuleDataWithDefaults

`func NewCustomPromotionRuleDataWithDefaults() *CustomPromotionRuleData`

NewCustomPromotionRuleDataWithDefaults instantiates a new CustomPromotionRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomPromotionRuleData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomPromotionRuleData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomPromotionRuleData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CustomPromotionRuleData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomPromotionRuleData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CustomPromotionRuleData) GetAttributes() GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomPromotionRuleData) GetAttributesOk() (*GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomPromotionRuleData) SetAttributes(v GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomPromotionRuleData) GetRelationships() CustomPromotionRuleDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomPromotionRuleData) GetRelationshipsOk() (*CustomPromotionRuleDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomPromotionRuleData) SetRelationships(v CustomPromotionRuleDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomPromotionRuleData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


