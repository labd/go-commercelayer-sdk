# POSTParcels201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | **float32** | The parcel weight, used to automatically calculate the tax rates from the available carrier accounts. | 
**UnitOfWeight** | **string** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | 
**EelPfc** | Pointer to **string** | When shipping outside the US, you need to provide either an Exemption and Exclusion Legend (EEL) code or a Proof of Filing Citation (PFC). Which you need is based on the value of the goods being shipped. Value can be one of \&quot;EEL\&quot; o \&quot;PFC\&quot;. | [optional] 
**ContentsType** | Pointer to **string** | The type of item you are sending. Can be one of &#39;merchandise&#39;, &#39;gift&#39;, &#39;documents&#39;, &#39;returned_goods&#39;, &#39;sample&#39;, or &#39;other&#39;. | [optional] 
**ContentsExplanation** | Pointer to **string** | If you specify &#39;other&#39; in the &#39;contents_type&#39; attribute, you must supply a brief description in this attribute. | [optional] 
**CustomsCertify** | Pointer to **bool** | Indicates if the provided information is accurate | [optional] 
**CustomsSigner** | Pointer to **string** | This is the name of the person who is certifying that the information provided on the customs form is accurate. Use a name of the person in your organization who is responsible for this. | [optional] 
**NonDeliveryOption** | Pointer to **string** | In case the shipment cannot be delivered, this option tells the carrier what you want to happen to the parcel. You can pass either &#39;return&#39;, or &#39;abandon&#39;. The value defaults to &#39;return&#39;. If you pass &#39;abandon&#39;, you will not receive the parcel back if it cannot be delivered. | [optional] 
**RestrictionType** | Pointer to **string** | Describes if your parcel requires any special treatment or quarantine when entering the country. Can be one of &#39;none&#39;, &#39;other&#39;, &#39;quarantine&#39;, or &#39;sanitary_phytosanitary_inspection&#39;. | [optional] 
**RestrictionComments** | Pointer to **string** | If you specify &#39;other&#39; in the restriction type, you must supply a brief description of what is required. | [optional] 
**CustomsInfoRequired** | Pointer to **bool** | Indicates if the parcel requires customs info to get the shipping rates. | [optional] 
**TrackingNumber** | Pointer to **string** | The tracking number associated to this parcel. | [optional] 
**TrackingStatus** | Pointer to **string** | The tracking status for this parcel, automatically updated in real time by the shipping carrier. | [optional] 
**TrackingStatusDetail** | Pointer to **string** | Additional information about the tracking status, automatically updated in real time by the shipping carrier. | [optional] 
**TrackingStatusUpdatedAt** | Pointer to **string** | Time at which the parcel&#39;s tracking status was last updated. | [optional] 
**TrackingDetails** | Pointer to **string** | The parcel&#39;s full tracking history, automatically updated in real time by the shipping carrier. | [optional] 
**CarrierWeightOz** | Pointer to **string** | The weight of the parcel as measured by the carrier in ounces (if available) | [optional] 
**SignedBy** | Pointer to **string** | The name of the person who signed for the parcel (if available) | [optional] 
**Incoterm** | Pointer to **string** | The type of Incoterm (if available). | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTParcels201ResponseDataAttributes

`func NewPOSTParcels201ResponseDataAttributes(weight float32, unitOfWeight string, ) *POSTParcels201ResponseDataAttributes`

NewPOSTParcels201ResponseDataAttributes instantiates a new POSTParcels201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTParcels201ResponseDataAttributesWithDefaults

`func NewPOSTParcels201ResponseDataAttributesWithDefaults() *POSTParcels201ResponseDataAttributes`

NewPOSTParcels201ResponseDataAttributesWithDefaults instantiates a new POSTParcels201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *POSTParcels201ResponseDataAttributes) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *POSTParcels201ResponseDataAttributes) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *POSTParcels201ResponseDataAttributes) SetWeight(v float32)`

SetWeight sets Weight field to given value.


### GetUnitOfWeight

`func (o *POSTParcels201ResponseDataAttributes) GetUnitOfWeight() string`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *POSTParcels201ResponseDataAttributes) GetUnitOfWeightOk() (*string, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *POSTParcels201ResponseDataAttributes) SetUnitOfWeight(v string)`

SetUnitOfWeight sets UnitOfWeight field to given value.


### GetEelPfc

`func (o *POSTParcels201ResponseDataAttributes) GetEelPfc() string`

GetEelPfc returns the EelPfc field if non-nil, zero value otherwise.

### GetEelPfcOk

`func (o *POSTParcels201ResponseDataAttributes) GetEelPfcOk() (*string, bool)`

GetEelPfcOk returns a tuple with the EelPfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEelPfc

`func (o *POSTParcels201ResponseDataAttributes) SetEelPfc(v string)`

SetEelPfc sets EelPfc field to given value.

### HasEelPfc

`func (o *POSTParcels201ResponseDataAttributes) HasEelPfc() bool`

HasEelPfc returns a boolean if a field has been set.

### GetContentsType

`func (o *POSTParcels201ResponseDataAttributes) GetContentsType() string`

GetContentsType returns the ContentsType field if non-nil, zero value otherwise.

### GetContentsTypeOk

`func (o *POSTParcels201ResponseDataAttributes) GetContentsTypeOk() (*string, bool)`

GetContentsTypeOk returns a tuple with the ContentsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsType

`func (o *POSTParcels201ResponseDataAttributes) SetContentsType(v string)`

SetContentsType sets ContentsType field to given value.

### HasContentsType

`func (o *POSTParcels201ResponseDataAttributes) HasContentsType() bool`

HasContentsType returns a boolean if a field has been set.

### GetContentsExplanation

`func (o *POSTParcels201ResponseDataAttributes) GetContentsExplanation() string`

GetContentsExplanation returns the ContentsExplanation field if non-nil, zero value otherwise.

### GetContentsExplanationOk

`func (o *POSTParcels201ResponseDataAttributes) GetContentsExplanationOk() (*string, bool)`

GetContentsExplanationOk returns a tuple with the ContentsExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsExplanation

`func (o *POSTParcels201ResponseDataAttributes) SetContentsExplanation(v string)`

SetContentsExplanation sets ContentsExplanation field to given value.

### HasContentsExplanation

`func (o *POSTParcels201ResponseDataAttributes) HasContentsExplanation() bool`

HasContentsExplanation returns a boolean if a field has been set.

### GetCustomsCertify

`func (o *POSTParcels201ResponseDataAttributes) GetCustomsCertify() bool`

GetCustomsCertify returns the CustomsCertify field if non-nil, zero value otherwise.

### GetCustomsCertifyOk

`func (o *POSTParcels201ResponseDataAttributes) GetCustomsCertifyOk() (*bool, bool)`

GetCustomsCertifyOk returns a tuple with the CustomsCertify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsCertify

`func (o *POSTParcels201ResponseDataAttributes) SetCustomsCertify(v bool)`

SetCustomsCertify sets CustomsCertify field to given value.

### HasCustomsCertify

`func (o *POSTParcels201ResponseDataAttributes) HasCustomsCertify() bool`

HasCustomsCertify returns a boolean if a field has been set.

### GetCustomsSigner

`func (o *POSTParcels201ResponseDataAttributes) GetCustomsSigner() string`

GetCustomsSigner returns the CustomsSigner field if non-nil, zero value otherwise.

### GetCustomsSignerOk

`func (o *POSTParcels201ResponseDataAttributes) GetCustomsSignerOk() (*string, bool)`

GetCustomsSignerOk returns a tuple with the CustomsSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsSigner

`func (o *POSTParcels201ResponseDataAttributes) SetCustomsSigner(v string)`

SetCustomsSigner sets CustomsSigner field to given value.

### HasCustomsSigner

`func (o *POSTParcels201ResponseDataAttributes) HasCustomsSigner() bool`

HasCustomsSigner returns a boolean if a field has been set.

### GetNonDeliveryOption

`func (o *POSTParcels201ResponseDataAttributes) GetNonDeliveryOption() string`

GetNonDeliveryOption returns the NonDeliveryOption field if non-nil, zero value otherwise.

### GetNonDeliveryOptionOk

`func (o *POSTParcels201ResponseDataAttributes) GetNonDeliveryOptionOk() (*string, bool)`

GetNonDeliveryOptionOk returns a tuple with the NonDeliveryOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDeliveryOption

`func (o *POSTParcels201ResponseDataAttributes) SetNonDeliveryOption(v string)`

SetNonDeliveryOption sets NonDeliveryOption field to given value.

### HasNonDeliveryOption

`func (o *POSTParcels201ResponseDataAttributes) HasNonDeliveryOption() bool`

HasNonDeliveryOption returns a boolean if a field has been set.

### GetRestrictionType

`func (o *POSTParcels201ResponseDataAttributes) GetRestrictionType() string`

GetRestrictionType returns the RestrictionType field if non-nil, zero value otherwise.

### GetRestrictionTypeOk

`func (o *POSTParcels201ResponseDataAttributes) GetRestrictionTypeOk() (*string, bool)`

GetRestrictionTypeOk returns a tuple with the RestrictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionType

`func (o *POSTParcels201ResponseDataAttributes) SetRestrictionType(v string)`

SetRestrictionType sets RestrictionType field to given value.

### HasRestrictionType

`func (o *POSTParcels201ResponseDataAttributes) HasRestrictionType() bool`

HasRestrictionType returns a boolean if a field has been set.

### GetRestrictionComments

`func (o *POSTParcels201ResponseDataAttributes) GetRestrictionComments() string`

GetRestrictionComments returns the RestrictionComments field if non-nil, zero value otherwise.

### GetRestrictionCommentsOk

`func (o *POSTParcels201ResponseDataAttributes) GetRestrictionCommentsOk() (*string, bool)`

GetRestrictionCommentsOk returns a tuple with the RestrictionComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionComments

`func (o *POSTParcels201ResponseDataAttributes) SetRestrictionComments(v string)`

SetRestrictionComments sets RestrictionComments field to given value.

### HasRestrictionComments

`func (o *POSTParcels201ResponseDataAttributes) HasRestrictionComments() bool`

HasRestrictionComments returns a boolean if a field has been set.

### GetCustomsInfoRequired

`func (o *POSTParcels201ResponseDataAttributes) GetCustomsInfoRequired() bool`

GetCustomsInfoRequired returns the CustomsInfoRequired field if non-nil, zero value otherwise.

### GetCustomsInfoRequiredOk

`func (o *POSTParcels201ResponseDataAttributes) GetCustomsInfoRequiredOk() (*bool, bool)`

GetCustomsInfoRequiredOk returns a tuple with the CustomsInfoRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsInfoRequired

`func (o *POSTParcels201ResponseDataAttributes) SetCustomsInfoRequired(v bool)`

SetCustomsInfoRequired sets CustomsInfoRequired field to given value.

### HasCustomsInfoRequired

`func (o *POSTParcels201ResponseDataAttributes) HasCustomsInfoRequired() bool`

HasCustomsInfoRequired returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *POSTParcels201ResponseDataAttributes) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *POSTParcels201ResponseDataAttributes) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *POSTParcels201ResponseDataAttributes) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *POSTParcels201ResponseDataAttributes) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetTrackingStatus

`func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatus() string`

GetTrackingStatus returns the TrackingStatus field if non-nil, zero value otherwise.

### GetTrackingStatusOk

`func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatusOk() (*string, bool)`

GetTrackingStatusOk returns a tuple with the TrackingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatus

`func (o *POSTParcels201ResponseDataAttributes) SetTrackingStatus(v string)`

SetTrackingStatus sets TrackingStatus field to given value.

### HasTrackingStatus

`func (o *POSTParcels201ResponseDataAttributes) HasTrackingStatus() bool`

HasTrackingStatus returns a boolean if a field has been set.

### GetTrackingStatusDetail

`func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatusDetail() string`

GetTrackingStatusDetail returns the TrackingStatusDetail field if non-nil, zero value otherwise.

### GetTrackingStatusDetailOk

`func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatusDetailOk() (*string, bool)`

GetTrackingStatusDetailOk returns a tuple with the TrackingStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusDetail

`func (o *POSTParcels201ResponseDataAttributes) SetTrackingStatusDetail(v string)`

SetTrackingStatusDetail sets TrackingStatusDetail field to given value.

### HasTrackingStatusDetail

`func (o *POSTParcels201ResponseDataAttributes) HasTrackingStatusDetail() bool`

HasTrackingStatusDetail returns a boolean if a field has been set.

### GetTrackingStatusUpdatedAt

`func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatusUpdatedAt() string`

GetTrackingStatusUpdatedAt returns the TrackingStatusUpdatedAt field if non-nil, zero value otherwise.

### GetTrackingStatusUpdatedAtOk

`func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatusUpdatedAtOk() (*string, bool)`

GetTrackingStatusUpdatedAtOk returns a tuple with the TrackingStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusUpdatedAt

`func (o *POSTParcels201ResponseDataAttributes) SetTrackingStatusUpdatedAt(v string)`

SetTrackingStatusUpdatedAt sets TrackingStatusUpdatedAt field to given value.

### HasTrackingStatusUpdatedAt

`func (o *POSTParcels201ResponseDataAttributes) HasTrackingStatusUpdatedAt() bool`

HasTrackingStatusUpdatedAt returns a boolean if a field has been set.

### GetTrackingDetails

`func (o *POSTParcels201ResponseDataAttributes) GetTrackingDetails() string`

GetTrackingDetails returns the TrackingDetails field if non-nil, zero value otherwise.

### GetTrackingDetailsOk

`func (o *POSTParcels201ResponseDataAttributes) GetTrackingDetailsOk() (*string, bool)`

GetTrackingDetailsOk returns a tuple with the TrackingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingDetails

`func (o *POSTParcels201ResponseDataAttributes) SetTrackingDetails(v string)`

SetTrackingDetails sets TrackingDetails field to given value.

### HasTrackingDetails

`func (o *POSTParcels201ResponseDataAttributes) HasTrackingDetails() bool`

HasTrackingDetails returns a boolean if a field has been set.

### GetCarrierWeightOz

`func (o *POSTParcels201ResponseDataAttributes) GetCarrierWeightOz() string`

GetCarrierWeightOz returns the CarrierWeightOz field if non-nil, zero value otherwise.

### GetCarrierWeightOzOk

`func (o *POSTParcels201ResponseDataAttributes) GetCarrierWeightOzOk() (*string, bool)`

GetCarrierWeightOzOk returns a tuple with the CarrierWeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierWeightOz

`func (o *POSTParcels201ResponseDataAttributes) SetCarrierWeightOz(v string)`

SetCarrierWeightOz sets CarrierWeightOz field to given value.

### HasCarrierWeightOz

`func (o *POSTParcels201ResponseDataAttributes) HasCarrierWeightOz() bool`

HasCarrierWeightOz returns a boolean if a field has been set.

### GetSignedBy

`func (o *POSTParcels201ResponseDataAttributes) GetSignedBy() string`

GetSignedBy returns the SignedBy field if non-nil, zero value otherwise.

### GetSignedByOk

`func (o *POSTParcels201ResponseDataAttributes) GetSignedByOk() (*string, bool)`

GetSignedByOk returns a tuple with the SignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedBy

`func (o *POSTParcels201ResponseDataAttributes) SetSignedBy(v string)`

SetSignedBy sets SignedBy field to given value.

### HasSignedBy

`func (o *POSTParcels201ResponseDataAttributes) HasSignedBy() bool`

HasSignedBy returns a boolean if a field has been set.

### GetIncoterm

`func (o *POSTParcels201ResponseDataAttributes) GetIncoterm() string`

GetIncoterm returns the Incoterm field if non-nil, zero value otherwise.

### GetIncotermOk

`func (o *POSTParcels201ResponseDataAttributes) GetIncotermOk() (*string, bool)`

GetIncotermOk returns a tuple with the Incoterm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncoterm

`func (o *POSTParcels201ResponseDataAttributes) SetIncoterm(v string)`

SetIncoterm sets Incoterm field to given value.

### HasIncoterm

`func (o *POSTParcels201ResponseDataAttributes) HasIncoterm() bool`

HasIncoterm returns a boolean if a field has been set.

### GetReference

`func (o *POSTParcels201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTParcels201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTParcels201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTParcels201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTParcels201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTParcels201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTParcels201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTParcels201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTParcels201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTParcels201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTParcels201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTParcels201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


