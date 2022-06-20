# ReturnUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "returns"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**ReturnUpdateDataAttributes**](ReturnUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**PackageUpdateDataRelationships**](PackageUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewReturnUpdateData

`func NewReturnUpdateData(type_ string, id string, attributes ReturnUpdateDataAttributes, ) *ReturnUpdateData`

NewReturnUpdateData instantiates a new ReturnUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnUpdateDataWithDefaults

`func NewReturnUpdateDataWithDefaults() *ReturnUpdateData`

NewReturnUpdateDataWithDefaults instantiates a new ReturnUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReturnUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReturnUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReturnUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ReturnUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *ReturnUpdateData) GetAttributes() ReturnUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ReturnUpdateData) GetAttributesOk() (*ReturnUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ReturnUpdateData) SetAttributes(v ReturnUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ReturnUpdateData) GetRelationships() PackageUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ReturnUpdateData) GetRelationshipsOk() (*PackageUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ReturnUpdateData) SetRelationships(v PackageUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ReturnUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


