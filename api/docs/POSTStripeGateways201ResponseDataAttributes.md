# POSTStripeGateways201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The payment gateway&#39;s internal name. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**Login** | **interface{}** | The gateway login. | 
**PublishableKey** | Pointer to **interface{}** | The gateway publishable API key. | [optional] 
**AutoPayments** | Pointer to **interface{}** | Indicates if the gateway will accept payment methods enabled in the Stripe dashboard. | [optional] 

## Methods

### NewPOSTStripeGateways201ResponseDataAttributes

`func NewPOSTStripeGateways201ResponseDataAttributes(name interface{}, login interface{}, ) *POSTStripeGateways201ResponseDataAttributes`

NewPOSTStripeGateways201ResponseDataAttributes instantiates a new POSTStripeGateways201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStripeGateways201ResponseDataAttributesWithDefaults

`func NewPOSTStripeGateways201ResponseDataAttributesWithDefaults() *POSTStripeGateways201ResponseDataAttributes`

NewPOSTStripeGateways201ResponseDataAttributesWithDefaults instantiates a new POSTStripeGateways201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTStripeGateways201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTStripeGateways201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTStripeGateways201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTStripeGateways201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTStripeGateways201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *POSTStripeGateways201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTStripeGateways201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTStripeGateways201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTStripeGateways201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTStripeGateways201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTStripeGateways201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTStripeGateways201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTStripeGateways201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTStripeGateways201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTStripeGateways201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTStripeGateways201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTStripeGateways201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTStripeGateways201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTStripeGateways201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTStripeGateways201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTStripeGateways201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTStripeGateways201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTStripeGateways201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetLogin

`func (o *POSTStripeGateways201ResponseDataAttributes) GetLogin() interface{}`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *POSTStripeGateways201ResponseDataAttributes) GetLoginOk() (*interface{}, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *POSTStripeGateways201ResponseDataAttributes) SetLogin(v interface{})`

SetLogin sets Login field to given value.


### SetLoginNil

`func (o *POSTStripeGateways201ResponseDataAttributes) SetLoginNil(b bool)`

 SetLoginNil sets the value for Login to be an explicit nil

### UnsetLogin
`func (o *POSTStripeGateways201ResponseDataAttributes) UnsetLogin()`

UnsetLogin ensures that no value is present for Login, not even an explicit nil
### GetPublishableKey

`func (o *POSTStripeGateways201ResponseDataAttributes) GetPublishableKey() interface{}`

GetPublishableKey returns the PublishableKey field if non-nil, zero value otherwise.

### GetPublishableKeyOk

`func (o *POSTStripeGateways201ResponseDataAttributes) GetPublishableKeyOk() (*interface{}, bool)`

GetPublishableKeyOk returns a tuple with the PublishableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishableKey

`func (o *POSTStripeGateways201ResponseDataAttributes) SetPublishableKey(v interface{})`

SetPublishableKey sets PublishableKey field to given value.

### HasPublishableKey

`func (o *POSTStripeGateways201ResponseDataAttributes) HasPublishableKey() bool`

HasPublishableKey returns a boolean if a field has been set.

### SetPublishableKeyNil

`func (o *POSTStripeGateways201ResponseDataAttributes) SetPublishableKeyNil(b bool)`

 SetPublishableKeyNil sets the value for PublishableKey to be an explicit nil

### UnsetPublishableKey
`func (o *POSTStripeGateways201ResponseDataAttributes) UnsetPublishableKey()`

UnsetPublishableKey ensures that no value is present for PublishableKey, not even an explicit nil
### GetAutoPayments

`func (o *POSTStripeGateways201ResponseDataAttributes) GetAutoPayments() interface{}`

GetAutoPayments returns the AutoPayments field if non-nil, zero value otherwise.

### GetAutoPaymentsOk

`func (o *POSTStripeGateways201ResponseDataAttributes) GetAutoPaymentsOk() (*interface{}, bool)`

GetAutoPaymentsOk returns a tuple with the AutoPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPayments

`func (o *POSTStripeGateways201ResponseDataAttributes) SetAutoPayments(v interface{})`

SetAutoPayments sets AutoPayments field to given value.

### HasAutoPayments

`func (o *POSTStripeGateways201ResponseDataAttributes) HasAutoPayments() bool`

HasAutoPayments returns a boolean if a field has been set.

### SetAutoPaymentsNil

`func (o *POSTStripeGateways201ResponseDataAttributes) SetAutoPaymentsNil(b bool)`

 SetAutoPaymentsNil sets the value for AutoPayments to be an explicit nil

### UnsetAutoPayments
`func (o *POSTStripeGateways201ResponseDataAttributes) UnsetAutoPayments()`

UnsetAutoPayments ensures that no value is present for AutoPayments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


