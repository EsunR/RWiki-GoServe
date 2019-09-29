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

## 2.1 框架风格

### 2.1.1 Koa

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

### 2.1.2 BeeGo

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

## 2.2 路由的解决方案

测试中将后台作为一个 RESTful API 服务器，路由则是必不可少的，

## npm发布

创建一个Cli工具，从而实现系统初始化的工作，创建cli系统需要用到的依赖：

- [commander.js](https://github.com/tj/commander.js)，可以自动的解析命令和参数，用于处理用户输入的命令。
- [download-git-repo](https://github.com/flipxfx/download-git-repo)，下载并提取 git 仓库，用于下载项目模板。
- [Inquirer.js](https://github.com/SBoudrias/Inquirer.js)，通用的命令行用户界面集合，用于和用户进行交互。
- [handlebars.js](https://github.com/wycats/handlebars.js)，模板引擎，将用户提交的信息动态填充到文件中。
- [ora](https://github.com/sindresorhus/ora)，下载过程久的话，可以用于显示下载中的动画效果。
- [chalk](https://github.com/chalk/chalk)，可以给终端的字体加上颜色。
- [log-symbols](https://github.com/sindresorhus/log-symbols)，可以在终端上显示出 √ 或 × 等的图标。



