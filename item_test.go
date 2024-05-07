package zabbix

import (
	"testing"
)

func TestItemCRUD(t *testing.T) {

	z := GetZabbixContext(t)
	defer DestroyContext(z)

	// Get
	testItemGet(t, z)
}

func testItemGet(t *testing.T, z *Context) []ItemObject {

	r := []ItemObject{}

	ItemObjects, err := z.ItemGet(ItemGetParams{
		Monitored: true,
		Tags: []ItemTagFilter{
			ItemTagFilter{
				Tag:   "compose",
				Value: "calls",
			},
		},
		GetParameters: GetParameters{
			//Limit: 1,
		},
	})

	if err != nil {
		t.Error("Item get error:", err)
	} else {
		r = ItemObjects
		if len(r) == 0 {
			t.Error("Item get error: unable to find history")
		} else {
			t.Logf("Item get: success")
		}
	}

	return r
}
