# 1. 框架选型

## 1.1 Koa

截至到报告书写，Express 框架在 Node 生态中仍保持极大的占有量，在其之后的依次为 Koa、Sails等。

Express 由于历史沉淀，在社区中拥有众多的插件可供使用，遇到问题会有更多的解决方案。但是作为其后来者 Koa 拥有着更新的特性，使用 async 函数代替了 Express 令人诟病的 Callback，解决了异步组合和异步异常捕获的问题。其次，Koa 相比于 Express 做了精简，移除了路由以及模板渲染等功能。

此外 Koa 的中间件机制与 Express 上也有着很大的区别，Express 的中间件执行顺序是直白的线性流程，但是 Koa 的中间件会按照洋葱模型来执行，这就为开发者提供了更多的操作空间，但同时其也会造成很多不必要的麻烦，如在异步情况下，其执行顺序会变得更为复杂，导致后方的中间件执行出错，如果在设计时能注意到这些问题，那么 Koa 会带来更完善的功能。

![image.png](https://i.loli.net/2019/09/29/LWr7Bnox65Gs9jm.png)

相比之下，Express 虽然拥有更多的历史沉淀，但是也存在着设计落伍的问题，Koa 作为更新一代的框架，在设计时的目标就是取代 Express，虽然功能并没有如 Express 完善，但是却拥有更强的可扩展性，所以在测试中采用了Koa 框架作为 serve 端。

## 1.2 BeeGo

BeeGo 作为排名靠前的国产 GoWeb 框架，在社区上的评价褒贬不一，大多负面评价都是由于其框架的笨重以及 ORM 的糟糕。但是 BeeGo 拥有着完善的中文文档，标准的 RESTful API 风格支持，同时功能又更加完善，设计风格清晰，已经能够满足大多的应用需求，所以在测试中使用了 BeeGo 框架作为 Go Lang 端的解决方案。

BeeGo 遵循标准的 MVC 设计模式，有着完整的路由控制方案，同时也包含了一套 ORM 系统，在 View 层上也包含渲染模板的方案（但是在使用 bee 创建 api 风格的工程结构时会主动删除 View ）。此外，Beego 提供了 8 大模块，将 Go 原生的 http 等模块进行了封装，可以拥有更好的使用体验。

![image.png](https://i.loli.net/2019/09/29/qLbksKn81yCQGDl.png)

BeeGo 的 8 大模块每个模块相互独立，可以在 Controller 层中调用，同时在 Router 与 Controller 之间、Controller 与 Output 之间存在着过滤器机制，利用这一机制可以进行数据传入的过滤（如权限控制），以及输出数据的过滤，组中数据将返回顶层向外输出，其总体流程如下：

![image.png](https://i.loli.net/2019/09/29/1m9XFK6z3Y7QEcP.png)

# 2. 框架应用

Koa 与 BeeGo 都属于 MVC 框架，但是在具体的使用上拥有一些区别，具体如下：

## 2.1 风格

#### Koa 工程目录

```sh
├─app
│  ├─controller
│  ├─middleware
│  └─utils
├─database
├─routes
│  └─api
└─static
    ├─css
    ├─fonts
    ├─img
    └─js
```

Koa 的目录可以由用户自行定制，通过 Koa-Router 作为路由层，执行 Controller 层的逻辑，同时可以选择任意一种 ORM/ODM 框架，提供了高度可制定化空间。同时由于 Koa 的中间件特性，我们可以定制化一些中间件，在其外部 app.js 入口文件中使用。

#### Bee Go 工程目录

```sh
apiproject
├── conf
│   └── app.conf
├── controllers
│   └── object.go
│   └── user.go
├── docs
│   └── doc.go
├── main.go
├── models
│   └── object.go
│   └── user.go
├── routers
│   └── router.go
└── tests
    └── default_test.go             
```

BeeGo 的工程目录则是完全被定制好的，Model层提供了完整的 ORM 支持，但是缺点就如果用户需要使用别的 ORM 系统则需要进行手动调整，不如 Koa 的灵活方便，但是其还是有着 8 大模块的优势，可以完美的跟工程项目结合，不需要额外的解决方案，官方对每个模块的工具使用都有详细的说明。但是其相关插件则较少，如果想要实现额外的功能需要用户去手动调整，好在 BeeGo 的设计较为清晰和简单，可以较为轻松的实现自定义功能。

## 2.2 路由

测试中将后台作为一个 RESTful API 服务器，路由则是必不可少的。

#### 在 Koa 中使用 koa-router

在 Koa 中使用路由实质上是通过注册中间件获取上下文对象，进而再通过判断 `ctx.path` 与 `ctx.method` 来匹配访问路径以及访问方法的。但是如果使用 `koa-router` 可以更合理的简化这些步骤，其以中间件的形式向 Koa 提供了路由功能，提供了基本的 get、post 等方法的匹配，如下：

```js
router.get('/',async (ctx, next)=>{
    ctx.body="首页"; 
    await next()
})
```

同时在项目中需要尽可能的考虑到路由的分层设置，将路由系统分发到各个文件中再建立关联。可以使用 `router.use()` 或者配合 `router.prefix()` 的方法来建立多个分层路由，从而连接 Controller 层：

```go
// main.js
const test_router = require('./routes/api/test_router');
router.prefix("/api")
router.use('/test', test_router)
app.use(router.routes()).use(router.middleware())
```

```js
// test_router.js
const controller = require('../../app/controller/test_controler')
router.get('/testRouter', controller.testRouter)
module.exports = router.routes()
```

![image.png](https://i.loli.net/2019/09/30/WriHgZNeCBqE2KV.png)

最终我们会以中间件的形式将 Router 注册于 Koa 实例中，由于其本身性质是一个 Koa 中间件，也就是说会按照层级调用的顺序，在匹配 Router 前我们可以路由进行一些预操作。同时 koa-router 本身的 `router.use()` 方法实际也是嵌入中间件，可以更加自由对其进行开发。

#### Bee Go 的路由系统

Bee Go 的路由系统原理是利用了 Go 的执行过程，再引入 package router 时执行其 `init()` 函数，从而注册路由，再通过映射 URL 到 controller，如下是其创建一个简单的映射关系的示例：

```go
beego.Router("/user", &controllers.UserController{})
```

在一个 Controller 内部需要声明一个控制器，在这个控制器中内嵌 `beego.Controller` 然后再去重写其 `Init`、`Prepare`、`Post`、`Get`、`Delete`、`Head` 等方法，当路由进匹配时，就会访问其对应的方法。

Bee Go 提倡用户以 RESTful 的风格去书写 API，每一个 Controller 中控制一项路由的匹配，再去指定其不同的请求方法或者解析不同的参数。但是在针对不同体系的系统时，用户也可以自主选择是否使用这种风格，用户也可以按照嵌套路由的风格编写更为灵活的 API，为路由匹配自定义的函数方法：

```go
beego.Router("/api/list",&RestController{},"*:ListFood")
beego.Router("/api/create",&RestController{},"post:CreateFood")
beego.Router("/api/update",&RestController{},"put:UpdateFood")
beego.Router("/api/delete",&RestController{},"delete:DeleteFood")
```

同时 BeeGo 还提供了自动匹配路由与注解路由，可以让路由与方法的连接自动化，这点是远比 koa-router 更强大的，我们只需要设置底层的路由分层规则，对于最顶层的规则则会由自动路由自行匹配，如下则会自动匹配 `/staticblock` 的 GET 方法：

```go
// @router /staticblock [get]
func (this *CMSController) StaticBlock() {
	// do something
}
```

此外，BeeGo 也提供了路由的命名空间，并且可以自由的多级嵌套，这是为路由分层专门提供的方案，配合自动路由可以编写极少量的路由代码来完成复杂的路由配置。

## 2.3 用户鉴权

#### session-cookie 与 JWT

在用户鉴权上，有两种方案可选，一种是基于 session - cookie 的用户登录方案，还有一种是基于 JWT（JSON Web Token）的解决方案，session 是一种比较主流的解决方案，JWT 则是一个比较新的协议方案，网上对两种方案的选型也有很大的争议，大多推荐用户鉴权使用 Session 机制，因为 Session 在解决用户登录问题上有大量的实践实例和解决方案，针对于 JWT 的主要争论是其安全性问题上。

> 无状态 JWT Tokens 无法被单独地销毁或更新，取决于你如何存储，可能还会导致长度问题、安全隐患。有状态 JWT Tokens 在功能方面与 Session cookies 无异，但缺乏生产环境的验证、经过大量 Review 的实现，以及良好的客户端支持。 
>
> ——摘录自文章 [《Stop using JWT for sessions》](http://cryto.net/~joepie91/blog/2016/06/13/stop-using-jwt-for-sessions/)

但是，如果采用 session 作为登录方案，想要保存用户的登录状态，则需要考虑到其持久化问题，可以使用 Redis 来做为持久化的方案或者使用 MySQL 等数据库方案，但是这样就回增加系统部署的成本，同时在开发中如果使用 Mongodb，Node 端的 `koa-session`（当前主流的 Koa session 中间件工具） 并不能很好的支持将 Mongodb 作为 session 持久化存储媒介，到最后还是需要考虑搭建 Redis 服务作为持久化存储媒介。所以在测试中还是使用了 JWT 作为用户鉴权的方案。

>  备注：如果使用 session 作为登录方案，可以选择如下的中间件使用
>
> - koa-session
> - koa-passport
> - passport-local

在使用 JWT 作为用户鉴权时，解决其安全性问题的根本就是防止 Token 泄露。Token 本身有过期时长的设置，超过期限的 Token 将不会被解析使用，但是由于 Token 是天然无状态的，所以未过期的 Token 都可以被应用后台解析其 payload 并返回内同。那么如果后台派发给用户了多个 Token 就会造成可认证的 Token 泛滥，任何人劫取了用户的任意一个有效 Token 都可以制作一个伪造请求。

为了防止 Token 泛滥，其中一个方案是缩短 Token 的有效时间，同时设立一个 Fresh Token 用于专门刷新 Token，当检测到 Token 过期时，前端再向后台发送 Fresh Token 用于刷新当前的 Token，最后再重新发送请求，整体的流程如下：

![image.png](https://i.loli.net/2019/09/30/O4fKkjITYi9HBWq.png)

但是这样就造成一个问题就是刷新 Token 的过程过于复杂，每一个请求在在服务器端与客户端之间的接收与发送都需要进行预处理（这在 Koa 中是基于洋葱模型的的，在 BeeGo 中是基于过滤器实现的），同时 Fresh Token 本质上还是一个 Token，所以 Fresh Token 也有可能泛滥。

那么为了弥补 Fresh Token 的缺陷，另外一个方案就是去利用数据库去维护有效 Token，这一点与 Session 的机制有些类似，但是实现比较简单灵活。我们创建一个记录用户有效 Token 数据库，在关系型数据库中通常是一个一对多关系，在文档型数据库中则可以使用一个数组，同时规定我们的 web app 中允许多少个有效 Token。当一个新 Token 被创建后，就会替换掉最旧的 Token，这一特性在防止 Token 泛滥的同时还可以允许用户多平台登录。当 Token 即将过期时，客户端还可以主动发送一个请求想服务器申请新的 Token，同时让原有的 Token 失效，让用户的登录时限延长。

![image.png](https://i.loli.net/2019/09/30/B1vmtVh3EC6Z7xP.png)

#### 实现方案

根据标准的 JWT 规范，Token 应存放于浏览器的 localstroge 中，同时存放于每个 HTTP Header 的 Authorization 字段中发送给服务器进行解码认证。在客户端，可以通过 axios 的全局设置来将 Token 设置到 Authorization 字段中，同时还可以集中进行一些拦截操作、预处理操作、错误集中处理操作。

在服务器端，如果使用 Koa，可以使用 `jsonwebtoken` 与 `koa-jwt` 两个插件便捷的处理 Token。其中 `jsonwebtoken` 主要提供了对 Token 的加密，这个可以自由实现，或者对其进行二次封装使用。`koa-jwt` 则是作为一个中间件使用，其可以根据密钥主动将有效的 Token 进行解析，并挂载到 `ctx` 上下文对象上，这是一个很方便的特性，我们可以直接在 Controller 层调用 payload 信息，而不需要自己手动解密提取信息，同时 `koa-jwt` 还提供了。

如果使用 BeeGo，可以引入 `jwt-go` ，其提供了灵活的 token 加密与解密功能，并支持多种加密方案，其特点是支持通过一个 Token 字符串以及一套加密信息返回给用户一个 token 对象，用户可以自主选择如何转化 Token 对象的信息或者使用 asert 转其中 payload 数据的类型，但是有一点要注意的是其本身会将 payload 中任意类型的数字转化为 float64 类型。在测试中，可以将 `jwt-go` 单独封装为一个 utils，再通过 Beego 的过滤器机制来验证每一次的 token 信息是否有效，将 payload 信息解析后放置于一个全局变量中，在 Controller 层应用该全局变量也可以便捷的完成 Token 信息的获取。

Node 端与 Go 端对 Token Payload 信息的提取上由于 javascript 弱类型特性，可以便捷的进行类型转换与使用， Go 解析数据上则是严谨了许多，需要进行各种类型转换才可以使用，一下是截取了两段 js 与 go 获取 payload 中的 uid 的示例，可以明显看出 Go 的语言要冗长很多：

```js
let uid = ctx.state.user.uid
let tid = ctx.state.user.tid

findUserById(uid) // uid 会被解析为 Number 类型直接传入函数使用
```

```go
// 通过 Type Assert对 interface{} 变量进行类型断言
payload_uid := filters.TokenData["uid"].(float64) 
payload_tid := filters.TokenData["tid"].(string)

uid := int(payload_uid) // uid 需要进行从 float64 到 int 类型的转换
userInfo, err := models.GetUserInfo(uid)
```

## 2.4 跨域方案

对于跨域问题，主流的解决方案通常是 jsonp 与 CORS，相比之下如果不考虑旧平台的兼容性问题，在使用 CORS 是较为主流的选择。

在 Koa 中我们可以通过中间件机制，在当请求穿如路有层的逻辑之前，设置与 CORS 请求相关的字段，以下为一个简单的演示，仅设置了部分信息，在生产环境中的配置按需求配置：

```js
app.use(async (ctx, next) => {
  ctx.set("Access-Control-Allow-Origin", "*")
  ctx.set("Access-Control-Allow-Headers", "authorization")
  await next()
})
```

需要注意的是，在发送非简单请求时需要特殊的去处理 OPTIONS 预检请求，该请求会造成 404 错误或者 403 错误，分别是由于没有设置相应信息与跳过身份验证导致的，需要对其进行逻辑判断进行给予相应与跳过身份验证。同时使用 `koa-cors` 可以便捷的解决 404、403 这些问题，以及进行一些快速的设置。

在 Beego 中，添加 CORS 请求信息则是在过滤器阶段进行操作，Beego 提供了 CORS 插件，在过滤器的路有匹配前进行调用，也可以在此阶段设定判断选择性跳过身份认证等：

```go
beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
    AllowAllOrigins:  true,
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
    ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
    AllowCredentials: true,
}))

beego.InsertFilter("*", beego.BeforeRouter, func(context *context.Context) {
    // CORS
    if context.Request.Method == http.MethodOptions {
        context.Output.Status = http.StatusOK
        context.WriteString("ok")
    }else{
        // do some things
    }
}
```

## 2.5 数据库

在数据库的测试阶段实验了两种场景，一种是用户信息与 Token 登录信息的关联，一种是记录 wiki 系统中词条项目信息。同时也测试了关系型数据库 MySQL 与文档行数据局 Mongodb 的对比，以及其 ORM 系统与 ODM 系统设计使用上的区别。

在 Node 端，Mongodb 可以与 Node 很好的结合在一起，同时采用 Mongoose ODM 系统可以将 Document 转化为 Object 形成 MVC 设计中的 Model 层，可以更好的的进行数据库操作。

而在 Go 端，测试所使用的 Beego 自带的 ORM 系统显而易见的不会支持 Mongodb 这种文档型数据库，所以在 Beego 上仍然使用 MySQL 作为测试数据库。

#### 针对于用户信息的场景设计

针对于用户信息，按照之前所阐述的用户鉴权机制，需要将用户的有效 Token 与用户的基本信息进行关联，在传统的关系型数据库中，其之间存在的是一对多关系，总体的设计如下：

![image.png](https://i.loli.net/2019/09/30/IqLOglPJ9bRe3nT.png)

如果换做 Mongodb 的设计方案，由于 Mongodb 支持创建数组类型的数据，可以直接创建数组用于记录 Token 信息，没有必要再另外创建一张 Token 表（在文档型数据库中是 Document）来用于记录 Token 信息，设计示例如下：

```json
{
    "id": ObjectId("5d84952398f7744200231d4f"),
    "name": "tester",
    "identity": "user",
    "password": "$2a$04$7dJOVf1H8EIayld/jC5FauwWN0J1YruV.TG2xUxwsHeuidwxSbxam",
    "tokens": ["1569206269638", "1569291746447"]
}
```

可以看出，对于创建一对多关系时，文档型数据库的结构相比于关系型数据库简化了很多，在数据存放量较小时不需要再去多设计一层关系。

#### 针对 wiki 文档的场景设计

wiki 系统中是一个典型的文档结构，首先按照关系型数据库的思维模式来设计，其存在如下几个关系：

- 一个用户可以创建多个项目（词条），所以用户与项目之间存在一对多关系；
- 一个项目（词条）可以拥有多个贡献者，一个用户也可以参与多个项目，所以用户（贡献者）与项目之间也存在多对多关系；
- 一个项目（词条）可以包含多篇描述文章，项目与文章之间是一对多关系。

![image.png](https://i.loli.net/2019/09/30/t2mBUYpXe5E6P9H.png)

Projects 通过 user_id 来建立一对多关系，通过维护一张 Users 与 Projects 的关系表来创建 Users 与 Projects 的多对多关系，通过 Article 的 project_id 来建立一对多关系。

如果将上面的关系用文档模型创建，则需要去创建一个 Project 文档即可完成对上述关系的描述：

```json
{
    "id": ObjectId("12343424234"),
    "project_name": "test project",
    "creator": ObjectId("123"),
    "contributer": [
        ObjectId("123"),
        ObjectId("456"),
        ObjectId("789")
    ],
    "articles": [
        {
            "title": "title",
            "content": "article_content",
            "tag": ["tag1", "tag2", "tag3"]
            "writer": ObjectId("123")
    	},
		{
            "title": "title2",
            "content": "article_content",
            "tag": ["tag1", "tag2", "tag3"]
            "writer": ObjectId("123")
        }
    ]
}
```

文档型数据结构在 wiki 系统的文档结构中很清晰，没有多复杂的关系连接。Mongodb 中的两个很常用的概念嵌套文档与引用，如上例，文章（Article）与项目（Projects）属于单一归属关系，所以可以很自然的使用嵌套文档将 Articles 作为 Projects 的一部分。然而 Users 与 Articles 和 Projects 之间都不是单一归属关系，一个 Users 可以创建多个 Projects 与 Articles ，所以在创建时为其设置为引用关系，Mongodb 可以通过向上查找的方式找到引用对象的相关信息。

## 2.6 应用部署

#### 打包发布

Beego 提供了打包服务的功能，通过指令可以打包运行至各平台：

```sh
bee pack -be GOOS=linux
bee pack -be GOOS=windows
```

同样的 Node 也有类似的打包服务，可以通过使用pkg打包Node.js应用：

```sh
npm install pkg --save-dev
pkg [options] <input>
```

#### npm发布

同时还有一种方案就是制作一个 cli 工具将其上传 npm，当用户全局安装该 cli 工具之后，通过命令行指令，可以自动拉取应用并完成一些基础配置。cli 可以从工作中总结繁杂、有规律可循、或者简单重复劳动的工作用 cli 来完成，只需一些命令，快速完成简单基础劳动，有利于项目的二次开发。

编写一个 Cli 工具通常需要考虑到用户与命令界面的交互问题，经过整理，开发 cli 的常用工具有以下几个：

- [commander.js](https://github.com/tj/commander.js)，可以自动的解析命令和参数，用于处理用户输入的命令。
- [download-git-repo](https://github.com/flipxfx/download-git-repo)，下载并提取 git 仓库，用于下载项目模板。
- [Inquirer.js](https://github.com/SBoudrias/Inquirer.js)，通用的命令行用户界面集合，用于和用户进行交互。
- [handlebars.js](https://github.com/wycats/handlebars.js)，模板引擎，将用户提交的信息动态填充到文件中。
- [ora](https://github.com/sindresorhus/ora)，下载过程久的话，可以用于显示下载中的动画效果。
- [chalk](https://github.com/chalk/chalk)，可以给终端的字体加上颜色。
- [log-symbols](https://github.com/sindresorhus/log-symbols)，可以在终端上显示出 √ 或 × 等的图标。

# 3. Go 与 Node

经过几天的开发测试明显感受到了 Go 与 Node 各自的优势与缺点，两者都非常适合用来开发 RESTful API，但具体的使用方案需要按照项目的需求而定。

## 3.1 包管理机制

在包管理方面，node 有 npm 的加持是比较有优势的，同时 node 在社区中仍保持着较高的活跃度，有大量的开发工具可以使用，npm 的包管理机制也是比较优秀的。

而 Go 的包管理机制则比较复杂，Go对包管理一定有自己的理解。对于包的获取，就是用go get命令从远程代码库(GitHub, Bitbucket, Google Code, Launchpad)拉取。并且它支持根据import package分析来递归拉取。这样做的好处是，直接跳过了包管理中央库的的约束，让代码的拉取直接基于版本控制库，但同时不能自由的拷贝项目，并且没有依赖的版本控制。

但是在 1.11 版本的 Go 发布之后，官方引入了 Go Mod 作为 Go 的包管理工具，但是需要用户手动开启 Go Mod，开启后项目会改变查找依赖的方式，依赖包的存放位置变更为`$GOPATH/pkg`，允许同一个package多个版本并存，且多个项目可以共享缓存的 module，当用户拉取项目时会根据 `.mod` 文件自动安装依赖。但是要注意的是如果是针对脚手架工具，如 `bee` ，需要手动关闭 go mod 模式，再进行安装，因为开启 go mod 后不会去关联 bin 文件，导致我们无法安装脚手架工具。

## 3.2 语言类型

Node 最明显的劣势就是由于其是一门动态、弱类型的语言，javascript 是在运行时解析的，没有办法在调试的时候发现错误。但是也有好处就是处理数据时非常灵活，可以忽略验证数据类型的繁琐，利用隐式转换可以便捷的完成很多事情。因此 node 非常适合用来搭建 api 服务，与前端的数据流通上完全统一。

go 属于静态语言，也是强类型语言，可以在编译时查找错误，并且对代码的规范相对严格，这都有助于我们编码时对数据处理的完整性。但是缺点就是繁杂的类型转换与大量的 Type Asert。

我们向前端返回 json 格式的数据，虽然 Beego 也提供了较好的封装，但是相比之下仍显笨重：

```go
var data = map[string]interface{}{
    "msg": "ok",
    "data": map[string]interface{}{
        "num": number.(int64),
        "slice": []string{"1","2"},
    },
}
controller.Ctx.Output.JSON(data, false, false)
```

相比之下 js 在传递数据上则有天然的优势：

```js
var data = {
    msg: "ok",
    data: {
        num,
        slice: ["1", "2"]
    }
}
ctx.send(data)
```

## 3.3 并发处理

node只适合IO密集型，它没有提供太多的并发基元。唯一能同时运行的是I/O程序和定时器等，并不适合CPU密集型。如果遇到计算密集型的任务，因为node是单线程，就会阻塞主线程直到该任务执行完毕才会往下执行，所以node不适合做CPU密集型。

> nodejs 的 worker_threads 还处于实验性阶段

Go 适合 IO 密集型同样也适合 CPU 密集型，用户可以在程序运行的任何阶段，创建 goruntine 去实现并发，并且go提供了 channel 来实现协程间通信，这让 go 在并发编程上有很强的原生优势。

## 3.4 错误处理

在异步方法的错误处理上，Node 原始方案是将错误在回调中层层冒泡，但是也可以利用 Promis then catch 来捕捉错误。同时对于 js 的异常捕获，我们需要考虑的往往更多，特别是在异步操作中，如果在任务中引起异常却没有设置捕获异常的方法，会导致整个 node 进程退出。

相对于 Node，Go 提供了更完善的错误处理能力，在 Go 语言的编程风格中，每一个函数方法返回的第二个参数都可以是一个 error 对象，用户可以通过判断 error 对象的返回值是否为空来提前 return 或执行另外的代码语句，同时 error 还可以使用断言来判断错误类型，从而进行错误的判断。同时对于异常（panic），Go 也有相对应处理方案，与 error 不同 painc 会导致代码的提前退出造成系统异常，但是 Go 在执行完 painc 后会执行函数体内的 defer 函数，我们可以在 defer 函数中通过设置 revcover 来对 panic 进行处理防止系统崩溃。

## 3.5 测试

Node 本身是不支持单元测试的，需要引入第三方的测试框架 Mocha 以及第三方的断言库。

而 Go 自身便拥有进行单元测试的能力，还额外提供了性能测试，帮助开发者分析代码消耗性能的部分以优化代码。相比之下 Go 的测试能力是 Node 不及的。

# 4. 附录

## 4.1 node 测试项目使用到的依赖包

```
"art-template": "^4.13.2",
"ejs": "^2.6.1",
"jsonwebtoken": "^8.5.1",
"koa": "^2.7.0",
"koa-art-template": "^1.1.1",
"koa-bodyparser": "^4.2.1",
"koa-jwt": "^3.6.0",
"koa-passport": "^4.1.3",
"koa-router": "^7.4.0",
"koa-session": "^5.12.0",
"koa-static": "^5.0.0",
"mongodb": "^3.3.2",
"mongoose": "^5.7.1",
"passport": "^0.4.0",
"passport-local": "^1.0.0"
```

## 4.2 GoLang 测试项目是用到的依赖包

```
github.com/astaxie/beego v1.12.0
github.com/dgrijalva/jwt-go v3.2.0+incompatible
github.com/go-sql-driver/mysql v1.4.1
github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
golang.org/x/crypto v0.0.0-20190927123631-a832865fa7ad
golang.org/x/net v0.0.0-20190926025831-c00fd9afed17 // indirect
google.golang.org/appengine v1.6.4 // indirect
gopkg.in/yaml.v2 v2.2.2 // indirect
```

