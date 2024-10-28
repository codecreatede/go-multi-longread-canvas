package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date: 2024-10-29


A multi long read canvas profiling golang application that allows you to scan for the long reads
using multiple patterns and remove them from the long reads. extension to the trimmomatic multi
pattern.


*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}

var (
	longread    string
	motiflooker string
)

var rootCmd = &cobra.Command{
	Use:  "longread",
	Long: "look for the matching patterns",
	Run:  joinFunc,
}

func init() {
	rootCmd.Flags().
		StringVarP(&longread, "longread", "L", "path to the long read file", "long read file to be checked")
	rootCmd.Flags().
		StringVarP(&motiflooker, "pattern", "P", "path to the file containing the patterns", "pattern file")
}

func joinFunc(cmd *cobra.Command, args []string) {
	type pacbiofileID struct {
		id string
	}
	type pacbiofileSeq struct {
		seq string
	}
	pacbioIDConstruct := []pacbiofileID{}
	pacbioSeqConstruct := []pacbiofileSeq{}

	fpacbio, err := os.Open(longread)
	if err != nil {
		log.Fatal(err)
	}
	Opacbio := bufio.NewScanner(fpacbio)
	for Opacbio.Scan() {
		line := Opacbio.Text()
		if strings.HasPrefix(string(line), "@") {
			pacbioIDConstruct = append(pacbioIDConstruct, pacbiofileID{
				id: strings.ReplaceAll(strings.Split(string(line), " ")[0], "@", ""),
			})
		}
		if strings.HasPrefix(string(line), "A") || strings.HasPrefix(string(line), "T") ||
			strings.HasPrefix(string(line), "G") ||
			strings.HasPrefix(string(line), "C)") {
			pacbioSeqConstruct = append(pacbioSeqConstruct, pacbiofileSeq{
				seq: string(line),
			})
		}
	}

	patterns := []string{}
	pOpen, err := os.Open(motiflooker)
	if err != nil {
		log.Fatal(err)
	}
	pRead := bufio.NewScanner(pOpen)
	for pRead.Scan() {
		line := pRead.Text()
		patterns = append(patterns, string(line))
	}

	fID := []string{}
	fSeq := []string{}

	for i := range pacbioIDConstruct {
		fID = append(fID, pacbioIDConstruct[i].id)
		fSeq = append(fSeq, pacbioSeqConstruct[i].seq)
	}

	// in the previous single pattern only one pattern can be defined. I implemented this
	// for this multi pattern, with additional data structure so that it remvoes them.

	pacbiofinal := []string{}
	for i := range len(patterns) - 1 {
		for j := range len(fSeq) {
			start := strings.Index(fSeq[j], patterns[i])
			end := strings.Index(fSeq[j], patterns[i]) + len(patterns[i])
			if start == -1 {
				return
			}
			sliceIter := fSeq[j][:start] + fSeq[j][end:]
			start1 := strings.Index(sliceIter, patterns[i+1])
			end1 := strings.Index(sliceIter, patterns[i+1]) + len(patterns[i+1])
			sliceIterFinal := sliceIter[:start1] + sliceIter[end1:]
			pacbiofinal = append(pacbiofinal, sliceIterFinal)
		}
	}

	file, err := os.Create("canvased.fastq")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for i := range fID {
		file.WriteString("@" + fID[i] + "\n" + pacbiofinal[i] + "\n")
	}
	fmt.Println("The canavssed reads have been written")

}
