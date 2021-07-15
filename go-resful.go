package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"io"
	"net/http"
	"time"
)

func main()  {
	//create web service
	webservice := new(restful.WebService)
	webservice.Route(webservice.GET("/ping").To((pingTime)))
	restful.Add(webservice)
	http.ListenAndServe(":8000", nill)
}

func pingTime(req *restful.Request, resp *restful.Response)  {
	//write to the response
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}