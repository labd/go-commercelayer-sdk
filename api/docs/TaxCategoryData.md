# TaxCategoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETTaxCategoriesTaxCategoryId200ResponseDataAttributes**](GETTaxCategoriesTaxCategoryId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**TaxCategoryDataRelationships**](TaxCategoryDataRelationships.md) |  | [optional] 

## Methods

### NewTaxCategoryData

`func NewTaxCategoryData(type_ interface{}, attributes GETTaxCategoriesTaxCategoryId200ResponseDataAttributes, ) *TaxCategoryData`

NewTaxCategoryData instantiates a new TaxCategoryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCategoryDataWithDefaults

`func NewTaxCategoryDataWithDefaults() *TaxCategoryData`

NewTaxCategoryDataWithDefaults instantiates a new TaxCategoryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxCategoryData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxCategoryData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxCategoryData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *TaxCategoryData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TaxCategoryData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *TaxCategoryData) GetAttributes() GETTaxCategoriesTaxCategoryId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxCategoryData) GetAttributesOk() (*GETTaxCategoriesTaxCategoryId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxCategoryData) SetAttributes(v GETTaxCategoriesTaxCategoryId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxCategoryData) GetRelationships() TaxCategoryDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxCategoryData) GetRelationshipsOk() (*TaxCategoryDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxCategoryData) SetRelationships(v TaxCategoryDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxCategoryData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


