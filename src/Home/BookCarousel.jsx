import React from 'react';
import { Image } from 'antd';

const BookImageDisplay = ({ book }) => {
  // Ensure there's at least one image in the array
  if (!book.files || book.files.length === 0) {
    return <div>No image available</div>;
  }

  return (
    <div
      style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        width: '70%',
        height: '100%',
      }}
    >
      <div
        style={{
          width: '1200px',
          height: '400px',
          border: '2px solid gray',
          borderRadius: '15px',
          boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
          overflow: 'hidden',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          padding: '10px',
        }}
        className="image-container"
      >
        <Image
          src={book.files[0]}
          width={380}
          height={380}
          style={{ borderRadius: '10px', objectFit: 'cover' }}
          alt="Book cover"
        />
      </div>
    </div>
  );
};

export default BookImageDisplay;