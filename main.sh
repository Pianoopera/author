#!/bin/bash

# 検索対象ディレクトリとして最初の引数を受け取る
SEARCH_DIR="$1"
# 検索対象ディレクトリが指定されていない、または無効な場合は警告を出してスクリプトを続行
if [[ ! -d "$SEARCH_DIR" ]]; then
    echo "警告: 指定されたディレクトリが存在しないか、無効です。"
fi

shift  # 最初の引数を削除し、残りの引数（アカウント名）を扱う

# アカウント名が指定されていない場合は警告を出してスクリプトを続行
if [ "$#" -eq 0 ]; then
    echo "警告: アカウント名が指定されていません。"
fi

# 指定されたディレクトリ以下で`.git`ディレクトリを検索
find "$SEARCH_DIR" -type d -name ".git" | while read repo; do
    repo_path=$(dirname "$repo")  # .gitを含むパスから.gitを削除してリポジトリのパスを取得
    echo "▲▲▲ Repository: $repo_path"
    cd "$repo_path" || continue   # そのリポジトリのパスに移動

    # 各アカウントについてコミットを検索
    for author in "$@"; do
        echo "▲▲▲ Account: $author"
        git log --since="3 months ago" --author="$author" --pretty=format:"%H" | while read commit; do
            # コミットが属するブランチを表示（複数ある場合があるため、一例を示す）
            echo '▲▲▲▲▲▲ Commit Message: '  `git log -1 --pretty=format:"%s" $commit`
            echo "▲▲▲▲▲▲ Commit Date: " `git log -1 --pretty=format:"%ad" --date=short`
            echo "▲▲▲▲▲▲ Commit Hash: $commit"
            # git branch --contains "$commit" | while read branch; do
            #     echo "▲▲▲▲▲▲▲▲▲ Branch: $branch"
            # done
        done
        echo ""
    done
    cd - > /dev/null  # 元のディレクトリに戻る
done
