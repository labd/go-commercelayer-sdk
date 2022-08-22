# SkuListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "sku_lists"]
**Attributes** | [**GETSkuLists200ResponseDataInnerAttributes**](GETSkuLists200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**GETSkuLists200ResponseDataInnerRelationships**](GETSkuLists200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewSkuListData

`func NewSkuListData(type_ string, attributes GETSkuLists200ResponseDataInnerAttributes, ) *SkuListData`

NewSkuListData instantiates a new SkuListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuListDataWithDefaults

`func NewSkuListDataWithDefaults() *SkuListData`

NewSkuListDataWithDefaults instantiates a new SkuListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuListData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuListData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuListData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SkuListData) GetAttributes() GETSkuLists200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuListData) GetAttributesOk() (*GETSkuLists200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuListData) SetAttributes(v GETSkuLists200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuListData) GetRelationships() GETSkuLists200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuListData) GetRelationshipsOk() (*GETSkuLists200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuListData) SetRelationships(v GETSkuLists200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuListData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


