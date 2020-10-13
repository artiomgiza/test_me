
Intro
-----

This package helps us to calculate BBQ price for N people.

The total BBQ price is the combination of (see `coolPriceCalculator` implementation):
- meet price (beef or chicken if no beef available)
- mangal price
- coal price

_Total price_ = _meet price_ + _mangal price_ + _coal price_

Purpose
-------

Our purpose is to see BDT testing approach based on Ginkgo infra.

We'll try to prepare best possible test for the `coolPriceCalculator` (see .../BBQ.go).

---
Notes:

- [example with structure explanation](./ginkgo_me)
- run test cmd: `ginkgo ./ginkgo_me -r`
