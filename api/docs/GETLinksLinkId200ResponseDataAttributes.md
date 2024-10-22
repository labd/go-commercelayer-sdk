# GETLinksLinkId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The link internal name. | [optional] 
**ClientId** | Pointer to **interface{}** | The link application client id, used to fetch JWT. | [optional] 
**Scope** | Pointer to **interface{}** | The link application scope, used to fetch JWT. | [optional] 
**StartsAt** | Pointer to **interface{}** | The activation date/time of this link. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | The expiration date/time of this link (must be after starts_at). | [optional] 
**Active** | Pointer to **interface{}** | Indicates if the link is active (enabled and not expired). | [optional] 
**Status** | Pointer to **interface{}** | The link status. One of &#39;disabled&#39;, &#39;expired&#39;, &#39;pending&#39;, or &#39;active&#39;. | [optional] 
**Domain** | Pointer to **interface{}** | The link URL second level domain. | [optional] 
**Url** | Pointer to **interface{}** | The link URL. | [optional] 
**ItemType** | Pointer to **interface{}** | The type of the associated item. One of &#39;orders&#39;, &#39;skus&#39;, or &#39;sku_lists&#39;. | [optional] 
**Params** | Pointer to **interface{}** | The link params to be passed in URL the query string. | [optional] 
**DisabledAt** | Pointer to **interface{}** | Time at which this resource was disabled. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETLinksLinkId200ResponseDataAttributes

`func NewGETLinksLinkId200ResponseDataAttributes() *GETLinksLinkId200ResponseDataAttributes`

NewGETLinksLinkId200ResponseDataAttributes instantiates a new GETLinksLinkId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETLinksLinkId200ResponseDataAttributesWithDefaults

`func NewGETLinksLinkId200ResponseDataAttributesWithDefaults() *GETLinksLinkId200ResponseDataAttributes`

NewGETLinksLinkId200ResponseDataAttributesWithDefaults instantiates a new GETLinksLinkId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETLinksLinkId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETLinksLinkId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETLinksLinkId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetClientId

`func (o *GETLinksLinkId200ResponseDataAttributes) GetClientId() interface{}`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetClientIdOk() (*interface{}, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GETLinksLinkId200ResponseDataAttributes) SetClientId(v interface{})`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GETLinksLinkId200ResponseDataAttributes) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetScope

`func (o *GETLinksLinkId200ResponseDataAttributes) GetScope() interface{}`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetScopeOk() (*interface{}, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GETLinksLinkId200ResponseDataAttributes) SetScope(v interface{})`

SetScope sets Scope field to given value.

### HasScope

`func (o *GETLinksLinkId200ResponseDataAttributes) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetStartsAt

`func (o *GETLinksLinkId200ResponseDataAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GETLinksLinkId200ResponseDataAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GETLinksLinkId200ResponseDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *GETLinksLinkId200ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETLinksLinkId200ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETLinksLinkId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetActive

`func (o *GETLinksLinkId200ResponseDataAttributes) GetActive() interface{}`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetActiveOk() (*interface{}, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GETLinksLinkId200ResponseDataAttributes) SetActive(v interface{})`

SetActive sets Active field to given value.

### HasActive

`func (o *GETLinksLinkId200ResponseDataAttributes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetStatus

`func (o *GETLinksLinkId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETLinksLinkId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETLinksLinkId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDomain

`func (o *GETLinksLinkId200ResponseDataAttributes) GetDomain() interface{}`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetDomainOk() (*interface{}, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GETLinksLinkId200ResponseDataAttributes) SetDomain(v interface{})`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GETLinksLinkId200ResponseDataAttributes) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetUrl

`func (o *GETLinksLinkId200ResponseDataAttributes) GetUrl() interface{}`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetUrlOk() (*interface{}, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GETLinksLinkId200ResponseDataAttributes) SetUrl(v interface{})`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GETLinksLinkId200ResponseDataAttributes) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetItemType

`func (o *GETLinksLinkId200ResponseDataAttributes) GetItemType() interface{}`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetItemTypeOk() (*interface{}, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *GETLinksLinkId200ResponseDataAttributes) SetItemType(v interface{})`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *GETLinksLinkId200ResponseDataAttributes) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetParams

`func (o *GETLinksLinkId200ResponseDataAttributes) GetParams() interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetParamsOk() (*interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GETLinksLinkId200ResponseDataAttributes) SetParams(v interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *GETLinksLinkId200ResponseDataAttributes) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetDisabledAt

`func (o *GETLinksLinkId200ResponseDataAttributes) GetDisabledAt() interface{}`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *GETLinksLinkId200ResponseDataAttributes) SetDisabledAt(v interface{})`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *GETLinksLinkId200ResponseDataAttributes) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### SetDisabledAtNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetCreatedAt

`func (o *GETLinksLinkId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETLinksLinkId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETLinksLinkId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETLinksLinkId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETLinksLinkId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETLinksLinkId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETLinksLinkId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETLinksLinkId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETLinksLinkId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETLinksLinkId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETLinksLinkId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETLinksLinkId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETLinksLinkId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETLinksLinkId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETLinksLinkId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETLinksLinkId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETLinksLinkId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETLinksLinkId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


