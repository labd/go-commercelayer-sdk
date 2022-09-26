# CarrierAccountDataAttributes

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

### NewCarrierAccountDataAttributes

`func NewCarrierAccountDataAttributes() *CarrierAccountDataAttributes`

NewCarrierAccountDataAttributes instantiates a new CarrierAccountDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarrierAccountDataAttributesWithDefaults

`func NewCarrierAccountDataAttributesWithDefaults() *CarrierAccountDataAttributes`

NewCarrierAccountDataAttributesWithDefaults instantiates a new CarrierAccountDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CarrierAccountDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CarrierAccountDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CarrierAccountDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CarrierAccountDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEasypostType

`func (o *CarrierAccountDataAttributes) GetEasypostType() string`

GetEasypostType returns the EasypostType field if non-nil, zero value otherwise.

### GetEasypostTypeOk

`func (o *CarrierAccountDataAttributes) GetEasypostTypeOk() (*string, bool)`

GetEasypostTypeOk returns a tuple with the EasypostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostType

`func (o *CarrierAccountDataAttributes) SetEasypostType(v string)`

SetEasypostType sets EasypostType field to given value.

### HasEasypostType

`func (o *CarrierAccountDataAttributes) HasEasypostType() bool`

HasEasypostType returns a boolean if a field has been set.

### GetEasypostId

`func (o *CarrierAccountDataAttributes) GetEasypostId() string`

GetEasypostId returns the EasypostId field if non-nil, zero value otherwise.

### GetEasypostIdOk

`func (o *CarrierAccountDataAttributes) GetEasypostIdOk() (*string, bool)`

GetEasypostIdOk returns a tuple with the EasypostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostId

`func (o *CarrierAccountDataAttributes) SetEasypostId(v string)`

SetEasypostId sets EasypostId field to given value.

### HasEasypostId

`func (o *CarrierAccountDataAttributes) HasEasypostId() bool`

HasEasypostId returns a boolean if a field has been set.

### GetId

`func (o *CarrierAccountDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CarrierAccountDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CarrierAccountDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CarrierAccountDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CarrierAccountDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CarrierAccountDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CarrierAccountDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CarrierAccountDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CarrierAccountDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CarrierAccountDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CarrierAccountDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CarrierAccountDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *CarrierAccountDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CarrierAccountDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CarrierAccountDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CarrierAccountDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *CarrierAccountDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *CarrierAccountDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *CarrierAccountDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *CarrierAccountDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *CarrierAccountDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CarrierAccountDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CarrierAccountDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CarrierAccountDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


