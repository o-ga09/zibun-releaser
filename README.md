# zibun-releaser

「じぶんりりーすノート」のリリースを自動化するGithub Actionsのカスタムアクションです。

---

**「じぶんりりーすノート」とは**

「じぶんリリースノート」とは、個人が自身の活動や成果、成長などを記録するためのノート、またはその記録方法のことです。ソフトウェアの更新履歴を記したリリースノートになぞらえて、自分自身のバージョンアップを記録していくというコンセプトから生まれました。

- 記載する内容

具体的な内容としては、以下のようなものが挙げられます。

- 活動記録: その月に取り組んだプロジェクト、学習内容、イベント参加などを記録
- 成果: 達成できた目標、身につけたスキル、資格取得などを記録
- 反省点: 課題として残ったこと、改善点などを記録
- 目標設定: 次の期間に向けての目標、やりたいことなどを記録
- KPT: Keep(継続すること), Problem(課題), Try(挑戦すること)をまとめる振り返り手法
- その他: 感想、感情、気づきなどを自由に記録

## 参考イメージ

![参考イメージ](https://github.com/user-attachments/assets/c67b5df8-0790-40fc-9327-267450c7facd)

## 使い方

じぶんリリースノートのフォーマットは、[example.md](./_example/example.md)に従ってください。

- `secrets.PAT`には、リポジトリ環境変数に`personal access token`を設定してください。
- `secrets.APIKEY`には、Gemini API KEYをリポジトリ環境変数に設定してください。
(Open AI APIにも対応する予定)
- `filepath`は、リポジトリルートからの相対パスで指定してください。
- `release-version`は、セマンティックバージョニングに従ってください。

```
- name: Use Custom Actions
uses: o-ga09/zibun-releaser@v0.0.15
with:
    release-version: '<リリースしたいバージョン>'
    github-token: ${{ secrets.PAT }}
    filepath: "<じぶんリリースノートの元となるmarkdownファイル>"
    apikey: ${{ secrets.APIKEY }}
```

## Buils and Run

makefile中の環境変数を設定してください

`<Gemini APIKEY>`は、Gemini APIKEYを設定する。
`_example/example.md`は、任意のじぶんリリースの元となるファイルが格納された場所を指定する

```
export APIKEY=<Gemini APIKEY>
export ENV="CI"
export FILEPATH="_example/example.md"
```

### Build

```
$ make build
```

### Run

```
$ make run
```

実行後、リリースノートが、/tmp/releasenote.mdに作成されます。

## Contribute

Issue、Pull Request受け付けています。
バグ修正、Feature Requestをお願いいたします。

## Roadmap

- [ ] SNSに自動で投稿する
- [ ] note自動で投稿する
- [ ] windows環境での実行に対応する

## Licence

MIT Licese .

2024 @o-ga09
