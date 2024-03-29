package zabbix

// For `UsermacroObject` field: `Type`
const (
	UsermacroTypeText   = 0
	UsermacroTypeSecret = 1
)

// UsermacroObject struct is used to store hostmacro and globalmacro operations results.
// In API docs Global and Host it is a two different object types that are joined in this package
// into one object `UsermacroObject` that includes fields form both API objects.
// The reason is the some other objects do not separates this types.
//
// see: https://www.zabbix.com/documentation/6.0/manual/api/reference/usermacro/object#host_macro
// and https://www.zabbix.com/documentation/6.0/manual/api/reference/usermacro/object#global_macro
type UsermacroObject struct {

	// Gobal macro fields only
	GlobalmacroID int `json:"globalmacroid,omitempty"`

	// Host macro fields only
	HostmacroID int    `json:"hostmacroid,omitempty"`
	HostID      string `json:"hostid,omitempty"`

	// Common fields
	Macro       string `json:"macro,omitempty"`
	Value       string `json:"value,omitempty"`
	Type        int    `json:"type,omitempty"` // has defined consts, see above
	Description string `json:"description,omitempty"`

	Groups    []HostgroupObject `json:"groups,omitempty"`
	Hosts     []HostObject      `json:"hosts,omitempty"`
	Templates []TemplateObject  `json:"templates,omitempty"`
}

// UsermacroGetParams struct is used for hostmacro get requests
//
// see: https://www.zabbix.com/documentation/6.0/manual/api/reference/usermacro/get#parameters
type UsermacroGetParams struct {
	GetParameters

	Globalmacro    bool     `json:"globalmacro,omitempty"`
	GlobalmacroIDs []string `json:"globalmacroids,omitempty"`
	GroupIDs       []string `json:"groupids,omitempty"`
	HostIDs        []string `json:"hostids,omitempty"`
	HostmacroIDs   []string `json:"hostmacroids,omitempty"`
	TemplateIDs    []string `json:"templateids,omitempty"`

	SelectGroups    SelectQuery `json:"selectGroups,omitempty"`
	SelectHosts     SelectQuery `json:"selectHosts,omitempty"`
	SelectTemplates SelectQuery `json:"selectTemplates,omitempty"`
}

// Structure to store creation result
type hostmacroCreateResult struct {
	HostmacroIDs []string `json:"hostmacroids"`
}

// Structure to store creation global macros result
type globalmacroCreateResult struct {
	GlobalmacroIDs []string `json:"globalmacroids"`
}

// Structure to store deletion result
type hostmacroDeleteResult struct {
	HostmacroIDs []string `json:"hostmacroids"`
}

// Structure to store deletion global macros result
type globalmacroDeleteResult struct {
	GlobalmacroIDs []string `json:"globalmacroids"`
}

// UsermacroGet gets global or host macros according to the given parameters
func (z *Context) UsermacroGet(params UsermacroGetParams) ([]UsermacroObject, error) {

	var result []UsermacroObject

	err := z.request("usermacro.get", params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HostmacroCreate creates new hostmacros
func (z *Context) HostmacroCreate(params []UsermacroObject) ([]string, error) {

	var result hostmacroCreateResult

	err := z.request("usermacro.create", params, &result)
	if err != nil {
		return nil, err
	}

	return result.HostmacroIDs, nil
}

// GlobalmacroCreate creates new globalmacros
func (z *Context) GlobalmacroCreate(params []UsermacroObject) ([]string, error) {

	var result globalmacroCreateResult

	err := z.request("usermacro.createglobal", params, &result)
	if err != nil {
		return nil, err
	}

	return result.GlobalmacroIDs, nil
}

// HostmacroDelete deletes hostmacros
func (z *Context) HostmacroDelete(hostmacroIDs []string) ([]string, error) {

	var result hostmacroDeleteResult

	err := z.request("usermacro.delete", hostmacroIDs, &result)
	if err != nil {
		return nil, err
	}

	return result.HostmacroIDs, nil
}

// GlobalmacroDelete deletes globalmacros
func (z *Context) GlobalmacroDelete(globalmacroIDs []string) ([]string, error) {

	var result globalmacroDeleteResult

	err := z.request("usermacro.deleteglobal", globalmacroIDs, &result)
	if err != nil {
		return nil, err
	}

	return result.GlobalmacroIDs, nil
}
