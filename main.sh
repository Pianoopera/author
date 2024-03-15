#!/bin/bash

# 検索対象ディレクトリとして最初の引数を受け取る
SEARCH_DIR="$1"
# 検索対象ディレクトリが指定されていない、または無効な場合は警告を出してスクリプトを続行
if [[ ! -d "$SEARCH_DIR" ]]; then
    echo "警告: 指定されたディレクトリが存在しないか、無効です。"
    echo "使い方: $0 <検索対象ディレクトリ> <アカウント名1> <アカウント名2> ..."
    exit 0
fi

shift  # 最初の引数を削除し、残りの引数（アカウント名）を扱う

# アカウント名が指定されていない場合は警告を出してスクリプトを続行
if [ "$#" -eq 0 ]; then
    echo "警告: アカウント名が指定されていません。"
    echo "使い方: $0 <検索対象ディレクトリ> <アカウント名1> <アカウント名2> ..."
    exit 0
fi

# アカウント名の表示を囲むための関数
print_boxed() {
    local -r msg="$1"
    local edge=$(printf '%*s' "${#msg}")
    edge=${edge// /-}
    echo "+$edge+"
    echo "|$msg|"
    echo "+$edge+"
}


# 指定されたディレクトリ以下で`.git`ディレクトリを検索
find "$SEARCH_DIR" -type d -name ".git" | while read repo; do
    repo_path=$(dirname "$repo")  # .gitを含むパスから.gitを削除してリポジトリのパスを取得
    print_boxed " Repository: $repo_path "
    cd "$repo_path" || continue   # そのリポジトリのパスに移動

    # 各アカウントについてコミットを検索
    for author in "$@"; do
        echo " <<< Account: $author >>>"
        echo ""
        git log --pretty=format:"%h %cd [%Creset%s%C(yellow)%d%C(reset)]"  --author="$author" --graph --date=short --decorate --all $commit
        echo ""
    done
    cd - > /dev/null  # 元のディレクトリに戻る
done
