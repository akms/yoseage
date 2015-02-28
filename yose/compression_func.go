package yose

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func MakeFile() (gw *gzip.Writer, tw *tar.Writer, file *os.File) {
	var (
		err error
	)
	if file, err = os.Create("/tmp/hoge.tar.gz"); err != nil {
		log.Fatal(err)
	}
	gw = gzip.NewWriter(file)
	tw = tar.NewWriter(gw)
	return
}

func CheckTarget(dirPaths []string) {
	var (
		fileinfo []os.FileInfo
		err      error
		gw       *gzip.Writer
		tw       *tar.Writer
		file     *os.File
	)
	for _, dirpath := range dirPaths {
		ChangeDir(dirpath)
		_,dirname := filepath.Split(dirpath)
		if fileinfo, err = ioutil.ReadDir(dirname); err != nil {
			log.Fatal(err)
		}
		gw, tw, file = MakeFile()
		CompressionFile(tw, fileinfo, dirname)
	}
	defer file.Close()
	defer gw.Close()
	defer tw.Close()
}

/*func walkFn(path string,info os.FileInfo,err error) error {
	fmt.Println(filepath.Base(info.Name()))
	return nil
}*/

func CompressionFile(tw *tar.Writer, fileinfo []os.FileInfo, dirname string) {
	var (
		err          error
		tmp_fileinfo []os.FileInfo
	)
	for _, file := range fileinfo {
		if file.IsDir() == true {
			if tmp_fileinfo, err = ioutil.ReadDir(file.Name()); err != nil {
				log.Fatal(err)
			}
			//		err = filepath.Walk(file.Name(),walkFn)
			ChangeDir(file.Name())
			filepath.Join(dirname,file.Name())
			CompressionFile(tw, tmp_fileinfo, dirname)
			dirname,_ = filepath.Split(dirname)
			ChangeDir("../")
		} else {
			tmpname := filepath.Join(dirname,file.Name())
			//fmt.Println(filepath.Base(file.Name()))
			fmt.Println(tmpname)
			body, _ := ioutil.ReadFile(tmpname)
			if err = tw.WriteHeader(&tar.Header{Mode: int64(file.Mode()), Size: file.Size(), ModTime: file.ModTime(), Name: tmpname}); err != nil {
				log.Fatal(err)
			}
			if _, err = tw.Write(body); err != nil {
				log.Fatal(err)
			}
		}
	}
}
