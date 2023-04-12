# BundleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETBundlesBundleId200ResponseDataAttributes**](GETBundlesBundleId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BundleDataRelationships**](BundleDataRelationships.md) |  | [optional] 

## Methods

### NewBundleData

`func NewBundleData(type_ interface{}, attributes GETBundlesBundleId200ResponseDataAttributes, ) *BundleData`

NewBundleData instantiates a new BundleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleDataWithDefaults

`func NewBundleDataWithDefaults() *BundleData`

NewBundleDataWithDefaults instantiates a new BundleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BundleData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BundleData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BundleData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *BundleData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BundleData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *BundleData) GetAttributes() GETBundlesBundleId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BundleData) GetAttributesOk() (*GETBundlesBundleId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BundleData) SetAttributes(v GETBundlesBundleId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BundleData) GetRelationships() BundleDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BundleData) GetRelationshipsOk() (*BundleDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BundleData) SetRelationships(v BundleDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BundleData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


