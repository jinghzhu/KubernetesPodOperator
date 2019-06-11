package sitespeed

const (
	// CountRule is num of sitespeed rules.
	CountRule = 39
	// Threshold50 is used for analyze the percent of data above scroe 50.
	Threshold50 = 50
	// WorstScore is the bottom score of performance budget.
	WorstScore = 0
	// BestScore is the top score of performance budget.
	BestScore = 90
	// RulesNum is the number of all rules.
	RulesNum = 39
)

// The following 3 strcuts are used to analyze performance budget based on sitespeed.io
// For more details, please read example test.json

type PerformanceBudget struct {
	CookieUrl string          `json:"cookieUrl,omitempty"`
	BasUrl    string          `json:"baseUrl"`
	Timeout   int             `json:"mochaTimeout,omitempty"`
	Id        string          `json:"id,omitempty"`
	Pw        string          `json:"pw,omitempty"`
	TestCases []JsonTestCases `json:"testCases"`
}

type JsonTestCases struct {
	Description string                 `json:"description,omitempty"`
	Pathname    string                 `json:"pathname"`
	UrlParams   map[string]interface{} `json:"urlParams"`
	Budget      JsonBudget             `json:"budget"`
}

type JsonBudget struct {
	Rules JSONRules `json:"rules"`
}

// JSONRules includes all sitespeed.io rules.
type JSONRules struct {
	Criticalpath                int `json:"criticalpath"`
	Spof                        int `json:"spof"`
	Cssnumreq                   int `json:"cssnumreq"`
	Cssimagesnumreq             int `json:"cssimagesnumreq"`
	Jsnumreq                    int `json:"jsnumreq"`
	Yemptysrc                   int `json:"yemptysrc"`
	Ycompress                   int `json:"ycompress"`
	Ycsstop                     int `json:"ycsstop"`
	Yjsbottom                   int `json:"yjsbottom"`
	Yexpressions                int `json:"yexpressions"`
	Ydns                        int `json:"ydns"`
	Yminify                     int `json:"yminify"`
	Redirects                   int `json:"redirects"`
	Noduplicates                int `json:"noduplicates"`
	Yetags                      int `json:"yetags"`
	Yxhr                        int `json:"yxhr"`
	Yxhrmethod                  int `json:"yxhrmethod"`
	Mindom                      int `json:"mindom"`
	Yno404                      int `json:"yno404"`
	Ymincookie                  int `json:"ymincookie"`
	Ycookiefree                 int `json:"ycookiefree"`
	Ynofilter                   int `json:"ynofilter"`
	Avoidscalingimages          int `json:"avoidscalingimages"`
	Yfavicon                    int `json:"yfavicon"`
	Thirdparyasyncjs            int `json:"thirdpartyasyncjs"`
	Csspring                    int `json:"cssprint"`
	Cssinheaddomain             int `json:"cssinheaddomain"`
	Syncjsinhead                int `json:"syncjsinhead"`
	Avoidfont                   int `json:"avoidfont"`
	Totalrequests               int `json:"totalrequests"`
	Expiresmod                  int `json:"expiresmod"`
	Longexpirehead              int `json:"longexpirehead"`
	Nodnslookupswhenfewrequests int `json:"nodnslookupswhenfewrequests"`
	Inlinecsswhenfewrequest     int `json:"inlinecsswhenfewrequest"`
	Textcontent                 int `json:"textcontent"`
	Thirdpartyversions          int `json:"thirdpartyversions"`
	Ycdn                        int `json:"ycdn"`
	Connectionclose             int `json:"connectionclose"`
	Privateheaders              int `json:"privateheaders"`
}
