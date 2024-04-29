# Go语言编程技巧
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-22-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
Go tips from [Phuong Le](https://twitter.com/func25).

翻译后的站点: [Go语言编程技巧](https://colobu.com/gotips/) (自动生成)

**微信讨论群:**

<img src="src/images/wechat.png" width="200px">

**翻译进度:**： 79/79 = 100%

## 翻译规范

1. 创建一个认领issue,认领标题中注明tip号和标题。不允许一人同时认领多个任务。翻译完一项任务后才能认领下一个任务。
2. 任务请搜索 Phuong Le 的Twitter账号，找到对应的tip，然后翻译成中文。(可能有的tip作者已经删除，经咨询，作者反馈内容会有增删，所以如果某个tip不存在，请空一个，寻找下一个tip)
3. fork 本项目，然后在 fork 后的项目中进行翻译认领的一个tip,翻译完成后提交pull request
4. 任务请在一周内完成，未完成的任务将被释放，其他人可以认领。

- 项目自动统计贡献者
- 翻译请保持语句通顺，可以使用AI协助翻译，但是一定要避免生硬和机翻的感觉

> **怎么提交Pull Request? **
>
> 如果大家以前没有提交过 Pull Request,可以探索下。
> 首先点击项目右上角的 Fork 按钮，将项目 Fork 到自己的仓库。
> 在github浏览你fork的项目，你会看到一个 "Create pull request" 按钮，点击它
> 填写相关的内容提交即可。
> 后续等待项目的维护者合并你的PR即可。


## 翻译模板

```markdown
# Tip #76 函数调用的结果回传

>  原始链接：[Golang Tip #76: Result forwarding in function call](https://twitter.com/func25/status/1779128931586850890)
>

当我刚开始使用go语言的时候，我发现有一个概念比较棘手：函数调用结果的回传。

....
```

- 在src下创建翻译文件: `xxx.md`, 内容模板如上
- 图片放入 `src/images` 文件夹中
- 在 `src/SUMMARY.md` 文件中加入你翻译的一项
- 在 `README` 文件中**翻译任务认领表格** 填写你翻译的项目

## 翻译任务认领表格

| tip序号 | 标题 | 译者         |
| --- | --- |------------|
| 1 | Measure the execution time of a function in just one line of code. | smallnest  |
| 2 | Multistage defer | smallnest  |
| 3 | Pre-allocate slices for performance | smallnest  |
| 4 | Parse an Array into a Slice | smallnest  |
| 5 | Method Chaining | smallnest  |
| 6 | Underscore Import | icyfire    |
| 7 | 作者已删除|            |
| 8 | Wrapping Errors | icyfire    |
| 9 | Compile-Time Interface Verification | icyfire    |
| 10 | Avoid Naked Parameters  | smallnest  |
| 11 | Numeric separators | icyfire    |
| 12 | Avoid using math/rand, use crypto/rand for keys instead | icyfire    |
| 13 | Empty slice or, even better, NIL SLICE | icyfire    |
| 14 | Error messages should not be capitalized or end with punctuation | icyfire    |
| 15 | When to use Dot (.) Import and Blank (_) Import? |      qcrao      |
| 16 | Don't Return -1 or nil to Indicate Error. | kagaya85   |
| 17 | Understanding "Return fast, return early" to avoid nested code | icyfire    |
| 18 | Define interfaces in the consumer package, not the producer | syjs10        |
| 19 | Avoid named results unless necessary for documentation. |     smallnest       |
| 20 | Pass values, not pointers | smallnest  |
| 21 | Prefer using a pointer receiver when defining methods |     QingyaFan       |
| 22 | Simplify function signatures with structs or variadic options | zhubiaook  |
| 23 | Skip the 'Get' prefix for getters | HBUzxl     |
| 24 | Avoid repetition in naming |   smallnest         |
| 25 | Prefer 'chan struct{}' over 'chan bool' for signaling between goroutines | justlorain |
| 26 | Explicitly ignore values with blank identifier (_) instead of silently ignoring them |    smallnest       |
| 27 | Filter without any allocation | devin7788 |
| 28 | Converting multiple if-else statements into switch cases |  zzzpppy |
| 29 | Avoid context.Background(), make your goroutines promisable. | stonemax |
| 30 | Keep contexts going with context.WithoutCancel() | smallnest |
| 31 | Loop labels for cleaner breaks and continues | zhubiaook |
| 32 | Scheduling functions after context cancellation with context.AfterFunc |     LinPr,smallnest       |
| 33 | Just... Don’t Panic() |  baxiang |
| 34 | Lead with context, end with options, and always close with an error | lylex |
| 35 | Prefer strconv over fmt for converting to/from string |     jjjjjim      |
| 36 | Naming Unexported Global Variables with an Underscore (_) Prefix |  baxiang  |
| 37 | Using Unexported Empty Struct as Context Key | baxiang |
| 38 | Make your errors clear with fmt.Errorf, don't just leave them bare | smallnest |
| 39 | Avoid defer in loops, or your memory might blow up | devin7788 |
| 40 | Handle errors while using defer to prevent silent failures | smallnest |
| 41 | Sort your fields in your struct from largest to smallest | justlorain |
| 42 | Single Touch Error Handling, Less Noise. | zhubiaook |
| 43 | Gracefully Shut Down Your Application |    LinPr,smallnest       |
| 44 | Intentionally Stop with Must Functions | syjs10    |
| 45 | Always Manage Your Goroutine Lifetime. | stonemax |
| 46 | Avoid using break in switch cases, except when paired with labels | baxiang |
| 47 | Table-driven tests, subtests, and parallel tests | devin7788 |
| 48 | Avoid Global Variables, Especially Mutable Ones. | vcheckzen          |
| 49 | Give the Caller the Right to Make Decisions | vcheckzen          |
| 50 | Make Structs Non-comparable. | smallnest |
| 51 | Avoid using init() |  richzw  |
| 52 | Adjusting GOMAXPROCS for Containerized Env (Kubernetes, Docker, etc.) | baxiang |
| 53 | Enums start from 1 for categorization and 0 for default cases | baxiang  |
| 54 | Only define errors (var Err = errors.New) when it's necessary for your client | lylex |
| 55 | Prevent Struct Unkeyed Literals by Using an Empty Field  | cannian1 |
| 56 | Simplify interfaces and only ask for what you really need  | cannian1 |
| 57 | Golang Tip #57: Flag Enums in Go |  baxiang         |
| 58 | Keep the mutex close to the data it's protecting |  richzw  |
| 59 | If a parameter isn't needed, either drop it or ignore it on purpose | TravisRoad |
| 60 | sync.Once is the best way to do things once | smallnest  |
| 61 | Making a Type with Built-In Locking (sync.Mutex embedding) | richzw     |
| 62 | context.Value is not our friend | hawkinggg  |
| 63 | Avoid time.Sleep(), it's not context-aware and can't be interrupted | richzw     |
| 64 | Make main() clean and testable. |    syjs10        |
| 65 | Returning Pointers Made Easy with Generics | miniLCT    |
| 66 | Simplify Your Error Messages in fmt.Errorf | smallnest  |
| 67 | How to deal with long function signatures | richzw     |
| 68 | Use the deadcode tool to find and remove unused functions | richzw     |
| 69 | Manage multiple goroutines with errgroup | richzw     |
| 70 | Implement a context-aware sleep function | hxzhouh    |
| 71 | sync.Pool, make it typed-safe with generics | QingyaFan  |
| 72 | Case-Insensitive string comparison with strings.EqualFold | syjs10     |
| 73 | Implement String() for enum with the stringer tool | syjs10     |
| 74 | Make time.Duration clear and easy to understand | richzw     |
| 75 | Optimize multiple calls with singleflight | hxzhouh    |
| 76 | Result forwarding in function calls | syjs10     |
| 77 | Buffered channels as semaphores to limit goroutine execution | QingyaFan  |
| 78 | Non-blocking channel send trick | hxzhouh    |
| 79 | If doing something unusual, comment why |    smallnest        |
| 80 |  |            |

## 生成文档

如果你想在本地编译， 请安装[mdbook](https://github.com/rust-lang/mdBook)工具。

### 本地预览
在本地`mdbook serve`可以生成本地网站访问。

http://localhost:3000

### 生成静态网站
`mdbook build` 生成网站到`book`目录下。





## 贡献者 ✨


<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://colobu.com/"><img src="https://avatars.githubusercontent.com/u/865763?v=4?s=100" width="100px;" alt="smallnest"/><br /><sub><b>smallnest</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=smallnest" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/hxzhouh"><img src="https://avatars.githubusercontent.com/u/25883521?v=4?s=100" width="100px;" alt="hxzhouh"/><br /><sub><b>hxzhouh</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=hxzhouh" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/icyfire"><img src="https://avatars.githubusercontent.com/u/1171180?v=4?s=100" width="100px;" alt="icyfire"/><br /><sub><b>icyfire</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=icyfire" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/syjs10"><img src="https://avatars.githubusercontent.com/u/15065304?v=4?s=100" width="100px;" alt="JS"/><br /><sub><b>JS</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=syjs10" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="http://stackoverflow.com/users/3011380/zangw"><img src="https://avatars.githubusercontent.com/u/1590890?v=4?s=100" width="100px;" alt="richzw"/><br /><sub><b>richzw</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=richzw" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/QingyaFan"><img src="https://avatars.githubusercontent.com/u/10420579?v=4?s=100" width="100px;" alt="cheerfun"/><br /><sub><b>cheerfun</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=QingyaFan" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/miniLCT"><img src="https://avatars.githubusercontent.com/u/45364609?v=4?s=100" width="100px;" alt="rkmdsxmds"/><br /><sub><b>rkmdsxmds</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=miniLCT" title="Code">💻</a></td>
    </tr>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/devin7788"><img src="https://avatars.githubusercontent.com/u/39721466?v=4?s=100" width="100px;" alt="devin7788"/><br /><sub><b>devin7788</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=devin7788" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/cannian1"><img src="https://avatars.githubusercontent.com/u/59365095?v=4?s=100" width="100px;" alt="Cannian"/><br /><sub><b>Cannian</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=cannian1" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://blog.lxythan2lxy.cn"><img src="https://avatars.githubusercontent.com/u/46567145?v=4?s=100" width="100px;" alt="LU XIUYUAN"/><br /><sub><b>LU XIUYUAN</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=TravisRoad" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/zzzpppy"><img src="https://avatars.githubusercontent.com/u/49774236?v=4?s=100" width="100px;" alt="zzzpppy"/><br /><sub><b>zzzpppy</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=zzzpppy" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="http://blog.kagaya.fun"><img src="https://avatars.githubusercontent.com/u/28755005?v=4?s=100" width="100px;" alt="Kagaya"/><br /><sub><b>Kagaya</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=kagaya85" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/stonemax"><img src="https://avatars.githubusercontent.com/u/4516172?v=4?s=100" width="100px;" alt="Zhang Jinlong"/><br /><sub><b>Zhang Jinlong</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=stonemax" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/HBUzxl"><img src="https://avatars.githubusercontent.com/u/86550622?v=4?s=100" width="100px;" alt="HBUzxl"/><br /><sub><b>HBUzxl</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=HBUzxl" title="Code">💻</a></td>
    </tr>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="http://yuque.com/baxiang"><img src="https://avatars.githubusercontent.com/u/2994323?v=4?s=100" width="100px;" alt="羊羽"/><br /><sub><b>羊羽</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=baxiang" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/zhubiaook"><img src="https://avatars.githubusercontent.com/u/30363122?v=4?s=100" width="100px;" alt="zhubiaook"/><br /><sub><b>zhubiaook</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=zhubiaook" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://logi.im"><img src="https://avatars.githubusercontent.com/u/18008498?v=4?s=100" width="100px;" alt="LOGI"/><br /><sub><b>LOGI</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=vcheckzen" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/justlorain"><img src="https://avatars.githubusercontent.com/u/87760338?v=4?s=100" width="100px;" alt="Lorain"/><br /><sub><b>Lorain</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=justlorain" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/lylex"><img src="https://avatars.githubusercontent.com/u/4638962?v=4?s=100" width="100px;" alt="lylex"/><br /><sub><b>lylex</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=lylex" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/LinPr"><img src="https://avatars.githubusercontent.com/u/56944601?v=4?s=100" width="100px;" alt="Lin"/><br /><sub><b>Lin</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=LinPr" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="http://qcrao.com"><img src="https://avatars.githubusercontent.com/u/7698088?v=4?s=100" width="100px;" alt="qcrao"/><br /><sub><b>qcrao</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=qcrao" title="Code">💻</a></td>
    </tr>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://ryan961.github.io"><img src="https://avatars.githubusercontent.com/u/32825698?v=4?s=100" width="100px;" alt="ryan961"/><br /><sub><b>ryan961</b></sub></a><br /><a href="https://github.com/smallnest/gotips/commits?author=ryan961" title="Code">💻</a></td>
    </tr>
  </tbody>
  <tfoot>
    <tr>
      <td align="center" size="13px" colspan="7">
        <img src="https://raw.githubusercontent.com/all-contributors/all-contributors-cli/1b8533af435da9854653492b1327a23a4dbd0a10/assets/logo-small.svg">
          <a href="https://all-contributors.js.org/docs/en/bot/usage">Add your contributions</a>
        </img>
      </td>
    </tr>
  </tfoot>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

