# POSTSkuListsRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The SKU list&#39;s internal name. | 
**Description** | Pointer to **interface{}** | An internal description of the SKU list. | [optional] 
**ImageUrl** | Pointer to **interface{}** | The URL of an image that represents the SKU list. | [optional] 
**Manual** | Pointer to **interface{}** | Indicates if the SKU list is populated manually. | [optional] 
**SkuCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the SKU codes. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTSkuListsRequestDataAttributes

`func NewPOSTSkuListsRequestDataAttributes(name interface{}, ) *POSTSkuListsRequestDataAttributes`

NewPOSTSkuListsRequestDataAttributes instantiates a new POSTSkuListsRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkuListsRequestDataAttributesWithDefaults

`func NewPOSTSkuListsRequestDataAttributesWithDefaults() *POSTSkuListsRequestDataAttributes`

NewPOSTSkuListsRequestDataAttributesWithDefaults instantiates a new POSTSkuListsRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTSkuListsRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTSkuListsRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTSkuListsRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTSkuListsRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTSkuListsRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *POSTSkuListsRequestDataAttributes) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *POSTSkuListsRequestDataAttributes) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *POSTSkuListsRequestDataAttributes) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *POSTSkuListsRequestDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *POSTSkuListsRequestDataAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *POSTSkuListsRequestDataAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *POSTSkuListsRequestDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *POSTSkuListsRequestDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *POSTSkuListsRequestDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *POSTSkuListsRequestDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *POSTSkuListsRequestDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *POSTSkuListsRequestDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetManual

`func (o *POSTSkuListsRequestDataAttributes) GetManual() interface{}`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *POSTSkuListsRequestDataAttributes) GetManualOk() (*interface{}, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *POSTSkuListsRequestDataAttributes) SetManual(v interface{})`

SetManual sets Manual field to given value.

### HasManual

`func (o *POSTSkuListsRequestDataAttributes) HasManual() bool`

HasManual returns a boolean if a field has been set.

### SetManualNil

`func (o *POSTSkuListsRequestDataAttributes) SetManualNil(b bool)`

 SetManualNil sets the value for Manual to be an explicit nil

### UnsetManual
`func (o *POSTSkuListsRequestDataAttributes) UnsetManual()`

UnsetManual ensures that no value is present for Manual, not even an explicit nil
### GetSkuCodeRegex

`func (o *POSTSkuListsRequestDataAttributes) GetSkuCodeRegex() interface{}`

GetSkuCodeRegex returns the SkuCodeRegex field if non-nil, zero value otherwise.

### GetSkuCodeRegexOk

`func (o *POSTSkuListsRequestDataAttributes) GetSkuCodeRegexOk() (*interface{}, bool)`

GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCodeRegex

`func (o *POSTSkuListsRequestDataAttributes) SetSkuCodeRegex(v interface{})`

SetSkuCodeRegex sets SkuCodeRegex field to given value.

### HasSkuCodeRegex

`func (o *POSTSkuListsRequestDataAttributes) HasSkuCodeRegex() bool`

HasSkuCodeRegex returns a boolean if a field has been set.

### SetSkuCodeRegexNil

`func (o *POSTSkuListsRequestDataAttributes) SetSkuCodeRegexNil(b bool)`

 SetSkuCodeRegexNil sets the value for SkuCodeRegex to be an explicit nil

### UnsetSkuCodeRegex
`func (o *POSTSkuListsRequestDataAttributes) UnsetSkuCodeRegex()`

UnsetSkuCodeRegex ensures that no value is present for SkuCodeRegex, not even an explicit nil
### GetReference

`func (o *POSTSkuListsRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTSkuListsRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTSkuListsRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTSkuListsRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTSkuListsRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTSkuListsRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTSkuListsRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTSkuListsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTSkuListsRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTSkuListsRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTSkuListsRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTSkuListsRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTSkuListsRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTSkuListsRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTSkuListsRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTSkuListsRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTSkuListsRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTSkuListsRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


