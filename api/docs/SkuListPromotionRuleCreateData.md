# SkuListPromotionRuleCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "sku_list_promotion_rules"]
**Attributes** | [**SkuListPromotionRuleCreateDataAttributes**](SkuListPromotionRuleCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**SkuListPromotionRuleCreateDataRelationships**](SkuListPromotionRuleCreateDataRelationships.md) |  | [optional] 

## Methods

### NewSkuListPromotionRuleCreateData

`func NewSkuListPromotionRuleCreateData(type_ string, attributes SkuListPromotionRuleCreateDataAttributes, ) *SkuListPromotionRuleCreateData`

NewSkuListPromotionRuleCreateData instantiates a new SkuListPromotionRuleCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuListPromotionRuleCreateDataWithDefaults

`func NewSkuListPromotionRuleCreateDataWithDefaults() *SkuListPromotionRuleCreateData`

NewSkuListPromotionRuleCreateDataWithDefaults instantiates a new SkuListPromotionRuleCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuListPromotionRuleCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuListPromotionRuleCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuListPromotionRuleCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SkuListPromotionRuleCreateData) GetAttributes() SkuListPromotionRuleCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuListPromotionRuleCreateData) GetAttributesOk() (*SkuListPromotionRuleCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuListPromotionRuleCreateData) SetAttributes(v SkuListPromotionRuleCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuListPromotionRuleCreateData) GetRelationships() SkuListPromotionRuleCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuListPromotionRuleCreateData) GetRelationshipsOk() (*SkuListPromotionRuleCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuListPromotionRuleCreateData) SetRelationships(v SkuListPromotionRuleCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuListPromotionRuleCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


