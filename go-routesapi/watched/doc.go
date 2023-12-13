// package watched provides read access to workload routes backed by a retrying client
//
// The retry methodology we use is to
//
//  1. divide Recv into an outer receive, which wraps the grpc (inner) watch receives.
//  1. keep an in memory model of the routes upon every successful _outer_ receive.
//  1. keep a buffer encoding a diff against the above model which is returned
//     with priority and in order with outer receive.
//  1. Upon error of inner receive, we retry-re-call the inner watch, get synced,
//     and then encode the diff buffer.  Upon error, we keep retrying with backoff.
//  1. The in-memory model is then given an exported interface for getting routes
//     and checking routing key membership
//
// The above guarantees that retries are not visible to the outer Recv until
// synced.  This, in turn, guarantees that we can provide only ADDs until sync.
package watched
