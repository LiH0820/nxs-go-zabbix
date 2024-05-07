package zabbix

// ItemObject represents a Zabbix Item returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/6.0/manual/api/reference/item/object
type ItemObject struct {
	ItemId          string        `json:"itemid"`
	Type            string        `json:"type"`
	SnmpOid         string        `json:"snmp_oid"`
	HostId          string        `json:"hostid"`
	Name            string        `json:"name"`
	Key             string        `json:"key_"`
	Delay           string        `json:"delay"`
	History         string        `json:"history"`
	Trends          string        `json:"trends"`
	Status          string        `json:"status"`
	ValueType       string        `json:"value_type"`
	TrapperHosts    string        `json:"trapper_hosts"`
	Units           string        `json:"units"`
	Formula         string        `json:"formula"`
	LogTimeFmt      string        `json:"logtimefmt"`
	TemplateId      string        `json:"templateid"`
	ValuemapId      string        `json:"valuemapid"`
	IpmiSensor      string        `json:"ipmi_sensor"`
	AuthType        string        `json:"authtype"`
	Username        string        `json:"username"`
	Password        string        `json:"password"`
	PublicKey       string        `json:"publickey"`
	PrivateKey      string        `json:"privatekey"`
	Flags           string        `json:"flags"`
	InterfaceId     string        `json:"interfaceid"`
	InventoryLink   string        `json:"inventory_link"`
	Lifetime        string        `json:"lifetime"`
	EvalType        string        `json:"evaltype"`
	JmxEndpoint     string        `json:"jmx_endpoint"`
	MasterItemId    string        `json:"master_itemid"`
	Timeout         string        `json:"timeout"`
	Url             string        `json:"url"`
	QueryFields     []interface{} `json:"query_fields"`
	StatusCodes     string        `json:"status_codes"`
	FollowRedirects string        `json:"follow_redirects"`
	PostType        string        `json:"post_type"`
	HttpProxy       string        `json:"http_proxy"`
	RetrieveMode    string        `json:"retrieve_mode"`
	RequestMethod   string        `json:"request_method"`
	OutputFormat    string        `json:"output_format"`
	SslCertFile     string        `json:"ssl_cert_file"`
	SslKeyFile      string        `json:"ssl_key_file"`
	SslKeyPassword  string        `json:"ssl_key_password"`
	VerifyPeer      string        `json:"verify_peer"`
	VerifyHost      string        `json:"verify_host"`
	AllowTraps      string        `json:"allow_traps"`
	Uuid            string        `json:"uuid"`
	State           string        `json:"state"`
	Error           string        `json:"error"`
	Params          string        `json:"params"`
	Description     string        `json:"description"`
	Posts           string        `json:"posts"`
	Headers         []interface{} `json:"headers"`
	Parameters      []interface{} `json:"parameters"`
	LastClock       string        `json:"lastclock"`
	LastNs          string        `json:"lastns"`
	LastValue       string        `json:"lastvalue"`
	PrevValue       string        `json:"prevvalue"`
}

type ItemTagFilter struct {
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Operator int    `json:"operator"`
}

type ItemGetParams struct {
	GetParameters

	// ItemIDs filters search results to items with the given Item ID's.
	ItemIDs []string `json:"itemids,omitempty"`

	// GroupIDs filters search results to items belong to the hosts
	// of the given Group ID's.
	GroupIDs []string `json:"groupids,omitempty"`

	// TemplateIDs filters search results to items belong to the
	// given templates of the given Template ID's.
	TemplateIDs []string `json:"templateids,omitempty"`

	// HostIDs filters search results to items belong to the
	// given Host ID's.
	HostIDs []string `json:"hostids,omitempty"`

	// ProxyIDs filters search results to items that are
	// monitored by the given Proxy ID's.
	ProxyIDs []string `json:"proxyids,omitempty"`

	// InterfaceIDs filters search results to items that use
	// the given host Interface ID's.
	InterfaceIDs []string `json:"interfaceids,omitempty"`

	// GraphIDs filters search results to items that are used
	// in the given graph ID's.
	GraphIDs []string `json:"graphids,omitempty"`

	// TriggerIDs filters search results to items that are used
	// in the given Trigger ID's.
	TriggerIDs []string `json:"triggerids,omitempty"`

	// ApplicationIDs filters search results to items that
	// belong to the given Applications ID's.
	ApplicationIDs []string `json:"applicationids,omitempty"`

	// WebItems flag includes web items in the result.
	WebItems bool `json:"webitems,omitempty"`

	// Inherited flag return only items inherited from a template
	// if set to 'true'.
	Inherited bool `json:"inherited,omitempty"`

	// Templated flag return only items that belong to templates
	// if set to 'true'.
	Templated bool `json:"templated,omitempty"`

	// Monitored flag return only enabled items that belong to
	// monitored hosts if set to 'true'.
	Monitored bool `json:"monitored,omitempty"`

	// Group filters search results to items belong to a group
	// with the given name.
	Group string `json:"group,omitempty"`

	// Host filters search results to items that belong to a host
	// with the given name.
	Host string `json:"host,omitempty"`

	// Application filters search results to items that belong to
	// an application with the given name.
	Application string `json:"application,omitempty"`

	// WithTriggers flag return only items that are used in triggers
	WithTriggers bool `json:"with_triggers,omitempty"`

	// Filter by tags
	Tags []ItemTagFilter `json:"tags,omitempty"`
}

// ItemGet gets items
func (z *Context) ItemGet(params ItemGetParams) ([]ItemObject, error) {
	var result []ItemObject

	err := z.request("item.get", params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
