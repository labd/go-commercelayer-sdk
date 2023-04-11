# GETStripeGatewaysStripeGatewayId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The payment gateway&#39;s internal name. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**AutoPayments** | Pointer to **interface{}** | Indicates if the gateway will accept payment methods enabled in the Stripe dashboard. | [optional] 
**WebhookEndpointId** | Pointer to **interface{}** | The gateway webhook endpoint ID, generated automatically. | [optional] 
**WebhookEndpointSecret** | Pointer to **interface{}** | The gateway webhook endpoint secret, generated automatically. | [optional] 
**WebhookEndpointUrl** | Pointer to **interface{}** | The gateway webhook URL, generated automatically. | [optional] 

## Methods

### NewGETStripeGatewaysStripeGatewayId200ResponseDataAttributes

`func NewGETStripeGatewaysStripeGatewayId200ResponseDataAttributes() *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes`

NewGETStripeGatewaysStripeGatewayId200ResponseDataAttributes instantiates a new GETStripeGatewaysStripeGatewayId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETStripeGatewaysStripeGatewayId200ResponseDataAttributesWithDefaults

`func NewGETStripeGatewaysStripeGatewayId200ResponseDataAttributesWithDefaults() *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes`

NewGETStripeGatewaysStripeGatewayId200ResponseDataAttributesWithDefaults instantiates a new GETStripeGatewaysStripeGatewayId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedAt

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAutoPayments

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetAutoPayments() interface{}`

GetAutoPayments returns the AutoPayments field if non-nil, zero value otherwise.

### GetAutoPaymentsOk

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetAutoPaymentsOk() (*interface{}, bool)`

GetAutoPaymentsOk returns a tuple with the AutoPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPayments

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetAutoPayments(v interface{})`

SetAutoPayments sets AutoPayments field to given value.

### HasAutoPayments

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasAutoPayments() bool`

HasAutoPayments returns a boolean if a field has been set.

### SetAutoPaymentsNil

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetAutoPaymentsNil(b bool)`

 SetAutoPaymentsNil sets the value for AutoPayments to be an explicit nil

### UnsetAutoPayments
`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetAutoPayments()`

UnsetAutoPayments ensures that no value is present for AutoPayments, not even an explicit nil
### GetWebhookEndpointId

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetWebhookEndpointId() interface{}`

GetWebhookEndpointId returns the WebhookEndpointId field if non-nil, zero value otherwise.

### GetWebhookEndpointIdOk

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetWebhookEndpointIdOk() (*interface{}, bool)`

GetWebhookEndpointIdOk returns a tuple with the WebhookEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpointId

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetWebhookEndpointId(v interface{})`

SetWebhookEndpointId sets WebhookEndpointId field to given value.

### HasWebhookEndpointId

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasWebhookEndpointId() bool`

HasWebhookEndpointId returns a boolean if a field has been set.

### SetWebhookEndpointIdNil

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetWebhookEndpointIdNil(b bool)`

 SetWebhookEndpointIdNil sets the value for WebhookEndpointId to be an explicit nil

### UnsetWebhookEndpointId
`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetWebhookEndpointId()`

UnsetWebhookEndpointId ensures that no value is present for WebhookEndpointId, not even an explicit nil
### GetWebhookEndpointSecret

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetWebhookEndpointSecret() interface{}`

GetWebhookEndpointSecret returns the WebhookEndpointSecret field if non-nil, zero value otherwise.

### GetWebhookEndpointSecretOk

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetWebhookEndpointSecretOk() (*interface{}, bool)`

GetWebhookEndpointSecretOk returns a tuple with the WebhookEndpointSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpointSecret

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetWebhookEndpointSecret(v interface{})`

SetWebhookEndpointSecret sets WebhookEndpointSecret field to given value.

### HasWebhookEndpointSecret

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasWebhookEndpointSecret() bool`

HasWebhookEndpointSecret returns a boolean if a field has been set.

### SetWebhookEndpointSecretNil

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetWebhookEndpointSecretNil(b bool)`

 SetWebhookEndpointSecretNil sets the value for WebhookEndpointSecret to be an explicit nil

### UnsetWebhookEndpointSecret
`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetWebhookEndpointSecret()`

UnsetWebhookEndpointSecret ensures that no value is present for WebhookEndpointSecret, not even an explicit nil
### GetWebhookEndpointUrl

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetWebhookEndpointUrl() interface{}`

GetWebhookEndpointUrl returns the WebhookEndpointUrl field if non-nil, zero value otherwise.

### GetWebhookEndpointUrlOk

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetWebhookEndpointUrlOk() (*interface{}, bool)`

GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpointUrl

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetWebhookEndpointUrl(v interface{})`

SetWebhookEndpointUrl sets WebhookEndpointUrl field to given value.

### HasWebhookEndpointUrl

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasWebhookEndpointUrl() bool`

HasWebhookEndpointUrl returns a boolean if a field has been set.

### SetWebhookEndpointUrlNil

`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetWebhookEndpointUrlNil(b bool)`

 SetWebhookEndpointUrlNil sets the value for WebhookEndpointUrl to be an explicit nil

### UnsetWebhookEndpointUrl
`func (o *GETStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetWebhookEndpointUrl()`

UnsetWebhookEndpointUrl ensures that no value is present for WebhookEndpointUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


