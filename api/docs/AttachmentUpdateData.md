# AttachmentUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHAttachmentsAttachmentIdRequestDataAttributes**](PATCHAttachmentsAttachmentIdRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**AttachmentUpdateDataRelationships**](AttachmentUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewAttachmentUpdateData

`func NewAttachmentUpdateData(type_ interface{}, id interface{}, attributes PATCHAttachmentsAttachmentIdRequestDataAttributes, ) *AttachmentUpdateData`

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

`func (o *AttachmentUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AttachmentUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AttachmentUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *AttachmentUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *AttachmentUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AttachmentUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *AttachmentUpdateData) GetAttributes() PATCHAttachmentsAttachmentIdRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AttachmentUpdateData) GetAttributesOk() (*PATCHAttachmentsAttachmentIdRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AttachmentUpdateData) SetAttributes(v PATCHAttachmentsAttachmentIdRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AttachmentUpdateData) GetRelationships() AttachmentUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AttachmentUpdateData) GetRelationshipsOk() (*AttachmentUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AttachmentUpdateData) SetRelationships(v AttachmentUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AttachmentUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


