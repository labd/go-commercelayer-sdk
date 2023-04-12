# ManualTaxCalculatorUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes**](PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ManualTaxCalculatorCreateDataRelationships**](ManualTaxCalculatorCreateDataRelationships.md) |  | [optional] 

## Methods

### NewManualTaxCalculatorUpdateData

`func NewManualTaxCalculatorUpdateData(type_ interface{}, id interface{}, attributes PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes, ) *ManualTaxCalculatorUpdateData`

NewManualTaxCalculatorUpdateData instantiates a new ManualTaxCalculatorUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualTaxCalculatorUpdateDataWithDefaults

`func NewManualTaxCalculatorUpdateDataWithDefaults() *ManualTaxCalculatorUpdateData`

NewManualTaxCalculatorUpdateDataWithDefaults instantiates a new ManualTaxCalculatorUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManualTaxCalculatorUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManualTaxCalculatorUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManualTaxCalculatorUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ManualTaxCalculatorUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ManualTaxCalculatorUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *ManualTaxCalculatorUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualTaxCalculatorUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualTaxCalculatorUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ManualTaxCalculatorUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ManualTaxCalculatorUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *ManualTaxCalculatorUpdateData) GetAttributes() PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManualTaxCalculatorUpdateData) GetAttributesOk() (*PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManualTaxCalculatorUpdateData) SetAttributes(v PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ManualTaxCalculatorUpdateData) GetRelationships() ManualTaxCalculatorCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ManualTaxCalculatorUpdateData) GetRelationshipsOk() (*ManualTaxCalculatorCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ManualTaxCalculatorUpdateData) SetRelationships(v ManualTaxCalculatorCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ManualTaxCalculatorUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


