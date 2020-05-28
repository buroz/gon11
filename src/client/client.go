package client

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	XMLName xml.Name `xml:"x:Envelope"`
	X       string   `xml:"xmlns:x,attr"`
	Sch     string   `xml:"xmlns:sch,attr"`
	Header  struct{} `xml:"x:Header"`
	Body    struct {
		Request interface{}
	} `xml:"x:Body"`
}

func (this *Request) Call(service string, request interface{}) []byte {
	this.X = "http://schemas.xmlsoap.org/soap/envelope/"
	this.Sch = "http://www.n11.com/ws/schemas"
	this.Body.Request = request

	out, _ := xml.MarshalIndent(&this, " ", "  ")
	body := string(out)

	fmt.Println(body)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	response, err := client.Post(service, "text/xml", bytes.NewBufferString(body))

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	s := strings.TrimSpace(string(content))
	fmt.Println(s)

	return content
}
