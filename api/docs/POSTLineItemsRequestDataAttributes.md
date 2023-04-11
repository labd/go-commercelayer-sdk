# POSTLineItemsRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**BundleCode** | Pointer to **interface{}** | The code of the associated bundle. | [optional] 
**Quantity** | **interface{}** | The line item quantity. | 
**ExternalPrice** | Pointer to **interface{}** | When creating or updating a new line item, set this attribute to &#39;1&#39; if you want to inject the unit_amount_cents price from an external source. | [optional] 
**UpdateQuantity** | Pointer to **interface{}** | When creating a new line item, set this attribute to &#39;1&#39; if you want to update the line item quantity (if present) instead of creating a new line item for the same SKU. | [optional] 
**UnitAmountCents** | Pointer to **interface{}** | The unit amount of the line item, in cents. Can be specified without an item, otherwise is automatically populated from the price list associated to the order&#39;s market. | [optional] 
**Name** | Pointer to **interface{}** | The name of the line item. When blank, it gets populated with the name of the associated item (if present). | [optional] 
**ImageUrl** | Pointer to **interface{}** | The image_url of the line item. When blank, it gets populated with the image_url of the associated item (if present, SKU only). | [optional] 
**ItemType** | Pointer to **interface{}** | The type of the associate item. Can be one of &#39;skus&#39;, &#39;bundles&#39;, &#39;shipments&#39;, &#39;payment_methods&#39;, &#39;adjustments&#39;, &#39;gift_cards&#39;, or a valid promotion type. | [optional] 
**Frequency** | Pointer to **interface{}** | The frequency which generates a subscription. Must be supported by existing associated subscription_model. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTLineItemsRequestDataAttributes

`func NewPOSTLineItemsRequestDataAttributes(quantity interface{}, ) *POSTLineItemsRequestDataAttributes`

NewPOSTLineItemsRequestDataAttributes instantiates a new POSTLineItemsRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTLineItemsRequestDataAttributesWithDefaults

`func NewPOSTLineItemsRequestDataAttributesWithDefaults() *POSTLineItemsRequestDataAttributes`

NewPOSTLineItemsRequestDataAttributesWithDefaults instantiates a new POSTLineItemsRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *POSTLineItemsRequestDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *POSTLineItemsRequestDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *POSTLineItemsRequestDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *POSTLineItemsRequestDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *POSTLineItemsRequestDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *POSTLineItemsRequestDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetBundleCode

`func (o *POSTLineItemsRequestDataAttributes) GetBundleCode() interface{}`

GetBundleCode returns the BundleCode field if non-nil, zero value otherwise.

### GetBundleCodeOk

`func (o *POSTLineItemsRequestDataAttributes) GetBundleCodeOk() (*interface{}, bool)`

GetBundleCodeOk returns a tuple with the BundleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleCode

`func (o *POSTLineItemsRequestDataAttributes) SetBundleCode(v interface{})`

SetBundleCode sets BundleCode field to given value.

### HasBundleCode

`func (o *POSTLineItemsRequestDataAttributes) HasBundleCode() bool`

HasBundleCode returns a boolean if a field has been set.

### SetBundleCodeNil

`func (o *POSTLineItemsRequestDataAttributes) SetBundleCodeNil(b bool)`

 SetBundleCodeNil sets the value for BundleCode to be an explicit nil

### UnsetBundleCode
`func (o *POSTLineItemsRequestDataAttributes) UnsetBundleCode()`

UnsetBundleCode ensures that no value is present for BundleCode, not even an explicit nil
### GetQuantity

`func (o *POSTLineItemsRequestDataAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *POSTLineItemsRequestDataAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *POSTLineItemsRequestDataAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.


### SetQuantityNil

`func (o *POSTLineItemsRequestDataAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *POSTLineItemsRequestDataAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetExternalPrice

`func (o *POSTLineItemsRequestDataAttributes) GetExternalPrice() interface{}`

GetExternalPrice returns the ExternalPrice field if non-nil, zero value otherwise.

### GetExternalPriceOk

`func (o *POSTLineItemsRequestDataAttributes) GetExternalPriceOk() (*interface{}, bool)`

GetExternalPriceOk returns a tuple with the ExternalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrice

`func (o *POSTLineItemsRequestDataAttributes) SetExternalPrice(v interface{})`

SetExternalPrice sets ExternalPrice field to given value.

### HasExternalPrice

`func (o *POSTLineItemsRequestDataAttributes) HasExternalPrice() bool`

HasExternalPrice returns a boolean if a field has been set.

### SetExternalPriceNil

`func (o *POSTLineItemsRequestDataAttributes) SetExternalPriceNil(b bool)`

 SetExternalPriceNil sets the value for ExternalPrice to be an explicit nil

### UnsetExternalPrice
`func (o *POSTLineItemsRequestDataAttributes) UnsetExternalPrice()`

UnsetExternalPrice ensures that no value is present for ExternalPrice, not even an explicit nil
### GetUpdateQuantity

`func (o *POSTLineItemsRequestDataAttributes) GetUpdateQuantity() interface{}`

GetUpdateQuantity returns the UpdateQuantity field if non-nil, zero value otherwise.

### GetUpdateQuantityOk

`func (o *POSTLineItemsRequestDataAttributes) GetUpdateQuantityOk() (*interface{}, bool)`

GetUpdateQuantityOk returns a tuple with the UpdateQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateQuantity

`func (o *POSTLineItemsRequestDataAttributes) SetUpdateQuantity(v interface{})`

SetUpdateQuantity sets UpdateQuantity field to given value.

### HasUpdateQuantity

`func (o *POSTLineItemsRequestDataAttributes) HasUpdateQuantity() bool`

HasUpdateQuantity returns a boolean if a field has been set.

### SetUpdateQuantityNil

`func (o *POSTLineItemsRequestDataAttributes) SetUpdateQuantityNil(b bool)`

 SetUpdateQuantityNil sets the value for UpdateQuantity to be an explicit nil

### UnsetUpdateQuantity
`func (o *POSTLineItemsRequestDataAttributes) UnsetUpdateQuantity()`

UnsetUpdateQuantity ensures that no value is present for UpdateQuantity, not even an explicit nil
### GetUnitAmountCents

`func (o *POSTLineItemsRequestDataAttributes) GetUnitAmountCents() interface{}`

GetUnitAmountCents returns the UnitAmountCents field if non-nil, zero value otherwise.

### GetUnitAmountCentsOk

`func (o *POSTLineItemsRequestDataAttributes) GetUnitAmountCentsOk() (*interface{}, bool)`

GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountCents

`func (o *POSTLineItemsRequestDataAttributes) SetUnitAmountCents(v interface{})`

SetUnitAmountCents sets UnitAmountCents field to given value.

### HasUnitAmountCents

`func (o *POSTLineItemsRequestDataAttributes) HasUnitAmountCents() bool`

HasUnitAmountCents returns a boolean if a field has been set.

### SetUnitAmountCentsNil

`func (o *POSTLineItemsRequestDataAttributes) SetUnitAmountCentsNil(b bool)`

 SetUnitAmountCentsNil sets the value for UnitAmountCents to be an explicit nil

### UnsetUnitAmountCents
`func (o *POSTLineItemsRequestDataAttributes) UnsetUnitAmountCents()`

UnsetUnitAmountCents ensures that no value is present for UnitAmountCents, not even an explicit nil
### GetName

`func (o *POSTLineItemsRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTLineItemsRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTLineItemsRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *POSTLineItemsRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *POSTLineItemsRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTLineItemsRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImageUrl

`func (o *POSTLineItemsRequestDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *POSTLineItemsRequestDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *POSTLineItemsRequestDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *POSTLineItemsRequestDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *POSTLineItemsRequestDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *POSTLineItemsRequestDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetItemType

`func (o *POSTLineItemsRequestDataAttributes) GetItemType() interface{}`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *POSTLineItemsRequestDataAttributes) GetItemTypeOk() (*interface{}, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *POSTLineItemsRequestDataAttributes) SetItemType(v interface{})`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *POSTLineItemsRequestDataAttributes) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *POSTLineItemsRequestDataAttributes) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *POSTLineItemsRequestDataAttributes) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetFrequency

`func (o *POSTLineItemsRequestDataAttributes) GetFrequency() interface{}`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *POSTLineItemsRequestDataAttributes) GetFrequencyOk() (*interface{}, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *POSTLineItemsRequestDataAttributes) SetFrequency(v interface{})`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *POSTLineItemsRequestDataAttributes) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *POSTLineItemsRequestDataAttributes) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *POSTLineItemsRequestDataAttributes) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetReference

`func (o *POSTLineItemsRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTLineItemsRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTLineItemsRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTLineItemsRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTLineItemsRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTLineItemsRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTLineItemsRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTLineItemsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTLineItemsRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTLineItemsRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTLineItemsRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTLineItemsRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTLineItemsRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTLineItemsRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTLineItemsRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTLineItemsRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTLineItemsRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTLineItemsRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


