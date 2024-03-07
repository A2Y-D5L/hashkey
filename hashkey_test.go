package hashkey

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	for _, test := range []struct {
		name    string
		s       string
		want    string
		wantErr bool
	}{
		{
			name:    "empty string",
			s:       "",
			want:    "cbf29ce484222325",
			wantErr: false,
		},
		{
			name:    "kenobi",
			s:       "Hello there.",
			want:    "892362dd056bbf55",
			wantErr: false,
		},
		{
			name:    "random utf-8",
			s:       ">jҚy#️⃣&󫷌_\\󰋉ѭ󖧅:셺ጦ🿃G˕ݞ",
			want:    "3757e1d2ea2b1546",
			wantErr: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := String(test.s)
			if (err != nil) != test.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("String() = %v, want %v", got, test.want)
			}
		})
	}
}

func ExampleString() {
	h, err := String("Hello there.")
	if err != nil {
		panic(err)
	}
	fmt.Println(h)
	// Output: 892362dd056bbf55
}

func TestUint64(t *testing.T) {
	for _, test := range []struct {
		name    string
		s       string
		want    uint64
		wantErr bool
	}{
		{
			name:    "empty string",
			s:       "",
			want:    0xcbf29ce484222325,
			wantErr: false,
		},
		{
			name:    "grevious",
			s:       "General Kenobi! You are a bold one.",
			want:    0x58d72f7088dae07,
			wantErr: false,
		},
		{
			name:    "random utf-8",
			s:       ">jҚy#️⃣&󫷌_\\󰋉ѭ󖧅:셺ጦ🿃G˕ݞ",
			want:    0x3757e1d2ea2b1546,
			wantErr: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := Uint64(test.s)
			if (err != nil) != test.wantErr {
				t.Errorf("Uint64() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("Uint64() = %v, want %v", got, test.want)
			}
		})
	}
}

func ExampleUint64() {
	h, err := Uint64("General Kenobi! You are a bold one.")
	if err != nil {
		panic(err)
	}
	fmt.Println(h)
	// Output: 400102347231833607
}
