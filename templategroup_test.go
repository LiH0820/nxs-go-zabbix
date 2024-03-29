package zabbix

import (
	"reflect"
	"testing"
)

const (
	testTemplategroupName = "testTemplategroup"
)

func TestTemplategroupCRUD(t *testing.T) {

	z := GetZabbixContext(t)
	defer DestroyContext(z)

	tgCreatedIDs := testTemplategroupCreate(t, z)
	defer testTemplategroupDelete(t, z, tgCreatedIDs)

	testTemplategroupGet(t, z)
}

func testTemplategroupCreate(t *testing.T, z *Context) []string {

	tgCreatedIDs, err := z.TemplategroupCreate([]TemplategroupObject{
		{
			Name: testTemplategroupName,
		},
	})
	if err != nil {
		t.Fatal("Templategroup create error:", err)
	}

	if len(tgCreatedIDs) == 0 {
		t.Fatal("Templategroup create err: empty IDs array")
	}

	t.Logf("Templategroup create: success")

	return tgCreatedIDs
}

func testTemplategroupDelete(t *testing.T, z *Context, tgCreatedIDs []string) []string {

	tgDeletedIDs, err := z.TemplategroupDelete(tgCreatedIDs)
	if err != nil {
		t.Fatal("Templategroup delete error:", err)
	}

	if len(tgDeletedIDs) == 0 {
		t.Fatal("Templategroup delete error: empty IDs array")
	}

	if reflect.DeepEqual(tgDeletedIDs, tgCreatedIDs) == false {
		t.Fatal("Template delete error: IDs arrays for created and deleted template are mismatch")
	}

	t.Logf("Template delete: success")

	return tgDeletedIDs
}

func testTemplategroupGet(t *testing.T, z *Context) []TemplategroupObject {

	tObjects, err := z.TemplategroupGet(TemplategroupGetParams{
		GetParameters: GetParameters{
			Filter: map[string]interface{}{
				"name": testTemplategroupName,
			},
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Templategroup get error:", err)
	} else {
		if len(tObjects) == 0 {
			t.Error("Templategroup get error: unable to find created template")
		} else {
			t.Logf("Templategroup get: success")
		}
	}

	return tObjects
}
