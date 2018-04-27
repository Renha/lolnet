package tofile

import (
	"fmt"
	"os"

	"github.com/lexfrei/lolnet"
)

var _ lolnet.Recipient = &tofile{}

type tofile struct {
	file string
}

// NewFileRecipient provides new recipient object
func NewFileRecipient(file string) *tofile {
	return &tofile{
		file: file,
	}
}

func (tofile) Remove() error {
	return nil
}

func (tf tofile) Add(bl *lolnet.Blood) error {
	f, err := os.Create(tf.file)
	if err != nil {
		return err
	}
	defer f.Close()
	var out string
	for _, i := range bl.IPs {
		out += fmt.Sprintf("%s\n", i)
	}
	for _, o := range bl.Nets {
		out += fmt.Sprintf("%s\n", o)
	}
	_, err = f.Write([]byte(out))
	if err != nil {
		return err
	}
	return nil
}

func (tf tofile) Update(bl *lolnet.Blood) error {
	if err := tf.Remove(); err != nil {
		return err
	}
	return tf.Add(bl)
}
