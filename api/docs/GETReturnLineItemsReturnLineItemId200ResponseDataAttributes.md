# GETReturnLineItemsReturnLineItemId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**BundleCode** | Pointer to **interface{}** | The code of the associated bundle. | [optional] 
**Quantity** | Pointer to **interface{}** | The return line item quantity. | [optional] 
**Name** | Pointer to **interface{}** | The name of the line item. | [optional] 
**ImageUrl** | Pointer to **interface{}** | The image_url of the associated line item. | [optional] 
**TotalAmountCents** | Pointer to **interface{}** | Calculated as line item unit amount x returned quantity and applied discounts, if any. | [optional] 
**TotalAmountFloat** | Pointer to **interface{}** | The return line item total amount, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedTotalAmount** | Pointer to **interface{}** | The return line item total amount, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**ReturnReason** | Pointer to **interface{}** | Set of key-value pairs that you can use to add details about return reason. | [optional] 
**RestockedAt** | Pointer to **interface{}** | Time at which the return line item was restocked. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETReturnLineItemsReturnLineItemId200ResponseDataAttributes

`func NewGETReturnLineItemsReturnLineItemId200ResponseDataAttributes() *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes`

NewGETReturnLineItemsReturnLineItemId200ResponseDataAttributes instantiates a new GETReturnLineItemsReturnLineItemId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETReturnLineItemsReturnLineItemId200ResponseDataAttributesWithDefaults

`func NewGETReturnLineItemsReturnLineItemId200ResponseDataAttributesWithDefaults() *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes`

NewGETReturnLineItemsReturnLineItemId200ResponseDataAttributesWithDefaults instantiates a new GETReturnLineItemsReturnLineItemId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetBundleCode

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetBundleCode() interface{}`

GetBundleCode returns the BundleCode field if non-nil, zero value otherwise.

### GetBundleCodeOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetBundleCodeOk() (*interface{}, bool)`

GetBundleCodeOk returns a tuple with the BundleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleCode

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetBundleCode(v interface{})`

SetBundleCode sets BundleCode field to given value.

### HasBundleCode

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasBundleCode() bool`

HasBundleCode returns a boolean if a field has been set.

### SetBundleCodeNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetBundleCodeNil(b bool)`

 SetBundleCodeNil sets the value for BundleCode to be an explicit nil

### UnsetBundleCode
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetBundleCode()`

UnsetBundleCode ensures that no value is present for BundleCode, not even an explicit nil
### GetQuantity

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetName

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImageUrl

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetTotalAmountCents

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetTotalAmountCents() interface{}`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetTotalAmountCentsOk() (*interface{}, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetTotalAmountCents(v interface{})`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### SetTotalAmountCentsNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetTotalAmountCentsNil(b bool)`

 SetTotalAmountCentsNil sets the value for TotalAmountCents to be an explicit nil

### UnsetTotalAmountCents
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetTotalAmountCents()`

UnsetTotalAmountCents ensures that no value is present for TotalAmountCents, not even an explicit nil
### GetTotalAmountFloat

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetTotalAmountFloat() interface{}`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetTotalAmountFloatOk() (*interface{}, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetTotalAmountFloat(v interface{})`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### SetTotalAmountFloatNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetTotalAmountFloatNil(b bool)`

 SetTotalAmountFloatNil sets the value for TotalAmountFloat to be an explicit nil

### UnsetTotalAmountFloat
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetTotalAmountFloat()`

UnsetTotalAmountFloat ensures that no value is present for TotalAmountFloat, not even an explicit nil
### GetFormattedTotalAmount

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetFormattedTotalAmount() interface{}`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetFormattedTotalAmountOk() (*interface{}, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetFormattedTotalAmount(v interface{})`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### SetFormattedTotalAmountNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetFormattedTotalAmountNil(b bool)`

 SetFormattedTotalAmountNil sets the value for FormattedTotalAmount to be an explicit nil

### UnsetFormattedTotalAmount
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetFormattedTotalAmount()`

UnsetFormattedTotalAmount ensures that no value is present for FormattedTotalAmount, not even an explicit nil
### GetReturnReason

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReturnReason() interface{}`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReturnReasonOk() (*interface{}, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetReturnReason(v interface{})`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### SetReturnReasonNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetReturnReasonNil(b bool)`

 SetReturnReasonNil sets the value for ReturnReason to be an explicit nil

### UnsetReturnReason
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetReturnReason()`

UnsetReturnReason ensures that no value is present for ReturnReason, not even an explicit nil
### GetRestockedAt

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetRestockedAt() interface{}`

GetRestockedAt returns the RestockedAt field if non-nil, zero value otherwise.

### GetRestockedAtOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetRestockedAtOk() (*interface{}, bool)`

GetRestockedAtOk returns a tuple with the RestockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockedAt

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetRestockedAt(v interface{})`

SetRestockedAt sets RestockedAt field to given value.

### HasRestockedAt

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasRestockedAt() bool`

HasRestockedAt returns a boolean if a field has been set.

### SetRestockedAtNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetRestockedAtNil(b bool)`

 SetRestockedAtNil sets the value for RestockedAt to be an explicit nil

### UnsetRestockedAt
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetRestockedAt()`

UnsetRestockedAt ensures that no value is present for RestockedAt, not even an explicit nil
### GetCreatedAt

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


