package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/config"
  "github.com/aws/aws-sdk-go-v2/credentials"
  "github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	accountid := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	bucket := os.Getenv("CLOUDFLARE_BUCKET")
	key := os.Getenv("CLOUDFLARE_KEY")
	secret := os.Getenv("CLOUDFLARE_SECRET")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
    config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")),
    config.WithRegion("auto"),
  )
  if err != nil {
    log.Fatal(err)
  }

  client := s3.NewFromConfig(cfg, func(o *s3.Options) {
      o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountid))
  })

  listObjectsOutput, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
    Bucket: &bucket,
  })
  if err != nil {
    log.Fatal(err)
  }

  file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalf("failed to open file, %v", err)
	}
	defer file.Close()

  uploadInput := &s3.PutObjectInput{
  	Bucket: aws.String(bucket),
  	Key: aws.String(file.Name()),
  	Body: file,
  }

	_, err = client.PutObject(context.TODO(), uploadInput)

	if err != nil {
		log.Fatal("failed to upload file", err)
	}


  for _, object := range listObjectsOutput.Contents {
    obj, _ := json.MarshalIndent(object, "", "\t")
    fmt.Println(string(obj))
  }


  listBucketsOutput, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
  if err != nil {
    log.Fatal(err)
  }

  for _, object := range listBucketsOutput.Buckets {
    obj, _ := json.MarshalIndent(object, "", "\t")
    fmt.Println(string(obj))
  }

}
