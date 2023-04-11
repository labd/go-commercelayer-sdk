# POSTSkuOptionsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTSkuOptionsRequestDataAttributes**](POSTSkuOptionsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships**](PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTSkuOptionsRequestData

`func NewPOSTSkuOptionsRequestData(type_ interface{}, attributes POSTSkuOptionsRequestDataAttributes, ) *POSTSkuOptionsRequestData`

NewPOSTSkuOptionsRequestData instantiates a new POSTSkuOptionsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkuOptionsRequestDataWithDefaults

`func NewPOSTSkuOptionsRequestDataWithDefaults() *POSTSkuOptionsRequestData`

NewPOSTSkuOptionsRequestDataWithDefaults instantiates a new POSTSkuOptionsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTSkuOptionsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTSkuOptionsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTSkuOptionsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTSkuOptionsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTSkuOptionsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTSkuOptionsRequestData) GetAttributes() POSTSkuOptionsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTSkuOptionsRequestData) GetAttributesOk() (*POSTSkuOptionsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTSkuOptionsRequestData) SetAttributes(v POSTSkuOptionsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTSkuOptionsRequestData) GetRelationships() PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTSkuOptionsRequestData) GetRelationshipsOk() (*PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTSkuOptionsRequestData) SetRelationships(v PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTSkuOptionsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


