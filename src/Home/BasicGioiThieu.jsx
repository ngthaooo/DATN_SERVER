import React from 'react';

const NhaNamProfile = () => {
    return (
        <div style={{ margin: '0 auto', padding: '2rem', maxWidth: '960px' }}>
            <div style={{
                background: 'linear-gradient(to right, #3b82f6, #8b5cf6)',
                color: 'white',
                padding: '1.5rem',
                borderRadius: '0.5rem',
                boxShadow: '0 4px 10px rgba(0, 0, 0, 0.1)',
                marginBottom: '2rem'
            }}>
                <h1 style={{ fontSize: '2rem', fontWeight: 'bold', marginBottom: '1rem' }}>Ts Shop</h1>
                <p style={{ fontSize: '1.125rem' }}>
                    Công ty Cổ phần Văn hóa và Truyền thông Ts shop, gia nhập thị trường sách Việt Nam vào tháng 2 năm 2005 với tác phẩm "Balzac và cô bé thợ may Trung hoa" của Đới Tư Kiệt.
                </p>
            </div>

            <div style={{ marginBottom: '2rem' }}>
                <Section title="GIẢI THƯỞNG CHO DOANH NGHIỆP">
                    <ListItem>
                        2016: Bằng khen của Bộ trưởng Bộ TT&TT cho tập thể đạt thành tích xuất sắc trong phong trào thi đua trong lĩnh vực xuất bản in và phát hành.
                    </ListItem>
                </Section>

                <Section title="GIẢI THƯỞNG CHO CÁ NHÂN">
                    <ListItem>
                        2018: Ngài Étienne Rolland-Piègue, Tham tán hợp tác và hoạt động văn hóa Đại sứ quán Pháp tại Việt Nam, đã trao tặng Huân chương Hiệp sĩ Văn học và Nghệ thuật cho ông Nguyễn Nhật Anh, Tổng Giám đốc Công ty Cổ phần Văn hóa và Truyền thông Ts shop.
                    </ListItem>
                </Section>

                <Section title="GIẢI THƯỞNG CHO SÁCH">
                    <h3 style={{ fontSize: '1.25rem', fontWeight: '600', marginTop: '1.5rem', marginBottom: '0.5rem', color: '#4338ca' }}>Giải Sách Hay 2022</h3>
                    <ListItem>

                        Năm 2022, Ts Shop tự hào nhận Giải Sách Hay cho tác phẩm "Hành trình về phương Đông" của tác giả Nguyễn Nhật Anh. Tác phẩm này không chỉ ghi dấu ấn mạnh mẽ trong lòng độc giả mà còn khẳng định vị trí của Ts Shop trong ngành xuất bản Việt Nam. Cuốn sách kể về những trải nghiệm kỳ diệu và sâu sắc của một người trẻ tuổi trong hành trình khám phá bản thân và thế giới xung quanh. Sự kết hợp tinh tế giữa văn phong lôi cuốn và thông điệp sâu sắc đã mang lại cho độc giả những phút giây thư giãn và suy ngẫm. Giải thưởng này một lần nữa khẳng định cam kết của Ts Shop trong việc cung cấp những tác phẩm chất lượng, góp phần làm phong phú thêm văn hóa đọc trong cộng đồng.

                        Ngoài ra, năm 2022 cũng ghi nhận thành công của Ts Shop với việc xuất bản tác phẩm "Mùa hè của những giấc mơ" của tác giả Lê Minh. Tác phẩm này đã nhận được Giải Sách Hay với nội dung sâu sắc về tình bạn và những giá trị cuộc sống. Cuốn sách đã thu hút được đông đảo bạn đọc và được giới thiệu rộng rãi tại nhiều hội sách lớn trong cả nước. Sự thành công này không chỉ nâng cao thương hiệu của Ts Shop mà còn thể hiện sự đa dạng trong các thể loại sách mà công ty đang phát triển.
                    </ListItem>

                    <h3 style={{ fontSize: '1.25rem', fontWeight: '600', marginTop: '1.5rem', marginBottom: '0.5rem', color: '#4338ca' }}>Giải thưởng của Hội Nhà văn Việt Nam</h3>
                    <ListItem>Giải Văn học dịch, năm 2005: Cuộc đời của Pi (Yann Martel, do Trịnh Lữ dịch).</ListItem>
                    <ListItem>Giải Văn học dịch, năm 2008: Tên tôi là Đỏ (Orhan Pamuk, do Phạm Viêm Phương và Huỳnh Kim Anh dịch).</ListItem>
                    <ListItem>Giải Văn học dịch, năm 2016: Lâu đài sói (Hilary Mantel, do Nguyễn Chí Hoan dịch).</ListItem>
                </Section>
            </div>
        </div>
    );
};

const Section = ({ title, children }) => (
    <div>
        <h2 style={{ fontSize: '1.5rem', fontWeight: 'bold', marginBottom: '1rem', color: '#4b5563', borderBottom: '2px solid #4f46e5', paddingBottom: '0.5rem' }}>{title}</h2>
        <div style={{ marginLeft: '1rem' }}>{children}</div>
    </div>
);

const ListItem = ({ children }) => (
    <div style={{ marginBottom: '0.5rem', display: 'flex', alignItems: 'flex-start' }}>
        <div style={{ width: '0.5rem', height: '0.5rem', backgroundColor: '#4f46e5', borderRadius: '50%', marginTop: '0.5rem', marginRight: '0.5rem' }}></div>
        <p>{children}</p>
    </div>
);

const YearAwards = ({ year }) => {
    return (
        <div style={{ marginBottom: '1rem' }}>
            <h4 style={{ fontWeight: '600', fontSize: '1.125rem', color: '#4b5563' }}>{year}</h4>
            <ListItem>Giải thưởng mẫu cho năm {year}</ListItem>
        </div>
    );
};

export default NhaNamProfile;
