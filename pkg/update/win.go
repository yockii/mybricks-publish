//go:build windows
// +build windows

package update

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"io"
	"net/http"
)

func update(url string, sum []byte) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	wc := writeSumCounter{
		hash: sha256.New(),
	}
	rsp, err := io.ReadAll(io.TeeReader(resp.Body, &wc))
	if err != nil {
		return err
	}
	if !bytes.Equal(wc.hash.Sum(nil), sum) {
		return errors.New("文件已损坏")
	}
	// 更新文件，原文件名为 manatee-publish.exe，进行更名备份后，再写入该文件中
	// 将当前运行目录下的manatee-publish.exe 更名为 manatee-publish.exe.bak
	//_ = os.Rename("manatee-publish.exe", "manatee-publish.exe.bak")
	// 将下载的文件写入当前运行目录下的manatee-publish.exe
	err, _ = fromStream(bytes.NewReader(rsp))
	//reader, _ :=
	//	zip.NewReader(bytes.NewReader(rsp), resp.ContentLength)
	//file, err := reader.Open("manatee-publish.exe")
	//if err != nil {
	//	return err
	//}
	//err, _ = fromStream(file)
	if err != nil {
		return err
	}
	return nil
}
