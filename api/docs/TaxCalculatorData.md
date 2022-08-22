# TaxCalculatorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "tax_calculators"]
**Attributes** | [**GETManualTaxCalculators200ResponseDataInnerAttributes**](GETManualTaxCalculators200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationships**](GETAvalaraAccounts200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewTaxCalculatorData

`func NewTaxCalculatorData(type_ string, attributes GETManualTaxCalculators200ResponseDataInnerAttributes, ) *TaxCalculatorData`

NewTaxCalculatorData instantiates a new TaxCalculatorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCalculatorDataWithDefaults

`func NewTaxCalculatorDataWithDefaults() *TaxCalculatorData`

NewTaxCalculatorDataWithDefaults instantiates a new TaxCalculatorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxCalculatorData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxCalculatorData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxCalculatorData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TaxCalculatorData) GetAttributes() GETManualTaxCalculators200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxCalculatorData) GetAttributesOk() (*GETManualTaxCalculators200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxCalculatorData) SetAttributes(v GETManualTaxCalculators200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxCalculatorData) GetRelationships() GETAvalaraAccounts200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxCalculatorData) GetRelationshipsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxCalculatorData) SetRelationships(v GETAvalaraAccounts200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxCalculatorData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


