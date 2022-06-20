# ReturnLineItemDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **string** | The code of the associated sku. | [optional] 
**BundleCode** | Pointer to **string** | The code of the associated bundle. | [optional] 
**Name** | Pointer to **string** | The name of the line item. | [optional] 
**Quantity** | Pointer to **int32** | The line item quantity. | [optional] 
**ReturnReason** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can use to add details about return reason. | [optional] 
**RestockedAt** | Pointer to **string** | Time at which the return line item was restocked. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewReturnLineItemDataAttributes

`func NewReturnLineItemDataAttributes() *ReturnLineItemDataAttributes`

NewReturnLineItemDataAttributes instantiates a new ReturnLineItemDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnLineItemDataAttributesWithDefaults

`func NewReturnLineItemDataAttributesWithDefaults() *ReturnLineItemDataAttributes`

NewReturnLineItemDataAttributesWithDefaults instantiates a new ReturnLineItemDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *ReturnLineItemDataAttributes) GetSkuCode() string`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *ReturnLineItemDataAttributes) GetSkuCodeOk() (*string, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *ReturnLineItemDataAttributes) SetSkuCode(v string)`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *ReturnLineItemDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### GetBundleCode

`func (o *ReturnLineItemDataAttributes) GetBundleCode() string`

GetBundleCode returns the BundleCode field if non-nil, zero value otherwise.

### GetBundleCodeOk

`func (o *ReturnLineItemDataAttributes) GetBundleCodeOk() (*string, bool)`

GetBundleCodeOk returns a tuple with the BundleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleCode

`func (o *ReturnLineItemDataAttributes) SetBundleCode(v string)`

SetBundleCode sets BundleCode field to given value.

### HasBundleCode

`func (o *ReturnLineItemDataAttributes) HasBundleCode() bool`

HasBundleCode returns a boolean if a field has been set.

### GetName

`func (o *ReturnLineItemDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReturnLineItemDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReturnLineItemDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReturnLineItemDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *ReturnLineItemDataAttributes) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReturnLineItemDataAttributes) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReturnLineItemDataAttributes) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReturnLineItemDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReturnReason

`func (o *ReturnLineItemDataAttributes) GetReturnReason() map[string]interface{}`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *ReturnLineItemDataAttributes) GetReturnReasonOk() (*map[string]interface{}, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *ReturnLineItemDataAttributes) SetReturnReason(v map[string]interface{})`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *ReturnLineItemDataAttributes) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetRestockedAt

`func (o *ReturnLineItemDataAttributes) GetRestockedAt() string`

GetRestockedAt returns the RestockedAt field if non-nil, zero value otherwise.

### GetRestockedAtOk

`func (o *ReturnLineItemDataAttributes) GetRestockedAtOk() (*string, bool)`

GetRestockedAtOk returns a tuple with the RestockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockedAt

`func (o *ReturnLineItemDataAttributes) SetRestockedAt(v string)`

SetRestockedAt sets RestockedAt field to given value.

### HasRestockedAt

`func (o *ReturnLineItemDataAttributes) HasRestockedAt() bool`

HasRestockedAt returns a boolean if a field has been set.

### GetId

`func (o *ReturnLineItemDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnLineItemDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnLineItemDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReturnLineItemDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReturnLineItemDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReturnLineItemDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReturnLineItemDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReturnLineItemDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ReturnLineItemDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReturnLineItemDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReturnLineItemDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReturnLineItemDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *ReturnLineItemDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ReturnLineItemDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ReturnLineItemDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ReturnLineItemDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ReturnLineItemDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ReturnLineItemDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ReturnLineItemDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ReturnLineItemDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ReturnLineItemDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReturnLineItemDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReturnLineItemDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ReturnLineItemDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


