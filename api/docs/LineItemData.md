# LineItemData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**GETLineItems200ResponseDataInnerAttributes**](GETLineItems200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**LineItemDataRelationships**](LineItemDataRelationships.md) |  | [optional] 

## Methods

### NewLineItemData

`func NewLineItemData(type_ string, attributes GETLineItems200ResponseDataInnerAttributes, ) *LineItemData`

NewLineItemData instantiates a new LineItemData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemDataWithDefaults

`func NewLineItemDataWithDefaults() *LineItemData`

NewLineItemDataWithDefaults instantiates a new LineItemData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LineItemData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LineItemData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LineItemData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *LineItemData) GetAttributes() GETLineItems200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LineItemData) GetAttributesOk() (*GETLineItems200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LineItemData) SetAttributes(v GETLineItems200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *LineItemData) GetRelationships() LineItemDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LineItemData) GetRelationshipsOk() (*LineItemDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LineItemData) SetRelationships(v LineItemDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *LineItemData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


