# TaxjarAccountCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**TaxjarAccountCreateDataAttributes**](TaxjarAccountCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**AvalaraAccountCreateDataRelationships**](AvalaraAccountCreateDataRelationships.md) |  | [optional] 

## Methods

### NewTaxjarAccountCreateData

`func NewTaxjarAccountCreateData(type_ string, attributes TaxjarAccountCreateDataAttributes, ) *TaxjarAccountCreateData`

NewTaxjarAccountCreateData instantiates a new TaxjarAccountCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxjarAccountCreateDataWithDefaults

`func NewTaxjarAccountCreateDataWithDefaults() *TaxjarAccountCreateData`

NewTaxjarAccountCreateDataWithDefaults instantiates a new TaxjarAccountCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxjarAccountCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxjarAccountCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxjarAccountCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TaxjarAccountCreateData) GetAttributes() TaxjarAccountCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxjarAccountCreateData) GetAttributesOk() (*TaxjarAccountCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxjarAccountCreateData) SetAttributes(v TaxjarAccountCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxjarAccountCreateData) GetRelationships() AvalaraAccountCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxjarAccountCreateData) GetRelationshipsOk() (*AvalaraAccountCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxjarAccountCreateData) SetRelationships(v AvalaraAccountCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxjarAccountCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


