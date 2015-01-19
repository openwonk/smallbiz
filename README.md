US Small Business Admin API Client
===================

Library provides access to the [US SBA API](http://data.worldbank.org/developers).  No registration key is necessary.

###Usage
```go
package main

import (
	biz "github.com/openwonk/smallbiz"
)

func main() {
	// Series{city, county, state, format}
	s := biz.Series{"Seattle", "King", "WA", "json"}

	s.OneCity().Write("seattle.json")
	s.OneCounty().Write("king.json")
	s.AllState().Write("wa.json")
}
```
<br>
<br>

<hr>
<small>
<strong>OpenWonk &copy; 2015 [MIT License](https://github.com/openwonk/mit_license)</strong>
</small>

