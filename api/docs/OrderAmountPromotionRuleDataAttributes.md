# OrderAmountPromotionRuleDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**OrderAmountCents** | Pointer to **int32** | Apply the promotion only when order is over this amount, in cents. | [optional] 
**OrderAmountFloat** | Pointer to **float32** | Apply the promotion only when order is over this amount, float. | [optional] 
**FormattedOrderAmount** | Pointer to **string** | Apply the promotion only when order is over this amount, formatted. | [optional] 

## Methods

### NewOrderAmountPromotionRuleDataAttributes

`func NewOrderAmountPromotionRuleDataAttributes() *OrderAmountPromotionRuleDataAttributes`

NewOrderAmountPromotionRuleDataAttributes instantiates a new OrderAmountPromotionRuleDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmountPromotionRuleDataAttributesWithDefaults

`func NewOrderAmountPromotionRuleDataAttributesWithDefaults() *OrderAmountPromotionRuleDataAttributes`

NewOrderAmountPromotionRuleDataAttributesWithDefaults instantiates a new OrderAmountPromotionRuleDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderAmountPromotionRuleDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderAmountPromotionRuleDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderAmountPromotionRuleDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderAmountPromotionRuleDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderAmountPromotionRuleDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderAmountPromotionRuleDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderAmountPromotionRuleDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderAmountPromotionRuleDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrderAmountPromotionRuleDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrderAmountPromotionRuleDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrderAmountPromotionRuleDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrderAmountPromotionRuleDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *OrderAmountPromotionRuleDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *OrderAmountPromotionRuleDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *OrderAmountPromotionRuleDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *OrderAmountPromotionRuleDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *OrderAmountPromotionRuleDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *OrderAmountPromotionRuleDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *OrderAmountPromotionRuleDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *OrderAmountPromotionRuleDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderAmountPromotionRuleDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderAmountPromotionRuleDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderAmountPromotionRuleDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderAmountPromotionRuleDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOrderAmountCents

`func (o *OrderAmountPromotionRuleDataAttributes) GetOrderAmountCents() int32`

GetOrderAmountCents returns the OrderAmountCents field if non-nil, zero value otherwise.

### GetOrderAmountCentsOk

`func (o *OrderAmountPromotionRuleDataAttributes) GetOrderAmountCentsOk() (*int32, bool)`

GetOrderAmountCentsOk returns a tuple with the OrderAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountCents

`func (o *OrderAmountPromotionRuleDataAttributes) SetOrderAmountCents(v int32)`

SetOrderAmountCents sets OrderAmountCents field to given value.

### HasOrderAmountCents

`func (o *OrderAmountPromotionRuleDataAttributes) HasOrderAmountCents() bool`

HasOrderAmountCents returns a boolean if a field has been set.

### GetOrderAmountFloat

`func (o *OrderAmountPromotionRuleDataAttributes) GetOrderAmountFloat() float32`

GetOrderAmountFloat returns the OrderAmountFloat field if non-nil, zero value otherwise.

### GetOrderAmountFloatOk

`func (o *OrderAmountPromotionRuleDataAttributes) GetOrderAmountFloatOk() (*float32, bool)`

GetOrderAmountFloatOk returns a tuple with the OrderAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountFloat

`func (o *OrderAmountPromotionRuleDataAttributes) SetOrderAmountFloat(v float32)`

SetOrderAmountFloat sets OrderAmountFloat field to given value.

### HasOrderAmountFloat

`func (o *OrderAmountPromotionRuleDataAttributes) HasOrderAmountFloat() bool`

HasOrderAmountFloat returns a boolean if a field has been set.

### GetFormattedOrderAmount

`func (o *OrderAmountPromotionRuleDataAttributes) GetFormattedOrderAmount() string`

GetFormattedOrderAmount returns the FormattedOrderAmount field if non-nil, zero value otherwise.

### GetFormattedOrderAmountOk

`func (o *OrderAmountPromotionRuleDataAttributes) GetFormattedOrderAmountOk() (*string, bool)`

GetFormattedOrderAmountOk returns a tuple with the FormattedOrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedOrderAmount

`func (o *OrderAmountPromotionRuleDataAttributes) SetFormattedOrderAmount(v string)`

SetFormattedOrderAmount sets FormattedOrderAmount field to given value.

### HasFormattedOrderAmount

`func (o *OrderAmountPromotionRuleDataAttributes) HasFormattedOrderAmount() bool`

HasFormattedOrderAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


