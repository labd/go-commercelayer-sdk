# GETStockLocations200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the stock location (numeric) | [optional] 
**Name** | Pointer to **interface{}** | The stock location&#39;s internal name. | [optional] 
**LabelFormat** | Pointer to **interface{}** | The shipping label format for this stock location. Can be one of &#39;PDF&#39;, &#39;ZPL&#39;, &#39;EPL2&#39;, or &#39;PNG&#39; | [optional] 
**SuppressEtd** | Pointer to **interface{}** | Flag it if you want to skip the electronic invoice creation when generating the customs info for this stock location shipments. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETStockLocations200ResponseDataInnerAttributes

`func NewGETStockLocations200ResponseDataInnerAttributes() *GETStockLocations200ResponseDataInnerAttributes`

NewGETStockLocations200ResponseDataInnerAttributes instantiates a new GETStockLocations200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETStockLocations200ResponseDataInnerAttributesWithDefaults

`func NewGETStockLocations200ResponseDataInnerAttributesWithDefaults() *GETStockLocations200ResponseDataInnerAttributes`

NewGETStockLocations200ResponseDataInnerAttributesWithDefaults instantiates a new GETStockLocations200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETStockLocations200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETStockLocations200ResponseDataInnerAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetName

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETStockLocations200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETStockLocations200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLabelFormat

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetLabelFormat() interface{}`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetLabelFormatOk() (*interface{}, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetLabelFormat(v interface{})`

SetLabelFormat sets LabelFormat field to given value.

### HasLabelFormat

`func (o *GETStockLocations200ResponseDataInnerAttributes) HasLabelFormat() bool`

HasLabelFormat returns a boolean if a field has been set.

### SetLabelFormatNil

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetLabelFormatNil(b bool)`

 SetLabelFormatNil sets the value for LabelFormat to be an explicit nil

### UnsetLabelFormat
`func (o *GETStockLocations200ResponseDataInnerAttributes) UnsetLabelFormat()`

UnsetLabelFormat ensures that no value is present for LabelFormat, not even an explicit nil
### GetSuppressEtd

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetSuppressEtd() interface{}`

GetSuppressEtd returns the SuppressEtd field if non-nil, zero value otherwise.

### GetSuppressEtdOk

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetSuppressEtdOk() (*interface{}, bool)`

GetSuppressEtdOk returns a tuple with the SuppressEtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressEtd

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetSuppressEtd(v interface{})`

SetSuppressEtd sets SuppressEtd field to given value.

### HasSuppressEtd

`func (o *GETStockLocations200ResponseDataInnerAttributes) HasSuppressEtd() bool`

HasSuppressEtd returns a boolean if a field has been set.

### SetSuppressEtdNil

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetSuppressEtdNil(b bool)`

 SetSuppressEtdNil sets the value for SuppressEtd to be an explicit nil

### UnsetSuppressEtd
`func (o *GETStockLocations200ResponseDataInnerAttributes) UnsetSuppressEtd()`

UnsetSuppressEtd ensures that no value is present for SuppressEtd, not even an explicit nil
### GetCreatedAt

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETStockLocations200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETStockLocations200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETStockLocations200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETStockLocations200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETStockLocations200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETStockLocations200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETStockLocations200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETStockLocations200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETStockLocations200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETStockLocations200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETStockLocations200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETStockLocations200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


