# PackageUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name for the package | [optional] 
**Code** | Pointer to **string** | The package identifying code | [optional] 
**Length** | Pointer to **float32** | The package length, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**Width** | Pointer to **float32** | The package width, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**Height** | Pointer to **float32** | The package height, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**UnitOfLength** | Pointer to **string** | The unit of length. Can be one of &#39;cm&#39;, or &#39;in&#39;. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPackageUpdateDataAttributes

`func NewPackageUpdateDataAttributes() *PackageUpdateDataAttributes`

NewPackageUpdateDataAttributes instantiates a new PackageUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageUpdateDataAttributesWithDefaults

`func NewPackageUpdateDataAttributesWithDefaults() *PackageUpdateDataAttributes`

NewPackageUpdateDataAttributesWithDefaults instantiates a new PackageUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PackageUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackageUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *PackageUpdateDataAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PackageUpdateDataAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PackageUpdateDataAttributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PackageUpdateDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLength

`func (o *PackageUpdateDataAttributes) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PackageUpdateDataAttributes) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PackageUpdateDataAttributes) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *PackageUpdateDataAttributes) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWidth

`func (o *PackageUpdateDataAttributes) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PackageUpdateDataAttributes) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PackageUpdateDataAttributes) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PackageUpdateDataAttributes) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *PackageUpdateDataAttributes) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PackageUpdateDataAttributes) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PackageUpdateDataAttributes) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *PackageUpdateDataAttributes) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetUnitOfLength

`func (o *PackageUpdateDataAttributes) GetUnitOfLength() string`

GetUnitOfLength returns the UnitOfLength field if non-nil, zero value otherwise.

### GetUnitOfLengthOk

`func (o *PackageUpdateDataAttributes) GetUnitOfLengthOk() (*string, bool)`

GetUnitOfLengthOk returns a tuple with the UnitOfLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfLength

`func (o *PackageUpdateDataAttributes) SetUnitOfLength(v string)`

SetUnitOfLength sets UnitOfLength field to given value.

### HasUnitOfLength

`func (o *PackageUpdateDataAttributes) HasUnitOfLength() bool`

HasUnitOfLength returns a boolean if a field has been set.

### GetReference

`func (o *PackageUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PackageUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PackageUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PackageUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PackageUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PackageUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PackageUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PackageUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PackageUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PackageUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PackageUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PackageUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


