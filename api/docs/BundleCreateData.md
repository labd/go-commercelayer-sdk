# BundleCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTBundles201ResponseDataAttributes**](POSTBundles201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BundleCreateDataRelationships**](BundleCreateDataRelationships.md) |  | [optional] 

## Methods

### NewBundleCreateData

`func NewBundleCreateData(type_ interface{}, attributes POSTBundles201ResponseDataAttributes, ) *BundleCreateData`

NewBundleCreateData instantiates a new BundleCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleCreateDataWithDefaults

`func NewBundleCreateDataWithDefaults() *BundleCreateData`

NewBundleCreateDataWithDefaults instantiates a new BundleCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BundleCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BundleCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BundleCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *BundleCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BundleCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *BundleCreateData) GetAttributes() POSTBundles201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BundleCreateData) GetAttributesOk() (*POSTBundles201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BundleCreateData) SetAttributes(v POSTBundles201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BundleCreateData) GetRelationships() BundleCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BundleCreateData) GetRelationshipsOk() (*BundleCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BundleCreateData) SetRelationships(v BundleCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BundleCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


