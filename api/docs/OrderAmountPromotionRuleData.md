# OrderAmountPromotionRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETOrderAmountPromotionRules200ResponseDataInnerAttributes**](GETOrderAmountPromotionRules200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**OrderAmountPromotionRuleDataRelationships**](OrderAmountPromotionRuleDataRelationships.md) |  | [optional] 

## Methods

### NewOrderAmountPromotionRuleData

`func NewOrderAmountPromotionRuleData(type_ interface{}, attributes GETOrderAmountPromotionRules200ResponseDataInnerAttributes, ) *OrderAmountPromotionRuleData`

NewOrderAmountPromotionRuleData instantiates a new OrderAmountPromotionRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmountPromotionRuleDataWithDefaults

`func NewOrderAmountPromotionRuleDataWithDefaults() *OrderAmountPromotionRuleData`

NewOrderAmountPromotionRuleDataWithDefaults instantiates a new OrderAmountPromotionRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderAmountPromotionRuleData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderAmountPromotionRuleData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderAmountPromotionRuleData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OrderAmountPromotionRuleData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderAmountPromotionRuleData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *OrderAmountPromotionRuleData) GetAttributes() GETOrderAmountPromotionRules200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderAmountPromotionRuleData) GetAttributesOk() (*GETOrderAmountPromotionRules200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderAmountPromotionRuleData) SetAttributes(v GETOrderAmountPromotionRules200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderAmountPromotionRuleData) GetRelationships() OrderAmountPromotionRuleDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderAmountPromotionRuleData) GetRelationshipsOk() (*OrderAmountPromotionRuleDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderAmountPromotionRuleData) SetRelationships(v OrderAmountPromotionRuleDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderAmountPromotionRuleData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


