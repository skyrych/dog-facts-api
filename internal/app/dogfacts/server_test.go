package dogfacts

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewFactServer(t *testing.T) {
	needyFacts := NewFactServer()
	if needyFacts == nil || needyFacts.Facts == nil || needyFacts.Rand == nil {
		t.Fatalf("Failed to get an instance of FactServer struct {}")
	}
	if len(needyFacts.Facts) == 0 {
		t.Fatalf("Facts slice is empty for no obvious reasons")
	}
	if len(needyFacts.Facts) != 15 {
		t.Errorf("The slice with Needy facts was modified.")
	}

}

func TestFactsHandler(t *testing.T) {
	Rand := rand.New(rand.NewSource(1))
	Facts := []string{"test_fact_1", "test_fact_2", "test_fact_3"}
	factServerInstance := &FactServer{Facts: Facts, Rand: Rand}
	rr1 := httptest.NewRecorder()
	mockRequest, err := http.NewRequest("GET", "/facts", nil)
	if err != nil {
		t.Fatalf("Failed to create a mock request")
	}
	factServerInstance.factsHandler(rr1, mockRequest)
	if rr1.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr1.Code, http.StatusOK)
	}
	if rr1.Body.String() != "test_fact_3\n" { // <--- Expected: test_fact_3
		t.Errorf("handler returned wrong value: got %q want %q",
			rr1.Body.String(), "test_fact_3\n")
	}

	rr2 := httptest.NewRecorder()

	factServerInstance.factsHandler(rr2, mockRequest)
	if rr2.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr2.Code, http.StatusOK)
	}
	if rr2.Body.String() != "test_fact_1\n" { // <--- Expected: test_fact_1
		t.Errorf("response returned wrong value: got %q want %q",
			rr2.Body.String(), "test_fact_1\n")
	}

}

func TestFactsHandler_EmptyFacts(t *testing.T) {
	Rand := rand.New(rand.NewSource(0))
	Facts := []string{}
	factServerInstanceEmpty := &FactServer{Facts: Facts, Rand: Rand}
	rrEmpty := httptest.NewRecorder()
	mockRequestEmpty, err := http.NewRequest("GET", "/facts", nil)
	if err != nil {
		t.Fatalf("Failed to create a mock request")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic when expected (empty facts slice).")
		} else {
			t.Logf("Successfully recovered from expected panic: %v", r)
		}
	}()
	factServerInstanceEmpty.factsHandler(rrEmpty, mockRequestEmpty)
}
