package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
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
