package thumbnail

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"os/exec"
)

type ExtentionHelper interface {
	Extention(ext string, path string) image.Image
	WriteThumb(ext string, out *os.File, m image.Image)
	MakeThumbnail(path string, key string)
}

func Extention(ext string, path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
	}
	if ext == ".jpg" {
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		return img
	} else if ext == ".png" {
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		return img
	} else if ext == ".gif" {
		img, err := gif.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		return img
	}
	return nil
}

func MakeThumbnail(path string, destination string) error {
	// ext := filepath.Ext(path)
	// img := Extention(ext, path)
	// thumb := resize.Resize(300, 400, img, resize.Lanczos3)
	// out, err := os.Create(destination)
	// if err != nil {
	// 	return err
	// }
	// defer out.Close()
	// WriteThumb(ext, out, thumb)
	// return nil
	// ffmpeg -i image.JPG -vf scale=1920:1280 -q:v 1 output_1920x1280.jpg
	ffmpeg := exec.Command("ffmpeg", "-i", path, "-vf", "scale=300:400", "-q:v", "1", destination)
	if err := ffmpeg.Run(); err != nil {
		return err
	}
	return nil
}

//ffmpeg -i ceacd225-dc83-4bc3-a0fa-a1085b0a3907.mp4 -vf fps=5,scale=300:400 -compression_level 0 -y result/result.gif
func MakeThumbnailVideo(path string, destination string) error {
	ffmpeg := exec.Command("ffmpeg", "-i", path, "-vf", "fps=5,scale=300:400", "-compression_level", "0", "-y", destination)
	if err := ffmpeg.Run(); err != nil {
		return err
	}
	return nil
}

func WriteThumb(ext string, out *os.File, m image.Image) {
	if ext == ".jpg" {
		jpeg.Encode(out, m, nil)
	} else if ext == ".png" {
		png.Encode(out, m)
	} else if ext == ".gif" {
		gif.Encode(out, m, nil)
	}
}
