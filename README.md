# koron/palicnv

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron/palicnv)](https://pkg.go.dev/github.com/koron/palicnv)
[![Actions/Go](https://github.com/koron/palicnv/workflows/Go/badge.svg)](https://github.com/koron/palicnv/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron/palicnv)](https://goreportcard.com/report/github.com/koron/palicnv)

Convert image files to a size and color model compatible with PaliGemma.

## Getting Started

Install or update:

```console
$ go install github.com/koron/palicnv@latest
```


## 説明 (Description in Japanese)

画像をPaliGemmaに適したサイズおよびカラーモデルに変換するコンバーターです。

PaliGemmaは縦と横が同じ正方形サイズの画像を入力とします。
コンバーターではそれに合わせて画像サイズを変換します。

対応する画像のフォーマットはGIF/JPEG/PNGで拡張子により自動判定します。
アニメーションGIFについては、画像情報量(エントロピー)が最も大きくなるフレームを代表フレームとして自動的に選択します。

本コマンドの使用例は以下の通りです。

```console
$ palicnv path/to/foobar.png
```

この時、出力は、自動で拡張子を取り除き接尾辞 `_224s.jpg` を付与して、 `path/to/foobar_224s.jpg` となります。
出力ファイル名は別途オプションで変更できます。

より一般的な使用方法は以下のようになります。

```console
$ palicnv [OPTIONS] {INPUT FILE}
```

サポートしているオプションは以下の通りです。

*   `-output {出力ファイル名}` 出力ファイル名を指定する。
    省略時は `{入力ファイル名のベース部分}_{SIZE}s.jpg` となる。
    拡張子によりフォーマットを指定できる。
    対応している拡張子は次の通り: `.jpeg`, `.jpg`, `.png`, `.gif`
*   `-size {数字}` 出力する画像のサイズを指定する。
    デフォルトは224。PaliGemmaは224, 448, 896のいずれかを推奨している。
