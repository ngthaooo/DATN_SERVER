class MessageParser {
    constructor(actionProvider) {
        this.actionProvider = actionProvider;
    }

    parse(message) {
        const lowerCaseMessage = message.toLowerCase();

        if (lowerCaseMessage.includes('xin chào')) {
            this.actionProvider.handleDefault();
        } else if (
            lowerCaseMessage.includes('đơn hàng nào đang giao') || 
            lowerCaseMessage.includes('đơn hàng của tôi')
        ) {
            this.actionProvider.handleCheckOrder();  // Gọi hàm kiểm tra đơn hàng đang giao
        }else if(
            lowerCaseMessage.includes('giờ làm việc của cửa hàng')
        ){
            this.actionProvider.giolamViec()
        } else {
            this.actionProvider.handleSearch(message);  // Pass the original message
        }
    }
}

export default MessageParser;
