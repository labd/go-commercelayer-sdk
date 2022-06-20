# ExternalTaxCalculatorUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "external_tax_calculators"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**ExternalTaxCalculatorUpdateDataAttributes**](ExternalTaxCalculatorUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**AvalaraAccountCreateDataRelationships**](AvalaraAccountCreateDataRelationships.md) |  | [optional] 

## Methods

### NewExternalTaxCalculatorUpdateData

`func NewExternalTaxCalculatorUpdateData(type_ string, id string, attributes ExternalTaxCalculatorUpdateDataAttributes, ) *ExternalTaxCalculatorUpdateData`

NewExternalTaxCalculatorUpdateData instantiates a new ExternalTaxCalculatorUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalTaxCalculatorUpdateDataWithDefaults

`func NewExternalTaxCalculatorUpdateDataWithDefaults() *ExternalTaxCalculatorUpdateData`

NewExternalTaxCalculatorUpdateDataWithDefaults instantiates a new ExternalTaxCalculatorUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalTaxCalculatorUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalTaxCalculatorUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalTaxCalculatorUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ExternalTaxCalculatorUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalTaxCalculatorUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalTaxCalculatorUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *ExternalTaxCalculatorUpdateData) GetAttributes() ExternalTaxCalculatorUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalTaxCalculatorUpdateData) GetAttributesOk() (*ExternalTaxCalculatorUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalTaxCalculatorUpdateData) SetAttributes(v ExternalTaxCalculatorUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExternalTaxCalculatorUpdateData) GetRelationships() AvalaraAccountCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalTaxCalculatorUpdateData) GetRelationshipsOk() (*AvalaraAccountCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalTaxCalculatorUpdateData) SetRelationships(v AvalaraAccountCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalTaxCalculatorUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


