package yose

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	gw   *gzip.Writer
	tw   *tar.Writer
	file *os.File
)

func MakeFile() (*gzip.Writer, *tar.Writer, *os.File) {
	var (
		hostname string
		err      error
	)
	if hostname, err = os.Hostname(); err != nil {
		log.Fatal(err)
	}
	hostname = "/tmp/" + hostname + ".tar.gz"
	if file, err = os.Create(hostname); err != nil {
		log.Fatal(err)
	}
	gw = gzip.NewWriter(file)
	tw = tar.NewWriter(gw)
	return gw, tw, file
}

func MatchTarget(name string) bool {
	str := strings.Fields(`^lost\+found$ ^proc$ ^sys$ ^dev$ ^mnt$ ^var$ ^run$`)
	for _, s := range str {
		dirRegexp := regexp.MustCompile(s)
		if dirRegexp.MatchString(name) {
			return false
		}
	}
	return true
}

func CheckTarget(dirpath string) {
	var (
		tmpfileinfo, fileinfo []os.FileInfo
		err                   error
	)
	ChangeDir(dirpath)
	if tmpfileinfo, err = ioutil.ReadDir(dirpath); err != nil {
		log.Fatal(err)
	}
	for _, info := range tmpfileinfo {
		if MatchTarget(info.Name()) {
			fileinfo = append(fileinfo, info)
		}
	}
	_, dirname := filepath.Split(dirpath)
	gw, tw, file = MakeFile()
	CompressionFile(tw, fileinfo, dirname)
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
		err            error
		tmp_fileinfo   []os.FileInfo
		change_dirpath string
	)
	for _, infile := range fileinfo {
		if infile.IsDir() == true {
			if tmp_fileinfo, err = ioutil.ReadDir(infile.Name()); err != nil {
				log.Fatal(err)
			}
			/*if err = infile.Walk(file.Name(),walkFn); err != nil {
				log.Fatal(err)
			}*/
			change_dirpath, _ = filepath.Abs(infile.Name())
			fmt.Println(change_dirpath)
			ChangeDir(change_dirpath)
			dirname = filepath.Join(dirname, infile.Name())
			CompressionFile(tw, tmp_fileinfo, dirname)
			dirname, _ = filepath.Split(dirname)
			change_dirpath, _ = filepath.Split(change_dirpath)
			//fmt.Println(change_dirpath)
			ChangeDir(change_dirpath)
		} else {
			if infile.Mode()&os.ModeSymlink == os.ModeSymlink {
				tmpname := filepath.Join(dirname, infile.Name())
				evalsym, _ := os.Readlink(infile.Name())
				hdr, _ := tar.FileInfoHeader(infile, evalsym)
				hdr.Typeflag = tar.TypeSymlink
				hdr.Name = tmpname
				if err = tw.WriteHeader(hdr); err != nil {
					log.Fatal(err)
				}
			} else {
				tmpname := filepath.Join(dirname, infile.Name())
				fmt.Println(tmpname)
				body, _ := ioutil.ReadFile(infile.Name())
				if err = tw.WriteHeader(&tar.Header{Mode: int64(infile.Mode()), Size: infile.Size(), ModTime: infile.ModTime(), Name: tmpname}); err != nil {
					log.Fatal(err)
				}
				if _, err = tw.Write(body); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
