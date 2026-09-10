package site

import "resume-site/internal/resume"

// NoIndex controls whether every page ships
// <meta name="robots" content="noindex, nofollow">. Currently true by
// design — the site is reachable by direct link but not meant to be
// search-indexed. Don't flip this without asking the account owner first;
// it changes public search-engine visibility of privacy-redacted personal
// content.
const NoIndex = true

const initials = "HW"

const favicon = `<link rel="icon" href="data:image/svg+xml,` +
	`<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22>` +
	`<text y=%22.9em%22 font-size=%2290%22>%F0%9F%93%84</text></svg>">`

const googleFonts = `<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&display=swap" rel="stylesheet">`

var printLabel = map[string]string{
	"en": "Print / Save PDF",
	"ja": "印刷 / PDF保存",
	"zh": "打印 / 导出PDF",
}

type langTab struct {
	Code, Label, Target string
}

var langTabs = []langTab{
	{"en", "EN", "index.html"},
	{"ja", "日本語", "ja/index.html"},
	{"zh", "中文", "zh/index.html"},
}

// Page is one entry in the site build — mirrors a PAGES dict from the
// Python generator, but with a direct pointer to its data instead of a
// string module name (Go has no importlib.import_module equivalent, and
// doesn't need one: all three data vars are compile-time-known).
type Page struct {
	Data             *resume.ResumeData
	Lang             string
	Output           string // dist-relative, POSIX slashes
	Title            string
	Detailed         bool
	ShowPageSwitch   bool
	PageSwitchActive string // "resume" | "cv" | "" (was None in Python)
}

var Pages = []Page{
	{
		Data: &resume.EN, Lang: "en", Output: "index.html",
		Title: "Hongcheng Wang — Resume", Detailed: false,
		ShowPageSwitch: true, PageSwitchActive: "resume",
	},
	{
		Data: &resume.EN, Lang: "en", Output: "cv.html",
		Title: "Hongcheng Wang — Detailed CV", Detailed: true,
		ShowPageSwitch: true, PageSwitchActive: "cv",
	},
	{
		Data: &resume.JA, Lang: "ja", Output: "ja/index.html",
		Title: "職務経歴書 — 王鴻乗", Detailed: true,
		ShowPageSwitch: false, PageSwitchActive: "",
	},
	{
		Data: &resume.ZH, Lang: "zh", Output: "zh/index.html",
		Title: "个人简历 — 王鸿乘", Detailed: true,
		ShowPageSwitch: false, PageSwitchActive: "",
	},
}
