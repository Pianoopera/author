#############
# Test Case #
#############

test1:
	# 正常系：
	./main.sh ~/git/Git-Author-Log-History/ Pianoopera

test2:
	# 引数を何も指定しないケース：
	./main.sh

test3:
	# 検索対象ディレクトリが存在しない場合：
	./main.sh non_existent_directory

test4:
	# 検索対象ディレクトリが有効 & アカウントが存在しない場合：
	./main.sh ~/git non_existent_account

test5:
	# アカウント名が指定されていない場合：
	./main.sh ~/git