package westcn

const (
	// 封装常用的请求Path

	// 获取域名价格
	PATH_INFO_GETPRICE = "/info/?act=getprice"

	// 查询域名
	PATH_DOMAIN_QUERY = "/domain/query/"

	// 获取域名列表
	PATH_DOMAIN_GETDOMAINS = "/domain/?act=getdomains"

	// 域名续费
	PATH_DOMAIN_RENEW = "/domain/?act=renew"

	//  获取域名管理密码
	PATH_DOMAIN_GETPWD = "/domain/?act=getpwd"

	//  修改域名管理密码
	PATH_DOMAIN_MODPWD = "/domain/?act=modpwd"

	// 修改域名DNS
	PATH_DOMAIN_MODDNS = "/domain/?act=moddns"

	// 添加域名解析
	PATH_DOMAIN_ADD_DNS_RECORD = "/domain/?act=adddnsrecord"

	// 修改域名解析
	PATH_DOMAIN_MOD_DNS_RECORD = "/domain/?act=moddnsrecord"

	// 删除域名解析
	PATH_DOMAIN_DEL_DNS_RECORD = "/domain/?act=deldnsrecord"

	// 获取域名解析记录
	PATH_DOMAIN_GET_DNS_RECORD = "/domain/?act=getdnsrecord"

	// 获取域名证书
	PATH_DOMAIN_CERT = "/domain/?act=cert"

	// 设置域名锁定状态
	PATH_DOMAIN_SETLOCK = "/domain/?act=setlock"

	// 获取Whois
	PATH_DOMAIN_WHOIS = "/domain/?act=whois"

	// 获取续费价格
	PATH_DOMAIN_GET_RENEW_PRICE = "/domain/?act=getrenprice"

	// 获取离线记录
	PATH_DOMAIN_GET_OFFLINE = "/domain/?act=offline"

	// 设置注册局安全锁(此功能收费)
	PATH_DOMAIN_BUY_SECURITYLOCK = " /domain/?act=buysecuritylock"

	// 获取注册局安全锁价格
	PATH_DOMAIN_SECURITYLOCK_PRICE = "/domain/?act=securitylockprice"

	// ------------------ 域名模板相关---------------

	// 创建模板
	PATH_AUDIT_SUB = "/audit/?act=auditsub"

	// 修改模板
	PATH_AUDIT_MOD = "/audit/?act=auditmod"

	// 删除模板
	PATH_AUDIT_DEL = "/audit/?act=auditdel"

	// 批量获取模板列表
	PATH_AUDIT_LIST = "/audit/?act=auditlist"

	// 获取模板信息
	PATH_AUDIT_INFO = "/audit/?act=auditinfo"

	// 注册域名
	PATH_AUDIT_REGDOMAIN = "/audit/?act=regdomain"

	// 修改域名信息
	PATH_AUDIT_DOMAIN_MODIFY = "/audit/?act=domainmodisub"

	// 获取域名信息
	PATH_AUDIT_DOMAIN_INFO = "/audit/?act=domaininfo"

	// 获取实名TOKEN
	PATH_AUDIT_UPLOAD_WCF_TOKEN = " /audit/?act=uploadwcftoken"

	// 域名模板过户
	PATH_AUDIT_GHSUB = "/audit/?act=auditghsub"

	// 查询域名过户状态
	PATH_AUDIT_GHLIST = "/audit/?act=ghlist"

	// ------------------ 域名转移相关---------------

)
