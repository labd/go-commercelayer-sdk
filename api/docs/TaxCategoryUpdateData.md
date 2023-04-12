# TaxCategoryUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes**](PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**TaxCategoryUpdateDataRelationships**](TaxCategoryUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewTaxCategoryUpdateData

`func NewTaxCategoryUpdateData(type_ interface{}, id interface{}, attributes PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes, ) *TaxCategoryUpdateData`

NewTaxCategoryUpdateData instantiates a new TaxCategoryUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCategoryUpdateDataWithDefaults

`func NewTaxCategoryUpdateDataWithDefaults() *TaxCategoryUpdateData`

NewTaxCategoryUpdateDataWithDefaults instantiates a new TaxCategoryUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxCategoryUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxCategoryUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxCategoryUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *TaxCategoryUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TaxCategoryUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *TaxCategoryUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxCategoryUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxCategoryUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *TaxCategoryUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaxCategoryUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *TaxCategoryUpdateData) GetAttributes() PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxCategoryUpdateData) GetAttributesOk() (*PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxCategoryUpdateData) SetAttributes(v PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxCategoryUpdateData) GetRelationships() TaxCategoryUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxCategoryUpdateData) GetRelationshipsOk() (*TaxCategoryUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxCategoryUpdateData) SetRelationships(v TaxCategoryUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxCategoryUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


