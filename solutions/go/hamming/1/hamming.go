package hamming

import "errors"

func Distance(a, b string) (int, error) {
	// panic("Implement the Distance function")
    if len(a) != len(b){
        return 0, errors.New("Strands must be of equal length")
    }

    distance :=0
    for i := range a{
        if a[i] != b[i]{
            distance++
        }
    }
    return distance, nil
}
