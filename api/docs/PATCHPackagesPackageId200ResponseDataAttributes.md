# PATCHPackagesPackageId200ResponseDataAttributes

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

### NewPATCHPackagesPackageId200ResponseDataAttributes

`func NewPATCHPackagesPackageId200ResponseDataAttributes() *PATCHPackagesPackageId200ResponseDataAttributes`

NewPATCHPackagesPackageId200ResponseDataAttributes instantiates a new PATCHPackagesPackageId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHPackagesPackageId200ResponseDataAttributesWithDefaults

`func NewPATCHPackagesPackageId200ResponseDataAttributesWithDefaults() *PATCHPackagesPackageId200ResponseDataAttributes`

NewPATCHPackagesPackageId200ResponseDataAttributesWithDefaults instantiates a new PATCHPackagesPackageId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLength

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWidth

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetUnitOfLength

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetUnitOfLength() string`

GetUnitOfLength returns the UnitOfLength field if non-nil, zero value otherwise.

### GetUnitOfLengthOk

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetUnitOfLengthOk() (*string, bool)`

GetUnitOfLengthOk returns a tuple with the UnitOfLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfLength

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetUnitOfLength(v string)`

SetUnitOfLength sets UnitOfLength field to given value.

### HasUnitOfLength

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasUnitOfLength() bool`

HasUnitOfLength returns a boolean if a field has been set.

### GetReference

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


