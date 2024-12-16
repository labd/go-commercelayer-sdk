/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes{}

// PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes struct for PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes
type PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The gateway merchant account.
	MerchantAccount interface{} `json:"merchant_account,omitempty"`
	// The gateway API key.
	ApiKey interface{} `json:"api_key,omitempty"`
	// The public key linked to your API credential.
	PublicKey interface{} `json:"public_key,omitempty"`
	// The prefix of the endpoint used for live transactions.
	LiveUrlPrefix interface{} `json:"live_url_prefix,omitempty"`
	// The checkout API version, supported range is from 66 to 71, default is 71.
	ApiVersion interface{} `json:"api_version,omitempty"`
	// Indicates if the gateway will leverage on the Adyen notification webhooks, using latest API version.
	AsyncApi interface{} `json:"async_api,omitempty"`
	// Indicates if the gateway will use the native customer payment sources.
	NativeCustomerPaymentSources interface{} `json:"native_customer_payment_sources,omitempty"`
	// The gateway webhook endpoint secret, generated by Adyen customer area.
	WebhookEndpointSecret interface{} `json:"webhook_endpoint_secret,omitempty"`
}

// NewPATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes instantiates a new PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes() *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes {
	this := PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes{}
	return &this
}

// NewPATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributesWithDefaults() *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes {
	this := PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetMerchantAccount returns the MerchantAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetMerchantAccount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetMerchantAccountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantAccount) {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// HasMerchantAccount returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasMerchantAccount() bool {
	if o != nil && IsNil(o.MerchantAccount) {
		return true
	}

	return false
}

// SetMerchantAccount gets a reference to the given interface{} and assigns it to the MerchantAccount field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetMerchantAccount(v interface{}) {
	o.MerchantAccount = v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetApiKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetApiKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return &o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasApiKey() bool {
	if o != nil && IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given interface{} and assigns it to the ApiKey field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetApiKey(v interface{}) {
	o.ApiKey = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetPublicKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetPublicKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return &o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasPublicKey() bool {
	if o != nil && IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given interface{} and assigns it to the PublicKey field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetPublicKey(v interface{}) {
	o.PublicKey = v
}

// GetLiveUrlPrefix returns the LiveUrlPrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetLiveUrlPrefix() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LiveUrlPrefix
}

// GetLiveUrlPrefixOk returns a tuple with the LiveUrlPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetLiveUrlPrefixOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LiveUrlPrefix) {
		return nil, false
	}
	return &o.LiveUrlPrefix, true
}

// HasLiveUrlPrefix returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasLiveUrlPrefix() bool {
	if o != nil && IsNil(o.LiveUrlPrefix) {
		return true
	}

	return false
}

// SetLiveUrlPrefix gets a reference to the given interface{} and assigns it to the LiveUrlPrefix field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetLiveUrlPrefix(v interface{}) {
	o.LiveUrlPrefix = v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetApiVersion() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetApiVersionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return &o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasApiVersion() bool {
	if o != nil && IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given interface{} and assigns it to the ApiVersion field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetApiVersion(v interface{}) {
	o.ApiVersion = v
}

// GetAsyncApi returns the AsyncApi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetAsyncApi() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AsyncApi
}

// GetAsyncApiOk returns a tuple with the AsyncApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetAsyncApiOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AsyncApi) {
		return nil, false
	}
	return &o.AsyncApi, true
}

// HasAsyncApi returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasAsyncApi() bool {
	if o != nil && IsNil(o.AsyncApi) {
		return true
	}

	return false
}

// SetAsyncApi gets a reference to the given interface{} and assigns it to the AsyncApi field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetAsyncApi(v interface{}) {
	o.AsyncApi = v
}

// GetNativeCustomerPaymentSources returns the NativeCustomerPaymentSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetNativeCustomerPaymentSources() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NativeCustomerPaymentSources
}

// GetNativeCustomerPaymentSourcesOk returns a tuple with the NativeCustomerPaymentSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetNativeCustomerPaymentSourcesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NativeCustomerPaymentSources) {
		return nil, false
	}
	return &o.NativeCustomerPaymentSources, true
}

// HasNativeCustomerPaymentSources returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasNativeCustomerPaymentSources() bool {
	if o != nil && IsNil(o.NativeCustomerPaymentSources) {
		return true
	}

	return false
}

// SetNativeCustomerPaymentSources gets a reference to the given interface{} and assigns it to the NativeCustomerPaymentSources field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetNativeCustomerPaymentSources(v interface{}) {
	o.NativeCustomerPaymentSources = v
}

// GetWebhookEndpointSecret returns the WebhookEndpointSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetWebhookEndpointSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WebhookEndpointSecret
}

// GetWebhookEndpointSecretOk returns a tuple with the WebhookEndpointSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetWebhookEndpointSecretOk() (*interface{}, bool) {
	if o == nil || IsNil(o.WebhookEndpointSecret) {
		return nil, false
	}
	return &o.WebhookEndpointSecret, true
}

// HasWebhookEndpointSecret returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasWebhookEndpointSecret() bool {
	if o != nil && IsNil(o.WebhookEndpointSecret) {
		return true
	}

	return false
}

// SetWebhookEndpointSecret gets a reference to the given interface{} and assigns it to the WebhookEndpointSecret field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetWebhookEndpointSecret(v interface{}) {
	o.WebhookEndpointSecret = v
}

func (o PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MerchantAccount != nil {
		toSerialize["merchant_account"] = o.MerchantAccount
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	if o.LiveUrlPrefix != nil {
		toSerialize["live_url_prefix"] = o.LiveUrlPrefix
	}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.AsyncApi != nil {
		toSerialize["async_api"] = o.AsyncApi
	}
	if o.NativeCustomerPaymentSources != nil {
		toSerialize["native_customer_payment_sources"] = o.NativeCustomerPaymentSources
	}
	if o.WebhookEndpointSecret != nil {
		toSerialize["webhook_endpoint_secret"] = o.WebhookEndpointSecret
	}
	return toSerialize, nil
}

type NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes struct {
	value *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) Get() *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) Set(val *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes(val *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) *NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes {
	return &NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
