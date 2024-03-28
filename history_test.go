package zabbix

import (
	"testing"
)

const (
	testHistoryItemID = 45503
	testHistoryType   = 0
)

func TestHistoryCRUD(t *testing.T) {

	z := GetZabbixContext(t)
	defer DestroyContext(z)

	// Get
	testHistoryGet(t, z)
}

func testHistoryGet(t *testing.T, z *Context) []HistoryFloatObject {

	r := []HistoryFloatObject{}

	hObjects, err := z.HistoryGet(HistoryGetParams{
		History: HistoryObjectTypeFloat,
		//ItemIDs: []int{testHistoryItemID},
		GetParameters: GetParameters{
			Limit: 1,
		},
	})

	if err != nil {
		t.Error("History get error:", err)
	} else {
		r = *hObjects.(*[]HistoryFloatObject)
		if len(r) == 0 {
			t.Error("History get error: unable to find history")
		} else {
			t.Logf("History get: success")
		}
	}

	return r
}
