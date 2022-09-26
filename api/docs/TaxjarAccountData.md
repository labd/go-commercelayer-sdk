# TaxjarAccountData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**ManualTaxCalculatorDataAttributes**](ManualTaxCalculatorDataAttributes.md) |  | 
**Relationships** | Pointer to [**AvalaraAccountDataRelationships**](AvalaraAccountDataRelationships.md) |  | [optional] 

## Methods

### NewTaxjarAccountData

`func NewTaxjarAccountData(type_ string, attributes ManualTaxCalculatorDataAttributes, ) *TaxjarAccountData`

NewTaxjarAccountData instantiates a new TaxjarAccountData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxjarAccountDataWithDefaults

`func NewTaxjarAccountDataWithDefaults() *TaxjarAccountData`

NewTaxjarAccountDataWithDefaults instantiates a new TaxjarAccountData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxjarAccountData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxjarAccountData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxjarAccountData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TaxjarAccountData) GetAttributes() ManualTaxCalculatorDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxjarAccountData) GetAttributesOk() (*ManualTaxCalculatorDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxjarAccountData) SetAttributes(v ManualTaxCalculatorDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxjarAccountData) GetRelationships() AvalaraAccountDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxjarAccountData) GetRelationshipsOk() (*AvalaraAccountDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxjarAccountData) SetRelationships(v AvalaraAccountDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxjarAccountData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


