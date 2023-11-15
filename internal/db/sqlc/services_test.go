package sqlc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateService(t *testing.T) {
	createUser(t)

	providers, err := testStore.ListProviders(context.Background())
	require.NoError(t, err)

	categoryIDs, err := testStore.ListCategoryIDs(context.Background())
	require.NoError(t, err)

	/*
		Table categories:
		id | title
		---+----------------
		0  | Phụ kiện
		1  | Dinh dưỡng và thức ăn
		2  | Y tế và chăm sóc sức khỏe
		3  | Grooming
		4  | Đào tạo và huấn luyện
		5  | Khác
	*/

	// -----------------------------------------------------
	// --------------------Phụ Kiện-------------------------
	// -----------------------------------------------------
	// Service 1
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Thiết kế lồng chim theo yêu cầu",
		Description:       "Dịch vụ Thiết kế lồng chim theo yêu cầu mang đến trải nghiệm cá nhân hóa với việc tạo lồng chim độc đáo, tuân theo ý kiến và yêu cầu cụ thể của từng khách hàng. Chuyên gia thiết kế đảm bảo rằng mỗi chiếc lồng không chỉ phản ánh sở thích cá nhân mà còn tối ưu hóa môi trường sống cho chim cảnh.",
		Price:             199_000,
		ImagePath:         "/static/img/services_img/phu-kien.jpg",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 2
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Sửa chữa & Nâng cấp lồng chim",
		Description:       "Dịch vụ Nâng cấp lồng chim cũ mang lại sự mới mẻ và hiện đại cho lồng chim đã sử dụng, bằng cách thêm các tính năng như đèn trang trí, hệ thống lọc không khí, nhằm cải thiện môi trường sống và tăng cường sự thoải mái cho chim cảnh. Điều này không chỉ tiết kiệm chi phí mà còn tối ưu hóa trải nghiệm cho cả chủ nhân và chim nuôi.",
		Price:             500_000-1_000_000,
		ImagePath:         "/static/img/services_img/phu-kien-2.png",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 3
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Sửa chữa và thay mới đồ chơi",
		Description:       "Dịch vụ Sửa chữa và thay mới đồ chơi cho chim cung cấp giải pháp kỹ thuật chuyên nghiệp để sửa chữa hoặc thay thế các đồ chơi đã hỏng hoặc xuống cấp trong lồng chim. Chuyên gia sẽ đảm bảo rằng đồ chơi mới không chỉ đáp ứng nhu cầu giải trí của chim mà còn đảm bảo an toàn và khích lệ hoạt động tinh thần tích cực của chúng.",
		Price:             299_000-399_000,
		ImagePath:         "/static/img/services_img/phu-kien-3.png",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 4
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Lắp đặt lồng chim tại nhà theo yêu cầu",
		Description:       "Dịch vụ Lắp đặt lồng chim tại nhà theo yêu cầu mang lại sự tiện lợi với chuyên gia đến tận nơi, đo lường và lắp đặt lồng chim theo yêu cầu đặc biệt của khách hàng, tạo ra một môi trường sống hoàn hảo cho chim cảnh.",
		Price:             499_000,
		ImagePath:         "/static/img/services_img/phu-kien-4.jpg",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 5
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Cung cấp đèn nhiệt dành cho chim",
		Description:       "Dịch vụ Cung cấp đèn nhiệt dành cho chim giúp tạo ra môi trường ấm áp và thoải mái cho chim cảnh, đồng thời cung cấp ánh sáng chất lượng để hỗ trợ chu kỳ sinh học và tăng cường sức khỏe của chúng.",
		Price:             300_000,
		ImagePath:         "/static/img/services_img/phu-kien-5.jfif",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 6
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Sách hướng dẫn nuôi chim yến",
		Description:       "Sách hướng dẫn nuôi chim cảnh cung cấp kiến thức chi tiết và hữu ích về cách chăm sóc, nuôi dưỡng, và huấn luyện chim cảnh. Tài liệu này là nguồn thông tin đáng tin cậy, giúp chủ nhân hiểu rõ hơn về nhu cầu cụ thể của từng loại chim, tạo điều kiện sống tốt nhất và cung cấp sự chăm sóc đặc biệt để duy trì sức khỏe và hạnh phúc cho loài chim nuôi.",
		Price:            	149_000,
		ImagePath:         "/static/img/services_img/phu-kien-6.jpg",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)

	// -----------------------------------------------------
	// --------------------Dinh dưỡng và Thức ăn------------
	// -----------------------------------------------------
	// Service 7
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Thức ăn chế biến chất lượng cao",
		Description:       "Thức ăn chế biến chất lượng cao là sự kết hợp độc đáo của các thành phần dinh dưỡng cần thiết, được chế biến cẩn thận để đáp ứng đầy đủ nhu cầu dinh dưỡng của chim cảnh. Sản phẩm này giúp hỗ trợ sức khỏe, tăng cường màu sắc và đảm bảo sự phát triển khỏe mạnh của loài chim.",
		Price:             89_000,
		ImagePath:         "/static/img/services_img/dinh-duong-1.jpg",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 8
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Thức ăn tự nhiên và hữu cơ",
		Description:       "Thức Ăn Tự Nhiên và Hữu Cơ là sự lựa chọn ưu tiên nguồn nguyên liệu tự nhiên và hữu cơ, không chứa hóa chất độc hại. Sản phẩm này cung cấp một chế độ dinh dưỡng an toàn và tự nhiên cho chim cảnh, giúp tạo nên một môi trường sống khỏe mạnh và thiên nhiên trong lồng.",
		Price:             79_000,
		ImagePath:         "/static/img/services_img/dinh-duong-2.png",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 9
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Tư vấn dinh dưỡng cho chim theo yêu cầu",
		Description:       "Dịch vụ Tư vấn dinh dưỡng cho chim mang đến sự chuyên nghiệp và thông tin chính xác về việc lựa chọn thức ăn phù hợp cho loài chim cảnh. Chuyên gia dinh dưỡng sẽ cung cấp hướng dẫn chi tiết, đáp ứng đúng nhu cầu dinh dưỡng của từng loại chim, giúp chủ nhân hiểu rõ và quản lý chế độ ăn sao cho chim có được sức khỏe và sinh sản tốt nhất.",
		Price:             199_000-299_000,
		ImagePath:         "/static/img/services_img/dinh-duong-3.png",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 10
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Thức ăn dành cho chim non",
		Description:       "Thức ăn dành cho chim non là một sự kết hợp cân đối của các thành phần dinh dưỡng quan trọng, thiết kế đặc biệt để đáp ứng nhu cầu phát triển của chim non. Sản phẩm này thường chứa các hạt nhỏ, dễ ăn, và cung cấp năng lượng, vitamin, và khoáng chất cần thiết để hỗ trợ quá trình lớn lên và hình thành sức khỏe toàn diện của chim non.",
		Price:             49_000,
		ImagePath:         "/static/img/services_img/dinh-duong-4.png",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 11
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Hỗ trợ tư vấn dinh dưỡng cho chim cảnh theo mùa",
		Description:       "Dịch vụ Hỗ trợ Tư vấn dinh dưỡng theo mùa cung cấp sự linh hoạt và tư vấn đặc biệt, điều chỉnh chế độ ăn cho chim theo các yếu tố biến đổi theo mùa vụ. Chuyên gia sẽ hướng dẫn chủ nhân về sự thay đổi cần thiết trong dinh dưỡng để đảm bảo rằng chim nhận được các dạng thức ăn phù hợp với nhu cầu thay đổi của chúng qua từng mùa. Điều này giúp duy trì sức khỏe và sự phát triển ổn định của loài chim trong mọi điều kiện thời tiết.",
		Price:             200_000-400_000,
		ImagePath:         "/static/img/services_img/dinh-duong-5.png",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 12
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Hướng dẫn xây dựng khóa biểu dinh dưỡng theo yêu cầu",
		Description:       "Hướng dẫn xây dựng khóa biểu dinh dưỡng theo yêu cầu cung cấp bước hướng dẫn chi tiết để tạo ra một lịch trình ăn cho chim chính xác với nhu cầu dinh dưỡng và lối sống của chúng. Hướng dẫn này sẽ giúp chủ nhân xây dựng một bảng chế độ ăn đa dạng và phù hợp, đồng thời lập kế hoạch điều chỉnh theo yêu cầu cụ thể của từng loại chim để duy trì sức khỏe tốt nhất và giảm nguy cơ các vấn đề dinh dưỡng.",
		Price:             200_000-300_000,
		ImagePath:         "/static/img/services_img/dinh-duong-6.png",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)

	// -----------------------------------------------------------
	// --------------------Y tế và Chăm sóc sức khỏe--------------
	// -----------------------------------------------------------
	// Service 13
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Kiểm tra sức khỏe cho chim",
		Description:       "Dịch vụ Kiểm tra sức khỏe cho chim bao gồm các quy trình chẩn đoán chuyên sâu để đảm bảo sức khỏe toàn diện của chim cảnh. Bác sĩ thú y sẽ thực hiện các kiểm tra về ngoại hình, lối sống, và kiểm tra y tế, đồng thời cung cấp đánh giá về tình trạng sức khỏe và đề xuất các biện pháp điều trị hoặc bảo vệ phù hợp. Kiểm tra sức khỏe định kỳ giúp phát hiện sớm các vấn đề y tế và duy trì sự trạng thái khỏe mạnh cho chim.",
		Price:             700_000,
		ImagePath:         "/static/img/services_img/suc-khoe-1.png",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 14
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Tiêm phòng Vacxin",
		Description:       "Dịch vụ Tiêm phòng Vaccin bao gồm quá trình cung cấp một loạt các vaccin để bảo vệ chim khỏi các bệnh truyền nhiễm. Bác sĩ thú y sẽ tiêm vaccin theo lịch trình phù hợp với loại chim và môi trường sống của chúng. Quá trình này không chỉ giúp ngăn chặn sự lây lan của bệnh mà còn tăng cường hệ thống miễn dịch, bảo vệ chim khỏi các mối đe dọa tiềm ẩn cho sức khỏe của chúng.",
		Price:             1_000_000,
		ImagePath:         "/static/img/services_img/suc-khoe-2.png",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 15
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Khám bệnh & điều trị tại nhà",
		Description:       "Dịch vụ Khám bệnh tại nhà mang đến sự thuận tiện cho chủ nhân và chim cảnh bằng cách đưa bác sĩ thú y đến tận nơi để thực hiện kiểm tra sức khỏe và chẩn đoán tình trạng y tế của chim. Quá trình này giảm stress cho loài chim và tạo điều kiện thuận lợi cho việc chăm sóc và điều trị khi cần thiết.",
		Price:             500_000-1_000_000,
		ImagePath:         "/static/img/services_img/suc-khoe-3.png",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 16
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Thuốc và phụ kiện y tế",
		Description:       "Dịch vụ Thuốc và Phụ kiện Y tế cung cấp thuốc theo bệnh của chim cưng của khách hàng để đảm bảo chăm sóc tối ưu. Đội ngũ bác sĩ thú y chuyên nghiệp của chúng tôi sẽ thực hiện kiểm tra kỹ lưỡng, đưa ra chẩn đoán chính xác và đề xuất phương pháp điều trị hiệu quả nhất. Chúng tôi hiểu rằng mỗi loài chim có nhu cầu y tế riêng biệt, và với sự chuyên nghiệp và tận tâm, chúng tôi cam kết mang lại sức khỏe và hạnh phúc cho chim cưng của bạn.",
		Price:             300_000,
		ImagePath:         "/static/img/services_img/suc-khoe-4.png",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 17
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Dịch vụ điều trị thường niên",
		Description:       "Chương trình Dịch vụ Điều trị Thường niên của chúng tôi là cam kết duy trì sức khỏe tốt nhất cho chim cảnh của bạn. Chúng tôi sẽ thực hiện kiểm tra định kỳ, tiêm phòng, và cung cấp các liệu pháp y tế cần thiết để ngăn chặn các vấn đề sức khỏe potentional. Đội ngũ bác sĩ thú y chuyên nghiệp của chúng tôi sẽ tư vấn và xây dựng kế hoạch điều trị phù hợp với nhu cầu riêng biệt của loài chim cưng, để đảm bảo chúng luôn duy trì một tình trạng sức khỏe và hạnh phúc toàn diện.",
		Price:             1_500_000-2_000_000,
		ImagePath:         "/static/img/services_img/suc-khoe-5.jfif",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 18
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Tư Vấn Sức Khỏe và Dinh Dưỡng",
		Description:       "Chúng tôi mang đến dịch vụ Tư vấn Sức khỏe và Dinh dưỡng để hỗ trợ chủ nhân hiểu rõ và quản lý tốt nhất sức khỏe của chim cảnh. Đội ngũ chuyên gia của chúng tôi sẽ cung cấp thông tin chi tiết về chế độ dinh dưỡng, lên lịch trình ăn, và các khía cạnh quan trọng khác liên quan đến sức khỏe của loài chim. Chúng tôi cam kết hỗ trợ bạn xây dựng một môi trường sống lành mạnh và đảm bảo chim cảnh của bạn luôn tràn đầy sức sống và hạnh phúc.",
		Price:             400_000,
		ImagePath:         "/static/img/services_img/suc-khoe-6.png",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)

	// -----------------------------------------------------------
	// --------------------Grooming-------------------------------
	// -----------------------------------------------------------
	// Service 19
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Dịch vụ tỉa lông",
		Description:       "Chúng tôi cung cấp dịch vụ tỉa lông chuyên nghiệp để đảm bảo bộ lông của chim cảnh luôn trong tình trạng sạch sẽ và đẹp mắt. Đội ngũ chăm sóc của chúng tôi sẽ thực hiện quy trình tỉa lông cẩn thận, tùy chỉnh theo yêu cầu và loại chim, nhằm tạo ra vẻ ngoại hình hoàn hảo và giữ cho lông luôn khỏe mạnh. Dịch vụ tỉa lông của chúng tôi không chỉ mang lại vẻ đẹp bên ngoài mà còn tăng cường sự thoải mái và sức khỏe cho loài chim của bạn.",
		Price:             200_000,
		ImagePath:         "/static/img/services_img/gromming-1.png",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 20
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Dịch vụ tắm gội",
		Description:       "Dịch vụ tắm gội của chúng tôi là trải nghiệm spa tuyệt vời dành cho chim cảnh của bạn. Đội ngũ chăm sóc tận tâm sẽ chăm sóc từng chi tiết trong quá trình tắm, sử dụng các sản phẩm chăm sóc chất lượng cao để làm sạch và nuôi dưỡng lông. Dịch vụ này không chỉ mang lại cho chim cảnh vẻ sạch sẽ và lông mềm mại, mà còn tạo ra một trải nghiệm thư giãn và dễ chịu, giúp tăng cường tinh thần và sức khỏe chung của loài chim cưng.",
		Price:             300_000,
		ImagePath:         "/static/img/services_img/gromming-2.png",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 21
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Dịch vụ Spa cho chim",
		Description:       "Dịch vụ Spa của chúng tôi không chỉ tạo ra một không gian làm đẹp mà còn mang đến trải nghiệm thư giãn và chăm sóc toàn diện. Dịch vụ Spa của chúng tôi bao gồm tắm gọi chuyên nghiệp với các sản phẩm chăm sóc lông hàng đầu, tạo nên một bộ lông mềm mại, sạch sẽ và mùi hương dễ chịu. ",
		Price:             500_000,
		ImagePath:         "/static/img/services_img/gromming-3.jpg",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 22
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Gromming tại nhà",
		Description:       "Dịch vụ Grooming Chim Cảnh Tại Nhà là sự kết hợp hoàn hảo giữa chăm sóc chuyên nghiệp và tiện ích tại ngôi nhà của bạn. Đội ngũ chăm sóc tận tâm của chúng tôi sẽ đến địa điểm của bạn với đầy đủ thiết bị và kỹ năng để biến việc làm đẹp cho chim cảnh thành một trải nghiệm thuận tiện và thoải mái. Từ tắm gội, tỉa lông, đến chăm sóc móng, chúng tôi cam kết mang đến vẻ đẹp tối ưu cho loài chim cưng của bạn, mà không cần chúng phải rời khỏi tổ ấm ưa thích của mình. Hãy để chúng tôi tạo nên một trải nghiệm làm đẹp tốt nhất cho chim cảnh của bạn ngay tại ngôi nhà của mình.",
		Price:             200_000-600_000,
		ImagePath:         "/static/img/services_img/gromming-4.png",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 23
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Cắt & giũa móng cho chim",
		Description:       "Dịch vụ Cắt & Giũa Móng Cho Chim là một dịch vụ chăm sóc và duy trì sức khỏe cho các loại chim cảnh, như các loại chim cảnh như vẹt, chích chòe, hay các loại chim khác trong môi trường nuôi nhốt. Dịch vụ này tập trung vào việc chăm sóc móng của chim, một phần quan trọng để đảm bảo sự thoải mái và an toàn cho chúng.",
		Price:             150_000,
		ImagePath:         "/static/img/services_img/gromming-5.png",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 24
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Tạo ảnh đẹp cho chim cảnh",
		Description:       "Làm thú cưng trở nên lôi cuốn và đẹp hơn bao giờ hết với dịch vụ Tạo Ảnh Đẹp Cho Chim Cảnh. Chúng tôi không chỉ làm đẹp thú cưng của bạn mà còn ghi lại những khoảnh khắc đáng yêu và độc đáo qua ống kính chuyên nghiệp.",
		Price:             49_000-99_000,
		ImagePath:         "/static/img/services_img/gromming-6.png",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)

	// -----------------------------------------------------------
	// --------------------Đào tạo và Huấn luyện------------------
	// -----------------------------------------------------------
	// Service 25
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Huấn luyện theo yêu cầu",
		Description:       "Huấn Luyện Chim Theo Yêu Cầu là dịch vụ linh hoạt và tùy chỉnh, được thiết kế để đáp ứng đầy đủ nhu cầu và mong muốn riêng biệt của chủ nhân và chim cảnh. Chúng tôi cung cấp một quy trình đào tạo cá nhân hóa, bắt đầu từ việc đánh giá tình hình hiện tại và mục tiêu mong muốn. Đội ngũ chuyên gia sẽ xây dựng một chương trình huấn luyện chuyên sâu, tập trung vào các kỹ năng cụ thể hoặc thách thức mà bạn muốn định hình cho chim của mình. Từ việc huấn luyện lệnh cơ bản đến các kỹ thuật biểu diễn nâng cao, chúng tôi cam kết đưa ra giải pháp hiệu quả và đem lại trải nghiệm huấn luyện tích cực cho cả chủ nhân và loài chim cưng.",
		Price:             300_000-500_000,
		ImagePath:         "/static/img/services_img/trainning-1.jpg",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 26
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Huấn luyện kỹ năng cơ bản cho chim",
		Description:       "Dịch vụ Huấn Luyện Cơ Bản Cho Chim là hành trình khám phá những kỹ năng cơ bản giúp tạo ra một mối quan hệ chặt chẽ và giao tiếp hiệu quả giữa bạn và chim cảnh của mình. Chúng tôi tập trung vào việc huấn luyện những kỹ năng như trả lời, nghe lời, và hiểu lệnh cơ bản. Đội ngũ huấn luyện chuyên nghiệp sẽ tận tâm hướng dẫn bạn về cách thiết lập các lệnh cơ bản, tạo ra một môi trường tích cực để chim phát triển. Dịch vụ này không chỉ giúp chim cảnh của bạn trở nên ngoan ngoãn hơn mà còn tăng cường sự kết nối và hiểu biết giữa bạn và người bạn lông xù của mình.",
		Price:             500_000,
		ImagePath:         "/static/img/services_img/trainning-2.png",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 27
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Sách huấn luyện cho chim Nhồng nói",
		Description:       "Sách Huấn Luyện Cho Chim Nhồng Nói là nguồn tài nguyên quý báu giúp chủ nhân hiểu rõ hơn về quá trình huấn luyện và tương tác với loài chim Nhồng Nói đáng yêu của họ. Bằng cách chi tiết và dễ hiểu, sách này cung cấp những phương pháp hiệu quả để phát triển kỹ năng nói và giao tiếp của chim. Từ việc xây dựng từ vựng đến luyện nghe, chủ nhân sẽ được hướng dẫn cách tạo ra môi trường tích cực và khích lệ sự tương tác thông minh với chim Nhồng Nói. Sách còn chia sẻ các kinh nghiệm thực tế và lời khuyên từ những chuyên gia huấn luyện, làm cho quá trình học trở nên thú vị và mang lại kết quả tích cực. Đây là hướng dẫn không thể thiếu cho những ai muốn xây dựng mối quan hệ đặc biệt và đồng điệu với loài chim thông minh này.",
		Price:             100_000,
		ImagePath:         "/static/img/services_img/trainning-3.jpg",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 28
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Sách tự huấn luyện chim",
		Description:       "Là nguồn thông tin toàn diện để chủ nhân tự mình đào tạo và tạo ra một môi trường tích cực cho sự phát triển của chim cảnh. Tận dụng kiến thức chuyên sâu về hành vi và tâm lý của chim, cuốn sách cung cấp bước đi rõ ràng để huấn luyện từ những lệnh cơ bản đến những kỹ năng phức tạp hơn.",
		Price:             100_000,
		ImagePath:         "/static/img/services_img/trainning-4.jpg",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 29
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Tư vấn và dạy huấn luyện chim cảnh",
		Description:       "Dịch vụ Tư vấn và Dạy Huấn Luyện Chim Cảnh cung cấp sự chăm sóc toàn diện cho thú cưng, bao gồm tư vấn chăm sóc cơ bản, đánh giá hành vi, lập kế hoạch huấn luyện, và dạy kỹ năng cơ bản và nâng cao. Chúng tôi hỗ trợ giải quyết vấn đề hành vi và cung cấp tư vấn trực tuyến hoặc tận nơi, nhằm tạo ra một môi trường tích cực và giao tiếp tốt giữa chủ nhân và chim cảnh.",
		Price:             200_000-500_000,
		ImagePath:         "/static/img/services_img/trainning-5.jpg",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 30
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Huấn luyện nâng cao cho chim cảnh",
		Description:       "Dịch vụ Huấn Luyện Nâng Cao Cho Chim Cảnh không chỉ là quá trình đào tạo thông thường mà còn là hành trình chăm sóc toàn diện để phát triển tối đa kỹ năng và mối quan hệ giữa chủ nhân và chim cảnh. Chúng tôi đặt sự chú ý vào việc phát triển những kỹ năng đặc biệt, từ việc học lệnh phức tạp đến việc thực hiện các động tác và nhiệm vụ đặc biệt.",
		Price:             200_000,
		ImagePath:         "/static/img/services_img/trainning-6.jpg",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)

	// ------------------------------------------------
	// --------------------Khác------------------------
	// ------------------------------------------------
	// Service 31
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Tìm chim thất lạc",
		Description:       "Tìm Chim Thất Lạc là dịch vụ đặc biệt thiết kế để hỗ trợ chủ nhân trong việc đối mặt với tình huống khẩn cấp khi loài chim cưng bị mất tích. Đội ngũ chuyên nghiệp của chúng tôi sẽ ngay lập tức hành động khi nhận được thông báo về việc chim cảnh của bạn mất tích. Chúng tôi cam kết đem lại sự an tâm và hỗ trợ tận tình cho chủ nhân trong những tình huống căng thẳng, với hy vọng đưa loài chim yêu quý trở lại với gia đình của mình một cách an toàn và nhanh chóng.",
		Price:             300_000,
		ImagePath:         "/static/img/services_img/other-1.png",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 32
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Chăm sóc thú cưng trong thời gian chủ vắng mặt",
		Description:       "Dịch vụ Chăm sóc thú cưng trong thời gian chủ vắng mặt là sự giải pháp đáng tin cậy để bảo đảm thú cưng của bạn nhận được sự chăm sóc tốt nhất khi bạn không thể ở bên. Chúng tôi đảm bảo thức ăn, vận động, và tình cảm đều được đáp ứng, mang lại sự an tâm cho bạn và sự thoải mái cho thú cưng của bạn trong suốt thời gian bạn vắng nhà.",
		Price:             1_000_000-3_000_000,
		ImagePath:         "/static/img/services_img/other-2.jpg",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 33
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Tìm chủ nuôi mới cho chim cảnh",
		Description:       "Dịch vụ Tìm Chủ Nuôi Mới Cho Chim Cảnh giúp tìm kiếm một tổ ấm mới và yêu thương cho thú cưng của bạn khi bạn không thể tiếp tục chăm sóc. Chúng tôi chuyên nghiệp trong quá trình kết nối chủ nhân hiện tại và những người yêu thú cưng tiềm năng, đảm bảo rằng chim cảnh của bạn sẽ được chăm sóc tốt và tiếp tục nhận được tình yêu và quan tâm.",
		Price:             200_000,
		ImagePath:         "/static/img/services_img/other-3.jpg",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 34
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Dịch vụ Chăm sóc Tận Nơi cho Chim Cảnh",
		Description:       "Dịch vụ Chăm sóc Tận Nơi cho Chim Cảnh là sự lựa chọn hoàn hảo để đảm bảo thú cưng của bạn nhận được sự chăm sóc tốt nhất ngay tại tổ ấm của chúng. Đội ngũ chuyên gia chăm sóc sẽ đến trực tiếp tại nhà bạn, cung cấp dịch vụ chu đáo, bao gồm việc cung cấp thức ăn, vận động, và tận tâm chăm sóc để đảm bảo sự thoải mái và hạnh phúc cho thú cưng trong môi trường quen thuộc của chúng.",
		Price:             200_000,
		ImagePath:         "/static/img/services_img/other-4.jpg",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 35
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Tân trang khuôn viên nhà thành sân chơi cho chim cảnh",
		Description:       "Dịch vụ Biến Đổi Khu Vực Nhà Thành Sân Chơi cho Chim Cảnh mang đến trải nghiệm tận hưởng cho thú cưng và chủ nhân. Chúng tôi chuyển đổi không gian nhà bạn thành một sân chơi đa dạng với đủ hoạt động giải trí và vận động cho chim cảnh của bạn. Từ các khu vực leo lên, đến những nơi ẩn náu và đồ chơi kích thích, mỗi góc của sân chơi đều được thiết kế để kích thích sự tò mò và niềm vui cho thú cưng của bạn. Đồng thời, không gian cũng được tối ưu hóa để tạo điều kiện sống an toàn và thoải mái nhất cho chim cảnh yêu quý của bạn.",
		Price:             200_000,
		ImagePath:         "/static/img/services_img/other-5.jpg",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 36
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Dịch vụ thông Tin Về Sự Kiện Chim Cảnh",
		Description:       "Dịch vụ Thông Tin Về Sự Kiện Chim Cảnh là nguồn thông tin đáng tin cậy cho những người yêu thú cưng và đặc biệt là những người quan tâm đến chim cảnh. Chúng tôi cung cấp thông báo chi tiết và chính xác về hội chợ, triển lãm, và các sự kiện liên quan đến thế giới chim cảnh. Bạn sẽ không bỏ lỡ cơ hội tham gia vào cộng đồng, trao đổi thông tin, và tìm hiểu về những xu hướng mới nhất trong lĩnh vực nuôi chim cảnh. Dịch vụ của chúng tôi giúp kết nối những người yêu thú cưng và tạo ra cơ hội gặp gỡ, học hỏi, và chia sẻ đam mê với cộng đồng đang ngày càng phát triển này.",
		Price:             200_000,
		ImagePath:         "/static/img/services_img/other-6.jpg",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)
	// -----------------------------------------------------
	// -------------End of adding another service-----------
	// -----------------------------------------------------
	// Create admin
	adminPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	err = testStore.CreateAdmin(context.Background(), CreateAdminParams{
		Email:    "admin@gmail.com",
		Password: string(adminPassword),
	})
}
