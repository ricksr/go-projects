package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func Handler(request Request) (Response, error) {
	fmt.Println("🚀 Server started, enter endpoint to execute Lambda")
	return Response{
		Message: fmt.Sprintf("Process Request ID %f", request.ID),
		Ok:      true,
	}, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
	fmt.Println("Base Hit 🌟")
}

func main() {
	http.HandleFunc("/", homePage)
	lambda.Start(Handler)
}
