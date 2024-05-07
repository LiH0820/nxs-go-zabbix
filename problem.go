package zabbix

type ProblemSourceType int64

const (
	ProblemSourceTypeTrigger             ProblemSourceType = 0
	ProblemSourceTypeInternal            ProblemSourceType = 3
	ProblemSourceTypeServiceStatusUpdate ProblemSourceType = 4
)

type ProblemObjectType int64

const (
	ProblemObjectTypeTrigger ProblemObjectType = 0
	ProblemObjectTypeItem    ProblemObjectType = 4
	ProblemObjectTypeLLDRule ProblemObjectType = 5
	ProblemObjectTypeService ProblemObjectType = 6
)

type ProblemAcknowledgeType int64

const (
	ProblemAcknowledgeTypeFalse ProblemAcknowledgeType = 0
	ProblemAcknowledgeTypeTrue  ProblemAcknowledgeType = 1
)

type SeverityType int64

const (
	ProblemSeverityTypeNotClassified SeverityType = 0
	ProblemSeverityTypeInformation   SeverityType = 1
	ProblemSeverityTypeWarning       SeverityType = 2
	ProblemSeverityTypeAverage       SeverityType = 3
	ProblemSeverityTypeHigh          SeverityType = 4
	ProblemSeverityTypeDisaster      SeverityType = 5
)

type EvalType int64

const (
	EvalTypeAndOr EvalType = 0
	EvalTypeOR    EvalType = 2
)

// ProblemObject struct is used to store problem operations results
//
// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/problem/object
type ProblemObject struct {
	EventID       string                     `json:"eventid,omitempty"`
	Source        ProblemSourceType          `json:"source,omitempty"`
	Object        ProblemObjectType          `json:"object,omitempty"`
	ObjectID      string                     `json:"objectid,omitempty"`
	Clock         int64                      `json:"clock,omitempty"`
	Ns            int64                      `json:"ns,omitempty"`
	REventID      string                     `json:"r_eventid,omitempty"`
	RClock        int64                      `json:"r_clock,omitempty"`
	RNs           int64                      `json:"r_ns,omitempty"`
	CauseEventID  string                     `json:"cause_eventid,omitempty"`
	CorrelationID string                     `json:"correlationid,omitempty"`
	UserID        string                     `json:"userid,omitempty"`
	Name          string                     `json:"name,omitempty"`
	Acknowledged  ProblemAcknowledgeType     `json:"acknowledged,omitempty"`
	Severity      SeverityType               `json:"severity,omitempty"`
	Suppressed    bool                       `json:"suppressed,omitempty"`
	OpData        string                     `json:"opdata,omitempty"`
	URLs          []MediaTypeURLObject       `json:"urls,omitempty"`
	Acknowledges  []ProblemAcknowledgeObject `json:"acknowledges,omitempty"`
	Tags          []TagObject                `json:"tags,omitempty"`
	Suppression   []ProblemSuppressionObject `json:"suppression_data,omitempty"`
}

type ProblemAcknowledgeObject struct {
	AcknowledgeID string                     `json:"acknowledgeid,omitempty"`
	UserID        string                     `json:"userid,omitempty"`
	EventID       string                     `json:"eventid,omitempty"`
	Clock         int64                      `json:"clock,omitempty"`
	Message       string                     `json:"message,omitempty"`
	Action        EventAcknowledgeActionType `json:"action,omitempty"`
	OldSeverity   SeverityType               `json:"old_severity,omitempty"`
	NewSeverity   SeverityType               `json:"new_severity,omitempty"`
	SuppressUntil int64                      `json:"suppress_until,omitempty"`
	TaskID        string                     `json:"taskid,omitempty"`
}

type MediaTypeURLObject struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type TagObject struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

type ProblemSuppressionObject struct {
	MaintenanceID string `json:"maintenanceid,omitempty"`
	UserID        string `json:"userid,omitempty"`
	SuppressUntil int64  `json:"suppress_until,omitempty"`
}

// ProblemGetParams struct is used for problem get requests
//
// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/problem/get#parameters
type ProblemGetParams struct {
	GetParameters

	EventIDs              []string          `json:"eventids,omitempty"`
	GroupIDs              []string          `json:"groupids,omitempty"`
	HostIDs               []string          `json:"hostids,omitempty"`
	ObjectIDs             []string          `json:"objectids,omitempty"`
	ApplicationIDs        []string          `json:"applicationids,omitempty"`
	Source                ProblemSourceType `json:"source,omitempty"`
	Object                ProblemObjectType `json:"object,omitempty"`
	Acknowledged          bool              `json:"acknowledged,omitempty"`
	Suppressed            bool              `json:"suppressed,omitempty"`
	Severities            []SeverityType    `json:"severities,omitempty"`
	Evaltype              EvalType          `json:"evaltype,omitempty"`
	Tags                  []TagObject       `json:"tags,omitempty"`
	Recent                bool              `json:"recent,omitempty"`
	EventIDFrom           int64             `json:"eventid_from,omitempty"`
	EventIDTill           int64             `json:"eventid_till,omitempty"`
	TimeFrom              int64             `json:"time_from,omitempty"`
	TimeTill              int64             `json:"time_till,omitempty"`
	SelectAcknowledges    SelectQuery       `json:"selectAcknowledges,omitempty"`
	SelectTags            SelectQuery       `json:"selectTags,omitempty"`
	SelectSuppressionData SelectQuery       `json:"selectSuppressionData,omitempty"`
	SortField             []string          `json:"sortfield,omitempty"`
}

// ProblemGet gets problems
func (z *Context) ProblemGet(params ProblemGetParams) ([]ProblemObject, error) {

	var result []ProblemObject

	err := z.request("problem.get", params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
