package usecase

import (
	"reflect"
	"shoe_shop_server/core/entities"
	"testing"
)

func Test_parseOrderDetails(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    entities.OrderDetails
		wantErr bool
	}{
		{
			name: "Trường hợp chuẩn",
			args: args{
				data: "tôi muốn mua cuốn sách Đắc Nhân Tâm số lượng 2 gửi về địa chỉ 123 Nguyễn Du, Hà Nội, địa chỉ email example@gmail.com số điện thoại 0912345678",
			},
			want: entities.OrderDetails{
				NameBook:    "đắc nhân tâm",
				Quantity:    2,
				Address:     "123 nguyễn du, hà nội",
				Email:       "example@gmail.com",
				PhoneNumber: "0912345678",
				Index:       4,
			},
			wantErr: true,
		},
		{
			name: "Sử dụng từ khóa khác",
			args: args{
				data: "đặt sách Bố Già với số lượng 1 giao hàng tới số 45 Lê Lợi, TP.HCM, mail customer@example.com sđt 098765432",
			},
			want: entities.OrderDetails{
				NameBook:    "bố già",
				Quantity:    1,
				Address:     "số 45 lê lợi, tp.hcm",
				Email:       "customer@example.com",
				PhoneNumber: "098765432",
				Index:       4,
			},
			wantErr: false,
		},
		{
			name: "Thiếu email",
			args: args{
				data: "mua sách Harry Potter số lượng 3 gửi về địa chỉ 789 Trần Hưng Đạo, Đà Nẵng, điện thoại 0887654321",
			},
			want: entities.OrderDetails{
				NameBook:    "harry potter",
				Quantity:    3,
				Address:     "789 trần hưng đạo, đà nẵng",
				PhoneNumber: "0887654321",
				Index:       4,
			},
			wantErr: false,
		},
		{
			name: "Thiếu số lượng",
			args: args{
				data: "tôi muốn mua cuốn sách Sherlock Holmes gửi về địa chỉ 456 Lý Thường Kiệt, Huế, email sherlock@example.com",
			},
			want:    entities.OrderDetails{},
			wantErr: true,
		},
		{
			name: "Thiếu địa chỉ",
			args: args{
				data: "đặt sách Doraemon số lượng 5, email doraemon@example.com số điện thoại 0923456789",
			},
			want:    entities.OrderDetails{},
			wantErr: true,
		},
		{
			name: "Thiếu cả email và số điện thoại",
			args: args{
				data: "mua cuốn sách Tôi Thấy Hoa Vàng Trên Cỏ Xanh số lượng 1 gửi về địa chỉ 321 Nguyễn Trãi, Hải Phòng",
			},
			want:    entities.OrderDetails{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOrderDetails(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOrderDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOrderDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
