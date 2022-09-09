# TaxRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "tax_rules"]
**Attributes** | [**GETTaxRules200ResponseDataInnerAttributes**](GETTaxRules200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**TaxRuleDataRelationships**](TaxRuleDataRelationships.md) |  | [optional] 

## Methods

### NewTaxRuleData

`func NewTaxRuleData(type_ string, attributes GETTaxRules200ResponseDataInnerAttributes, ) *TaxRuleData`

NewTaxRuleData instantiates a new TaxRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRuleDataWithDefaults

`func NewTaxRuleDataWithDefaults() *TaxRuleData`

NewTaxRuleDataWithDefaults instantiates a new TaxRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxRuleData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxRuleData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxRuleData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TaxRuleData) GetAttributes() GETTaxRules200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxRuleData) GetAttributesOk() (*GETTaxRules200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxRuleData) SetAttributes(v GETTaxRules200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxRuleData) GetRelationships() TaxRuleDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxRuleData) GetRelationshipsOk() (*TaxRuleDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxRuleData) SetRelationships(v TaxRuleDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxRuleData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


