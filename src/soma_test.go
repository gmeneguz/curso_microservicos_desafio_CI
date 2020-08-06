package main
import "testing"

func TestSoma( t *testing.T ) {
	zero := soma(0,0)
	if(zero != 0){
		t.Errorf("Zero sum failed, exptect 0 but got %v", zero)
	}
	ten := soma(5,5)
	if(ten != 10){
		t.Errorf("5 + 5 sum failed, exptect 10 but got %v", zero)
	}

}