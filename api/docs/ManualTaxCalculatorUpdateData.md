# ManualTaxCalculatorUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "manual_tax_calculators"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes**](PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTManualTaxCalculators201ResponseDataRelationships**](POSTManualTaxCalculators201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewManualTaxCalculatorUpdateData

`func NewManualTaxCalculatorUpdateData(type_ string, id string, attributes PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes, ) *ManualTaxCalculatorUpdateData`

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

`func (o *ManualTaxCalculatorUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManualTaxCalculatorUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManualTaxCalculatorUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ManualTaxCalculatorUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualTaxCalculatorUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualTaxCalculatorUpdateData) SetId(v string)`

SetId sets Id field to given value.


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

`func (o *ManualTaxCalculatorUpdateData) GetRelationships() POSTManualTaxCalculators201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ManualTaxCalculatorUpdateData) GetRelationshipsOk() (*POSTManualTaxCalculators201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ManualTaxCalculatorUpdateData) SetRelationships(v POSTManualTaxCalculators201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ManualTaxCalculatorUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


