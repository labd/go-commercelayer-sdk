# OrderAmountPromotionRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**OrderAmountPromotionRuleDataAttributes**](OrderAmountPromotionRuleDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderAmountPromotionRuleDataRelationships**](OrderAmountPromotionRuleDataRelationships.md) |  | [optional] 

## Methods

### NewOrderAmountPromotionRuleData

`func NewOrderAmountPromotionRuleData(type_ string, attributes OrderAmountPromotionRuleDataAttributes, ) *OrderAmountPromotionRuleData`

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

`func (o *OrderAmountPromotionRuleData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderAmountPromotionRuleData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderAmountPromotionRuleData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OrderAmountPromotionRuleData) GetAttributes() OrderAmountPromotionRuleDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderAmountPromotionRuleData) GetAttributesOk() (*OrderAmountPromotionRuleDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderAmountPromotionRuleData) SetAttributes(v OrderAmountPromotionRuleDataAttributes)`

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


