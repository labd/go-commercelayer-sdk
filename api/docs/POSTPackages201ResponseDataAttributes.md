# POSTPackages201ResponseDataAttributes

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

### NewPOSTPackages201ResponseDataAttributes

`func NewPOSTPackages201ResponseDataAttributes(name interface{}, length interface{}, width interface{}, height interface{}, unitOfLength interface{}, ) *POSTPackages201ResponseDataAttributes`

NewPOSTPackages201ResponseDataAttributes instantiates a new POSTPackages201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPackages201ResponseDataAttributesWithDefaults

`func NewPOSTPackages201ResponseDataAttributesWithDefaults() *POSTPackages201ResponseDataAttributes`

NewPOSTPackages201ResponseDataAttributesWithDefaults instantiates a new POSTPackages201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTPackages201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTPackages201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTPackages201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTPackages201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTPackages201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *POSTPackages201ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *POSTPackages201ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *POSTPackages201ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *POSTPackages201ResponseDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *POSTPackages201ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *POSTPackages201ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetLength

`func (o *POSTPackages201ResponseDataAttributes) GetLength() interface{}`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *POSTPackages201ResponseDataAttributes) GetLengthOk() (*interface{}, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *POSTPackages201ResponseDataAttributes) SetLength(v interface{})`

SetLength sets Length field to given value.


### SetLengthNil

`func (o *POSTPackages201ResponseDataAttributes) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *POSTPackages201ResponseDataAttributes) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetWidth

`func (o *POSTPackages201ResponseDataAttributes) GetWidth() interface{}`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *POSTPackages201ResponseDataAttributes) GetWidthOk() (*interface{}, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *POSTPackages201ResponseDataAttributes) SetWidth(v interface{})`

SetWidth sets Width field to given value.


### SetWidthNil

`func (o *POSTPackages201ResponseDataAttributes) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *POSTPackages201ResponseDataAttributes) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *POSTPackages201ResponseDataAttributes) GetHeight() interface{}`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *POSTPackages201ResponseDataAttributes) GetHeightOk() (*interface{}, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *POSTPackages201ResponseDataAttributes) SetHeight(v interface{})`

SetHeight sets Height field to given value.


### SetHeightNil

`func (o *POSTPackages201ResponseDataAttributes) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *POSTPackages201ResponseDataAttributes) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetUnitOfLength

`func (o *POSTPackages201ResponseDataAttributes) GetUnitOfLength() interface{}`

GetUnitOfLength returns the UnitOfLength field if non-nil, zero value otherwise.

### GetUnitOfLengthOk

`func (o *POSTPackages201ResponseDataAttributes) GetUnitOfLengthOk() (*interface{}, bool)`

GetUnitOfLengthOk returns a tuple with the UnitOfLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfLength

`func (o *POSTPackages201ResponseDataAttributes) SetUnitOfLength(v interface{})`

SetUnitOfLength sets UnitOfLength field to given value.


### SetUnitOfLengthNil

`func (o *POSTPackages201ResponseDataAttributes) SetUnitOfLengthNil(b bool)`

 SetUnitOfLengthNil sets the value for UnitOfLength to be an explicit nil

### UnsetUnitOfLength
`func (o *POSTPackages201ResponseDataAttributes) UnsetUnitOfLength()`

UnsetUnitOfLength ensures that no value is present for UnitOfLength, not even an explicit nil
### GetReference

`func (o *POSTPackages201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTPackages201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTPackages201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTPackages201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTPackages201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTPackages201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTPackages201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTPackages201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTPackages201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTPackages201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTPackages201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTPackages201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTPackages201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTPackages201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTPackages201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTPackages201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTPackages201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTPackages201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


