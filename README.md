# 数据

来源: [IPinfo](https://ipinfo.io/account/data-downloads), [IP2Location](https://lite.ip2location.com/database-download), [mmdb_china_ip_list](https://github.com/alecthw/mmdb_china_ip_list)

# 文件

## country.mmdb(ip2location+cn)

[https://github.com/xream/geoip/releases/latest/download/country.mmdb](https://github.com/xream/geoip/releases/latest/download/country.mmdb)

## geoip.dat(ip2location+cn)

[https://github.com/xream/geoip/releases/latest/download/geoip.dat](https://github.com/xream/geoip/releases/latest/download/geoip.dat)

## ipinfo.country.mmdb

[https://github.com/xream/geoip/releases/latest/download/ipinfo.country.mmdb](https://github.com/xream/geoip/releases/latest/download/ipinfo.country.mmdb)

## ipinfo.asn.mmdb

[https://github.com/xream/geoip/releases/latest/download/ipinfo.asn.mmdb](https://github.com/xream/geoip/releases/latest/download/ipinfo.asn.mmdb)

## 社群

👏🏻 欢迎加入社群进行交流讨论

👥 群组 [折腾啥(群组)](https://t.me/zhetengsha_group)

📢 频道 [折腾啥(频道)](https://t.me/zhetengsha)

> 以下为原项目的 README.md

# 简介

本项目每周四自动生成 GeoIP 文件，同时提供命令行界面（CLI）供用户自行定制 GeoIP 文件，包括但不限于 V2Ray dat 格式路由规则文件 `geoip.dat` 和 MaxMind mmdb 格式文件 `Country.mmdb`。

This project releases GeoIP files automatically every Thursday. It also provides a command line interface(CLI) for users to customize their own GeoIP files, included but not limited to V2Ray dat format file `geoip.dat` and MaxMind mmdb format file `Country.mmdb`.

## 与官方版 GeoIP 的区别

- 中国大陆 IPv4 地址数据融合了 [IPIP.net](https://github.com/17mon/china_ip_list/blob/master/china_ip_list.txt) 和 [@gaoyifan/china-operator-ip](https://github.com/gaoyifan/china-operator-ip/blob/ip-lists/china.txt)
- 中国大陆 IPv6 地址数据融合了 MaxMind GeoLite2 和 [@gaoyifan/china-operator-ip](https://github.com/gaoyifan/china-operator-ip/blob/ip-lists/china6.txt)
- 新增类别（方便有特殊需求的用户使用）：
  - `geoip:cloudflare`（`GEOIP,CLOUDFLARE`）
  - `geoip:cloudfront`（`GEOIP,CLOUDFRONT`）
  - `geoip:facebook`（`GEOIP,FACEBOOK`）
  - `geoip:fastly`（`GEOIP,FASTLY`）
  - `geoip:google`（`GEOIP,GOOGLE`）
  - `geoip:netflix`（`GEOIP,NETFLIX`）
  - `geoip:telegram`（`GEOIP,TELEGRAM`）
  - `geoip:twitter`（`GEOIP,TWITTER`）

## 参考配置

在 [V2Ray](https://github.com/v2fly/v2ray-core) 中使用本项目 `.dat` 格式文件的参考配置：

```json
"routing": {
  "rules": [
    {
      "type": "field",
      "outboundTag": "Direct",
      "ip": [
        "geoip:cn",
        "geoip:private",
        "ext:cn.dat:cn",
        "ext:private.dat:private",
        "ext:geoip-only-cn-private.dat:cn",
        "ext:geoip-only-cn-private.dat:private"
      ]
    },
    {
      "type": "field",
      "outboundTag": "Proxy",
      "ip": [
        "geoip:us",
        "geoip:jp",
        "geoip:facebook",
        "geoip:telegram",
        "ext:geoip-asn.dat:facebook",
        "ext:geoip-asn.dat:telegram"
      ]
    }
  ]
}
```

在 [Clash](https://github.com/Dreamacro/clash) 中使用本项目 `.mmdb` 格式文件的参考配置：

```yaml
rules:
  - GEOIP,PRIVATE,policy,no-resolve
  - GEOIP,FACEBOOK,policy
  - GEOIP,CN,policy,no-resolve
```

在 [Leaf](https://github.com/eycorsican/leaf) 中使用本项目 `.mmdb` 格式文件的参考配置，查看[官方 README](https://github.com/eycorsican/leaf/blob/master/README.zh.md#geoip)。

## 下载地址

> 如果无法访问域名 `raw.githubusercontent.com`，可以使用第二个地址 `cdn.jsdelivr.net`。
> *.sha256sum 为校验文件。

### V2Ray dat 格式路由规则文件

> 适用于 [V2Ray](https://github.com/v2fly/v2ray-core)、[Xray-core](https://github.com/XTLS/Xray-core) 和 [Trojan-Go](https://github.com/p4gefau1t/trojan-go)。

- **geoip.dat**：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip.dat](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip.dat)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip.dat](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip.dat)
- **geoip.dat.sha256sum**：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip.dat.sha256sum](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip.dat.sha256sum)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip.dat.sha256sum](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip.dat.sha256sum)
- **geoip-only-cn-private.dat**（精简版 GeoIP，只包含 `geoip:cn` 和 `geoip:private`）：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip-only-cn-private.dat](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip-only-cn-private.dat)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip-only-cn-private.dat](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip-only-cn-private.dat)
- **geoip-only-cn-private.dat.sha256sum**：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip-only-cn-private.dat.sha256sum](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip-only-cn-private.dat.sha256sum)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip-only-cn-private.dat.sha256sum](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip-only-cn-private.dat.sha256sum)
- **geoip-asn.dat**（精简版 GeoIP，只包含上述新增类别）：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip-asn.dat](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip-asn.dat)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip-asn.dat](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip-asn.dat)
- **geoip-asn.dat.sha256sum**：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip-asn.dat.sha256sum](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/geoip-asn.dat.sha256sum)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip-asn.dat.sha256sum](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/geoip-asn.dat.sha256sum)
- **cn.dat**（精简版 GeoIP，只包含 `geoip:cn`）：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/cn.dat](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/cn.dat)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/cn.dat](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/cn.dat)
- **cn.dat.sha256sum**：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/cn.dat.sha256sum](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/cn.dat.sha256sum)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/cn.dat.sha256sum](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/cn.dat.sha256sum)
- **private.dat**（精简版 GeoIP，只包含 `geoip:private`）：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/private.dat](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/private.dat)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/private.dat](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/private.dat)
- **private.dat.sha256sum**：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/private.dat.sha256sum](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/private.dat.sha256sum)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/private.dat.sha256sum](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/private.dat.sha256sum)

### MaxMind mmdb 格式文件

> 适用于 [Clash](https://github.com/Dreamacro/clash) 和 [Leaf](https://github.com/eycorsican/leaf)。

- **Country.mmdb**：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country.mmdb](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country.mmdb)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country.mmdb](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country.mmdb)
- **Country.mmdb.sha256sum**：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country.mmdb.sha256sum](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country.mmdb.sha256sum)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country.mmdb.sha256sum](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country.mmdb.sha256sum)
- **Country-only-cn-private.mmdb**（精简版 GeoIP，只包含 `GEOIP,CN` 和 `GEOIP,PRIVATE`）：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country-only-cn-private.mmdb](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country-only-cn-private.mmdb)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country-only-cn-private.mmdb](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country-only-cn-private.mmdb)
- **Country-only-cn-private.mmdb.sha256sum**：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country-only-cn-private.mmdb.sha256sum](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country-only-cn-private.mmdb.sha256sum)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country-only-cn-private.mmdb.sha256sum](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country-only-cn-private.mmdb.sha256sum)
- **Country-asn.mmdb**（精简版 GeoIP，只包含上述新增类别）：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country-asn.mmdb](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country-asn.mmdb)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country-asn.mmdb](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country-asn.mmdb)
- **Country-asn.mmdb.sha256sum**：
  - [https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country-asn.mmdb.sha256sum](https://raw.githubusercontent.com/Loyalsoldier/geoip/release/Country-asn.mmdb.sha256sum)
  - [https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country-asn.mmdb.sha256sum](https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country-asn.mmdb.sha256sum)

## 定制 GeoIP 文件

可通过以下几种方式定制 GeoIP 文件：

- **在线生成**：[Fork](https://github.com/Loyalsoldier/geoip/fork) 本仓库后，修改自己仓库内的配置文件 `config.json` 和 GitHub Workflow `.github/workflows/build.yml`
- **本地生成**：
  - 安装 [Golang](https://golang.org/dl/) 和 [Git](https://git-scm.com)
  - 拉取项目代码: `git clone https://github.com/Loyalsoldier/geoip.git`
  - 进入项目根目录：`cd geoip`
  - 修改配置文件 `config.json`
  - 运行代码：`go run ./`

**特别说明：**

- **在线生成**：[Fork](https://github.com/Loyalsoldier/geoip/fork) 本项目后，如果需要使用 MaxMind GeoLite2 Country CSV 数据文件，需要在自己仓库的 **[Settings]** 选项卡的 **[Secrets]** 页面中添加一个名为 **MAXMIND_GEOLITE2_LICENSE** 的 secret，否则 GitHub Actions 会运行失败。这个 secret 的值为 MAXMIND 账号的 LICENSE KEY，需要[**注册 MAXMIND 账号**](https://www.maxmind.com/en/geolite2/signup)后，在[**个人账号管理页面**](https://www.maxmind.com/en/account)左侧边栏的 **[Services]** 项下的 **[My License Key]** 里生成。
- **本地生成**：如果需要使用 MaxMind GeoLite2 Country CSV 数据文件（`GeoLite2-Country-CSV.zip`），需要提前从 MaxMind 下载，或从本项目 [release 分支](https://github.com/Loyalsoldier/geoip/tree/release)[下载](https://github.com/Loyalsoldier/geoip/raw/release/GeoLite2-Country-CSV.zip)，并解压缩到名为 `geolite2` 的目录。

### 概念解析

本项目有两个概念：`input` 和 `output`。`input` 指数据源（data source）及其输入格式，`output` 指数据的去向（data destination）及其输出格式。CLI 的作用就是通过读取配置文件中的选项，聚合用户提供的所有数据源，去重，将其转换为目标格式，并输出到文件。

These two concepts are notable: `input` and `output`. The `input` is the data source and its input format, whereas the `output` is the destination of the converted data and its output format. What the CLI does is to aggregate all input format data, then convert them to output format and write them to GeoIP files by using the options in the config file.

### 支持的格式

关于每种格式所支持的配置选项，查看本项目 [`config-example.json`](https://github.com/Loyalsoldier/geoip/blob/HEAD/config-example.json) 文件。

支持的 `input` 输入格式：

- **text**：纯文本 IP 和 CIDR（例如：`1.1.1.1` 或 `1.0.0.0/24`）
- **private**：局域网和私有网络 CIDR（例如：`192.168.0.0/16` 和 `127.0.0.0/8`）
- **cutter**：用于裁剪前置步骤中的数据
- **v2rayGeoIPDat**：V2Ray GeoIP dat 格式（`geoip.dat`）
- **maxmindMMDB**：MaxMind mmdb 数据格式（`GeoLite2-Country.mmdb`）
- **maxmindGeoLite2CountryCSV**：MaxMind GeoLite2 country CSV 数据（`GeoLite2-Country-CSV.zip`）
- **clashRuleSetClassical**：[classical 类型的 Clash RuleSet](https://github.com/Dreamacro/clash/wiki/premium-core-features#classical)
- **clashRuleSet**：[ipcidr 类型的 Clash RuleSet](https://github.com/Dreamacro/clash/wiki/premium-core-features#ipcidr)
- **surgeRuleSet**：[Surge RuleSet](https://manual.nssurge.com/rule/ruleset.html)

支持的 `output` 输出格式：

- **text**：纯文本 CIDR（例如：`1.0.0.0/24`）
- **v2rayGeoIPDat**：V2Ray GeoIP dat 格式（`geoip.dat`，适用于 [V2Ray](https://github.com/v2fly/v2ray-core)、[Xray-core](https://github.com/XTLS/Xray-core) 和 [Trojan-Go](https://github.com/p4gefau1t/trojan-go)）
- **maxmindMMDB**：MaxMind mmdb 数据格式（`GeoLite2-Country.mmdb`，适用于 [Clash](https://github.com/Dreamacro/clash) 和 [Leaf](https://github.com/eycorsican/leaf)）
- **clashRuleSetClassical**：[classical 类型的 Clash RuleSet](https://github.com/Dreamacro/clash/wiki/premium-core-features#classical)
- **clashRuleSet**：[ipcidr 类型的 Clash RuleSet](https://github.com/Dreamacro/clash/wiki/premium-core-features#ipcidr)
- **surgeRuleSet**：[Surge RuleSet](https://manual.nssurge.com/rule/ruleset.html)

### 注意事项

由于 MaxMind mmdb 文件格式的限制，当不同列表的 IP 或 CIDR 数据有交集或重复项时，后写入的列表的 IP 或 CIDR 数据会覆盖（overwrite）之前已写入的列表的数据。譬如，IP `1.1.1.1` 同属于列表 `AU` 和列表 `Cloudflare`。如果 `Cloudflare` 在 `AU` 之后写入，则 IP `1.1.1.1` 归属于列表 `Cloudflare`。

为了确保某些指定的列表、被修改的列表一定囊括属于它的所有 IP 或 CIDR 数据，可在 `output` 输出格式为 `maxmindMMDB` 的配置中增加选项 `overwriteList`，该选项中指定的列表会在最后逐一写入，列表中最后一项优先级最高。若已设置选项 `wantedList`，则无需设置 `overwriteList`。`wantedList` 中指定的列表会在最后逐一写入，列表中最后一项优先级最高。

## CLI 功能展示

可通过 `go install -v github.com/Loyalsoldier/geoip@latest` 直接安装 CLI。

```bash
$ ./geoip -h
Usage of ./geoip:
  -c string
    	URI of the JSON format config file, support both local file path and remote HTTP(S) URL (default "config.json")
  -l	List all available input and output formats

$ ./geoip -c config.json
2021/08/29 12:11:35 ✅ [v2rayGeoIPDat] geoip.dat --> output/dat
2021/08/29 12:11:35 ✅ [v2rayGeoIPDat] geoip-only-cn-private.dat --> output/dat
2021/08/29 12:11:35 ✅ [v2rayGeoIPDat] geoip-asn.dat --> output/dat
2021/08/29 12:11:35 ✅ [v2rayGeoIPDat] cn.dat --> output/dat
2021/08/29 12:11:35 ✅ [v2rayGeoIPDat] private.dat --> output/dat
2021/08/29 12:11:39 ✅ [maxmindMMDB] Country.mmdb --> output/maxmind
2021/08/29 12:11:39 ✅ [maxmindMMDB] Country-only-cn-private.mmdb --> output/maxmind
2021/08/29 12:11:39 ✅ [text] netflix.txt --> output/text
2021/08/29 12:11:39 ✅ [text] telegram.txt --> output/text
2021/08/29 12:11:39 ✅ [text] cn.txt --> output/text
2021/08/29 12:11:39 ✅ [text] cloudflare.txt --> output/text
2021/08/29 12:11:39 ✅ [text] cloudfront.txt --> output/text
2021/08/29 12:11:39 ✅ [text] facebook.txt --> output/text
2021/08/29 12:11:39 ✅ [text] fastly.txt --> output/text

$ ./geoip -l
All available input formats:
  - v2rayGeoIPDat (Convert V2Ray GeoIP dat to other formats)
  - maxmindMMDB (Convert MaxMind mmdb database to other formats)
  - maxmindGeoLite2CountryCSV (Convert MaxMind GeoLite2 country CSV data to other formats)
  - private (Convert LAN and private network CIDR to other formats)
  - text (Convert plaintext IP & CIDR to other formats)
  - clashRuleSetClassical (Convert classical type of Clash RuleSet to other formats (just processing IP & CIDR lines))
  - clashRuleSet (Convert ipcidr type of Clash RuleSet to other formats)
  - surgeRuleSet (Convert Surge RuleSet to other formats (just processing IP & CIDR lines))
  - cutter (Remove data from previous steps)
  - test (Convert specific CIDR to other formats (for test only))
All available output formats:
  - v2rayGeoIPDat (Convert data to V2Ray GeoIP dat format)
  - maxmindMMDB (Convert data to MaxMind mmdb database format)
  - clashRuleSetClassical (Convert data to classical type of Clash RuleSet)
  - clashRuleSet (Convert data to ipcidr type of Clash RuleSet)
  - surgeRuleSet (Convert data to Surge RuleSet)
  - text (Convert data to plaintext CIDR format)
```

## License

[CC-BY-SA-4.0](https://creativecommons.org/licenses/by-sa/4.0/)

This product includes GeoLite2 data created by MaxMind, available from [MaxMind](http://www.maxmind.com).

## 项目 Star 数增长趋势

[![Stargazers over time](https://starchart.cc/Loyalsoldier/geoip.svg)](https://starchart.cc/Loyalsoldier/geoip)
