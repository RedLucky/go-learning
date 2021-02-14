package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	// "time"
)

func main() {

	/* =================VARIABLE=============================*/

	// komentar inline

	/*
		komentar multiline
	*/

	/*
		_ (underscore) merupakan variabel tampungan yang nantinya dihapus. variabel ini tidak bisa diprint
		var firstName string = "Lucky" // manifest typing
		lastName := "fernanda" //  type inference
		var fourth, fifth, sixth string = "empat", "lima", "enam" // multi variable manifest typing
		seventh, eight, ninth := "tujuh", "delapan", "sembilan" // multi variable type inference
		one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello" // difference value

		fmt.Printf("halo %s!\n", firstName)

		fmt.Println("Hai " + lastName)

		var message = `Nama saya "John Wick".
		Salam kenal.
		Mari belajar "Golang".`

		fmt.Println(message)
	*/

	/* ================== TIPE DATA ========================= */

	/* //numeric non decimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644 // Compiler secara cerdas akan menentukan tipe data variabel tersebut sebagai int32 (karena angka tersebut masuk ke cakupan tipe data int32)

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// numeric decimal
	var decimalNumber = 2.62
	decimal := 2.5

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimal) */

	/* ========================== KONDISI =================================== */

	/* var point = 8

	if point == 10 {
			fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
			fmt.Println("lulus")
	} else if point == 4 {
			fmt.Println("hampir lulus")
	} else {
			fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	var tempPoint = 8840.0
	// variabel Temporary Pada if
	if percent := tempPoint / 100; percent >= 100 {
			fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
			fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
			fmt.Printf("%.1f%s not bad\n", percent, "%")
	} */

	/* ================== Array ========================== */

	/* var names [4]string
		names[0] = "trafalgar"
		names[1] = "d"
		names[2] = "water"
		names[3] = "law"
		fmt.Println(names[0], names[1], names[2], names[3])

		var fruits = [4]string{"apple", "grape", "banana", "melon"}
		fmt.Println("Jumlah element \t\t", len(fruits))
		fmt.Println("Isi semua element \t", fruits)

		// tidak di deklarasi jumlah elemen
		var numbers = [...]int{2, 3, 2, 4, 3}
		fmt.Println("data array \t:", numbers)
		fmt.Println("jumlah elemen \t:", len(numbers))

		// 2 dimensi
		var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
		var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
		fmt.Println("numbers1", numbers1)
		fmt.Println("numbers2", numbers2)

		for i := 0; i < len(fruits); i++ {
	    fmt.Printf("elemen fruits %d : %s\n", i, fruits[i])
		}

		// menggunakan range
		for i, fruit := range fruits {
	    fmt.Printf("elemen %d : %s\n", i, fruit)
		}

		// ketika i tidak diperlukan. maka ganti dengan _ (underscore)
		for _, fruit := range fruits {
	    fmt.Printf("nama buah : %s\n", fruit)
		}

		// ketika fruit tidak diperlukan. maka ganti dengan _ (underscore)
		for i, _ := range fruits {
	    fmt.Printf("nama buah : %d\n", i)
		} */

	/* =========================== SLICE ============================== */
	/* 	var fruitsA = []string{"apple", "grape"}      // slice
	   	var fruitsB = [2]string{"banana", "melon"}    // array
	   	fmt.Println("nama buah", fruitsA)
	   	fmt.Println("nama buah", fruitsB)

	   	var fruitsC = []string{"apple", "grape", "banana", "melon"}
	   	var newFruits = fruitsC[0:2]
	   	newFruitsB := fruitsC[:]

	   	fmt.Println(newFruits) // ["apple", "grape"]
	   	fmt.Println(newFruitsB) // ["apple", "grape"] */

	/* ============================ MAP =============================== */
	/* 	var months map[string]int
	   	months = map[string]int{}

	   	months["jan"] = 50
	   	months["feb"] = 40
	   	fmt.Println("jan", months["jan"]) // jan 50
	   	fmt.Println("mei", months["mei"]) // mei 0

	   	var months1 = map[string]int{"jan": 50, "februari": 40}
	   	fmt.Println(months1) // mei 0

	   	var chickens = map[string]int{
	   		"jan": 50,
	   		"feb": 40,
	   		"mar": 34,
	   		"apr": 67,
	   	}

	   	for key, val := range chickens {
	   		fmt.Println(key, "  \t:", val)
	   	} */

	/* =========================== FUNGSI =============================== */
	/* var names = []string{"Lucky", "Fernanda"}
	printMessage("helo", names)

	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	var area2, circumference2 = calculate2(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f \n", area2)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference2)

	var avg = calculate3(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	// pakai slice
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	avg = calculate3(numbers...)
	msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	yourHobbies("wick", "sleeping", "eating")

	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...) */

	/* ============================== FUNGSI VARIADIC ==============================*/
	/* var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
				switch {
				case i == 0:
						max, min = e, e
				case e > max:
						max = e
				case e < min:
						min = e
				}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max) */

	/* =============================== FUNGSI SEBAGAI PARAMETER ===========================*/
	/* var data = []string{"wick", "jason", "ethan"}
	dataContainsO := filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLenght5 = filter(data, func(each string) bool {
			return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [jason ethan] */

	/* =================================== POINTER =============================== */
	/* 
		- Pointer adalah reference atau alamat memori. Nilai default variabel pointer adalah nil (kosong)
		- Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand (&) tepat sebelum nama variabel. 
			Metode ini disebut dengan referencing.
		- Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variabel. 
			Metode ini disebut dengan dereferencing.
	*/
	/* number4 := 4
	var number5 *int = &number4

	fmt.Println("Number4 (value) :", number4) // 4
	fmt.Println("Number4 (address on memory) :" , &number4) // 0xc00018c008

	fmt.Println("Number4 (value) :", *number5) // 4
	fmt.Println("Number4 (address on memory) :" , number5) // 0xc00018c008


	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address on memory) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address on memory) :", numberB)

	fmt.Println("-----------------------------------")

	numberA = 5

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	number := 15
	fmt.Println("before :", number) // 15

	change(&number, 10)
  fmt.Println("after  :", number) // 10 */

	/* ========================================== STRUCT ================================== */
	/* 
	- Go tidak memiliki class yang ada di bahasa-bahasa strict OOP lain. 
		Tapi Go memiliki tipe data struktur yang disebut dengan Struct.

	- Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu. 
		Property dalam struct, tipe datanya bisa bervariasi. 
		Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.

	- Dari sebuah struct, kita bisa buat variabel baru, yang memiliki atribut sesuai skema struct tersebut. 
		Kita sepakati dalam buku ini, variabel tersebut dipanggil dengan istilah object atau object struct. 
	*/

	var john student
	john.name = "john atan"
	john.grade = 3


	fmt.Println("name  :", john.name)
	fmt.Println("grade :", john.grade)

}

// function standard
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// jika semua parameter type data sama, tipe data bisa dideklarasikan di akhir params
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func randomWithRange2(min int, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

// multiple return
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

// nilai balik bisa didefinisikan di-awal
func calculate2(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}

// variadic function (parameter tak terbatas)
func calculate3(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func filter(data []string, callback func(string) bool) (results []string)  {
		for _, each := range data {
			if filtered := callback(each); filtered{
				results = append(results, each)
			}
		}
		return
}

type filterCallback func(string) bool

func filterAlias(data []string, callback filterCallback) (results []string)  {
	for _, each := range data {
		if filtered := callback(each); filtered{
			results = append(results, each)
		}
	}
	return
}

func change(original *int, value int) {
	*original = value
}

type student struct {
	name string
	grade int
}
