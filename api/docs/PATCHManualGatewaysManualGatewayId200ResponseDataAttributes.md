# PATCHManualGatewaysManualGatewayId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**RequireCapture** | Pointer to **bool** | Indicates if the gateway requires to captured. | [optional] 

## Methods

### NewPATCHManualGatewaysManualGatewayId200ResponseDataAttributes

`func NewPATCHManualGatewaysManualGatewayId200ResponseDataAttributes() *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes`

NewPATCHManualGatewaysManualGatewayId200ResponseDataAttributes instantiates a new PATCHManualGatewaysManualGatewayId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHManualGatewaysManualGatewayId200ResponseDataAttributesWithDefaults

`func NewPATCHManualGatewaysManualGatewayId200ResponseDataAttributesWithDefaults() *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes`

NewPATCHManualGatewaysManualGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHManualGatewaysManualGatewayId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReference

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRequireCapture

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) GetRequireCapture() bool`

GetRequireCapture returns the RequireCapture field if non-nil, zero value otherwise.

### GetRequireCaptureOk

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) GetRequireCaptureOk() (*bool, bool)`

GetRequireCaptureOk returns a tuple with the RequireCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCapture

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) SetRequireCapture(v bool)`

SetRequireCapture sets RequireCapture field to given value.

### HasRequireCapture

`func (o *PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) HasRequireCapture() bool`

HasRequireCapture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


