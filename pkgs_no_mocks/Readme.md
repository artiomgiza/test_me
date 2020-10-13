
### Intro

Here we have providers of some functionality.

Each provider define `Interface` - a commitment of what it able to do.

Pay attention, that how the provider implements this commitment doesn't matter for us.

---

### Mocks

The way we can play with the provided functionality in tests.


Mock generation cmd:

```
mockgen -source=PATH/FILENAME.go -destination=PATH/mock/FILENAME_mock.go
```

Example:
```
mockgen -source=pkgs/chicken-farm/chicken_farm.go -destination=pkgs/chicken-farm/mock/chicken_farm_mock.go
```