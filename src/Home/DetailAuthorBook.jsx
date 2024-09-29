import { useEffect, useState } from "react";
import axios from "axios";
import { Col, Image, Row } from "antd";
import styles from './DetailAuthorBook.module.css'; // Import CSS Module

function DetailAuthorBook({ authorBooName }) {
    const [author, setAuthor] = useState(null); // State to store author data
    const [loading, setLoading] = useState(true); // State to track loading status
    const [error, setError] = useState(null); // State to track errors

    useEffect(() => {
        const author = "Nam Cao";
        const fetchAuthorDetails = async () => {
            try {
                const response = await axios.get(
                    'http://127.0.0.1:8080/manager/author_book/details',
                    { params: { name: author } }
                );

                if (response.data.code === 0) {
                    setAuthor(response.data.body); // Set author data
                } else {
                    setError("Không thể lấy thông tin tác giả.");
                }
            } catch (err) {
                setError(err.message || "Đã xảy ra lỗi.");
            } finally {
                setLoading(false);
            }
        };

        fetchAuthorDetails(); // Call the function on component mount
    }, []);

    if (loading) {
        return <div>Đang tải...</div>;
    }

    if (error) {
        return <div>Lỗi: {error}</div>;
    }

    return (
        <div className={styles.container}>
            {author ? (
                <div>
                    <div className={styles.authorDetails}>
                        <Col style={{ marginLeft: '250px', marginTop: '50px' }} className={styles.column}>
                            <Image width={600} height={400} src={author.avatar} />
                        </Col>
                        <Col style={{ marginLeft: '100px', marginTop: '-50px' }} className={styles.column}>
                            <h1>Chi tiết Tác giả</h1>
                            <p style={{ fontSize: '18px', color: '#555', marginBottom: '10px' }}>
                                Tên: {author.name}</p>
                            <p style={{ fontSize: '18px', color: '#555', marginBottom: '10px' }}>
                                Ngày sinh: {author.birth_date}</p>
                            <p style={{ fontSize: '18px', color: '#555', marginBottom: '10px' }}>
                                Quốc tịch: {author.nationality}</p>
                            <p style={{
                                fontSize: '18px',
                                color: '#555',
                                marginBottom: '10px',
                                fontFamily: 'Georgia, serif', // Custom font
                                textShadow: '2px 2px 4px rgba(0, 0, 0, 0.2)', // Shadow effect
                                fontWeight: 'normal' // Optional, if you want a normal font-weight
                            }}>
                                <strong>Tiểu sử:</strong> {author.biography}
                            </p>

                        </Col>

                    </div>
                </div>

            ) : (
                <p>Không có thông tin tác giả.</p>
            )}
        </div>
    );
}

export default DetailAuthorBook;
