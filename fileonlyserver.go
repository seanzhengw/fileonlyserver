package fileonlyserver

import (
	"net/http"
	"os"
)

type fileOnlyServer struct {
	fs http.FileSystem
}

func (fos fileOnlyServer) Open(name string) (http.File, error) {
	f, err := fos.fs.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, os.ErrNotExist
	}
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}

// Serve return a handler that serves HTTP requests
// with the contents of the file system rooted at root without directory listings.
//
// Use like http.FileServer:
//
//     http.Handle("/", staticserver.Serve(http.Dir("/tmp")))
//
func Serve(root http.FileSystem) http.Handler {
	return http.FileServer(fileOnlyServer{root})
}
