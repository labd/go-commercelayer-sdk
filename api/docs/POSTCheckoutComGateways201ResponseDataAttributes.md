# POSTCheckoutComGateways201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The payment gateway&#39;s internal name. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**SecretKey** | **interface{}** | The gateway secret key. | 
**PublicKey** | **interface{}** | The gateway public key. | 

## Methods

### NewPOSTCheckoutComGateways201ResponseDataAttributes

`func NewPOSTCheckoutComGateways201ResponseDataAttributes(name interface{}, secretKey interface{}, publicKey interface{}, ) *POSTCheckoutComGateways201ResponseDataAttributes`

NewPOSTCheckoutComGateways201ResponseDataAttributes instantiates a new POSTCheckoutComGateways201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCheckoutComGateways201ResponseDataAttributesWithDefaults

`func NewPOSTCheckoutComGateways201ResponseDataAttributesWithDefaults() *POSTCheckoutComGateways201ResponseDataAttributes`

NewPOSTCheckoutComGateways201ResponseDataAttributesWithDefaults instantiates a new POSTCheckoutComGateways201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTCheckoutComGateways201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTCheckoutComGateways201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTCheckoutComGateways201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTCheckoutComGateways201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSecretKey

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetSecretKey() interface{}`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetSecretKeyOk() (*interface{}, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetSecretKey(v interface{})`

SetSecretKey sets SecretKey field to given value.


### SetSecretKeyNil

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *POSTCheckoutComGateways201ResponseDataAttributes) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetPublicKey

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetPublicKey() interface{}`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetPublicKeyOk() (*interface{}, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetPublicKey(v interface{})`

SetPublicKey sets PublicKey field to given value.


### SetPublicKeyNil

`func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *POSTCheckoutComGateways201ResponseDataAttributes) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


