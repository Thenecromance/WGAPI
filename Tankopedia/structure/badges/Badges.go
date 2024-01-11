package structure
type Badges struct {
BadgeId int32/*numeric*/ `json:"badge_id" xml:"badge_id" ` // Badge ID

Description string/*string*/ `json:"description" xml:"description" ` // Badge description

Name string/*string*/ `json:"name" xml:"name" ` // Badge name

Images *Images`json:"images" xml:"images" ` // Image links
}