# TaxRuleUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHTaxRulesTaxRuleId200ResponseDataAttributes**](PATCHTaxRulesTaxRuleId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**TaxRuleUpdateDataRelationships**](TaxRuleUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewTaxRuleUpdateData

`func NewTaxRuleUpdateData(type_ string, id string, attributes PATCHTaxRulesTaxRuleId200ResponseDataAttributes, ) *TaxRuleUpdateData`

NewTaxRuleUpdateData instantiates a new TaxRuleUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRuleUpdateDataWithDefaults

`func NewTaxRuleUpdateDataWithDefaults() *TaxRuleUpdateData`

NewTaxRuleUpdateDataWithDefaults instantiates a new TaxRuleUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxRuleUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxRuleUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxRuleUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *TaxRuleUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxRuleUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxRuleUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *TaxRuleUpdateData) GetAttributes() PATCHTaxRulesTaxRuleId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxRuleUpdateData) GetAttributesOk() (*PATCHTaxRulesTaxRuleId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxRuleUpdateData) SetAttributes(v PATCHTaxRulesTaxRuleId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxRuleUpdateData) GetRelationships() TaxRuleUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxRuleUpdateData) GetRelationshipsOk() (*TaxRuleUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxRuleUpdateData) SetRelationships(v TaxRuleUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxRuleUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


