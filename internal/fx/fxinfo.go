package fx

var defalutPlugins = []Plugin{

	{
		Id:          "fx-2021-1001",
		Query:       "Google Reverse",
		RuleName:    "Google反代服务器",
		RuleEnglish: "Google Reverse proxy",
		Description: "不用挂代理就可以访问的Google搜索，但搜索记录可能会被记录。",
		Author:      "fofa",
		FofaQuery:   `body="var c = Array.prototype.slice.call(arguments, 1);return function() {var d=c.slice();"`,
		Tag:         []string{"google"},
		Type:        TypeInline,
		Source:      "https://tp.wjx.top/m/67114073.aspx",
	},
	{
		Id:          "fx-2021-1002",
		Query:       "Python SimpleHTTP",
		RuleName:    "Python SimpleHTTP服务器",
		RuleEnglish: "Python SimpleHTTP Server",
		Description: "Python SimpleHTTP临时服务器",
		Author:      "fofa",
		FofaQuery:   `(server="SimpleHTTP/0.6 Python/3.6.3" || server="SimpleHTTP/0.6 Python/3.6.8" || server="SimpleHTTP/0.6 Python/3.7.0" || server="SimpleHTTP/0.6 Python/3.7.1" || server="SimpleHTTP/0.6 Python/3.7.2" || server="SimpleHTTP/0.6 Python/3.7.3" || server="SimpleHTTP/0.6 Python/3.7.4" || server="SimpleHTTP/0.6 Python/3.7.5" || server="SimpleHTTP/0.6 Python/3.7.6" || server="SimpleHTTP/0.6 Python/3.7.7" || server="SimpleHTTP/0.6 Python/3.7.8" || server="SimpleHTTP/0.6 Python/3.7.9" || server="SimpleHTTP/0.6 Python/3.8.1" || server="SimpleHTTP/0.6 Python/3.8.2" || server="SimpleHTTP/0.6 Python/3.8.3" || server="SimpleHTTP/0.6 Python/3.8.4" || server="SimpleHTTP/0.6 Python/3.8.5" || server="SimpleHTTP/0.6 Python/3.8.6" || server="SimpleHTTP/0.6 Python/3.8.7" || server="SimpleHTTP/0.6 Python/3.8.8" || server="SimpleHTTP/0.6 Python/3.8.9" || server="SimpleHTTP/0.6 Python/3.9.1" || server="SimpleHTTP/0.6 Python/3.9.2" || server="SimpleHTTP/0.6 Python/3.9.3" || server="SimpleHTTP/0.6 Python/3.9.4" || server="SimpleHTTP/0.6 Python/3.9.5" || server="SimpleHTTP/0.6 Python/3.9.6" || server="SimpleHTTP/0.6 Python/3.9.7" || server="SimpleHTTP/0.6 Python/3.9.8" || server="SimpleHTTP/0.6 Python/3.9.9") && title="Directory listing for"`,
		Tag:         []string{"file"},
		Type:        TypeInline,
		Source:      "https://tp.wjx.top/m/67114073.aspx",
	},
	{
		Id:          "fx-2021-1003",
		Query:       "xiangqin",
		RuleName:    "相亲约会？",
		RuleEnglish: "Blind Date",
		Description: "相亲约会？",
		Author:      "fofa",
		FofaQuery:   `body="/tpl/static/varpop/css/box.css"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1004",
		Query:       "hfs rce",
		RuleName:    "存在命令执行的HFS服务",
		RuleEnglish: "Presence of HFS services for command execution",
		Description: "这个语法可搜索存在命令执行的HFS服务，使用者多数为抓鸡黑客，可以直接上线能捡到不少有趣的东西~",
		Author:      "fofa",
		FofaQuery:   `body="HttpFileServer v2.3 beta 287"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1005",
		Query:       "Satellites FTP",
		RuleName:    "一键日卫星FTP？",
		RuleEnglish: "Satellites FTP",
		Description: "一键日卫星FTP？",
		Author:      "fofa",
		FofaQuery:   `banner="Cobham SATCOM"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1006",
		Query:       "mk Mining",
		RuleName:    "mk路由器全球挖矿感染",
		RuleEnglish: "mk router global mining infection",
		Description: "mk路由器全球挖矿感染",
		Author:      "fofa",
		FofaQuery:   `app="Mikrotik-HttpProxy"&&(body="CoinHive.Anonymous" || body="CRLT.Anonymous" || body=" WMP.Anonymous(")`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1007",
		Query:       "ss-Manager Login",
		RuleName:    "ss-Manager 登录",
		RuleEnglish: "ss-Manager Login",
		Description: "Shadowsocks-Manager登录界面",
		Author:      "fofa",
		FofaQuery:   `body="indeterminate" && body="MainController" && header="X-Powered-By: Express"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1008",
		Query:       "Heating monitoring",
		RuleName:    "供暖监控系统",
		RuleEnglish: "Heating monitoring systems",
		Description: "供暖监控系统",
		Author:      "fofa",
		FofaQuery:   `body="s1v13.htm"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1009",
		Query:       "Free Proxy",
		RuleName:    "免费代理池",
		RuleEnglish: "Free Proxy Pool",
		Description: "获取免费代理池",
		Author:      "fofa",
		FofaQuery:   `body="get all proxy from proxy pool"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1010",
		Query:       "honeypot-01",
		RuleName:    "蜜罐",
		RuleEnglish: "honeypot 01",
		Description: "蜜罐",
		Author:      "fofa",
		FofaQuery:   `(header="uc-httpd 1.0.0" && server="JBoss-5.0") || server="Apache,Tomcat,Jboss,weblogic,phpstudy,struts"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1011",
		Query:       "hacked by",
		RuleName:    "被挂黑的站点",
		RuleEnglish: "hacked by xxx",
		Description: "被挂黑的站点",
		Author:      "fofa",
		FofaQuery:   `body="hacked by"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1012",
		Query:       "jupyter unauthorized",
		RuleName:    "jupyter 未授权访问",
		RuleEnglish: "jupyter unauthorized",
		Description: "jupyter 未授权访问",
		Author:      "xiecat",
		FofaQuery:   `(body="ipython-main-app" && title="Home Page - Select or create a notebook")"`,
		Tag:         []string{"xiecat"},
		Type:        TypeInline,
		Source:      "",
	},
}
