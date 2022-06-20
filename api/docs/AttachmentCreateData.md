# AttachmentCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "attachments"]
**Attributes** | [**AttachmentCreateDataAttributes**](AttachmentCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**AttachmentCreateDataRelationships**](AttachmentCreateDataRelationships.md) |  | [optional] 

## Methods

### NewAttachmentCreateData

`func NewAttachmentCreateData(type_ string, attributes AttachmentCreateDataAttributes, ) *AttachmentCreateData`

NewAttachmentCreateData instantiates a new AttachmentCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentCreateDataWithDefaults

`func NewAttachmentCreateDataWithDefaults() *AttachmentCreateData`

NewAttachmentCreateDataWithDefaults instantiates a new AttachmentCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttachmentCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AttachmentCreateData) GetAttributes() AttachmentCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AttachmentCreateData) GetAttributesOk() (*AttachmentCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AttachmentCreateData) SetAttributes(v AttachmentCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AttachmentCreateData) GetRelationships() AttachmentCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AttachmentCreateData) GetRelationshipsOk() (*AttachmentCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AttachmentCreateData) SetRelationships(v AttachmentCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AttachmentCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


