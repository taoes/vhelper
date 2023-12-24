package proxy

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	port    int
	host    string
	debug   bool
	prefix  string
	Command = &cobra.Command{
		Use:   "proxy",
		Short: "在本地设备创建Http代理服务",
		RunE: func(cmd *cobra.Command, args []string) error {
			if strings.Compare(host, "") == 0 {
				return errors.New("启动代理服务失败，目标网站未设置,请使用 --host=xxx 设置host!")
			}
			log.Println(fmt.Sprintf("代理服务运行中... 端口[%v]", port))
			err := http.ListenAndServe(":"+strconv.Itoa(port), http.HandlerFunc(DoHandle))
			return err
		},
	}
)

func DoHandle(writer http.ResponseWriter, request *http.Request) {
	// 创建一个HttpClient用于转发请求
	cli := &http.Client{}
	localHost := request.Host
	url := request.URL.String()
	log.Println(fmt.Sprintf("Proxy: [%s][%s]->[%s]", request.Method, localHost+url, host+url))

	// 读取请求的Body
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println("读取请求体发生错误")
		// 响应状态码
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// 转发的URL
	reqURL := host + prefix + url

	// 创建转发用的请求
	reqProxy, err := http.NewRequest(request.Method, reqURL, strings.NewReader(string(body)))
	if err != nil {
		log.Println("创建转发请求发生错误")
		// 响应状态码
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// 转发请求的 Header
	for k, v := range request.Header {
		reqProxy.Header.Set(k, v[0])
	}

	// 发起请求
	responseProxy, err := cli.Do(reqProxy)
	if err != nil {
		log.Println("转发请求发生错误")
		// 响应状态码
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer responseProxy.Body.Close()

	// 转发响应的 Header
	for k, v := range responseProxy.Header {
		writer.Header().Set(k, v[0])
	}

	// 转发响应的Body数据
	var data []byte

	// 读取转发响应的Body
	data, err = ioutil.ReadAll(responseProxy.Body)
	if err != nil {
		log.Println("读取响应体发生错误")
		// 响应状态码
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// 转发响应的输出数据
	if debug {
		var dataOutput []byte
		// gzip压缩判断
		isGzipped := isGzipped(responseProxy.Header)
		// gzip压缩编码数据
		if isGzipped {
			// 读取后 r.Body 即关闭，无法再次读取
			// 若需要再次读取，需要用读取到的内容再次构建Reader
			resProxyGzippedBody := ioutil.NopCloser(bytes.NewBuffer(data))
			defer resProxyGzippedBody.Close() // 延时关闭

			// gzip Reader
			gr, err := gzip.NewReader(resProxyGzippedBody)
			if err != nil {
				log.Println("创建gzip读取器发生错误")
				// 响应状态码
				writer.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			defer gr.Close()

			// 读取gzip对象内容
			dataOutput, err = ioutil.ReadAll(gr)
			if err != nil {
				log.Println("读取gzip对象内容发生错误")
				// 响应状态码
				writer.WriteHeader(http.StatusServiceUnavailable)
				return
			}
		} else { // 非gzip压缩
			dataOutput = data
		}
		// 打印转发响应的Body数据，查看转发响应的响应数据时需要。
		log.Println(string(dataOutput))
	}

	// response的Body不能多次读取，
	// 上面已经被读取过一次，需要重新生成可读取的Body数据。
	resProxyBody := ioutil.NopCloser(bytes.NewBuffer(data))
	defer resProxyBody.Close() // 延时关闭

	// 响应状态码
	writer.WriteHeader(responseProxy.StatusCode)
	// 复制转发的响应Body到响应Body
	io.Copy(writer, resProxyBody)

}

const headerContentEncoding = "Content-Encoding"
const encodingGzip = "gzip"

func isGzipped(header http.Header) bool {
	if header == nil {
		return false
	}

	contentEncoding := header.Get(headerContentEncoding)
	isGzipped := false
	if strings.Contains(contentEncoding, encodingGzip) {
		isGzipped = true
	}

	return isGzipped
}

func init() {
	Command.Flags().IntVar(&port, "port", 3364, "运行端口")
	Command.Flags().StringVar(&host, "host", "", "目标主机地址,eg: https://www.baidu.com")
	Command.Flags().BoolVar(&debug, "debug", false, "调试模式")
	Command.Flags().StringVar(&prefix, "prefix", "", "转发前缀")
}
