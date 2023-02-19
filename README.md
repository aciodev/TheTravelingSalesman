# The Traveling Salesman
This repository features implementations of The Traveling Salesman Problem (TSP) in both sequential and parallel approaches. The goal is to compare and contrast:
* Pros/cons of the various approaches
* Method execution time
* Method memory consumption

# Graduate Research
* Master of Science, Computer Science.
* COMP620 - Computer Architecture
* California State University, Northridge. Spring 2023.

# Data format
Each file in the `data` folder represents an NxN grid of integers uniformly distributed in the range of 1 to 100. A node cannot be adjacent to itself, so the diagonal value is equal to -1.

To reduce variance in the results, we create a strongly connected graph to guarantee a threshold amount of work.

# Useful commands

Run the benchmark
```bash
go test -bench=. -benchtime=10x -benchmem -count 10 | tee res.txt
```

Analyze the information
```bash
benchstat res.txt
```