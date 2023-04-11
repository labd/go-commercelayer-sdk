# POSTAvalaraAccountsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTAvalaraAccountsRequestDataAttributes**](POSTAvalaraAccountsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTAvalaraAccountsRequestDataRelationships**](POSTAvalaraAccountsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTAvalaraAccountsRequestData

`func NewPOSTAvalaraAccountsRequestData(type_ interface{}, attributes POSTAvalaraAccountsRequestDataAttributes, ) *POSTAvalaraAccountsRequestData`

NewPOSTAvalaraAccountsRequestData instantiates a new POSTAvalaraAccountsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTAvalaraAccountsRequestDataWithDefaults

`func NewPOSTAvalaraAccountsRequestDataWithDefaults() *POSTAvalaraAccountsRequestData`

NewPOSTAvalaraAccountsRequestDataWithDefaults instantiates a new POSTAvalaraAccountsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTAvalaraAccountsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTAvalaraAccountsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTAvalaraAccountsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTAvalaraAccountsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTAvalaraAccountsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTAvalaraAccountsRequestData) GetAttributes() POSTAvalaraAccountsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTAvalaraAccountsRequestData) GetAttributesOk() (*POSTAvalaraAccountsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTAvalaraAccountsRequestData) SetAttributes(v POSTAvalaraAccountsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTAvalaraAccountsRequestData) GetRelationships() POSTAvalaraAccountsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTAvalaraAccountsRequestData) GetRelationshipsOk() (*POSTAvalaraAccountsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTAvalaraAccountsRequestData) SetRelationships(v POSTAvalaraAccountsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTAvalaraAccountsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


