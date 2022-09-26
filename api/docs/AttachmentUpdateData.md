# AttachmentUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHAttachmentsAttachmentId200ResponseDataAttributes**](PATCHAttachmentsAttachmentId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AttachmentDataRelationships**](AttachmentDataRelationships.md) |  | [optional] 

## Methods

### NewAttachmentUpdateData

`func NewAttachmentUpdateData(type_ string, id string, attributes PATCHAttachmentsAttachmentId200ResponseDataAttributes, ) *AttachmentUpdateData`

NewAttachmentUpdateData instantiates a new AttachmentUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentUpdateDataWithDefaults

`func NewAttachmentUpdateDataWithDefaults() *AttachmentUpdateData`

NewAttachmentUpdateDataWithDefaults instantiates a new AttachmentUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttachmentUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AttachmentUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AttachmentUpdateData) GetAttributes() PATCHAttachmentsAttachmentId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AttachmentUpdateData) GetAttributesOk() (*PATCHAttachmentsAttachmentId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AttachmentUpdateData) SetAttributes(v PATCHAttachmentsAttachmentId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AttachmentUpdateData) GetRelationships() AttachmentDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AttachmentUpdateData) GetRelationshipsOk() (*AttachmentDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AttachmentUpdateData) SetRelationships(v AttachmentDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AttachmentUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


