# POSTSkuLists201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The SKU list&#39;s internal name. | 
**Description** | Pointer to **string** | An internal description of the SKU list. | [optional] 
**ImageUrl** | Pointer to **string** | The URL of an image that represents the SKU list. | [optional] 
**Manual** | Pointer to **bool** | Indicates if the SKU list is populated manually. | [optional] 
**SkuCodeRegex** | Pointer to **string** | The regex that will be evaluated to match the SKU codes. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTSkuLists201ResponseDataAttributes

`func NewPOSTSkuLists201ResponseDataAttributes(name string, ) *POSTSkuLists201ResponseDataAttributes`

NewPOSTSkuLists201ResponseDataAttributes instantiates a new POSTSkuLists201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkuLists201ResponseDataAttributesWithDefaults

`func NewPOSTSkuLists201ResponseDataAttributesWithDefaults() *POSTSkuLists201ResponseDataAttributes`

NewPOSTSkuLists201ResponseDataAttributesWithDefaults instantiates a new POSTSkuLists201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTSkuLists201ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTSkuLists201ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTSkuLists201ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *POSTSkuLists201ResponseDataAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *POSTSkuLists201ResponseDataAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *POSTSkuLists201ResponseDataAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *POSTSkuLists201ResponseDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageUrl

`func (o *POSTSkuLists201ResponseDataAttributes) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *POSTSkuLists201ResponseDataAttributes) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *POSTSkuLists201ResponseDataAttributes) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *POSTSkuLists201ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetManual

`func (o *POSTSkuLists201ResponseDataAttributes) GetManual() bool`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *POSTSkuLists201ResponseDataAttributes) GetManualOk() (*bool, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *POSTSkuLists201ResponseDataAttributes) SetManual(v bool)`

SetManual sets Manual field to given value.

### HasManual

`func (o *POSTSkuLists201ResponseDataAttributes) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetSkuCodeRegex

`func (o *POSTSkuLists201ResponseDataAttributes) GetSkuCodeRegex() string`

GetSkuCodeRegex returns the SkuCodeRegex field if non-nil, zero value otherwise.

### GetSkuCodeRegexOk

`func (o *POSTSkuLists201ResponseDataAttributes) GetSkuCodeRegexOk() (*string, bool)`

GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCodeRegex

`func (o *POSTSkuLists201ResponseDataAttributes) SetSkuCodeRegex(v string)`

SetSkuCodeRegex sets SkuCodeRegex field to given value.

### HasSkuCodeRegex

`func (o *POSTSkuLists201ResponseDataAttributes) HasSkuCodeRegex() bool`

HasSkuCodeRegex returns a boolean if a field has been set.

### GetReference

`func (o *POSTSkuLists201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTSkuLists201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTSkuLists201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTSkuLists201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTSkuLists201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTSkuLists201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTSkuLists201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTSkuLists201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTSkuLists201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTSkuLists201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTSkuLists201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTSkuLists201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


