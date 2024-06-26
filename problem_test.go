package zabbix

import (
	"testing"
)

func TestProblemCRUD(t *testing.T) {

	z := GetZabbixContext(t)
	defer DestroyContext(z)

	// Get
	testProblemGet(t, z)
}

func testProblemGet(t *testing.T, z *Context) {

	pObjects, err := z.ProblemGet(ProblemGetParams{
		//ObjectIDs: []int{20143},
		// ... Add other fields as needed
	})

	if err != nil {
		t.Error("Problem get error:", err)
	} else {
		if len(pObjects) == 0 {
			t.Error("Problem get error: unable to find problems")
		} else {
			t.Logf("Problem get: success")
		}
	}
}
