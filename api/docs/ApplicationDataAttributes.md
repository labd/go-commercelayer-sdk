# ApplicationDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The application&#39;s internal name. | [optional] 
**Kind** | Pointer to **string** | The application&#39;s kind, can be one of: &#39;sales_channel&#39;, &#39;integration&#39; and &#39;webapp&#39;. | [optional] 
**PublicAccess** | Pointer to **bool** | Indicates if the application has public access. | [optional] 
**RedirectUri** | Pointer to **string** | The application&#39;s redirect URI. | [optional] 
**Scopes** | Pointer to **string** | The application&#39;s scopes. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewApplicationDataAttributes

`func NewApplicationDataAttributes() *ApplicationDataAttributes`

NewApplicationDataAttributes instantiates a new ApplicationDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDataAttributesWithDefaults

`func NewApplicationDataAttributesWithDefaults() *ApplicationDataAttributes`

NewApplicationDataAttributesWithDefaults instantiates a new ApplicationDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *ApplicationDataAttributes) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApplicationDataAttributes) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApplicationDataAttributes) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApplicationDataAttributes) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPublicAccess

`func (o *ApplicationDataAttributes) GetPublicAccess() bool`

GetPublicAccess returns the PublicAccess field if non-nil, zero value otherwise.

### GetPublicAccessOk

`func (o *ApplicationDataAttributes) GetPublicAccessOk() (*bool, bool)`

GetPublicAccessOk returns a tuple with the PublicAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccess

`func (o *ApplicationDataAttributes) SetPublicAccess(v bool)`

SetPublicAccess sets PublicAccess field to given value.

### HasPublicAccess

`func (o *ApplicationDataAttributes) HasPublicAccess() bool`

HasPublicAccess returns a boolean if a field has been set.

### GetRedirectUri

`func (o *ApplicationDataAttributes) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *ApplicationDataAttributes) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *ApplicationDataAttributes) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *ApplicationDataAttributes) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetScopes

`func (o *ApplicationDataAttributes) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApplicationDataAttributes) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApplicationDataAttributes) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApplicationDataAttributes) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetId

`func (o *ApplicationDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ApplicationDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApplicationDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *ApplicationDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ApplicationDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ApplicationDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ApplicationDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ApplicationDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ApplicationDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ApplicationDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ApplicationDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ApplicationDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


