package zabbix

import "testing"

func TestEventCRUD(t *testing.T) {

	z := GetZabbixContext(t)
	defer DestroyContext(z)

	// Get
	testEventGet(t, z)
}

func testEventGet(t *testing.T, z *Context) {

	pObjects, err := z.EventGet(EventGetParams{
		Object:      EventObjectTypeTrigger,
		Value:       []EventValue{TriggerEventValueProblem},
		SelectHosts: "extend",
	})

	if err != nil {
		t.Error("Event get error:", err)
	} else {
		if len(pObjects) == 0 {
			t.Error("Event get error: unable to find events")
		} else {
			t.Logf("Event get: success")
		}
	}
}
