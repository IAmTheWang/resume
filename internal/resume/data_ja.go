package resume

// JA is the Japanese resume content, feeding ja/index.html only (no short
// form exists for this language — see internal/resume/CLAUDE.md).
var JA = ResumeData{
	Name:          "王 鴻乗（Hongcheng Wang）",
	Tagline:       "フロントエンドエンジニア　｜　日英バイリンガル（JLPT N1）",
	Location:      "千葉県、日本",
	Email:         "wanghongcheng10@gmail.com",
	LinkedInLabel: "linkedin.com/in/hongcheng-wang-1597a6381",
	LinkedInURL:   "https://linkedin.com/in/hongcheng-wang-1597a6381",

	SummaryLong: []string{
		"中国・日本で7年間、医療・電力・製造・地理空間の各分野において本番運用のWebアプリケーション開発に従事。" +
			"現在は東京にて、医師向け健康記録（PHR）プラットフォームの開発を担当。",
		"TypeScript（Vue 3、React、Next.js、Nuxt）によるフロントエンド開発を中核としつつ、" +
			"S3・CloudFront・Terraform・GitHub Actions OIDCを用いたAWS配信パイプラインの構築・運用にも従事。",
		"日本人関係者と日本語で現地対応しながら、本番アーキテクチャ設計文書を日本語・英語の両方で執筆。" +
			"アジャイル・ウォーターフォール双方の開発サイクルにおいて、バックログ整理から本番リリースまでを一貫して担当。",
		"2023年からAI活用開発を実践：中国では日々の開発業務でCursorを使用し、日本移住後はClaudeおよび" +
			"ChatGPTを、コード作成・デバッグ、テスト生成、日英バイリンガルのアーキテクチャ設計文書作成、" +
			"コードレビューに活用。",
	},

	CoreStrengths: []string{
		"バックログ整理から本番リリースまで、7年間一貫して本番Webアプリケーションの提供を担当。",
		"TypeScriptによるフロントエンド実務経験——Vue 3（Composition API、Pinia、Vue Router）" +
			"およびReactを用いた本番SPA開発。",
		"医療・電力・製造・地理空間分野にまたがる業務経験。",
		"AWS上でのInfrastructure as Code——Terraform、CloudFormation、OIDC認証によるGitHub Actionsパイプライン。",
		"2023年からのAI活用開発——中国ではCursor、日本移住後はClaudeおよびChatGPTを、日々のコーディング・" +
			"デバッグ・テスト・ドキュメント作成・コードレビューに組み込んで活用。",
		"最大4名のエンジニア・関係部門とのチームリードとしての経験。",
	},

	Skills: []SkillEntry{
		{"プログラミング言語", "TypeScript, JavaScript, Go, Shell"},
		{"フロントエンド", "Vue 3（Composition API, Pinia, Vue Router, vue-i18n）, React, Next.js, Nuxt, " +
			"Angular, Chart.js, Vite, Sass/SCSS"},
		{"クラウド・DevOps", "AWS（S3, CloudFront, WAF, Route 53, IAM/OIDC）, Terraform, CloudFormation, " +
			"GitHub Actions, CodeDeploy, CI/CDパイプライン設計"},
		{"テスト・品質", "Vitest, Vue Test Utils, ESLint, Prettier, Husky"},
		{"ツール・プロセス", "Git / GitHub, Jira, Confluence, Gitea, アジャイル / スクラム"},
		{"AI活用開発", "Cursor（2023年、中国での実務コード開発）；Claude、ChatGPT（2024年〜現在、日本）—— " +
			"コード作成・デバッグ、テスト生成、ドキュメント作成、コードレビューに活用"},
	},

	Experience: []Job{
		{
			Company:  "Welby, Inc.",
			Role:     "フルスタックエンジニア（サービス開発部）",
			Period:   "2026年3月〜現在",
			Location: "日本・東京",
			Meta:     "使用言語：日本語",
			Blurb: "慢性疾患診療向けPHR（Personal Health Record）プラットフォーム。" +
				"Vue 3 + TypeScriptによるSPA。東証グロース市場上場企業。",
			Responsibilities: []string{
				"Vue 3、TypeScript、Pinia、Chart.jsを用いて、医師向けSPAの機能開発・保守を担当。",
				"mykarte SPAのAWS配信アーキテクチャを設計し、実装前にチームレビューを実施。",
				"実装と並行してVitestおよびVue Test Utilsによるテストを作成。",
			},
			Achievements: []string{
				"慢性腎臓病（CKD）患者向けのLTEP（Long-Term eGFR Progression）チャートを構築。" +
					"ドラッグ操作可能な介入ポイント、設定可能なベースライン、専用の印刷レイアウトを備えた" +
					"Chart.jsによるインタラクティブな可視化機能を、4か月間で12のフィーチャーブランチにわたり提供。",
				"mykarte SPAのS3＋CloudFront＋GitHub Actionsによる配信アーキテクチャを設計。" +
					"長期間有効なAWSキーをOIDCによる短期認証情報に置き換え、SPAのフォールバック応答も含めて" +
					"カバーするCloudFront Response Headers Policy（HSTS、CSP、X-Frame-Options）を適用。" +
					"WAFv2をCOUNTモードで導入し、本番環境へのリリースを型入力による確認フローの背後に配置。",
				"既存のTerraformモジュール群（CloudFront、S3、WAF、Route 53、CDNログ、SNS/Chatbot）を" +
					"再利用し、新規インフラコードを書かずに新環境を構築。",
				"チャート操作、eGFR/CKD計算ユーティリティ、患者詳細コンポーネントをカバーする" +
					"Vitestテストスイートを15件以上追加。",
				"血糖値・体重・血圧のライフログのPDFエクスポート機能を提供。",
				"Claudeを一次執筆パートナーとして日本語・英語両方のアーキテクチャ設計文書を起草し、" +
					"チームレビューを経て実装に着手する前に推敲。",
			},
		},
		{
			Company:  "個人事業主（Independent Contractor）",
			Role:     "フロントエンド・APIエンジニア",
			Period:   "2025年9月〜2026年2月",
			Location: "日本・東京",
			Meta:     "使用言語：日本語",
			Blurb: "大手電力会社案件（業務委託）。顧客向け契約情報ポータルおよび社内契約管理システム。" +
				"チーム規模15名。",
			Responsibilities: []string{
				"クライアントのプロダクトオーナーとバックログの整理を行い、合意事項を実装可能な仕様に落とし込み。",
				"変化するプロダクト要件に応じてライブラリの選定・評価を実施。",
			},
			Achievements: []string{
				"Reactを用いて、契約情報ポータル向け機能を設計・実装からテスト・関係者へのデモまで" +
					"エンドツーエンドで提供。",
				"Gitea、Jira、Confluenceでのソース管理・チケット運用・技術文書管理を担当し、" +
					"各リリースのチェックリストを主導。",
			},
		},
		{
			Company:  "ESS Co., Ltd.",
			Role:     "フロントエンド・APIエンジニア",
			Period:   "2024年7月〜2025年8月",
			Location: "日本・東京",
			Meta:     "使用言語：日本語",
			Blurb:    "同一の電力会社案件に正社員として従事。契約終了後は個人事業主として継続。",
			Responsibilities: []string{
				"各リリースにおける技術仕様書、プロセス文書、リリースノートを作成。",
				"全編日本語で実施される現地での設計討議に参加。",
			},
			Achievements: []string{
				"クライアントのプロダクトオーナー・業務担当者と日本語で直接要件・UXフローを検証。",
				"顧客向けポータルおよび社内管理システムの両方について、TypeScriptで" +
					"機能の設計・実装・テスト・デモを担当。",
				"変化する要件に応じてライブラリを選定・評価し、チームが実装のベースとする" +
					"処理詳細仕様書を作成。",
				"日本移住後、ClaudeおよびChatGPTを日々の開発ワークフローに導入し、コード作成・デバッグ、" +
					"テストコード作成、技術文書作成に活用。",
			},
		},
		{
			Company:  "INTSIG Information Co., Ltd.",
			Role:     "ソフトウェアエンジニア（保守開発部）",
			Period:   "2022年3月〜2024年5月",
			Location: "中国・上海",
			Blurb:    "製造業向け部品管理プラットフォームおよび教育市場向けOCR製品。チーム規模11名。上場企業。",
			Responsibilities: []string{
				"ユーザーおよびQAからの不具合報告をトリアージし、機能バックログと優先順位付け。",
				"2つのプロダクトラインにまたがるリリース切替を事業部門と調整。",
			},
			Achievements: []string{
				"TypeScript、React（Next.js）を用いて、" +
					"部品管理プラットフォームの保守および新規機能開発を担当。",
				"ユーザー・QAから報告されたフロントエンド側の本番障害を調査し、運用負荷とプロジェクトコスト全体を" +
					"削減する修正を提供。",
				"上海の中学生約4,000人が利用する宿題認識OCRモバイルアプリの開発に貢献。",
				"事業部門がエンジニアリング側の支援なしに新リリースへ切り替えられるよう、" +
					"運用ドキュメントを常に最新に維持。",
				"2023年よりAI支援型IDEとしてCursorを導入し、部品管理プラットフォームの日々の" +
					"機能開発・不具合修正に活用。",
			},
		},
		{
			Company:  "GTMap Information Industry Co., Ltd.",
			Role:     "フルスタックエンジニア（Webマップ開発部）",
			Period:   "2019年7月〜2022年2月",
			Location: "中国・南京",
			Blurb: "統計・地理空間データを可視化するモバイル・ブラウザ向け地図アプリケーション。" +
				"チーム規模14名。上場企業。",
			Responsibilities: []string{
				"法人顧客からのフィードバックを収集し、プロダクトバックログへ反映。",
			},
			Achievements: []string{
				"約5,000社の法人顧客に採用された地図プラットフォームをリリースし、遠隔地域でも" +
					"統計地図データを利用可能に。",
				"TypeScriptとVue、Nuxtで地図機能を構築。",
				"1か月単位のスプリントを計画し、プロダクト関係者とバックログの優先順位付けを実施。",
				"エンドツーエンドのQAサイクルを運用し、ユーザーフィードバックを設計変更・本番パッチへ反映。",
			},
		},
	},

	Education: Education{
		Degree: "学士（工学）　情報工学",
		School: "蘇州理工学院　2015年9月〜2019年6月",
	},

	Languages: []LangEntry{
		{"日本語", "JLPT N1；直近2年間の日常業務言語"},
		{"英語", "ビジネスレベル；TOEIC 790点（Listening 450 / Reading 340、2026年8月）；英語での設計文書執筆実績あり"},
		{"中国語", "ネイティブ"},
		{"就労資格", "技術・人文知識・国際業務ビザ（有効期限：2029年7月16日）。就労ビザのスポンサーシップ不要。"},
	},

	PersonalStatement: "フロントエンド開発を中心に7年の経験を積む中で、設計文書やテストコード、" +
		"運用手順書など、リリース後も残る記録を残すことを習慣としてきました。日本での現地就業を通じて、" +
		"設計討議は日本語で進めながら、成果物としてのドキュメントは英語でも読める状態を保つ働き方を" +
		"身につけ、これによりバイリンガルなチームの認識を揃えてきました。アーキテクチャ設計から" +
		"本番リリースまで、機能をエンドツーエンドで担当し続けられるチームを探しています。",
}
