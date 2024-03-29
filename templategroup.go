package zabbix

// TemplategroupObject struct is used to store template group operations results
//
// see: https://www.zabbix.com/documentation/7.0/manual/api/reference/template/object
type TemplategroupObject struct {
	GroupID string `json:"groupid,omitempty"`
	Name    string `json:"name,omitempty"`
	Uuid    string `json:"uuid,omitempty"`
}

type TemplategroupGetParams struct {
	GetParameters

	GroupIDs          []string `json:"groupids,omitempty"`
	TemplateIDs       []string `json:"templateids,omitempty"`
	ParentTemplateIDs []string `json:"parentTemplateids,omitempty"`
	HostIDs           []string `json:"hostids,omitempty"`
	GraphIDs          []string `json:"graphids,omitempty"`
	ItemIDs           []string `json:"itemids,omitempty"`
	TriggerIDs        []string `json:"triggerids,omitempty"`

	WithItems          bool `json:"with_items,omitempty"`
	WithItemPrototypes bool `json:"with_item_prototypes,omitempty"`
	WithTemplates      bool `json:"with_templates,omitempty"`
	WithTriggers       bool `json:"with_triggers,omitempty"`
	WithGraphs         bool `json:"with_graphs,omitempty"`
}

// Structure to store creation result
type templategroupCreateResult struct {
	GroupIDs []string `json:"groupids"`
}

// Structure to store deletion result
type templategroupDeleteResult struct {
	GroupIDs []string `json:"groupids"`
}

// TemplategroupGet gets template groups
func (z *Context) TemplategroupGet(params TemplategroupGetParams) (result []TemplategroupObject, err error) {
	err = z.request("templategroup.get", params, &result)
	return
}

// TemplategroupCreate create template groups
func (z *Context) TemplategroupCreate(params []TemplategroupObject) ([]string, error) {

	var result templategroupCreateResult

	err := z.request("templategroup.create", params, &result)
	if err != nil {
		return nil, err
	}

	return result.GroupIDs, nil
}

// TemplategroupDelete delete template groups
func (z *Context) TemplategroupDelete(groupIDs []string) ([]string, error) {

	var result templategroupDeleteResult

	err := z.request("templategroup.delete", groupIDs, &result)
	if err != nil {
		return nil, err
	}

	return result.GroupIDs, nil
}
