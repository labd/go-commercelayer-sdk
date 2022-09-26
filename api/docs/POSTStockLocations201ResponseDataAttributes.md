# POSTStockLocations201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The stock location&#39;s internal name. | 
**LabelFormat** | Pointer to **string** | The shipping label format for this stock location. Can be one of &#39;PDF&#39;, &#39;ZPL&#39;, &#39;EPL2&#39;, or &#39;PNG&#39; | [optional] 
**SuppressEtd** | Pointer to **bool** | Flag it if you want to skip the electronic invoice creation when generating the customs info for this stock location shipments. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTStockLocations201ResponseDataAttributes

`func NewPOSTStockLocations201ResponseDataAttributes(name string, ) *POSTStockLocations201ResponseDataAttributes`

NewPOSTStockLocations201ResponseDataAttributes instantiates a new POSTStockLocations201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockLocations201ResponseDataAttributesWithDefaults

`func NewPOSTStockLocations201ResponseDataAttributesWithDefaults() *POSTStockLocations201ResponseDataAttributes`

NewPOSTStockLocations201ResponseDataAttributesWithDefaults instantiates a new POSTStockLocations201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTStockLocations201ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTStockLocations201ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTStockLocations201ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetLabelFormat

`func (o *POSTStockLocations201ResponseDataAttributes) GetLabelFormat() string`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *POSTStockLocations201ResponseDataAttributes) GetLabelFormatOk() (*string, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *POSTStockLocations201ResponseDataAttributes) SetLabelFormat(v string)`

SetLabelFormat sets LabelFormat field to given value.

### HasLabelFormat

`func (o *POSTStockLocations201ResponseDataAttributes) HasLabelFormat() bool`

HasLabelFormat returns a boolean if a field has been set.

### GetSuppressEtd

`func (o *POSTStockLocations201ResponseDataAttributes) GetSuppressEtd() bool`

GetSuppressEtd returns the SuppressEtd field if non-nil, zero value otherwise.

### GetSuppressEtdOk

`func (o *POSTStockLocations201ResponseDataAttributes) GetSuppressEtdOk() (*bool, bool)`

GetSuppressEtdOk returns a tuple with the SuppressEtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressEtd

`func (o *POSTStockLocations201ResponseDataAttributes) SetSuppressEtd(v bool)`

SetSuppressEtd sets SuppressEtd field to given value.

### HasSuppressEtd

`func (o *POSTStockLocations201ResponseDataAttributes) HasSuppressEtd() bool`

HasSuppressEtd returns a boolean if a field has been set.

### GetReference

`func (o *POSTStockLocations201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTStockLocations201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTStockLocations201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTStockLocations201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTStockLocations201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTStockLocations201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTStockLocations201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTStockLocations201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTStockLocations201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTStockLocations201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTStockLocations201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTStockLocations201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


