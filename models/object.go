package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	Objects map[string]*ProcessedObject
)

type Object struct {
	ObjectID          int           `json:"objectID"`
	IsHighlight       bool          `json:"isHighlight"`
	AccessionNumber   string        `json:"accessionNumber"`
	IsPublicDomain    bool          `json:"isPublicDomain"`
	PrimaryImage      string        `json:"primaryImage"`
	PrimaryImageSmall string        `json:"primaryImageSmall"`
	AdditionalImages  []interface{} `json:"additionalImages"`
	Constituents      []struct {
		Role string `json:"role"`
		Name string `json:"name"`
	} `json:"constituents"`
	Department            string    `json:"department"`
	ObjectName            string    `json:"objectName"`
	Title                 string    `json:"title"`
	Culture               string    `json:"culture"`
	Period                string    `json:"period"`
	Dynasty               string    `json:"dynasty"`
	Reign                 string    `json:"reign"`
	Portfolio             string    `json:"portfolio"`
	ArtistRole            string    `json:"artistRole"`
	ArtistPrefix          string    `json:"artistPrefix"`
	ArtistDisplayName     string    `json:"artistDisplayName"`
	ArtistDisplayBio      string    `json:"artistDisplayBio"`
	ArtistSuffix          string    `json:"artistSuffix"`
	ArtistAlphaSort       string    `json:"artistAlphaSort"`
	ArtistNationality     string    `json:"artistNationality"`
	ArtistBeginDate       string    `json:"artistBeginDate"`
	ArtistEndDate         string    `json:"artistEndDate"`
	ObjectDate            string    `json:"objectDate"`
	ObjectBeginDate       int       `json:"objectBeginDate"`
	ObjectEndDate         int       `json:"objectEndDate"`
	Medium                string    `json:"medium"`
	Dimensions            string    `json:"dimensions"`
	CreditLine            string    `json:"creditLine"`
	GeographyType         string    `json:"geographyType"`
	City                  string    `json:"city"`
	State                 string    `json:"state"`
	County                string    `json:"county"`
	Country               string    `json:"country"`
	Region                string    `json:"region"`
	Subregion             string    `json:"subregion"`
	Locale                string    `json:"locale"`
	Locus                 string    `json:"locus"`
	Excavation            string    `json:"excavation"`
	River                 string    `json:"river"`
	Classification        string    `json:"classification"`
	RightsAndReproduction string    `json:"rightsAndReproduction"`
	LinkResource          string    `json:"linkResource"`
	MetadataDate          time.Time `json:"metadataDate"`
	Repository            string    `json:"repository"`
	ObjectURL             string    `json:"objectURL"`
	Tags                  []string  `json:"tags"`
}

type ProcessedObject struct {
	Name       string `json:"name"`
	Title      string `json:"title"`
	Period     string `json:"period"`
	ObjectDate string `json:"objectDate"`
}

func init() {
	Objects = make(map[string]*ProcessedObject)
	Objects["Painting"] = &ProcessedObject{}

}

func GetOne(ObjectId string) (object *ProcessedObject, err error) {

	url := "https://collectionapi.metmuseum.org/public/collection/v1/objects/" + ObjectId

	metMusClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "BackStory")

	res, getErr := metMusClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	metMusResponse := Object{}
	jsonErr := json.Unmarshal(body, &metMusResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return transformer(&metMusResponse), nil

}

func transformer(object *Object) *ProcessedObject {
	transformedObject := ProcessedObject{}
	transformedObject.Name = object.Constituents[0].Name
	transformedObject.ObjectDate = object.ObjectDate
	transformedObject.Period = object.Period
	transformedObject.Title = object.Title
	return &transformedObject
}
