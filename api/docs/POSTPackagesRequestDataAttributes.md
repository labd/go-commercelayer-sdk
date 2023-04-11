# POSTPackagesRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | Unique name for the package | 
**Code** | Pointer to **interface{}** | The package identifying code | [optional] 
**Length** | **interface{}** | The package length, used to automatically calculate the tax rates from the available carrier accounts. | 
**Width** | **interface{}** | The package width, used to automatically calculate the tax rates from the available carrier accounts. | 
**Height** | **interface{}** | The package height, used to automatically calculate the tax rates from the available carrier accounts. | 
**UnitOfLength** | **interface{}** | The unit of length. Can be one of &#39;cm&#39;, or &#39;in&#39;. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTPackagesRequestDataAttributes

`func NewPOSTPackagesRequestDataAttributes(name interface{}, length interface{}, width interface{}, height interface{}, unitOfLength interface{}, ) *POSTPackagesRequestDataAttributes`

NewPOSTPackagesRequestDataAttributes instantiates a new POSTPackagesRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPackagesRequestDataAttributesWithDefaults

`func NewPOSTPackagesRequestDataAttributesWithDefaults() *POSTPackagesRequestDataAttributes`

NewPOSTPackagesRequestDataAttributesWithDefaults instantiates a new POSTPackagesRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTPackagesRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTPackagesRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTPackagesRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTPackagesRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTPackagesRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *POSTPackagesRequestDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *POSTPackagesRequestDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *POSTPackagesRequestDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *POSTPackagesRequestDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *POSTPackagesRequestDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *POSTPackagesRequestDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetLength

`func (o *POSTPackagesRequestDataAttributes) GetLength() interface{}`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *POSTPackagesRequestDataAttributes) GetLengthOk() (*interface{}, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *POSTPackagesRequestDataAttributes) SetLength(v interface{})`

SetLength sets Length field to given value.


### SetLengthNil

`func (o *POSTPackagesRequestDataAttributes) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *POSTPackagesRequestDataAttributes) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetWidth

`func (o *POSTPackagesRequestDataAttributes) GetWidth() interface{}`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *POSTPackagesRequestDataAttributes) GetWidthOk() (*interface{}, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *POSTPackagesRequestDataAttributes) SetWidth(v interface{})`

SetWidth sets Width field to given value.


### SetWidthNil

`func (o *POSTPackagesRequestDataAttributes) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *POSTPackagesRequestDataAttributes) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *POSTPackagesRequestDataAttributes) GetHeight() interface{}`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *POSTPackagesRequestDataAttributes) GetHeightOk() (*interface{}, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *POSTPackagesRequestDataAttributes) SetHeight(v interface{})`

SetHeight sets Height field to given value.


### SetHeightNil

`func (o *POSTPackagesRequestDataAttributes) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *POSTPackagesRequestDataAttributes) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetUnitOfLength

`func (o *POSTPackagesRequestDataAttributes) GetUnitOfLength() interface{}`

GetUnitOfLength returns the UnitOfLength field if non-nil, zero value otherwise.

### GetUnitOfLengthOk

`func (o *POSTPackagesRequestDataAttributes) GetUnitOfLengthOk() (*interface{}, bool)`

GetUnitOfLengthOk returns a tuple with the UnitOfLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfLength

`func (o *POSTPackagesRequestDataAttributes) SetUnitOfLength(v interface{})`

SetUnitOfLength sets UnitOfLength field to given value.


### SetUnitOfLengthNil

`func (o *POSTPackagesRequestDataAttributes) SetUnitOfLengthNil(b bool)`

 SetUnitOfLengthNil sets the value for UnitOfLength to be an explicit nil

### UnsetUnitOfLength
`func (o *POSTPackagesRequestDataAttributes) UnsetUnitOfLength()`

UnsetUnitOfLength ensures that no value is present for UnitOfLength, not even an explicit nil
### GetReference

`func (o *POSTPackagesRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTPackagesRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTPackagesRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTPackagesRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTPackagesRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTPackagesRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTPackagesRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTPackagesRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTPackagesRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTPackagesRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTPackagesRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTPackagesRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTPackagesRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTPackagesRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTPackagesRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTPackagesRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTPackagesRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTPackagesRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


