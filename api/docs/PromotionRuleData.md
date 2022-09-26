# PromotionRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**BillingInfoValidationRuleDataAttributes**](BillingInfoValidationRuleDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderAmountPromotionRuleDataRelationships**](OrderAmountPromotionRuleDataRelationships.md) |  | [optional] 

## Methods

### NewPromotionRuleData

`func NewPromotionRuleData(type_ string, attributes BillingInfoValidationRuleDataAttributes, ) *PromotionRuleData`

NewPromotionRuleData instantiates a new PromotionRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionRuleDataWithDefaults

`func NewPromotionRuleDataWithDefaults() *PromotionRuleData`

NewPromotionRuleDataWithDefaults instantiates a new PromotionRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PromotionRuleData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromotionRuleData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromotionRuleData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PromotionRuleData) GetAttributes() BillingInfoValidationRuleDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PromotionRuleData) GetAttributesOk() (*BillingInfoValidationRuleDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PromotionRuleData) SetAttributes(v BillingInfoValidationRuleDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PromotionRuleData) GetRelationships() OrderAmountPromotionRuleDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PromotionRuleData) GetRelationshipsOk() (*OrderAmountPromotionRuleDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PromotionRuleData) SetRelationships(v OrderAmountPromotionRuleDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PromotionRuleData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


