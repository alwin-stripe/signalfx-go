/*
 * Integrations API
 *
 * APIs for creating, retrieving, updating, and deleting SignalFx integrations to the systems you use.<br> An integration provides SignalFx with information from the external system that you're connecting to. You'll need to retrieve this information from the external system before you use the API. Each external system is different, so to see a summary of its requirements and procedures, view its request body description. # Authentication To create, update, delete, or validate an integration, you need to authenticate your request using a session token associated with a SignalFx administrator. To **retrieve** an integration, your session token doesn't need to be associated with an administrator. You can also retrieve integrations using an org token.<br> In the web UI, session tokens are known as <strong>user access</strong> tokens, and org tokens are known as <strong>access tokens</strong>. <br> To learn more about authentication tokens, see the topic [Authentication Tokens](https://developers.signalfx.com/administration/access_tokens_overview.html) in the Developers Guide. # Supported service types SignalFx offers integrations for the following:<br>   * Data collection from other monitoring systems such as AWS CloudWatch   * Authentication using your existing Single Sign-On (**SSO**) system   * Sending alerts using your preferred messaging, chat, or incident management service <br> To use one of these integrations, you first register it with SignalFx. After that, you configure the integration to communicate between the system you're using and SignalFx. ## Data collection SignalFx integrations APIs support data collection for the following services:<br>   * Amazon Web Services (**AWS**)   * Google Cloud Platform (**GCP**)   * Microsoft Azure   * NewRelic  ## Authentication using SSO SignalFx integration APIs support SAML-based SSO integrations for the following services:<br>   * Microsoft Active Directory Federation Services (**ADFS**)   * Bitium   * Okta   * OneLogin   * PingOne  ## Alerts using message, chat, or incident management services SignalFx integration APIs support alert notifications using the following services: <br>   * BigPanda   * Office 365   * Opsgenie   * PagerDuty   * ServiceNow   * Slack   * VictorOps   * Webhook   * xMatters<br>  **NOTE:** You can't create Office 365 integrations using the API, and your ability to update them in a **PUT** request is limited, but you can retrieve their data or delete them. To create an Office 365 integration, use the the web UI. <br> # Viewing request body documentation The *request* body format for the following operations depends on the type of integration you use:<br>   * POST `/integration`   * PUT `/integration/{id}`<br>  The *response* body format for the following operations also depends on the type of integration you use:<br>   * GET `/integration`   * GET `/integration/{id}`  <br>  To see the request or response body format for an integration: <br>   1. Find the endpoint and method.   2. For a request body, find the section *REQUEST BODY SCHEMA*. For a     response body, find the section *RESPONSE SCHEMA*.   3. Scroll down to the `type` property.   4. At the end of the description for `type`, find the dropdown box that      contains the integration type. By default, it's set to *AWSCloudWatch*.   5. To see a complete list of integrations, click the down arrow. A list      with a vertical scroll bar appears.   6. Select the integration type from the list. The request body properties      for this integration type now appear.
 *
 * API version: 3.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package integration

// Specifies the data collection integration between Google Cloud Platform and SignalFx, in the form of a JSON object.
type GcpIntegrationModel struct {
	// The creation date and time for the integration object, in Unix time UTC-relative. The system sets this value, and you can't modify it.
	Created int64 `json:"created,omitempty"`
	// SignalFx-assigned user ID of the user that created the integration object. If the system created the object, the value is \"AAAAAAAAAA\". The system sets this value, and you can't modify it.
	Creator string `json:"creator,omitempty"`
	// Flag that indicates the state of the integration object. If  `true`, the integration is enabled. If `false`, the integration is disabled, and you must enable it by setting \"enabled\" to `true` in a **PUT** request that updates the object. <br> **NOTE:** SignalFx always sets the flag to `true` when you call  **POST** `/integration` to create an integration.
	Enabled bool `json:"enabled,omitempty"`
	// SignalFx-assigned ID of an integration you create in the web UI or API. Use this property to retrieve an integration using the **GET**, **PUT**, or **DELETE** `/integration/{id}` endpoints or the **GET** `/integration/validate{id}/` endpoint, as described in this topic.
	Id string `json:"id,omitempty"`
	// The last time the integration was updated, in Unix time UTC-relative. This value is \"read-only\".
	LastUpdated int64 `json:"lastUpdated,omitempty"`
	// SignalFx-assigned ID of the last user who updated the integration. If the last update was by the system, the value is \"AAAAAAAAAA\". This value is \"read-only\".
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	// A human-readable label for the integration. This property helps you identify a specific integration when you're using multiple integrations for the same service.
	Name     string   `json:"name,omitempty"`
	Type     Type     `json:"type"`
	PollRate PollRate `json:"pollRate,omitempty"`
	// Array of GCP services that you want SignalFx to monitor. SignalFx only supports certain services, and if you specify an unsupported one, you receive an API error. The supported services are: <br>   * appengine   * bigquery   * bigtable   * cloudfunctions   * cloudiot   * cloudsql   * cloudtasks   * compute   * container   * dataflow   * datastore   * firebasedatabase   * firebasehosting   * interconnect   * loadbalancing   * logging   * ml   * monitoring   * pubsub   * router   * serviceruntime   * spanner   * storage   * vpn
	Services []string `json:"services,omitempty"`
	// List of GCP project that you want SignalFx to monitor, in the form of a JSON array of objects
	ProjectServiceKeys []*GcpProject `json:"projectServiceKeys,omitempty"`
	// List of GCP metadata names that you want SignalFx to collect from the data incoming from the GCP integration, in the form of a JSON array. Refer to Google's GCP documentation to find out the names you want to whitelist.
	Whitelist []string `json:"whitelist,omitempty"`
}
