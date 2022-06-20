# BillingInfoValidationRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "billing_info_validation_rules"]
**Attributes** | [**BillingInfoValidationRuleDataAttributes**](BillingInfoValidationRuleDataAttributes.md) |  | 
**Relationships** | Pointer to [**BillingInfoValidationRuleDataRelationships**](BillingInfoValidationRuleDataRelationships.md) |  | [optional] 

## Methods

### NewBillingInfoValidationRuleData

`func NewBillingInfoValidationRuleData(type_ string, attributes BillingInfoValidationRuleDataAttributes, ) *BillingInfoValidationRuleData`

NewBillingInfoValidationRuleData instantiates a new BillingInfoValidationRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInfoValidationRuleDataWithDefaults

`func NewBillingInfoValidationRuleDataWithDefaults() *BillingInfoValidationRuleData`

NewBillingInfoValidationRuleDataWithDefaults instantiates a new BillingInfoValidationRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BillingInfoValidationRuleData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingInfoValidationRuleData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingInfoValidationRuleData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BillingInfoValidationRuleData) GetAttributes() BillingInfoValidationRuleDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BillingInfoValidationRuleData) GetAttributesOk() (*BillingInfoValidationRuleDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BillingInfoValidationRuleData) SetAttributes(v BillingInfoValidationRuleDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BillingInfoValidationRuleData) GetRelationships() BillingInfoValidationRuleDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BillingInfoValidationRuleData) GetRelationshipsOk() (*BillingInfoValidationRuleDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BillingInfoValidationRuleData) SetRelationships(v BillingInfoValidationRuleDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BillingInfoValidationRuleData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


