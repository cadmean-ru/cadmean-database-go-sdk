package database

import (
	"log"
	"testing"
)

func TestClient_NewApiRequest(t *testing.T) {
	var db = NewClient("http://127.0.0.1:69", "ecopoints", "ZTNVclBNTEhJdFZPRVo1Y1draXdtQlNXc3ZjT0hCVUQ=")

	docs, err := db.Collection("taxiOrders").FindDocs()
	if err != nil {
		t.Fatal(err)
	}

	arr := docs.ToArray()

	log.Println(arr)
}
