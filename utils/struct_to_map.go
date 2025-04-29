package utils

import (
	"fmt"
	"reflect"
)

// StructToMapString mengubah sebuah struct menjadi map dengan key dan value berupa string.
// Fungsi ini akan memetakan field-field dari struct ke dalam map dengan nama field sebagai key dan
// nilai field sebagai value. Nama field akan diambil dari tag struct yang diberikan,
// atau jika tag tidak ada, nama field itu sendiri yang digunakan.
// Fungsi hanya akan memetakan field yang diekspos (exported) dan memiliki tipe yang dapat di-interfacing.
//
// Parameter:
//   - input: Struct yang ingin diubah menjadi map. Dapat berupa pointer ke struct atau langsung struct.
//   - tag: Nama tag struct yang digunakan untuk mengekstrak nama field. Jika tidak ada tag, nama field digunakan.
//
// Return:
//   - map[string]string: Map yang berisi nama field sebagai key dan nilai field sebagai value dalam bentuk string.
//
// Catatan:
//   - Hanya field yang diekspos (public) yang akan diproses, sehingga field yang diawali dengan huruf kecil akan diabaikan.
//   - Nilai field akan dikonversi menjadi string menggunakan fmt.Sprintf.
//   - Jika input bukan merupakan struct atau pointer ke struct, maka fungsi akan mengembalikan map kosong.

func StructToMapString(input any, tag string) map[string]string {
	result := make(map[string]string)
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	// pastikan input adalah struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	if v.Kind() != reflect.Struct {
		return result
	}

	for i := range v.NumField() {
		field := v.Field(i)
		fieldType := t.Field(i)
		fieldName := fieldType.Tag.Get(tag)
		if fieldName == "" {
			fieldName = fieldType.Name //
		}
		// hanya ambil field exported
		if field.CanInterface() {
			result[fieldName] = fmt.Sprintf("%v", field.Interface())
		}
	}
	return result
}
