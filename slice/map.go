package iter

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

// Map iterates through slc and, for each element, calls fn with its index
// and the element itself. if fn returns a non-nil error, Map returns immediately
// with (nil, <the_error>). Otherwise, Map assigns the first return value to a new
// slice at the same index and moves on. If all calls to fn return nil errors,
// the final slice will be returned along with a nil error
//
// Example usage of this function:
//
//	slc := []int{1, 2, 3, 4, 5}
//	Map(slc, func(_ uint, val int) (int, error) {
//		return val+1, nil
//	})
func Map[T any, U any](slc []T, fn func(i uint, t T) (U, error)) ([]U, error) {
	ret := make([]U, len(slc))
	for i, t := range slc {
		u, err := fn(uint(i), t)
		if err != nil {
			return nil, err
		}
		ret[i] = u
	}
	return ret, nil
}

// ParMap is similar to Map, except calls fn in a separate goroutine for 
// each element in slc. If any one of the calls to fn returns an error,
// the first that returns an error will have that error returned, and nil will
// be returned for the slice. fn will be passed a context that is derived from 
// the input ctx.
//
// Common use of this function is to do operations on a slice that can be
// done concurrently. Often this applies to "embarassingly parallel" problems.
//
// Example usage:
//
//	var mut sync.Mutex
//	slc := []int{1, 2, 3, 4, 5}
//	ParMap(context.Background(), slc, func(_ context.Context, _ uint, val int) (string, error) {
//		return strconv.Itoa(val), nil
//	})
func ParMap[T any, U any](
	ctx context.Context,
	slc []T,
	fn func(context.Context, uint, T) (U, error),
) ([]U, error) {

	g, ctx := errgroup.WithContext(ctx)
	ret := make([]U, len(slc))
	for idx, v := range slc {
		i, v := uint(idx), v
		g.Go(func() error {
			r, err := fn(ctx, i, v)
			if err == nil {
				ret[i] = r
			}
			return err
		})
	}
	
	if err := g.Wait(); err != nil {
		return nil, err
	}
	
	return ret, nil
}