# POSTTaxjarAccountsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTTaxjarAccountsRequestDataAttributes**](POSTTaxjarAccountsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTAvalaraAccountsRequestDataRelationships**](POSTAvalaraAccountsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTTaxjarAccountsRequestData

`func NewPOSTTaxjarAccountsRequestData(type_ interface{}, attributes POSTTaxjarAccountsRequestDataAttributes, ) *POSTTaxjarAccountsRequestData`

NewPOSTTaxjarAccountsRequestData instantiates a new POSTTaxjarAccountsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTTaxjarAccountsRequestDataWithDefaults

`func NewPOSTTaxjarAccountsRequestDataWithDefaults() *POSTTaxjarAccountsRequestData`

NewPOSTTaxjarAccountsRequestDataWithDefaults instantiates a new POSTTaxjarAccountsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTTaxjarAccountsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTTaxjarAccountsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTTaxjarAccountsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTTaxjarAccountsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTTaxjarAccountsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTTaxjarAccountsRequestData) GetAttributes() POSTTaxjarAccountsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTTaxjarAccountsRequestData) GetAttributesOk() (*POSTTaxjarAccountsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTTaxjarAccountsRequestData) SetAttributes(v POSTTaxjarAccountsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTTaxjarAccountsRequestData) GetRelationships() POSTAvalaraAccountsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTTaxjarAccountsRequestData) GetRelationshipsOk() (*POSTAvalaraAccountsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTTaxjarAccountsRequestData) SetRelationships(v POSTAvalaraAccountsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTTaxjarAccountsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


