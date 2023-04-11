# POSTStockLocationsRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The stock location&#39;s internal name. | 
**LabelFormat** | Pointer to **interface{}** | The shipping label format for this stock location. Can be one of &#39;PDF&#39;, &#39;ZPL&#39;, &#39;EPL2&#39;, or &#39;PNG&#39; | [optional] 
**SuppressEtd** | Pointer to **interface{}** | Flag it if you want to skip the electronic invoice creation when generating the customs info for this stock location shipments. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTStockLocationsRequestDataAttributes

`func NewPOSTStockLocationsRequestDataAttributes(name interface{}, ) *POSTStockLocationsRequestDataAttributes`

NewPOSTStockLocationsRequestDataAttributes instantiates a new POSTStockLocationsRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockLocationsRequestDataAttributesWithDefaults

`func NewPOSTStockLocationsRequestDataAttributesWithDefaults() *POSTStockLocationsRequestDataAttributes`

NewPOSTStockLocationsRequestDataAttributesWithDefaults instantiates a new POSTStockLocationsRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTStockLocationsRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTStockLocationsRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTStockLocationsRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTStockLocationsRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTStockLocationsRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLabelFormat

`func (o *POSTStockLocationsRequestDataAttributes) GetLabelFormat() interface{}`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *POSTStockLocationsRequestDataAttributes) GetLabelFormatOk() (*interface{}, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *POSTStockLocationsRequestDataAttributes) SetLabelFormat(v interface{})`

SetLabelFormat sets LabelFormat field to given value.

### HasLabelFormat

`func (o *POSTStockLocationsRequestDataAttributes) HasLabelFormat() bool`

HasLabelFormat returns a boolean if a field has been set.

### SetLabelFormatNil

`func (o *POSTStockLocationsRequestDataAttributes) SetLabelFormatNil(b bool)`

 SetLabelFormatNil sets the value for LabelFormat to be an explicit nil

### UnsetLabelFormat
`func (o *POSTStockLocationsRequestDataAttributes) UnsetLabelFormat()`

UnsetLabelFormat ensures that no value is present for LabelFormat, not even an explicit nil
### GetSuppressEtd

`func (o *POSTStockLocationsRequestDataAttributes) GetSuppressEtd() interface{}`

GetSuppressEtd returns the SuppressEtd field if non-nil, zero value otherwise.

### GetSuppressEtdOk

`func (o *POSTStockLocationsRequestDataAttributes) GetSuppressEtdOk() (*interface{}, bool)`

GetSuppressEtdOk returns a tuple with the SuppressEtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressEtd

`func (o *POSTStockLocationsRequestDataAttributes) SetSuppressEtd(v interface{})`

SetSuppressEtd sets SuppressEtd field to given value.

### HasSuppressEtd

`func (o *POSTStockLocationsRequestDataAttributes) HasSuppressEtd() bool`

HasSuppressEtd returns a boolean if a field has been set.

### SetSuppressEtdNil

`func (o *POSTStockLocationsRequestDataAttributes) SetSuppressEtdNil(b bool)`

 SetSuppressEtdNil sets the value for SuppressEtd to be an explicit nil

### UnsetSuppressEtd
`func (o *POSTStockLocationsRequestDataAttributes) UnsetSuppressEtd()`

UnsetSuppressEtd ensures that no value is present for SuppressEtd, not even an explicit nil
### GetReference

`func (o *POSTStockLocationsRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTStockLocationsRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTStockLocationsRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTStockLocationsRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTStockLocationsRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTStockLocationsRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTStockLocationsRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTStockLocationsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTStockLocationsRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTStockLocationsRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTStockLocationsRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTStockLocationsRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTStockLocationsRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTStockLocationsRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTStockLocationsRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTStockLocationsRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTStockLocationsRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTStockLocationsRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


