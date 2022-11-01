package qychan

import (
	"fmt"
	"github.com/heyangwenphp/qychan/util"
	"strconv"
	"time"
)

//创建合约
func InitLedger(appid string, secrect string, name string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["name"] = strconv.QuoteToASCII(name)
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	res := util.PostJsonRequest("initLedger", parameter)
	return res
}

//创建用户账户
func CreateAccountAddress(appid string, secrect string, name string, mobile string, idc string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["name"] = strconv.QuoteToASCII(name)
	parameter["mobile"] = mobile
	parameter["idc"] = idc
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	res := util.PostJsonRequest("createAccountAddress", parameter)
	return res
}

//创建一个nft资源
func CreateNftInfo(appid string, secrect string, author string, title string, series_name string, series_id string, url string) interface{} {

	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["author"] = strconv.QuoteToASCII(author)
	parameter["title"] = strconv.QuoteToASCII(title)
	parameter["series_name"] = strconv.QuoteToASCII(series_name)
	parameter["series_id"] = series_id
	parameter["url"] = url
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	res := util.PostJsonRequest("createNftInfo", parameter)
	return res
}

//创建nft并且上链
func CreationMine(appid string, secrect string, address string, name string, order_sn string, url string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["address"] = address
	parameter["name"] = strconv.QuoteToASCII(name)
	parameter["order_sn"] = strconv.QuoteToASCII(order_sn)
	parameter["url"] = url
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	fmt.Println(parameter)
	res := util.PostJsonRequest("creationMine", parameter)
	return res
}

//获取nft实际持有人
func GetOwnerof(appid string, secrect string, token_id string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["token_id"] = token_id
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	res := util.PostJsonRequest("getOwnerof", parameter)
	return res
}

//获取nft的url
func GetTokenUrl(appid string, secrect string, token_id string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["token_id"] = token_id
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	res := util.PostJsonRequest("getTokenUrl", parameter)
	return res
}

//查询某个账户拥有的nft个数
func GetBalanceof(appid string, secrect string, address string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["address"] = address
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	res := util.PostJsonRequest("getBalanceof", parameter)
	return res
}

//查询哈希存证，获取上链的存证信息
func Query(appid string, secrect string, token_id string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["token_id"] = token_id
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	res := util.PostJsonRequest("query", parameter)
	return res
}

//转移nft到另一个账户
func TransFrom(appid string, secrect string, address_from string, address_to string, token_id string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["address_from"] = address_from
	parameter["address_to"] = address_to
	parameter["token_id"] = token_id
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	res := util.PostJsonRequest("transFrom", parameter)
	return res
}

//获取某种类型的nft的基本信息
func GetTokenBasic(appid string, secrect string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	res := util.PostJsonRequest("getTokenBasic", parameter)
	return res
}

//获取某种类型的nft的名称
func GetName(appid string, secrect string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	res := util.PostJsonRequest("getName", parameter)
	return res
}
