# Go Interfaces ![gha build](https://github.com/karantan/go-interfaces/workflows/Go/badge.svg)

When I started writing tests in Go, I sometimes used interfaces to mock
things in tests. Then I discovered that this is not the correct way:

- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments#interfaces)
- [SOLID Go Design](https://dave.cheney.net/2016/08/20/solid-go-design)
- [Accept Interfaces Return Struct in Go](https://mycodesmells.com/post/accept-interfaces-return-struct-in-go)

Similar to [github.com:karantan/go-concurrency](https://github.com/karantan/go-concurrency),
I decided to create (another) "lessons learned/good practices" repo that shows how to
accepting interfaces and returning structs in practice.

_Preemptive abstractions make systems complex._

I won't explain in detail why it's a bad practice to return interfaces because it's
already been explained in the posts linked above. Please read them before going
through the code here.

Additional resources:

- [Learn Go with test-driven development](https://github.com/quii/learn-go-with-tests)
- [Writing Good Unit Tests; Don’t Mock Database Connections](https://qvault.io/clean-code/writing-good-unit-tests-dont-mock-database-connections/)
