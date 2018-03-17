package main

import "github.com/itorisaias/html"
import "time"
import "fmt"

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url1)
	c3 := html.Titulo(url1)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(900 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"

	}
}

func main() {
	campeao := oMaisRapido(
		"https://backoffice.beblue.com.br",
		"https://www.google.com",
		"https://www.amazon.com",
	)

	fmt.Println(campeao)
}