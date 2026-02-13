package responseContentDTO

type RatingSystem string

const (
	STAR    RatingSystem = "STAR"
	DIAMOND RatingSystem = "DIAMOND"
)

type (
	// * Indicates a public reward, prize, or title that expresses appreciation for any kind of achievement.
	AwardsResponse struct {
		Name         string       `json:"name"`
		ProviderName string       `json:"providerName"` // example: Michelin Name of the provider who has bestowed the honor or award
		Rating       string       `json:"rating"`       // Describes the rating value and its recognitions for the received honor or award
		Description  string       `json:"description"`  // Describes the rating value and its recognitions for the received honor or award
		DateGranted  string       `json:"dateGranted"`  // date on which the hounor was bestowed upon.Format YYYY-MM-DD (ISO 8601)
		RatingSystem RatingSystem `json:"ratingSystem"` // It is a way to evaluate a restaurant's quality using symbols or other notations. For Instance Star Rating from Michelin Stars or Local Star Rating, Diamond from AAA
	}
)
