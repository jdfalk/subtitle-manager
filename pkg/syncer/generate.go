// file: pkg/syncer/generate.go

package syncer

//go:generate go run github.com/vektra/mockery/v2 --name=Transcriber --output=mocks --outpkg=mocks --filename=transcriber.go
//go:generate go run github.com/vektra/mockery/v2 --name=SubtitleExtractor --output=mocks --outpkg=mocks --filename=subtitle_extractor.go
