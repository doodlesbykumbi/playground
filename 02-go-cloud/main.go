// Command blobber saves and retrieves files to blob storage on File, GCP and AWS.
package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Define our input.
	cloud := flag.String("cloud", "", "Cloud storage to use")
	bucketName := flag.String("bucket", "go-cloud-bucket", "Name of bucket")
	download := flag.Bool("d", false, "Download vs upload")
	flag.Parse()
	if flag.NArg() != 2 {
		log.Fatal("Failed to provide file and where to put it")
	}
	file := flag.Arg(0)
	outputFile := flag.Arg(1)

	ctx := context.Background()
	// Open a connection to the bucket.
	b, err := setupBucket(context.Background(), *cloud, *bucketName)
	if err != nil {
		log.Fatalf("Failed to setup bucket: %s", err)
	}

	if !*download {
		// Prepare the file for upload.
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read file: %s", err)
		}

		w, err := b.NewWriter(ctx, outputFile, nil)
		if err != nil {
			log.Fatalf("Failed to obtain writer: %s", err)
		}
		_, err = w.Write(data)
		if err != nil {
			log.Fatalf("Failed to write to bucket: %s", err)
		}
		if err = w.Close(); err != nil {
			log.Fatalf("Failed to close: %s", err)
		}
	} else {
		r, err := b.NewReader(ctx, file, nil)
		if err != nil {
			log.Fatalf("Failed to obtain reader: %s", err)
		}

		f, err := os.OpenFile(outputFile, os.O_WRONLY | os.O_CREATE, 0666)
		defer f.Close()

		if err != nil {
			log.Fatalf("Failed to setup bucket reader: %s", err)
		}

		w := bufio.NewWriter(f)
		defer w.Flush()

		_, err = io.Copy(w, r)
		if err != nil {
			log.Fatalf("Failed to write download file: %s", err)
		}

	}
}
