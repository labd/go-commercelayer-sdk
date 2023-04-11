# POSTTaxRulesRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTTaxRulesRequestDataAttributes**](POSTTaxRulesRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTTaxRulesRequestDataRelationships**](POSTTaxRulesRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTTaxRulesRequestData

`func NewPOSTTaxRulesRequestData(type_ interface{}, attributes POSTTaxRulesRequestDataAttributes, ) *POSTTaxRulesRequestData`

NewPOSTTaxRulesRequestData instantiates a new POSTTaxRulesRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTTaxRulesRequestDataWithDefaults

`func NewPOSTTaxRulesRequestDataWithDefaults() *POSTTaxRulesRequestData`

NewPOSTTaxRulesRequestDataWithDefaults instantiates a new POSTTaxRulesRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTTaxRulesRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTTaxRulesRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTTaxRulesRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTTaxRulesRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTTaxRulesRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTTaxRulesRequestData) GetAttributes() POSTTaxRulesRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTTaxRulesRequestData) GetAttributesOk() (*POSTTaxRulesRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTTaxRulesRequestData) SetAttributes(v POSTTaxRulesRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTTaxRulesRequestData) GetRelationships() POSTTaxRulesRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTTaxRulesRequestData) GetRelationshipsOk() (*POSTTaxRulesRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTTaxRulesRequestData) SetRelationships(v POSTTaxRulesRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTTaxRulesRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


