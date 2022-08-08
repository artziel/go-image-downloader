package ImageDownloader

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"strings"

	_ "golang.org/x/image/webp"

	"net/http"
	"os"
)

const (
	mimeJPEG = "image/jpeg"
	mimeGIF  = "image/gif"
	mimePNG  = "image/png"
	mimeWEBP = "image/webp"
)

type Response struct {
	StatusCode int
	Mime       string
	Message    string
}

func DownloadImage(dest string, url string) (Response, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Response{Message: err.Error()}, err
	}

	req.Header = http.Header{
		"accept": {
			"image/avif,image/webp,image/jpeg,image/jpg,image/apng,image/png,image/*",
		},
		"accept-encoding": {"gzip, deflate, br"},
		"accept-language": {"es-MX,es;q=0.9,en-US;q=0.8,en;q=0.7,es-419;q=0.6"},
		"user-agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36"},
	}

	resp, err := client.Do(req)
	if err != nil {
		return Response{StatusCode: resp.StatusCode, Message: err.Error()}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return Response{StatusCode: resp.StatusCode, Message: ErrServerError.Error()}, ErrServerError
	}

	cType := strings.TrimSpace(resp.Header.Get("Content-Type"))
	if cType != mimeGIF && cType != mimeJPEG && cType != mimePNG && cType != mimeWEBP {
		return Response{StatusCode: 415, Message: ErrInvalidDataType.Error()}, ErrInvalidDataType
	}

	out, err := os.Create(dest)
	if err != nil {
		return Response{}, err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return Response{}, err
	}

	return Response{StatusCode: 200, Mime: cType}, nil
}
