# ExternalTaxCalculatorCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "external_tax_calculators"]
**Attributes** | [**POSTExternalTaxCalculators201ResponseDataAttributes**](POSTExternalTaxCalculators201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AvalaraAccountCreateDataRelationships**](AvalaraAccountCreateDataRelationships.md) |  | [optional] 

## Methods

### NewExternalTaxCalculatorCreateData

`func NewExternalTaxCalculatorCreateData(type_ string, attributes POSTExternalTaxCalculators201ResponseDataAttributes, ) *ExternalTaxCalculatorCreateData`

NewExternalTaxCalculatorCreateData instantiates a new ExternalTaxCalculatorCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalTaxCalculatorCreateDataWithDefaults

`func NewExternalTaxCalculatorCreateDataWithDefaults() *ExternalTaxCalculatorCreateData`

NewExternalTaxCalculatorCreateDataWithDefaults instantiates a new ExternalTaxCalculatorCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalTaxCalculatorCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalTaxCalculatorCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalTaxCalculatorCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ExternalTaxCalculatorCreateData) GetAttributes() POSTExternalTaxCalculators201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalTaxCalculatorCreateData) GetAttributesOk() (*POSTExternalTaxCalculators201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalTaxCalculatorCreateData) SetAttributes(v POSTExternalTaxCalculators201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExternalTaxCalculatorCreateData) GetRelationships() AvalaraAccountCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalTaxCalculatorCreateData) GetRelationshipsOk() (*AvalaraAccountCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalTaxCalculatorCreateData) SetRelationships(v AvalaraAccountCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalTaxCalculatorCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


