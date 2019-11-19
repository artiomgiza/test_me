
Short introduction
------------------

This package helps us to calculate BBQ price for N people.

The total BBQ price is the combination of (see `coolPriceCalculator` implementation):
- meet price (beef or chicken if no beef available)
- mangal price
- coal price

_Total price_ = _meet price_ + _mangal price_ + _coal price_

Purpose
-------

Our purpose is to compare different kind of test approaches for the same code base.

We'll try to prepare best possible tests for the same `coolPriceCalculator` (see .../BBQ.go) with each approach.

Then compare and discuss.

Currently we have:
----
1. ginkgo example (./ginkgo_me/BBQ_test.go) - run cmd: `ginkgo ./ginkgo_me -r`
2. ...

How to add more:
----------------
Please add directory for your approach (native, testify, ...) in the project root (near `ginkgo_me` dir) and copy the `ginkgo_me/BBQ.go` file there.
Now add tests (by creating PR).

Once there will be some examples to compare we might set up discussion meeting...
