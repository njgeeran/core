package utils

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//压缩文件夹&文件-不包含当前文件夹
func CompressDirNotIn(source_dir string,dest string) error {
	f,err := os.Open(source_dir)
	if err != nil {
		return err
	}
	info, err := f.Stat()
	if info.IsDir() {
		fileInfos, err := f.Readdir(-1)
		if err != nil {
			return err
		}
		fs := []*os.File{}
		for _, fi := range fileInfos {
			cf, err := os.Open(f.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			fs = append(fs, cf)
		}
		return Compress(dest,fs...)
	}else {
		return Compress(dest,f)
	}
}
//压缩文件夹&文件-包含当前文件夹
func CompressDir(source_dir string,dest string) error {
	f,err := os.Open(source_dir)
	if err != nil {
		return err
	}
	return Compress(dest,f)
}
//压缩
func Compress(dest string,files ...*os.File) error {
	d, _ := os.Create(dest)
	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	for _, file := range files {
		err := compress(file, "", tw)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}
func compress(file *os.File, prefix string, tw *tar.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压
func DeCompressZip(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		decodeName := ""
		if file.Flags == 0 {
			//如果标致位是0  则是默认的本地编码   默认为gbk
			i:= bytes.NewReader([]byte(file.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
			content,_:= ioutil.ReadAll(decoder)
			decodeName = string(content)
		}else{
			//如果标志为是 1 << 11也就是 2048  则是utf-8编码
			decodeName = file.Name
		}
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + decodeName
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}
func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}
func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}
