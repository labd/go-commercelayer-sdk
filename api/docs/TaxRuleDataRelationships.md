# TaxRuleDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManualTaxCalculator** | Pointer to [**TaxRuleDataRelationshipsManualTaxCalculator**](TaxRuleDataRelationshipsManualTaxCalculator.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewTaxRuleDataRelationships

`func NewTaxRuleDataRelationships() *TaxRuleDataRelationships`

NewTaxRuleDataRelationships instantiates a new TaxRuleDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRuleDataRelationshipsWithDefaults

`func NewTaxRuleDataRelationshipsWithDefaults() *TaxRuleDataRelationships`

NewTaxRuleDataRelationshipsWithDefaults instantiates a new TaxRuleDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManualTaxCalculator

`func (o *TaxRuleDataRelationships) GetManualTaxCalculator() TaxRuleDataRelationshipsManualTaxCalculator`

GetManualTaxCalculator returns the ManualTaxCalculator field if non-nil, zero value otherwise.

### GetManualTaxCalculatorOk

`func (o *TaxRuleDataRelationships) GetManualTaxCalculatorOk() (*TaxRuleDataRelationshipsManualTaxCalculator, bool)`

GetManualTaxCalculatorOk returns a tuple with the ManualTaxCalculator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualTaxCalculator

`func (o *TaxRuleDataRelationships) SetManualTaxCalculator(v TaxRuleDataRelationshipsManualTaxCalculator)`

SetManualTaxCalculator sets ManualTaxCalculator field to given value.

### HasManualTaxCalculator

`func (o *TaxRuleDataRelationships) HasManualTaxCalculator() bool`

HasManualTaxCalculator returns a boolean if a field has been set.

### GetVersions

`func (o *TaxRuleDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *TaxRuleDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *TaxRuleDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *TaxRuleDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


