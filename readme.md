# Cached-Translate

Cached-Translate 是一款使用 Go 语言编写的全语种翻译器。它内置有缓存功能，以提高翻译的效率和速度。

## 功能

- 使用 [Google Translate API](https://github.com/bregydoc/gtranslate) 进行翻译
- 使用 SQLite 数据库缓存翻译结果，避免重复翻译
- 支持从 HTTP 请求中获取用户的首选语言进行翻译

## 使用

以下是一个简单的使用例子：

```go
text := "Hello World"
translated := translate.I18n(text, "en", "ja")
fmt.Println("translated:", translated)
```

您也可以从 HTTP 请求中获取用户的首选语言进行翻译：

```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
    text := "Hello World"
    translated := translate.QuickI18nFromRequest(text, r)
    fmt.Fprintln(w, "translated:", translated)
}
```

## 安装

要安装此库，请运行以下命令：

```shell
go get github.com/atmshang/cached-translate
```

## 测试

要运行测试，请运行以下命令：

```shell
go test ./...
```

## 依赖

此项目使用以下库：

- [Google Translate API](https://github.com/bregydoc/gtranslate)
- [gorm](https://gorm.io/gorm)
- [SQLite driver for gorm by Gleb Ares](https://github.com/glebarez/sqlite)
- [plog](https://github.com/atmshang/plog)

## 许可证

此项目使用 MIT 许可证。有关详细信息，请参阅 [LICENSE](LICENSE)。

## 贡献

欢迎任何形式的贡献。如果您发现错误，或者有任何改进建议，请提交 issue 或 pull request。# Cached-Translate

Cached-Translate 是一款使用 Go 语言编写的全语种翻译器。它内置有缓存功能，以提高翻译的效率和速度。

## 功能

- 使用 [Google Translate API](https://github.com/bregydoc/gtranslate) 进行翻译
- 使用 SQLite 数据库缓存翻译结果，避免重复翻译
- 支持从 HTTP 请求中获取用户的首选语言进行翻译

## 使用

以下是一个简单的使用例子：

```go
text := "Hello World"
translated := translate.I18n(text, "en", "ja")
fmt.Println("translated:", translated)
```

您也可以从 HTTP 请求中获取用户的首选语言进行翻译：

```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
    text := "Hello World"
    translated := translate.QuickI18nFromRequest(text, r)
    fmt.Fprintln(w, "translated:", translated)
}
```

## 安装

要安装此库，请运行以下命令：

```shell
go get github.com/atmshang/cached-translate
```

## 测试

要运行测试，请运行以下命令：

```shell
go test ./...
```

## 依赖

此项目使用以下库：

- [Google Translate API](https://github.com/bregydoc/gtranslate)
- [gorm](https://gorm.io/gorm)
- [SQLite driver for gorm by Gleb Ares](https://github.com/glebarez/sqlite)
- [plog](https://github.com/atmshang/plog)

## 许可证

此项目使用 MIT 许可证。有关详细信息，请参阅 [LICENSE](LICENSE)。

## 贡献

欢迎任何形式的贡献。如果您发现错误，或者有任何改进建议，请提交 issue 或 pull request。