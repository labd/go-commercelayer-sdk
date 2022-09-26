# PATCHLineItemsLineItemId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **string** | The code of the associated SKU. | [optional] 
**BundleCode** | Pointer to **string** | The code of the associated bundle. | [optional] 
**Quantity** | Pointer to **int32** | The line item quantity. | [optional] 
**Name** | Pointer to **string** | The name of the line item. When blank, it gets populated with the name of the associated item (if present). | [optional] 
**ImageUrl** | Pointer to **string** | The image_url of the line item. When blank, it gets populated with the image_url of the associated item (if present, SKU only). | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHLineItemsLineItemId200ResponseDataAttributes

`func NewPATCHLineItemsLineItemId200ResponseDataAttributes() *PATCHLineItemsLineItemId200ResponseDataAttributes`

NewPATCHLineItemsLineItemId200ResponseDataAttributes instantiates a new PATCHLineItemsLineItemId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHLineItemsLineItemId200ResponseDataAttributesWithDefaults

`func NewPATCHLineItemsLineItemId200ResponseDataAttributesWithDefaults() *PATCHLineItemsLineItemId200ResponseDataAttributes`

NewPATCHLineItemsLineItemId200ResponseDataAttributesWithDefaults instantiates a new PATCHLineItemsLineItemId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetSkuCode() string`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetSkuCodeOk() (*string, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetSkuCode(v string)`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### GetBundleCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetBundleCode() string`

GetBundleCode returns the BundleCode field if non-nil, zero value otherwise.

### GetBundleCodeOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetBundleCodeOk() (*string, bool)`

GetBundleCodeOk returns a tuple with the BundleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetBundleCode(v string)`

SetBundleCode sets BundleCode field to given value.

### HasBundleCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasBundleCode() bool`

HasBundleCode returns a boolean if a field has been set.

### GetQuantity

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetName

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImageUrl

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetReference

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


