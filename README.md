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
+ `v  [command] help` / `v [command] --help` / `v  [command] -h`


## 配置
+ `v config --lang=xx`

## 发布日志

## 命令列表

### 计算
#### 进制转换
`v hex 123` 将输出16/10/8/2 进制的数据，参数如下
1. base： 给定数据的进制，默认为10进制，如将16进制数据转换 `v hex FF --base=16`


### 安全
+ 编码 `v encode [content] [--type=base64]`
+ 解码 `v decode [content] [--type=base64]`
+ 密码 `v password [--count=5] [--len=10] [--numshow=true]`

### 工具
+ JSON

+ 随机数字

+ 时间 `v time [--after=4h]`


### 网络

+ 查询IP `v ip [--domain=www.baidu.com]`
+ 静态资源服务 `v web [--host=:3364] [--path=.]`
>  1. host: 开启的款口和主机地址，比如 `v web --host=:1234` 将使服务启动在`3364`端口
>  2. path：静态资源地址，默认值是当前路径，`v web --path=/tmp`

+ 正向代理服务器 `v proxy --host=xxxx [--port=3364] [--debug=false] [--prefix=xx]`
>  1. port: 服务启动端口,默认值`3364`
>  2. host: 必填参数, 目标主机的主机名，如 `v proxy --host=https://www.baidu.com`
>  3. debug： 是否是调试模式, 默认值 false。调试模式将输出更多细节日志
>  4. prefix: 转发前缀,将在请求URI添加的前缀

+ WebSocket `v ws --url=ws://xxxx` 或者
+ DNS `v dns [--refresh=true] [--domain=xxx]` 刷新DNS或者解析DNS
+ Ping `v ping host --count=10` PING 主机