# Test Case

# 引数を何も指定しないケース：
./main.sh

# 検索対象ディレクトリが存在しない場合：
./main.sh non_existent_directory

# 検索対象ディレクトリが有効な場合：
./main.sh valid_directory account_name

# アカウント名が指定されていない場合：
./main.sh valid_directory

# アカウント名が指定されている場合：
./main.sh valid_directory account_name

# ディレクトリが存在しないディレクトリを指定した場合：
./main.sh directory_without_git

# ディレクトリが存在するディレクトリを指定した場合：
./main.sh directory_with_git account_name