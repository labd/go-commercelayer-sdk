# PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**SecretKey** | Pointer to **string** | The gateway secret key. | [optional] 
**PublicKey** | Pointer to **string** | The gateway public key. | [optional] 

## Methods

### NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes

`func NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes() *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes`

NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes instantiates a new PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributesWithDefaults

`func NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributesWithDefaults() *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes`

NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReference

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSecretKey

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


