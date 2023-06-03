# go_study

> https://www.youtube.com/watch?v=9fYqg6uo-UU&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&ab_channel=HiteshChoudhary

## Memory Management

---

> new()
> Allocate Memory but not initialize
> zeroed storage (cannot input any datas)

> make()
> Allocate Memory and initialize
> non-zeroed storage (can input any datas)

---

GC(Garbage Collector) happens automatically. got errored and reload entire GC -> pf sky-looped

- runtime

1. GO GC variable sets the initial garbage collection target percentage. the collection is triggered when the ratio of freshly allocated data
   to live data remaining the previous collection reaches this percentage.
