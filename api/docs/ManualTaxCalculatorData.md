# ManualTaxCalculatorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**ManualTaxCalculatorDataAttributes**](ManualTaxCalculatorDataAttributes.md) |  | 
**Relationships** | Pointer to [**ManualTaxCalculatorDataRelationships**](ManualTaxCalculatorDataRelationships.md) |  | [optional] 

## Methods

### NewManualTaxCalculatorData

`func NewManualTaxCalculatorData(type_ string, attributes ManualTaxCalculatorDataAttributes, ) *ManualTaxCalculatorData`

NewManualTaxCalculatorData instantiates a new ManualTaxCalculatorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualTaxCalculatorDataWithDefaults

`func NewManualTaxCalculatorDataWithDefaults() *ManualTaxCalculatorData`

NewManualTaxCalculatorDataWithDefaults instantiates a new ManualTaxCalculatorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManualTaxCalculatorData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManualTaxCalculatorData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManualTaxCalculatorData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ManualTaxCalculatorData) GetAttributes() ManualTaxCalculatorDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManualTaxCalculatorData) GetAttributesOk() (*ManualTaxCalculatorDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManualTaxCalculatorData) SetAttributes(v ManualTaxCalculatorDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ManualTaxCalculatorData) GetRelationships() ManualTaxCalculatorDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ManualTaxCalculatorData) GetRelationshipsOk() (*ManualTaxCalculatorDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ManualTaxCalculatorData) SetRelationships(v ManualTaxCalculatorDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ManualTaxCalculatorData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


