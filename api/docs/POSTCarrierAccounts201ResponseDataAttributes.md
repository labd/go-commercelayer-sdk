# POSTCarrierAccounts201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The carrier account internal name. | 
**EasypostType** | **interface{}** | The Easypost service carrier type. | 
**Credentials** | **interface{}** | The Easypost carrier accounts credentials fields. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTCarrierAccounts201ResponseDataAttributes

`func NewPOSTCarrierAccounts201ResponseDataAttributes(name interface{}, easypostType interface{}, credentials interface{}, ) *POSTCarrierAccounts201ResponseDataAttributes`

NewPOSTCarrierAccounts201ResponseDataAttributes instantiates a new POSTCarrierAccounts201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCarrierAccounts201ResponseDataAttributesWithDefaults

`func NewPOSTCarrierAccounts201ResponseDataAttributesWithDefaults() *POSTCarrierAccounts201ResponseDataAttributes`

NewPOSTCarrierAccounts201ResponseDataAttributesWithDefaults instantiates a new POSTCarrierAccounts201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTCarrierAccounts201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEasypostType

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetEasypostType() interface{}`

GetEasypostType returns the EasypostType field if non-nil, zero value otherwise.

### GetEasypostTypeOk

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetEasypostTypeOk() (*interface{}, bool)`

GetEasypostTypeOk returns a tuple with the EasypostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostType

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetEasypostType(v interface{})`

SetEasypostType sets EasypostType field to given value.


### SetEasypostTypeNil

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetEasypostTypeNil(b bool)`

 SetEasypostTypeNil sets the value for EasypostType to be an explicit nil

### UnsetEasypostType
`func (o *POSTCarrierAccounts201ResponseDataAttributes) UnsetEasypostType()`

UnsetEasypostType ensures that no value is present for EasypostType, not even an explicit nil
### GetCredentials

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetCredentials() interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetCredentialsOk() (*interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetCredentials(v interface{})`

SetCredentials sets Credentials field to given value.


### SetCredentialsNil

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *POSTCarrierAccounts201ResponseDataAttributes) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetReference

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTCarrierAccounts201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTCarrierAccounts201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTCarrierAccounts201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTCarrierAccounts201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTCarrierAccounts201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTCarrierAccounts201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTCarrierAccounts201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTCarrierAccounts201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


