
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

We want to compare different approaches.
We'll try to prepare best possible test for the `coolPriceCalculator` (see .../BBQ.go).

In this repo everyone can add folder with his own idea of how given BBQ functionality should be tested.

For example:
1) ./ginkgo_me [folder](./ginkgo_me) includes an example of BDT approach using Ginkgo
2) ...

Note:
- Each folder should contain Readme file with sufficient explanation
- Please see closed PRs to read the discussions about added approaches (annotate -> open in github -> ...)
- If you see/fix issues - please create PR to contribute to main repo (if this is too hard, write me PM)