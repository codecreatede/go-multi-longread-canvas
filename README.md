# go-multi-longread-canvas

- multi-trimmomatic for the long reads, no heap memory allocation required, runs on goroutines and later adding the waitgroup.
- in this version, it supports the multiple patterns as in the last release of trimmomatic.
- parsed 100GB fastq file in few seconds. 
```
╭─gauavsablok@gauravsablok ~/Desktop/go/go--multi-longread-canvas ‹main●›
╰─$ go run main.go -h
look for the matching patterns

Usage:
  longread [flags]

Flags:
  -h, --help              help for longread
  -L, --longread string   long read file to be checked (default "path to the long read file")
  -P, --pattern string    pattern file (default "path to the file containing the patterns")
exit status 1
```
-detailed usage

```
╭─gauavsablok@gauravsablok ~/Desktop/go/go-multi-longread-canvas ‹main●›
╰─$ go run main.go -L sample.fastq -P pattern.txt
The canavssed reads have been written
exit status 1
╭─gauavsablok@gauravsablok ~/Desktop/go/go-multi-longread-canvas ‹main●›
╰─$ cat sample.fastq
@ERR10930361.1 magdelm64071_201030_115446/27
ATACTTTAAATTTTAGTTACTATTATTATTATTTAAAAAAAAAAAAAAAATTGAAAGTATATCCAAACTA
@ERR10930361.1 magdelm64071_201030_115446/27
ATACTTTAAATTTTAGTTACTATTATTATTATTTAAAAAAAAAAAAAAAATTGAAAGTATATCCAAACTAGCACTCAATTAATGCAAACAAT

╭─gauavsablok@gauravsablok ~/Desktop/go/go-multi-longread-canvas ‹main●›
╰─$ cat canvased.fastq
@ERR10930361.1
ATACTTTAAATTTTAGTTACTATTATTATTATTTTTTCCAAACTA
@ERR10930361.1
ATACTTTAAATTTTAGTTACTATTATTATTATTTTTTCCAAACTAGCACTCAATTAATGCAAACAAT

```

Gaurav Sablok
