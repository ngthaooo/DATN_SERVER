package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

// AddressUserOrder represents the structure of the order details.
type AddressUserOrder struct {
	Index       int     `json:"index"`
	NameBook    string  `json:"name_book"`
	Quantity    int     `json:"quantity"`
	PriceBook   string  `json:"price_book"`
	Email       string  `json:"email"`
	Address     string  `json:"address"`
	PhoneNumber string  `json:"phone_number"`
	Amount      float64 `json:"amount"`
}

// extractDetails extracts book name or order details based on the input data.
func extractDetails(data string, index int) (interface{}, int) {
	// Define regex patterns to capture book name, quantity, address, email, and phone number
	rePrice := regexp.MustCompile(`(?i)cuốn sách (.+?) giá`)
	reAvailable := regexp.MustCompile(`(?i)cuốn sách (.+?) còn`)
	reAuthorBook := regexp.MustCompile(`(?i)tác giả cuốn sách (.+?) là`)

	// Try to extract the book name
	if match := rePrice.FindStringSubmatch(data); len(match) > 1 {
		return match[1], 1
	} else if match := reAvailable.FindStringSubmatch(data); len(match) > 1 {
		return match[1], 2
	} else if match := reAuthorBook.FindStringSubmatch(data); len(match) > 1 {
		return match[1], 3
	}

	// If no book name is found, continue to extract order details
	reBookName := regexp.MustCompile(`cuốn sách (.+?) số lượng`)
	reQuantity := regexp.MustCompile(`số lượng (\d+)`)
	reAddress := regexp.MustCompile(`địa chỉ (.+?) ,`)
	reEmail := regexp.MustCompile(`email (.+?) số`)
	rePhoneNumber := regexp.MustCompile(`số điện thoại[: ]+(\d+)`)

	order := AddressUserOrder{
		Index: index,
	}

	// Extract book name for the order
	if match := reBookName.FindStringSubmatch(data); len(match) > 1 {
		order.NameBook = match[1]
	}

	// Extract quantity
	if match := reQuantity.FindStringSubmatch(data); len(match) > 1 {
		quantity, _ := strconv.Atoi(match[1])
		order.Quantity = quantity
	}

	// Extract address
	if match := reAddress.FindStringSubmatch(data); len(match) > 1 {
		order.Address = match[1]
	}

	// Extract email
	if match := reEmail.FindStringSubmatch(data); len(match) > 1 {
		order.Email = match[1]
	}

	// Extract phone number
	if match := rePhoneNumber.FindStringSubmatch(data); len(match) > 1 {
		order.PhoneNumber = match[1]
	}

	// Return order as JSON string
	jsonOrder, _ := json.Marshal(order)
	return string(jsonOrder), 4
}

func main() {
	// Test cases for book name and order details
	testData := []string{
		"cuốn syyách 'Tôi và chúng ta' giá bao nhiêu?",
		"cuốn sácddh 'Tôi và chúng ta' còn không?",
		"tác giảd cuốn sách 'Tôi và chúng ta 88' là ai?",
		"tôi muốn mua cuốn sách ABCXY số lượng 2 gửi về địa chỉ Hà Nội Bắc Từ Liêm Nguyên Xá , địa chỉ email hhhhh@hmail.com số điện thoại 098143601",
	}

	for index, data := range testData {
		result, _ := extractDetails(data, index)

		switch v := result.(type) {
		case string:
			fmt.Println("Tên sách:", v)
		case AddressUserOrder:
			fmt.Println("Chi tiết đơn hàng:")
			orderJSON, _ := json.Marshal(v)
			fmt.Println(string(orderJSON))
		default:
			fmt.Println("Không tìm thấy thông tin phù hợp.")
		}
		fmt.Println()
	}
}
