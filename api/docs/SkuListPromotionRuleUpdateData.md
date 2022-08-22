# SkuListPromotionRuleUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "sku_list_promotion_rules"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**POSTSkuListPromotionRules201ResponseDataAttributes**](POSTSkuListPromotionRules201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataRelationships**](PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewSkuListPromotionRuleUpdateData

`func NewSkuListPromotionRuleUpdateData(type_ string, id string, attributes POSTSkuListPromotionRules201ResponseDataAttributes, ) *SkuListPromotionRuleUpdateData`

NewSkuListPromotionRuleUpdateData instantiates a new SkuListPromotionRuleUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuListPromotionRuleUpdateDataWithDefaults

`func NewSkuListPromotionRuleUpdateDataWithDefaults() *SkuListPromotionRuleUpdateData`

NewSkuListPromotionRuleUpdateDataWithDefaults instantiates a new SkuListPromotionRuleUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuListPromotionRuleUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuListPromotionRuleUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuListPromotionRuleUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *SkuListPromotionRuleUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkuListPromotionRuleUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkuListPromotionRuleUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *SkuListPromotionRuleUpdateData) GetAttributes() POSTSkuListPromotionRules201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuListPromotionRuleUpdateData) GetAttributesOk() (*POSTSkuListPromotionRules201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuListPromotionRuleUpdateData) SetAttributes(v POSTSkuListPromotionRules201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuListPromotionRuleUpdateData) GetRelationships() PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuListPromotionRuleUpdateData) GetRelationshipsOk() (*PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuListPromotionRuleUpdateData) SetRelationships(v PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuListPromotionRuleUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


