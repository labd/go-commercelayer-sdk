# GETCustomerGroups200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The customer group&#39;s internal name | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETCustomerGroups200ResponseDataInnerAttributes

`func NewGETCustomerGroups200ResponseDataInnerAttributes() *GETCustomerGroups200ResponseDataInnerAttributes`

NewGETCustomerGroups200ResponseDataInnerAttributes instantiates a new GETCustomerGroups200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCustomerGroups200ResponseDataInnerAttributesWithDefaults

`func NewGETCustomerGroups200ResponseDataInnerAttributesWithDefaults() *GETCustomerGroups200ResponseDataInnerAttributes`

NewGETCustomerGroups200ResponseDataInnerAttributesWithDefaults instantiates a new GETCustomerGroups200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCustomerGroups200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


