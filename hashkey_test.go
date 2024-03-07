package hashkey

import "testing"

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
