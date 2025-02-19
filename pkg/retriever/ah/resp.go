package ah

type bonusResT struct {
	SectionType          string `json:"sectionType"`
	SectionDescription   string `json:"sectionDescription"`
	BonusGroupOrProducts []struct {
        BonusGroup bonusGroupT `json:"bonusGroup,omitempty"`
	} `json:"bonusGroupOrProducts"`
	SectionImage []any `json:"sectionImage"`
}


type bonusGroupT struct {
    ID                 string `json:"id"`
    OfferID            int    `json:"offerId"`
    OfferStartDate     string `json:"offerStartDate"`
    SegmentID          int    `json:"segmentId"`
    SegmentDescription string `json:"segmentDescription"`
    BonusStartDate     string `json:"bonusStartDate"`
    BonusEndDate       string `json:"bonusEndDate"`
    ValidityPeriod     struct {
        Start       string `json:"start"`
        End         string `json:"end"`
        Description string `json:"description"`
        Code        string `json:"code"`
    } `json:"validityPeriod"`
    BonusType           string `json:"bonusType"`
    PromotionType       string `json:"promotionType"`
    SegmentType         string `json:"segmentType"`
    DiscountDescription string `json:"discountDescription"`
    Images              []struct {
        Width  int    `json:"width"`
        Height int    `json:"height"`
        URL    string `json:"url"`
    } `json:"images"`
    Category               string  `json:"category"`
    Future                 bool    `json:"future"`
    Products               []any   `json:"products"`
    ShopType               string  `json:"shopType"`
    BonusPeriodDescription string  `json:"bonusPeriodDescription"`
    ExampleFromPrice       float64 `json:"exampleFromPrice"`
    ExampleForPrice        float64 `json:"exampleForPrice"`
    ExampleHasListPrice    bool    `json:"exampleHasListPrice"`
    IsStapelBonus          bool    `json:"isStapelBonus"`
    ExtraDescriptions      []any   `json:"extraDescriptions"`
    DiscountLabels         []struct {
        Code               string  `json:"code"`
        DefaultDescription string  `json:"defaultDescription"`
        Count              int     `json:"count"`
        Price              float64 `json:"price"`
    } `json:"discountLabels"`
    StoreOnlyPromotion bool `json:"storeOnlyPromotion"`
} 
