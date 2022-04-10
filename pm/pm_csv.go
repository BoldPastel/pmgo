//!Inefficent implementation that reads entire file it; this is due to the field permission problem; will correct in future
//*Purpose is to read the CSV and Return Permission Struct

package pmgo

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

func clear(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

var permissionsArray []Permissions

func OpenCSV(filename string) []Permissions { //havent decided what the return is here/
	var header []string
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	f := csv.NewReader(csvFile)

	//define an array of permissions --> this needs to decompose to multiple files.

	var permission Permissions
	permissionsArraytemp := &permissionsArray

	for {
		record, err := f.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if header == nil {
			header = record
			//create the map of permissions
			for _, v := range header[2:] {
				permission.XMLName.Local = "PermissionSet"
				permission.XMLName.Space = "http://soap.sforce.com/2006/04/metadata"
				//permission.UserLicense = "Salesforce"
				permission.Label = v
				permission.HasActivationRequired = "false"
				permissionsArray = append(*permissionsArraytemp, permission)
				//map enum to the number -> so i can find the right column. ; no enum -> String better. array should have a key and value
			}
			continue
			//have to define enum? --> the permission names
		}

		for i, cust := range record[2:] { //over the columns of the csv

			activeperm := &permissionsArray[i]
			switch record[0] {
			case "App":
				{
					if cust == "Yes" {
						var app ApplicationVisibilities
						app.Application = record[1]
						app.Visible = "true"
						activeperm.ApplicationVisibilities = append(activeperm.ApplicationVisibilities, app)
						clear(&app)
					}

				}

			case "Object":
				{
					if cust != "" {
						var obj ObjectPermissions
						obj.AllowCreate = "false"
						obj.AllowDelete = "false"
						obj.AllowEdit = "false"
						obj.AllowRead = "false"
						obj.ModifyAllRecords = "false"
						obj.ViewAllRecords = "false"
						obj.Object = record[1]

						for _, x := range strings.Split(cust, ";") {
							switch x {
							case "R":
								obj.AllowRead = "true"
							case "C":
								obj.AllowCreate = "true"
								obj.AllowRead = "true"
							case "E":
								obj.AllowEdit = "true"
								obj.AllowRead = "true"
							case "D":
								obj.AllowDelete = "true"
								obj.AllowRead = "true"
							}
						}

						activeperm.ObjectPermissions = append(activeperm.ObjectPermissions, obj)
						clear(&obj)
					}

					//might be slightly more tricky and may have to write a func for this.
					// might be easier just matching to value combinations.

				}

			case "Tab":
				{
					if cust != "" {
						var tab TabSettings
						tab.Tab = record[1]
						tab.Visibility = cust

						activeperm.TabSettings = append(activeperm.TabSettings, tab)

						clear(&tab)
					}
				}

			case "RecordType":
				{

					if cust == "Yes" {
						var rt RecordTypeVisibilities
						rt.RecordType = record[1]
						rt.Visible = "true"
						activeperm.RecordTypeVisibilities = append(activeperm.RecordTypeVisibilities, rt)
						clear(&rt)
					}
				}

			case "Field":
				{
					if cust != "" {
						var field FieldPermissions
						field.Field = record[1]

						for _, x := range strings.Split(cust, ";") {
							switch x {
							case "R":
								field.Readable = "true"
								field.Editable = "false"
							case "E":
								field.Editable = "true"
								field.Readable = "true"
							}
						}

						activeperm.FieldPermissions = append(activeperm.FieldPermissions, field)
						clear(&field)
					}
				}

			case "Class":
				{
					if cust == "Yes" {
						var ca ClassAccesses
						ca.ApexClass = record[1]
						ca.Visible = "true"
						activeperm.ClassAccesses = append(activeperm.ClassAccesses, ca)
						clear(&ca)
					}
				}

			case "CustomPerm":
				{
					if cust == "Yes" {
						var ca CustomPermissions
						ca.Name = record[1]
						ca.Enabled = "true"
						activeperm.CustomPermissions = append(activeperm.CustomPermissions, ca)
						clear(&ca)
					}
				}

			case "ExternalDSource":
				{
					if cust == "Yes" {
						var ca ExternalDataSourceAccesses
						ca.ExternalDataSource = record[1]
						ca.Enabled = "true"
						activeperm.ExternalDataSourceAccesses = append(activeperm.ExternalDataSourceAccesses, ca)
						clear(&ca)
					}
				}

			case "Page":
				{
					if cust == "Yes" {
						var ca PageAccesses
						ca.ApexPage = record[1]
						ca.Enabled = "true"
						activeperm.PageAccesses = append(activeperm.PageAccesses, ca)
						clear(&ca)
					}
				}

			case "UserPerm":
				{
					if cust == "Yes" {
						var ca UserPermissions
						ca.Name = record[1]
						ca.Enabled = "true"
						activeperm.UserPermissions = append(activeperm.UserPermissions, ca)
						clear(&ca)
					}
				}

				//clear(&activeperm) //clear out the permission, maybe would have to clear in the inner loop !Would this clear the underlying reference --> maybe should avoid this.
			}
		}

	}
	return *permissionsArraytemp
}

//SCREWS//

/*

func csvUnpack(filename string) chan []string {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	//r := csv.NewReader(csvIn) //* Ideal if headers are identified


		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			// processing here.
		}


	records := make(chan []string)

	go func() {
		p := csv.NewReader(csvFile)

		defer close(records)

		for {
			record, err := p.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			records <- record
		}
	}()

	return records
}

func csvLineConvert(records chan []string) {
	for record := range records {

		switch record[0] {
		case "App":
			//potentially a new function here to map to struct
		}

	}


func appvisconvert(line string)

}
*/
