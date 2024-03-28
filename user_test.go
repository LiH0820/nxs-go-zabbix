package zabbix

import (
	"reflect"
	"testing"
)

const (
	testUserAlias         = "testUserAlias"
	testUserName          = "testUserName"
	testUserSurname       = "testUserSurname"
	testUserPasswd        = "testUserPasswd"
	testUserLang          = "ru_RU"
	testUserMediaEmail    = "test_user@domain.com"
	testUserMediaSeverity = 63
	testUserMediaPeriod   = "1-7,00:00-24:00"
)

func TestUserCRUD(t *testing.T) {

	z := GetZabbixContext(t)
	defer DestroyContext(z)

	// Preparing auxiliary data
	ugCreatedIDs := testUsergroupCreate(t, z)
	defer testUsergroupDelete(t, z, ugCreatedIDs)

	// Create and delete
	uCreatedIDs := testUserCreate(t, z, ugCreatedIDs)
	defer testUserDelete(t, z, uCreatedIDs)

	// Get
	testUserGet(t, z, uCreatedIDs)
}

func testUserCreate(t *testing.T, z *Context, ugCreatedIDs []int) []int {

	var usergroups []UsergroupObject

	// Add usergroups to user
	for _, e := range ugCreatedIDs {
		usergroups = append(usergroups, UsergroupObject{
			UsrgrpID: e,
		})
	}

	uCreatedIDs, err := z.UserCreate([]UserObject{
		{
			Alias:      testUserAlias,
			Name:       testUserName,
			Surname:    testUserSurname,
			Passwd:     testUserPasswd,
			AutoLogin:  UserAutoLoginDisabled,
			AutoLogout: "15m",
			Lang:       testUserLang,
			Type:       UserTypeUser,
			Refresh:    "90s",
			Usrgrps:    usergroups,
			UserMedias: []MediaObject{
				{
					MediaTypeID: 1,
					SendTo:      []string{testUserMediaEmail},
					Active:      MediaActiveEnabled,
					Severity:    testUserMediaSeverity,
					Period:      testUserMediaPeriod,
				},
			},
		},
	})
	if err != nil {
		t.Fatal("Username create error:", err)
	}

	if len(uCreatedIDs) == 0 {
		t.Fatal("Username create error: empty IDs array")
	}

	t.Logf("Username create: success")

	return uCreatedIDs
}

func testUserDelete(t *testing.T, z *Context, uCreatedIDs []int) []int {

	uDeletedIDs, err := z.UserDelete(uCreatedIDs)
	if err != nil {
		t.Fatal("Username delete error:", err)
	}

	if len(uDeletedIDs) == 0 {
		t.Fatal("Username delete error: empty IDs array")
	}

	if reflect.DeepEqual(uDeletedIDs, uCreatedIDs) == false {
		t.Fatal("Username delete error: IDs arrays for created and deleted user are mismatch")
	}

	t.Logf("Username delete: success")

	return uDeletedIDs
}

func testUserGet(t *testing.T, z *Context, uCreatedIDs []int) []UserObject {

	uObjects, err := z.UserGet(UserGetParams{
		UserIDs:          uCreatedIDs,
		SelectMedias:     SelectExtendedOutput,
		SelectMediatypes: SelectExtendedOutput,
		SelectUsrgrps:    SelectExtendedOutput,
		GetParameters: GetParameters{
			Filter: map[string]interface{}{
				"alias":   testUserAlias,
				"name":    testUserName,
				"surname": testUserSurname,
			},
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Username get error:", err)
	} else {
		if len(uObjects) == 0 {
			t.Error("Username get error: unable to find created user")
		} else {
			t.Logf("Username get: success")
		}
	}

	return uObjects
}
