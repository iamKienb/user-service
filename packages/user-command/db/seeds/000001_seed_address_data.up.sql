INSERT INTO countries (id, name, code) VALUES (1, 'Việt Nam', 'VN');

INSERT INTO cities (id, country_id, name, type) VALUES 
(1, 1, 'Thành phố Hà Nội', 'Thành phố Trung ương'),
(79, 1, 'Thành phố Hồ Chí Minh', 'Thành phố Trung ương');

INSERT INTO districts (id, city_id, name, type) VALUES 
-- Hà Nội
(1, 1, 'Quận Ba Đình', 'Quận'),
(2, 1, 'Quận Hoàn Kiếm', 'Quận'),
(3, 1, 'Quận Tây Hồ', 'Quận'),
(4, 1, 'Quận Cầu Giấy', 'Quận'),
(5, 1, 'Quận Đống Đa', 'Quận'),
(6, 1, 'Quận Hai Bà Trưng', 'Quận'),
(7, 1, 'Quận Hoàng Mai', 'Quận'),
(8, 1, 'Quận Thanh Xuân', 'Quận'),
(9, 1, 'Quận Long Biên', 'Quận'),
-- TP. Hồ Chí Minh
(760, 79, 'Quận 1', 'Quận'),
(761, 79, 'Quận 12', 'Quận'),
(764, 79, 'Quận Gò Vấp', 'Quận'),
(765, 79, 'Quận Bình Thạnh', 'Quận'),
(766, 79, 'Quận Tân Bình', 'Quận'),
(767, 79, 'Quận Tân Phú', 'Quận'),
(768, 79, 'Quận Phú Nhuận', 'Quận'),
(769, 79, 'Thành phố Thủ Đức', 'Thành phố');

INSERT INTO wards (id, district_id, name, type) VALUES 
-- Thuộc Quận Ba Đình (Hà Nội)
(1, 1, 'Phường Phúc Xá', 'Phường'),
(4, 1, 'Phường Trúc Bạch', 'Phường'),
(7, 1, 'Phường Vĩnh Phúc', 'Phường'),
-- Thuộc Quận Cầu Giấy (Hà Nội)
(157, 4, 'Phường Nghĩa Đô', 'Phường'),
(160, 4, 'Phường Nghĩa Tân', 'Phường'),
(163, 4, 'Phường Mai Dịch', 'Phường'),
-- Thuộc Quận 1 (TP. HCM)
(26734, 760, 'Phường Tân Định', 'Phường'),
(26737, 760, 'Phường Đa Kao', 'Phường'),
(26740, 760, 'Phường Bến Nghé', 'Phường'),
-- Thuộc Thành phố Thủ Đức (TP. HCM)
(26743, 769, 'Phường Linh Xuân', 'Phường'),
(26746, 769, 'Phường Linh Trung', 'Phường'),
(26749, 769, 'Phường Tam Phú', 'Phường');
