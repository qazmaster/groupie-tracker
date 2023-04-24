package artist

type Artist struct {
	id           int       `json:"id"`
	imageURL     string    `json:"image"`
	name         string    `json:"name"`
	members      []string  `json:"[]string"`
	creationDate int       `json:"creationDate"`
	firstAlbum   string    `json:"firstAlbum"`
	locations    Locations `json:"locations"`
	concertDates Dates     `json:"concertDates""`
	relation     Relation  `json:"relations"`
}

type Locations struct {
	id        int
	locations []string
}

type Dates struct {
	id    int
	dates string
}

type Relation struct {
	id             int
	datesLocations map[string][]string
}
