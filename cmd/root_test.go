package cmd_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/Pianoopera/author/cmd"
)

var originalArgs = os.Args

// 複数の引数を渡す場合は、setArgsを複数回呼び出す
func setArgs(args ...string) {
	os.Args = append(originalArgs, args...)
}

func resetArgs() {
	os.Args = originalArgs
}

func Test_versionSubCmd(t *testing.T) {
	setArgs("version")
	defer resetArgs()

	got := PickStdout(t, func() { cmd.Execute() })
	want := "author version 0.1.2"
	if got != want {
		t.Errorf("subCmd.Execute() = %v, want = %v", got, want)
	}
}

// 引数の指定がされている場合のテスト
func Test_rootSubCmd(t *testing.T) {
	// 第一引数：ディレクトリのパス 第二引数：ユーザー名
	setArgs("./", "Hoge")
	defer resetArgs()

	got := PickStdout(t, func() { cmd.Execute() })
	// コマンドが無事終了したかステータスコードで確認
	want := " Done!!!!!!!!"
	if got != want {
		t.Errorf("subCmd.Execute() = %v, want = %v", got, want)
	}
}

func PickStdout(t *testing.T, fnc func()) string {
	t.Helper()
	backup := os.Stdout
	defer func() {
		os.Stdout = backup
	}()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("fail pipe: %v", err)
	}
	os.Stdout = w
	fnc()
	w.Close()
	var buffer bytes.Buffer
	if n, err := buffer.ReadFrom(r); err != nil {
		t.Fatalf("fail read buf: %v - number: %v", err, n)
	}
	s := buffer.String()
	return s[:len(s)-1]
}
