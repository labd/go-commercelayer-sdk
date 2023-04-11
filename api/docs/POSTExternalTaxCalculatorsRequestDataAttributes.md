# POSTExternalTaxCalculatorsRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The tax calculator&#39;s internal name. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**TaxCalculatorUrl** | **interface{}** | The URL to the service that will compute the taxes. | 

## Methods

### NewPOSTExternalTaxCalculatorsRequestDataAttributes

`func NewPOSTExternalTaxCalculatorsRequestDataAttributes(name interface{}, taxCalculatorUrl interface{}, ) *POSTExternalTaxCalculatorsRequestDataAttributes`

NewPOSTExternalTaxCalculatorsRequestDataAttributes instantiates a new POSTExternalTaxCalculatorsRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTExternalTaxCalculatorsRequestDataAttributesWithDefaults

`func NewPOSTExternalTaxCalculatorsRequestDataAttributesWithDefaults() *POSTExternalTaxCalculatorsRequestDataAttributes`

NewPOSTExternalTaxCalculatorsRequestDataAttributesWithDefaults instantiates a new POSTExternalTaxCalculatorsRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetTaxCalculatorUrl

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) GetTaxCalculatorUrl() interface{}`

GetTaxCalculatorUrl returns the TaxCalculatorUrl field if non-nil, zero value otherwise.

### GetTaxCalculatorUrlOk

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) GetTaxCalculatorUrlOk() (*interface{}, bool)`

GetTaxCalculatorUrlOk returns a tuple with the TaxCalculatorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculatorUrl

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) SetTaxCalculatorUrl(v interface{})`

SetTaxCalculatorUrl sets TaxCalculatorUrl field to given value.


### SetTaxCalculatorUrlNil

`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) SetTaxCalculatorUrlNil(b bool)`

 SetTaxCalculatorUrlNil sets the value for TaxCalculatorUrl to be an explicit nil

### UnsetTaxCalculatorUrl
`func (o *POSTExternalTaxCalculatorsRequestDataAttributes) UnsetTaxCalculatorUrl()`

UnsetTaxCalculatorUrl ensures that no value is present for TaxCalculatorUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


