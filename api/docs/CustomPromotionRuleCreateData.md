# CustomPromotionRuleCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTCustomPromotionRules201ResponseDataAttributes**](POSTCustomPromotionRules201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomPromotionRuleCreateDataRelationships**](CustomPromotionRuleCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCustomPromotionRuleCreateData

`func NewCustomPromotionRuleCreateData(type_ interface{}, attributes POSTCustomPromotionRules201ResponseDataAttributes, ) *CustomPromotionRuleCreateData`

NewCustomPromotionRuleCreateData instantiates a new CustomPromotionRuleCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPromotionRuleCreateDataWithDefaults

`func NewCustomPromotionRuleCreateDataWithDefaults() *CustomPromotionRuleCreateData`

NewCustomPromotionRuleCreateDataWithDefaults instantiates a new CustomPromotionRuleCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomPromotionRuleCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomPromotionRuleCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomPromotionRuleCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CustomPromotionRuleCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomPromotionRuleCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CustomPromotionRuleCreateData) GetAttributes() POSTCustomPromotionRules201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomPromotionRuleCreateData) GetAttributesOk() (*POSTCustomPromotionRules201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomPromotionRuleCreateData) SetAttributes(v POSTCustomPromotionRules201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomPromotionRuleCreateData) GetRelationships() CustomPromotionRuleCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomPromotionRuleCreateData) GetRelationshipsOk() (*CustomPromotionRuleCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomPromotionRuleCreateData) SetRelationships(v CustomPromotionRuleCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomPromotionRuleCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


