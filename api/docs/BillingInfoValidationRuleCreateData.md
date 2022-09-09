# BillingInfoValidationRuleCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "billing_info_validation_rules"]
**Attributes** | [**POSTAdyenPayments201ResponseDataAttributes**](POSTAdyenPayments201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BillingInfoValidationRuleCreateDataRelationships**](BillingInfoValidationRuleCreateDataRelationships.md) |  | [optional] 

## Methods

### NewBillingInfoValidationRuleCreateData

`func NewBillingInfoValidationRuleCreateData(type_ string, attributes POSTAdyenPayments201ResponseDataAttributes, ) *BillingInfoValidationRuleCreateData`

NewBillingInfoValidationRuleCreateData instantiates a new BillingInfoValidationRuleCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInfoValidationRuleCreateDataWithDefaults

`func NewBillingInfoValidationRuleCreateDataWithDefaults() *BillingInfoValidationRuleCreateData`

NewBillingInfoValidationRuleCreateDataWithDefaults instantiates a new BillingInfoValidationRuleCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BillingInfoValidationRuleCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingInfoValidationRuleCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingInfoValidationRuleCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BillingInfoValidationRuleCreateData) GetAttributes() POSTAdyenPayments201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BillingInfoValidationRuleCreateData) GetAttributesOk() (*POSTAdyenPayments201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BillingInfoValidationRuleCreateData) SetAttributes(v POSTAdyenPayments201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BillingInfoValidationRuleCreateData) GetRelationships() BillingInfoValidationRuleCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BillingInfoValidationRuleCreateData) GetRelationshipsOk() (*BillingInfoValidationRuleCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BillingInfoValidationRuleCreateData) SetRelationships(v BillingInfoValidationRuleCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BillingInfoValidationRuleCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


