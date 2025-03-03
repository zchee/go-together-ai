// Code generated by github.com/zchee/go-openapi-tools/cmd/oapi-generator. DO NOT EDIT.

package togetherai

// FileObject represents a model of FileObject.
type FileObject struct {
	Filename string `json:"filename,omitempty"`
	ID       string `json:"id,omitempty"`
	Object   string `json:"object,omitempty"`
	Size     int32  `json:"size,omitempty"`
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (f *FileObject) GetFilename() (ret string) {
	if f == nil {
		return ret
	}
	return f.Filename
}

// GetID returns the ID field value if set, zero value otherwise.
func (f *FileObject) GetID() (ret string) {
	if f == nil {
		return ret
	}
	return f.ID
}

// GetObject returns the Object field value if set, zero value otherwise.
func (f *FileObject) GetObject() (ret string) {
	if f == nil {
		return ret
	}
	return f.Object
}

// GetSize returns the Size field value if set, zero value otherwise.
func (f *FileObject) GetSize() (ret int32) {
	if f == nil {
		return ret
	}
	return f.Size
}
