package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"go.seankhliao.com/goindex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	GoindexURL = "goindex.seankhliao.com:443"
)

func main() {
	ctx := context.Background()

	var goindexURL string
	var insecure, semver bool
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.StringVar(&goindexURL, "server", GoindexURL, "url of index server")
	fs.BoolVar(&insecure, "insecure", false, "use insecure")
	fs.BoolVar(&semver, "semver", false, "only list vX.Y.Z")
	fs.Parse(os.Args[1:])

	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	var opts []grpc.DialOption
	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		creds, err := credentials.NewClientTLSFromFile("/etc/ca-certificates/extracted/tls-ca-bundle.pem", "")
		if err != nil {
			log.Fatal().Err(err).Msg("get server cert")
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	cc, err := grpc.Dial(goindexURL, opts...)
	if err != nil {
		log.Fatal().Err(err).Msg("dial index")
	}
	client := goindex.NewGoindexClient(cc)

	args := fs.Args()
	if len(args) == 0 {
		log.Fatal().Msg("no subcommand passed")
	}
	switch args[0] {
	case "versions":
		for i, mod := range args[1:] {
			pv, err := client.Versions(ctx, &goindex.VersionsRequest{
				Project: mod,
				Semver:  semver,
			})
			if err != nil {
				log.Error().Err(err).Int("i", i).Str("mod", mod).Msg("get versions")
				continue
			}
			fmt.Println(pv.Project)
			for _, ir := range pv.Versions {
				fmt.Println("\t", ir.Path, "\t", ir.Version)
			}

		}
	}

}
