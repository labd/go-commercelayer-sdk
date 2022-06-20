# CheckoutComGatewayDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The payment gateway&#39;s internal name. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**WebhookEndpointId** | Pointer to **string** | The gateway webhook endpoint ID, generated automatically. | [optional] 
**WebhookEndpointSecret** | Pointer to **string** | The gateway webhook endpoint secret, generated automatically. | [optional] 
**WebhookEndpointUrl** | Pointer to **string** | The gateway webhook URL, generated automatically. | [optional] 

## Methods

### NewCheckoutComGatewayDataAttributes

`func NewCheckoutComGatewayDataAttributes() *CheckoutComGatewayDataAttributes`

NewCheckoutComGatewayDataAttributes instantiates a new CheckoutComGatewayDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutComGatewayDataAttributesWithDefaults

`func NewCheckoutComGatewayDataAttributesWithDefaults() *CheckoutComGatewayDataAttributes`

NewCheckoutComGatewayDataAttributesWithDefaults instantiates a new CheckoutComGatewayDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CheckoutComGatewayDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckoutComGatewayDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckoutComGatewayDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckoutComGatewayDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *CheckoutComGatewayDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutComGatewayDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutComGatewayDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckoutComGatewayDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CheckoutComGatewayDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckoutComGatewayDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckoutComGatewayDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CheckoutComGatewayDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CheckoutComGatewayDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CheckoutComGatewayDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CheckoutComGatewayDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CheckoutComGatewayDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *CheckoutComGatewayDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CheckoutComGatewayDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CheckoutComGatewayDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CheckoutComGatewayDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *CheckoutComGatewayDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *CheckoutComGatewayDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *CheckoutComGatewayDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *CheckoutComGatewayDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *CheckoutComGatewayDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckoutComGatewayDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckoutComGatewayDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckoutComGatewayDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetWebhookEndpointId

`func (o *CheckoutComGatewayDataAttributes) GetWebhookEndpointId() string`

GetWebhookEndpointId returns the WebhookEndpointId field if non-nil, zero value otherwise.

### GetWebhookEndpointIdOk

`func (o *CheckoutComGatewayDataAttributes) GetWebhookEndpointIdOk() (*string, bool)`

GetWebhookEndpointIdOk returns a tuple with the WebhookEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpointId

`func (o *CheckoutComGatewayDataAttributes) SetWebhookEndpointId(v string)`

SetWebhookEndpointId sets WebhookEndpointId field to given value.

### HasWebhookEndpointId

`func (o *CheckoutComGatewayDataAttributes) HasWebhookEndpointId() bool`

HasWebhookEndpointId returns a boolean if a field has been set.

### GetWebhookEndpointSecret

`func (o *CheckoutComGatewayDataAttributes) GetWebhookEndpointSecret() string`

GetWebhookEndpointSecret returns the WebhookEndpointSecret field if non-nil, zero value otherwise.

### GetWebhookEndpointSecretOk

`func (o *CheckoutComGatewayDataAttributes) GetWebhookEndpointSecretOk() (*string, bool)`

GetWebhookEndpointSecretOk returns a tuple with the WebhookEndpointSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpointSecret

`func (o *CheckoutComGatewayDataAttributes) SetWebhookEndpointSecret(v string)`

SetWebhookEndpointSecret sets WebhookEndpointSecret field to given value.

### HasWebhookEndpointSecret

`func (o *CheckoutComGatewayDataAttributes) HasWebhookEndpointSecret() bool`

HasWebhookEndpointSecret returns a boolean if a field has been set.

### GetWebhookEndpointUrl

`func (o *CheckoutComGatewayDataAttributes) GetWebhookEndpointUrl() string`

GetWebhookEndpointUrl returns the WebhookEndpointUrl field if non-nil, zero value otherwise.

### GetWebhookEndpointUrlOk

`func (o *CheckoutComGatewayDataAttributes) GetWebhookEndpointUrlOk() (*string, bool)`

GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpointUrl

`func (o *CheckoutComGatewayDataAttributes) SetWebhookEndpointUrl(v string)`

SetWebhookEndpointUrl sets WebhookEndpointUrl field to given value.

### HasWebhookEndpointUrl

`func (o *CheckoutComGatewayDataAttributes) HasWebhookEndpointUrl() bool`

HasWebhookEndpointUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


