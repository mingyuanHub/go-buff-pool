package main

import (
	"compress/gzip"
	"io/ioutil"
	"net"
	"net/http"
	"bytes"
	"io"
	"errors"
	"fmt"
	"time"
)

func main()  {
	fmt.Println("------------START------------")

	for i := 1; i < 100; i ++ {
		res := post(i)
		if res == 1 {
			break
		}
	}

	fmt.Println("------------END------------")
}

//var apiUrl = "https://openbid.pubmatic.com/translator?pubId=161276" //测试
//var apiUrl = "https://openbid.pubmatic.com/translator?pubId=160692" //正式
//pubmatic video rewarded 激励视频
//var bytesData = ""

//pubmatic video instl =1 插屏视频
//var bytesData = ""

//pubmatic banner instl =1 插屏图片
//var bytesData = ""

//pubmatic banner instl =0 Banner
//var bytesData = "{\"id\": \"1f7cb80e-b994-bba6-69ed-07cf7fa66b1e\", \"app\": {\"id\": \"FDC48B1D9D9E1F5CBD0C327159C8191C\", \"name\": \"tradplusdemo\", \"bundle\": \"com.unstall.meetdeleteapp\", \"ver\": \"2.0\", \"ext\": {\"orientation\": 1 } }, \"test\": 0, \"tmax\": 1000, \"cur\": [\"USD\"], \"at\": 1, \"imp\": [{\"id\": \"1\", \"banner\": {\"w\": 320, \"h\": 50, \"topframe\": 1, \"api\": [3, 5, 6 ], \"id\": \"1\"}, \"displaymanager\": \"TradPlus\", \"displaymanagerver\": \"6.5\", \"tagid\": \"4112977\", \"bidfloorcur\": \"USD\", \"secure\": 1, \"exp\": 10800, \"ext\": {}, \"iframebuster\": 0 } ], \"device\": {\"ua\": \"Mozilla/5.0 (Linux; Android 7.1.1; vivo Y75A Build/N6F26Q; wv)AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 MobileSafari/537.36\", \"geo\": {\"lat\": 55.9197, \"lon\": 37.8281, \"type\": 2, \"country\": \"RUS\", \"ipservice\": 1 }, \"ip\": \"59.144.134.146\", \"devicetype\": 1, \"make\": \"google\", \"model\": \"Nexus 6\", \"os\": \"Android\", \"osv\": \"7.1.1\", \"hwv\": \"shamu\", \"h\": 2392, \"w\": 1440, \"ppi\": 5, \"pxratio\": 3.5, \"js\": 1, \"connectiontype\": 2, \"ifa\": \"9d8fe0a9-c0dd-4482-b16b-5709b00c608d\", \"ext\": {} }, \"user\": {\"id\": \"UID-edf3508b-494d-4394-bcc5-9583db5f7b55\", \"ext\": {\"gdpr\": 0, \"consent\": \"1\"} }, \"regs\": {\"coppa\": 0, \"ext\": {\"gdpr\": 0 } }, \"ext\": {} }"

//pubmatic video 贴片
//var bytesData = "{\"id\":\"2a21e3b5-bb2d-43e1-1e98-bcc13230914f\",\"imp\":[{\"id\":\"1\",\"video\":{\"ext\":{\"rewarded\":\"0\"},\"linearity\":1,\"placement\":1,\"companiontype\":[1,2],\"h\":640,\"skip\":1,\"skipmin\":1,\"minduration\":3,\"mimes\":[\"video/MP4\",\"video/AVI\"],\"maxduration\":120,\"w\":320,\"startdelay\":5,\"api\":[1,2,3,4,5,6],\"protocols\":[1,2,3,4,5,6,7,8,9,10]},\"instl\":0,\"tagid\":\"3943799\",\"bidfloor\":0.01,\"bidfloorcur\":\"\",\"clickbrowser\":0,\"secure\":1,\"exp\":0,\"iframebuster\":0,\"metric\":null}],\"app\":{\"id\":\"119102\",\"bundle\":\"pampam.ibf2\",\"domain\":\"tastypill.com\",\"storeurl\":\"https://play.google.com/store/apps/details?id=pampam.ibf2\",\"cat\":[\"IAB1\"]},\"device\":{\"ua\":\"Mozilla/5.0 (Linux; Android 9; vivo 1902 Build/PPR1.180610.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36\",\"geo\":{\"country\":\"IDN\",\"city\":\"bandung\"},\"ip\":\"203.78.112.246\",\"devicetype\":4,\"make\":\"vivo\",\"model\":\"1902\",\"os\":\"android\",\"osv\":\"9.0.0\",\"connectiontype\":4,\"ifa\":\"07745767-e48a-4b2f-bbc8-2e043760310c\"},\"test\":0,\"at\":1,\"regs\":{}}"



//var apiUrl = "http://api.w.inmobi.com/ortb"

//inmobi banner 【成功】
//var bytesData = "{\"id\": \"f582b483-7970-406f-d86e-ea538c040791\", \"app\": {\"id\": \"FDC48B1D9D9E1F5CBD0C327159C8191C\", \"name\": \"tradplusdemo\", \"bundle\": \"com.unstall.meetdeleteapp\", \"ver\": \"2.0\", \"ext\": {\"orientation\": 1 } }, \"test\": 0, \"tmax\": 1000, \"cur\": [\"USD\"], \"regs\": {\"ext\": {} }, \"at\": 1, \"ext\": {}, \"imp\": [{\"id\": \"1\", \"banner\": {\"w\": 320, \"h\": 50, \"topframe\": 1, \"api\": [3, 5, 6 ], \"id\": \"1\"}, \"displaymanager\": \"TradPlus\", \"displaymanagerver\": \"6.5\", \"tagid\": \"1634743063620\", \"bidfloorcur\": \"USD\", \"secure\": 1, \"exp\": 10800, \"ext\": {\"placementid\": \"1634743063620\"} } ], \"device\": {\"ua\": \"Mozilla/5.0 (Linux; Android 7.1.1; vivo Y75A Build/N6F26Q; wv)AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 MobileSafari/537.36\", \"geo\": {\"lat\": 55.9197, \"lon\": 37.8281, \"type\": 2, \"country\": \"RUS\", \"ipservice\": 1 }, \"ip\": \"59.144.134.146\", \"devicetype\": 1, \"make\": \"google\", \"model\": \"Nexus 6\", \"os\": \"Android\", \"osv\": \"7.1.1\", \"hwv\": \"shamu\", \"h\": 2392, \"w\": 1440, \"ppi\": 5, \"pxratio\": 3.5, \"js\": 1, \"connectiontype\": 2, \"ifa\": \"9d8fe0a9-c0dd-4482-b16b-5709b00c608d\", \"ext\": {} }, \"user\": {\"id\": \"UID-edf3508b-494d-4394-bcc5-9583db5f7b55\", \"ext\": {\"consent\": \"1\"} } }"

//inmobi reward video 【成功】
//var bytesData = "{\"id\": \"b2e2e4f2-7204-a87a-2f5e-c3abdf170423\", \"app\": {\"id\": \"FDC48B1D9D9E1F5CBD0C327159C8191C\", \"name\": \"tradplusdemo\", \"bundle\": \"com.unstall.meetdeleteapp\", \"ver\": \"2.0\", \"ext\": {\"orientation\": 1 } }, \"test\": 0, \"tmax\": 1000, \"cur\": [\"USD\"], \"regs\": {\"ext\": {} }, \"at\": 1, \"ext\": {}, \"imp\": [{\"id\": \"1\", \"video\": {\"mimes\": [\"video/mp4\"], \"minduration\": 3, \"maxDuration\": 50, \"protocols\": [2, 3 ], \"w\": 2392, \"h\": 1440, \"placement\": 5, \"skip\": 1, \"skipmin\": 10, \"skipafter\": 3, \"boxingallowed\": 1, \"playbackmethod\": [1 ], \"playBackend\": 1, \"delivery\": [3 ], \"pos\": 7, \"api\": [3, 5, 6 ], \"companiontype\": [1 ], \"ext\": {\"rewarded\": 1 } }, \"displaymanager\": \"TradPlus\", \"displaymanagerver\": \"6.5\", \"tagid\": \"1638014035526\", \"bidfloorcur\": \"USD\", \"secure\": 1, \"inst\":1, \"exp\": 10800, \"ext\": {\"placementid\": \"1638014035526\"} } ], \"device\": {\"ua\": \"Mozilla/5.0 (Linux; Android 7.1.1; vivo Y75A Build/N6F26Q; wv)AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 MobileSafari/537.36\", \"geo\": {\"lat\": 55.9197, \"lon\": 37.8281, \"type\": 2, \"country\": \"RUS\", \"ipservice\": 1 }, \"ip\": \"59.144.134.146\", \"devicetype\": 1, \"make\": \"google\", \"model\": \"Nexus 6\", \"os\": \"Android\", \"osv\": \"7.1.1\", \"hwv\": \"shamu\", \"h\": 2392, \"w\": 1440, \"ppi\": 5, \"pxratio\": 3.5, \"js\": 1, \"connectiontype\": 2, \"ifa\": \"2da994ba-0e4c-4618-a96a-10ca1ad1abe1\", \"ext\": {} }, \"user\": {\"id\": \"UID-edf3508b-494d-4394-bcc5-9583db5f7b55\", \"ext\": {\"consent\": \"1\"} } }"

//inmobi native 【】
//var bytesData = "{\"id\": \"a3e8ebec-37df-a047-0f0e-3c5a829abee5\", \"app\": {\"id\": \"FDC48B1D9D9E1F5CBD0C327159C8191C\", \"name\": \"tradplusdemo\", \"bundle\": \"com.unstall.meetdeleteapp\", \"ver\": \"2.0\", \"ext\": {\"orientation\": 1 } }, \"test\": 0, \"tmax\": 1000, \"cur\": [\"USD\"], \"regs\": {\"ext\": {} }, \"at\": 1, \"ext\": {}, \"imp\": [{\"id\": \"1\", \"native\": {\"request\": \"{\\\"native\\\":{\\\"plcmtcnt\\\":1,\\\"ver\\\":\\\"1\\\",\\\"assets\\\":[{\\\"img\\\":{\\\"w\\\":\\\"720\\\",\\\"h\\\":\\\"1280\\\",\\\"hmin\\\":\\\"200\\\",\\\"type\\\":\\\"2\\\",\\\"wmin\\\":\\\"382\\\"},\\\"id\\\":123,\\\"video\\\":{\\\"maxduration\\\":15,\\\"protocols\\\":[2,3],\\\"minduration\\\":5,\\\"mimes\\\":[\\\"video/mp4”\\\",\\\"video/xflv\\\"]},\\\"title\\\":{\\\"len\\\":140},\\\"required\\\":0}],\\\"plcmttype\\\":11,\\\"privacy\\\":\\\"1\\\",\\\"eventtrackers\\\":[],\\\"seq\\\":0}}\", \"ver\": \"1\", \"api\": [3, 5, 6 ] }, \"displaymanager\": \"TradPlus\", \"displaymanagerver\": \"6.5\", \"tagid\": \"1637517570002\", \"bidfloorcur\": \"USD\", \"secure\": 1, \"exp\": 10800, \"ext\": {\"placementid\": \"1637517570002\"} } ], \"device\": {\"ua\": \"Mozilla/5.0 (Linux; Android 7.1.1; vivo Y75A Build/N6F26Q; wv)AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 MobileSafari/537.36\", \"geo\": {\"lat\": 55.9197, \"lon\": 37.8281, \"type\": 2, \"country\": \"RUS\", \"ipservice\": 1 }, \"ip\": \"59.144.134.146\", \"devicetype\": 1, \"make\": \"google\", \"model\": \"Nexus 6\", \"os\": \"Android\", \"osv\": \"7.1.1\", \"hwv\": \"shamu\", \"h\": 2392, \"w\": 1440, \"ppi\": 5, \"pxratio\": 3.5, \"js\": 1, \"connectiontype\": 2, \"ifa\": \"9d8fe0a9-c0dd-4482-b16b-5709b00c608d\", \"ext\": {} }, \"user\": {\"id\": \"UID-edf3508b-494d-4394-bcc5-9583db5f7b55\", \"ext\": {\"consent\": \"1\"} } }"

//inmobi banner instl=1 【成功】
//var bytesData = "{\"id\": \"fcadeae6-3374-e2e0-8021-fedacb01fdb6\", \"app\": {\"id\": \"FDC48B1D9D9E1F5CBD0C327159C8191C\", \"name\": \"tradplusdemo\", \"bundle\": \"com.unstall.meetdeleteapp\", \"ver\": \"2.0\", \"ext\": {\"orientation\": 1 } }, \"test\": 0, \"tmax\": 1000, \"cur\": [\"USD\"], \"regs\": {\"ext\": {} }, \"at\": 1, \"ext\": {}, \"imp\": [{\"id\": \"1\", \"banner\": {\"topframe\": 1, \"api\": [3, 5, 6 ], \"id\": \"1\", \"w\":\"1440\", \"h\":\"2392\"}, \"displaymanager\": \"TradPlus\", \"displaymanagerver\": \"6.5\", \"instl\": 1, \"tagid\": \"1637932216721\", \"bidfloorcur\": \"USD\", \"secure\": 1, \"exp\": 10800, \"ext\": {\"placementid\": \"1637932216721\"} } ], \"device\": {\"ua\": \"Mozilla/5.0 (Linux; Android 7.1.1; vivo Y75A Build/N6F26Q; wv)AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 MobileSafari/537.36\", \"geo\": {\"lat\": 55.9197, \"lon\": 37.8281, \"type\": 2, \"country\": \"RUS\", \"ipservice\": 1 }, \"ip\": \"59.144.134.146\", \"devicetype\": 1, \"make\": \"google\", \"model\": \"Nexus 6\", \"os\": \"Android\", \"osv\": \"7.1.1\", \"hwv\": \"shamu\", \"h\": 2392, \"w\": 1440, \"ppi\": 5, \"pxratio\": 3.5, \"js\": 1, \"connectiontype\": 2, \"ifa\": \"9d8fe0a9-c0dd-4482-b16b-5709b00c608d\", \"ext\": {} }, \"user\": {\"id\": \"UID-edf3508b-494d-4394-bcc5-9583db5f7b55\", \"ext\": {\"consent\": \"1\"} } }"


//inmobi reward instl=1 【成功】
//var bytesData = "{\"id\": \"4e4d1b2f-f12c-6a73-0411-ce00e401fb6e\", \"app\": {\"id\": \"FDC48B1D9D9E1F5CBD0C327159C8191C\", \"name\": \"tradplusdemo\", \"bundle\": \"com.unstall.meetdeleteapp\", \"ver\": \"2.0\", \"ext\": {\"orientation\": 1 } }, \"test\": 0, \"tmax\": 1000, \"cur\": [\"USD\"], \"regs\": {\"ext\": {} }, \"at\": 1, \"ext\": {}, \"imp\": [{\"id\": \"1\", \"video\": {\"mimes\": [\"video/mp4\"], \"minduration\": 3, \"maxDuration\": 50, \"protocols\": [2, 3 ], \"w\": 2392, \"h\": 1440, \"placement\": 5, \"skip\": 1, \"skipmin\": 10, \"skipafter\": 3, \"boxingallowed\": 1, \"playbackmethod\": [1 ], \"playbackend\": 1, \"delivery\": [3 ], \"pos\": 7, \"api\": [3, 5, 6 ], \"companiontype\": [1 ], \"ext\": {\"rewarded\": 2 } }, \"displaymanager\": \"TradPlus\", \"displaymanagerver\": \"6.5\", \"instl\": 1, \"tagid\": \"1638014035526\", \"bidfloorcur\": \"USD\", \"secure\": 1, \"exp\": 10800, \"ext\": {\"placementid\": \"1638014035526\"} } ], \"device\": {\"ua\": \"Mozilla/5.0 (Linux; Android 7.1.1; vivo Y75A Build/N6F26Q; wv)AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 MobileSafari/537.36\", \"geo\": {\"lat\": 55.9197, \"lon\": 37.8281, \"type\": 2, \"country\": \"RUS\", \"ipservice\": 1 }, \"ip\": \"59.144.134.146\", \"devicetype\": 1, \"make\": \"google\", \"model\": \"Nexus 6\", \"os\": \"Android\", \"osv\": \"7.1.1\", \"hwv\": \"shamu\", \"h\": 2392, \"w\": 1440, \"ppi\": 5, \"pxratio\": 3.5, \"js\": 1, \"connectiontype\": 2, \"ifa\": \"9d8fe0a9-c0dd-4482-b16b-5709b00c608d\", \"ext\": {} }, \"user\": {\"id\": \"UID-edf3508b-494d-4394-bcc5-9583db5f7b55\", \"ext\": {\"consent\": \"1\"} } }"


//inmobi reward instl=1 【】 测试inmobi激励视频广告位
//var bytesData = "{\"id\": \"4e4d1b2f-f12c-6a73-0411-ce00e401fb6e\", \"app\": {\"id\": \"FDC48B1D9D9E1F5CBD0C327159C8191C\", \"name\": \"tradplusdemo\", \"bundle\": \"com.unstall.meetdeleteapp\", \"ver\": \"2.0\", \"ext\": {\"orientation\": 1 } }, \"test\": 0, \"tmax\": 1000, \"cur\": [\"USD\"], \"regs\": {\"ext\": {} }, \"at\": 1, \"ext\": {}, \"imp\": [{\"id\": \"1\", \"video\": {\"mimes\": [\"video/mp4\"], \"minduration\": 3, \"maxDuration\": 50, \"protocols\": [2, 3 ], \"w\": 2392, \"h\": 1440, \"placement\": 5, \"skip\": 1, \"skipmin\": 10, \"skipafter\": 3, \"boxingallowed\": 1, \"playbackmethod\": [1 ], \"playbackend\": 1, \"delivery\": [3 ], \"pos\": 7, \"api\": [3, 5, 6 ], \"companiontype\": [1 ], \"ext\": {\"rewarded\": 1 } }, \"displaymanager\": \"TradPlus\", \"displaymanagerver\": \"6.5\", \"instl\": 0, \"tagid\": \"1635440803917\", \"bidfloorcur\": \"USD\", \"secure\": 1, \"exp\": 10800, \"ext\": {\"placementid\": \"1635440803917\"} } ], \"device\": {\"ua\": \"Mozilla/5.0 (Linux; Android 7.1.1; vivo Y75A Build/N6F26Q; wv)AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 MobileSafari/537.36\", \"geo\": {\"lat\": 55.9197, \"lon\": 37.8281, \"type\": 2, \"country\": \"RUS\", \"ipservice\": 1 }, \"ip\": \"59.144.134.146\", \"devicetype\": 1, \"make\": \"google\", \"model\": \"Nexus 6\", \"os\": \"Android\", \"osv\": \"7.1.1\", \"hwv\": \"shamu\", \"h\": 2392, \"w\": 1440, \"ppi\": 5, \"pxratio\": 3.5, \"js\": 1, \"connectiontype\": 2, \"ifa\": \"9d8fe0a9-c0dd-4482-b16b-5709b00c608d\", \"ext\": {} }, \"user\": {\"id\": \"UID-edf3508b-494d-4394-bcc5-9583db5f7b55\", \"ext\": {\"consent\": \"1\"} } }"


func post(i int) int {

	var header = make(map[string]string)

	var httpClient2000 = createHTTPClient(2000)

	response, body, err := HttpPostRequest(apiUrl, []byte(bytesData), header, httpClient2000)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	fmt.Println(i, response.StatusCode)

	if response.StatusCode == 200 {
		fmt.Println(string(body))
		return 1
	}

	return 0
}


func createHTTPClient(requestTimeout int) *http.Client {
	transport := &http.Transport{
		MaxIdleConns: 200,
		MaxIdleConnsPerHost: 1000,
		IdleConnTimeout:       300 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
		Timeout: time.Duration(requestTimeout) * time.Millisecond,
	}
	return client
}

func HttpPostRequest(apiUrl string, bytesData []byte, headers map[string]string, httpClient *http.Client)  (*http.Response, []byte, error){
	var err error

	var isGzip = false
	if len(headers) > 0 {
		for key, item := range headers {
			if key == "Accept-Encoding" && item == "gzip" {
				isGzip = true
				break
			}
		}
	}

	var reader *bytes.Buffer

	if isGzip {
		var zBuf bytes.Buffer
		zw := gzip.NewWriter(&zBuf)
		if _, err = zw.Write(bytesData); err != nil {
			return nil, []byte{}, errors.New(fmt.Sprintf("gzip error='%s'", err))
		}
		zw.Close()
		reader = &zBuf
	} else {
		reader = bytes.NewBuffer(bytesData)

	}

	request, err := http.NewRequest("POST", apiUrl, reader)

	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("newRequest error='%s'", err))
	}

	if len(headers) > 0 {
		for key, item := range headers {
			request.Header.Set(key, item)
		}
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("clientDo error='%s'", err))
	}

	defer response.Body.Close()

	body := response.Body

	//if response.Header.Get("Content-Encoding") == "gzip" {
	//	body, err = gzip.NewReader(response.Body)
	//	if err != nil {
	//		return nil, []byte{}, errors.New(fmt.Sprintf("unzip error='%s'", err))
	//	}
	//}

	data, err := ioutil.ReadAll(body)
	io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("ioutilReadAll error='%s'", err))
	}

	return response, data, nil
}
