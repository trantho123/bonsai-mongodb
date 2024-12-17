db.products.insertMany([
    {
      name: "Cây Kim Ngân",
      price: 300000,
      description: "Kim Ngân là một loại cây ưa sống ở vùng đầm lầy, có nguồn gốc từ Trung Quốc. Cây Kim Ngân ngoài tự nhiên có thể cao đến 5 – 6 m. Tuy nhiên, khi được sử dụng làm cây để bàn trong văn phòng, Kim Ngân thường được cắt, tỉa với vóc dáng phù hợp, gốc đơn hoặc gốc thắt bím 3 – 5 thân.",
      quantity: 20,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-12-696x696.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 4.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9010"), name: "Phong thủy" },
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" }
      ]
    },
    {
      name: "Cây Kim Tiền",
      price: 250000,
      description: "Kim Tiền là một loại cây dạng bụi, thân và lá đều khá “mập mạp”, là loại cây dễ chăm sóc, không cần ánh sáng quá nhiều và thích hợp để trồng &nbsp;trong nhà. Không chỉ có tác dụng trang trí cho ngôi nhà, cây Kim Tiền còn giúp cung cấp oxi, thanh lọc không khí và đem lại nhiều tài lộc cho gia chủ.",
      quantity: 30,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-18-696x522.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 4.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" },
        { id: ObjectId("64c7efc8af1d72b47f6b9012"), name: "Nội thất" }
      ]
    },
    {
      name: "Cây Bàng Singapore",
      price: 200000,
      description: "Là một trong những loại cây văn phòng được ưa thích nhất hiện nay, Bàng Singapore (tên khoa học Ficus Lyrata) phát triển mạnh mẽ trong các rừng mưa nhiệt đới phía Tây châu Phi và được lai tạo, nhân giống ở Singapore nên vẫn được gọi là Bàng Singapore.",
      quantity: 25,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-29-696x686.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 3.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" },
        { id: ObjectId("64c7efc8af1d72b47f6b9012"), name: "Nội thất" }
      ]
    },
    {
      name: "Cây Lưỡi Hổ",
      price: 400000,
      description: "Cây Lưỡi Hổ (tên khoa học: Sansevieria Trifasciata), họ Măng tây, có nguồn gốc từ châu Phi và có nhiều tên gọi khác nhau: cây lưỡi cọp, cây hổ vĩ mép vàng,… Nhưng chung quy lại, cây lưỡi hổ đều có đặc điểm là cây mọc thành bụi, lá mọng nước và lá có 2 màu vàng – xanh hoặc trắng xám – xanh xen kẽ.",
      quantity: 15,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-8-696x928.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 4.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9010"), name: "Phong thủy" },
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" }
      ]
    },
    {
      name: "Cây Hạnh Phúc",
      price: 150000,
      description: "Cây Hạnh Phúc (Radermachera sinica) là loại cây được phát hiện lần đầu trong các rừng mưa nhiệt đới ở Đông Nam Á và Trung Quốc. Cây dạng thân gỗ, tán lá dày và xanh mướt, thích hợp để trang trí &nbsp;trong nhà giúp tăng hòa khí và cải thiện không gian sống. Cây có thể cao tới 1,5 – 2,5 m, tuy nhiên, tùy vào mục đích sử dụng và vị trí trưng bày mà người ta có thể cắt tỉa cây với chiều cao phù hợp.",
      quantity: 50,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-24-696x696.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 4.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" },
        { id: ObjectId("64c7efc8af1d72b47f6b9013"), name: "Quà tặng" }
      ]
    },
    {
      name: "Cây Hồng Môn",
      price: 500000,
      description: "Cây Hồng Môn (Anthurium) là một loại cây có nhiều ở vùng Trung và Nam Mĩ, còn được biết đến với những cái tên khác như: Vĩ Hoa Tròn, Buồm Đỏ,… Cây mọc dạng bụi, lá và hoa có hình tim. Hoa Hồng Môn thường có màu đỏ ngọc, cam hoặc màu hồng.",
      quantity: 10,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-11-696x696.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 4.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9012"), name: "Nội thất" },
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" }
      ]
    },
    {
      name: "Cây Kim Ngân",
      price: 300000,
      description: "Kim Ngân là một loại cây ưa sống ở vùng đầm lầy, có nguồn gốc từ Trung Quốc. Cây Kim Ngân ngoài tự nhiên có thể cao đến 5 – 6 m. Tuy nhiên, khi được sử dụng làm cây để bàn trong văn phòng, Kim Ngân thường được cắt, tỉa với vóc dáng phù hợp, gốc đơn hoặc gốc thắt bím 3 – 5 thân.",
      quantity: 20,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-12-696x696.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 4.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9010"), name: "Phong thủy" },
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" }
      ]
    },
    {
      name: "Cây Kim Tiền",
      price: 250000,
      description: "Kim Tiền là một loại cây dạng bụi, thân và lá đều khá “mập mạp”, là loại cây dễ chăm sóc, không cần ánh sáng quá nhiều và thích hợp để trồng &nbsp;trong nhà. Không chỉ có tác dụng trang trí cho ngôi nhà, cây Kim Tiền còn giúp cung cấp oxi, thanh lọc không khí và đem lại nhiều tài lộc cho gia chủ.",
      quantity: 30,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-18-696x522.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 4.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" },
        { id: ObjectId("64c7efc8af1d72b47f6b9012"), name: "Nội thất" }
      ]
    },
    {
      name: "Cây Bàng Singapore",
      price: 200000,
      description: "Là một trong những loại cây văn phòng được ưa thích nhất hiện nay, Bàng Singapore (tên khoa học Ficus Lyrata) phát triển mạnh mẽ trong các rừng mưa nhiệt đới phía Tây châu Phi và được lai tạo, nhân giống ở Singapore nên vẫn được gọi là Bàng Singapore.",
      quantity: 25,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-29-696x686.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 3.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" },
        { id: ObjectId("64c7efc8af1d72b47f6b9012"), name: "Nội thất" }
      ]
    },
    {
      name: "Cây Lưỡi Hổ",
      price: 400000,
      description: "Cây Lưỡi Hổ (tên khoa học: Sansevieria Trifasciata), họ Măng tây, có nguồn gốc từ châu Phi và có nhiều tên gọi khác nhau: cây lưỡi cọp, cây hổ vĩ mép vàng,… Nhưng chung quy lại, cây lưỡi hổ đều có đặc điểm là cây mọc thành bụi, lá mọng nước và lá có 2 màu vàng – xanh hoặc trắng xám – xanh xen kẽ.",
      quantity: 15,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-8-696x928.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 4.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9010"), name: "Phong thủy" },
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" }
      ]
    },
    {
      name: "Cây Hạnh Phúc",
      price: 150000,
      description: "Cây Hạnh Phúc (Radermachera sinica) là loại cây được phát hiện lần đầu trong các rừng mưa nhiệt đới ở Đông Nam Á và Trung Quốc. Cây dạng thân gỗ, tán lá dày và xanh mướt, thích hợp để trang trí &nbsp;trong nhà giúp tăng hòa khí và cải thiện không gian sống. Cây có thể cao tới 1,5 – 2,5 m, tuy nhiên, tùy vào mục đích sử dụng và vị trí trưng bày mà người ta có thể cắt tỉa cây với chiều cao phù hợp.",
      quantity: 50,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-24-696x696.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 4.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" },
        { id: ObjectId("64c7efc8af1d72b47f6b9013"), name: "Quà tặng" }
      ]
    },
    {
      name: "Cây Hồng Môn",
      price: 500000,
      description: "Cây Hồng Môn (Anthurium) là một loại cây có nhiều ở vùng Trung và Nam Mĩ, còn được biết đến với những cái tên khác như: Vĩ Hoa Tròn, Buồm Đỏ,… Cây mọc dạng bụi, lá và hoa có hình tim. Hoa Hồng Môn thường có màu đỏ ngọc, cam hoặc màu hồng.",
      quantity: 10,
      image: "https://bloganchoi.com/wp-content/uploads/2021/02/cay-canh-van-phong-11-696x696.jpg",
      created_at: new Date("2024-12-02T12:00:00Z"),
      updated_at: new Date("2024-12-02T12:00:00Z"),
      rating: 4.5,
      tags: [
        { id: ObjectId("64c7efc8af1d72b47f6b9012"), name: "Nội thất" },
        { id: ObjectId("64c7efc8af1d72b47f6b9011"), name: "Cây cảnh" }
      ]
    }
    // ... Tiếp tục thêm các sản phẩm khác tương tự
  ]);