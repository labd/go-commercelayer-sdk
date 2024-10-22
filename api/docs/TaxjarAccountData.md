# TaxjarAccountData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes**](GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**TaxjarAccountDataRelationships**](TaxjarAccountDataRelationships.md) |  | [optional] 

## Methods

### NewTaxjarAccountData

`func NewTaxjarAccountData(type_ interface{}, attributes GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes, ) *TaxjarAccountData`

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

`func (o *TaxjarAccountData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxjarAccountData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxjarAccountData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *TaxjarAccountData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TaxjarAccountData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *TaxjarAccountData) GetAttributes() GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxjarAccountData) GetAttributesOk() (*GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxjarAccountData) SetAttributes(v GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxjarAccountData) GetRelationships() TaxjarAccountDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxjarAccountData) GetRelationshipsOk() (*TaxjarAccountDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxjarAccountData) SetRelationships(v TaxjarAccountDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxjarAccountData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


