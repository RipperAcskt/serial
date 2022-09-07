package serial

import (
	"fmt"
	"os"

	"github.com/tarm/serial"
)

func Read(name string, baund int) (string, error) {
	s, err := initPort(name, baund)
	if err != nil {
		fmt.Println(os.Stderr, "Error while openning port: %v-%v\n", name, err)
		return "", err
	}

	buf := make([]byte, 128)
	_, err = s.Read(buf)
	if err != nil {
		fmt.Println(os.Stderr, "Error while readding: %v-%v\n", name, err)
		return "", err
	}
	return string(buf), nil
}

func Write(name, text string, baund int) error {
	s, err := initPort(name, baund)
	if err != nil {
		fmt.Println(os.Stderr, "Error while openning port: %v\n", err)
		return err
	}

	_, err = s.Write([]byte(text))
	if err != nil {
		fmt.Println(os.Stderr, "Error while writting: %v-%v\n", name, err)
		return err
	}
	return nil

}

func initPort(name string, baund int) (*serial.Port, error) {
	c := &serial.Config{Name: name, Baud: baund}
	return serial.OpenPort(c)
}
