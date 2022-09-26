# ManualTaxCalculatorCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**ManualTaxCalculatorCreateDataAttributes**](ManualTaxCalculatorCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**ManualTaxCalculatorCreateDataRelationships**](ManualTaxCalculatorCreateDataRelationships.md) |  | [optional] 

## Methods

### NewManualTaxCalculatorCreateData

`func NewManualTaxCalculatorCreateData(type_ string, attributes ManualTaxCalculatorCreateDataAttributes, ) *ManualTaxCalculatorCreateData`

NewManualTaxCalculatorCreateData instantiates a new ManualTaxCalculatorCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualTaxCalculatorCreateDataWithDefaults

`func NewManualTaxCalculatorCreateDataWithDefaults() *ManualTaxCalculatorCreateData`

NewManualTaxCalculatorCreateDataWithDefaults instantiates a new ManualTaxCalculatorCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManualTaxCalculatorCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManualTaxCalculatorCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManualTaxCalculatorCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ManualTaxCalculatorCreateData) GetAttributes() ManualTaxCalculatorCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManualTaxCalculatorCreateData) GetAttributesOk() (*ManualTaxCalculatorCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManualTaxCalculatorCreateData) SetAttributes(v ManualTaxCalculatorCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ManualTaxCalculatorCreateData) GetRelationships() ManualTaxCalculatorCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ManualTaxCalculatorCreateData) GetRelationshipsOk() (*ManualTaxCalculatorCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ManualTaxCalculatorCreateData) SetRelationships(v ManualTaxCalculatorCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ManualTaxCalculatorCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


