package resume

// ZH is the Chinese resume content, feeding zh/index.html only (no short
// form exists for this language — see internal/resume/CLAUDE.md).
var ZH = ResumeData{
	Name:          "王鸿乘（Hongcheng Wang）",
	Tagline:       "前端工程师　·　AWS交付基建　·　日语N1 / 英语商务水平",
	Location:      "日本千叶",
	Email:         "wanghongcheng10@gmail.com",
	LinkedInLabel: "linkedin.com/in/hongcheng-wang-1597a6381",
	LinkedInURL:   "https://linkedin.com/in/hongcheng-wang-1597a6381",

	SummaryLong: []string{
		"拥有7年生产环境Web应用开发经验，涉及医疗、电力、制造、地理空间等多个行业领域，" +
			"横跨中国与日本。目前在东京负责医生端健康档案（PHR）平台的开发。",
		"以TypeScript（Vue 3、React、Next.js、Nuxt）为核心的前端深度交付经验，包括在" +
			"Next.js/Nuxt服务端路由中的BFF层工程实践——跨微服务的服务聚合、流式代理、鉴权" +
			"中间件与服务端数据整形——同时负责基于S3、CloudFront、Terraform、GitHub Actions " +
			"OIDC的AWS交付流水线搭建。",
		"在日本现场用日语与相关方对接，同时用日语与英语撰写生产环境架构设计文档；在敏捷" +
			"与瀑布两种开发模式下，全程负责从待办事项梳理到生产环境发布的完整流程。",
		"自2023年起践行AI辅助开发：在中国期间使用Cursor进行日常开发；赴日后使用Claude与" +
			"ChatGPT，用于代码编写与调试、测试生成、日英双语架构文档撰写与代码审查。",
	},

	CoreStrengths: []string{
		"7年生产环境Web应用交付经验，全程负责从待办梳理到生产发布的完整流程。",
		"前端以TypeScript为核心的深厚实战能力——Vue 3（Composition API、Pinia、" +
			"Vue Router）与React——端到端构建生产环境SPA。",
		"在Next.js与Nuxt服务端路由中的BFF层工程实践——多服务聚合、流式代理、鉴权中间件、" +
			"服务端数据整形——将前端职责延伸至后端服务边界。",
		"涉猎医疗、电力、制造、地理空间等多个行业领域的项目经验。",
		"AWS上的基础设施即代码（IaC）交付：Terraform、CloudFormation，以及基于OIDC认证的" +
			"GitHub Actions流水线。",
		"2023年起的AI辅助开发实践——中国期间使用Cursor，赴日后使用Claude与ChatGPT——" +
			"融入日常编码、调试、测试、文档撰写与代码审查全流程。",
		"曾带领最多4人规模的工程师团队及跨部门相关方（项目负责人经验）。",
	},

	Skills: []SkillEntry{
		{"编程语言", "TypeScript, JavaScript, Go, Shell"},
		{"前端", "Vue 3（Composition API、Pinia、Vue Router、vue-i18n）、React、Next.js、Nuxt、" +
			"Angular、Chart.js、Vite、Sass/SCSS"},
		{"BFF / 服务端", "Next.js（API Routes、App Router、Middleware）、Nuxt 3（Nitro Server Routes、" +
			"Server Middleware）、SSE流式代理、服务聚合、服务端缓存、数据脱敏/DTO裁剪"},
		{"云与DevOps", "AWS（S3、CloudFront、WAF、Route 53、IAM/OIDC）、Terraform、CloudFormation、" +
			"GitHub Actions、CodeDeploy、CI/CD流水线设计"},
		{"测试与质量", "Vitest、Vue Test Utils、ESLint、Prettier、Husky"},
		{"工具与流程", "Git/GitHub、Jira、Confluence、Gitea、敏捷/Scrum"},
		{"AI辅助开发", "Cursor（2023年，中国期间生产代码开发）；Claude、ChatGPT（2024年至今，日本）" +
			"——用于代码编写、调试、测试生成、文档撰写与代码审查"},
	},

	Experience: []Job{
		{
			Company:  "Welby, Inc.（日本）",
			Role:     "前端工程师（服务开发部）",
			Period:   "2026年3月 – 至今",
			Location: "日本·东京",
			Meta:     "工作语言：日语",
			Blurb: "面向慢性病诊疗的PHR（个人健康档案）平台，基于Vue 3 + TypeScript的单页应用。" +
				"东京证券交易所Growth市场上市公司。",
			Responsibilities: []string{
				"使用Vue 3、TypeScript、Pinia、Chart.js开发与维护医生端SPA的各项功能。",
				"负责mykarte SPA在AWS上的交付架构设计，并在实现前提交团队评审。",
				"在开发功能的同时同步编写Vitest与Vue Test Utils测试用例，而非事后补充。",
			},
			Achievements: []string{
				"为慢性肾脏病（CKD）患者构建LTEP（长期eGFR预测）图表——基于Chart.js的交互式" +
					"可视化功能，支持可拖拽干预节点、可配置基线与专用打印布局，4个月内完成12个" +
					"功能分支的迭代交付。",
				"设计mykarte SPA的S3+CloudFront+GitHub Actions交付架构：用OIDC短期凭证" +
					"替换长期有效的AWS密钥，应用覆盖SPA兜底响应在内的CloudFront响应头策略" +
					"（HSTS、CSP、X-Frame-Options），以COUNT模式部署WAFv2，并将生产环境发布" +
					"置于二次确认流程之后。",
				"复用现有Terraform模块集（CloudFront、S3、WAF、Route 53、CDN日志、" +
					"SNS/Chatbot）搭建新环境，无需编写新的基础设施代码。",
				"新增15个以上Vitest测试套件，覆盖图表交互、eGFR/CKD计算工具与患者详情组件。",
				"上线血糖、体重、血压生活日志的PDF导出功能。",
				"以Claude作为初稿撰写伙伴，起草日语与英语双语的架构设计文档，经本人精修后" +
					"提交团队评审、进入实现阶段。",
			},
		},
		{
			Company:  "个人事业主（自由职业开发者，日本）",
			Role:     "前端与API开发工程师",
			Period:   "2025年9月 – 2026年2月",
			Location: "日本·东京",
			Meta:     "工作语言：日语",
			Blurb: "承接日本大型电力公司项目（业务委托形式）——面向客户的合同信息门户及" +
				"公司内部合同管理系统。团队规模15人。",
			Responsibilities: []string{
				"与客户产品负责人梳理需求，将确认事项转化为可实施的规格文档。",
				"根据不断变化的产品需求评估并选型技术库。",
			},
			Achievements: []string{
				"使用React端到端交付前端功能——从设计、实现到测试与客户演示。",
				"在Gitea、Jira、Confluence上管理源码、工单流程与技术文档，并负责每次交付的" +
					"发布检查清单。",
			},
		},
		{
			Company:  "ESS Co., Ltd.（日本）",
			Role:     "前端与API开发工程师",
			Period:   "2024年7月 – 2025年8月",
			Location: "日本·东京",
			Meta:     "工作语言：日语",
			Blurb: "以正式员工身份参与同一电力公司项目，合同结束后转为个人事业主（自由职业）" +
				"继续该项目。",
			Responsibilities: []string{
				"为每次交付撰写技术规格书、流程文档与发布说明。",
				"参与全程以日语进行的现场设计讨论。",
			},
			Achievements: []string{
				"直接用日语与客户产品负责人及业务人员核实需求与用户体验流程。",
				"用TypeScript为客户门户与内部管理系统设计、实现、测试并演示前端功能。",
				"根据不断变化的需求评估并选型技术库，并撰写团队据以实现的处理细节规格书。",
				"赴日后将Claude与ChatGPT纳入日常开发流程，用于代码编写与调试、测试用例编写" +
					"及技术文档撰写。",
			},
		},
		{
			Company:  "上海合合信息科技股份有限公司（INTSIG）",
			Role:     "前端与API开发工程师（维护开发部）",
			Period:   "2022年3月 – 2024年5月",
			Location: "中国·上海",
			Blurb: "面向制造业的零部件管理平台，及面向教育市场的OCR产品。团队规模11人；" +
				"上市公司。",
			Responsibilities: []string{
				"对用户与QA提交的缺陷报告进行分诊，并按优先级纳入功能待办列表。",
				"跨两条产品线协调发布切换事宜，与业务团队对接。",
			},
			Achievements: []string{
				"使用TypeScript与React（Next.js）负责零部件管理平台的前端维护与新功能开发。",
				"调查用户与QA反馈的生产环境缺陷，并交付修复方案，降低运维负担与项目整体" +
					"成本。",
				"参与一款服务于上海约4000名初中生的作业识别OCR移动应用的开发。",
				"持续维护运维文档，使业务团队无需工程支持即可完成新版本切换。",
				"自2023年起采用Cursor作为AI辅助IDE，用于零部件管理平台的日常功能开发与" +
					"缺陷修复。",
			},
		},
		{
			Company:  "南京国图信息有限公司（GTMap）",
			Role:     "前端工程师（Web地图开发部）",
			Period:   "2019年7月 – 2022年2月",
			Location: "中国·南京",
			Blurb: "面向统计与地理空间数据可视化的移动端与浏览器端地图应用。团队规模14人；" +
				"上市公司。",
			Responsibilities: []string{
				"收集并分诊企业客户反馈，纳入产品待办列表。",
			},
			Achievements: []string{
				"上线服务约5000家企业客户的地图平台，使偏远地区也能获取统计地图数据。",
				"用TypeScript结合Vue与Nuxt构建地图功能。",
				"规划以一个月为周期的迭代排期，并与产品相关方共同确定待办优先级。",
				"执行端到端QA循环，将用户反馈转化为设计改动与生产环境补丁。",
			},
		},
	},

	Education: Education{
		Degree: "工学学士，计算机科学与技术",
		School: "苏州理工学院　·　2015年9月 – 2019年6月",
	},

	Languages: []LangEntry{
		{"日语", "JLPT N1；近2年的日常工作语言"},
		{"英语", "商务水平；TOEIC 790分（听力 450 / 阅读 340，2026年8月）；具备英文生产环境架构文档撰写经验"},
	},

	PersonalStatement: "7年前端开发经验之外，我还积累了AWS交付基础设施搭建与AI辅助后端维护的实战经验，" +
		"并保持撰写设计文档、测试用例与运维手册的习惯——这些内容在版本发布后依然能长期" +
		"沉淀下来。在日本现场工作的经历让我习惯于用日语推进设计讨论，同时保证书面成果物" +
		"在英语环境下同样可读，这也是我协调多语言团队认知一致的方式。我希望加入一个能让" +
		"我从架构设计到生产发布全程负责功能交付的团队。",
}
