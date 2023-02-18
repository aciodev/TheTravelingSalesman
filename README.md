# COMP620 Term Project

Run the benchmark
```bash
go test -bench=. -benchtime=100x -benchmem -count 10 | tee res.txt
```

Analyze the information
```bash
benchstat r_5.txt r_10.txt r_25.txt r_50.txt r_100.txt
```