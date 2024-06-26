# HTMX-Golang Projesi


![Go Version](https://img.shields.io/github/go-mod/go-version/emrekndl/htmx-golang)
![HTMX Version](https://img.shields.io/badge/htmx-2.0.0-blue)

Bu proje, Go dilini kullanarak `htmx` ile dinamik bir web uygulaması oluşturmayı göstermektedir. Go tarafında `echo` framework'ü ile sunucu sağlanmaktadır.

## İçindekiler
- [Kurulum](#kurulum)
- [Kullanım](#kullanım)
- [Proje Yapısı](#proje-yapısı)
- [HTMX Nedir?](#htmx-nedir)
- [Golang Nedir?](#golang-nedir)
- [Air](#air)
- [Örnekler](#örnekler)

## Kurulum

Projenizi klonladıktan sonra gerekli bağımlılıkları yüklemek için aşağıdaki adımları izleyin:

```bash
git clone https://github.com/emrekndl/htmx-golang.git
cd htmx-golang
go mod tidy
```

## Kullanım

Sunucuyu başlatmak için:

```bash
go run main.go
```

Tarayıcınızdan `http://localhost:1323` adresine giderek uygulamayı görüntüleyebilirsiniz.

## Proje Yapısı

```
htmx-golang/
├── cmd/
│   ├── main.go
│   └── class-examples/
│       └── blocks/
│           └── main.go
├── views/
│   ├── index.html
│   ├── blocks.html
│   └── contacts.html
├── css/
│   └── index.css
└── images/
```

## HTMX Nedir?

HTMX, HTML özniteliklerini kullanarak modern web uygulamaları geliştirmeyi sağlayan bir JavaScript kütüphanesidir. Daha fazla bilgi için [htmx.org](https://htmx.org) adresini ziyaret edebilirsiniz.

![HTMX Banner](https://img.shields.io/badge/%3C/%3E%20htmx-3D72D7?style=for-the-badge&logo=mysl&logoColor=white)

## Golang Nedir?

Golang, Google tarafından geliştirilmiş açık kaynaklı bir programlama dilidir. Daha fazla bilgi için [golang.org](https://golang.org) adresini ziyaret edebilirsiniz.

![Golang Banner](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

## Air

[Air](https://github.com/cosmtrek/air) paketi, Go uygulamalarını geliştirme sürecinde canlı yeniden yükleme (live reload) özelliği sağlar. 

Kurmak için:

```bash
go install github.com/air-verse/air@latest
```

## Örnekler

### Basit Bir HTMX Kullanımı

`index.html` dosyasında bir butona tıklandığında bir GET isteği gönderen basit bir örnek:

```html
<button hx-get="/message" hx-target="#message">Mesajı Göster</button>
<div id="message"></div>
```

### Go Tarafında Echo ile Yanıt Verme

`main.go` dosyasında ilgili endpoint'i tanımlama:

```go
package main

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "index.html", nil)
    })

    e.GET("/message", func(c echo.Context) error {
        return c.String(http.StatusOK, "Merhaba, HTMX ile Go!")
    })

    e.Logger.Fatal(e.Start(":8080"))
}
```
