module go-vuln-demo

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3 // Contains CVE-2020-28483
	github.com/tidwall/gjson v1.6.0 // Contains CVE-2020-36067
	golang.org/x/text v0.3.0 // Contains CVE-2020-14040
)