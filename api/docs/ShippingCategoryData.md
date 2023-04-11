# ShippingCategoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETShippingCategories200ResponseDataInnerAttributes**](GETShippingCategories200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**ShippingCategoryDataRelationships**](ShippingCategoryDataRelationships.md) |  | [optional] 

## Methods

### NewShippingCategoryData

`func NewShippingCategoryData(type_ interface{}, attributes GETShippingCategories200ResponseDataInnerAttributes, ) *ShippingCategoryData`

NewShippingCategoryData instantiates a new ShippingCategoryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingCategoryDataWithDefaults

`func NewShippingCategoryDataWithDefaults() *ShippingCategoryData`

NewShippingCategoryDataWithDefaults instantiates a new ShippingCategoryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShippingCategoryData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingCategoryData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingCategoryData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ShippingCategoryData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ShippingCategoryData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ShippingCategoryData) GetAttributes() GETShippingCategories200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShippingCategoryData) GetAttributesOk() (*GETShippingCategories200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShippingCategoryData) SetAttributes(v GETShippingCategories200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ShippingCategoryData) GetRelationships() ShippingCategoryDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShippingCategoryData) GetRelationshipsOk() (*ShippingCategoryDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShippingCategoryData) SetRelationships(v ShippingCategoryDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShippingCategoryData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


