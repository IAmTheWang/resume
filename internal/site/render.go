package site

import (
	"embed"
	"fmt"
	"html"
	"html/template"
	"strings"

	"resume-site/internal/resume"
)

//go:embed templates/*.tmpl.html
var templateFS embed.FS

var tmpl = template.Must(template.New("root").Funcs(template.FuncMap{
	// Raw HTML fragments the data never controls — safe to mark trusted,
	// matching the two spots the Python generator interpolates unescaped.
	"dot": func() template.HTML { return template.HTML("&nbsp;·&nbsp;") },
	"br":  func() template.HTML { return template.HTML("<br>") },
	// esc mirrors Python's esc() (a plain html.escape() wrapper) exactly,
	// via Go's stdlib "html" package rather than html/template's own
	// context-derived escaper — html/template additionally escapes ASCII
	// '+' to "&#43;" in body text, which Python's html.escape does not do,
	// and this resume's content is full of literal '+' (tech-stack notation
	// like "Vue 3 + TypeScript", "15+ suites"). Escaping every resume-data
	// field through this func instead of relying on bare {{.}} keeps output
	// byte-identical to the Python generator (the apostrophe encoding,
	// &#39; vs &#x27;, is the one accepted difference — see internal/site/CLAUDE.md)
	// while still escaping everything that actually matters for HTML safety
	// (& < > ' ").
	"esc": func(s string) template.HTML { return template.HTML(html.EscapeString(s)) },
}).ParseFS(templateFS, "templates/*.tmpl.html"))

// NavLink is one rendered nav item (language tab or Resume/CV switch link).
type NavLink struct {
	Href   string
	Label  string
	Active bool
}

// PageView is what the templates render from — every field is precomputed
// in Go so templates contain no business logic, only presentation.
type PageView struct {
	Data          *resume.ResumeData
	Page          *Page
	CSSHref       string
	NavLinks      []NavLink
	SwitchLinks   []NavLink
	PrintLabel    string
	Robots        template.HTML
	Favicon       template.HTML
	GoogleFonts   template.HTML
	Initials      string
	OGDescription string
}

type experienceItemView struct {
	Job                     resume.Job
	ShowResponsibilities    bool
	ShowAchievementsSubhead bool
}

func ogDescription(d *resume.ResumeData, detailed bool) string {
	var text string
	if detailed || !d.HasSummary {
		text = d.SummaryLong[0]
	} else {
		text = d.Summary
	}
	return truncate(text, 200)
}

func buildNavLinks(page *Page) []NavLink {
	links := make([]NavLink, 0, len(langTabs))
	for _, lt := range langTabs {
		links = append(links, NavLink{
			Href:   relhref(lt.Target, page.Output),
			Label:  lt.Label,
			Active: lt.Code == page.Lang,
		})
	}
	return links
}

func buildSwitchLinks(page *Page) []NavLink {
	if !page.ShowPageSwitch {
		return nil
	}
	return []NavLink{
		{Href: relhref("index.html", page.Output), Label: "Resume", Active: page.PageSwitchActive == "resume"},
		{Href: relhref("cv.html", page.Output), Label: "Detailed CV", Active: page.PageSwitchActive == "cv"},
	}
}

func execTemplate(name string, data any) (string, error) {
	var buf strings.Builder
	if err := tmpl.ExecuteTemplate(&buf, name, data); err != nil {
		return "", fmt.Errorf("execute template %q: %w", name, err)
	}
	return buf.String(), nil
}

// Render builds the full HTML document for a page, mirroring build_page()
// in the Python generator. The outer document/head/body skeleton is plain
// Go string assembly (it carries no resume content, so there's nothing to
// auto-escape); every section holding resume data goes through html/template
// so repeated, arbitrary strings from internal/resume get consistent,
// automatic escaping.
func Render(page *Page, cssHash string) (string, error) {
	robots := template.HTML("")
	if NoIndex {
		robots = template.HTML(`<meta name="robots" content="noindex, nofollow">` + "\n")
	}

	view := PageView{
		Data:          page.Data,
		Page:          page,
		CSSHref:       relhref("style.css", page.Output) + "?v=" + cssHash,
		NavLinks:      buildNavLinks(page),
		SwitchLinks:   buildSwitchLinks(page),
		PrintLabel:    printLabel[page.Lang],
		Robots:        robots,
		Favicon:       template.HTML(favicon),
		GoogleFonts:   template.HTML(googleFonts),
		Initials:      initials,
		OGDescription: ogDescription(page.Data, page.Detailed),
	}

	headHTML, err := execTemplate("head", view)
	if err != nil {
		return "", err
	}
	heroHTML, err := execTemplate("hero", view)
	if err != nil {
		return "", err
	}
	summaryHTML, err := execTemplate("summary", view)
	if err != nil {
		return "", err
	}
	skillsHTML, err := execTemplate("skills", view)
	if err != nil {
		return "", err
	}
	educationLanguagesHTML, err := execTemplate("educationLanguages", view)
	if err != nil {
		return "", err
	}

	itemStrings := make([]string, len(page.Data.Experience))
	for i, job := range page.Data.Experience {
		itemView := experienceItemView{
			Job:                     job,
			ShowResponsibilities:    page.Detailed && len(job.Responsibilities) > 0,
			ShowAchievementsSubhead: page.Detailed,
		}
		itemHTML, err := execTemplate("experienceItem", itemView)
		if err != nil {
			return "", err
		}
		itemStrings[i] = itemHTML
	}
	experienceHTML := fmt.Sprintf(
		"<section>\n  <h2>Experience</h2>\n  <div class=\"timeline\">\n%s\n  </div>\n</section>",
		strings.Join(itemStrings, "\n"),
	)

	sections := []string{heroHTML, summaryHTML}
	if page.Detailed {
		coreStrengthsHTML, err := execTemplate("coreStrengths", view)
		if err != nil {
			return "", err
		}
		sections = append(sections, coreStrengthsHTML)
	}
	sections = append(sections, skillsHTML, experienceHTML, educationLanguagesHTML)
	if page.Detailed {
		personalStatementHTML, err := execTemplate("personalStatement", view)
		if err != nil {
			return "", err
		}
		sections = append(sections, personalStatementHTML)
	}
	body := strings.Join(sections, "\n")

	full := fmt.Sprintf(
		"<!doctype html>\n<html lang=\"%s\">\n%s\n<body>\n<div class=\"page\">\n%s\n</div>\n</body>\n</html>\n",
		page.Lang, headHTML, body,
	)
	return full, nil
}
