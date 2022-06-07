# Intro

Prolly trees is probabilistic merkle b-trees.

Prolly trees is similar to merkle trees:

* Verifiable
* Proofs of inclusion/exclusion
* Efficient and correct diff/sync/merge

It's also similar to b-trees:

* Efficient random reads/writes
* Efficient ordered scans
* Tightly controlled blocksize

Here's some properties of prolly trees:

|                                        | b-tree        | Prolly trees             |
| -------------------------------------- | ------------- | ------------------------ |
| random read                            | $\log_kn$     | $\log_kn$                |
| random write                           | $\log_kn$     | $(1+\frac{k}{w})\log_kn$ |
| Ordered scan of one item with size $z$ | $\frac{z}{k}$ | $\frac{z}{k}$            |
| Calculate diff of size $d$             | $n$           | $d$                      |
| Verification, proofs                   | $\times$      | $\surd$                  |
| Structured sharing                     | $\times$      | $\surd$                  |


