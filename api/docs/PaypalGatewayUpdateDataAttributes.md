# PaypalGatewayUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**ClientId** | Pointer to **string** | The gateway client ID. | [optional] 
**ClientSecret** | Pointer to **string** | The gateway client secret. | [optional] 

## Methods

### NewPaypalGatewayUpdateDataAttributes

`func NewPaypalGatewayUpdateDataAttributes() *PaypalGatewayUpdateDataAttributes`

NewPaypalGatewayUpdateDataAttributes instantiates a new PaypalGatewayUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalGatewayUpdateDataAttributesWithDefaults

`func NewPaypalGatewayUpdateDataAttributesWithDefaults() *PaypalGatewayUpdateDataAttributes`

NewPaypalGatewayUpdateDataAttributesWithDefaults instantiates a new PaypalGatewayUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PaypalGatewayUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaypalGatewayUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaypalGatewayUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaypalGatewayUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReference

`func (o *PaypalGatewayUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaypalGatewayUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaypalGatewayUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaypalGatewayUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PaypalGatewayUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PaypalGatewayUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PaypalGatewayUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PaypalGatewayUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PaypalGatewayUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaypalGatewayUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaypalGatewayUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaypalGatewayUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetClientId

`func (o *PaypalGatewayUpdateDataAttributes) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PaypalGatewayUpdateDataAttributes) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PaypalGatewayUpdateDataAttributes) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PaypalGatewayUpdateDataAttributes) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *PaypalGatewayUpdateDataAttributes) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PaypalGatewayUpdateDataAttributes) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PaypalGatewayUpdateDataAttributes) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PaypalGatewayUpdateDataAttributes) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


