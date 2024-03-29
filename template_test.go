package zabbix

import (
	"reflect"
	"testing"
)

const (
	testTemplateName = "testTemplate"
)

func TestTemplateCRUD(t *testing.T) {

	z := GetZabbixContext(t)
	defer DestroyContext(z)

	// Preparing auxiliary data
	hgCreatedIDs := testHostgroupCreate(t, z)
	defer testHostgroupDelete(t, z, hgCreatedIDs)

	// Create and delete
	tCreatedIDs := testTemplateCreate(t, z, hgCreatedIDs)
	defer testTemplateDelete(t, z, tCreatedIDs)

	// Get
	testTemplateGet(t, z, tCreatedIDs, hgCreatedIDs)
}

func testTemplateCreate(t *testing.T, z *Context, tgCreatedIDs []string) []string {

	var groups []TemplategroupObject

	// Add groups to template
	for _, e := range tgCreatedIDs {
		groups = append(groups, TemplategroupObject{
			GroupID: e,
		})
	}

	tCreatedIDs, err := z.TemplateCreate([]TemplateObject{
		{
			Host:   testTemplateName,
			Groups: groups,
		},
	})
	if err != nil {
		t.Fatal("Template create error:", err)
	}

	if len(tCreatedIDs) == 0 {
		t.Fatal("Template create error: empty IDs array")
	}

	t.Logf("Template create: success")

	return tCreatedIDs
}

func testTemplateDelete(t *testing.T, z *Context, tCreatedIDs []string) []string {

	tDeletedIDs, err := z.TemplateDelete(tCreatedIDs)
	if err != nil {
		t.Fatal("Template delete error:", err)
	}

	if len(tDeletedIDs) == 0 {
		t.Fatal("Template delete error: empty IDs array")
	}

	if reflect.DeepEqual(tDeletedIDs, tCreatedIDs) == false {
		t.Fatal("Template delete error: IDs arrays for created and deleted template are mismatch")
	}

	t.Logf("Template delete: success")

	return tDeletedIDs
}

func testTemplateGet(t *testing.T, z *Context, tCreatedIDs, hgCreatedIDs []string) []TemplateObject {

	tObjects, err := z.TemplateGet(TemplateGetParams{
		TemplateIDs: tCreatedIDs,
		GroupIDs:    hgCreatedIDs,
		GetParameters: GetParameters{
			Filter: map[string]interface{}{
				"name": testTemplateName,
			},
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Template get error:", err)
	} else {
		if len(tObjects) == 0 {
			t.Error("Template get error: unable to find created template")
		} else {
			t.Logf("Template get: success")
		}
	}

	return tObjects
}
