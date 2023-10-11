package portfolioapp

type Image struct {
	Url     string  `json:"url"`
	Layout  *string `json:"layout"`
	Caption *string `json:"caption"`
}

type SubSection struct {
	Type    string   `json:"type"`
	Heading string   `json:"heading"`
	Content string   `json:"content"`
	Images  *[]Image `json:"images"`
}

type Section struct {
	Heading     *string      `json:"heading"`
	SubSections []SubSection `json:"subSections"`
}

type Label struct {
	Name string `json:"name"`
}

type Meta struct {
	Client   string `json:"client"`
	Role     string `json:"role"`
	Duration string `json:"duration"`
	Location string `json:"location"`
}

type CaseStudy struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Lead        string    `json:"lead"`
	HeroImgUrl  string    `json:"heroImgUrl"`
	Labels      []string  `json:"labels"`
	Meta        Meta      `json:"meta"`
	Sections    []Section `json:"sections"`
}
