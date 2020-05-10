package util_test

import (
	"testing"
	"time"

	util "ccs/server/util"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCopy(t *testing.T) {
	Convey("Copier", t, func() {
		type nestedType struct {
			AddressLine1 string
			AddressLine2 string
			City         string
			State        string
		}

		type type1 struct {
			ID                  string
			Name                string
			Avatar              string
			Phone               string
			Zip                 string
			ContactFirstName    string
			ContactLastName     string
			ContactEmailAddress string
			ContactPhone        string
			ContactPhoneExt     string
			Facilities          []nestedType
			CreateAt            time.Time
		}

		type type2 struct {
			ID           string
			Name         string
			Avatar       string
			Phone        string
			Zip          string
			FirstName    string
			LastName     string
			EmailAddress string
			PhoneExt     string
			Facilities   []nestedType
			CreateAt     time.Time
		}

		type1Int := type1{
			Phone: "12121",
			Name:  "Name1",
			Facilities: []nestedType{
				{
					AddressLine1: "AddressLine1",
					AddressLine2: "AddressLine2",
				},
			},
			Zip: "121212",
		}

		type2Int := type2{}

		err := util.Copy(&type2Int, &type1Int)
		So(err, ShouldBeNil)
		So(type2Int.Zip, ShouldEqual, type1Int.Zip)
	})
}
