gakki-api
===

がっきーアプリケーションインターフェース

## 使い方
* http requestを投げるといい感じにがっきーの画像を返してくれます

## 構成
* がっきーの画像を集める君
    * Google画像検索にクエリを投げる
    * DBに入れる

* 本当にがっきーの画像か判定する君
    * 人による温かみのある判定
    * 未判定のものを表示する
    * 判定結果を保存する

* がっき画像を返す君
    * リクエストに対して判定済みのがっきー画像を返す
