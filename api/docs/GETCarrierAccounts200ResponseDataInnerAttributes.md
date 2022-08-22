# GETCarrierAccounts200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The carrier account internal name. | [optional] 
**EasypostType** | Pointer to **string** | The Easypost service carrier type. | [optional] 
**EasypostId** | Pointer to **string** | The Easypost internal reference ID. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETCarrierAccounts200ResponseDataInnerAttributes

`func NewGETCarrierAccounts200ResponseDataInnerAttributes() *GETCarrierAccounts200ResponseDataInnerAttributes`

NewGETCarrierAccounts200ResponseDataInnerAttributes instantiates a new GETCarrierAccounts200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCarrierAccounts200ResponseDataInnerAttributesWithDefaults

`func NewGETCarrierAccounts200ResponseDataInnerAttributesWithDefaults() *GETCarrierAccounts200ResponseDataInnerAttributes`

NewGETCarrierAccounts200ResponseDataInnerAttributesWithDefaults instantiates a new GETCarrierAccounts200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEasypostType

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetEasypostType() string`

GetEasypostType returns the EasypostType field if non-nil, zero value otherwise.

### GetEasypostTypeOk

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetEasypostTypeOk() (*string, bool)`

GetEasypostTypeOk returns a tuple with the EasypostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostType

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) SetEasypostType(v string)`

SetEasypostType sets EasypostType field to given value.

### HasEasypostType

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) HasEasypostType() bool`

HasEasypostType returns a boolean if a field has been set.

### GetEasypostId

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetEasypostId() string`

GetEasypostId returns the EasypostId field if non-nil, zero value otherwise.

### GetEasypostIdOk

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetEasypostIdOk() (*string, bool)`

GetEasypostIdOk returns a tuple with the EasypostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostId

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) SetEasypostId(v string)`

SetEasypostId sets EasypostId field to given value.

### HasEasypostId

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) HasEasypostId() bool`

HasEasypostId returns a boolean if a field has been set.

### GetId

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCarrierAccounts200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


