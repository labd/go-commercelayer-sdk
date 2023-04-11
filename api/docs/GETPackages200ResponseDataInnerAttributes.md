# GETPackages200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | Unique name for the package | [optional] 
**Code** | Pointer to **interface{}** | The package identifying code | [optional] 
**Length** | Pointer to **interface{}** | The package length, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**Width** | Pointer to **interface{}** | The package width, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**Height** | Pointer to **interface{}** | The package height, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**UnitOfLength** | Pointer to **interface{}** | The unit of length. Can be one of &#39;cm&#39;, or &#39;in&#39;. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETPackages200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPackages200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETPackages200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETPackages200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETPackages200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *GETPackages200ResponseDataInnerAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GETPackages200ResponseDataInnerAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *GETPackages200ResponseDataInnerAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GETPackages200ResponseDataInnerAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GETPackages200ResponseDataInnerAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetLength

`func (o *GETPackages200ResponseDataInnerAttributes) GetLength() interface{}`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetLengthOk() (*interface{}, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *GETPackages200ResponseDataInnerAttributes) SetLength(v interface{})`

SetLength sets Length field to given value.

### HasLength

`func (o *GETPackages200ResponseDataInnerAttributes) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *GETPackages200ResponseDataInnerAttributes) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *GETPackages200ResponseDataInnerAttributes) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetWidth

`func (o *GETPackages200ResponseDataInnerAttributes) GetWidth() interface{}`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetWidthOk() (*interface{}, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *GETPackages200ResponseDataInnerAttributes) SetWidth(v interface{})`

SetWidth sets Width field to given value.

### HasWidth

`func (o *GETPackages200ResponseDataInnerAttributes) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *GETPackages200ResponseDataInnerAttributes) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *GETPackages200ResponseDataInnerAttributes) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *GETPackages200ResponseDataInnerAttributes) GetHeight() interface{}`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetHeightOk() (*interface{}, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *GETPackages200ResponseDataInnerAttributes) SetHeight(v interface{})`

SetHeight sets Height field to given value.

### HasHeight

`func (o *GETPackages200ResponseDataInnerAttributes) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *GETPackages200ResponseDataInnerAttributes) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *GETPackages200ResponseDataInnerAttributes) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetUnitOfLength

`func (o *GETPackages200ResponseDataInnerAttributes) GetUnitOfLength() interface{}`

GetUnitOfLength returns the UnitOfLength field if non-nil, zero value otherwise.

### GetUnitOfLengthOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetUnitOfLengthOk() (*interface{}, bool)`

GetUnitOfLengthOk returns a tuple with the UnitOfLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfLength

`func (o *GETPackages200ResponseDataInnerAttributes) SetUnitOfLength(v interface{})`

SetUnitOfLength sets UnitOfLength field to given value.

### HasUnitOfLength

`func (o *GETPackages200ResponseDataInnerAttributes) HasUnitOfLength() bool`

HasUnitOfLength returns a boolean if a field has been set.

### SetUnitOfLengthNil

`func (o *GETPackages200ResponseDataInnerAttributes) SetUnitOfLengthNil(b bool)`

 SetUnitOfLengthNil sets the value for UnitOfLength to be an explicit nil

### UnsetUnitOfLength
`func (o *GETPackages200ResponseDataInnerAttributes) UnsetUnitOfLength()`

UnsetUnitOfLength ensures that no value is present for UnitOfLength, not even an explicit nil
### GetCreatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPackages200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPackages200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPackages200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPackages200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPackages200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPackages200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPackages200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPackages200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPackages200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPackages200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPackages200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPackages200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPackages200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPackages200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPackages200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETPackages200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPackages200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPackages200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPackages200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPackages200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPackages200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


