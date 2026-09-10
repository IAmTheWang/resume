// Package resume holds the per-language resume content — the single source
// of truth for what's on the site, one file per language.
package resume

type SkillEntry struct {
	Label string
	Value string
}

type LangEntry struct {
	Label string
	Value string
}

type Job struct {
	Company          string
	Role             string
	Period           string
	Location         string
	Meta             string // "" means absent (job.get("meta") was falsy in the Python source)
	Blurb            string
	Responsibilities []string // only rendered when the page is Detailed ("Key Contributions")
	Achievements     []string // always rendered
}

type Education struct {
	Degree string
	School string
}

type ResumeData struct {
	Name          string
	Tagline       string
	Location      string
	Email         string
	LinkedInLabel string
	LinkedInURL   string

	Summary    string // short-form summary; EN only
	HasSummary bool   // true only for EN — JA/ZH have no short form at all

	SummaryLong       []string
	CoreStrengths     []string
	Skills            []SkillEntry
	Experience        []Job // newest first
	Education         Education
	Languages         []LangEntry // length varies by language: EN/JA=4, ZH=2
	PersonalStatement string
}
