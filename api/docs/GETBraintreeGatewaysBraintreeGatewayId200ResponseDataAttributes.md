# GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The payment gateway&#39;s internal name. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**DescriptorName** | Pointer to **interface{}** | The dynamic descriptor name. Must be composed by business name (3, 7 or 12 chars), an asterisk (*) and the product name (18, 14 or 9 chars), for a total length of 22 chars. | [optional] 
**DescriptorPhone** | Pointer to **interface{}** | The dynamic descriptor phone number. Must be 10-14 characters and can only contain numbers, dashes, parentheses and periods. | [optional] 
**DescriptorUrl** | Pointer to **interface{}** | The dynamic descriptor URL. | [optional] 
**WebhookEndpointUrl** | Pointer to **interface{}** | The gateway webhook URL, generated automatically. | [optional] 

## Methods

### NewGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes

`func NewGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes() *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes`

NewGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes instantiates a new GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributesWithDefaults

`func NewGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributesWithDefaults() *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes`

NewGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributesWithDefaults instantiates a new GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedAt

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDescriptorName

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorName() interface{}`

GetDescriptorName returns the DescriptorName field if non-nil, zero value otherwise.

### GetDescriptorNameOk

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorNameOk() (*interface{}, bool)`

GetDescriptorNameOk returns a tuple with the DescriptorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorName

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorName(v interface{})`

SetDescriptorName sets DescriptorName field to given value.

### HasDescriptorName

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorName() bool`

HasDescriptorName returns a boolean if a field has been set.

### SetDescriptorNameNil

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorNameNil(b bool)`

 SetDescriptorNameNil sets the value for DescriptorName to be an explicit nil

### UnsetDescriptorName
`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetDescriptorName()`

UnsetDescriptorName ensures that no value is present for DescriptorName, not even an explicit nil
### GetDescriptorPhone

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorPhone() interface{}`

GetDescriptorPhone returns the DescriptorPhone field if non-nil, zero value otherwise.

### GetDescriptorPhoneOk

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorPhoneOk() (*interface{}, bool)`

GetDescriptorPhoneOk returns a tuple with the DescriptorPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorPhone

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorPhone(v interface{})`

SetDescriptorPhone sets DescriptorPhone field to given value.

### HasDescriptorPhone

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorPhone() bool`

HasDescriptorPhone returns a boolean if a field has been set.

### SetDescriptorPhoneNil

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorPhoneNil(b bool)`

 SetDescriptorPhoneNil sets the value for DescriptorPhone to be an explicit nil

### UnsetDescriptorPhone
`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetDescriptorPhone()`

UnsetDescriptorPhone ensures that no value is present for DescriptorPhone, not even an explicit nil
### GetDescriptorUrl

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorUrl() interface{}`

GetDescriptorUrl returns the DescriptorUrl field if non-nil, zero value otherwise.

### GetDescriptorUrlOk

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorUrlOk() (*interface{}, bool)`

GetDescriptorUrlOk returns a tuple with the DescriptorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorUrl

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorUrl(v interface{})`

SetDescriptorUrl sets DescriptorUrl field to given value.

### HasDescriptorUrl

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorUrl() bool`

HasDescriptorUrl returns a boolean if a field has been set.

### SetDescriptorUrlNil

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorUrlNil(b bool)`

 SetDescriptorUrlNil sets the value for DescriptorUrl to be an explicit nil

### UnsetDescriptorUrl
`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetDescriptorUrl()`

UnsetDescriptorUrl ensures that no value is present for DescriptorUrl, not even an explicit nil
### GetWebhookEndpointUrl

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetWebhookEndpointUrl() interface{}`

GetWebhookEndpointUrl returns the WebhookEndpointUrl field if non-nil, zero value otherwise.

### GetWebhookEndpointUrlOk

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetWebhookEndpointUrlOk() (*interface{}, bool)`

GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpointUrl

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetWebhookEndpointUrl(v interface{})`

SetWebhookEndpointUrl sets WebhookEndpointUrl field to given value.

### HasWebhookEndpointUrl

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasWebhookEndpointUrl() bool`

HasWebhookEndpointUrl returns a boolean if a field has been set.

### SetWebhookEndpointUrlNil

`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetWebhookEndpointUrlNil(b bool)`

 SetWebhookEndpointUrlNil sets the value for WebhookEndpointUrl to be an explicit nil

### UnsetWebhookEndpointUrl
`func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetWebhookEndpointUrl()`

UnsetWebhookEndpointUrl ensures that no value is present for WebhookEndpointUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


