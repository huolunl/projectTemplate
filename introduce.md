api # 接口文档
│
│
assets # 项目使用的其他资源 (图片、CSS、JavaScript 等)。
│
│
cmd  # 一个项目有很多组件，可以把组件 main 函数所在的文件夹统一放在/cmd 目录下,每个组件的目录名应该跟你期望的可执行文件名是一致的。
│
│
configs # 这个目录用来配置文件模板或默认配置。例如，可以在这里存放 confd 或 consul-template 模板文件,这里有一点要注意，配置中不能携带敏感信息，这些敏感信息，我们可以用占位符来替代，例如username: ${CONFIG_USER_USERNAME}
│
│
deployments # 用来存放 Iaas、PaaS 系统和容器编排部署配置和模板（Docker-Compose，Kubernetes/Helm，Mesos，Terraform，Bosh）。
│
│
docs # 文档
├── devel                            # 开发文档，可以提前规划好，英文版文档和中文版文档
│   ├── en-US/                       # 英文版文档，可以根据需要组织文件结构
│   └── zh-CN                        # 中文版文档，可以根据需要组织文件结构
│       └── development.md           # 开发手册，可以说明如何编译、构建、运行项目
├── guide                            # 用户文档
│   ├── en-US/                       # 英文版文档，可以根据需要组织文件结构
│   └── zh-CN                        # 中文版文档，可以根据需要组织文件结构
│       ├── api/                     # API文档
│       ├── best-practice            # 最佳实践，存放一些比较重要的实践文章
│       │   └── authorization.md
│       ├── faq                      # 常见问题
│       │   ├── iam-apiserver
│       │   └── installation
│       ├── installation             # 安装文档
│       │   └── installation.md
│       ├── introduction/            # 产品介绍文档
│       ├── operation-guide          # 操作指南，里面可以根据RESTful资源再划分为更细的子目录，用来存放系统核心/全部功能的操作手册
│       │   ├── policy.md
│       │   ├── secret.md
│       │   └── user.md
│       ├── quickstart               # 快速入门
│       │   └── quickstart.md
│       ├── README.md                # 用户文档入口文件
│       └── sdk                      # SDK文档
│           └── golang.md
└── images                           # 图片存放目录
|    └── 部署架构v1.png
│
│
examples # 存放应用程序或者公共包的示例代码 
│
│
internal #存放私有应用和库代码。如果一些代码，你不希望在其他应用和库中被导入，可以将这部分代码放在/internal 目录下，在引入其它项目 internal 下的包时，Go 语言会在编译时报错
├── apiserver # 该目录中存放真实的应用代码(服务类型的应用)
│   ├── api # HTTP API 接口的具体实现，主要用来做 HTTP 请求的解包、参数校验、业务逻辑处理、返回。注意这里的业务逻辑处理应该是轻量级的，如果业务逻辑比较复杂，代码量比较多，建议放到 /internal/apiserver/service 目录下。该源码文件主要用来串流程。
│   │    └── v1
│   │ 
│   ├── options # 应用的 command flag。
│   │ 
│   ├── config # 根据命令行参数创建应用配置。
│   │ 
│   ├── service # 存放应用复杂业务处理代码。
│   │ 
│   └── store # 一个应用可能要持久化的存储一些数据，这里主要存放跟数据库交互的代码，比如 Create、Update、Delete、Get、List 等。
├── job # 该目录中存放真实的应用代码(job类型的应用)
│ 
├── pkg # 存放项目内可共享，项目外不共享的包。这些包提供了比较基础、通用的功能，例如工具、错误码、用户验证等功能。
│   ├── code # 错误码
│   │ 
│   ├── validation # 一些通用的验证函数。
│   │ 
│   └── middleware # HTTP 处理链。
│
pkg # 该目录中存放可以被外部应用使用的代码库，其他项目可以直接通过 import 导入这里的代码
│
│
vendor # 需要注意的是，如果是一个 Go 库，不要提交 vendor 依赖包。
│
│
scripts 
│   ├── make-rules # 用来存放 makefile 文件，实现 /Makefile 文件中的各个功能。
│   │ 
│   ├── lib # shell 库，用来存放 shell 脚本。一个大型项目中有很多自动化任务，比如发布、更新文档、生成代码等，所以要写很多 shell 脚本，这些 shell 脚本会有一些通用功能，可以抽象成库，存放在/scripts/lib 目录下，比如 logging.sh，util.sh 等
│   │ 
│   └── install # 如果项目支持自动化部署，可以将自动化部署脚本放在此目录下。如果部署脚本简单，也可以直接放在 /scripts 目录下。
test # 用于存放其他外部测试应用和测试数据。
│ 
│
third_party # 外部帮助工具，分支代码或其他第三方应用（例如 Swagger UI）。比如我们 fork 了一个第三方 go 包，并做了一些小的改动，我们可以放在目录 /third_party/forked 下
│
│ 
tools # 存放这个项目的支持工具。这些工具可导入来自 /pkg 和 /internal 目录的代码。
│
│
web # 前端代码存放目录，主要用来存放 Web 静态资源，服务端模板和单页应用（SPAs）。
│
│
build # 这里存放安装包和持续集成相关的文件。这个目录下有 3 个大概率会使用到的目录，在设计目录结构时可以考虑进去
│   ├── package # 存放容器（Docker）、系统（deb, rpm, pkg）的包配置和脚本。
│   │ 
│   ├── ci # 存放 CI（travis，circle，drone）的配置文件和脚本。
│   │ 
│   └── docker # 存放子项目各个组件的 Dockerfile 文件。
.gitignore # git忽略的文件
│ 
│
CONTRIBUTING.md # 描述如何贡献代码
│ 
│
README.md # 项目整体介绍
│ 
│
Makefile # 项目管理
