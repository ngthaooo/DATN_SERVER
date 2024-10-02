package usecase

import (
	"context"
	"math/rand"
	"regexp"
	errors "shoe_shop_server/common/error"
	"shoe_shop_server/common/log"
	"shoe_shop_server/core/domain"
	"shoe_shop_server/core/entities"
	"strconv"
	"strings"
	"time"
)

type TransPortUserCase struct {
	book     domain.RepositoryBook
	typeBook domain.RepositoryGenre
}

func NewTransPortUserCase(book domain.RepositoryBook, typeBook domain.RepositoryGenre) *TransPortUserCase {
	return &TransPortUserCase{
		book:     book,
		typeBook: typeBook,
	}
}

func (u *TransPortUserCase) GetInforMationForChatBot(ctx context.Context, text *entities.DataContent) (interface{}, errors.Error) {

	respWhenSearhEng := []interface{}{}

	nameBook, index := u.extractDetails(text.Content)
	if index > 0 {
		if index == 1 {
			book, err := u.book.GetBookByUserNameBook(ctx, nameBook.(string))
			if err != nil {
				log.Error(err, "error")
				return nil, errors.ErrSystem
			}
			if book != nil {
				respWhenSearhEng = append(respWhenSearhEng, &entities.PriceBook{
					Index:    1,
					Price:    book.Price,
					NameBook: nameBook.(string),
				})
			} else {
				orderInfor, _ := parseOrderDetails(text.Content)
				book, err := u.book.GetBookByUserNameBook(ctx, orderInfor.NameBook)
				if err != nil {
					log.Error(err, "error")
					return nil, errors.ErrSystem
				}
				if book == nil {
					respWhenSearhEng = append(respWhenSearhEng, entities.OrderDetails{
						Index: 41,
					})
				} else {
					respWhenSearhEng = append(respWhenSearhEng, entities.OrderDetails{
						Index:       4,
						NameBook:    orderInfor.NameBook,
						Quantity:    orderInfor.Quantity,
						Address:     orderInfor.Address,
						Email:       orderInfor.Email,
						PhoneNumber: orderInfor.PhoneNumber,
					})
					log.Infof("data ", respWhenSearhEng...)

				}

			}
		}

		if index == 2 {
			book, err := u.book.GetCountBookByUserNameBook(ctx, nameBook.(string))
			if err != nil {
				log.Error(err, "error")
				return nil, errors.ErrSystem
			}

			respWhenSearhEng = append(respWhenSearhEng, entities.CountBook{
				Index:    2,
				Quality:  book.Quantity,
				NameBook: book.Title,
			})
		}
		if index == 3 {
			book, err := u.book.GetAuthorBookByBook(ctx, nameBook.(string))
			if err != nil {
				log.Error(err, "error")
				return nil, errors.ErrSystem
			}
			if book != nil {
				respWhenSearhEng = append(respWhenSearhEng, &entities.AuthorBook{
					Index:       3,
					Author:      book.Name,
					Nationality: book.Nationality,
					BirthDate:   book.BirthDate,
				})
			}
		}

		if index == 5 {

			var typeBooks []*entities.ListBookByTypeBook
			listTypeBook, _ := u.book.GetListBookByTypeBookBotUseBot(ctx, nameBook.(string))

			// Kiểm tra xem có loại sách nào được tìm thấy không
			if len(listTypeBook) > 0 {
				// Trộn danh sách loại sách
				rand.Seed(time.Now().UnixNano())
				rand.Shuffle(len(listTypeBook), func(i, j int) {
					listTypeBook[i], listTypeBook[j] = listTypeBook[j], listTypeBook[i]
				})

				// Lấy số lượng ngẫu nhiên từ 2 đến 5
				numTypeBooks := rand.Intn(4) + 2 // Tạo số ngẫu nhiên từ 2 đến 5
				if numTypeBooks > len(listTypeBook) {
					numTypeBooks = len(listTypeBook) // Đảm bảo không lấy quá số lượng loại sách hiện có
				}

				// Chỉ lấy một số lượng ngẫu nhiên loại sách từ listTypeBook
				for i := 0; i < numTypeBooks; i++ {
					typeBooks = append(typeBooks, &entities.ListBookByTypeBook{
						Name: listTypeBook[i].Title,
					})
				}

				// Thêm phản hồi với danh sách loại sách đã chọn
				respWhenSearhEng = append(respWhenSearhEng, &entities.TypeBookResp{
					Index:              5,
					ListBookByTypeBook: typeBooks,
				})
			} else {
				// Nếu không tìm thấy loại sách nào
				respWhenSearhEng = append(respWhenSearhEng, &entities.TypeBookResp{
					Index: 1, // Index cho thông báo không tìm thấy
				})
			}
		}

	}
	return respWhenSearhEng, nil
}

func (u *TransPortUserCase) extractDetails(data string) (interface{}, int) {
	// Define regex patterns to capture book name, quantity, author, and book type
	rePrice := regexp.MustCompile(`(?i)cuốn sách (.+?) g`)
	reAvailable := regexp.MustCompile(`(?i)cuốn sách (.+?) c`) // số lượng
	reAuthorBook := regexp.MustCompile(`(?i)tác giả sách (.+?) l`)
	reTypeBook := regexp.MustCompile(`(?i)loại sách (.+)`) // tìm loại sách sau "loại sách"

	// Try to extract the book name
	if match := rePrice.FindStringSubmatch(data); len(match) > 1 {
		return match[1], 1
	} else if match := reAvailable.FindStringSubmatch(data); len(match) > 1 {
		return match[1], 2
	} else if match := reAuthorBook.FindStringSubmatch(data); len(match) > 1 {
		return match[1], 3
	} else if match := reTypeBook.FindStringSubmatch(data); len(match) > 1 {
		return match[1], 5 // trả về tên loại sách sau "loại sách"
	}

	return 0, 0
}

func parseOrderDetails(data string) (entities.OrderDetails, error) {
	order := entities.OrderDetails{}

	// Chuẩn hóa dữ liệu đầu vào
	data = strings.ToLower(data)
	data = strings.ReplaceAll(data, "  ", " ")

	// Cải thiện regex patterns
	patterns := map[string]*regexp.Regexp{
		"bookName":    regexp.MustCompile(`(?:mua|đặt)(?:\s+cuốn)?(?:\s+sách)?\s+(.+?)(?:\s+(?:số lượng|với số lượng|gửi về|giao hàng|tới địa chỉ))`),
		"quantity":    regexp.MustCompile(`(?:số lượng|với số lượng)\s+(\d+)`),
		"address":     regexp.MustCompile(`(?:gửi về|giao hàng tới|tới địa chỉ)\s+(.+?)(?:\s*(?:email|mail|số điện thoại|sđt|$))`),
		"email":       regexp.MustCompile(`(?:mail|email)\s+(\S+@\S+\.\S+)`),
		"phoneNumber": regexp.MustCompile(`(?:số điện thoại|sđt)\s+(\d{8,12})`),
	}

	// Xử lý từng trường thông tin
	for field, pattern := range patterns {
		if match := pattern.FindStringSubmatch(data); len(match) > 1 {
			switch field {
			case "bookName":
				order.NameBook = strings.TrimSpace(match[1])
			case "quantity":
				quantity, err := strconv.Atoi(match[1])
				if err != nil {
					return order, errors.NewResourceNotFoundWithCode("không thể xác định số lượng sách")
				}
				order.Quantity = quantity
			case "address":
				order.Address = strings.TrimSpace(match[1])
			case "email":
				order.Email = strings.TrimSpace(match[1])
			case "phoneNumber":
				order.PhoneNumber = strings.TrimSpace(match[1])
			}
		}
	}

	// Kiểm tra tính hợp lệ của đơn hàng
	if order.NameBook == "" {
		return order, errors.NewResourceNotFoundWithCode("không tìm thấy tên sách trong đơn hàng")
	}
	if order.Quantity == 0 {
		return order, errors.NewResourceNotFoundWithCode("số lượng sách không hợp lệ")
	}
	if order.Address == "" {
		return order, errors.NewResourceNotFoundWithCode("không tìm thấy địa chỉ giao hàng")
	}
	if order.Email == "" && order.PhoneNumber == "" {
		return order, errors.NewResourceNotFoundWithCode("cần cung cấp ít nhất một phương thức liên hệ (email hoặc số điện thoại)")
	}

	order.Index = 4
	return order, nil
}
