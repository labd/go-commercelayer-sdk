# VoidData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "voids"]
**Attributes** | [**GETRefunds200ResponseDataInnerAttributes**](GETRefunds200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**VoidDataRelationships**](VoidDataRelationships.md) |  | [optional] 

## Methods

### NewVoidData

`func NewVoidData(type_ string, attributes GETRefunds200ResponseDataInnerAttributes, ) *VoidData`

NewVoidData instantiates a new VoidData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidDataWithDefaults

`func NewVoidDataWithDefaults() *VoidData`

NewVoidDataWithDefaults instantiates a new VoidData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VoidData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VoidData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VoidData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *VoidData) GetAttributes() GETRefunds200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VoidData) GetAttributesOk() (*GETRefunds200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VoidData) SetAttributes(v GETRefunds200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *VoidData) GetRelationships() VoidDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *VoidData) GetRelationshipsOk() (*VoidDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *VoidData) SetRelationships(v VoidDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *VoidData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


