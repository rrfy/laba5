package main

import "fmt"

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (r *RealSubject) Do() string {
	return "RealSubject: executing operation"
}

type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Do() string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	result := "Proxy: calling RealSubject to execute operation\n"
	result += p.realSubject.Do()
	return result
}

func main() {
	proxy := &Proxy{}

	result := proxy.Do()
	fmt.Println(result)
}
