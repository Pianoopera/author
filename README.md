# Git-Author-Log-History
Search Git Author Log Commit History

## シンボリックリンクの作成
毎回shellを実行するのも面倒なのでシンボリックの作成すると楽です。

スクリプトの実行権限の修正

```shell
chmod +x /path/to/main.sh
```

システムのパスが設定されているディレクトリにシンボリックを作成する
```shell
ln -n /path/to/main.sh /usr/local/bin/author
```

※pathは自分の環境に合わせて修正してください。