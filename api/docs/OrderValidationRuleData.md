# OrderValidationRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes**](GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BillingInfoValidationRuleDataRelationships**](BillingInfoValidationRuleDataRelationships.md) |  | [optional] 

## Methods

### NewOrderValidationRuleData

`func NewOrderValidationRuleData(type_ interface{}, attributes GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes, ) *OrderValidationRuleData`

NewOrderValidationRuleData instantiates a new OrderValidationRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderValidationRuleDataWithDefaults

`func NewOrderValidationRuleDataWithDefaults() *OrderValidationRuleData`

NewOrderValidationRuleDataWithDefaults instantiates a new OrderValidationRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderValidationRuleData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderValidationRuleData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderValidationRuleData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OrderValidationRuleData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderValidationRuleData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *OrderValidationRuleData) GetAttributes() GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderValidationRuleData) GetAttributesOk() (*GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderValidationRuleData) SetAttributes(v GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderValidationRuleData) GetRelationships() BillingInfoValidationRuleDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderValidationRuleData) GetRelationshipsOk() (*BillingInfoValidationRuleDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderValidationRuleData) SetRelationships(v BillingInfoValidationRuleDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderValidationRuleData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


