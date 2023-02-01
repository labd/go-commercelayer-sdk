# PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**AutoPayments** | Pointer to **bool** | Indicates if the gateway will accept payment methods enabled in the Stripe dashboard. | [optional] 

## Methods

### NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes

`func NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes() *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes`

NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes instantiates a new PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributesWithDefaults

`func NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributesWithDefaults() *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes`

NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReference

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAutoPayments

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetAutoPayments() bool`

GetAutoPayments returns the AutoPayments field if non-nil, zero value otherwise.

### GetAutoPaymentsOk

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetAutoPaymentsOk() (*bool, bool)`

GetAutoPaymentsOk returns a tuple with the AutoPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPayments

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetAutoPayments(v bool)`

SetAutoPayments sets AutoPayments field to given value.

### HasAutoPayments

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasAutoPayments() bool`

HasAutoPayments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


