# OrderAmountPromotionRuleCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**POSTOrderAmountPromotionRules201ResponseDataAttributes**](POSTOrderAmountPromotionRules201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderAmountPromotionRuleCreateDataRelationships**](OrderAmountPromotionRuleCreateDataRelationships.md) |  | [optional] 

## Methods

### NewOrderAmountPromotionRuleCreateData

`func NewOrderAmountPromotionRuleCreateData(type_ string, attributes POSTOrderAmountPromotionRules201ResponseDataAttributes, ) *OrderAmountPromotionRuleCreateData`

NewOrderAmountPromotionRuleCreateData instantiates a new OrderAmountPromotionRuleCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmountPromotionRuleCreateDataWithDefaults

`func NewOrderAmountPromotionRuleCreateDataWithDefaults() *OrderAmountPromotionRuleCreateData`

NewOrderAmountPromotionRuleCreateDataWithDefaults instantiates a new OrderAmountPromotionRuleCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderAmountPromotionRuleCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderAmountPromotionRuleCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderAmountPromotionRuleCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OrderAmountPromotionRuleCreateData) GetAttributes() POSTOrderAmountPromotionRules201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderAmountPromotionRuleCreateData) GetAttributesOk() (*POSTOrderAmountPromotionRules201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderAmountPromotionRuleCreateData) SetAttributes(v POSTOrderAmountPromotionRules201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderAmountPromotionRuleCreateData) GetRelationships() OrderAmountPromotionRuleCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderAmountPromotionRuleCreateData) GetRelationshipsOk() (*OrderAmountPromotionRuleCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderAmountPromotionRuleCreateData) SetRelationships(v OrderAmountPromotionRuleCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderAmountPromotionRuleCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


