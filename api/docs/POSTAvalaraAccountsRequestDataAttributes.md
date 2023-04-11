# POSTAvalaraAccountsRequestDataAttributes

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

### NewPOSTAvalaraAccountsRequestDataAttributes

`func NewPOSTAvalaraAccountsRequestDataAttributes(name interface{}, username interface{}, password interface{}, companyCode interface{}, ) *POSTAvalaraAccountsRequestDataAttributes`

NewPOSTAvalaraAccountsRequestDataAttributes instantiates a new POSTAvalaraAccountsRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTAvalaraAccountsRequestDataAttributesWithDefaults

`func NewPOSTAvalaraAccountsRequestDataAttributesWithDefaults() *POSTAvalaraAccountsRequestDataAttributes`

NewPOSTAvalaraAccountsRequestDataAttributesWithDefaults instantiates a new POSTAvalaraAccountsRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTAvalaraAccountsRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTAvalaraAccountsRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTAvalaraAccountsRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTAvalaraAccountsRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTAvalaraAccountsRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTAvalaraAccountsRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTAvalaraAccountsRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetUsername

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetUsername() interface{}`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetUsernameOk() (*interface{}, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetUsername(v interface{})`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *POSTAvalaraAccountsRequestDataAttributes) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetPassword() interface{}`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetPasswordOk() (*interface{}, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetPassword(v interface{})`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *POSTAvalaraAccountsRequestDataAttributes) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetCompanyCode

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetCompanyCode() interface{}`

GetCompanyCode returns the CompanyCode field if non-nil, zero value otherwise.

### GetCompanyCodeOk

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetCompanyCodeOk() (*interface{}, bool)`

GetCompanyCodeOk returns a tuple with the CompanyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCode

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetCompanyCode(v interface{})`

SetCompanyCode sets CompanyCode field to given value.


### SetCompanyCodeNil

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetCompanyCodeNil(b bool)`

 SetCompanyCodeNil sets the value for CompanyCode to be an explicit nil

### UnsetCompanyCode
`func (o *POSTAvalaraAccountsRequestDataAttributes) UnsetCompanyCode()`

UnsetCompanyCode ensures that no value is present for CompanyCode, not even an explicit nil
### GetCommitInvoice

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetCommitInvoice() interface{}`

GetCommitInvoice returns the CommitInvoice field if non-nil, zero value otherwise.

### GetCommitInvoiceOk

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetCommitInvoiceOk() (*interface{}, bool)`

GetCommitInvoiceOk returns a tuple with the CommitInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitInvoice

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetCommitInvoice(v interface{})`

SetCommitInvoice sets CommitInvoice field to given value.

### HasCommitInvoice

`func (o *POSTAvalaraAccountsRequestDataAttributes) HasCommitInvoice() bool`

HasCommitInvoice returns a boolean if a field has been set.

### SetCommitInvoiceNil

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetCommitInvoiceNil(b bool)`

 SetCommitInvoiceNil sets the value for CommitInvoice to be an explicit nil

### UnsetCommitInvoice
`func (o *POSTAvalaraAccountsRequestDataAttributes) UnsetCommitInvoice()`

UnsetCommitInvoice ensures that no value is present for CommitInvoice, not even an explicit nil
### GetDdp

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetDdp() interface{}`

GetDdp returns the Ddp field if non-nil, zero value otherwise.

### GetDdpOk

`func (o *POSTAvalaraAccountsRequestDataAttributes) GetDdpOk() (*interface{}, bool)`

GetDdpOk returns a tuple with the Ddp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdp

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetDdp(v interface{})`

SetDdp sets Ddp field to given value.

### HasDdp

`func (o *POSTAvalaraAccountsRequestDataAttributes) HasDdp() bool`

HasDdp returns a boolean if a field has been set.

### SetDdpNil

`func (o *POSTAvalaraAccountsRequestDataAttributes) SetDdpNil(b bool)`

 SetDdpNil sets the value for Ddp to be an explicit nil

### UnsetDdp
`func (o *POSTAvalaraAccountsRequestDataAttributes) UnsetDdp()`

UnsetDdp ensures that no value is present for Ddp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


