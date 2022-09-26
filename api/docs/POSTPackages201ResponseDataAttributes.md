# POSTPackages201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name for the package | 
**Code** | Pointer to **string** | The package identifying code | [optional] 
**Length** | **float32** | The package length, used to automatically calculate the tax rates from the available carrier accounts. | 
**Width** | **float32** | The package width, used to automatically calculate the tax rates from the available carrier accounts. | 
**Height** | **float32** | The package height, used to automatically calculate the tax rates from the available carrier accounts. | 
**UnitOfLength** | **string** | The unit of length. Can be one of &#39;cm&#39;, or &#39;in&#39;. | 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTPackages201ResponseDataAttributes

`func NewPOSTPackages201ResponseDataAttributes(name string, length float32, width float32, height float32, unitOfLength string, ) *POSTPackages201ResponseDataAttributes`

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

`func (o *POSTPackages201ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTPackages201ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTPackages201ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *POSTPackages201ResponseDataAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *POSTPackages201ResponseDataAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *POSTPackages201ResponseDataAttributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *POSTPackages201ResponseDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLength

`func (o *POSTPackages201ResponseDataAttributes) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *POSTPackages201ResponseDataAttributes) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *POSTPackages201ResponseDataAttributes) SetLength(v float32)`

SetLength sets Length field to given value.


### GetWidth

`func (o *POSTPackages201ResponseDataAttributes) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *POSTPackages201ResponseDataAttributes) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *POSTPackages201ResponseDataAttributes) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *POSTPackages201ResponseDataAttributes) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *POSTPackages201ResponseDataAttributes) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *POSTPackages201ResponseDataAttributes) SetHeight(v float32)`

SetHeight sets Height field to given value.


### GetUnitOfLength

`func (o *POSTPackages201ResponseDataAttributes) GetUnitOfLength() string`

GetUnitOfLength returns the UnitOfLength field if non-nil, zero value otherwise.

### GetUnitOfLengthOk

`func (o *POSTPackages201ResponseDataAttributes) GetUnitOfLengthOk() (*string, bool)`

GetUnitOfLengthOk returns a tuple with the UnitOfLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfLength

`func (o *POSTPackages201ResponseDataAttributes) SetUnitOfLength(v string)`

SetUnitOfLength sets UnitOfLength field to given value.


### GetReference

`func (o *POSTPackages201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTPackages201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTPackages201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTPackages201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTPackages201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTPackages201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTPackages201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTPackages201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTPackages201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTPackages201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTPackages201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTPackages201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


