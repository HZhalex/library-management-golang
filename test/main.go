package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Tạo một handler đơn giản

	
// Ở đây: / là đường dẫn gốc (tức là khi người dùng vào http://localhost:8080/).
// Tham số thứ hai là một hàm xử lý (handler) nhận vào hai đối số:

// w http.ResponseWriter: cho phép ghi dữ liệu trả về cho trình duyệt.

// r *http.Request: chứa thông tin về request từ phía client (người dùng gửi lên).

// func(w http.ResponseWriter, r *http.Request) { ... }
// Đây là hàm ẩn danh (anonymous function) dùng làm handler xử lý mỗi khi có request tới đường /.

// fmt.Fprintf(w, "Chào mừng...")
// Ghi dòng "Chào mừng đến với Library Management System!" vào w, tức là gửi về cho trình duyệt của client.

// Giống như nói: "In dòng chữ này ra giao diện web."
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Chào mừng đến với Library Management System!")
	})


// http.ListenAndServe(":8080", nil)
// Khởi động HTTP server và bắt đầu lắng nghe tại cổng 8080.

// "8080" là port thường được dùng để chạy ứng dụng web thử nghiệm (có thể đổi thành 80 hoặc 3000...).

// nil: tức là dùng DefaultServeMux (trình quản lý route mặc định, nơi bạn đã đăng ký / ở bước trên).

// log.Fatal(...)
// Nếu server gặp lỗi (ví dụ port 8080 bị chiếm), log.Fatal() sẽ:

// In lỗi ra màn hình.

// Dừng chương trình ngay lập tức.
	// Khởi động server tại port 8080
	fmt.Println("🚀 Server đang chạy tại http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	/*
	Khái niệm: HTTP Handler là gì?
	Trong Go (và cả trong lập trình web nói chung), một HTTP handler là một 
	hàm hoặc đối tượng có thể xử lý một HTTP request.
	Nói đơn giản: Khi người dùng gửi yêu cầu đến server (truy cập URL), 
	handler sẽ quyết định server phải làm gì và trả lại gì.
	*/
}
