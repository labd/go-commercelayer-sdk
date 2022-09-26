# SkuListItemCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**POSTSkuListItems201ResponseDataAttributes**](POSTSkuListItems201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**SkuListItemCreateDataRelationships**](SkuListItemCreateDataRelationships.md) |  | [optional] 

## Methods

### NewSkuListItemCreateData

`func NewSkuListItemCreateData(type_ string, attributes POSTSkuListItems201ResponseDataAttributes, ) *SkuListItemCreateData`

NewSkuListItemCreateData instantiates a new SkuListItemCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuListItemCreateDataWithDefaults

`func NewSkuListItemCreateDataWithDefaults() *SkuListItemCreateData`

NewSkuListItemCreateDataWithDefaults instantiates a new SkuListItemCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuListItemCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuListItemCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuListItemCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SkuListItemCreateData) GetAttributes() POSTSkuListItems201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuListItemCreateData) GetAttributesOk() (*POSTSkuListItems201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuListItemCreateData) SetAttributes(v POSTSkuListItems201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuListItemCreateData) GetRelationships() SkuListItemCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuListItemCreateData) GetRelationshipsOk() (*SkuListItemCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuListItemCreateData) SetRelationships(v SkuListItemCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuListItemCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


