package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/matumoto1234/go-compro-library/graph"
	"github.com/matumoto1234/go-compro-library/graph/dijkstra"
)

func doMain(t *testing.T, f func(), path string) ([]byte, time.Duration) {
	t.Helper()

	// 標準入出力の差し替え
	backupIn, backupOut := os.Stdin, os.Stdout
	defer func() {
		os.Stdin = backupIn
		os.Stdout = backupOut
	}()

	in, err := os.Open(path)
	if err != nil {
		t.Fatalf("%+v\n", err)
	}

	os.Stdin = in

	tmp, err := os.CreateTemp(t.TempDir(), "tmp")
	if err != nil {
		t.Fatalf("%+v\n", err)
	}

	defer func() {
		tmp.Close()
		os.Remove(tmp.Name())
	}()

	defer func() {
		if t.Failed() {
			os.Exit(1)
		}
	}()

	os.Stdout = tmp

	now := time.Now()

	f()

	duration := time.Since(now)

	tmp.Seek(0, 0)

	var buffer bytes.Buffer
	if _, err := buffer.ReadFrom(tmp); err != nil {
		t.Fatalf("%+v\n", err)
	}

	return buffer.Bytes(), duration
}

type JudgeStatus string

const (
	Accepted    JudgeStatus = "Accepted"
	WrongAnswer JudgeStatus = "WrongAnswer"
)

// out : i行目に「i番目の頂点までの最短経路長」がある
// ansにある辺を使って最短経路長を計算し、outの総和と一緒ならAC
func specialJudge(t *testing.T, in, out, ans io.Reader) JudgeStatus {
	t.Helper()

	bufin := bufio.NewReader(in)
	bufout := bufio.NewReader(out)
	bufans := bufio.NewReader(ans)

	// 入力から辺の読み込み
	var n, m int
	fmt.Fscan(bufin, &n, &m)

	edges := make([]*graph.Edge[int], 0)

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(bufin, &a, &b, &c)

		a--
		b--

		edges = append(edges, graph.NewEdge(a, b, c))
	}

	// 回答の辺の読み込み
	line, err := bufans.ReadString('\n')
	if err != nil {
		t.Fatalf("%+v\n", err)
	}

	line = strings.Trim(line, "\n")
	edgeIndexes := strings.Split(line, " ")

	g := graph.NewAdjacencyList[int](n)

	for _, v := range edgeIndexes {
		idx, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return WrongAnswer
		}

		g.AddEdge(edges[idx-1])

		re := graph.NewEdge(edges[idx-1].To, edges[idx-1].From, edges[idx-1].Cost)
		g.AddEdge(re)
	}

	// 回答から、各頂点(1<=i<n)への最短経路長の和を計算
	sumAns := 0

	d := dijkstra.NewOrdered[int](g, 0, 1<<60)
	for v := 1; v < n; v++ {
		sumAns += d.Distances[v]
	}

	// outから、各頂点(0<=i<n)への最短経路長の和を計算
	sumOut := 0

	for v := 0; v < n; v++ {
		line, err := bufout.ReadString('\n')
		if err != nil {
			t.Fatalf("%+v\n", err)
		}

		line = strings.Trim(line, "\n")
		dist, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			t.Fatalf("%+v\n", err)
		}

		sumOut += int(dist)
	}

	if sumAns == sumOut {
		return Accepted
	}

	return WrongAnswer
}

func TestMain(t *testing.T) {
	err := filepath.Walk("./test", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) != ".in" {
			return nil
		}

		ans, duration := doMain(t, main, path)

		in, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}

		out, err := os.Open(strings.Replace(path, ".in", ".out", 1))
		if err != nil {
			t.Fatal(err)
		}

		status := specialJudge(t, in, out, bytes.NewReader(ans))

		log.Printf("[INFO] name : %20s status : %20s time : %5v[ms]\n", path, status, duration.Milliseconds())

		return nil
	})

	if err != nil {
		t.Fatalf("%+v\n", err)
	}
}
