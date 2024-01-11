package structure

type Crewroles struct {
	Name string/*string*/ `json:"name" xml:"name" ` // Сrew qualification name

	Role string/*string*/ `json:"role" xml:"role" ` // Сrew qualification ID

	Skills []string/*list of strings*/ `json:"skills" xml:"skills" ` // List of crew member qualifications

}
