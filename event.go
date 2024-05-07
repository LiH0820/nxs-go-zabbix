package zabbix

type EventAcknowledgeActionType int64

const (
	EventAcknowledgeActionTypeClose                  EventAcknowledgeActionType = 1
	EventAcknowledgeActionTypeAck                    EventAcknowledgeActionType = 2
	EventAcknowledgeActionTypeAddMessage             EventAcknowledgeActionType = 4
	EventAcknowledgeActionTypeChangeSeverity         EventAcknowledgeActionType = 8
	EventAcknowledgeActionTypeUnack                  EventAcknowledgeActionType = 16
	EventAcknowledgeActionTypeSuppress               EventAcknowledgeActionType = 32
	EventAcknowledgeActionTypeUnsuppress             EventAcknowledgeActionType = 64
	EventAcknowledgeActionTypeChangeEventRankCause   EventAcknowledgeActionType = 128
	EventAcknowledgeActionTypeChangeEventRankSymptom EventAcknowledgeActionType = 256
)

type EventSourceType int64

const (
	EventSourceTrigger EventSourceType = iota
	EventSourceDiscoveryRule
	EventSourceAutoRegistration
	EventSourceInternal
)

type EventObjectType int64

const (
	EventObjectTypeTrigger EventObjectType = iota
	EventObjectTypeDiscoveredHost
	EventObjectTypeDiscoveredService
	EventObjectTypeAutoRegisteredHost
	EventObjectTypeItem
	EventObjectTypeLLDRule
)

type EventValue int64

const (
	// TriggerEventValueOK indicates that the Object related to an EventObject with
	// Source type EventSourceTrigger is in an "OK" state.
	TriggerEventValueOK EventValue = iota

	// TriggerEventValueProblem indicates that the Object related to an EventObject with
	// Source type EventSourceTrigger is in a "Problem" state.
	TriggerEventValueProblem
)

const (
	// DiscoveryEventValueUp indicates that the Host or Service related to an
	// EventObject with Source type EventSourceDiscoveryRule is in an "Up" state.
	DiscoveryEventValueUp EventValue = iota

	// DiscoveryEventValueDown indicates that the Host or Service related to an
	// EventObject with Source type EventSourceDiscoveryRule is in a "Down" state.
	DiscoveryEventValueDown

	// DiscoveryEventValueDiscovered indicates that the Host or Service related
	// to an EventObject with Source type EventSourceDiscoveryRule is in a
	// "Discovered" state.
	DiscoveryEventValueDiscovered

	// DiscoveryEventValueLost indicates that the Host or Service related to an
	// EventObject with Source type EventSourceDiscoveryRule is in a "Lost" state.
	DiscoveryEventValueLost
)

const (
	// InternalEventValueNormal indicates that the Object related to an EventObject
	// with Source type EventSourceInternal is in a "Normal" state.
	InternalEventValueNormal EventValue = iota

	// InternalEventValueNotSupported indicates that the Object related to an
	// EventObject with Source type EventSourceInternal is in an "Unknown" or
	// "Not supported" state.
	InternalEventValueNotSupported
)

// EventObject represents a Zabbix EventObject returned from the Zabbix API. Events are
// readonly as they may only be created by the Zabbix server.
//
// See: https://www.zabbix.com/documentation/2.2/manual/config/events
type EventObject struct {
	EventId       string               `json:"eventId,omitempty"`
	Source        EventSourceType      `json:"source,omitempty"`
	Object        EventObjectType      `json:"object,omitempty"`
	ObjectId      string               `json:"objectid,omitempty"`
	Acknowledged  bool                 `json:"acknowledged,omitempty"`
	Clock         int64                `json:"clock,omitempty"`
	Ns            int64                `json:"ns,omitempty"`
	Name          string               `json:"name,omitempty"`
	Value         EventValue           `json:"value,omitempty"`
	Severity      SeverityType         `json:"severity,omitempty"`
	REventId      string               `json:"r_eventid,omitempty"`
	CEventId      string               `json:"c_eventid,omitempty"`
	CorrelationId string               `json:"correlationid,omitempty"`
	UserId        string               `json:"userid,omitempty"`
	Suppressed    bool                 `json:"suppressed,omitempty"`
	OpData        string               `json:"opdata,omitempty"`
	Urls          []MediaTypeURLObject `json:"urls,omitempty"`
	Hosts         []HostObject         `json:"hosts,omitempty"`
}

// EventGetParams is query params for event.get call
type EventGetParams struct {
	GetParameters

	EventIDs               []string        `json:"eventids,omitempty"`
	GroupIDs               []string        `json:"groupids,omitempty"`
	HostIDs                []string        `json:"hostids,omitempty"`
	ObjectIDs              []string        `json:"objectids,omitempty"`
	Object                 EventObjectType `json:"object,omitempty"`
	AcknowledgedOnly       bool            `json:"acknowledged,omitempty"`
	Suppressed             bool            `json:"suppressed,omitempty"`
	Severity               SeverityType    `json:"severity,omitempty"`
	EvalType               EvalType        `json:"evaltype,omitempty"`
	Tags                   []TagObject     `json:"tags,omitempty"`
	MinEventID             string          `json:"eventid_from,omitempty"`
	MaxEventID             string          `json:"eventid_till,omitempty"`
	MinTime                int64           `json:"time_from,omitempty"`
	MaxTime                int64           `json:"time_till,omitempty"`
	ProblemTimeFrom        int64           `json:"problem_time_from,omitempty"`
	ProblemTimeTill        int64           `json:"problem_time_till,omitempty"`
	Value                  []EventValue    `json:"value,omitempty"`
	SelectHosts            SelectQuery     `json:"selectHosts,omitempty"`
	SelectRelatedObject    SelectQuery     `json:"selectRelatedObject,omitempty"`
	SelectAlerts           SelectQuery     `json:"select_alerts,omitempty"`
	SelectAcknowledgements SelectQuery     `json:"select_acknowledges,omitempty"`
	SelectTags             SelectQuery     `json:"selectTags,omitempty"`
}

// EventGet get events
func (z *Context) EventGet(params EventGetParams) ([]EventObject, error) {

	var result []EventObject

	err := z.request("event.get", params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
