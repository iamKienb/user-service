INSERT INTO countries (id, name, code) VALUES (1, 'Việt Nam', 'VN');

INSERT INTO provinces (id, country_id, name, type) VALUES 
(1, 1, 'Thành phố Hà Nội', 'Thành phố Trung ương'),
(79, 1, 'Thành phố Hồ Chí Minh', 'Thành phố Trung ương');

INSERT INTO wards (id, province_id, name, type) VALUES 
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
