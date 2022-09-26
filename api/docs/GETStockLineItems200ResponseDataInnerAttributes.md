# GETStockLineItems200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **string** | The code of the associated SKU. | [optional] 
**BundleCode** | Pointer to **string** | The code of the associated bundle. | [optional] 
**Quantity** | Pointer to **int32** | The line item quantity. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETStockLineItems200ResponseDataInnerAttributes

`func NewGETStockLineItems200ResponseDataInnerAttributes() *GETStockLineItems200ResponseDataInnerAttributes`

NewGETStockLineItems200ResponseDataInnerAttributes instantiates a new GETStockLineItems200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETStockLineItems200ResponseDataInnerAttributesWithDefaults

`func NewGETStockLineItems200ResponseDataInnerAttributesWithDefaults() *GETStockLineItems200ResponseDataInnerAttributes`

NewGETStockLineItems200ResponseDataInnerAttributesWithDefaults instantiates a new GETStockLineItems200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetSkuCode() string`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetSkuCodeOk() (*string, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *GETStockLineItems200ResponseDataInnerAttributes) SetSkuCode(v string)`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *GETStockLineItems200ResponseDataInnerAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### GetBundleCode

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetBundleCode() string`

GetBundleCode returns the BundleCode field if non-nil, zero value otherwise.

### GetBundleCodeOk

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetBundleCodeOk() (*string, bool)`

GetBundleCodeOk returns a tuple with the BundleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleCode

`func (o *GETStockLineItems200ResponseDataInnerAttributes) SetBundleCode(v string)`

SetBundleCode sets BundleCode field to given value.

### HasBundleCode

`func (o *GETStockLineItems200ResponseDataInnerAttributes) HasBundleCode() bool`

HasBundleCode returns a boolean if a field has been set.

### GetQuantity

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GETStockLineItems200ResponseDataInnerAttributes) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GETStockLineItems200ResponseDataInnerAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetId

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETStockLineItems200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETStockLineItems200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETStockLineItems200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETStockLineItems200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETStockLineItems200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETStockLineItems200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETStockLineItems200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETStockLineItems200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETStockLineItems200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETStockLineItems200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETStockLineItems200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETStockLineItems200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETStockLineItems200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


