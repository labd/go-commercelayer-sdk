# PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**AutoPayments** | Pointer to **interface{}** | Indicates if the gateway will accept payment methods enabled in the Stripe dashboard. | [optional] 

## Methods

### NewPATCHStripeGatewaysStripeGatewayIdRequestDataAttributes

`func NewPATCHStripeGatewaysStripeGatewayIdRequestDataAttributes() *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes`

NewPATCHStripeGatewaysStripeGatewayIdRequestDataAttributes instantiates a new PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHStripeGatewaysStripeGatewayIdRequestDataAttributesWithDefaults

`func NewPATCHStripeGatewaysStripeGatewayIdRequestDataAttributesWithDefaults() *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes`

NewPATCHStripeGatewaysStripeGatewayIdRequestDataAttributesWithDefaults instantiates a new PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAutoPayments

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetAutoPayments() interface{}`

GetAutoPayments returns the AutoPayments field if non-nil, zero value otherwise.

### GetAutoPaymentsOk

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetAutoPaymentsOk() (*interface{}, bool)`

GetAutoPaymentsOk returns a tuple with the AutoPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPayments

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetAutoPayments(v interface{})`

SetAutoPayments sets AutoPayments field to given value.

### HasAutoPayments

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) HasAutoPayments() bool`

HasAutoPayments returns a boolean if a field has been set.

### SetAutoPaymentsNil

`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetAutoPaymentsNil(b bool)`

 SetAutoPaymentsNil sets the value for AutoPayments to be an explicit nil

### UnsetAutoPayments
`func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) UnsetAutoPayments()`

UnsetAutoPayments ensures that no value is present for AutoPayments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


