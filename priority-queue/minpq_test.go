package priority_queue

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestReadFile(t *testing.T) {
	file, err := os.Open("test.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	// bufio.Scanner
	//scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	fmt.Println(line)
	//}
	//
	//if err = scanner.Err(); err != nil {
	//	t.Fatal(err)
	//}

	// file.Read
	//buffer := make([]byte, 64)
	//for {
	//	n, err := file.Read(buffer)
	//	if n > 0 {
	//		fmt.Println(string(buffer[:n]))
	//	}
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//}

	// bufio.Reader
	//reader := bufio.NewReader(file)
	//for {
	//	line, err := reader.ReadString('\n')
	//	if line != "" {
	//		fmt.Println(line)
	//	}
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//}

	// csv
	reader := csv.NewReader(file)
	for {
		r, err := reader.Read()
		if len(r) > 0 {
			fmt.Println(r)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestNewMinPQ(t *testing.T) {
	M := 10
	minPQ := NewMinPQ[Transaction](M + 1)

	file, err := os.Open("test.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// reader
	reader := csv.NewReader(file)
	for {
		ss, err := reader.Read()
		if len(ss) == 3 {
			date, err := time.Parse("20060102", ss[1])
			if err != nil {
				t.Fatal(err)
			}
			amount, err := strconv.ParseFloat(ss[2], 64)
			if err != nil {
				t.Fatal(err)
			}

			transaction := Transaction{
				Name:   ss[0],
				Date:   date,
				Amount: amount,
			}

			minPQ.Insert(transaction)
			if minPQ.Size() > M {
				minPQ.Delete()
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
	}

	minPQ.Range(func(_ int, a Transaction) bool {
		fmt.Println(a)
		return true
	})
}
