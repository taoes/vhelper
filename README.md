[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
![GoLang](https://camo.githubusercontent.com/3cc2c758e723cb171cec95cc8535c9642e1322584da8bd6d218ef0390ab49d04/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2d476f4c616e672d3030414444383f7374796c653d666c61742d737175617265266c6f676f3d676f266c6f676f436f6c6f723d7768697465)
![Linux](https://camo.githubusercontent.com/dbe944dadb1ba77b539d3e12cf20e400b90d8097a42e67a9389227d75acba4c4/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2d4c696e75782d4643433632343f7374796c653d666c61742d737175617265266c6f676f3d6c696e7578266c6f676f436f6c6f723d626c61636b)
![Docker](https://camo.githubusercontent.com/204410115a0bb658668e7446bfc6a7eadb6a96a98d81daba65ddaaa541e95f58/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2d446f636b65722d3234393645443f7374796c653d666c61742d737175617265266c6f676f3d646f636b6572266c6f676f436f6c6f723d7768697465)
![GNU Shell](https://camo.githubusercontent.com/d7bf14575b678a9e2a3df9916ea4d66a41e8ad226dc160b4ae07955ff021521e/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2d474e555f426173682d3445414132353f7374796c653d666c61742d737175617265266c6f676f3d676e7562617368266c6f676f436f6c6f723d7768697465)
![HomeBrew](https://camo.githubusercontent.com/5dd3c75d4f830f93385f93af2afb1c8f0789ce91b3fad658a5a890b604e4f5ff/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2d486f6d65427265772d4642423034303f7374796c653d666c61742d737175617265266c6f676f3d686f6d6562726577266c6f676f436f6c6f723d7768697465)
![Markdown](https://camo.githubusercontent.com/b4ffd17afd4f5133a29621bb201dd41f29436c88952d29adab7a96ecbb59cc96/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2d4d61726b646f776e2d3331353241303f7374796c653d666c61742d737175617265266c6f676f3d6d61726b646f776e266c6f676f436f6c6f723d7768697465)
![Git](https://camo.githubusercontent.com/561f3d4fd727fcca82984c91a65eca069ff34a435072158f6947c4ca52370eae/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2d4769742d4630353033323f7374796c653d666c61742d737175617265266c6f676f3d676974266c6f676f436f6c6f723d7768697465)
![Github](https://camo.githubusercontent.com/b620c6ad3a16345749694c16a7c06a101c9c7757179e6072352e4035fa562837/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2d4769746875622d3138313731373f7374796c653d666c61742d737175617265266c6f676f3d676974687562266c6f676f436f6c6f723d7768697465)

中文 | [英文](./README.md)

## 关于
一款转为程序员设计的终端命令行工具集合，基于Go 语言设计，支持Windows, Mac以及Linux  等操作系统,  目前正在开发中, 欢迎您的关注

## 使用帮助
- [x] `v  [command] help` / `v [command] --help` / `v  [command] -h`
```text
为开发者提供的CLI助手,So you Know, love Command Line MORE!

Usage:
   [flags]
   [command]

Available Commands:
  agent       分析Http请求头中的Agent信息
  completion  Generate the autocompletion script for the specified shell
  decode      字符串编码工具
  encode      编码字符串或者文件内容
  help        Help about any command
  hex         提供数据之间的进制转换
  ip          输出当前设备的网络IP信息
  json        JSON格式化
  password    生成随机密码
  proxy       在本地设备创建Http代理服务
  qrcode      生成二维码文件
  random      生成随机数，支持字符串，整数，浮点数
  search      快速打开搜索引擎
  time        输出当前或者未来的时间戳
  url         快速打开网页
  web         在本地设备运行静态资源
  ws          建立WebSocket连接，并进行调试

Flags:
      --author string         作者: 燕归来兮 https://www.zhoutao123.com
  -h, --help                  help for this command
      --lang string           国际化语言
      --searchEngine string   默认搜索引擎

Use " [command] --help" for more information about a command.
```

## 命令列表

### 文件
- [x] 二维码 `v qrcode (content) [--level=1] [--size=256] [--file=/tmp/xx.png]`

### 计算
- [x] 进制转换 `v hex (content) [--base=16]`


### 安全
- [x] 编码内容 `v encode (content) [--type=base64]`
- [x] 解码内容 `v decode (content) [--type=base64]`
- [x] 生成密码 `v password [--count=5] [--len=10] [--numshow=true]`

### 工具
- [x] JSON美化 `v json`
- [x] 随机数字 `v random [--count=4] [--len=10] [--type=string]`
- [x] 查看时间 `v time`
- [x] 复制内容到剪切 `v copy [--file=xxx] [--stdout=true]` 
- [x] 重定向内容到剪切板 `cat xxx | v copy [--stdout=true]`
- [x] 将剪切板内容输出到文件 `v board [--file=xxx] [--stdout=true]`
- [ ] Cron表达式 `v cron '0 * * * * ?'`
- [ ] 文件下载 `v download (fileUrl)`

### 网络
- [x] 查询IP地址 `v ip [--domain=www.baidu.com]`
- [x] 静态资源服务器 `v web [--host=127.0.0.1:3364] [--path=.]`
- [x] 正向代理服务器 `v proxy --host=xxxx [--port=3364] [--debug=false]`
- [ ] WebSocket `v ws --url=ws://xxxx`
- [x] 打开网页 `v url (webUrl)`


## 更新日志
- [ ] 暂无