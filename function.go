package main

import "fmt"

func main(){
    var car = map[string]string{}
    car["name"] = "BWM"
    car["color"] = "Black"

    // buat 2 buah fungsi :
    // 1 => fungsi yang mengembalikan sebuah string
    // pada fungsi ini terjadi pengolahan kata sehingga menghasilkan kata : Mobil BMW berwarna Black

    // 2 => fungsi yang menampilkan hasil dari kembalian string
    // fungsi ini hanya bertugas untuk menampilkan kata

    // alur
    // simpan hasil dari return function kedalam sebuah variable message
	message := getCar(car)
    // tampilkan hasil dari variable message
	fmt.Println(message)

    // output => Mobil BMW berwarna Black

}

func getCar(car map[string]string) (message string){
	printMessage(car)
	return 
}

func printMessage(car map[string]string) {
	fmt.Println("Mobil " + car["name"] + " berwarna " + car["color"])
}