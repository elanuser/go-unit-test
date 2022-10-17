package main 

// reference 
// https://medium.com/yemeksepeti-teknoloji/how-to-write-unit-test-in-go-1df2b98ad510

import (
"testing"
"fmt"
)

func TestSayHello(t *testing.T) {
	name := "Mert"
	expect := fmt.Sprintf("hello %s", name)
	if sayHello(name) != expect {
		t.Errorf("\"sayHello('%s')\" FAILED, expected: %s, got: %s", name, expect, sayHello(name))
	} else {
		t.Logf("\"sayHello('%s')\" SUCCEED, expected: %s, got: %s", name, expect, sayHello(name))
	}
}

func Test_SayGoodBye(t *testing.T) {
	name := "Mert"
	expected := fmt.Sprintf("Bye Bye %s!", name)
	result := sayGoodBye(name)

	if result != expected {
		t.Errorf("\"sayGoodBye('%s')\" FAILED, expected -> %v, got -> %v", name, expected, result)
	} else {
		t.Logf("\"sayGoodBye('%s')\" SUCCEDED, expected -> %v, got -> %v", name, expected, result)
	}
}


func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sayHello("Benchmark")
	}
}