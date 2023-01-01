package dsalgo

// --- Directions
// Print out the n-th entry in the fibonacci series.
// The fibonacci series is an ordering of numbers where
// each number is the sum of the preceeding two.
// For example, the sequence
//  [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
// forms the first ten entries of the fibonacci series.
// Example:
//   fib(4) === 3

type IntInIntOut func(n int) int

// Dynamic Programming
// this function caches the result of the fib function to be efficient
// this returns a function to use to calculate fibonacci series
func memoize(fn IntInIntOut) IntInIntOut {
	cache := make(map[int]int)
	return func(n int) int {
		if _, ok := cache[n]; ok {
			return cache[n]
		}
		res := fn(n)
		cache[n] = res
		return res
	}
}

func Fib(n int) int {
	newFunc := memoize(slowFib)
	return newFunc(n)
}

func slowFib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
