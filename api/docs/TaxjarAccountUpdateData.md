# TaxjarAccountUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes**](PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AvalaraAccountCreateDataRelationships**](AvalaraAccountCreateDataRelationships.md) |  | [optional] 

## Methods

### NewTaxjarAccountUpdateData

`func NewTaxjarAccountUpdateData(type_ interface{}, id interface{}, attributes PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes, ) *TaxjarAccountUpdateData`

NewTaxjarAccountUpdateData instantiates a new TaxjarAccountUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxjarAccountUpdateDataWithDefaults

`func NewTaxjarAccountUpdateDataWithDefaults() *TaxjarAccountUpdateData`

NewTaxjarAccountUpdateDataWithDefaults instantiates a new TaxjarAccountUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxjarAccountUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxjarAccountUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxjarAccountUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *TaxjarAccountUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TaxjarAccountUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *TaxjarAccountUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxjarAccountUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxjarAccountUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *TaxjarAccountUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaxjarAccountUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *TaxjarAccountUpdateData) GetAttributes() PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxjarAccountUpdateData) GetAttributesOk() (*PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxjarAccountUpdateData) SetAttributes(v PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxjarAccountUpdateData) GetRelationships() AvalaraAccountCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxjarAccountUpdateData) GetRelationshipsOk() (*AvalaraAccountCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxjarAccountUpdateData) SetRelationships(v AvalaraAccountCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxjarAccountUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


