package main

//Mutex and Condition

//TODO: 1. Create a basic GIN server with a single GET route returning JSON.

//TODO: 2. Create a global variable `var counter int`.

//TODO: 3. Inside the GET handler, increment `counter++` and return it.

//TODO: 4. Run the app using `go run -race main.go` and spam the endpoint to trigger a crash/warning.

//TODO: 5. Import `sync` package and add `var mu sync.Mutex`.

//TODO: 6. Wrap the increment logic with `mu.Lock()` and `mu.Unlock()` to fix the race condition.

// Read/Write Efficiency

//TODO: 7. Create a struct `UserCache` containing a `map[string]string` and `sync.RWMutex`.

//TODO: 8. Create a POST handler to add users that uses `Lock()` (Write Lock).

//TODO: 9. Create a GET handler to read users that uses `RLock()` (Read Lock).

//TODO: 10. Test with multiple simultaneous requests to understand why RLock allows parallel reads.

// Parallel Execution

//TODO: 11. Create a "Dashboard" endpoint that needs to fetch data from 3 "slow" functions (simulate with time.Sleep).

//TODO: 12. Initialize `var wg sync.WaitGroup` inside the handler.

//TODO: 13. Use `wg.Add(3)` before starting the goroutines.

//TODO: 14. Wrap each slow function in `go func() { defer wg.Done(); ... }`.

//TODO: 15. Use `wg.Wait()` to block the main response until all background tasks finish.

// Architecture and Safety

//TODO: 16. Create a singleton DB connection function using `sync.Once` so it connects only once.

//TODO: 17. Create an endpoint that launches a background goroutine (e.g., logging or analytics).

//TODO: 18. Inside that handler, use `cCp := c.Copy()` before passing the context to the goroutine.

//TODO: 19. Attempt to use the original `c` in the goroutine to see why it fails (the "Context Trap").

// Project

//TODO: 20. Build a "Ticket Booking API" with 100 tickets.

//TODO: 21. Ensure valid concurrency: Prevent selling 101 tickets when 500 users click "Buy" at the same time.
