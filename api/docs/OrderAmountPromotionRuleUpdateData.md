# OrderAmountPromotionRuleUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**OrderAmountPromotionRuleCreateDataAttributes**](OrderAmountPromotionRuleCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderAmountPromotionRuleDataRelationships**](OrderAmountPromotionRuleDataRelationships.md) |  | [optional] 

## Methods

### NewOrderAmountPromotionRuleUpdateData

`func NewOrderAmountPromotionRuleUpdateData(type_ string, id string, attributes OrderAmountPromotionRuleCreateDataAttributes, ) *OrderAmountPromotionRuleUpdateData`

NewOrderAmountPromotionRuleUpdateData instantiates a new OrderAmountPromotionRuleUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmountPromotionRuleUpdateDataWithDefaults

`func NewOrderAmountPromotionRuleUpdateDataWithDefaults() *OrderAmountPromotionRuleUpdateData`

NewOrderAmountPromotionRuleUpdateDataWithDefaults instantiates a new OrderAmountPromotionRuleUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderAmountPromotionRuleUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderAmountPromotionRuleUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderAmountPromotionRuleUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *OrderAmountPromotionRuleUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderAmountPromotionRuleUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderAmountPromotionRuleUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *OrderAmountPromotionRuleUpdateData) GetAttributes() OrderAmountPromotionRuleCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderAmountPromotionRuleUpdateData) GetAttributesOk() (*OrderAmountPromotionRuleCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderAmountPromotionRuleUpdateData) SetAttributes(v OrderAmountPromotionRuleCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderAmountPromotionRuleUpdateData) GetRelationships() OrderAmountPromotionRuleDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderAmountPromotionRuleUpdateData) GetRelationshipsOk() (*OrderAmountPromotionRuleDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderAmountPromotionRuleUpdateData) SetRelationships(v OrderAmountPromotionRuleDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderAmountPromotionRuleUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


