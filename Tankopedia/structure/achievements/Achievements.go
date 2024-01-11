package structure

type Achievements struct {
	Condition string /*string*/ `json:"condition" xml:"condition" ` // Condition

	Description string /*string*/ `json:"description" xml:"description" ` // Achievement description

	HeroInfo string /*string*/ `json:"hero_info" xml:"hero_info" ` // Historical reference

	Image string /*string*/ `json:"image" xml:"image" ` // URL to image

	ImageBig string /*string*/ `json:"image_big" xml:"image_big" ` // 180x180px image

	Name string /*string*/ `json:"name" xml:"name" ` // Achievement name

	NameI18n string /*string*/ `json:"name_i18n" xml:"name_i18n" ` // Localized

	Order int32 /*numeric*/ `json:"order" xml:"order" ` // Achievement order in this section. Achievements with a lesser value of the Order field are displayed above achievements with a greater value.

	Outdated bool /*boolean*/ `json:"outdated" xml:"outdated" ` // Indicates, if achievement is outdated and cannot be received anymore

	Section string /*string*/ `json:"section" xml:"section" ` // Section

	SectionOrder int32 /*numeric*/ `json:"section_order" xml:"section_order" ` // Section order. Sections with a lesser value of the Section order field are displayed above sections with a greater value.

	Type string /*string*/ `json:"type" xml:"type" ` // Type

	Options []*Options `json:"options" xml:"options" ` // Service Record
}
