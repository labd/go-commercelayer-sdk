# POSTAvalaraAccounts201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The tax calculator&#39;s internal name. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**Username** | **interface{}** | The Avalara account username. | 
**Password** | **interface{}** | The Avalara account password. | 
**CompanyCode** | **interface{}** | The Avalara company code. | 
**CommitInvoice** | Pointer to **interface{}** | Indicates if the transaction will be recorded and visible on the Avalara website. | [optional] 
**Ddp** | Pointer to **interface{}** | Indicates if the seller is responsible for paying/remitting the customs duty &amp; import tax to the customs authorities. | [optional] 

## Methods

### NewPOSTAvalaraAccounts201ResponseDataAttributes

`func NewPOSTAvalaraAccounts201ResponseDataAttributes(name interface{}, username interface{}, password interface{}, companyCode interface{}, ) *POSTAvalaraAccounts201ResponseDataAttributes`

NewPOSTAvalaraAccounts201ResponseDataAttributes instantiates a new POSTAvalaraAccounts201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTAvalaraAccounts201ResponseDataAttributesWithDefaults

`func NewPOSTAvalaraAccounts201ResponseDataAttributesWithDefaults() *POSTAvalaraAccounts201ResponseDataAttributes`

NewPOSTAvalaraAccounts201ResponseDataAttributesWithDefaults instantiates a new POSTAvalaraAccounts201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTAvalaraAccounts201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTAvalaraAccounts201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTAvalaraAccounts201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTAvalaraAccounts201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetUsername

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetUsername() interface{}`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetUsernameOk() (*interface{}, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetUsername(v interface{})`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *POSTAvalaraAccounts201ResponseDataAttributes) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetPassword() interface{}`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetPasswordOk() (*interface{}, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetPassword(v interface{})`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *POSTAvalaraAccounts201ResponseDataAttributes) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetCompanyCode

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetCompanyCode() interface{}`

GetCompanyCode returns the CompanyCode field if non-nil, zero value otherwise.

### GetCompanyCodeOk

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetCompanyCodeOk() (*interface{}, bool)`

GetCompanyCodeOk returns a tuple with the CompanyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCode

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetCompanyCode(v interface{})`

SetCompanyCode sets CompanyCode field to given value.


### SetCompanyCodeNil

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetCompanyCodeNil(b bool)`

 SetCompanyCodeNil sets the value for CompanyCode to be an explicit nil

### UnsetCompanyCode
`func (o *POSTAvalaraAccounts201ResponseDataAttributes) UnsetCompanyCode()`

UnsetCompanyCode ensures that no value is present for CompanyCode, not even an explicit nil
### GetCommitInvoice

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetCommitInvoice() interface{}`

GetCommitInvoice returns the CommitInvoice field if non-nil, zero value otherwise.

### GetCommitInvoiceOk

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetCommitInvoiceOk() (*interface{}, bool)`

GetCommitInvoiceOk returns a tuple with the CommitInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitInvoice

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetCommitInvoice(v interface{})`

SetCommitInvoice sets CommitInvoice field to given value.

### HasCommitInvoice

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) HasCommitInvoice() bool`

HasCommitInvoice returns a boolean if a field has been set.

### SetCommitInvoiceNil

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetCommitInvoiceNil(b bool)`

 SetCommitInvoiceNil sets the value for CommitInvoice to be an explicit nil

### UnsetCommitInvoice
`func (o *POSTAvalaraAccounts201ResponseDataAttributes) UnsetCommitInvoice()`

UnsetCommitInvoice ensures that no value is present for CommitInvoice, not even an explicit nil
### GetDdp

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetDdp() interface{}`

GetDdp returns the Ddp field if non-nil, zero value otherwise.

### GetDdpOk

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetDdpOk() (*interface{}, bool)`

GetDdpOk returns a tuple with the Ddp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdp

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetDdp(v interface{})`

SetDdp sets Ddp field to given value.

### HasDdp

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) HasDdp() bool`

HasDdp returns a boolean if a field has been set.

### SetDdpNil

`func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetDdpNil(b bool)`

 SetDdpNil sets the value for Ddp to be an explicit nil

### UnsetDdp
`func (o *POSTAvalaraAccounts201ResponseDataAttributes) UnsetDdp()`

UnsetDdp ensures that no value is present for Ddp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


