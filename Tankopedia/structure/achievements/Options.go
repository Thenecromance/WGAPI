package structure

type Options struct {
	Description string /*string*/ `json:"description" xml:"description" ` // Achievement description

	Image string /*string*/ `json:"image" xml:"image" ` // URL to image

	ImageBig string /*string*/ `json:"image_big" xml:"image_big" ` // 180x180px image

	NameI18n string /*string*/ `json:"name_i18n" xml:"name_i18n" ` // Localized

	NationImages *NationImages `json:"nation_images" xml:"nation_images" ` // Information about nation emblems
}
