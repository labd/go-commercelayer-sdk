# ReturnData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "returns"]
**Attributes** | [**GETReturns200ResponseDataInnerAttributes**](GETReturns200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**ReturnDataRelationships**](ReturnDataRelationships.md) |  | [optional] 

## Methods

### NewReturnData

`func NewReturnData(type_ string, attributes GETReturns200ResponseDataInnerAttributes, ) *ReturnData`

NewReturnData instantiates a new ReturnData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnDataWithDefaults

`func NewReturnDataWithDefaults() *ReturnData`

NewReturnDataWithDefaults instantiates a new ReturnData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReturnData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReturnData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReturnData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ReturnData) GetAttributes() GETReturns200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ReturnData) GetAttributesOk() (*GETReturns200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ReturnData) SetAttributes(v GETReturns200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ReturnData) GetRelationships() ReturnDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ReturnData) GetRelationshipsOk() (*ReturnDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ReturnData) SetRelationships(v ReturnDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ReturnData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


