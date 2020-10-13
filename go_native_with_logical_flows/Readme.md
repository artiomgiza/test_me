I believe that tests are needed not only for code verification. They are also a kind of documentation.  
Therefore, they should reflect the business logic of the code under test, and not just if/else structure.


P.S.
---
A package should have interfaces and mocks for external dependencies within itself.  
We don't need these dependencies:
```
import (
	beeffarm "github.com/artiomgiza/test_me/pkgs/beef-farm"
	chickenfarm "github.com/artiomgiza/test_me/pkgs/chicken-farm"
	coalfarm "github.com/artiomgiza/test_me/pkgs/coal-farm"
	mangalstore "github.com/artiomgiza/test_me/pkgs/mangal-store"
)
```
We don't need these stuff:
```
type Provider interface {
	CalculatePrice(peopleCounter int) (int, error)
}
var Instance Provider = coolPriceCalculator{
	// ... inject real fields ...
}
```