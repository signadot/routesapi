package retrypolicy

import (
	"math/rand"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

// BackoffNexter takes a k8s Backoff and turns it into a Nexter
type BackoffNexter wait.Backoff

func (b *BackoffNexter) Next() time.Duration {
	return (*wait.Backoff)(b).Step()
}

// LubyNexter provides a Nexter which follows the Luby series (1,1,2,1,1,2,4,1,1,2,4,8,...)
// which has the property that after every time a new max > 1 is reached in the sequence,
// the sum of all the 1s is equal to the sum of all the 2s and so forth for all the values
// in the sequence.  This provides a "balanced" or "reluctant" exponential backoff: it
// still will backoff exponentially, but the frequency of the backoff reduces exponentially
// at the same time.  Such a Nexter will tend to succeed much sooner than standard
// exponential backoff in the face of error states of unpredictable duration, it will
// still tend to increase pauses between tries over time, but periodically burst at
// faster rates, with period increasing over time.
//
// For more info on the Luby series, please see
//
//  1. Donald Knuth The Art of Computer Programming, Volume 4, section 7.2.2.2 -- Satisfiability, pages 306-308, 2015.
//  2. Optimal Speedup of Las Vegas Algorithms, M Luby et al, 1993 (https://www.cs.utexas.edu/~diz/Sub%20Websites/Research/optimal_speedup_of_las_vegas_algorithms.pdf)
//  3. The original luby algorithm published to Twitter in 2009: https://github.com/go-air/gini/blob/7aed76bfd77fb2e84a145453b1097f6179acc6e6/internal/xo/luby.go#L15
type LubyNexter struct {
	BaseStep      time.Duration
	MaxTimeout    time.Duration
	JitterPercent uint
	exp, turns    uint
}

func DefaultLubyNexter() *LubyNexter {
	return &LubyNexter{
		BaseStep:      time.Second,
		JitterPercent: 10,
		MaxTimeout:    30 * time.Second,
	}
}

func (l *LubyNexter) stepNext() uint {
	res := uint(1 << l.exp)
	if res&l.turns == 0 {
		l.exp = 0
		l.turns++
	} else {
		l.exp++
	}
	return res
}

func (l *LubyNexter) Next() time.Duration {
	scale := time.Duration(l.stepNext())
	target := scale * l.BaseStep
	// only consider max if specified and with respect
	// to the non-jitterized next duration
	if l.MaxTimeout != 0 && target > l.MaxTimeout {
		l.exp = 0
		l.turns = 0
		scale = time.Duration(l.stepNext())
		target = scale * l.BaseStep
		// in case MaxTimeout is less than base step
		if target > l.MaxTimeout {
			target = l.MaxTimeout
		}
	}
	return target + l.jitter(target)
}

func (l *LubyNexter) jitter(d time.Duration) time.Duration {
	j := rand.Int63n(int64(d/100)) * int64(l.JitterPercent/2)
	if rand.Intn(2) == 1 {
		j = -j
	}
	return time.Duration(j)
}
