# GETBraintreeGateways200ResponseDataInnerAttributes

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
**DescriptorName** | Pointer to **string** | The dynamic descriptor name. Must be composed by business name (3, 7 or 12 chars), an asterisk (*) and the product name (18, 14 or 9 chars), for a total length of 22 chars. | [optional] 
**DescriptorPhone** | Pointer to **string** | The dynamic descriptor phone number. Must be 10-14 characters and can only contain numbers, dashes, parentheses and periods. | [optional] 
**DescriptorUrl** | Pointer to **string** | The dynamic descriptor URL. | [optional] 
**WebhookEndpointUrl** | Pointer to **string** | The gateway webhook URL, generated automatically. | [optional] 

## Methods

### NewGETBraintreeGateways200ResponseDataInnerAttributes

`func NewGETBraintreeGateways200ResponseDataInnerAttributes() *GETBraintreeGateways200ResponseDataInnerAttributes`

NewGETBraintreeGateways200ResponseDataInnerAttributes instantiates a new GETBraintreeGateways200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETBraintreeGateways200ResponseDataInnerAttributesWithDefaults

`func NewGETBraintreeGateways200ResponseDataInnerAttributesWithDefaults() *GETBraintreeGateways200ResponseDataInnerAttributes`

NewGETBraintreeGateways200ResponseDataInnerAttributesWithDefaults instantiates a new GETBraintreeGateways200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDescriptorName

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorName() string`

GetDescriptorName returns the DescriptorName field if non-nil, zero value otherwise.

### GetDescriptorNameOk

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorNameOk() (*string, bool)`

GetDescriptorNameOk returns a tuple with the DescriptorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorName

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetDescriptorName(v string)`

SetDescriptorName sets DescriptorName field to given value.

### HasDescriptorName

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasDescriptorName() bool`

HasDescriptorName returns a boolean if a field has been set.

### GetDescriptorPhone

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorPhone() string`

GetDescriptorPhone returns the DescriptorPhone field if non-nil, zero value otherwise.

### GetDescriptorPhoneOk

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorPhoneOk() (*string, bool)`

GetDescriptorPhoneOk returns a tuple with the DescriptorPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorPhone

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetDescriptorPhone(v string)`

SetDescriptorPhone sets DescriptorPhone field to given value.

### HasDescriptorPhone

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasDescriptorPhone() bool`

HasDescriptorPhone returns a boolean if a field has been set.

### GetDescriptorUrl

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorUrl() string`

GetDescriptorUrl returns the DescriptorUrl field if non-nil, zero value otherwise.

### GetDescriptorUrlOk

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorUrlOk() (*string, bool)`

GetDescriptorUrlOk returns a tuple with the DescriptorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorUrl

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetDescriptorUrl(v string)`

SetDescriptorUrl sets DescriptorUrl field to given value.

### HasDescriptorUrl

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasDescriptorUrl() bool`

HasDescriptorUrl returns a boolean if a field has been set.

### GetWebhookEndpointUrl

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetWebhookEndpointUrl() string`

GetWebhookEndpointUrl returns the WebhookEndpointUrl field if non-nil, zero value otherwise.

### GetWebhookEndpointUrlOk

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetWebhookEndpointUrlOk() (*string, bool)`

GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpointUrl

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetWebhookEndpointUrl(v string)`

SetWebhookEndpointUrl sets WebhookEndpointUrl field to given value.

### HasWebhookEndpointUrl

`func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasWebhookEndpointUrl() bool`

HasWebhookEndpointUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


