# ExternalTaxCalculatorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes**](GETExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExternalTaxCalculatorDataRelationships**](ExternalTaxCalculatorDataRelationships.md) |  | [optional] 

## Methods

### NewExternalTaxCalculatorData

`func NewExternalTaxCalculatorData(type_ interface{}, attributes GETExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes, ) *ExternalTaxCalculatorData`

NewExternalTaxCalculatorData instantiates a new ExternalTaxCalculatorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalTaxCalculatorDataWithDefaults

`func NewExternalTaxCalculatorDataWithDefaults() *ExternalTaxCalculatorData`

NewExternalTaxCalculatorDataWithDefaults instantiates a new ExternalTaxCalculatorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalTaxCalculatorData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalTaxCalculatorData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalTaxCalculatorData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ExternalTaxCalculatorData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalTaxCalculatorData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ExternalTaxCalculatorData) GetAttributes() GETExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalTaxCalculatorData) GetAttributesOk() (*GETExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalTaxCalculatorData) SetAttributes(v GETExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExternalTaxCalculatorData) GetRelationships() ExternalTaxCalculatorDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalTaxCalculatorData) GetRelationshipsOk() (*ExternalTaxCalculatorDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalTaxCalculatorData) SetRelationships(v ExternalTaxCalculatorDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalTaxCalculatorData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


