# PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUrl** | Pointer to **interface{}** | The url to redirect the customer after the payment flow is completed. | [optional] 
**Refresh** | Pointer to **interface{}** | Send this attribute if you want to refresh all the pending transactions, can be used as webhooks fallback logic. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes

`func NewPATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes() *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes`

NewPATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes instantiates a new PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributesWithDefaults

`func NewPATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributesWithDefaults() *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes`

NewPATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributesWithDefaults instantiates a new PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUrl

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) GetRedirectUrl() interface{}`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) GetRedirectUrlOk() (*interface{}, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) SetRedirectUrl(v interface{})`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetRefresh

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) GetRefresh() interface{}`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) GetRefreshOk() (*interface{}, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) SetRefresh(v interface{})`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### SetRefreshNil

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) SetRefreshNil(b bool)`

 SetRefreshNil sets the value for Refresh to be an explicit nil

### UnsetRefresh
`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) UnsetRefresh()`

UnsetRefresh ensures that no value is present for Refresh, not even an explicit nil
### GetReference

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHSatispayPaymentsSatispayPaymentIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


