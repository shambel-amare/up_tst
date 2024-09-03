package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type request struct {
	Array []int32 `json:"array"`
}
type response struct {
	Result int32 `json:"result"`
}

// helloHandler handles HTTP requests to the /hello endpoint
func handleNumberSum(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	req := request{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}
	sum := 0
	for _, b := range req.Array {
		sum += int(b)
	}
	resp := response{Result: int32(sum)}
	// Write the response back to the client
	jRsp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jRsp)
}

func main() {
	http.HandleFunc("/add", handleNumberSum)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
