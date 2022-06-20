# BundleCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "bundles"]
**Attributes** | [**BundleCreateDataAttributes**](BundleCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**BundleCreateDataRelationships**](BundleCreateDataRelationships.md) |  | [optional] 

## Methods

### NewBundleCreateData

`func NewBundleCreateData(type_ string, attributes BundleCreateDataAttributes, ) *BundleCreateData`

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

`func (o *BundleCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BundleCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BundleCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BundleCreateData) GetAttributes() BundleCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BundleCreateData) GetAttributesOk() (*BundleCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BundleCreateData) SetAttributes(v BundleCreateDataAttributes)`

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


