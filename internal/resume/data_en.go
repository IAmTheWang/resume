package resume

// EN is the English resume content — the single source of truth for both
// index.html (short Resume) and cv.html (Detailed CV).
var EN = ResumeData{
	Name:          "Hongcheng Wang",
	Tagline:       "Frontend & BFF Engineer  ·  Bilingual EN / JA (JLPT N1)",
	Location:      "Chiba, Japan",
	Email:         "wanghongcheng10@gmail.com",
	LinkedInLabel: "linkedin.com/in/hongcheng-wang-1597a6381",
	LinkedInURL:   "https://linkedin.com/in/hongcheng-wang-1597a6381",

	HasSummary: true,
	Summary: "Frontend engineer with 7 years of experience building production web applications in " +
		"China and Japan, currently shipping a clinician-facing health-record platform in Tokyo. " +
		"Strongest in TypeScript across Vue 3 and React, with BFF-layer experience in Next.js and " +
		"Nuxt — service aggregation, streaming proxying, and auth middleware — plus hands-on " +
		"ownership of AWS delivery pipelines built on S3, CloudFront, Terraform, and GitHub Actions " +
		"OIDC. Work day-to-day in Japanese with clinical and product stakeholders while authoring " +
		"production architecture documents in both English and Japanese. AI-assisted development " +
		"practice since 2023: Cursor for production code in China, Claude and ChatGPT since " +
		"relocating to Japan for code authoring, debugging, test generation, documentation, and " +
		"code review.",

	SummaryLong: []string{
		"Frontend engineer with 7 years of experience delivering production web applications " +
			"across healthcare, energy, manufacturing, and geospatial domains in China and Japan.",
		"Deep front-end delivery in TypeScript (Vue 3, React, Next.js, Nuxt), including BFF-layer " +
			"engineering in Next.js/Nuxt server routes — service aggregation across microservices, " +
			"streaming proxying, auth middleware, and server-side data shaping — paired with hands-on " +
			"AWS delivery-pipeline ownership built on S3, CloudFront, Terraform, and GitHub Actions " +
			"OIDC.",
		"Works onsite with Japanese stakeholders in Japanese while authoring production " +
			"architecture documents in both English and Japanese; operates in both Agile and Waterfall " +
			"lifecycles, from backlog grooming through production release.",
		"AI-assisted development practice since 2023: Cursor for day-to-day feature work at " +
			"INTSIG in China; Claude and ChatGPT after relocating to Japan, used for code authoring " +
			"and debugging, test generation, drafting bilingual architecture documentation, and code " +
			"review.",
	},

	CoreStrengths: []string{
		"7 years shipping production web applications end to end, from backlog grooming through " +
			"release.",
		"Deep hands-on expertise in TypeScript — Vue 3 (Composition API, Pinia, Vue Router) and " +
			"React — building production SPAs end to end.",
		"BFF-layer engineering within Next.js and Nuxt server routes — multi-service aggregation, " +
			"streaming proxying, auth middleware, and server-side data shaping — extending front-end " +
			"ownership up to the edge of backend services.",
		"Domain exposure spanning healthcare, energy, manufacturing, and geospatial products.",
		"Infrastructure-as-code delivery on AWS: Terraform, CloudFormation, and GitHub Actions " +
			"pipelines authenticated with OIDC.",
		"AI-assisted development since 2023 — Cursor, then Claude and ChatGPT after relocating " +
			"to Japan — integrated into daily coding, debugging, testing, documentation, and code " +
			"review.",
		"Project leadership for teams of up to four engineers and cross-functional partners.",
	},

	Skills: []SkillEntry{
		{"Languages", "TypeScript, JavaScript, Go, Shell"},
		{"Frontend", "Vue 3 (Composition API, Pinia, Vue Router, vue-i18n), React, Next.js, Nuxt, " +
			"Angular, Chart.js, Vite, Sass/SCSS"},
		{"BFF / Server-Side", "Next.js (API Routes, App Router, Middleware), Nuxt 3 (Nitro Server Routes, " +
			"Server Middleware), SSE streaming/proxying, service aggregation, server-side " +
			"caching, data masking/DTO shaping"},
		{"Cloud & DevOps", "AWS (S3, CloudFront, WAF, Route 53, IAM/OIDC), Terraform, CloudFormation, " +
			"GitHub Actions, CodeDeploy, CI/CD pipeline design"},
		{"Testing & Quality", "Vitest, Vue Test Utils, ESLint, Prettier, Husky"},
		{"Tools & Process", "Git / GitHub, Jira, Confluence, Gitea, Agile / Scrum"},
		{"AI-Assisted Development", "Cursor (2023, production code in China); Claude, ChatGPT (2024–present, Japan) for " +
			"code authoring, debugging, test generation, documentation, and code review"},
	},

	Experience: []Job{
		{
			Company:  "Welby, Inc.",
			Role:     "Frontend & BFF Engineer, Service Development Dept.",
			Period:   "Mar 2026 – Present",
			Location: "Tokyo, Japan",
			Meta:     "Working language: Japanese",
			Blurb: "Clinician-facing PHR (Personal Health Record) platform for chronic-disease care, " +
				"built as a Vue 3 + TypeScript SPA. Tokyo Stock Exchange Growth-listed company.",
			Responsibilities: []string{
				"Build and maintain features across the clinician-facing SPA using Vue 3, " +
					"TypeScript, Pinia, and Chart.js.",
				"Design delivery architecture for the mykarte SPA on AWS and take it through " +
					"team review before implementation.",
				"Write Vitest and Vue Test Utils coverage alongside every feature rather than " +
					"afterwards.",
				"Implemented server-side data masking and DTO trimming in Next.js for medical " +
					"data returned by backend microservices, passing through only the fields the UI " +
					"needs to render.",
				"Added BFF-layer caching in Next.js for infrequently-changing patient " +
					"health-chart and dashboard data to reduce load on downstream microservices.",
				"Implemented timeout control and fallback default values in Next.js API Routes " +
					"for slow or intermittently-timing-out third-party medical services.",
			},
			Achievements: []string{
				"Hardened the Next.js BFF layer for the clinician-facing platform: server-side " +
					"masking of sensitive medical fields before they reach the client, BFF-layer " +
					"caching for infrequently-changing health-chart data, and timeout/fallback " +
					"handling for slow third-party medical services.",
				"Built the LTEP (Long-Term eGFR Progression) chart for chronic kidney disease " +
					"patients — an interactive Chart.js visualization with draggable intervention " +
					"points, a configurable baseline, and a dedicated print layout — shipped across " +
					"12 feature branches in 4 months.",
				"Authored the S3 + CloudFront + GitHub Actions delivery architecture for the " +
					"mykarte SPA: replaced long-lived AWS keys with OIDC short-lived credentials, " +
					"applied a CloudFront Response Headers Policy (HSTS, CSP, X-Frame-Options) that " +
					"also covers SPA fallback responses, deployed WAFv2 in COUNT mode, and gated " +
					"production behind a typed-confirmation workflow.",
				"Reused the existing Terraform module set (CloudFront, S3, WAF, Route 53, CDN " +
					"logs, SNS/Chatbot) to stand up the new environment without writing new " +
					"infrastructure code.",
				"Added 15+ Vitest suites covering chart interaction, eGFR/CKD calculation " +
					"utilities, and patient-detail components.",
				"Shipped PDF export for blood-glucose, weight, and blood-pressure lifelogs.",
				"Draft bilingual (Japanese/English) architecture design documents with Claude as " +
					"a first-pass writing partner, then refine them before they pass team review and " +
					"implementation begins.",
			},
		},
		{
			Company:  "Independent Contractor (個人事業主)",
			Role:     "Frontend & API Developer",
			Period:   "Sep 2025 – Feb 2026",
			Location: "Tokyo, Japan",
			Meta:     "Working language: Japanese",
			Blurb: "Contracted to a major Japanese electric power utility — customer-facing contract " +
				"information portal and internal contract management system. Team of 15.",
			Responsibilities: []string{
				"Groomed the backlog with the client's product owners and turned agreed items " +
					"into implementation-ready specifications.",
				"Evaluated and selected libraries against evolving product requirements.",
			},
			Achievements: []string{
				"Delivered React features end to end for the contract information portal, from " +
					"design and implementation through testing and stakeholder demos.",
				"Ran source control, ticket flow, and technical documentation in Gitea, Jira, and " +
					"Confluence, and owned the release checklist for each delivery.",
			},
		},
		{
			Company:  "ESS Co., Ltd.",
			Role:     "Frontend & API Developer",
			Period:   "Jul 2024 – Aug 2025",
			Location: "Tokyo, Japan",
			Meta:     "Working language: Japanese",
			Blurb: "Full-time employee on the same electric-utility engagement, continued afterward " +
				"as an independent contractor.",
			Responsibilities: []string{
				"Produced technical specifications, process documentation, and release notes for " +
					"each delivery.",
				"Took part in onsite design discussions conducted entirely in Japanese.",
			},
			Achievements: []string{
				"Validated requirements and UX flows directly in Japanese with the client's " +
					"product owners and business users.",
				"Designed, implemented, tested, and demoed TypeScript features for both the " +
					"customer portal and the internal management system.",
				"Evaluated and selected libraries against changing requirements, and wrote the " +
					"processing-detail specifications the team built from.",
				"Adopted Claude and ChatGPT into daily development workflow after relocating to " +
					"Japan, for code authoring and debugging, test writing, and technical " +
					"documentation.",
			},
		},
		{
			Company:  "INTSIG Information Co., Ltd.",
			Role:     "Software Engineer, Maintenance & Development Division",
			Period:   "Mar 2022 – May 2024",
			Location: "Shanghai, China",
			Blurb: "Manufacturing parts management platform, plus OCR products for the education " +
				"market. Team of 11; publicly listed company.",
			Responsibilities: []string{
				"Triaged incoming defect reports from users and QA and prioritized them against " +
					"the feature backlog.",
				"Coordinated release cutovers with business teams across two product lines.",
				"Built parallel-request aggregation in Next.js API Routes for the home and " +
					"document-detail pages, combining data from the OCR engine, document-history, " +
					"and permission services into a single frontend-ready payload.",
				"Implemented SSE streaming proxying in Next.js API Routes for calls to backend " +
					"AI/LLM and file-processing services, keeping internal service domains hidden " +
					"from the client and working around cross-origin restrictions.",
				"Wrote Next.js Middleware for centralized cookie/token validation with silent " +
					"refresh, intercepting invalid requests before they reached application routes.",
			},
			Achievements: []string{
				"Owned maintenance and net-new feature development on the parts-management " +
					"platform using TypeScript and React (Next.js).",
				"Extended the Next.js frontend into a lightweight BFF layer — service " +
					"aggregation across OCR, document-history, and permission services, streaming " +
					"proxying for AI/LLM calls, and centralized auth middleware — to keep internal " +
					"service topology out of the client.",
				"Investigated frontend production defects reported by users and QA, and shipped " +
					"fixes that cut operational workload and overall project cost.",
				"Contributed to a homework-recognition OCR mobile application used by " +
					"approximately 4,000 middle-school students in Shanghai.",
				"Kept operational documentation current so business teams could cut over to new " +
					"releases without engineering support.",
				"Adopted Cursor as an AI-assisted IDE starting in 2023, using it for day-to-day " +
					"feature development and defect fixes on the parts-management platform.",
			},
		},
		{
			Company:  "GTMap Information Industry Co., Ltd.",
			Role:     "Full-Stack Engineer, Web Map Development Dept.",
			Period:   "Jul 2019 – Feb 2022",
			Location: "Nanjing, China",
			Blurb: "Mobile and browser map applications visualizing statistical and geospatial data. " +
				"Team of 14; publicly listed company.",
			Responsibilities: []string{
				"Collected and triaged enterprise-customer feedback into the product backlog.",
				"Used Nuxt Server Routes (`/server/api`) to aggregate layer metadata, " +
					"permission config, and map marker data from multiple GIS endpoints into a " +
					"single response for page load.",
				"Performed spatial coordinate-system conversion (e.g., GCJ-02 to WGS-84) and " +
					"lightweight data filtering on the Nuxt server side to keep heavy computation " +
					"off the browser.",
				"Used Nitro Server Middleware to uniformly intercept map-layer requests, attach " +
					"a shared API key, and forward them to the geo-info backend service.",
			},
			Achievements: []string{
				"Launched a map platform adopted by approximately 5,000 enterprise customers, " +
					"making statistical map data available even for remote regions.",
				"Built map features in TypeScript with Vue and Nuxt.",
				"Built a Nuxt/Nitro BFF layer for the map platform: server-route aggregation of " +
					"GIS endpoints, GCJ-02/WGS-84 coordinate conversion, and a shared middleware " +
					"gateway for geo-info backend calls.",
				"Planned one-month sprints and prioritized the backlog with product stakeholders.",
				"Ran end-to-end QA cycles and turned user feedback into design changes and " +
					"production patches.",
			},
		},
	},

	Education: Education{
		Degree: "Bachelor of Engineering, Computer Science and Technology",
		School: "Suzhou Institute of Technology, Sep 2015 – Jun 2019",
	},

	Languages: []LangEntry{
		{"Japanese", "JLPT N1; daily working language for the past 2 years"},
		{"English", "business level; TOEIC 790 (Listening 450 / Reading 340, Aug 2026); authoring production design documents in English"},
		{"Chinese", "native"},
		{"Work authorization", "Engineer / Specialist in Humanities / International Services visa, " +
			"valid to July 16, 2029. No sponsorship required."},
	},

	PersonalStatement: "I pair seven years of front-end development with a habit of writing things " +
		"down — design documents, test suites, and runbooks that outlive any single release. " +
		"Working onsite in Japan taught me to run design discussions in Japanese while keeping " +
		"written artifacts readable in English, which is how I keep bilingual teams aligned. " +
		"I am looking for a team where I can keep owning features end to end, from architecture " +
		"through production.",
}
