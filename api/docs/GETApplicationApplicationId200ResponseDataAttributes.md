# GETApplicationApplicationId200ResponseDataAttributes

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

### NewGETApplicationApplicationId200ResponseDataAttributes

`func NewGETApplicationApplicationId200ResponseDataAttributes() *GETApplicationApplicationId200ResponseDataAttributes`

NewGETApplicationApplicationId200ResponseDataAttributes instantiates a new GETApplicationApplicationId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETApplicationApplicationId200ResponseDataAttributesWithDefaults

`func NewGETApplicationApplicationId200ResponseDataAttributesWithDefaults() *GETApplicationApplicationId200ResponseDataAttributes`

NewGETApplicationApplicationId200ResponseDataAttributesWithDefaults instantiates a new GETApplicationApplicationId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETApplicationApplicationId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETApplicationApplicationId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GETApplicationApplicationId200ResponseDataAttributes) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GETApplicationApplicationId200ResponseDataAttributes) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPublicAccess

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetPublicAccess() bool`

GetPublicAccess returns the PublicAccess field if non-nil, zero value otherwise.

### GetPublicAccessOk

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetPublicAccessOk() (*bool, bool)`

GetPublicAccessOk returns a tuple with the PublicAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccess

`func (o *GETApplicationApplicationId200ResponseDataAttributes) SetPublicAccess(v bool)`

SetPublicAccess sets PublicAccess field to given value.

### HasPublicAccess

`func (o *GETApplicationApplicationId200ResponseDataAttributes) HasPublicAccess() bool`

HasPublicAccess returns a boolean if a field has been set.

### GetRedirectUri

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *GETApplicationApplicationId200ResponseDataAttributes) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *GETApplicationApplicationId200ResponseDataAttributes) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetScopes

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GETApplicationApplicationId200ResponseDataAttributes) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GETApplicationApplicationId200ResponseDataAttributes) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetId

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETApplicationApplicationId200ResponseDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETApplicationApplicationId200ResponseDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETApplicationApplicationId200ResponseDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETApplicationApplicationId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETApplicationApplicationId200ResponseDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETApplicationApplicationId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETApplicationApplicationId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETApplicationApplicationId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETApplicationApplicationId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETApplicationApplicationId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETApplicationApplicationId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETApplicationApplicationId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETApplicationApplicationId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


