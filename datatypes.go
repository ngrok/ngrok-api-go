// Code generated for API Clients. DO NOT EDIT.

package ngrok

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

type Empty struct {
}

func (x *Empty) String() string {
	return x.GoString()
}

func (x *Empty) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Empty {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type Item struct {
	// a resource identifier
	ID string `json:"id,omitempty"`
}

func (x *Item) String() string {
	return fmt.Sprintf("Item{ID: %v}", x.ID)

}

func (x *Item) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Item {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type Paging struct {
	BeforeID *string `json:"before_id,omitempty"`
	Limit    *string `json:"limit,omitempty"`
}

func (x *Paging) String() string {
	return x.GoString()
}

func (x *Paging) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Paging {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tBeforeID\t%v\n", x.BeforeID)
	fmt.Fprintf(tw, "\tLimit\t%v\n", x.Limit)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type Error struct {
	ErrorCode  string            `json:"error_code,omitempty"`
	StatusCode int32             `json:"status_code,omitempty"`
	Msg        string            `json:"msg,omitempty"`
	Details    map[string]string `json:"details,omitempty"`
}

func (x *Error) String() string {
	return x.GoString()
}

func (x *Error) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Error {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tErrorCode\t%v\n", x.ErrorCode)
	fmt.Fprintf(tw, "\tStatusCode\t%v\n", x.StatusCode)
	fmt.Fprintf(tw, "\tMsg\t%v\n", x.Msg)
	fmt.Fprintf(tw, "\tDetails\t%v\n", x.Details)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type Ref struct {
	// a resource identifier
	ID string `json:"id,omitempty"`
	// a uri for locating a resource
	URI string `json:"uri,omitempty"`
}

func (x *Ref) String() string {
	return fmt.Sprintf("Ref{ID: %v}", x.ID)

}

func (x *Ref) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Ref {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AbuseReport struct {
	// ID of the abuse report
	ID string `json:"id,omitempty"`
	// URI of the abuse report API resource
	URI string `json:"uri,omitempty"`
	// timestamp that the abuse report record was created in RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// a list of URLs containing suspected abusive content
	URLs []string `json:"urls,omitempty"`
	// arbitrary user-defined data about this abuse report. Optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// Indicates whether ngrok has processed the abuse report. one of PENDING,
	// PROCESSED, or PARTIALLY_PROCESSED
	Status string `json:"status,omitempty"`
	// an array of hostname statuses related to the report
	Hostnames []AbuseReportHostname `json:"hostnames,omitempty"`
}

func (x *AbuseReport) String() string {
	return fmt.Sprintf("AbuseReport{ID: %v}", x.ID)

}

func (x *AbuseReport) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AbuseReport {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tURLs\t%v\n", x.URLs)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tStatus\t%v\n", x.Status)
	fmt.Fprintf(tw, "\tHostnames\t%v\n", x.Hostnames)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AbuseReportHostname struct {
	// the hostname ngrok has parsed out of one of the reported URLs in this abuse
	// report
	Hostname string `json:"hostname,omitempty"`
	// indicates what action ngrok has taken against the hostname. one of PENDING,
	// BANNED, UNBANNED, or IGNORE
	Status string `json:"status,omitempty"`
}

func (x *AbuseReportHostname) String() string {
	return x.GoString()
}

func (x *AbuseReportHostname) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AbuseReportHostname {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tHostname\t%v\n", x.Hostname)
	fmt.Fprintf(tw, "\tStatus\t%v\n", x.Status)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AbuseReportCreate struct {
	// a list of URLs containing suspected abusive content
	URLs []string `json:"urls,omitempty"`
	// arbitrary user-defined data about this abuse report. Optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
}

func (x *AbuseReportCreate) String() string {
	return x.GoString()
}

func (x *AbuseReportCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AbuseReportCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tURLs\t%v\n", x.URLs)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AgentIngressCreate struct {
	// human-readable description of the use of this Agent Ingress. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Agent Ingress. optional,
	// max 4096 bytes
	Metadata string `json:"metadata,omitempty"`
	// the domain that you own to be used as the base domain name to generate regional
	// agent ingress domains.
	Domain string `json:"domain,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled. Optional.
	CertificateManagementPolicy *AgentIngressCertPolicy `json:"certificate_management_policy,omitempty"`
}

func (x *AgentIngressCreate) String() string {
	return x.GoString()
}

func (x *AgentIngressCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AgentIngressCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tDomain\t%v\n", x.Domain)
	fmt.Fprintf(tw, "\tCertificateManagementPolicy\t%v\n", x.CertificateManagementPolicy)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AgentIngressUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of the use of this Agent Ingress. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Agent Ingress. optional,
	// max 4096 bytes
	Metadata *string `json:"metadata,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled. Optional.
	CertificateManagementPolicy *AgentIngressCertPolicy `json:"certificate_management_policy,omitempty"`
}

func (x *AgentIngressUpdate) String() string {
	return fmt.Sprintf("AgentIngressUpdate{ID: %v}", x.ID)

}

func (x *AgentIngressUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AgentIngressUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCertificateManagementPolicy\t%v\n", x.CertificateManagementPolicy)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AgentIngress struct {
	// unique Agent Ingress resource identifier
	ID string `json:"id,omitempty"`
	// URI to the API resource of this Agent ingress
	URI string `json:"uri,omitempty"`
	// human-readable description of the use of this Agent Ingress. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Agent Ingress. optional,
	// max 4096 bytes
	Metadata string `json:"metadata,omitempty"`
	// the domain that you own to be used as the base domain name to generate regional
	// agent ingress domains.
	Domain string `json:"domain,omitempty"`
	// a list of target values to use as the values of NS records for the domain
	// property these values will delegate control over the domain to ngrok
	NSTargets []string `json:"ns_targets,omitempty"`
	// a list of regional agent ingress domains that are subdomains of the value of
	// domain this value may increase over time as ngrok adds more regions
	RegionDomains []string `json:"region_domains,omitempty"`
	// timestamp when the Agent Ingress was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled
	CertificateManagementPolicy *AgentIngressCertPolicy `json:"certificate_management_policy,omitempty"`
	// status of the automatic certificate management for this domain, or null if
	// automatic management is disabled
	CertificateManagementStatus *AgentIngressCertStatus `json:"certificate_management_status,omitempty"`
}

func (x *AgentIngress) String() string {
	return fmt.Sprintf("AgentIngress{ID: %v}", x.ID)

}

func (x *AgentIngress) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AgentIngress {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tDomain\t%v\n", x.Domain)
	fmt.Fprintf(tw, "\tNSTargets\t%v\n", x.NSTargets)
	fmt.Fprintf(tw, "\tRegionDomains\t%v\n", x.RegionDomains)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tCertificateManagementPolicy\t%v\n", x.CertificateManagementPolicy)
	fmt.Fprintf(tw, "\tCertificateManagementStatus\t%v\n", x.CertificateManagementStatus)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AgentIngressList struct {
	// the list of Agent Ingresses owned by this account
	Ingresses []AgentIngress `json:"ingresses,omitempty"`
	// URI of the Agent Ingress list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *AgentIngressList) String() string {
	return x.GoString()
}

func (x *AgentIngressList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AgentIngressList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tIngresses\t%v\n", x.Ingresses)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AgentIngressCertPolicy struct {
	// certificate authority to request certificates from. The only supported value is
	// letsencrypt.
	Authority string `json:"authority,omitempty"`
	// type of private key to use when requesting certificates. Defaults to rsa, can be
	// either rsa or ecdsa.
	PrivateKeyType string `json:"private_key_type,omitempty"`
}

func (x *AgentIngressCertPolicy) String() string {
	return x.GoString()
}

func (x *AgentIngressCertPolicy) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AgentIngressCertPolicy {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tAuthority\t%v\n", x.Authority)
	fmt.Fprintf(tw, "\tPrivateKeyType\t%v\n", x.PrivateKeyType)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AgentIngressCertStatus struct {
	// timestamp when the next renewal will be requested, RFC 3339 format
	RenewsAt *string `json:"renews_at,omitempty"`
	// status of the certificate provisioning job, or null if the certificiate isn't
	// being provisioned or renewed
	ProvisioningJob *AgentIngressCertJob `json:"provisioning_job,omitempty"`
}

func (x *AgentIngressCertStatus) String() string {
	return x.GoString()
}

func (x *AgentIngressCertStatus) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AgentIngressCertStatus {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tRenewsAt\t%v\n", x.RenewsAt)
	fmt.Fprintf(tw, "\tProvisioningJob\t%v\n", x.ProvisioningJob)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AgentIngressCertJob struct {
	// if present, an error code indicating why provisioning is failing. It may be
	// either a temporary condition (INTERNAL_ERROR), or a permanent one the user must
	// correct (DNS_ERROR).
	ErrorCode *string `json:"error_code,omitempty"`
	// a message describing the current status or error
	Msg string `json:"msg,omitempty"`
	// timestamp when the provisioning job started, RFC 3339 format
	StartedAt string `json:"started_at,omitempty"`
	// timestamp when the provisioning job will be retried
	RetriesAt *string `json:"retries_at,omitempty"`
}

func (x *AgentIngressCertJob) String() string {
	return x.GoString()
}

func (x *AgentIngressCertJob) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AgentIngressCertJob {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tErrorCode\t%v\n", x.ErrorCode)
	fmt.Fprintf(tw, "\tMsg\t%v\n", x.Msg)
	fmt.Fprintf(tw, "\tStartedAt\t%v\n", x.StartedAt)
	fmt.Fprintf(tw, "\tRetriesAt\t%v\n", x.RetriesAt)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type APIKeyCreate struct {
	// human-readable description of what uses the API key to authenticate. optional,
	// max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined data of this API key. optional, max 4096 bytes
	Metadata string `json:"metadata,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
}

func (x *APIKeyCreate) String() string {
	return x.GoString()
}

func (x *APIKeyCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "APIKeyCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tOwnerID\t%v\n", x.OwnerID)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type APIKeyUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of what uses the API key to authenticate. optional,
	// max 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined data of this API key. optional, max 4096 bytes
	Metadata *string `json:"metadata,omitempty"`
}

func (x *APIKeyUpdate) String() string {
	return fmt.Sprintf("APIKeyUpdate{ID: %v}", x.ID)

}

func (x *APIKeyUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "APIKeyUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type APIKey struct {
	// unique API key resource identifier
	ID string `json:"id,omitempty"`
	// URI to the API resource of this API key
	URI string `json:"uri,omitempty"`
	// human-readable description of what uses the API key to authenticate. optional,
	// max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined data of this API key. optional, max 4096 bytes
	Metadata string `json:"metadata,omitempty"`
	// timestamp when the api key was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// the bearer token that can be placed into the Authorization header to
	// authenticate request to the ngrok API. This value is only available one time, on
	// the API response from key creation. Otherwise it is null.
	Token *string `json:"token,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
}

func (x *APIKey) String() string {
	return fmt.Sprintf("APIKey{ID: %v}", x.ID)

}

func (x *APIKey) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "APIKey {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tToken\t%v\n", x.Token)
	fmt.Fprintf(tw, "\tOwnerID\t%v\n", x.OwnerID)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type APIKeyList struct {
	// the list of API keys for this account
	Keys []APIKey `json:"keys,omitempty"`
	// URI of the API keys list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *APIKeyList) String() string {
	return x.GoString()
}

func (x *APIKeyList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "APIKeyList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tKeys\t%v\n", x.Keys)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ApplicationSession struct {
	// unique application session resource identifier
	ID string `json:"id,omitempty"`
	// URI of the application session API resource
	URI string `json:"uri,omitempty"`
	// URL of the hostport served by this endpoint
	PublicURL string `json:"public_url,omitempty"`
	// browser session details of the application session
	BrowserSession BrowserSession `json:"browser_session,omitempty"`
	// application user this session is associated with
	ApplicationUser *Ref `json:"application_user,omitempty"`
	// timestamp when the user was created in RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// timestamp when the user was last active in RFC 3339 format
	LastActive string `json:"last_active,omitempty"`
	// timestamp when session expires in RFC 3339 format
	ExpiresAt string `json:"expires_at,omitempty"`
	// ephemeral endpoint this session is associated with
	Endpoint *Ref `json:"endpoint,omitempty"`
	// edge this session is associated with, null if the endpoint is agent-initiated
	Edge *Ref `json:"edge,omitempty"`
	// route this session is associated with, null if the endpoint is agent-initiated
	Route *Ref `json:"route,omitempty"`
}

func (x *ApplicationSession) String() string {
	return fmt.Sprintf("ApplicationSession{ID: %v}", x.ID)

}

func (x *ApplicationSession) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ApplicationSession {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tPublicURL\t%v\n", x.PublicURL)
	fmt.Fprintf(tw, "\tBrowserSession\t%v\n", x.BrowserSession)
	fmt.Fprintf(tw, "\tApplicationUser\t%v\n", x.ApplicationUser)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tLastActive\t%v\n", x.LastActive)
	fmt.Fprintf(tw, "\tExpiresAt\t%v\n", x.ExpiresAt)
	fmt.Fprintf(tw, "\tEndpoint\t%v\n", x.Endpoint)
	fmt.Fprintf(tw, "\tEdge\t%v\n", x.Edge)
	fmt.Fprintf(tw, "\tRoute\t%v\n", x.Route)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ApplicationSessionList struct {
	// list of all application sessions on this account
	ApplicationSessions []ApplicationSession `json:"application_sessions,omitempty"`
	// URI of the application session list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *ApplicationSessionList) String() string {
	return x.GoString()
}

func (x *ApplicationSessionList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ApplicationSessionList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tApplicationSessions\t%v\n", x.ApplicationSessions)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type BrowserSession struct {
	// HTTP User-Agent data
	UserAgent UserAgent `json:"user_agent,omitempty"`
	// IP address
	IPAddress string `json:"ip_address,omitempty"`
	// IP geolocation data
	Location *Location `json:"location,omitempty"`
}

func (x *BrowserSession) String() string {
	return x.GoString()
}

func (x *BrowserSession) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "BrowserSession {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tUserAgent\t%v\n", x.UserAgent)
	fmt.Fprintf(tw, "\tIPAddress\t%v\n", x.IPAddress)
	fmt.Fprintf(tw, "\tLocation\t%v\n", x.Location)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type UserAgent struct {
	// raw User-Agent request header
	Raw string `json:"raw,omitempty"`
	// browser name (e.g. Chrome)
	BrowserName string `json:"browser_name,omitempty"`
	// browser version (e.g. 102)
	BrowserVersion string `json:"browser_version,omitempty"`
	// type of device (e.g. Desktop)
	DeviceType string `json:"device_type,omitempty"`
	// operating system name (e.g. MacOS)
	OSName string `json:"os_name,omitempty"`
	// operating system version (e.g. 10.15.7)
	OSVersion string `json:"os_version,omitempty"`
}

func (x *UserAgent) String() string {
	return x.GoString()
}

func (x *UserAgent) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "UserAgent {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tRaw\t%v\n", x.Raw)
	fmt.Fprintf(tw, "\tBrowserName\t%v\n", x.BrowserName)
	fmt.Fprintf(tw, "\tBrowserVersion\t%v\n", x.BrowserVersion)
	fmt.Fprintf(tw, "\tDeviceType\t%v\n", x.DeviceType)
	fmt.Fprintf(tw, "\tOSName\t%v\n", x.OSName)
	fmt.Fprintf(tw, "\tOSVersion\t%v\n", x.OSVersion)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type Location struct {
	// ISO country code
	CountryCode *string `json:"country_code,omitempty"`
	// geographical latitude
	Latitude *float64 `json:"latitude,omitempty"`
	// geographical longitude
	Longitude *float64 `json:"longitude,omitempty"`
	// accuracy radius of the geographical coordinates
	LatLongRadiusKm *uint64 `json:"lat_long_radius_km,omitempty"`
}

func (x *Location) String() string {
	return x.GoString()
}

func (x *Location) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Location {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tCountryCode\t%v\n", x.CountryCode)
	fmt.Fprintf(tw, "\tLatitude\t%v\n", x.Latitude)
	fmt.Fprintf(tw, "\tLongitude\t%v\n", x.Longitude)
	fmt.Fprintf(tw, "\tLatLongRadiusKm\t%v\n", x.LatLongRadiusKm)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ApplicationUser struct {
	// unique application user resource identifier
	ID string `json:"id,omitempty"`
	// URI of the application user API resource
	URI string `json:"uri,omitempty"`
	// identity provider that the user authenticated with
	IdentityProvider IdentityProvider `json:"identity_provider,omitempty"`
	// unique user identifier
	ProviderUserID string `json:"provider_user_id,omitempty"`
	// user username
	Username string `json:"username,omitempty"`
	// user email
	Email string `json:"email,omitempty"`
	// user common name
	Name string `json:"name,omitempty"`
	// timestamp when the user was created in RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// timestamp when the user was last active in RFC 3339 format
	LastActive string `json:"last_active,omitempty"`
	// timestamp when the user last signed-in in RFC 3339 format
	LastLogin string `json:"last_login,omitempty"`
}

func (x *ApplicationUser) String() string {
	return fmt.Sprintf("ApplicationUser{ID: %v}", x.ID)

}

func (x *ApplicationUser) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ApplicationUser {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tIdentityProvider\t%v\n", x.IdentityProvider)
	fmt.Fprintf(tw, "\tProviderUserID\t%v\n", x.ProviderUserID)
	fmt.Fprintf(tw, "\tUsername\t%v\n", x.Username)
	fmt.Fprintf(tw, "\tEmail\t%v\n", x.Email)
	fmt.Fprintf(tw, "\tName\t%v\n", x.Name)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tLastActive\t%v\n", x.LastActive)
	fmt.Fprintf(tw, "\tLastLogin\t%v\n", x.LastLogin)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ApplicationUserList struct {
	// list of all application users on this account
	ApplicationUsers []ApplicationUser `json:"application_users,omitempty"`
	// URI of the application user list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *ApplicationUserList) String() string {
	return x.GoString()
}

func (x *ApplicationUserList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ApplicationUserList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tApplicationUsers\t%v\n", x.ApplicationUsers)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IdentityProvider struct {
	// name of the identity provider (e.g. Google)
	Name string `json:"name,omitempty"`
	// URL of the identity provider (e.g. https://accounts.google.com
	// (https://accounts.google.com))
	URL string `json:"url,omitempty"`
}

func (x *IdentityProvider) String() string {
	return x.GoString()
}

func (x *IdentityProvider) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IdentityProvider {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tName\t%v\n", x.Name)
	fmt.Fprintf(tw, "\tURL\t%v\n", x.URL)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TunnelSession struct {
	// version of the ngrok agent that started this ngrok tunnel session
	AgentVersion string `json:"agent_version,omitempty"`
	// reference to the tunnel credential or ssh credential used by the ngrok agent to
	// start this tunnel session
	Credential Ref `json:"credential,omitempty"`
	// unique tunnel session resource identifier
	ID string `json:"id,omitempty"`
	// source ip address of the tunnel session
	IP string `json:"ip,omitempty"`
	// arbitrary user-defined data specified in the metadata property in the ngrok
	// configuration file. See the metadata configuration option
	Metadata string `json:"metadata,omitempty"`
	// operating system of the host the ngrok agent is running on
	OS string `json:"os,omitempty"`
	// the ngrok region identifier in which this tunnel session was started
	Region string `json:"region,omitempty"`
	// time when the tunnel session first connected to the ngrok servers
	StartedAt string `json:"started_at,omitempty"`
	// the transport protocol used to start the tunnel session. Either ngrok/v2 or ssh
	Transport string `json:"transport,omitempty"`
	// URI to the API resource of the tunnel session
	URI string `json:"uri,omitempty"`
}

func (x *TunnelSession) String() string {
	return fmt.Sprintf("TunnelSession{ID: %v}", x.ID)

}

func (x *TunnelSession) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TunnelSession {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tAgentVersion\t%v\n", x.AgentVersion)
	fmt.Fprintf(tw, "\tCredential\t%v\n", x.Credential)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tIP\t%v\n", x.IP)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tOS\t%v\n", x.OS)
	fmt.Fprintf(tw, "\tRegion\t%v\n", x.Region)
	fmt.Fprintf(tw, "\tStartedAt\t%v\n", x.StartedAt)
	fmt.Fprintf(tw, "\tTransport\t%v\n", x.Transport)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TunnelSessionList struct {
	// list of all tunnel sessions on this account
	TunnelSessions []TunnelSession `json:"tunnel_sessions,omitempty"`
	// URI to the API resource of the tunnel session list
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *TunnelSessionList) String() string {
	return x.GoString()
}

func (x *TunnelSessionList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TunnelSessionList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tTunnelSessions\t%v\n", x.TunnelSessions)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TunnelSessionsUpdate struct {
	ID string `json:"id,omitempty"`
}

func (x *TunnelSessionsUpdate) String() string {
	return fmt.Sprintf("TunnelSessionsUpdate{ID: %v}", x.ID)

}

func (x *TunnelSessionsUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TunnelSessionsUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type FailoverBackend struct {
	// unique identifier for this Failover backend
	ID string `json:"id,omitempty"`
	// URI of the FailoverBackend API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the backend was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// the ids of the child backends in order
	Backends []string `json:"backends,omitempty"`
}

func (x *FailoverBackend) String() string {
	return fmt.Sprintf("FailoverBackend{ID: %v}", x.ID)

}

func (x *FailoverBackend) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "FailoverBackend {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBackends\t%v\n", x.Backends)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type FailoverBackendCreate struct {
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// the ids of the child backends in order
	Backends []string `json:"backends,omitempty"`
}

func (x *FailoverBackendCreate) String() string {
	return x.GoString()
}

func (x *FailoverBackendCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "FailoverBackendCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBackends\t%v\n", x.Backends)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type FailoverBackendUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this backend. Optional
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata *string `json:"metadata,omitempty"`
	// the ids of the child backends in order
	Backends []string `json:"backends,omitempty"`
}

func (x *FailoverBackendUpdate) String() string {
	return fmt.Sprintf("FailoverBackendUpdate{ID: %v}", x.ID)

}

func (x *FailoverBackendUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "FailoverBackendUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBackends\t%v\n", x.Backends)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type FailoverBackendList struct {
	// the list of all Failover backends on this account
	Backends []FailoverBackend `json:"backends,omitempty"`
	// URI of the Failover backends list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *FailoverBackendList) String() string {
	return x.GoString()
}

func (x *FailoverBackendList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "FailoverBackendList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tBackends\t%v\n", x.Backends)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type HTTPResponseBackend struct {
	ID string `json:"id,omitempty"`
	// URI of the HTTPResponseBackend API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the backend was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// body to return as fixed content
	Body string `json:"body,omitempty"`
	// headers to return
	Headers map[string]string `json:"headers,omitempty"`
	// status code to return
	StatusCode int32 `json:"status_code,omitempty"`
}

func (x *HTTPResponseBackend) String() string {
	return fmt.Sprintf("HTTPResponseBackend{ID: %v}", x.ID)

}

func (x *HTTPResponseBackend) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "HTTPResponseBackend {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBody\t%v\n", x.Body)
	fmt.Fprintf(tw, "\tHeaders\t%v\n", x.Headers)
	fmt.Fprintf(tw, "\tStatusCode\t%v\n", x.StatusCode)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type HTTPResponseBackendCreate struct {
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// body to return as fixed content
	Body string `json:"body,omitempty"`
	// headers to return
	Headers map[string]string `json:"headers,omitempty"`
	// status code to return
	StatusCode *int32 `json:"status_code,omitempty"`
}

func (x *HTTPResponseBackendCreate) String() string {
	return x.GoString()
}

func (x *HTTPResponseBackendCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "HTTPResponseBackendCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBody\t%v\n", x.Body)
	fmt.Fprintf(tw, "\tHeaders\t%v\n", x.Headers)
	fmt.Fprintf(tw, "\tStatusCode\t%v\n", x.StatusCode)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type HTTPResponseBackendUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this backend. Optional
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata *string `json:"metadata,omitempty"`
	// body to return as fixed content
	Body *string `json:"body,omitempty"`
	// headers to return
	Headers map[string]string `json:"headers,omitempty"`
	// status code to return
	StatusCode *int32 `json:"status_code,omitempty"`
}

func (x *HTTPResponseBackendUpdate) String() string {
	return fmt.Sprintf("HTTPResponseBackendUpdate{ID: %v}", x.ID)

}

func (x *HTTPResponseBackendUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "HTTPResponseBackendUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBody\t%v\n", x.Body)
	fmt.Fprintf(tw, "\tHeaders\t%v\n", x.Headers)
	fmt.Fprintf(tw, "\tStatusCode\t%v\n", x.StatusCode)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type HTTPResponseBackendList struct {
	Backends    []HTTPResponseBackend `json:"backends,omitempty"`
	URI         string                `json:"uri,omitempty"`
	NextPageURI *string               `json:"next_page_uri,omitempty"`
}

func (x *HTTPResponseBackendList) String() string {
	return x.GoString()
}

func (x *HTTPResponseBackendList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "HTTPResponseBackendList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tBackends\t%v\n", x.Backends)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TunnelGroupBackend struct {
	// unique identifier for this TunnelGroup backend
	ID string `json:"id,omitempty"`
	// URI of the TunnelGroupBackend API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the backend was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// labels to watch for tunnels on, e.g. app->foo, dc->bar
	Labels map[string]string `json:"labels,omitempty"`
	// tunnels matching this backend
	Tunnels []Ref `json:"tunnels,omitempty"`
}

func (x *TunnelGroupBackend) String() string {
	return fmt.Sprintf("TunnelGroupBackend{ID: %v}", x.ID)

}

func (x *TunnelGroupBackend) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TunnelGroupBackend {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tLabels\t%v\n", x.Labels)
	fmt.Fprintf(tw, "\tTunnels\t%v\n", x.Tunnels)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TunnelGroupBackendCreate struct {
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// labels to watch for tunnels on, e.g. app->foo, dc->bar
	Labels map[string]string `json:"labels,omitempty"`
}

func (x *TunnelGroupBackendCreate) String() string {
	return x.GoString()
}

func (x *TunnelGroupBackendCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TunnelGroupBackendCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tLabels\t%v\n", x.Labels)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TunnelGroupBackendUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this backend. Optional
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata *string `json:"metadata,omitempty"`
	// labels to watch for tunnels on, e.g. app->foo, dc->bar
	Labels map[string]string `json:"labels,omitempty"`
}

func (x *TunnelGroupBackendUpdate) String() string {
	return fmt.Sprintf("TunnelGroupBackendUpdate{ID: %v}", x.ID)

}

func (x *TunnelGroupBackendUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TunnelGroupBackendUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tLabels\t%v\n", x.Labels)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TunnelGroupBackendList struct {
	// the list of all TunnelGroup backends on this account
	Backends []TunnelGroupBackend `json:"backends,omitempty"`
	// URI of the TunnelGroup backends list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *TunnelGroupBackendList) String() string {
	return x.GoString()
}

func (x *TunnelGroupBackendList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TunnelGroupBackendList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tBackends\t%v\n", x.Backends)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type WeightedBackend struct {
	// unique identifier for this Weighted backend
	ID string `json:"id,omitempty"`
	// URI of the WeightedBackend API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the backend was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// the ids of the child backends to their weights [0-10000]
	Backends map[string]int64 `json:"backends,omitempty"`
}

func (x *WeightedBackend) String() string {
	return fmt.Sprintf("WeightedBackend{ID: %v}", x.ID)

}

func (x *WeightedBackend) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "WeightedBackend {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBackends\t%v\n", x.Backends)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type WeightedBackendCreate struct {
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// the ids of the child backends to their weights [0-10000]
	Backends map[string]int64 `json:"backends,omitempty"`
}

func (x *WeightedBackendCreate) String() string {
	return x.GoString()
}

func (x *WeightedBackendCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "WeightedBackendCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBackends\t%v\n", x.Backends)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type WeightedBackendUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this backend. Optional
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata *string `json:"metadata,omitempty"`
	// the ids of the child backends to their weights [0-10000]
	Backends map[string]int64 `json:"backends,omitempty"`
}

func (x *WeightedBackendUpdate) String() string {
	return fmt.Sprintf("WeightedBackendUpdate{ID: %v}", x.ID)

}

func (x *WeightedBackendUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "WeightedBackendUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBackends\t%v\n", x.Backends)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type WeightedBackendList struct {
	// the list of all Weighted backends on this account
	Backends []WeightedBackend `json:"backends,omitempty"`
	// URI of the Weighted backends list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *WeightedBackendList) String() string {
	return x.GoString()
}

func (x *WeightedBackendList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "WeightedBackendList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tBackends\t%v\n", x.Backends)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type CertificateAuthorityCreate struct {
	// human-readable description of this Certificate Authority. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Certificate Authority.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// raw PEM of the Certificate Authority
	CAPEM string `json:"ca_pem,omitempty"`
}

func (x *CertificateAuthorityCreate) String() string {
	return x.GoString()
}

func (x *CertificateAuthorityCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "CertificateAuthorityCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCAPEM\t%v\n", x.CAPEM)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type CertificateAuthorityUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this Certificate Authority. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Certificate Authority.
	// optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

func (x *CertificateAuthorityUpdate) String() string {
	return fmt.Sprintf("CertificateAuthorityUpdate{ID: %v}", x.ID)

}

func (x *CertificateAuthorityUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "CertificateAuthorityUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type CertificateAuthority struct {
	// unique identifier for this Certificate Authority
	ID string `json:"id,omitempty"`
	// URI of the Certificate Authority API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the Certificate Authority was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this Certificate Authority. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Certificate Authority.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// raw PEM of the Certificate Authority
	CAPEM string `json:"ca_pem,omitempty"`
	// subject common name of the Certificate Authority
	SubjectCommonName string `json:"subject_common_name,omitempty"`
	// timestamp when this Certificate Authority becomes valid, RFC 3339 format
	NotBefore string `json:"not_before,omitempty"`
	// timestamp when this Certificate Authority becomes invalid, RFC 3339 format
	NotAfter string `json:"not_after,omitempty"`
	// set of actions the private key of this Certificate Authority can be used for
	KeyUsages []string `json:"key_usages,omitempty"`
	// extended set of actions the private key of this Certificate Authority can be
	// used for
	ExtendedKeyUsages []string `json:"extended_key_usages,omitempty"`
}

func (x *CertificateAuthority) String() string {
	return fmt.Sprintf("CertificateAuthority{ID: %v}", x.ID)

}

func (x *CertificateAuthority) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "CertificateAuthority {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCAPEM\t%v\n", x.CAPEM)
	fmt.Fprintf(tw, "\tSubjectCommonName\t%v\n", x.SubjectCommonName)
	fmt.Fprintf(tw, "\tNotBefore\t%v\n", x.NotBefore)
	fmt.Fprintf(tw, "\tNotAfter\t%v\n", x.NotAfter)
	fmt.Fprintf(tw, "\tKeyUsages\t%v\n", x.KeyUsages)
	fmt.Fprintf(tw, "\tExtendedKeyUsages\t%v\n", x.ExtendedKeyUsages)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type CertificateAuthorityList struct {
	// the list of all certificate authorities on this account
	CertificateAuthorities []CertificateAuthority `json:"certificate_authorities,omitempty"`
	// URI of the certificates authorities list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *CertificateAuthorityList) String() string {
	return x.GoString()
}

func (x *CertificateAuthorityList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "CertificateAuthorityList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tCertificateAuthorities\t%v\n", x.CertificateAuthorities)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type CredentialCreate struct {
	// human-readable description of who or what will use the credential to
	// authenticate. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this credential. Optional, max
	// 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL []string `json:"acl,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
}

func (x *CredentialCreate) String() string {
	return x.GoString()
}

func (x *CredentialCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "CredentialCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tACL\t%v\n", x.ACL)
	fmt.Fprintf(tw, "\tOwnerID\t%v\n", x.OwnerID)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type CredentialUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of who or what will use the credential to
	// authenticate. Optional, max 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this credential. Optional, max
	// 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL []string `json:"acl,omitempty"`
}

func (x *CredentialUpdate) String() string {
	return fmt.Sprintf("CredentialUpdate{ID: %v}", x.ID)

}

func (x *CredentialUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "CredentialUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tACL\t%v\n", x.ACL)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type Credential struct {
	// unique tunnel credential resource identifier
	ID string `json:"id,omitempty"`
	// URI of the tunnel credential API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the tunnel credential was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of who or what will use the credential to
	// authenticate. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this credential. Optional, max
	// 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// the credential's authtoken that can be used to authenticate an ngrok agent. This
	// value is only available one time, on the API response from credential creation,
	// otherwise it is null.
	Token *string `json:"token,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL []string `json:"acl,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
}

func (x *Credential) String() string {
	return fmt.Sprintf("Credential{ID: %v}", x.ID)

}

func (x *Credential) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Credential {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tToken\t%v\n", x.Token)
	fmt.Fprintf(tw, "\tACL\t%v\n", x.ACL)
	fmt.Fprintf(tw, "\tOwnerID\t%v\n", x.OwnerID)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type CredentialList struct {
	// the list of all tunnel credentials on this account
	Credentials []Credential `json:"credentials,omitempty"`
	// URI of the tunnel credential list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *CredentialList) String() string {
	return x.GoString()
}

func (x *CredentialList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "CredentialList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tCredentials\t%v\n", x.Credentials)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointWebhookValidation struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// a string indicating which webhook provider will be sending webhooks to this
	// endpoint. Value must be one of the supported providers defined at
	// https://ngrok.com/docs/cloud-edge/modules/webhook
	// (https://ngrok.com/docs/cloud-edge/modules/webhook)
	Provider string `json:"provider,omitempty"`
	// a string secret used to validate requests from the given provider. All providers
	// except AWS SNS require a secret
	Secret string `json:"secret,omitempty"`
}

func (x *EndpointWebhookValidation) String() string {
	return x.GoString()
}

func (x *EndpointWebhookValidation) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointWebhookValidation {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tProvider\t%v\n", x.Provider)
	fmt.Fprintf(tw, "\tSecret\t%v\n", x.Secret)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointCompression struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
}

func (x *EndpointCompression) String() string {
	return x.GoString()
}

func (x *EndpointCompression) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointCompression {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointMutualTLS struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// PEM-encoded CA certificates that will be used to validate. Multiple CAs may be
	// provided by concatenating them together.
	CertificateAuthorities []Ref `json:"certificate_authorities,omitempty"`
}

func (x *EndpointMutualTLS) String() string {
	return x.GoString()
}

func (x *EndpointMutualTLS) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointMutualTLS {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tCertificateAuthorities\t%v\n", x.CertificateAuthorities)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointMutualTLSMutate struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// list of certificate authorities that will be used to validate the TLS client
	// certificate presented by the initiator of the TLS connection
	CertificateAuthorityIDs []string `json:"certificate_authority_ids,omitempty"`
}

func (x *EndpointMutualTLSMutate) String() string {
	return x.GoString()
}

func (x *EndpointMutualTLSMutate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointMutualTLSMutate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tCertificateAuthorityIDs\t%v\n", x.CertificateAuthorityIDs)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointTLSTermination struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// edge if the ngrok edge should terminate TLS traffic, upstream if TLS traffic
	// should be passed through to the upstream ngrok agent / application server for
	// termination. if upstream is chosen, most other modules will be disallowed
	// because they rely on the ngrok edge being able to access the underlying traffic.
	TerminateAt string `json:"terminate_at,omitempty"`
	// The minimum TLS version used for termination and advertised to the client during
	// the TLS handshake. if unspecified, ngrok will choose an industry-safe default.
	// This value must be null if terminate_at is set to upstream.
	MinVersion *string `json:"min_version,omitempty"`
}

func (x *EndpointTLSTermination) String() string {
	return x.GoString()
}

func (x *EndpointTLSTermination) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointTLSTermination {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tTerminateAt\t%v\n", x.TerminateAt)
	fmt.Fprintf(tw, "\tMinVersion\t%v\n", x.MinVersion)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointTLSTerminationAtEdge struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// The minimum TLS version used for termination and advertised to the client during
	// the TLS handshake. if unspecified, ngrok will choose an industry-safe default.
	// This value must be null if terminate_at is set to upstream.
	MinVersion *string `json:"min_version,omitempty"`
}

func (x *EndpointTLSTerminationAtEdge) String() string {
	return x.GoString()
}

func (x *EndpointTLSTerminationAtEdge) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointTLSTerminationAtEdge {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tMinVersion\t%v\n", x.MinVersion)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointRequestHeaders struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// a map of header key to header value that will be injected into the HTTP Request
	// before being sent to the upstream application server
	Add map[string]string `json:"add,omitempty"`
	// a list of header names that will be removed from the HTTP Request before being
	// sent to the upstream application server
	Remove []string `json:"remove,omitempty"`
}

func (x *EndpointRequestHeaders) String() string {
	return x.GoString()
}

func (x *EndpointRequestHeaders) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointRequestHeaders {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tAdd\t%v\n", x.Add)
	fmt.Fprintf(tw, "\tRemove\t%v\n", x.Remove)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointResponseHeaders struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// a map of header key to header value that will be injected into the HTTP Response
	// returned to the HTTP client
	Add map[string]string `json:"add,omitempty"`
	// a list of header names that will be removed from the HTTP Response returned to
	// the HTTP client
	Remove []string `json:"remove,omitempty"`
}

func (x *EndpointResponseHeaders) String() string {
	return x.GoString()
}

func (x *EndpointResponseHeaders) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointResponseHeaders {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tAdd\t%v\n", x.Add)
	fmt.Fprintf(tw, "\tRemove\t%v\n", x.Remove)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointIPPolicy struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// list of all IP policies that will be used to check if a source IP is allowed
	// access to the endpoint
	IPPolicies []Ref `json:"ip_policies,omitempty"`
}

func (x *EndpointIPPolicy) String() string {
	return x.GoString()
}

func (x *EndpointIPPolicy) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointIPPolicy {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tIPPolicies\t%v\n", x.IPPolicies)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointIPPolicyMutate struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// list of all IP policies that will be used to check if a source IP is allowed
	// access to the endpoint
	IPPolicyIDs []string `json:"ip_policy_ids,omitempty"`
}

func (x *EndpointIPPolicyMutate) String() string {
	return x.GoString()
}

func (x *EndpointIPPolicyMutate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointIPPolicyMutate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tIPPolicyIDs\t%v\n", x.IPPolicyIDs)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointCircuitBreaker struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// Integer number of seconds after which the circuit is tripped to wait before
	// re-evaluating upstream health
	TrippedDuration uint32 `json:"tripped_duration,omitempty"`
	// Integer number of seconds in the statistical rolling window that metrics are
	// retained for.
	RollingWindow uint32 `json:"rolling_window,omitempty"`
	// Integer number of buckets into which metrics are retained. Max 128.
	NumBuckets uint32 `json:"num_buckets,omitempty"`
	// Integer number of requests in a rolling window that will trip the circuit.
	// Helpful if traffic volume is low.
	VolumeThreshold uint32 `json:"volume_threshold,omitempty"`
	// Error threshold percentage should be between 0 - 1.0, not 0-100.0
	ErrorThresholdPercentage float64 `json:"error_threshold_percentage,omitempty"`
}

func (x *EndpointCircuitBreaker) String() string {
	return x.GoString()
}

func (x *EndpointCircuitBreaker) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointCircuitBreaker {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tTrippedDuration\t%v\n", x.TrippedDuration)
	fmt.Fprintf(tw, "\tRollingWindow\t%v\n", x.RollingWindow)
	fmt.Fprintf(tw, "\tNumBuckets\t%v\n", x.NumBuckets)
	fmt.Fprintf(tw, "\tVolumeThreshold\t%v\n", x.VolumeThreshold)
	fmt.Fprintf(tw, "\tErrorThresholdPercentage\t%v\n", x.ErrorThresholdPercentage)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointOAuth struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// an object which defines the identity provider to use for authentication and
	// configuration for who may access the endpoint
	Provider EndpointOAuthProvider `json:"provider,omitempty"`
	// Do not enforce authentication on HTTP OPTIONS requests. necessary if you are
	// supporting CORS.
	OptionsPassthrough bool `json:"options_passthrough,omitempty"`
	// the prefix of the session cookie that ngrok sets on the http client to cache
	// authentication. default is 'ngrok.'
	CookiePrefix string `json:"cookie_prefix,omitempty"`
	// Integer number of seconds of inactivity after which if the user has not accessed
	// the endpoint, their session will time out and they will be forced to
	// reauthenticate.
	InactivityTimeout uint32 `json:"inactivity_timeout,omitempty"`
	// Integer number of seconds of the maximum duration of an authenticated session.
	// After this period is exceeded, a user must reauthenticate.
	MaximumDuration uint32 `json:"maximum_duration,omitempty"`
	// Integer number of seconds after which ngrok guarantees it will refresh user
	// state from the identity provider and recheck whether the user is still
	// authorized to access the endpoint. This is the preferred tunable to use to
	// enforce a minimum amount of time after which a revoked user will no longer be
	// able to access the resource.
	AuthCheckInterval uint32 `json:"auth_check_interval,omitempty"`
}

func (x *EndpointOAuth) String() string {
	return x.GoString()
}

func (x *EndpointOAuth) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointOAuth {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tProvider\t%v\n", x.Provider)
	fmt.Fprintf(tw, "\tOptionsPassthrough\t%v\n", x.OptionsPassthrough)
	fmt.Fprintf(tw, "\tCookiePrefix\t%v\n", x.CookiePrefix)
	fmt.Fprintf(tw, "\tInactivityTimeout\t%v\n", x.InactivityTimeout)
	fmt.Fprintf(tw, "\tMaximumDuration\t%v\n", x.MaximumDuration)
	fmt.Fprintf(tw, "\tAuthCheckInterval\t%v\n", x.AuthCheckInterval)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointOAuthProvider struct {
	// configuration for using github as the identity provider
	Github *EndpointOAuthGitHub `json:"github,omitempty"`
	// configuration for using facebook as the identity provider
	Facebook *EndpointOAuthFacebook `json:"facebook,omitempty"`
	// configuration for using microsoft as the identity provider
	Microsoft *EndpointOAuthMicrosoft `json:"microsoft,omitempty"`
	// configuration for using google as the identity provider
	Google *EndpointOAuthGoogle `json:"google,omitempty"`
	// configuration for using linkedin as the identity provider
	Linkedin *EndpointOAuthLinkedIn `json:"linkedin,omitempty"`
	// configuration for using gitlab as the identity provider
	Gitlab *EndpointOAuthGitLab `json:"gitlab,omitempty"`
	// configuration for using twitch as the identity provider
	Twitch *EndpointOAuthTwitch `json:"twitch,omitempty"`
	// configuration for using amazon as the identity provider
	Amazon *EndpointOAuthAmazon `json:"amazon,omitempty"`
}

func (x *EndpointOAuthProvider) String() string {
	return x.GoString()
}

func (x *EndpointOAuthProvider) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointOAuthProvider {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tGithub\t%v\n", x.Github)
	fmt.Fprintf(tw, "\tFacebook\t%v\n", x.Facebook)
	fmt.Fprintf(tw, "\tMicrosoft\t%v\n", x.Microsoft)
	fmt.Fprintf(tw, "\tGoogle\t%v\n", x.Google)
	fmt.Fprintf(tw, "\tLinkedin\t%v\n", x.Linkedin)
	fmt.Fprintf(tw, "\tGitlab\t%v\n", x.Gitlab)
	fmt.Fprintf(tw, "\tTwitch\t%v\n", x.Twitch)
	fmt.Fprintf(tw, "\tAmazon\t%v\n", x.Amazon)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointOAuthGitHub struct {
	// the OAuth app client ID. retrieve it from the identity provider's dashboard
	// where you created your own OAuth app. optional. if unspecified, ngrok will use
	// its own managed oauth application which has additional restrictions. see the
	// OAuth module docs for more details. if present, client_secret must be present as
	// well.
	ClientID *string `json:"client_id,omitempty"`
	// the OAuth app client secret. retrieve if from the identity provider's dashboard
	// where you created your own OAuth app. optional, see all of the caveats in the
	// docs for client_id.
	ClientSecret *string `json:"client_secret,omitempty"`
	// a list of provider-specific OAuth scopes with the permissions your OAuth app
	// would like to ask for. these may not be set if you are using the ngrok-managed
	// oauth app (i.e. you must pass both client_id and client_secret to set scopes)
	Scopes []string `json:"scopes,omitempty"`
	// a list of email addresses of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailAddresses []string `json:"email_addresses,omitempty"`
	// a list of email domains of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailDomains []string `json:"email_domains,omitempty"`
	// a list of github teams identifiers. users will be allowed access to the endpoint
	// if they are a member of any of these teams. identifiers should be in the 'slug'
	// format qualified with the org name, e.g. org-name/team-name
	Teams []string `json:"teams,omitempty"`
	// a list of github org identifiers. users who are members of any of the listed
	// organizations will be allowed access. identifiers should be the organization's
	// 'slug'
	Organizations []string `json:"organizations,omitempty"`
}

func (x *EndpointOAuthGitHub) String() string {
	return x.GoString()
}

func (x *EndpointOAuthGitHub) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointOAuthGitHub {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tClientID\t%v\n", x.ClientID)
	fmt.Fprintf(tw, "\tClientSecret\t%v\n", x.ClientSecret)
	fmt.Fprintf(tw, "\tScopes\t%v\n", x.Scopes)
	fmt.Fprintf(tw, "\tEmailAddresses\t%v\n", x.EmailAddresses)
	fmt.Fprintf(tw, "\tEmailDomains\t%v\n", x.EmailDomains)
	fmt.Fprintf(tw, "\tTeams\t%v\n", x.Teams)
	fmt.Fprintf(tw, "\tOrganizations\t%v\n", x.Organizations)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointOAuthFacebook struct {
	// the OAuth app client ID. retrieve it from the identity provider's dashboard
	// where you created your own OAuth app. optional. if unspecified, ngrok will use
	// its own managed oauth application which has additional restrictions. see the
	// OAuth module docs for more details. if present, client_secret must be present as
	// well.
	ClientID *string `json:"client_id,omitempty"`
	// the OAuth app client secret. retrieve if from the identity provider's dashboard
	// where you created your own OAuth app. optional, see all of the caveats in the
	// docs for client_id.
	ClientSecret *string `json:"client_secret,omitempty"`
	// a list of provider-specific OAuth scopes with the permissions your OAuth app
	// would like to ask for. these may not be set if you are using the ngrok-managed
	// oauth app (i.e. you must pass both client_id and client_secret to set scopes)
	Scopes []string `json:"scopes,omitempty"`
	// a list of email addresses of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailAddresses []string `json:"email_addresses,omitempty"`
	// a list of email domains of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailDomains []string `json:"email_domains,omitempty"`
}

func (x *EndpointOAuthFacebook) String() string {
	return x.GoString()
}

func (x *EndpointOAuthFacebook) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointOAuthFacebook {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tClientID\t%v\n", x.ClientID)
	fmt.Fprintf(tw, "\tClientSecret\t%v\n", x.ClientSecret)
	fmt.Fprintf(tw, "\tScopes\t%v\n", x.Scopes)
	fmt.Fprintf(tw, "\tEmailAddresses\t%v\n", x.EmailAddresses)
	fmt.Fprintf(tw, "\tEmailDomains\t%v\n", x.EmailDomains)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointOAuthMicrosoft struct {
	// the OAuth app client ID. retrieve it from the identity provider's dashboard
	// where you created your own OAuth app. optional. if unspecified, ngrok will use
	// its own managed oauth application which has additional restrictions. see the
	// OAuth module docs for more details. if present, client_secret must be present as
	// well.
	ClientID *string `json:"client_id,omitempty"`
	// the OAuth app client secret. retrieve if from the identity provider's dashboard
	// where you created your own OAuth app. optional, see all of the caveats in the
	// docs for client_id.
	ClientSecret *string `json:"client_secret,omitempty"`
	// a list of provider-specific OAuth scopes with the permissions your OAuth app
	// would like to ask for. these may not be set if you are using the ngrok-managed
	// oauth app (i.e. you must pass both client_id and client_secret to set scopes)
	Scopes []string `json:"scopes,omitempty"`
	// a list of email addresses of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailAddresses []string `json:"email_addresses,omitempty"`
	// a list of email domains of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailDomains []string `json:"email_domains,omitempty"`
}

func (x *EndpointOAuthMicrosoft) String() string {
	return x.GoString()
}

func (x *EndpointOAuthMicrosoft) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointOAuthMicrosoft {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tClientID\t%v\n", x.ClientID)
	fmt.Fprintf(tw, "\tClientSecret\t%v\n", x.ClientSecret)
	fmt.Fprintf(tw, "\tScopes\t%v\n", x.Scopes)
	fmt.Fprintf(tw, "\tEmailAddresses\t%v\n", x.EmailAddresses)
	fmt.Fprintf(tw, "\tEmailDomains\t%v\n", x.EmailDomains)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointOAuthGoogle struct {
	// the OAuth app client ID. retrieve it from the identity provider's dashboard
	// where you created your own OAuth app. optional. if unspecified, ngrok will use
	// its own managed oauth application which has additional restrictions. see the
	// OAuth module docs for more details. if present, client_secret must be present as
	// well.
	ClientID *string `json:"client_id,omitempty"`
	// the OAuth app client secret. retrieve if from the identity provider's dashboard
	// where you created your own OAuth app. optional, see all of the caveats in the
	// docs for client_id.
	ClientSecret *string `json:"client_secret,omitempty"`
	// a list of provider-specific OAuth scopes with the permissions your OAuth app
	// would like to ask for. these may not be set if you are using the ngrok-managed
	// oauth app (i.e. you must pass both client_id and client_secret to set scopes)
	Scopes []string `json:"scopes,omitempty"`
	// a list of email addresses of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailAddresses []string `json:"email_addresses,omitempty"`
	// a list of email domains of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailDomains []string `json:"email_domains,omitempty"`
}

func (x *EndpointOAuthGoogle) String() string {
	return x.GoString()
}

func (x *EndpointOAuthGoogle) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointOAuthGoogle {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tClientID\t%v\n", x.ClientID)
	fmt.Fprintf(tw, "\tClientSecret\t%v\n", x.ClientSecret)
	fmt.Fprintf(tw, "\tScopes\t%v\n", x.Scopes)
	fmt.Fprintf(tw, "\tEmailAddresses\t%v\n", x.EmailAddresses)
	fmt.Fprintf(tw, "\tEmailDomains\t%v\n", x.EmailDomains)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointOAuthLinkedIn struct {
	ClientID       *string  `json:"client_id,omitempty"`
	ClientSecret   *string  `json:"client_secret,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	EmailDomains   []string `json:"email_domains,omitempty"`
}

func (x *EndpointOAuthLinkedIn) String() string {
	return x.GoString()
}

func (x *EndpointOAuthLinkedIn) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointOAuthLinkedIn {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tClientID\t%v\n", x.ClientID)
	fmt.Fprintf(tw, "\tClientSecret\t%v\n", x.ClientSecret)
	fmt.Fprintf(tw, "\tScopes\t%v\n", x.Scopes)
	fmt.Fprintf(tw, "\tEmailAddresses\t%v\n", x.EmailAddresses)
	fmt.Fprintf(tw, "\tEmailDomains\t%v\n", x.EmailDomains)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointOAuthGitLab struct {
	ClientID       *string  `json:"client_id,omitempty"`
	ClientSecret   *string  `json:"client_secret,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	EmailDomains   []string `json:"email_domains,omitempty"`
}

func (x *EndpointOAuthGitLab) String() string {
	return x.GoString()
}

func (x *EndpointOAuthGitLab) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointOAuthGitLab {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tClientID\t%v\n", x.ClientID)
	fmt.Fprintf(tw, "\tClientSecret\t%v\n", x.ClientSecret)
	fmt.Fprintf(tw, "\tScopes\t%v\n", x.Scopes)
	fmt.Fprintf(tw, "\tEmailAddresses\t%v\n", x.EmailAddresses)
	fmt.Fprintf(tw, "\tEmailDomains\t%v\n", x.EmailDomains)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointOAuthTwitch struct {
	ClientID       *string  `json:"client_id,omitempty"`
	ClientSecret   *string  `json:"client_secret,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	EmailDomains   []string `json:"email_domains,omitempty"`
}

func (x *EndpointOAuthTwitch) String() string {
	return x.GoString()
}

func (x *EndpointOAuthTwitch) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointOAuthTwitch {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tClientID\t%v\n", x.ClientID)
	fmt.Fprintf(tw, "\tClientSecret\t%v\n", x.ClientSecret)
	fmt.Fprintf(tw, "\tScopes\t%v\n", x.Scopes)
	fmt.Fprintf(tw, "\tEmailAddresses\t%v\n", x.EmailAddresses)
	fmt.Fprintf(tw, "\tEmailDomains\t%v\n", x.EmailDomains)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointOAuthAmazon struct {
	ClientID       *string  `json:"client_id,omitempty"`
	ClientSecret   *string  `json:"client_secret,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	EmailDomains   []string `json:"email_domains,omitempty"`
}

func (x *EndpointOAuthAmazon) String() string {
	return x.GoString()
}

func (x *EndpointOAuthAmazon) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointOAuthAmazon {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tClientID\t%v\n", x.ClientID)
	fmt.Fprintf(tw, "\tClientSecret\t%v\n", x.ClientSecret)
	fmt.Fprintf(tw, "\tScopes\t%v\n", x.Scopes)
	fmt.Fprintf(tw, "\tEmailAddresses\t%v\n", x.EmailAddresses)
	fmt.Fprintf(tw, "\tEmailDomains\t%v\n", x.EmailDomains)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointSAML struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// Do not enforce authentication on HTTP OPTIONS requests. necessary if you are
	// supporting CORS.
	OptionsPassthrough bool `json:"options_passthrough,omitempty"`
	// the prefix of the session cookie that ngrok sets on the http client to cache
	// authentication. default is 'ngrok.'
	CookiePrefix string `json:"cookie_prefix,omitempty"`
	// Integer number of seconds of inactivity after which if the user has not accessed
	// the endpoint, their session will time out and they will be forced to
	// reauthenticate.
	InactivityTimeout uint32 `json:"inactivity_timeout,omitempty"`
	// Integer number of seconds of the maximum duration of an authenticated session.
	// After this period is exceeded, a user must reauthenticate.
	MaximumDuration uint32 `json:"maximum_duration,omitempty"`
	// The full XML IdP EntityDescriptor. Your IdP may provide this to you as a a file
	// to download or as a URL.
	IdPMetadata string `json:"idp_metadata,omitempty"`
	// If true, indicates that whenever we redirect a user to the IdP for
	// authentication that the IdP must prompt the user for authentication credentials
	// even if the user already has a valid session with the IdP.
	ForceAuthn bool `json:"force_authn,omitempty"`
	// If true, the IdP may initiate a login directly (e.g. the user does not need to
	// visit the endpoint first and then be redirected). The IdP should set the
	// RelayState parameter to the target URL of the resource they want the user to be
	// redirected to after the SAML login assertion has been processed.
	AllowIdPInitiated *bool `json:"allow_idp_initiated,omitempty"`
	// If present, only users who are a member of one of the listed groups may access
	// the target endpoint.
	AuthorizedGroups []string `json:"authorized_groups,omitempty"`
	// The SP Entity's unique ID. This always takes the form of a URL. In ngrok's
	// implementation, this URL is the same as the metadata URL. This will need to be
	// specified to the IdP as configuration.
	EntityID string `json:"entity_id,omitempty"`
	// The public URL of the SP's Assertion Consumer Service. This is where the IdP
	// will redirect to during an authentication flow. This will need to be specified
	// to the IdP as configuration.
	AssertionConsumerServiceURL string `json:"assertion_consumer_service_url,omitempty"`
	// The public URL of the SP's Single Logout Service. This is where the IdP will
	// redirect to during a single logout flow. This will optionally need to be
	// specified to the IdP as configuration.
	SingleLogoutURL string `json:"single_logout_url,omitempty"`
	// PEM-encoded x.509 certificate of the key pair that is used to sign all SAML
	// requests that the ngrok SP makes to the IdP. Many IdPs do not support request
	// signing verification, but we highly recommend specifying this in the IdP's
	// configuration if it is supported.
	RequestSigningCertificatePEM string `json:"request_signing_certificate_pem,omitempty"`
	// A public URL where the SP's metadata is hosted. If an IdP supports dynamic
	// configuration, this is the URL it can use to retrieve the SP metadata.
	MetadataURL string `json:"metadata_url,omitempty"`
	// Defines the name identifier format the SP expects the IdP to use in its
	// assertions to identify subjects. If unspecified, a default value of
	// urn:oasis:names:tc:SAML:2.0:nameid-format:persistent will be used. A subset of
	// the allowed values enumerated by the SAML specification are supported.
	NameIDFormat string `json:"nameid_format,omitempty"`
}

func (x *EndpointSAML) String() string {
	return x.GoString()
}

func (x *EndpointSAML) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointSAML {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tOptionsPassthrough\t%v\n", x.OptionsPassthrough)
	fmt.Fprintf(tw, "\tCookiePrefix\t%v\n", x.CookiePrefix)
	fmt.Fprintf(tw, "\tInactivityTimeout\t%v\n", x.InactivityTimeout)
	fmt.Fprintf(tw, "\tMaximumDuration\t%v\n", x.MaximumDuration)
	fmt.Fprintf(tw, "\tIdPMetadata\t%v\n", x.IdPMetadata)
	fmt.Fprintf(tw, "\tForceAuthn\t%v\n", x.ForceAuthn)
	fmt.Fprintf(tw, "\tAllowIdPInitiated\t%v\n", x.AllowIdPInitiated)
	fmt.Fprintf(tw, "\tAuthorizedGroups\t%v\n", x.AuthorizedGroups)
	fmt.Fprintf(tw, "\tEntityID\t%v\n", x.EntityID)
	fmt.Fprintf(tw, "\tAssertionConsumerServiceURL\t%v\n", x.AssertionConsumerServiceURL)
	fmt.Fprintf(tw, "\tSingleLogoutURL\t%v\n", x.SingleLogoutURL)
	fmt.Fprintf(tw, "\tRequestSigningCertificatePEM\t%v\n", x.RequestSigningCertificatePEM)
	fmt.Fprintf(tw, "\tMetadataURL\t%v\n", x.MetadataURL)
	fmt.Fprintf(tw, "\tNameIDFormat\t%v\n", x.NameIDFormat)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointSAMLMutate struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// Do not enforce authentication on HTTP OPTIONS requests. necessary if you are
	// supporting CORS.
	OptionsPassthrough bool `json:"options_passthrough,omitempty"`
	// the prefix of the session cookie that ngrok sets on the http client to cache
	// authentication. default is 'ngrok.'
	CookiePrefix string `json:"cookie_prefix,omitempty"`
	// Integer number of seconds of inactivity after which if the user has not accessed
	// the endpoint, their session will time out and they will be forced to
	// reauthenticate.
	InactivityTimeout uint32 `json:"inactivity_timeout,omitempty"`
	// Integer number of seconds of the maximum duration of an authenticated session.
	// After this period is exceeded, a user must reauthenticate.
	MaximumDuration uint32 `json:"maximum_duration,omitempty"`
	// The full XML IdP EntityDescriptor. Your IdP may provide this to you as a a file
	// to download or as a URL.
	IdPMetadata string `json:"idp_metadata,omitempty"`
	// If true, indicates that whenever we redirect a user to the IdP for
	// authentication that the IdP must prompt the user for authentication credentials
	// even if the user already has a valid session with the IdP.
	ForceAuthn bool `json:"force_authn,omitempty"`
	// If true, the IdP may initiate a login directly (e.g. the user does not need to
	// visit the endpoint first and then be redirected). The IdP should set the
	// RelayState parameter to the target URL of the resource they want the user to be
	// redirected to after the SAML login assertion has been processed.
	AllowIdPInitiated *bool `json:"allow_idp_initiated,omitempty"`
	// If present, only users who are a member of one of the listed groups may access
	// the target endpoint.
	AuthorizedGroups []string `json:"authorized_groups,omitempty"`
	// Defines the name identifier format the SP expects the IdP to use in its
	// assertions to identify subjects. If unspecified, a default value of
	// urn:oasis:names:tc:SAML:2.0:nameid-format:persistent will be used. A subset of
	// the allowed values enumerated by the SAML specification are supported.
	NameIDFormat string `json:"nameid_format,omitempty"`
}

func (x *EndpointSAMLMutate) String() string {
	return x.GoString()
}

func (x *EndpointSAMLMutate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointSAMLMutate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tOptionsPassthrough\t%v\n", x.OptionsPassthrough)
	fmt.Fprintf(tw, "\tCookiePrefix\t%v\n", x.CookiePrefix)
	fmt.Fprintf(tw, "\tInactivityTimeout\t%v\n", x.InactivityTimeout)
	fmt.Fprintf(tw, "\tMaximumDuration\t%v\n", x.MaximumDuration)
	fmt.Fprintf(tw, "\tIdPMetadata\t%v\n", x.IdPMetadata)
	fmt.Fprintf(tw, "\tForceAuthn\t%v\n", x.ForceAuthn)
	fmt.Fprintf(tw, "\tAllowIdPInitiated\t%v\n", x.AllowIdPInitiated)
	fmt.Fprintf(tw, "\tAuthorizedGroups\t%v\n", x.AuthorizedGroups)
	fmt.Fprintf(tw, "\tNameIDFormat\t%v\n", x.NameIDFormat)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointOIDC struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// Do not enforce authentication on HTTP OPTIONS requests. necessary if you are
	// supporting CORS.
	OptionsPassthrough bool `json:"options_passthrough,omitempty"`
	// the prefix of the session cookie that ngrok sets on the http client to cache
	// authentication. default is 'ngrok.'
	CookiePrefix string `json:"cookie_prefix,omitempty"`
	// Integer number of seconds of inactivity after which if the user has not accessed
	// the endpoint, their session will time out and they will be forced to
	// reauthenticate.
	InactivityTimeout uint32 `json:"inactivity_timeout,omitempty"`
	// Integer number of seconds of the maximum duration of an authenticated session.
	// After this period is exceeded, a user must reauthenticate.
	MaximumDuration uint32 `json:"maximum_duration,omitempty"`
	// URL of the OIDC "OpenID provider". This is the base URL used for discovery.
	Issuer string `json:"issuer,omitempty"`
	// The OIDC app's client ID and OIDC audience.
	ClientID string `json:"client_id,omitempty"`
	// The OIDC app's client secret.
	ClientSecret string `json:"client_secret,omitempty"`
	// The set of scopes to request from the OIDC identity provider.
	Scopes []string `json:"scopes,omitempty"`
}

func (x *EndpointOIDC) String() string {
	return x.GoString()
}

func (x *EndpointOIDC) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointOIDC {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tOptionsPassthrough\t%v\n", x.OptionsPassthrough)
	fmt.Fprintf(tw, "\tCookiePrefix\t%v\n", x.CookiePrefix)
	fmt.Fprintf(tw, "\tInactivityTimeout\t%v\n", x.InactivityTimeout)
	fmt.Fprintf(tw, "\tMaximumDuration\t%v\n", x.MaximumDuration)
	fmt.Fprintf(tw, "\tIssuer\t%v\n", x.Issuer)
	fmt.Fprintf(tw, "\tClientID\t%v\n", x.ClientID)
	fmt.Fprintf(tw, "\tClientSecret\t%v\n", x.ClientSecret)
	fmt.Fprintf(tw, "\tScopes\t%v\n", x.Scopes)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointBackend struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// backend to be used to back this endpoint
	Backend Ref `json:"backend,omitempty"`
}

func (x *EndpointBackend) String() string {
	return x.GoString()
}

func (x *EndpointBackend) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointBackend {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tBackend\t%v\n", x.Backend)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointBackendMutate struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// backend to be used to back this endpoint
	BackendID string `json:"backend_id,omitempty"`
}

func (x *EndpointBackendMutate) String() string {
	return x.GoString()
}

func (x *EndpointBackendMutate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointBackendMutate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	fmt.Fprintf(tw, "\tBackendID\t%v\n", x.BackendID)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointWebsocketTCPConverter struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
}

func (x *EndpointWebsocketTCPConverter) String() string {
	return x.GoString()
}

func (x *EndpointWebsocketTCPConverter) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointWebsocketTCPConverter {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEnabled\t%v\n", x.Enabled)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteItem struct {
	// unique identifier of this edge
	EdgeID string `json:"edge_id,omitempty"`
	// unique identifier of this edge route
	ID string `json:"id,omitempty"`
}

func (x *EdgeRouteItem) String() string {
	return fmt.Sprintf("EdgeRouteItem{ID: %v}", x.ID)

}

func (x *EdgeRouteItem) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteItem {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type HTTPSEdgeRouteCreate struct {
	// unique identifier of this edge
	EdgeID string `json:"edge_id,omitempty"`
	// Type of match to use for this route. Valid values are "exact_path" and
	// "path_prefix".
	MatchType string `json:"match_type,omitempty"`
	// Route selector: "/blog" or "example.com" or "example.com/blog"
	Match string `json:"match,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// backend module configuration or null
	Backend *EndpointBackendMutate `json:"backend,omitempty"`
	// ip restriction module configuration or null
	IPRestriction *EndpointIPPolicyMutate `json:"ip_restriction,omitempty"`
	// circuit breaker module configuration or null
	CircuitBreaker *EndpointCircuitBreaker `json:"circuit_breaker,omitempty"`
	// compression module configuration or null
	Compression *EndpointCompression `json:"compression,omitempty"`
	// request headers module configuration or null
	RequestHeaders *EndpointRequestHeaders `json:"request_headers,omitempty"`
	// response headers module configuration or null
	ResponseHeaders *EndpointResponseHeaders `json:"response_headers,omitempty"`
	// webhook verification module configuration or null
	WebhookVerification *EndpointWebhookValidation `json:"webhook_verification,omitempty"`
	// oauth module configuration or null
	OAuth *EndpointOAuth `json:"oauth,omitempty"`
	// saml module configuration or null
	SAML *EndpointSAMLMutate `json:"saml,omitempty"`
	// oidc module configuration or null
	OIDC *EndpointOIDC `json:"oidc,omitempty"`
	// websocket to tcp adapter configuration or null
	WebsocketTCPConverter *EndpointWebsocketTCPConverter `json:"websocket_tcp_converter,omitempty"`
}

func (x *HTTPSEdgeRouteCreate) String() string {
	return x.GoString()
}

func (x *HTTPSEdgeRouteCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "HTTPSEdgeRouteCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tMatchType\t%v\n", x.MatchType)
	fmt.Fprintf(tw, "\tMatch\t%v\n", x.Match)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBackend\t%v\n", x.Backend)
	fmt.Fprintf(tw, "\tIPRestriction\t%v\n", x.IPRestriction)
	fmt.Fprintf(tw, "\tCircuitBreaker\t%v\n", x.CircuitBreaker)
	fmt.Fprintf(tw, "\tCompression\t%v\n", x.Compression)
	fmt.Fprintf(tw, "\tRequestHeaders\t%v\n", x.RequestHeaders)
	fmt.Fprintf(tw, "\tResponseHeaders\t%v\n", x.ResponseHeaders)
	fmt.Fprintf(tw, "\tWebhookVerification\t%v\n", x.WebhookVerification)
	fmt.Fprintf(tw, "\tOAuth\t%v\n", x.OAuth)
	fmt.Fprintf(tw, "\tSAML\t%v\n", x.SAML)
	fmt.Fprintf(tw, "\tOIDC\t%v\n", x.OIDC)
	fmt.Fprintf(tw, "\tWebsocketTCPConverter\t%v\n", x.WebsocketTCPConverter)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type HTTPSEdgeRouteUpdate struct {
	// unique identifier of this edge
	EdgeID string `json:"edge_id,omitempty"`
	// unique identifier of this edge route
	ID string `json:"id,omitempty"`
	// Type of match to use for this route. Valid values are "exact_path" and
	// "path_prefix".
	MatchType string `json:"match_type,omitempty"`
	// Route selector: "/blog" or "example.com" or "example.com/blog"
	Match string `json:"match,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// backend module configuration or null
	Backend *EndpointBackendMutate `json:"backend,omitempty"`
	// ip restriction module configuration or null
	IPRestriction *EndpointIPPolicyMutate `json:"ip_restriction,omitempty"`
	// circuit breaker module configuration or null
	CircuitBreaker *EndpointCircuitBreaker `json:"circuit_breaker,omitempty"`
	// compression module configuration or null
	Compression *EndpointCompression `json:"compression,omitempty"`
	// request headers module configuration or null
	RequestHeaders *EndpointRequestHeaders `json:"request_headers,omitempty"`
	// response headers module configuration or null
	ResponseHeaders *EndpointResponseHeaders `json:"response_headers,omitempty"`
	// webhook verification module configuration or null
	WebhookVerification *EndpointWebhookValidation `json:"webhook_verification,omitempty"`
	// oauth module configuration or null
	OAuth *EndpointOAuth `json:"oauth,omitempty"`
	// saml module configuration or null
	SAML *EndpointSAMLMutate `json:"saml,omitempty"`
	// oidc module configuration or null
	OIDC *EndpointOIDC `json:"oidc,omitempty"`
	// websocket to tcp adapter configuration or null
	WebsocketTCPConverter *EndpointWebsocketTCPConverter `json:"websocket_tcp_converter,omitempty"`
}

func (x *HTTPSEdgeRouteUpdate) String() string {
	return fmt.Sprintf("HTTPSEdgeRouteUpdate{ID: %v}", x.ID)

}

func (x *HTTPSEdgeRouteUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "HTTPSEdgeRouteUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tMatchType\t%v\n", x.MatchType)
	fmt.Fprintf(tw, "\tMatch\t%v\n", x.Match)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBackend\t%v\n", x.Backend)
	fmt.Fprintf(tw, "\tIPRestriction\t%v\n", x.IPRestriction)
	fmt.Fprintf(tw, "\tCircuitBreaker\t%v\n", x.CircuitBreaker)
	fmt.Fprintf(tw, "\tCompression\t%v\n", x.Compression)
	fmt.Fprintf(tw, "\tRequestHeaders\t%v\n", x.RequestHeaders)
	fmt.Fprintf(tw, "\tResponseHeaders\t%v\n", x.ResponseHeaders)
	fmt.Fprintf(tw, "\tWebhookVerification\t%v\n", x.WebhookVerification)
	fmt.Fprintf(tw, "\tOAuth\t%v\n", x.OAuth)
	fmt.Fprintf(tw, "\tSAML\t%v\n", x.SAML)
	fmt.Fprintf(tw, "\tOIDC\t%v\n", x.OIDC)
	fmt.Fprintf(tw, "\tWebsocketTCPConverter\t%v\n", x.WebsocketTCPConverter)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type HTTPSEdgeRoute struct {
	// unique identifier of this edge
	EdgeID string `json:"edge_id,omitempty"`
	// unique identifier of this edge route
	ID string `json:"id,omitempty"`
	// timestamp when the edge configuration was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// Type of match to use for this route. Valid values are "exact_path" and
	// "path_prefix".
	MatchType string `json:"match_type,omitempty"`
	// Route selector: "/blog" or "example.com" or "example.com/blog"
	Match string `json:"match,omitempty"`
	// URI of the edge API resource
	URI string `json:"uri,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// backend module configuration or null
	Backend *EndpointBackend `json:"backend,omitempty"`
	// ip restriction module configuration or null
	IpRestriction *EndpointIPPolicy `json:"ip_restriction,omitempty"`
	// circuit breaker module configuration or null
	CircuitBreaker *EndpointCircuitBreaker `json:"circuit_breaker,omitempty"`
	// compression module configuration or null
	Compression *EndpointCompression `json:"compression,omitempty"`
	// request headers module configuration or null
	RequestHeaders *EndpointRequestHeaders `json:"request_headers,omitempty"`
	// response headers module configuration or null
	ResponseHeaders *EndpointResponseHeaders `json:"response_headers,omitempty"`
	// webhook verification module configuration or null
	WebhookVerification *EndpointWebhookValidation `json:"webhook_verification,omitempty"`
	// oauth module configuration or null
	OAuth *EndpointOAuth `json:"oauth,omitempty"`
	// saml module configuration or null
	SAML *EndpointSAML `json:"saml,omitempty"`
	// oidc module configuration or null
	OIDC *EndpointOIDC `json:"oidc,omitempty"`
	// websocket to tcp adapter configuration or null
	WebsocketTCPConverter *EndpointWebsocketTCPConverter `json:"websocket_tcp_converter,omitempty"`
}

func (x *HTTPSEdgeRoute) String() string {
	return fmt.Sprintf("HTTPSEdgeRoute{ID: %v}", x.ID)

}

func (x *HTTPSEdgeRoute) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "HTTPSEdgeRoute {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tMatchType\t%v\n", x.MatchType)
	fmt.Fprintf(tw, "\tMatch\t%v\n", x.Match)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tBackend\t%v\n", x.Backend)
	fmt.Fprintf(tw, "\tIpRestriction\t%v\n", x.IpRestriction)
	fmt.Fprintf(tw, "\tCircuitBreaker\t%v\n", x.CircuitBreaker)
	fmt.Fprintf(tw, "\tCompression\t%v\n", x.Compression)
	fmt.Fprintf(tw, "\tRequestHeaders\t%v\n", x.RequestHeaders)
	fmt.Fprintf(tw, "\tResponseHeaders\t%v\n", x.ResponseHeaders)
	fmt.Fprintf(tw, "\tWebhookVerification\t%v\n", x.WebhookVerification)
	fmt.Fprintf(tw, "\tOAuth\t%v\n", x.OAuth)
	fmt.Fprintf(tw, "\tSAML\t%v\n", x.SAML)
	fmt.Fprintf(tw, "\tOIDC\t%v\n", x.OIDC)
	fmt.Fprintf(tw, "\tWebsocketTCPConverter\t%v\n", x.WebsocketTCPConverter)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type HTTPSEdgeList struct {
	// the list of all HTTPS Edges on this account
	HTTPSEdges []HTTPSEdge `json:"https_edges,omitempty"`
	// URI of the HTTPS Edge list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *HTTPSEdgeList) String() string {
	return x.GoString()
}

func (x *HTTPSEdgeList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "HTTPSEdgeList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tHTTPSEdges\t%v\n", x.HTTPSEdges)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type HTTPSEdgeCreate struct {
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge; optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports []string `json:"hostports,omitempty"`
	// edge modules
	MutualTLS      *EndpointMutualTLSMutate      `json:"mutual_tls,omitempty"`
	TLSTermination *EndpointTLSTerminationAtEdge `json:"tls_termination,omitempty"`
}

func (x *HTTPSEdgeCreate) String() string {
	return x.GoString()
}

func (x *HTTPSEdgeCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "HTTPSEdgeCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tHostports\t%v\n", x.Hostports)
	fmt.Fprintf(tw, "\tMutualTLS\t%v\n", x.MutualTLS)
	fmt.Fprintf(tw, "\tTLSTermination\t%v\n", x.TLSTermination)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type HTTPSEdgeUpdate struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge; optional, max 4096
	// bytes.
	Metadata *string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports []string `json:"hostports,omitempty"`
	// edge modules
	MutualTLS      *EndpointMutualTLSMutate      `json:"mutual_tls,omitempty"`
	TLSTermination *EndpointTLSTerminationAtEdge `json:"tls_termination,omitempty"`
}

func (x *HTTPSEdgeUpdate) String() string {
	return fmt.Sprintf("HTTPSEdgeUpdate{ID: %v}", x.ID)

}

func (x *HTTPSEdgeUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "HTTPSEdgeUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tHostports\t%v\n", x.Hostports)
	fmt.Fprintf(tw, "\tMutualTLS\t%v\n", x.MutualTLS)
	fmt.Fprintf(tw, "\tTLSTermination\t%v\n", x.TLSTermination)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type HTTPSEdge struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge; optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// timestamp when the edge configuration was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// URI of the edge API resource
	URI string `json:"uri,omitempty"`
	// hostports served by this edge
	Hostports []string `json:"hostports,omitempty"`
	// edge modules
	MutualTls      *EndpointMutualTLS      `json:"mutual_tls,omitempty"`
	TlsTermination *EndpointTLSTermination `json:"tls_termination,omitempty"`
	// routes
	Routes []HTTPSEdgeRoute `json:"routes,omitempty"`
}

func (x *HTTPSEdge) String() string {
	return fmt.Sprintf("HTTPSEdge{ID: %v}", x.ID)

}

func (x *HTTPSEdge) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "HTTPSEdge {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tHostports\t%v\n", x.Hostports)
	fmt.Fprintf(tw, "\tMutualTls\t%v\n", x.MutualTls)
	fmt.Fprintf(tw, "\tTlsTermination\t%v\n", x.TlsTermination)
	fmt.Fprintf(tw, "\tRoutes\t%v\n", x.Routes)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeBackendReplace struct {
	ID     string                `json:"id,omitempty"`
	Module EndpointBackendMutate `json:"module,omitempty"`
}

func (x *EdgeBackendReplace) String() string {
	return fmt.Sprintf("EdgeBackendReplace{ID: %v}", x.ID)

}

func (x *EdgeBackendReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeBackendReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeIPRestrictionReplace struct {
	ID     string                 `json:"id,omitempty"`
	Module EndpointIPPolicyMutate `json:"module,omitempty"`
}

func (x *EdgeIPRestrictionReplace) String() string {
	return fmt.Sprintf("EdgeIPRestrictionReplace{ID: %v}", x.ID)

}

func (x *EdgeIPRestrictionReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeIPRestrictionReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeMutualTLSReplace struct {
	ID     string                  `json:"id,omitempty"`
	Module EndpointMutualTLSMutate `json:"module,omitempty"`
}

func (x *EdgeMutualTLSReplace) String() string {
	return fmt.Sprintf("EdgeMutualTLSReplace{ID: %v}", x.ID)

}

func (x *EdgeMutualTLSReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeMutualTLSReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeTLSTerminationReplace struct {
	ID     string                 `json:"id,omitempty"`
	Module EndpointTLSTermination `json:"module,omitempty"`
}

func (x *EdgeTLSTerminationReplace) String() string {
	return fmt.Sprintf("EdgeTLSTerminationReplace{ID: %v}", x.ID)

}

func (x *EdgeTLSTerminationReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeTLSTerminationReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeTLSTerminationAtEdgeReplace struct {
	ID     string                       `json:"id,omitempty"`
	Module EndpointTLSTerminationAtEdge `json:"module,omitempty"`
}

func (x *EdgeTLSTerminationAtEdgeReplace) String() string {
	return fmt.Sprintf("EdgeTLSTerminationAtEdgeReplace{ID: %v}", x.ID)

}

func (x *EdgeTLSTerminationAtEdgeReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeTLSTerminationAtEdgeReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteBackendReplace struct {
	EdgeID string                `json:"edge_id,omitempty"`
	ID     string                `json:"id,omitempty"`
	Module EndpointBackendMutate `json:"module,omitempty"`
}

func (x *EdgeRouteBackendReplace) String() string {
	return fmt.Sprintf("EdgeRouteBackendReplace{ID: %v}", x.ID)

}

func (x *EdgeRouteBackendReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteBackendReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteIPRestrictionReplace struct {
	EdgeID string                 `json:"edge_id,omitempty"`
	ID     string                 `json:"id,omitempty"`
	Module EndpointIPPolicyMutate `json:"module,omitempty"`
}

func (x *EdgeRouteIPRestrictionReplace) String() string {
	return fmt.Sprintf("EdgeRouteIPRestrictionReplace{ID: %v}", x.ID)

}

func (x *EdgeRouteIPRestrictionReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteIPRestrictionReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteRequestHeadersReplace struct {
	EdgeID string                 `json:"edge_id,omitempty"`
	ID     string                 `json:"id,omitempty"`
	Module EndpointRequestHeaders `json:"module,omitempty"`
}

func (x *EdgeRouteRequestHeadersReplace) String() string {
	return fmt.Sprintf("EdgeRouteRequestHeadersReplace{ID: %v}", x.ID)

}

func (x *EdgeRouteRequestHeadersReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteRequestHeadersReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteResponseHeadersReplace struct {
	EdgeID string                  `json:"edge_id,omitempty"`
	ID     string                  `json:"id,omitempty"`
	Module EndpointResponseHeaders `json:"module,omitempty"`
}

func (x *EdgeRouteResponseHeadersReplace) String() string {
	return fmt.Sprintf("EdgeRouteResponseHeadersReplace{ID: %v}", x.ID)

}

func (x *EdgeRouteResponseHeadersReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteResponseHeadersReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteCompressionReplace struct {
	EdgeID string              `json:"edge_id,omitempty"`
	ID     string              `json:"id,omitempty"`
	Module EndpointCompression `json:"module,omitempty"`
}

func (x *EdgeRouteCompressionReplace) String() string {
	return fmt.Sprintf("EdgeRouteCompressionReplace{ID: %v}", x.ID)

}

func (x *EdgeRouteCompressionReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteCompressionReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteCircuitBreakerReplace struct {
	EdgeID string                 `json:"edge_id,omitempty"`
	ID     string                 `json:"id,omitempty"`
	Module EndpointCircuitBreaker `json:"module,omitempty"`
}

func (x *EdgeRouteCircuitBreakerReplace) String() string {
	return fmt.Sprintf("EdgeRouteCircuitBreakerReplace{ID: %v}", x.ID)

}

func (x *EdgeRouteCircuitBreakerReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteCircuitBreakerReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteWebhookVerificationReplace struct {
	EdgeID string                    `json:"edge_id,omitempty"`
	ID     string                    `json:"id,omitempty"`
	Module EndpointWebhookValidation `json:"module,omitempty"`
}

func (x *EdgeRouteWebhookVerificationReplace) String() string {
	return fmt.Sprintf("EdgeRouteWebhookVerificationReplace{ID: %v}", x.ID)

}

func (x *EdgeRouteWebhookVerificationReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteWebhookVerificationReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteOAuthReplace struct {
	EdgeID string        `json:"edge_id,omitempty"`
	ID     string        `json:"id,omitempty"`
	Module EndpointOAuth `json:"module,omitempty"`
}

func (x *EdgeRouteOAuthReplace) String() string {
	return fmt.Sprintf("EdgeRouteOAuthReplace{ID: %v}", x.ID)

}

func (x *EdgeRouteOAuthReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteOAuthReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteSAMLReplace struct {
	EdgeID string             `json:"edge_id,omitempty"`
	ID     string             `json:"id,omitempty"`
	Module EndpointSAMLMutate `json:"module,omitempty"`
}

func (x *EdgeRouteSAMLReplace) String() string {
	return fmt.Sprintf("EdgeRouteSAMLReplace{ID: %v}", x.ID)

}

func (x *EdgeRouteSAMLReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteSAMLReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteOIDCReplace struct {
	EdgeID string       `json:"edge_id,omitempty"`
	ID     string       `json:"id,omitempty"`
	Module EndpointOIDC `json:"module,omitempty"`
}

func (x *EdgeRouteOIDCReplace) String() string {
	return fmt.Sprintf("EdgeRouteOIDCReplace{ID: %v}", x.ID)

}

func (x *EdgeRouteOIDCReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteOIDCReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EdgeRouteWebsocketTCPConverterReplace struct {
	EdgeID string                        `json:"edge_id,omitempty"`
	ID     string                        `json:"id,omitempty"`
	Module EndpointWebsocketTCPConverter `json:"module,omitempty"`
}

func (x *EdgeRouteWebsocketTCPConverterReplace) String() string {
	return fmt.Sprintf("EdgeRouteWebsocketTCPConverterReplace{ID: %v}", x.ID)

}

func (x *EdgeRouteWebsocketTCPConverterReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EdgeRouteWebsocketTCPConverterReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEdgeID\t%v\n", x.EdgeID)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tModule\t%v\n", x.Module)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TCPEdgeList struct {
	// the list of all TCP Edges on this account
	TCPEdges []TCPEdge `json:"tcp_edges,omitempty"`
	// URI of the TCP Edge list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *TCPEdgeList) String() string {
	return x.GoString()
}

func (x *TCPEdgeList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TCPEdgeList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tTCPEdges\t%v\n", x.TCPEdges)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TCPEdgeCreate struct {
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports []string `json:"hostports,omitempty"`
	// edge modules
	Backend       *EndpointBackendMutate  `json:"backend,omitempty"`
	IPRestriction *EndpointIPPolicyMutate `json:"ip_restriction,omitempty"`
}

func (x *TCPEdgeCreate) String() string {
	return x.GoString()
}

func (x *TCPEdgeCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TCPEdgeCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tHostports\t%v\n", x.Hostports)
	fmt.Fprintf(tw, "\tBackend\t%v\n", x.Backend)
	fmt.Fprintf(tw, "\tIPRestriction\t%v\n", x.IPRestriction)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TCPEdgeUpdate struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata *string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports []string `json:"hostports,omitempty"`
	// edge modules
	Backend       *EndpointBackendMutate  `json:"backend,omitempty"`
	IPRestriction *EndpointIPPolicyMutate `json:"ip_restriction,omitempty"`
}

func (x *TCPEdgeUpdate) String() string {
	return fmt.Sprintf("TCPEdgeUpdate{ID: %v}", x.ID)

}

func (x *TCPEdgeUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TCPEdgeUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tHostports\t%v\n", x.Hostports)
	fmt.Fprintf(tw, "\tBackend\t%v\n", x.Backend)
	fmt.Fprintf(tw, "\tIPRestriction\t%v\n", x.IPRestriction)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TCPEdge struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// timestamp when the edge was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// URI of the edge API resource
	URI string `json:"uri,omitempty"`
	// hostports served by this edge
	Hostports []string `json:"hostports,omitempty"`
	// edge modules
	Backend       *EndpointBackend  `json:"backend,omitempty"`
	IpRestriction *EndpointIPPolicy `json:"ip_restriction,omitempty"`
}

func (x *TCPEdge) String() string {
	return fmt.Sprintf("TCPEdge{ID: %v}", x.ID)

}

func (x *TCPEdge) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TCPEdge {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tHostports\t%v\n", x.Hostports)
	fmt.Fprintf(tw, "\tBackend\t%v\n", x.Backend)
	fmt.Fprintf(tw, "\tIpRestriction\t%v\n", x.IpRestriction)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TLSEdgeList struct {
	// the list of all TLS Edges on this account
	TLSEdges []TLSEdge `json:"tls_edges,omitempty"`
	// URI of the TLS Edge list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *TLSEdgeList) String() string {
	return x.GoString()
}

func (x *TLSEdgeList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TLSEdgeList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tTLSEdges\t%v\n", x.TLSEdges)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TLSEdgeCreate struct {
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports []string `json:"hostports,omitempty"`
	// edge modules
	Backend        *EndpointBackendMutate   `json:"backend,omitempty"`
	IPRestriction  *EndpointIPPolicyMutate  `json:"ip_restriction,omitempty"`
	MutualTLS      *EndpointMutualTLSMutate `json:"mutual_tls,omitempty"`
	TLSTermination *EndpointTLSTermination  `json:"tls_termination,omitempty"`
}

func (x *TLSEdgeCreate) String() string {
	return x.GoString()
}

func (x *TLSEdgeCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TLSEdgeCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tHostports\t%v\n", x.Hostports)
	fmt.Fprintf(tw, "\tBackend\t%v\n", x.Backend)
	fmt.Fprintf(tw, "\tIPRestriction\t%v\n", x.IPRestriction)
	fmt.Fprintf(tw, "\tMutualTLS\t%v\n", x.MutualTLS)
	fmt.Fprintf(tw, "\tTLSTermination\t%v\n", x.TLSTermination)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TLSEdgeUpdate struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata *string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports []string `json:"hostports,omitempty"`
	// edge modules
	Backend        *EndpointBackendMutate   `json:"backend,omitempty"`
	IPRestriction  *EndpointIPPolicyMutate  `json:"ip_restriction,omitempty"`
	MutualTLS      *EndpointMutualTLSMutate `json:"mutual_tls,omitempty"`
	TLSTermination *EndpointTLSTermination  `json:"tls_termination,omitempty"`
}

func (x *TLSEdgeUpdate) String() string {
	return fmt.Sprintf("TLSEdgeUpdate{ID: %v}", x.ID)

}

func (x *TLSEdgeUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TLSEdgeUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tHostports\t%v\n", x.Hostports)
	fmt.Fprintf(tw, "\tBackend\t%v\n", x.Backend)
	fmt.Fprintf(tw, "\tIPRestriction\t%v\n", x.IPRestriction)
	fmt.Fprintf(tw, "\tMutualTLS\t%v\n", x.MutualTLS)
	fmt.Fprintf(tw, "\tTLSTermination\t%v\n", x.TLSTermination)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TLSEdge struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// timestamp when the edge configuration was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// URI of the edge API resource
	URI string `json:"uri,omitempty"`
	// hostports served by this edge
	Hostports []string `json:"hostports,omitempty"`
	// edge modules
	Backend        *EndpointBackend        `json:"backend,omitempty"`
	IpRestriction  *EndpointIPPolicy       `json:"ip_restriction,omitempty"`
	MutualTls      *EndpointMutualTLS      `json:"mutual_tls,omitempty"`
	TlsTermination *EndpointTLSTermination `json:"tls_termination,omitempty"`
}

func (x *TLSEdge) String() string {
	return fmt.Sprintf("TLSEdge{ID: %v}", x.ID)

}

func (x *TLSEdge) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TLSEdge {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tHostports\t%v\n", x.Hostports)
	fmt.Fprintf(tw, "\tBackend\t%v\n", x.Backend)
	fmt.Fprintf(tw, "\tIpRestriction\t%v\n", x.IpRestriction)
	fmt.Fprintf(tw, "\tMutualTls\t%v\n", x.MutualTls)
	fmt.Fprintf(tw, "\tTlsTermination\t%v\n", x.TlsTermination)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type Endpoint struct {
	// unique endpoint resource identifier
	ID string `json:"id,omitempty"`
	// identifier of the region this endpoint belongs to
	Region string `json:"region,omitempty"`
	// timestamp when the endpoint was created in RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// timestamp when the endpoint was updated in RFC 3339 format
	UpdatedAt string `json:"updated_at,omitempty"`
	// URL of the hostport served by this endpoint
	PublicURL string `json:"public_url,omitempty"`
	// protocol served by this endpoint. one of http, https, tcp, or tls
	Proto string `json:"proto,omitempty"`
	// hostport served by this endpoint (hostname:port)
	Hostport string `json:"hostport,omitempty"`
	// whether the endpoint is ephemeral (served directly by an agent-initiated tunnel)
	// or edge (served by an edge)
	Type string `json:"type,omitempty"`
	// user-supplied metadata of the associated tunnel or edge object
	Metadata string `json:"metadata,omitempty"`
	// the domain reserved for this endpoint
	Domain *Ref `json:"domain,omitempty"`
	// the address reserved for this endpoint
	TCPAddr *Ref `json:"tcp_addr,omitempty"`
	// the tunnel serving requests to this endpoint, if this is an ephemeral endpoint
	Tunnel *Ref `json:"tunnel,omitempty"`
	// the edge serving requests to this endpoint, if this is an edge endpoint
	Edge *Ref `json:"edge,omitempty"`
}

func (x *Endpoint) String() string {
	return fmt.Sprintf("Endpoint{ID: %v}", x.ID)

}

func (x *Endpoint) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Endpoint {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tRegion\t%v\n", x.Region)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tUpdatedAt\t%v\n", x.UpdatedAt)
	fmt.Fprintf(tw, "\tPublicURL\t%v\n", x.PublicURL)
	fmt.Fprintf(tw, "\tProto\t%v\n", x.Proto)
	fmt.Fprintf(tw, "\tHostport\t%v\n", x.Hostport)
	fmt.Fprintf(tw, "\tType\t%v\n", x.Type)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tDomain\t%v\n", x.Domain)
	fmt.Fprintf(tw, "\tTCPAddr\t%v\n", x.TCPAddr)
	fmt.Fprintf(tw, "\tTunnel\t%v\n", x.Tunnel)
	fmt.Fprintf(tw, "\tEdge\t%v\n", x.Edge)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EndpointList struct {
	// the list of all active endpoints on this account
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	// URI of the endpoints list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *EndpointList) String() string {
	return x.GoString()
}

func (x *EndpointList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EndpointList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEndpoints\t%v\n", x.Endpoints)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventDestinationCreate struct {
	// Arbitrary user-defined machine-readable data of this Event Destination.
	// Optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// Human-readable description of the Event Destination. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// The output format you would like to serialize events into when sending to their
	// target. Currently the only accepted value is JSON.
	Format string `json:"format,omitempty"`
	// An object that encapsulates where and how to send your events. An event
	// destination must contain exactly one of the following objects, leaving the rest
	// null: kinesis, firehose, cloudwatch_logs, or s3.
	Target EventTarget `json:"target,omitempty"`
}

func (x *EventDestinationCreate) String() string {
	return x.GoString()
}

func (x *EventDestinationCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventDestinationCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tFormat\t%v\n", x.Format)
	fmt.Fprintf(tw, "\tTarget\t%v\n", x.Target)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventDestinationUpdate struct {
	// Unique identifier for this Event Destination.
	ID string `json:"id,omitempty"`
	// Arbitrary user-defined machine-readable data of this Event Destination.
	// Optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// Human-readable description of the Event Destination. Optional, max 255 bytes.
	Description *string `json:"description,omitempty"`
	// The output format you would like to serialize events into when sending to their
	// target. Currently the only accepted value is JSON.
	Format *string `json:"format,omitempty"`
	// An object that encapsulates where and how to send your events. An event
	// destination must contain exactly one of the following objects, leaving the rest
	// null: kinesis, firehose, cloudwatch_logs, or s3.
	Target *EventTarget `json:"target,omitempty"`
}

func (x *EventDestinationUpdate) String() string {
	return fmt.Sprintf("EventDestinationUpdate{ID: %v}", x.ID)

}

func (x *EventDestinationUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventDestinationUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tFormat\t%v\n", x.Format)
	fmt.Fprintf(tw, "\tTarget\t%v\n", x.Target)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventDestination struct {
	// Unique identifier for this Event Destination.
	ID string `json:"id,omitempty"`
	// Arbitrary user-defined machine-readable data of this Event Destination.
	// Optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// Timestamp when the Event Destination was created, RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// Human-readable description of the Event Destination. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// The output format you would like to serialize events into when sending to their
	// target. Currently the only accepted value is JSON.
	Format string `json:"format,omitempty"`
	// An object that encapsulates where and how to send your events. An event
	// destination must contain exactly one of the following objects, leaving the rest
	// null: kinesis, firehose, cloudwatch_logs, or s3.
	Target EventTarget `json:"target,omitempty"`
	// URI of the Event Destination API resource.
	URI string `json:"uri,omitempty"`
}

func (x *EventDestination) String() string {
	return fmt.Sprintf("EventDestination{ID: %v}", x.ID)

}

func (x *EventDestination) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventDestination {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tFormat\t%v\n", x.Format)
	fmt.Fprintf(tw, "\tTarget\t%v\n", x.Target)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventDestinationList struct {
	// The list of all Event Destinations on this account.
	EventDestinations []EventDestination `json:"event_destinations,omitempty"`
	// URI of the Event Destinations list API resource.
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page.
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *EventDestinationList) String() string {
	return x.GoString()
}

func (x *EventDestinationList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventDestinationList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEventDestinations\t%v\n", x.EventDestinations)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventTarget struct {
	// Configuration used to send events to Amazon Kinesis Data Firehose.
	Firehose *EventTargetFirehose `json:"firehose,omitempty"`
	// Configuration used to send events to Amazon Kinesis.
	Kinesis *EventTargetKinesis `json:"kinesis,omitempty"`
	// Configuration used to send events to Amazon CloudWatch Logs.
	CloudwatchLogs *EventTargetCloudwatchLogs `json:"cloudwatch_logs,omitempty"`
	// Configuration used to send events to Datadog.
	Datadog *EventTargetDatadog `json:"datadog,omitempty"`
}

func (x *EventTarget) String() string {
	return x.GoString()
}

func (x *EventTarget) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventTarget {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tFirehose\t%v\n", x.Firehose)
	fmt.Fprintf(tw, "\tKinesis\t%v\n", x.Kinesis)
	fmt.Fprintf(tw, "\tCloudwatchLogs\t%v\n", x.CloudwatchLogs)
	fmt.Fprintf(tw, "\tDatadog\t%v\n", x.Datadog)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventTargetFirehose struct {
	// Configuration for how to authenticate into your AWS account. Exactly one of role
	// or creds should be configured.
	Auth AWSAuth `json:"auth,omitempty"`
	// An Amazon Resource Name specifying the Firehose delivery stream to deposit
	// events into.
	DeliveryStreamARN string `json:"delivery_stream_arn,omitempty"`
}

func (x *EventTargetFirehose) String() string {
	return x.GoString()
}

func (x *EventTargetFirehose) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventTargetFirehose {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tAuth\t%v\n", x.Auth)
	fmt.Fprintf(tw, "\tDeliveryStreamARN\t%v\n", x.DeliveryStreamARN)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventTargetKinesis struct {
	// Configuration for how to authenticate into your AWS account. Exactly one of role
	// or creds should be configured.
	Auth AWSAuth `json:"auth,omitempty"`
	// An Amazon Resource Name specifying the Kinesis stream to deposit events into.
	StreamARN string `json:"stream_arn,omitempty"`
}

func (x *EventTargetKinesis) String() string {
	return x.GoString()
}

func (x *EventTargetKinesis) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventTargetKinesis {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tAuth\t%v\n", x.Auth)
	fmt.Fprintf(tw, "\tStreamARN\t%v\n", x.StreamARN)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventTargetCloudwatchLogs struct {
	// Configuration for how to authenticate into your AWS account. Exactly one of role
	// or creds should be configured.
	Auth AWSAuth `json:"auth,omitempty"`
	// An Amazon Resource Name specifying the CloudWatch Logs group to deposit events
	// into.
	LogGroupARN string `json:"log_group_arn,omitempty"`
}

func (x *EventTargetCloudwatchLogs) String() string {
	return x.GoString()
}

func (x *EventTargetCloudwatchLogs) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventTargetCloudwatchLogs {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tAuth\t%v\n", x.Auth)
	fmt.Fprintf(tw, "\tLogGroupARN\t%v\n", x.LogGroupARN)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventTargetDatadog struct {
	// Datadog API key to use.
	ApiKey *string `json:"api_key,omitempty"`
	// Tags to send with the event.
	Ddtags *string `json:"ddtags,omitempty"`
	// Service name to send with the event.
	Service *string `json:"service,omitempty"`
	// Datadog site to send event to.
	Ddsite *string `json:"ddsite,omitempty"`
}

func (x *EventTargetDatadog) String() string {
	return x.GoString()
}

func (x *EventTargetDatadog) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventTargetDatadog {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tApiKey\t%v\n", x.ApiKey)
	fmt.Fprintf(tw, "\tDdtags\t%v\n", x.Ddtags)
	fmt.Fprintf(tw, "\tService\t%v\n", x.Service)
	fmt.Fprintf(tw, "\tDdsite\t%v\n", x.Ddsite)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AWSAuth struct {
	// A role for ngrok to assume on your behalf to deposit events into your AWS
	// account.
	Role *AWSRole `json:"role,omitempty"`
	// Credentials to your AWS account if you prefer ngrok to sign in with long-term
	// access keys.
	Creds *AWSCredentials `json:"creds,omitempty"`
}

func (x *AWSAuth) String() string {
	return x.GoString()
}

func (x *AWSAuth) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AWSAuth {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tRole\t%v\n", x.Role)
	fmt.Fprintf(tw, "\tCreds\t%v\n", x.Creds)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AWSRole struct {
	// An ARN that specifies the role that ngrok should use to deliver to the
	// configured target.
	RoleARN string `json:"role_arn,omitempty"`
}

func (x *AWSRole) String() string {
	return x.GoString()
}

func (x *AWSRole) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AWSRole {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tRoleARN\t%v\n", x.RoleARN)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type AWSCredentials struct {
	// The ID portion of an AWS access key.
	AWSAccessKeyID string `json:"aws_access_key_id,omitempty"`
	// The secret portion of an AWS access key.
	AWSSecretAccessKey *string `json:"aws_secret_access_key,omitempty"`
}

func (x *AWSCredentials) String() string {
	return x.GoString()
}

func (x *AWSCredentials) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "AWSCredentials {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tAWSAccessKeyID\t%v\n", x.AWSAccessKeyID)
	fmt.Fprintf(tw, "\tAWSSecretAccessKey\t%v\n", x.AWSSecretAccessKey)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventSubscriptionCreate struct {
	// Arbitrary customer supplied information intended to be machine readable.
	// Optional, max 4096 chars.
	Metadata string `json:"metadata,omitempty"`
	// Arbitrary customer supplied information intended to be human readable. Optional,
	// max 255 chars.
	Description string `json:"description,omitempty"`
	// Sources containing the types for which this event subscription will trigger
	Sources []EventSourceReplace `json:"sources,omitempty"`
	// A list of Event Destination IDs which should be used for this Event
	// Subscription.
	DestinationIDs []string `json:"destination_ids,omitempty"`
}

func (x *EventSubscriptionCreate) String() string {
	return x.GoString()
}

func (x *EventSubscriptionCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventSubscriptionCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tSources\t%v\n", x.Sources)
	fmt.Fprintf(tw, "\tDestinationIDs\t%v\n", x.DestinationIDs)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventSubscriptionUpdate struct {
	// Unique identifier for this Event Subscription.
	ID string `json:"id,omitempty"`
	// Arbitrary customer supplied information intended to be machine readable.
	// Optional, max 4096 chars.
	Metadata *string `json:"metadata,omitempty"`
	// Arbitrary customer supplied information intended to be human readable. Optional,
	// max 255 chars.
	Description *string `json:"description,omitempty"`
	// Sources containing the types for which this event subscription will trigger
	Sources []EventSourceReplace `json:"sources,omitempty"`
	// A list of Event Destination IDs which should be used for this Event
	// Subscription.
	DestinationIDs []string `json:"destination_ids,omitempty"`
}

func (x *EventSubscriptionUpdate) String() string {
	return fmt.Sprintf("EventSubscriptionUpdate{ID: %v}", x.ID)

}

func (x *EventSubscriptionUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventSubscriptionUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tSources\t%v\n", x.Sources)
	fmt.Fprintf(tw, "\tDestinationIDs\t%v\n", x.DestinationIDs)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventSubscriptionList struct {
	// The list of all Event Subscriptions on this account.
	EventSubscriptions []EventSubscription `json:"event_subscriptions,omitempty"`
	// URI of the Event Subscriptions list API resource.
	URI string `json:"uri,omitempty"`
	// URI of next page, or null if there is no next page.
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *EventSubscriptionList) String() string {
	return x.GoString()
}

func (x *EventSubscriptionList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventSubscriptionList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tEventSubscriptions\t%v\n", x.EventSubscriptions)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventSubscription struct {
	// Unique identifier for this Event Subscription.
	ID string `json:"id,omitempty"`
	// URI of the Event Subscription API resource.
	URI string `json:"uri,omitempty"`
	// When the Event Subscription was created (RFC 3339 format).
	CreatedAt string `json:"created_at,omitempty"`
	// Arbitrary customer supplied information intended to be machine readable.
	// Optional, max 4096 chars.
	Metadata string `json:"metadata,omitempty"`
	// Arbitrary customer supplied information intended to be human readable. Optional,
	// max 255 chars.
	Description string `json:"description,omitempty"`
	// Sources containing the types for which this event subscription will trigger
	Sources []EventSource `json:"sources,omitempty"`
	// Destinations to which these events will be sent
	Destinations []Ref `json:"destinations,omitempty"`
}

func (x *EventSubscription) String() string {
	return fmt.Sprintf("EventSubscription{ID: %v}", x.ID)

}

func (x *EventSubscription) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventSubscription {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tSources\t%v\n", x.Sources)
	fmt.Fprintf(tw, "\tDestinations\t%v\n", x.Destinations)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventSourceReplace struct {
	// Type of event for which an event subscription will trigger
	Type string `json:"type,omitempty"`
}

func (x *EventSourceReplace) String() string {
	return x.GoString()
}

func (x *EventSourceReplace) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventSourceReplace {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tType\t%v\n", x.Type)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventSource struct {
	// Type of event for which an event subscription will trigger
	Type string `json:"type,omitempty"`
	// URI of the Event Source API resource.
	URI string `json:"uri,omitempty"`
}

func (x *EventSource) String() string {
	return x.GoString()
}

func (x *EventSource) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventSource {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tType\t%v\n", x.Type)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventSourceList struct {
	// The list of all Event Sources for an Event Subscription
	Sources []EventSource `json:"sources,omitempty"`
	// URI of the next page, or null if there is no next page.
	URI string `json:"uri,omitempty"`
}

func (x *EventSourceList) String() string {
	return x.GoString()
}

func (x *EventSourceList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventSourceList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tSources\t%v\n", x.Sources)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventSourceCreate struct {
	// The unique identifier for the Event Subscription that this Event Source is
	// attached to.
	SubscriptionID string `json:"subscription_id,omitempty"`
	// Type of event for which an event subscription will trigger
	Type string `json:"type,omitempty"`
}

func (x *EventSourceCreate) String() string {
	return x.GoString()
}

func (x *EventSourceCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventSourceCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tSubscriptionID\t%v\n", x.SubscriptionID)
	fmt.Fprintf(tw, "\tType\t%v\n", x.Type)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type EventSourceUpdate struct {
	// The unique identifier for the Event Subscription that this Event Source is
	// attached to.
	SubscriptionID string `json:"subscription_id,omitempty"`
	// Type of event for which an event subscription will trigger
	Type string `json:"type,omitempty"`
}

func (x *EventSourceUpdate) String() string {
	return x.GoString()
}

func (x *EventSourceUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventSourceUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tSubscriptionID\t%v\n", x.SubscriptionID)
	fmt.Fprintf(tw, "\tType\t%v\n", x.Type)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

// This is needed instead of Item because the parameters are different.
type EventSourceItem struct {
	// The unique identifier for the Event Subscription that this Event Source is
	// attached to.
	SubscriptionID string `json:"subscription_id,omitempty"`
	// Type of event for which an event subscription will trigger
	Type string `json:"type,omitempty"`
}

func (x *EventSourceItem) String() string {
	return x.GoString()
}

func (x *EventSourceItem) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventSourceItem {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tSubscriptionID\t%v\n", x.SubscriptionID)
	fmt.Fprintf(tw, "\tType\t%v\n", x.Type)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

// This is needed instead of Paging because the parameters are different. We also don't need the typical pagination params because pagination of this isn't necessary or supported.
type EventSourcePaging struct {
	// The unique identifier for the Event Subscription that this Event Source is
	// attached to.
	SubscriptionID string `json:"subscription_id,omitempty"`
}

func (x *EventSourcePaging) String() string {
	return x.GoString()
}

func (x *EventSourcePaging) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "EventSourcePaging {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tSubscriptionID\t%v\n", x.SubscriptionID)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPPolicyCreate struct {
	// human-readable description of the source IPs of this IP policy. optional, max
	// 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy. optional, max
	// 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
}

func (x *IPPolicyCreate) String() string {
	return x.GoString()
}

func (x *IPPolicyCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPPolicyCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPPolicyUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of the source IPs of this IP policy. optional, max
	// 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy. optional, max
	// 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

func (x *IPPolicyUpdate) String() string {
	return fmt.Sprintf("IPPolicyUpdate{ID: %v}", x.ID)

}

func (x *IPPolicyUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPPolicyUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPPolicy struct {
	// unique identifier for this IP policy
	ID string `json:"id,omitempty"`
	// URI of the IP Policy API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the IP policy was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of the source IPs of this IP policy. optional, max
	// 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy. optional, max
	// 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
}

func (x *IPPolicy) String() string {
	return fmt.Sprintf("IPPolicy{ID: %v}", x.ID)

}

func (x *IPPolicy) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPPolicy {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPPolicyList struct {
	// the list of all IP policies on this account
	IPPolicies []IPPolicy `json:"ip_policies,omitempty"`
	// URI of the IP policy list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *IPPolicyList) String() string {
	return x.GoString()
}

func (x *IPPolicyList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPPolicyList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tIPPolicies\t%v\n", x.IPPolicies)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPPolicyRuleCreate struct {
	// human-readable description of the source IPs of this IP rule. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy rule. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// an IP or IP range specified in CIDR notation. IPv4 and IPv6 are both supported.
	CIDR string `json:"cidr,omitempty"`
	// ID of the IP policy this IP policy rule will be attached to
	IPPolicyID string `json:"ip_policy_id,omitempty"`
	// the action to apply to the policy rule, either allow or deny
	Action *string `json:"action,omitempty"`
}

func (x *IPPolicyRuleCreate) String() string {
	return x.GoString()
}

func (x *IPPolicyRuleCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPPolicyRuleCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCIDR\t%v\n", x.CIDR)
	fmt.Fprintf(tw, "\tIPPolicyID\t%v\n", x.IPPolicyID)
	fmt.Fprintf(tw, "\tAction\t%v\n", x.Action)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPPolicyRuleUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of the source IPs of this IP rule. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy rule. optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// an IP or IP range specified in CIDR notation. IPv4 and IPv6 are both supported.
	CIDR *string `json:"cidr,omitempty"`
}

func (x *IPPolicyRuleUpdate) String() string {
	return fmt.Sprintf("IPPolicyRuleUpdate{ID: %v}", x.ID)

}

func (x *IPPolicyRuleUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPPolicyRuleUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCIDR\t%v\n", x.CIDR)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPPolicyRule struct {
	// unique identifier for this IP policy rule
	ID string `json:"id,omitempty"`
	// URI of the IP policy rule API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the IP policy rule was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of the source IPs of this IP rule. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy rule. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// an IP or IP range specified in CIDR notation. IPv4 and IPv6 are both supported.
	CIDR string `json:"cidr,omitempty"`
	// object describing the IP policy this IP Policy Rule belongs to
	IPPolicy Ref `json:"ip_policy,omitempty"`
	// the action to apply to the policy rule, either allow or deny
	Action string `json:"action,omitempty"`
}

func (x *IPPolicyRule) String() string {
	return fmt.Sprintf("IPPolicyRule{ID: %v}", x.ID)

}

func (x *IPPolicyRule) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPPolicyRule {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCIDR\t%v\n", x.CIDR)
	fmt.Fprintf(tw, "\tIPPolicy\t%v\n", x.IPPolicy)
	fmt.Fprintf(tw, "\tAction\t%v\n", x.Action)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPPolicyRuleList struct {
	// the list of all IP policy rules on this account
	IPPolicyRules []IPPolicyRule `json:"ip_policy_rules,omitempty"`
	// URI of the IP policy rule list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *IPPolicyRuleList) String() string {
	return x.GoString()
}

func (x *IPPolicyRuleList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPPolicyRuleList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tIPPolicyRules\t%v\n", x.IPPolicyRules)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPRestrictionCreate struct {
	// human-readable description of this IP restriction. optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP restriction. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// true if the IP restriction will be enforced. if false, only warnings will be
	// issued
	Enforced bool `json:"enforced,omitempty"`
	// the type of IP restriction. this defines what traffic will be restricted with
	// the attached policies. four values are currently supported: dashboard, api,
	// agent, and endpoints
	Type string `json:"type,omitempty"`
	// the set of IP policy identifiers that are used to enforce the restriction
	IPPolicyIDs []string `json:"ip_policy_ids,omitempty"`
}

func (x *IPRestrictionCreate) String() string {
	return x.GoString()
}

func (x *IPRestrictionCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPRestrictionCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tEnforced\t%v\n", x.Enforced)
	fmt.Fprintf(tw, "\tType\t%v\n", x.Type)
	fmt.Fprintf(tw, "\tIPPolicyIDs\t%v\n", x.IPPolicyIDs)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPRestrictionUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this IP restriction. optional, max 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP restriction. optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// true if the IP restriction will be enforced. if false, only warnings will be
	// issued
	Enforced *bool `json:"enforced,omitempty"`
	// the set of IP policy identifiers that are used to enforce the restriction
	IPPolicyIDs []string `json:"ip_policy_ids,omitempty"`
}

func (x *IPRestrictionUpdate) String() string {
	return fmt.Sprintf("IPRestrictionUpdate{ID: %v}", x.ID)

}

func (x *IPRestrictionUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPRestrictionUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tEnforced\t%v\n", x.Enforced)
	fmt.Fprintf(tw, "\tIPPolicyIDs\t%v\n", x.IPPolicyIDs)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPRestriction struct {
	// unique identifier for this IP restriction
	ID string `json:"id,omitempty"`
	// URI of the IP restriction API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the IP restriction was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this IP restriction. optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP restriction. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// true if the IP restriction will be enforced. if false, only warnings will be
	// issued
	Enforced bool `json:"enforced,omitempty"`
	// the type of IP restriction. this defines what traffic will be restricted with
	// the attached policies. four values are currently supported: dashboard, api,
	// agent, and endpoints
	Type string `json:"type,omitempty"`
	// the set of IP policies that are used to enforce the restriction
	IPPolicies []Ref `json:"ip_policies,omitempty"`
}

func (x *IPRestriction) String() string {
	return fmt.Sprintf("IPRestriction{ID: %v}", x.ID)

}

func (x *IPRestriction) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPRestriction {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tEnforced\t%v\n", x.Enforced)
	fmt.Fprintf(tw, "\tType\t%v\n", x.Type)
	fmt.Fprintf(tw, "\tIPPolicies\t%v\n", x.IPPolicies)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type IPRestrictionList struct {
	// the list of all IP restrictions on this account
	IPRestrictions []IPRestriction `json:"ip_restrictions,omitempty"`
	// URI of the IP restrictions list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *IPRestrictionList) String() string {
	return x.GoString()
}

func (x *IPRestrictionList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "IPRestrictionList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tIPRestrictions\t%v\n", x.IPRestrictions)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ReservedAddrCreate struct {
	// human-readable description of what this reserved address will be used for
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved address. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// reserve the address in this geographic ngrok datacenter. Optional, default is
	// us. (au, eu, ap, us, jp, in, sa)
	Region string `json:"region,omitempty"`
}

func (x *ReservedAddrCreate) String() string {
	return x.GoString()
}

func (x *ReservedAddrCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ReservedAddrCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tRegion\t%v\n", x.Region)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ReservedAddrUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of what this reserved address will be used for
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved address. Optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

func (x *ReservedAddrUpdate) String() string {
	return fmt.Sprintf("ReservedAddrUpdate{ID: %v}", x.ID)

}

func (x *ReservedAddrUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ReservedAddrUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ReservedAddr struct {
	// unique reserved address resource identifier
	ID string `json:"id,omitempty"`
	// URI of the reserved address API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the reserved address was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of what this reserved address will be used for
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved address. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// hostname:port of the reserved address that was assigned at creation time
	Addr string `json:"addr,omitempty"`
	// reserve the address in this geographic ngrok datacenter. Optional, default is
	// us. (au, eu, ap, us, jp, in, sa)
	Region string `json:"region,omitempty"`
}

func (x *ReservedAddr) String() string {
	return fmt.Sprintf("ReservedAddr{ID: %v}", x.ID)

}

func (x *ReservedAddr) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ReservedAddr {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tAddr\t%v\n", x.Addr)
	fmt.Fprintf(tw, "\tRegion\t%v\n", x.Region)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ReservedAddrList struct {
	// the list of all reserved addresses on this account
	ReservedAddrs []ReservedAddr `json:"reserved_addrs,omitempty"`
	// URI of the reserved address list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *ReservedAddrList) String() string {
	return x.GoString()
}

func (x *ReservedAddrList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ReservedAddrList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tReservedAddrs\t%v\n", x.ReservedAddrs)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ReservedDomainCreate struct {
	// hostname of the reserved domain
	Domain string `json:"domain,omitempty"`
	// reserve the domain in this geographic ngrok datacenter. Optional, default is us.
	// (au, eu, ap, us, jp, in, sa)
	Region string `json:"region,omitempty"`
	// human-readable description of what this reserved domain will be used for
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved domain. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// ID of a user-uploaded TLS certificate to use for connections to targeting this
	// domain. Optional, mutually exclusive with certificate_management_policy.
	CertificateID *string `json:"certificate_id,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled. Optional, mutually exclusive with
	// certificate_id.
	CertificateManagementPolicy *ReservedDomainCertPolicy `json:"certificate_management_policy,omitempty"`
}

func (x *ReservedDomainCreate) String() string {
	return x.GoString()
}

func (x *ReservedDomainCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ReservedDomainCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDomain\t%v\n", x.Domain)
	fmt.Fprintf(tw, "\tRegion\t%v\n", x.Region)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCertificateID\t%v\n", x.CertificateID)
	fmt.Fprintf(tw, "\tCertificateManagementPolicy\t%v\n", x.CertificateManagementPolicy)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ReservedDomainUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of what this reserved domain will be used for
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved domain. Optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// ID of a user-uploaded TLS certificate to use for connections to targeting this
	// domain. Optional, mutually exclusive with certificate_management_policy.
	CertificateID *string `json:"certificate_id,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled. Optional, mutually exclusive with
	// certificate_id.
	CertificateManagementPolicy *ReservedDomainCertPolicy `json:"certificate_management_policy,omitempty"`
}

func (x *ReservedDomainUpdate) String() string {
	return fmt.Sprintf("ReservedDomainUpdate{ID: %v}", x.ID)

}

func (x *ReservedDomainUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ReservedDomainUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCertificateID\t%v\n", x.CertificateID)
	fmt.Fprintf(tw, "\tCertificateManagementPolicy\t%v\n", x.CertificateManagementPolicy)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ReservedDomain struct {
	// unique reserved domain resource identifier
	ID string `json:"id,omitempty"`
	// URI of the reserved domain API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the reserved domain was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of what this reserved domain will be used for
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved domain. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// hostname of the reserved domain
	Domain string `json:"domain,omitempty"`
	// reserve the domain in this geographic ngrok datacenter. Optional, default is us.
	// (au, eu, ap, us, jp, in, sa)
	Region string `json:"region,omitempty"`
	// DNS CNAME target for a custom hostname, or null if the reserved domain is a
	// subdomain of *.ngrok.io
	CNAMETarget *string `json:"cname_target,omitempty"`
	// object referencing the TLS certificate used for connections to this domain. This
	// can be either a user-uploaded certificate, the most recently issued automatic
	// one, or null otherwise.
	Certificate *Ref `json:"certificate,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled
	CertificateManagementPolicy *ReservedDomainCertPolicy `json:"certificate_management_policy,omitempty"`
	// status of the automatic certificate management for this domain, or null if
	// automatic management is disabled
	CertificateManagementStatus *ReservedDomainCertStatus `json:"certificate_management_status,omitempty"`
	// DNS CNAME target for the host _acme-challenge.example.com, where example.com is
	// your reserved domain name. This is required to issue certificates for wildcard,
	// non-ngrok reserved domains. Must be null for non-wildcard domains and ngrok
	// subdomains.
	ACMEChallengeCNAMETarget *string `json:"acme_challenge_cname_target,omitempty"`
}

func (x *ReservedDomain) String() string {
	return fmt.Sprintf("ReservedDomain{ID: %v}", x.ID)

}

func (x *ReservedDomain) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ReservedDomain {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tDomain\t%v\n", x.Domain)
	fmt.Fprintf(tw, "\tRegion\t%v\n", x.Region)
	fmt.Fprintf(tw, "\tCNAMETarget\t%v\n", x.CNAMETarget)
	fmt.Fprintf(tw, "\tCertificate\t%v\n", x.Certificate)
	fmt.Fprintf(tw, "\tCertificateManagementPolicy\t%v\n", x.CertificateManagementPolicy)
	fmt.Fprintf(tw, "\tCertificateManagementStatus\t%v\n", x.CertificateManagementStatus)
	fmt.Fprintf(tw, "\tACMEChallengeCNAMETarget\t%v\n", x.ACMEChallengeCNAMETarget)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ReservedDomainList struct {
	// the list of all reserved domains on this account
	ReservedDomains []ReservedDomain `json:"reserved_domains,omitempty"`
	// URI of the reserved domain list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *ReservedDomainList) String() string {
	return x.GoString()
}

func (x *ReservedDomainList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ReservedDomainList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tReservedDomains\t%v\n", x.ReservedDomains)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ReservedDomainCertPolicy struct {
	// certificate authority to request certificates from. The only supported value is
	// letsencrypt.
	Authority string `json:"authority,omitempty"`
	// type of private key to use when requesting certificates. Defaults to rsa, can be
	// either rsa or ecdsa.
	PrivateKeyType string `json:"private_key_type,omitempty"`
}

func (x *ReservedDomainCertPolicy) String() string {
	return x.GoString()
}

func (x *ReservedDomainCertPolicy) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ReservedDomainCertPolicy {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tAuthority\t%v\n", x.Authority)
	fmt.Fprintf(tw, "\tPrivateKeyType\t%v\n", x.PrivateKeyType)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ReservedDomainCertStatus struct {
	// timestamp when the next renewal will be requested, RFC 3339 format
	RenewsAt *string `json:"renews_at,omitempty"`
	// status of the certificate provisioning job, or null if the certificiate isn't
	// being provisioned or renewed
	ProvisioningJob *ReservedDomainCertJob `json:"provisioning_job,omitempty"`
}

func (x *ReservedDomainCertStatus) String() string {
	return x.GoString()
}

func (x *ReservedDomainCertStatus) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ReservedDomainCertStatus {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tRenewsAt\t%v\n", x.RenewsAt)
	fmt.Fprintf(tw, "\tProvisioningJob\t%v\n", x.ProvisioningJob)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type ReservedDomainCertJob struct {
	// if present, an error code indicating why provisioning is failing. It may be
	// either a temporary condition (INTERNAL_ERROR), or a permanent one the user must
	// correct (DNS_ERROR).
	ErrorCode *string `json:"error_code,omitempty"`
	// a message describing the current status or error
	Msg string `json:"msg,omitempty"`
	// timestamp when the provisioning job started, RFC 3339 format
	StartedAt string `json:"started_at,omitempty"`
	// timestamp when the provisioning job will be retried
	RetriesAt *string `json:"retries_at,omitempty"`
}

func (x *ReservedDomainCertJob) String() string {
	return x.GoString()
}

func (x *ReservedDomainCertJob) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ReservedDomainCertJob {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tErrorCode\t%v\n", x.ErrorCode)
	fmt.Fprintf(tw, "\tMsg\t%v\n", x.Msg)
	fmt.Fprintf(tw, "\tStartedAt\t%v\n", x.StartedAt)
	fmt.Fprintf(tw, "\tRetriesAt\t%v\n", x.RetriesAt)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHCertificateAuthorityCreate struct {
	// human-readable description of this SSH Certificate Authority. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Certificate Authority.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// the type of private key to generate. one of rsa, ecdsa, ed25519
	PrivateKeyType string `json:"private_key_type,omitempty"`
	// the type of elliptic curve to use when creating an ECDSA key
	EllipticCurve string `json:"elliptic_curve,omitempty"`
	// the key size to use when creating an RSA key. one of 2048 or 4096
	KeySize int64 `json:"key_size,omitempty"`
}

func (x *SSHCertificateAuthorityCreate) String() string {
	return x.GoString()
}

func (x *SSHCertificateAuthorityCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHCertificateAuthorityCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tPrivateKeyType\t%v\n", x.PrivateKeyType)
	fmt.Fprintf(tw, "\tEllipticCurve\t%v\n", x.EllipticCurve)
	fmt.Fprintf(tw, "\tKeySize\t%v\n", x.KeySize)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHCertificateAuthorityUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this SSH Certificate Authority. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Certificate Authority.
	// optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

func (x *SSHCertificateAuthorityUpdate) String() string {
	return fmt.Sprintf("SSHCertificateAuthorityUpdate{ID: %v}", x.ID)

}

func (x *SSHCertificateAuthorityUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHCertificateAuthorityUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHCertificateAuthority struct {
	// unique identifier for this SSH Certificate Authority
	ID string `json:"id,omitempty"`
	// URI of the SSH Certificate Authority API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the SSH Certificate Authority API resource was created, RFC 3339
	// format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this SSH Certificate Authority. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Certificate Authority.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// raw public key for this SSH Certificate Authority
	PublicKey string `json:"public_key,omitempty"`
	// the type of private key for this SSH Certificate Authority
	KeyType string `json:"key_type,omitempty"`
}

func (x *SSHCertificateAuthority) String() string {
	return fmt.Sprintf("SSHCertificateAuthority{ID: %v}", x.ID)

}

func (x *SSHCertificateAuthority) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHCertificateAuthority {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tPublicKey\t%v\n", x.PublicKey)
	fmt.Fprintf(tw, "\tKeyType\t%v\n", x.KeyType)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHCertificateAuthorityList struct {
	// the list of all certificate authorities on this account
	SSHCertificateAuthorities []SSHCertificateAuthority `json:"ssh_certificate_authorities,omitempty"`
	// URI of the certificates authorities list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *SSHCertificateAuthorityList) String() string {
	return x.GoString()
}

func (x *SSHCertificateAuthorityList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHCertificateAuthorityList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tSSHCertificateAuthorities\t%v\n", x.SSHCertificateAuthorities)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHCredentialCreate struct {
	// human-readable description of who or what will use the ssh credential to
	// authenticate. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this ssh credential. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL []string `json:"acl,omitempty"`
	// the PEM-encoded public key of the SSH keypair that will be used to authenticate
	PublicKey string `json:"public_key,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
}

func (x *SSHCredentialCreate) String() string {
	return x.GoString()
}

func (x *SSHCredentialCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHCredentialCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tACL\t%v\n", x.ACL)
	fmt.Fprintf(tw, "\tPublicKey\t%v\n", x.PublicKey)
	fmt.Fprintf(tw, "\tOwnerID\t%v\n", x.OwnerID)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHCredentialUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of who or what will use the ssh credential to
	// authenticate. Optional, max 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this ssh credential. Optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL []string `json:"acl,omitempty"`
}

func (x *SSHCredentialUpdate) String() string {
	return fmt.Sprintf("SSHCredentialUpdate{ID: %v}", x.ID)

}

func (x *SSHCredentialUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHCredentialUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tACL\t%v\n", x.ACL)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHCredential struct {
	// unique ssh credential resource identifier
	ID string `json:"id,omitempty"`
	// URI of the ssh credential API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the ssh credential was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of who or what will use the ssh credential to
	// authenticate. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this ssh credential. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// the PEM-encoded public key of the SSH keypair that will be used to authenticate
	PublicKey string `json:"public_key,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL []string `json:"acl,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
}

func (x *SSHCredential) String() string {
	return fmt.Sprintf("SSHCredential{ID: %v}", x.ID)

}

func (x *SSHCredential) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHCredential {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tPublicKey\t%v\n", x.PublicKey)
	fmt.Fprintf(tw, "\tACL\t%v\n", x.ACL)
	fmt.Fprintf(tw, "\tOwnerID\t%v\n", x.OwnerID)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHCredentialList struct {
	// the list of all ssh credentials on this account
	SSHCredentials []SSHCredential `json:"ssh_credentials,omitempty"`
	// URI of the ssh credential list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *SSHCredentialList) String() string {
	return x.GoString()
}

func (x *SSHCredentialList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHCredentialList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tSSHCredentials\t%v\n", x.SSHCredentials)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHHostCertificateCreate struct {
	// the ssh certificate authority that is used to sign this ssh host certificate
	SSHCertificateAuthorityID string `json:"ssh_certificate_authority_id,omitempty"`
	// a public key in OpenSSH Authorized Keys format that this certificate signs
	PublicKey string `json:"public_key,omitempty"`
	// the list of principals included in the ssh host certificate. This is the list of
	// hostnames and/or IP addresses that are authorized to serve SSH traffic with this
	// certificate. Dangerously, if no principals are specified, this certificate is
	// considered valid for all hosts.
	Principals []string `json:"principals,omitempty"`
	// The time when the host certificate becomes valid, in RFC 3339 format. Defaults
	// to the current time if unspecified.
	ValidAfter string `json:"valid_after,omitempty"`
	// The time when this host certificate becomes invalid, in RFC 3339 format. If
	// unspecified, a default value of one year in the future will be used. The OpenSSH
	// certificates RFC calls this valid_before.
	ValidUntil string `json:"valid_until,omitempty"`
	// human-readable description of this SSH Host Certificate. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Host Certificate.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
}

func (x *SSHHostCertificateCreate) String() string {
	return x.GoString()
}

func (x *SSHHostCertificateCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHHostCertificateCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tSSHCertificateAuthorityID\t%v\n", x.SSHCertificateAuthorityID)
	fmt.Fprintf(tw, "\tPublicKey\t%v\n", x.PublicKey)
	fmt.Fprintf(tw, "\tPrincipals\t%v\n", x.Principals)
	fmt.Fprintf(tw, "\tValidAfter\t%v\n", x.ValidAfter)
	fmt.Fprintf(tw, "\tValidUntil\t%v\n", x.ValidUntil)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHHostCertificateUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this SSH Host Certificate. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Host Certificate.
	// optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

func (x *SSHHostCertificateUpdate) String() string {
	return fmt.Sprintf("SSHHostCertificateUpdate{ID: %v}", x.ID)

}

func (x *SSHHostCertificateUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHHostCertificateUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHHostCertificate struct {
	// unique identifier for this SSH Host Certificate
	ID string `json:"id,omitempty"`
	// URI of the SSH Host Certificate API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the SSH Host Certificate API resource was created, RFC 3339
	// format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this SSH Host Certificate. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Host Certificate.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// a public key in OpenSSH Authorized Keys format that this certificate signs
	PublicKey string `json:"public_key,omitempty"`
	// the key type of the public_key, one of rsa, ecdsa or ed25519
	KeyType string `json:"key_type,omitempty"`
	// the ssh certificate authority that is used to sign this ssh host certificate
	SSHCertificateAuthorityID string `json:"ssh_certificate_authority_id,omitempty"`
	// the list of principals included in the ssh host certificate. This is the list of
	// hostnames and/or IP addresses that are authorized to serve SSH traffic with this
	// certificate. Dangerously, if no principals are specified, this certificate is
	// considered valid for all hosts.
	Principals []string `json:"principals,omitempty"`
	// the time when the ssh host certificate becomes valid, in RFC 3339 format.
	ValidAfter string `json:"valid_after,omitempty"`
	// the time after which the ssh host certificate becomes invalid, in RFC 3339
	// format. the OpenSSH certificates RFC calls this valid_before.
	ValidUntil string `json:"valid_until,omitempty"`
	// the signed SSH certificate in OpenSSH Authorized Keys format. this value should
	// be placed in a -cert.pub certificate file on disk that should be referenced in
	// your sshd_config configuration file with a HostCertificate directive
	Certificate string `json:"certificate,omitempty"`
}

func (x *SSHHostCertificate) String() string {
	return fmt.Sprintf("SSHHostCertificate{ID: %v}", x.ID)

}

func (x *SSHHostCertificate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHHostCertificate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tPublicKey\t%v\n", x.PublicKey)
	fmt.Fprintf(tw, "\tKeyType\t%v\n", x.KeyType)
	fmt.Fprintf(tw, "\tSSHCertificateAuthorityID\t%v\n", x.SSHCertificateAuthorityID)
	fmt.Fprintf(tw, "\tPrincipals\t%v\n", x.Principals)
	fmt.Fprintf(tw, "\tValidAfter\t%v\n", x.ValidAfter)
	fmt.Fprintf(tw, "\tValidUntil\t%v\n", x.ValidUntil)
	fmt.Fprintf(tw, "\tCertificate\t%v\n", x.Certificate)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHHostCertificateList struct {
	// the list of all ssh host certificates on this account
	SSHHostCertificates []SSHHostCertificate `json:"ssh_host_certificates,omitempty"`
	// URI of the ssh host certificates list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *SSHHostCertificateList) String() string {
	return x.GoString()
}

func (x *SSHHostCertificateList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHHostCertificateList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tSSHHostCertificates\t%v\n", x.SSHHostCertificates)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHUserCertificateCreate struct {
	// the ssh certificate authority that is used to sign this ssh user certificate
	SSHCertificateAuthorityID string `json:"ssh_certificate_authority_id,omitempty"`
	// a public key in OpenSSH Authorized Keys format that this certificate signs
	PublicKey string `json:"public_key,omitempty"`
	// the list of principals included in the ssh user certificate. This is the list of
	// usernames that the certificate holder may sign in as on a machine authorizing
	// the signing certificate authority. Dangerously, if no principals are specified,
	// this certificate may be used to log in as any user.
	Principals []string `json:"principals,omitempty"`
	// A map of critical options included in the certificate. Only two critical options
	// are currently defined by OpenSSH: force-command and source-address. See the
	// OpenSSH certificate protocol spec
	// (https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys) for
	// additional details.
	CriticalOptions map[string]string `json:"critical_options,omitempty"`
	// A map of extensions included in the certificate. Extensions are additional
	// metadata that can be interpreted by the SSH server for any purpose. These can be
	// used to permit or deny the ability to open a terminal, do port forwarding, x11
	// forwarding, and more. If unspecified, the certificate will include limited
	// permissions with the following extension map: {"permit-pty": "",
	// "permit-user-rc": ""} OpenSSH understands a number of predefined extensions. See
	// the OpenSSH certificate protocol spec
	// (https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys) for
	// additional details.
	Extensions map[string]string `json:"extensions,omitempty"`
	// The time when the user certificate becomes valid, in RFC 3339 format. Defaults
	// to the current time if unspecified.
	ValidAfter string `json:"valid_after,omitempty"`
	// The time when this host certificate becomes invalid, in RFC 3339 format. If
	// unspecified, a default value of 24 hours will be used. The OpenSSH certificates
	// RFC calls this valid_before.
	ValidUntil string `json:"valid_until,omitempty"`
	// human-readable description of this SSH User Certificate. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH User Certificate.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
}

func (x *SSHUserCertificateCreate) String() string {
	return x.GoString()
}

func (x *SSHUserCertificateCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHUserCertificateCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tSSHCertificateAuthorityID\t%v\n", x.SSHCertificateAuthorityID)
	fmt.Fprintf(tw, "\tPublicKey\t%v\n", x.PublicKey)
	fmt.Fprintf(tw, "\tPrincipals\t%v\n", x.Principals)
	fmt.Fprintf(tw, "\tCriticalOptions\t%v\n", x.CriticalOptions)
	fmt.Fprintf(tw, "\tExtensions\t%v\n", x.Extensions)
	fmt.Fprintf(tw, "\tValidAfter\t%v\n", x.ValidAfter)
	fmt.Fprintf(tw, "\tValidUntil\t%v\n", x.ValidUntil)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHUserCertificateUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this SSH User Certificate. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH User Certificate.
	// optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

func (x *SSHUserCertificateUpdate) String() string {
	return fmt.Sprintf("SSHUserCertificateUpdate{ID: %v}", x.ID)

}

func (x *SSHUserCertificateUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHUserCertificateUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHUserCertificate struct {
	// unique identifier for this SSH User Certificate
	ID string `json:"id,omitempty"`
	// URI of the SSH User Certificate API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the SSH User Certificate API resource was created, RFC 3339
	// format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this SSH User Certificate. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH User Certificate.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// a public key in OpenSSH Authorized Keys format that this certificate signs
	PublicKey string `json:"public_key,omitempty"`
	// the key type of the public_key, one of rsa, ecdsa or ed25519
	KeyType string `json:"key_type,omitempty"`
	// the ssh certificate authority that is used to sign this ssh user certificate
	SSHCertificateAuthorityID string `json:"ssh_certificate_authority_id,omitempty"`
	// the list of principals included in the ssh user certificate. This is the list of
	// usernames that the certificate holder may sign in as on a machine authorizing
	// the signing certificate authority. Dangerously, if no principals are specified,
	// this certificate may be used to log in as any user.
	Principals []string `json:"principals,omitempty"`
	// A map of critical options included in the certificate. Only two critical options
	// are currently defined by OpenSSH: force-command and source-address. See the
	// OpenSSH certificate protocol spec
	// (https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys) for
	// additional details.
	CriticalOptions map[string]string `json:"critical_options,omitempty"`
	// A map of extensions included in the certificate. Extensions are additional
	// metadata that can be interpreted by the SSH server for any purpose. These can be
	// used to permit or deny the ability to open a terminal, do port forwarding, x11
	// forwarding, and more. If unspecified, the certificate will include limited
	// permissions with the following extension map: {"permit-pty": "",
	// "permit-user-rc": ""} OpenSSH understands a number of predefined extensions. See
	// the OpenSSH certificate protocol spec
	// (https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys) for
	// additional details.
	Extensions map[string]string `json:"extensions,omitempty"`
	// the time when the ssh host certificate becomes valid, in RFC 3339 format.
	ValidAfter string `json:"valid_after,omitempty"`
	// the time after which the ssh host certificate becomes invalid, in RFC 3339
	// format. the OpenSSH certificates RFC calls this valid_before.
	ValidUntil string `json:"valid_until,omitempty"`
	// the signed SSH certificate in OpenSSH Authorized Keys Format. this value should
	// be placed in a -cert.pub certificate file on disk that should be referenced in
	// your sshd_config configuration file with a HostCertificate directive
	Certificate string `json:"certificate,omitempty"`
}

func (x *SSHUserCertificate) String() string {
	return fmt.Sprintf("SSHUserCertificate{ID: %v}", x.ID)

}

func (x *SSHUserCertificate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHUserCertificate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tPublicKey\t%v\n", x.PublicKey)
	fmt.Fprintf(tw, "\tKeyType\t%v\n", x.KeyType)
	fmt.Fprintf(tw, "\tSSHCertificateAuthorityID\t%v\n", x.SSHCertificateAuthorityID)
	fmt.Fprintf(tw, "\tPrincipals\t%v\n", x.Principals)
	fmt.Fprintf(tw, "\tCriticalOptions\t%v\n", x.CriticalOptions)
	fmt.Fprintf(tw, "\tExtensions\t%v\n", x.Extensions)
	fmt.Fprintf(tw, "\tValidAfter\t%v\n", x.ValidAfter)
	fmt.Fprintf(tw, "\tValidUntil\t%v\n", x.ValidUntil)
	fmt.Fprintf(tw, "\tCertificate\t%v\n", x.Certificate)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type SSHUserCertificateList struct {
	// the list of all ssh user certificates on this account
	SSHUserCertificates []SSHUserCertificate `json:"ssh_user_certificates,omitempty"`
	// URI of the ssh user certificates list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *SSHUserCertificateList) String() string {
	return x.GoString()
}

func (x *SSHUserCertificateList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SSHUserCertificateList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tSSHUserCertificates\t%v\n", x.SSHUserCertificates)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TLSCertificateCreate struct {
	// human-readable description of this TLS certificate. optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this TLS certificate. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// chain of PEM-encoded certificates, leaf first. See Certificate Bundles
	// (https://ngrok.com/docs/cloud-edge/endpoints#certificate-chains).
	CertificatePEM string `json:"certificate_pem,omitempty"`
	// private key for the TLS certificate, PEM-encoded. See Private Keys
	// (https://ngrok.com/docs/cloud-edge/endpoints#private-keys).
	PrivateKeyPEM string `json:"private_key_pem,omitempty"`
}

func (x *TLSCertificateCreate) String() string {
	return x.GoString()
}

func (x *TLSCertificateCreate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TLSCertificateCreate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCertificatePEM\t%v\n", x.CertificatePEM)
	fmt.Fprintf(tw, "\tPrivateKeyPEM\t%v\n", x.PrivateKeyPEM)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TLSCertificateUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this TLS certificate. optional, max 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this TLS certificate. optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

func (x *TLSCertificateUpdate) String() string {
	return fmt.Sprintf("TLSCertificateUpdate{ID: %v}", x.ID)

}

func (x *TLSCertificateUpdate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TLSCertificateUpdate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TLSCertificate struct {
	// unique identifier for this TLS certificate
	ID string `json:"id,omitempty"`
	// URI of the TLS certificate API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the TLS certificate was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this TLS certificate. optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this TLS certificate. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// chain of PEM-encoded certificates, leaf first. See Certificate Bundles
	// (https://ngrok.com/docs/cloud-edge/endpoints#certificate-chains).
	CertificatePEM string `json:"certificate_pem,omitempty"`
	// subject common name from the leaf of this TLS certificate
	SubjectCommonName string `json:"subject_common_name,omitempty"`
	// subject alternative names (SANs) from the leaf of this TLS certificate
	SubjectAlternativeNames TLSCertificateSANs `json:"subject_alternative_names,omitempty"`
	// timestamp (in RFC 3339 format) when this TLS certificate was issued
	// automatically, or null if this certificate was user-uploaded
	IssuedAt *string `json:"issued_at,omitempty"`
	// timestamp when this TLS certificate becomes valid, RFC 3339 format
	NotBefore string `json:"not_before,omitempty"`
	// timestamp when this TLS certificate becomes invalid, RFC 3339 format
	NotAfter string `json:"not_after,omitempty"`
	// set of actions the private key of this TLS certificate can be used for
	KeyUsages []string `json:"key_usages,omitempty"`
	// extended set of actions the private key of this TLS certificate can be used for
	ExtendedKeyUsages []string `json:"extended_key_usages,omitempty"`
	// type of the private key of this TLS certificate. One of rsa, ecdsa, or ed25519.
	PrivateKeyType string `json:"private_key_type,omitempty"`
	// issuer common name from the leaf of this TLS certificate
	IssuerCommonName string `json:"issuer_common_name,omitempty"`
	// serial number of the leaf of this TLS certificate
	SerialNumber string `json:"serial_number,omitempty"`
	// subject organization from the leaf of this TLS certificate
	SubjectOrganization string `json:"subject_organization,omitempty"`
	// subject organizational unit from the leaf of this TLS certificate
	SubjectOrganizationalUnit string `json:"subject_organizational_unit,omitempty"`
	// subject locality from the leaf of this TLS certificate
	SubjectLocality string `json:"subject_locality,omitempty"`
	// subject province from the leaf of this TLS certificate
	SubjectProvince string `json:"subject_province,omitempty"`
	// subject country from the leaf of this TLS certificate
	SubjectCountry string `json:"subject_country,omitempty"`
}

func (x *TLSCertificate) String() string {
	return fmt.Sprintf("TLSCertificate{ID: %v}", x.ID)

}

func (x *TLSCertificate) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TLSCertificate {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tCreatedAt\t%v\n", x.CreatedAt)
	fmt.Fprintf(tw, "\tDescription\t%v\n", x.Description)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tCertificatePEM\t%v\n", x.CertificatePEM)
	fmt.Fprintf(tw, "\tSubjectCommonName\t%v\n", x.SubjectCommonName)
	fmt.Fprintf(tw, "\tSubjectAlternativeNames\t%v\n", x.SubjectAlternativeNames)
	fmt.Fprintf(tw, "\tIssuedAt\t%v\n", x.IssuedAt)
	fmt.Fprintf(tw, "\tNotBefore\t%v\n", x.NotBefore)
	fmt.Fprintf(tw, "\tNotAfter\t%v\n", x.NotAfter)
	fmt.Fprintf(tw, "\tKeyUsages\t%v\n", x.KeyUsages)
	fmt.Fprintf(tw, "\tExtendedKeyUsages\t%v\n", x.ExtendedKeyUsages)
	fmt.Fprintf(tw, "\tPrivateKeyType\t%v\n", x.PrivateKeyType)
	fmt.Fprintf(tw, "\tIssuerCommonName\t%v\n", x.IssuerCommonName)
	fmt.Fprintf(tw, "\tSerialNumber\t%v\n", x.SerialNumber)
	fmt.Fprintf(tw, "\tSubjectOrganization\t%v\n", x.SubjectOrganization)
	fmt.Fprintf(tw, "\tSubjectOrganizationalUnit\t%v\n", x.SubjectOrganizationalUnit)
	fmt.Fprintf(tw, "\tSubjectLocality\t%v\n", x.SubjectLocality)
	fmt.Fprintf(tw, "\tSubjectProvince\t%v\n", x.SubjectProvince)
	fmt.Fprintf(tw, "\tSubjectCountry\t%v\n", x.SubjectCountry)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TLSCertificateList struct {
	// the list of all TLS certificates on this account
	TLSCertificates []TLSCertificate `json:"tls_certificates,omitempty"`
	// URI of the TLS certificates list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *TLSCertificateList) String() string {
	return x.GoString()
}

func (x *TLSCertificateList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TLSCertificateList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tTLSCertificates\t%v\n", x.TLSCertificates)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TLSCertificateSANs struct {
	// set of additional domains (including wildcards) this TLS certificate is valid
	// for
	DNSNames []string `json:"dns_names,omitempty"`
	// set of IP addresses this TLS certificate is also valid for
	IPs []string `json:"ips,omitempty"`
}

func (x *TLSCertificateSANs) String() string {
	return x.GoString()
}

func (x *TLSCertificateSANs) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TLSCertificateSANs {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tDNSNames\t%v\n", x.DNSNames)
	fmt.Fprintf(tw, "\tIPs\t%v\n", x.IPs)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type Tunnel struct {
	// unique tunnel resource identifier
	ID string `json:"id,omitempty"`
	// URL of the ephemeral tunnel's public endpoint
	PublicURL string `json:"public_url,omitempty"`
	// timestamp when the tunnel was initiated in RFC 3339 format
	StartedAt string `json:"started_at,omitempty"`
	// user-supplied metadata for the tunnel defined in the ngrok configuration file.
	// See the tunnel metadata configuration option
	// (https://ngrok.com/docs/secure-tunnels/ngrok-agent/reference/config#common-tunnel-configuration-properties)
	// In API version 0, this value was instead pulled from the top-level metadata
	// configuration option
	// (https://ngrok.com/docs/secure-tunnels/ngrok-agent/reference/config#metadata).
	Metadata string `json:"metadata,omitempty"`
	// tunnel protocol for ephemeral tunnels. one of http, https, tcp or tls
	Proto string `json:"proto,omitempty"`
	// identifier of tune region where the tunnel is running
	Region string `json:"region,omitempty"`
	// reference object pointing to the tunnel session on which this tunnel was started
	TunnelSession Ref `json:"tunnel_session,omitempty"`
	// the ephemeral endpoint this tunnel is associated with, if this is an
	// agent-initiated tunnel
	Endpoint *Ref `json:"endpoint,omitempty"`
	// the labels the tunnel group backends will match against, if this is a backend
	// tunnel
	Labels map[string]string `json:"labels,omitempty"`
	// tunnel group backends served by this backend tunnel
	Backends []Ref `json:"backends,omitempty"`
	// upstream address the ngrok agent forwards traffic over this tunnel to. this may
	// be expressed as a URL or a network address.
	ForwardsTo string `json:"forwards_to,omitempty"`
}

func (x *Tunnel) String() string {
	return fmt.Sprintf("Tunnel{ID: %v}", x.ID)

}

func (x *Tunnel) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Tunnel {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tID\t%v\n", x.ID)
	fmt.Fprintf(tw, "\tPublicURL\t%v\n", x.PublicURL)
	fmt.Fprintf(tw, "\tStartedAt\t%v\n", x.StartedAt)
	fmt.Fprintf(tw, "\tMetadata\t%v\n", x.Metadata)
	fmt.Fprintf(tw, "\tProto\t%v\n", x.Proto)
	fmt.Fprintf(tw, "\tRegion\t%v\n", x.Region)
	fmt.Fprintf(tw, "\tTunnelSession\t%v\n", x.TunnelSession)
	fmt.Fprintf(tw, "\tEndpoint\t%v\n", x.Endpoint)
	fmt.Fprintf(tw, "\tLabels\t%v\n", x.Labels)
	fmt.Fprintf(tw, "\tBackends\t%v\n", x.Backends)
	fmt.Fprintf(tw, "\tForwardsTo\t%v\n", x.ForwardsTo)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}

type TunnelList struct {
	// the list of all online tunnels on this account
	Tunnels []Tunnel `json:"tunnels,omitempty"`
	// URI of the tunnels list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

func (x *TunnelList) String() string {
	return x.GoString()
}

func (x *TunnelList) GoString() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "TunnelList {\n")
	tw := tabwriter.NewWriter(&b, 0, 4, 0, ' ', 0)
	fmt.Fprintf(tw, "\tTunnels\t%v\n", x.Tunnels)
	fmt.Fprintf(tw, "\tURI\t%v\n", x.URI)
	fmt.Fprintf(tw, "\tNextPageURI\t%v\n", x.NextPageURI)
	tw.Flush()
	fmt.Fprintf(&b, "}\n")
	return b.String()
}
