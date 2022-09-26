# GETPackages200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name for the package | [optional] 
**Code** | Pointer to **string** | The package identifying code | [optional] 
**Length** | Pointer to **float32** | The package length, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**Width** | Pointer to **float32** | The package width, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**Height** | Pointer to **float32** | The package height, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**UnitOfLength** | Pointer to **string** | The unit of length. Can be one of &#39;cm&#39;, or &#39;in&#39;. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETPackages200ResponseDataInnerAttributes

`func NewGETPackages200ResponseDataInnerAttributes() *GETPackages200ResponseDataInnerAttributes`

NewGETPackages200ResponseDataInnerAttributes instantiates a new GETPackages200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPackages200ResponseDataInnerAttributesWithDefaults

`func NewGETPackages200ResponseDataInnerAttributesWithDefaults() *GETPackages200ResponseDataInnerAttributes`

NewGETPackages200ResponseDataInnerAttributesWithDefaults instantiates a new GETPackages200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETPackages200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPackages200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETPackages200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GETPackages200ResponseDataInnerAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GETPackages200ResponseDataInnerAttributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GETPackages200ResponseDataInnerAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLength

`func (o *GETPackages200ResponseDataInnerAttributes) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *GETPackages200ResponseDataInnerAttributes) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *GETPackages200ResponseDataInnerAttributes) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWidth

`func (o *GETPackages200ResponseDataInnerAttributes) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *GETPackages200ResponseDataInnerAttributes) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *GETPackages200ResponseDataInnerAttributes) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *GETPackages200ResponseDataInnerAttributes) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *GETPackages200ResponseDataInnerAttributes) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *GETPackages200ResponseDataInnerAttributes) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetUnitOfLength

`func (o *GETPackages200ResponseDataInnerAttributes) GetUnitOfLength() string`

GetUnitOfLength returns the UnitOfLength field if non-nil, zero value otherwise.

### GetUnitOfLengthOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetUnitOfLengthOk() (*string, bool)`

GetUnitOfLengthOk returns a tuple with the UnitOfLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfLength

`func (o *GETPackages200ResponseDataInnerAttributes) SetUnitOfLength(v string)`

SetUnitOfLength sets UnitOfLength field to given value.

### HasUnitOfLength

`func (o *GETPackages200ResponseDataInnerAttributes) HasUnitOfLength() bool`

HasUnitOfLength returns a boolean if a field has been set.

### GetId

`func (o *GETPackages200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETPackages200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETPackages200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETPackages200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPackages200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPackages200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETPackages200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPackages200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPackages200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETPackages200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPackages200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPackages200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


