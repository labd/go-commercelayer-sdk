# SkuListPromotionRuleUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**SkuListPromotionRuleCreateDataAttributes**](SkuListPromotionRuleCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**SkuListPromotionRuleUpdateDataRelationships**](SkuListPromotionRuleUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewSkuListPromotionRuleUpdateData

`func NewSkuListPromotionRuleUpdateData(type_ string, id string, attributes SkuListPromotionRuleCreateDataAttributes, ) *SkuListPromotionRuleUpdateData`

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

`func (o *SkuListPromotionRuleUpdateData) GetAttributes() SkuListPromotionRuleCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuListPromotionRuleUpdateData) GetAttributesOk() (*SkuListPromotionRuleCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuListPromotionRuleUpdateData) SetAttributes(v SkuListPromotionRuleCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuListPromotionRuleUpdateData) GetRelationships() SkuListPromotionRuleUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuListPromotionRuleUpdateData) GetRelationshipsOk() (*SkuListPromotionRuleUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuListPromotionRuleUpdateData) SetRelationships(v SkuListPromotionRuleUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuListPromotionRuleUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


