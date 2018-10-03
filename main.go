package main

import (
	_ "BeeMail/routers"
	"github.com/astaxie/beego"
)

func init() {
	// database.GetInstance()
}
func main() {
	// db, err := bolt.Open("my.db", 0600, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// err = db.Update(func(tx *bolt.Tx) error {
	// 	_, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
	// 	if err != nil {
	// 		return fmt.Errorf("create bucket: %s", err)
	// 	}
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = db.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("MyBucket"))
	// 	err := b.Put([]byte("answer"), []byte("42"))
	// 	return err
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = db.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("MyBucket"))
	// 	v := b.Get([]byte("answer"))
	// 	fmt.Printf("The answer is: %s\n", v)
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	beego.Run()
}
