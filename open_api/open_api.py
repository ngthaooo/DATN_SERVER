from whoosh.index import create_in
from whoosh.fields import Schema, TEXT
from whoosh.qparser import QueryParser
import os
from unidecode import unidecode
import re

# Tạo schema cho chỉ mục
schema = Schema(question=TEXT(stored=True))

# Tạo thư mục để lưu chỉ mục
index_dir = "indexdir"
if not os.path.exists(index_dir):
    os.mkdir(index_dir)

# Tạo chỉ mục
ix = create_in(index_dir, schema)
writer = ix.writer()

# Danh sách các câu hỏi liên quan đến sách
questions = [
   "Giá của cuốn sách XYZ bao nhiêu?",
   "Giá của cuốn sách ABC bao nhiêu?",
   "Sách XYZ có bản e-book không?",
]

# Thêm câu hỏi vào chỉ mục (đã chuẩn hóa không có dấu)
for question in questions:
    normalized_question = unidecode(question).lower()
    writer.add_document(question=normalized_question)  # Thêm câu hỏi vào chỉ mục
writer.commit()  # Lưu các thay đổi

def search_query(query_str):
    # Chuẩn hóa chuỗi truy vấn đầu vào để loại bỏ dấu và chuyển thành chữ thường
    normalized_query_str = unidecode(query_str).lower()
    print(f"Normalized Query String: '{normalized_query_str}'")  # Debugging

    # Tìm tên sách trong câu hỏi
    book_name_patterns = [
        r'cuốn sách (.+?) bao nhiêu',
        r'cuốn sách (.+?) giá',
        r'sách (.+?) có bản e-book',
        r'tác giả của (.+?) là ai'
    ]

    book_name = None
    for pattern in book_name_patterns:
        match = re.search(pattern, normalized_query_str)
        if match:
            book_name = match.group(1).strip()
            break

    if book_name:
        print(f"Tên sách bạn đã hỏi: '{book_name}'")  # Debugging

        # Thay thế {tên_sách} trong câu hỏi mẫu với tên sách đã tìm thấy
        search_templates = [
            "Giá của cuốn sách {tên_sách} là bao nhiêu?",
            "Sách {tên_sách} có bản e-book không?",
            "Tác giả của cuốn sách {tên_sách} là ai?"
        ]

        results = []
        with ix.searcher() as searcher:
            for template in search_templates:
                # Chuẩn hóa truy vấn với tên sách thực tế
                normalized_query = unidecode(template.replace("{tên_sách}", book_name)).lower()
                print(f"Search Template: '{normalized_query}'")  # Debugging
                query = QueryParser("question", ix.schema).parse(normalized_query)
                search_results = searcher.search(query)
                results.extend([hit['question'] for hit in search_results])
                print(f"Results Found for '{normalized_query}': {[hit['question'] for hit in search_results]}")  # Debugging

        return results

    # Nếu không tìm thấy tên sách, thực hiện tìm kiếm thông thường
    with ix.searcher() as searcher:
        query = QueryParser("question", ix.schema).parse(normalized_query_str)
        results = searcher.search(query)
        return [hit['question'] for hit in results]

if __name__ == "__main__":
    query_str = input("Nhập từ khóa để tìm kiếm câu hỏi: ")
    results = search_query(query_str)
    if results:
        print("Kết quả tìm kiếm:")
        for result in results:
            print(f"- {result}")
    else:
        print("Không tìm thấy câu hỏi nào phù hợp.")
