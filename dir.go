package ndash

import (
	"os"
)

func Mkdir(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		//fmt.Println("stat temp dir error,maybe is not exist, maybe not")
		if os.IsNotExist(err) {
			//fmt.Println("temp dir is not exist")
			err := os.MkdirAll(dir, os.ModePerm)
			//if err != nil {
			//	//fmt.Printf("mkdir failed![%v]\n", err)
			//	return err
			//}
			return err
		}

		//fmt.Println("stat file error")
		return nil
	}

	return err
}
