# PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU, replace the relationship | [optional] 
**StockThreshold** | Pointer to **interface{}** | The threshold at which to trigger the back in stock notification. | [optional] 
**Activate** | Pointer to **interface{}** | Send this attribute if you want to activate an inactive subscription. | [optional] 
**Deactivate** | Pointer to **interface{}** | Send this attribute if you want to dactivate an active subscription. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes

`func NewPATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes() *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes`

NewPATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes instantiates a new PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributesWithDefaults

`func NewPATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributesWithDefaults() *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes`

NewPATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributesWithDefaults instantiates a new PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetStockThreshold

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetStockThreshold() interface{}`

GetStockThreshold returns the StockThreshold field if non-nil, zero value otherwise.

### GetStockThresholdOk

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetStockThresholdOk() (*interface{}, bool)`

GetStockThresholdOk returns a tuple with the StockThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockThreshold

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetStockThreshold(v interface{})`

SetStockThreshold sets StockThreshold field to given value.

### HasStockThreshold

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) HasStockThreshold() bool`

HasStockThreshold returns a boolean if a field has been set.

### SetStockThresholdNil

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetStockThresholdNil(b bool)`

 SetStockThresholdNil sets the value for StockThreshold to be an explicit nil

### UnsetStockThreshold
`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) UnsetStockThreshold()`

UnsetStockThreshold ensures that no value is present for StockThreshold, not even an explicit nil
### GetActivate

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetActivate() interface{}`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetActivateOk() (*interface{}, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetActivate(v interface{})`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### SetActivateNil

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetActivateNil(b bool)`

 SetActivateNil sets the value for Activate to be an explicit nil

### UnsetActivate
`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) UnsetActivate()`

UnsetActivate ensures that no value is present for Activate, not even an explicit nil
### GetDeactivate

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetDeactivate() interface{}`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetDeactivateOk() (*interface{}, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetDeactivate(v interface{})`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### SetDeactivateNil

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetDeactivateNil(b bool)`

 SetDeactivateNil sets the value for Deactivate to be an explicit nil

### UnsetDeactivate
`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) UnsetDeactivate()`

UnsetDeactivate ensures that no value is present for Deactivate, not even an explicit nil
### GetReference

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


