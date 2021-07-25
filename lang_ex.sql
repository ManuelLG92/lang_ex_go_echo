-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jul 26, 2021 at 12:33 AM
-- Server version: 10.4.19-MariaDB
-- PHP Version: 8.0.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `lang_ex`
--

-- --------------------------------------------------------

--
-- Table structure for table `countries`
--

CREATE TABLE `countries` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `iso_code` varchar(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `countries`
--

INSERT INTO `countries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `iso_code`) VALUES
(1, NULL, NULL, NULL, 'Afghanistan', 'AF'),
(2, NULL, NULL, NULL, 'Albania', 'AL'),
(3, NULL, NULL, NULL, 'Algeria', 'DZ'),
(4, NULL, NULL, NULL, 'American Samoa', 'AS'),
(5, NULL, NULL, NULL, 'Andorra', 'AD'),
(6, NULL, NULL, NULL, 'Angola', 'AO'),
(7, NULL, NULL, NULL, 'Anguilla', 'AI'),
(8, NULL, NULL, NULL, 'Antarctica', 'AQ'),
(9, NULL, NULL, NULL, 'Antigua And Barbuda', 'AG'),
(10, NULL, NULL, NULL, 'Argentina', 'AR'),
(11, NULL, NULL, NULL, 'Armenia', 'AM'),
(12, NULL, NULL, NULL, 'Aruba', 'AW'),
(13, NULL, NULL, NULL, 'Australia', 'AU'),
(14, NULL, NULL, NULL, 'Austria', 'AT'),
(15, NULL, NULL, NULL, 'Azerbaijan', 'AZ'),
(16, NULL, NULL, NULL, 'Bahamas The', 'BS'),
(17, NULL, NULL, NULL, 'Bahrain', 'BH'),
(18, NULL, NULL, NULL, 'Bangladesh', 'BD'),
(19, NULL, NULL, NULL, 'Barbados', 'BB'),
(20, NULL, NULL, NULL, 'Belarus', 'BY'),
(21, NULL, NULL, NULL, 'Belgium', 'BE'),
(22, NULL, NULL, NULL, 'Belize', 'BZ'),
(23, NULL, NULL, NULL, 'Benin', 'BJ'),
(24, NULL, NULL, NULL, 'Bermuda', 'BM'),
(25, NULL, NULL, NULL, 'Bhutan', 'BT'),
(26, NULL, NULL, NULL, 'Bolivia', 'BO'),
(27, NULL, NULL, NULL, 'Bosnia and Herzegovina', 'BA'),
(28, NULL, NULL, NULL, 'Botswana', 'BW'),
(29, NULL, NULL, NULL, 'Bouvet Island', 'BV'),
(30, NULL, NULL, NULL, 'Brazil', 'BR'),
(31, NULL, NULL, NULL, 'British Indian Ocean Territory', 'IO'),
(32, NULL, NULL, NULL, 'Brunei', 'BN'),
(33, NULL, NULL, NULL, 'Bulgaria', 'BG'),
(34, NULL, NULL, NULL, 'Burkina Faso', 'BF'),
(35, NULL, NULL, NULL, 'Burundi', 'BI'),
(36, NULL, NULL, NULL, 'Cambodia', 'KH'),
(37, NULL, NULL, NULL, 'Cameroon', 'CM'),
(38, NULL, NULL, NULL, 'Canada', 'CA'),
(39, NULL, NULL, NULL, 'Cape Verde', 'CV'),
(40, NULL, NULL, NULL, 'Cayman Islands', 'KY'),
(41, NULL, NULL, NULL, 'Central African Republic', 'CF'),
(42, NULL, NULL, NULL, 'Chad', 'TD'),
(43, NULL, NULL, NULL, 'Chile', 'CL'),
(44, NULL, NULL, NULL, 'China', 'CN'),
(45, NULL, NULL, NULL, 'Christmas Island', 'CX'),
(46, NULL, NULL, NULL, 'Cocos (Keeling) Islands', 'CC'),
(47, NULL, NULL, NULL, 'Colombia', 'CO'),
(48, NULL, NULL, NULL, 'Comoros', 'KM'),
(49, NULL, NULL, NULL, 'Congo', 'CG'),
(50, NULL, NULL, NULL, 'Congo The Democratic Republic Of The', 'CD'),
(51, NULL, NULL, NULL, 'Cook Islands', 'CK'),
(52, NULL, NULL, NULL, 'Costa Rica', 'CR'),
(53, NULL, NULL, NULL, 'Cote D\'Ivoire (Ivory Coast)', 'CI'),
(54, NULL, NULL, NULL, 'Croatia (Hrvatska)', 'HR'),
(55, NULL, NULL, NULL, 'Cuba', 'CU'),
(56, NULL, NULL, NULL, 'Cyprus', 'CY'),
(57, NULL, NULL, NULL, 'Czech Republic', 'CZ'),
(58, NULL, NULL, NULL, 'Denmark', 'DK'),
(59, NULL, NULL, NULL, 'Djibouti', 'DJ'),
(60, NULL, NULL, NULL, 'Dominica', 'DM'),
(61, NULL, NULL, NULL, 'Dominican Republic', 'DO'),
(62, NULL, NULL, NULL, 'East Timor', 'TP'),
(63, NULL, NULL, NULL, 'Ecuador', 'EC'),
(64, NULL, NULL, NULL, 'Egypt', 'EG'),
(65, NULL, NULL, NULL, 'El Salvador', 'SV'),
(66, NULL, NULL, NULL, 'Equatorial Guinea', 'GQ'),
(67, NULL, NULL, NULL, 'Eritrea', 'ER'),
(68, NULL, NULL, NULL, 'Estonia', 'EE'),
(69, NULL, NULL, NULL, 'Ethiopia', 'ET'),
(70, NULL, NULL, NULL, 'External Territories of Australia', 'XA'),
(71, NULL, NULL, NULL, 'Falkland Islands', 'FK'),
(72, NULL, NULL, NULL, 'Faroe Islands', 'FO'),
(73, NULL, NULL, NULL, 'Fiji Islands', 'FJ'),
(74, NULL, NULL, NULL, 'Finland', 'FI'),
(75, NULL, NULL, NULL, 'France', 'FR'),
(76, NULL, NULL, NULL, 'French Guiana', 'GF'),
(77, NULL, NULL, NULL, 'French Polynesia', 'PF'),
(78, NULL, NULL, NULL, 'French Southern Territories', 'TF'),
(79, NULL, NULL, NULL, 'Gabon', 'GA'),
(80, NULL, NULL, NULL, 'Gambia The', 'GM'),
(81, NULL, NULL, NULL, 'Georgia', 'GE'),
(82, NULL, NULL, NULL, 'Germany', 'DE'),
(83, NULL, NULL, NULL, 'Ghana', 'GH'),
(84, NULL, NULL, NULL, 'Gibraltar', 'GI'),
(85, NULL, NULL, NULL, 'Greece', 'GR'),
(86, NULL, NULL, NULL, 'Greenland', 'GL'),
(87, NULL, NULL, NULL, 'Grenada', 'GD'),
(88, NULL, NULL, NULL, 'Guadeloupe', 'GP'),
(89, NULL, NULL, NULL, 'Guam', 'GU'),
(90, NULL, NULL, NULL, 'Guatemala', 'GT'),
(91, NULL, NULL, NULL, 'Guernsey and Alderney', 'XU'),
(92, NULL, NULL, NULL, 'Guinea', 'GN'),
(93, NULL, NULL, NULL, 'Guinea-Bissau', 'GW'),
(94, NULL, NULL, NULL, 'Guyana', 'GY'),
(95, NULL, NULL, NULL, 'Haiti', 'HT'),
(96, NULL, NULL, NULL, 'Heard and McDonald Islands', 'HM'),
(97, NULL, NULL, NULL, 'Honduras', 'HN'),
(98, NULL, NULL, NULL, 'Hong Kong S.A.R.', 'HK'),
(99, NULL, NULL, NULL, 'Hungary', 'HU'),
(100, NULL, NULL, NULL, 'Iceland', 'IS'),
(101, NULL, NULL, NULL, 'India', 'IN'),
(102, NULL, NULL, NULL, 'Indonesia', 'ID'),
(103, NULL, NULL, NULL, 'Iran', 'IR'),
(104, NULL, NULL, NULL, 'Iraq', 'IQ'),
(105, NULL, NULL, NULL, 'Ireland', 'IE'),
(106, NULL, NULL, NULL, 'Israel', 'IL'),
(107, NULL, NULL, NULL, 'Italy', 'IT'),
(108, NULL, NULL, NULL, 'Jamaica', 'JM'),
(109, NULL, NULL, NULL, 'Japan', 'JP'),
(110, NULL, NULL, NULL, 'Jersey', 'XJ'),
(111, NULL, NULL, NULL, 'Jordan', 'JO'),
(112, NULL, NULL, NULL, 'Kazakhstan', 'KZ'),
(113, NULL, NULL, NULL, 'Kenya', 'KE'),
(114, NULL, NULL, NULL, 'Kiribati', 'KI'),
(115, NULL, NULL, NULL, 'Korea North', 'KP'),
(116, NULL, NULL, NULL, 'Korea South', 'KR'),
(117, NULL, NULL, NULL, 'Kuwait', 'KW'),
(118, NULL, NULL, NULL, 'Kyrgyzstan', 'KG'),
(119, NULL, NULL, NULL, 'Laos', 'LA'),
(120, NULL, NULL, NULL, 'Latvia', 'LV'),
(121, NULL, NULL, NULL, 'Lebanon', 'LB'),
(122, NULL, NULL, NULL, 'Lesotho', 'LS'),
(123, NULL, NULL, NULL, 'Liberia', 'LR'),
(124, NULL, NULL, NULL, 'Libya', 'LY'),
(125, NULL, NULL, NULL, 'Liechtenstein', 'LI'),
(126, NULL, NULL, NULL, 'Lithuania', 'LT'),
(127, NULL, NULL, NULL, 'Luxembourg', 'LU'),
(128, NULL, NULL, NULL, 'Macau S.A.R.', 'MO'),
(129, NULL, NULL, NULL, 'Macedonia', 'MK'),
(130, NULL, NULL, NULL, 'Madagascar', 'MG'),
(131, NULL, NULL, NULL, 'Malawi', 'MW'),
(132, NULL, NULL, NULL, 'Malaysia', 'MY'),
(133, NULL, NULL, NULL, 'Maldives', 'MV'),
(134, NULL, NULL, NULL, 'Mali', 'ML'),
(135, NULL, NULL, NULL, 'Malta', 'MT'),
(136, NULL, NULL, NULL, 'Man (Isle of)', 'XM'),
(137, NULL, NULL, NULL, 'Marshall Islands', 'MH'),
(138, NULL, NULL, NULL, 'Martinique', 'MQ'),
(139, NULL, NULL, NULL, 'Mauritania', 'MR'),
(140, NULL, NULL, NULL, 'Mauritius', 'MU'),
(141, NULL, NULL, NULL, 'Mayotte', 'YT'),
(142, NULL, NULL, NULL, 'Mexico', 'MX'),
(143, NULL, NULL, NULL, 'Micronesia', 'FM'),
(144, NULL, NULL, NULL, 'Moldova', 'MD'),
(145, NULL, NULL, NULL, 'Monaco', 'MC'),
(146, NULL, NULL, NULL, 'Mongolia', 'MN'),
(147, NULL, NULL, NULL, 'Montserrat', 'MS'),
(148, NULL, NULL, NULL, 'Morocco', 'MA'),
(149, NULL, NULL, NULL, 'Mozambique', 'MZ'),
(150, NULL, NULL, NULL, 'Myanmar', 'MM'),
(151, NULL, NULL, NULL, 'Namibia', 'NA'),
(152, NULL, NULL, NULL, 'Nauru', 'NR'),
(153, NULL, NULL, NULL, 'Nepal', 'NP'),
(154, NULL, NULL, NULL, 'Netherlands Antilles', 'AN'),
(155, NULL, NULL, NULL, 'Netherlands The', 'NL'),
(156, NULL, NULL, NULL, 'New Caledonia', 'NC'),
(157, NULL, NULL, NULL, 'New Zealand', 'NZ'),
(158, NULL, NULL, NULL, 'Nicaragua', 'NI'),
(159, NULL, NULL, NULL, 'Niger', 'NE'),
(160, NULL, NULL, NULL, 'Nigeria', 'NG'),
(161, NULL, NULL, NULL, 'Niue', 'NU'),
(162, NULL, NULL, NULL, 'Norfolk Island', 'NF'),
(163, NULL, NULL, NULL, 'Northern Mariana Islands', 'MP'),
(164, NULL, NULL, NULL, 'Norway', 'NO'),
(165, NULL, NULL, NULL, 'Oman', 'OM'),
(166, NULL, NULL, NULL, 'Pakistan', 'PK'),
(167, NULL, NULL, NULL, 'Palau', 'PW'),
(168, NULL, NULL, NULL, 'Palestinian Territory Occupied', 'PS'),
(169, NULL, NULL, NULL, 'Panama', 'PA'),
(170, NULL, NULL, NULL, 'Papua new Guinea', 'PG'),
(171, NULL, NULL, NULL, 'Paraguay', 'PY'),
(172, NULL, NULL, NULL, 'Peru', 'PE'),
(173, NULL, NULL, NULL, 'Philippines', 'PH'),
(174, NULL, NULL, NULL, 'Pitcairn Island', 'PN'),
(175, NULL, NULL, NULL, 'Poland', 'PL'),
(176, NULL, NULL, NULL, 'Portugal', 'PT'),
(177, NULL, NULL, NULL, 'Puerto Rico', 'PR'),
(178, NULL, NULL, NULL, 'Qatar', 'QA'),
(179, NULL, NULL, NULL, 'Reunion', 'RE'),
(180, NULL, NULL, NULL, 'Romania', 'RO'),
(181, NULL, NULL, NULL, 'Russia', 'RU'),
(182, NULL, NULL, NULL, 'Rwanda', 'RW'),
(183, NULL, NULL, NULL, 'Saint Helena', 'SH'),
(184, NULL, NULL, NULL, 'Saint Kitts And Nevis', 'KN'),
(185, NULL, NULL, NULL, 'Saint Lucia', 'LC'),
(186, NULL, NULL, NULL, 'Saint Pierre and Miquelon', 'PM'),
(187, NULL, NULL, NULL, 'Saint Vincent And The Grenadines', 'VC'),
(188, NULL, NULL, NULL, 'Samoa', 'WS'),
(189, NULL, NULL, NULL, 'San Marino', 'SM'),
(190, NULL, NULL, NULL, 'Sao Tome and Principe', 'ST'),
(191, NULL, NULL, NULL, 'Saudi Arabia', 'SA'),
(192, NULL, NULL, NULL, 'Senegal', 'SN'),
(193, NULL, NULL, NULL, 'Serbia', 'RS'),
(194, NULL, NULL, NULL, 'Seychelles', 'SC'),
(195, NULL, NULL, NULL, 'Sierra Leone', 'SL'),
(196, NULL, NULL, NULL, 'Singapore', 'SG'),
(197, NULL, NULL, NULL, 'Slovakia', 'SK'),
(198, NULL, NULL, NULL, 'Slovenia', 'SI'),
(199, NULL, NULL, NULL, 'Smaller Territories of the UK', 'XG'),
(200, NULL, NULL, NULL, 'Solomon Islands', 'SB'),
(201, NULL, NULL, NULL, 'Somalia', 'SO'),
(202, NULL, NULL, NULL, 'South Africa', 'ZA'),
(203, NULL, NULL, NULL, 'South Georgia', 'GS'),
(204, NULL, NULL, NULL, 'South Sudan', 'SS'),
(205, NULL, NULL, NULL, 'Spain', 'ES'),
(206, NULL, NULL, NULL, 'Sri Lanka', 'LK'),
(207, NULL, NULL, NULL, 'Sudan', 'SD'),
(208, NULL, NULL, NULL, 'Suriname', 'SR'),
(209, NULL, NULL, NULL, 'Svalbard And Jan Mayen Islands', 'SJ'),
(210, NULL, NULL, NULL, 'Swaziland', 'SZ'),
(211, NULL, NULL, NULL, 'Sweden', 'SE'),
(212, NULL, NULL, NULL, 'Switzerland', 'CH'),
(213, NULL, NULL, NULL, 'Syria', 'SY'),
(214, NULL, NULL, NULL, 'Taiwan', 'TW'),
(215, NULL, NULL, NULL, 'Tajikistan', 'TJ'),
(216, NULL, NULL, NULL, 'Tanzania', 'TZ'),
(217, NULL, NULL, NULL, 'Thailand', 'TH'),
(218, NULL, NULL, NULL, 'Togo', 'TG'),
(219, NULL, NULL, NULL, 'Tokelau', 'TK'),
(220, NULL, NULL, NULL, 'Tonga', 'TO'),
(221, NULL, NULL, NULL, 'Trinidad And Tobago', 'TT'),
(222, NULL, NULL, NULL, 'Tunisia', 'TN'),
(223, NULL, NULL, NULL, 'Turkey', 'TR'),
(224, NULL, NULL, NULL, 'Turkmenistan', 'TM'),
(225, NULL, NULL, NULL, 'Turks And Caicos Islands', 'TC'),
(226, NULL, NULL, NULL, 'Tuvalu', 'TV'),
(227, NULL, NULL, NULL, 'Uganda', 'UG'),
(228, NULL, NULL, NULL, 'Ukraine', 'UA'),
(229, NULL, NULL, NULL, 'United Arab Emirates', 'AE'),
(230, NULL, NULL, NULL, 'United Kingdom', 'GB'),
(231, NULL, NULL, NULL, 'United States', 'US'),
(232, NULL, NULL, NULL, 'United States Minor Outlying Islands', 'UM'),
(233, NULL, NULL, NULL, 'Uruguay', 'UY'),
(234, NULL, NULL, NULL, 'Uzbekistan', 'UZ'),
(235, NULL, NULL, NULL, 'Vanuatu', 'VU'),
(236, NULL, NULL, NULL, 'Vatican City State (Holy See)', 'VA'),
(237, NULL, NULL, NULL, 'Venezuela', 'VE'),
(238, NULL, NULL, NULL, 'Vietnam', 'VN'),
(239, NULL, NULL, NULL, 'Virgin Islands (British)', 'VG'),
(240, NULL, NULL, NULL, 'Virgin Islands (US)', 'VI'),
(241, NULL, NULL, NULL, 'Wallis And Futuna Islands', 'WF'),
(242, NULL, NULL, NULL, 'Western Sahara', 'EH'),
(243, NULL, NULL, NULL, 'Yemen', 'YE'),
(244, NULL, NULL, NULL, 'Yugoslavia', 'YU'),
(245, NULL, NULL, NULL, 'Zambia', 'ZM'),
(246, NULL, NULL, NULL, 'Zimbabwe', 'ZW');

-- --------------------------------------------------------

--
-- Table structure for table `genders`
--

CREATE TABLE `genders` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `genders`
--

INSERT INTO `genders` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`) VALUES
(1, NULL, NULL, NULL, 'Male'),
(2, NULL, NULL, NULL, 'Female'),
(3, NULL, NULL, NULL, 'NULL');

-- --------------------------------------------------------

--
-- Table structure for table `languages`
--

CREATE TABLE `languages` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `iso_code` varchar(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `languages`
--

INSERT INTO `languages` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `iso_code`) VALUES
(1, NULL, NULL, NULL, 'English', 'en'),
(2, NULL, NULL, NULL, 'Afar', 'aa'),
(3, NULL, NULL, NULL, 'Abkhazian', 'ab'),
(4, NULL, NULL, NULL, 'Afrikaans', 'af'),
(5, NULL, NULL, NULL, 'Amharic', 'am'),
(6, NULL, NULL, NULL, 'Arabic', 'ar'),
(7, NULL, NULL, NULL, 'Assamese', 'as'),
(8, NULL, NULL, NULL, 'Aymara', 'ay'),
(9, NULL, NULL, NULL, 'Azerbaijani', 'az'),
(10, NULL, NULL, NULL, 'Bashkir', 'ba'),
(11, NULL, NULL, NULL, 'Belarusian', 'be'),
(12, NULL, NULL, NULL, 'Bulgarian', 'bg'),
(13, NULL, NULL, NULL, 'Bihari', 'bh'),
(14, NULL, NULL, NULL, 'Bislama', 'bi'),
(15, NULL, NULL, NULL, 'Bengali/Bangla', 'bn'),
(16, NULL, NULL, NULL, 'Tibetan', 'bo'),
(17, NULL, NULL, NULL, 'Breton', 'br'),
(18, NULL, NULL, NULL, 'Catalan', 'ca'),
(19, NULL, NULL, NULL, 'Corsican', 'co'),
(20, NULL, NULL, NULL, 'Czech', 'cs'),
(21, NULL, NULL, NULL, 'Welsh', 'cy'),
(22, NULL, NULL, NULL, 'Danish', 'da'),
(23, NULL, NULL, NULL, 'German', 'de'),
(24, NULL, NULL, NULL, 'Bhutani', 'dz'),
(25, NULL, NULL, NULL, 'Greek', 'el'),
(26, NULL, NULL, NULL, 'Esperanto', 'eo'),
(27, NULL, NULL, NULL, 'Spanish', 'es'),
(28, NULL, NULL, NULL, 'Estonian', 'et'),
(29, NULL, NULL, NULL, 'Basque', 'eu'),
(30, NULL, NULL, NULL, 'Persian', 'fa'),
(31, NULL, NULL, NULL, 'Finnish', 'fi'),
(32, NULL, NULL, NULL, 'Fiji', 'fj'),
(33, NULL, NULL, NULL, 'Faeroese', 'fo'),
(34, NULL, NULL, NULL, 'French', 'fr'),
(35, NULL, NULL, NULL, 'Frisian', 'fy'),
(36, NULL, NULL, NULL, 'Irish', 'ga'),
(37, NULL, NULL, NULL, 'Scots/Gaelic', 'gd'),
(38, NULL, NULL, NULL, 'Galician', 'gl'),
(39, NULL, NULL, NULL, 'Guarani', 'gn'),
(40, NULL, NULL, NULL, 'Gujarati', 'gu'),
(41, NULL, NULL, NULL, 'Hausa', 'ha'),
(42, NULL, NULL, NULL, 'Hindi', 'hi'),
(43, NULL, NULL, NULL, 'Croatian', 'hr'),
(44, NULL, NULL, NULL, 'Hungarian', 'hu'),
(45, NULL, NULL, NULL, 'Armenian', 'hy'),
(46, NULL, NULL, NULL, 'Interlingua', 'ia'),
(47, NULL, NULL, NULL, 'Interlingue', 'ie'),
(48, NULL, NULL, NULL, 'Inupiak', 'ik'),
(49, NULL, NULL, NULL, 'Indonesian', 'in'),
(50, NULL, NULL, NULL, 'Icelandic', 'is'),
(51, NULL, NULL, NULL, 'Italian', 'it'),
(52, NULL, NULL, NULL, 'Hebrew', 'iw'),
(53, NULL, NULL, NULL, 'Japanese', 'ja'),
(54, NULL, NULL, NULL, 'Yiddish', 'ji'),
(55, NULL, NULL, NULL, 'Javanese', 'jw'),
(56, NULL, NULL, NULL, 'Georgian', 'ka'),
(57, NULL, NULL, NULL, 'Kazakh', 'kk'),
(58, NULL, NULL, NULL, 'Greenlandic', 'kl'),
(59, NULL, NULL, NULL, 'Cambodian', 'km'),
(60, NULL, NULL, NULL, 'Kannada', 'kn'),
(61, NULL, NULL, NULL, 'Korean', 'ko'),
(62, NULL, NULL, NULL, 'Kashmiri', 'ks'),
(63, NULL, NULL, NULL, 'Kurdish', 'ku'),
(64, NULL, NULL, NULL, 'Kirghiz', 'ky'),
(65, NULL, NULL, NULL, 'Latin', 'la'),
(66, NULL, NULL, NULL, 'Lingala', 'ln'),
(67, NULL, NULL, NULL, 'Laothian', 'lo'),
(68, NULL, NULL, NULL, 'Lithuanian', 'lt'),
(69, NULL, NULL, NULL, 'Latvian/Lettish', 'lv'),
(70, NULL, NULL, NULL, 'Malagasy', 'mg'),
(71, NULL, NULL, NULL, 'Maori', 'mi'),
(72, NULL, NULL, NULL, 'Macedonian', 'mk'),
(73, NULL, NULL, NULL, 'Malayalam', 'ml'),
(74, NULL, NULL, NULL, 'Mongolian', 'mn'),
(75, NULL, NULL, NULL, 'Moldavian', 'mo'),
(76, NULL, NULL, NULL, 'Marathi', 'mr'),
(77, NULL, NULL, NULL, 'Malay', 'ms'),
(78, NULL, NULL, NULL, 'Maltese', 'mt'),
(79, NULL, NULL, NULL, 'Burmese', 'my'),
(80, NULL, NULL, NULL, 'Nauru', 'na'),
(81, NULL, NULL, NULL, 'Nepali', 'ne'),
(82, NULL, NULL, NULL, 'Dutch', 'nl'),
(83, NULL, NULL, NULL, 'Norwegian', 'no'),
(84, NULL, NULL, NULL, 'Occitan', 'oc'),
(85, NULL, NULL, NULL, '(Afan)/Oromoor/Oriya', 'om'),
(86, NULL, NULL, NULL, 'Punjabi', 'pa'),
(87, NULL, NULL, NULL, 'Polish', 'pl'),
(88, NULL, NULL, NULL, 'Pashto/Pushto', 'ps'),
(89, NULL, NULL, NULL, 'Portuguese', 'pt'),
(90, NULL, NULL, NULL, 'Quechua', 'qu'),
(91, NULL, NULL, NULL, 'Rhaeto-Romance', 'rm'),
(92, NULL, NULL, NULL, 'Kirundi', 'rn'),
(93, NULL, NULL, NULL, 'Romanian', 'ro'),
(94, NULL, NULL, NULL, 'Russian', 'ru'),
(95, NULL, NULL, NULL, 'Kinyarwanda', 'rw'),
(96, NULL, NULL, NULL, 'Sanskrit', 'sa'),
(97, NULL, NULL, NULL, 'Sindhi', 'sd'),
(98, NULL, NULL, NULL, 'Sangro', 'sg'),
(99, NULL, NULL, NULL, 'Serbo-Croatian', 'sh'),
(100, NULL, NULL, NULL, 'Singhalese', 'si'),
(101, NULL, NULL, NULL, 'Slovak', 'sk'),
(102, NULL, NULL, NULL, 'Slovenian', 'sl'),
(103, NULL, NULL, NULL, 'Samoan', 'sm'),
(104, NULL, NULL, NULL, 'Shona', 'sn'),
(105, NULL, NULL, NULL, 'Somali', 'so'),
(106, NULL, NULL, NULL, 'Albanian', 'sq'),
(107, NULL, NULL, NULL, 'Serbian', 'sr'),
(108, NULL, NULL, NULL, 'Siswati', 'ss'),
(109, NULL, NULL, NULL, 'Sesotho', 'st'),
(110, NULL, NULL, NULL, 'Sundanese', 'su'),
(111, NULL, NULL, NULL, 'Swedish', 'sv'),
(112, NULL, NULL, NULL, 'Swahili', 'sw'),
(113, NULL, NULL, NULL, 'Tamil', 'ta'),
(114, NULL, NULL, NULL, 'Telugu', 'te'),
(115, NULL, NULL, NULL, 'Tajik', 'tg'),
(116, NULL, NULL, NULL, 'Thai', 'th'),
(117, NULL, NULL, NULL, 'Tigrinya', 'ti'),
(118, NULL, NULL, NULL, 'Turkmen', 'tk'),
(119, NULL, NULL, NULL, 'Tagalog', 'tl'),
(120, NULL, NULL, NULL, 'Setswana', 'tn'),
(121, NULL, NULL, NULL, 'Tonga', 'to'),
(122, NULL, NULL, NULL, 'Turkish', 'tr'),
(123, NULL, NULL, NULL, 'Tsonga', 'ts'),
(124, NULL, NULL, NULL, 'Tatar', 'tt'),
(125, NULL, NULL, NULL, 'Twi', 'tw'),
(126, NULL, NULL, NULL, 'Ukrainian', 'uk'),
(127, NULL, NULL, NULL, 'Urdu', 'ur'),
(128, NULL, NULL, NULL, 'Uzbek', 'uz'),
(129, NULL, NULL, NULL, 'Vietnamese', 'vi'),
(130, NULL, NULL, NULL, 'Volapuk', 'vo'),
(131, NULL, NULL, NULL, 'Wolof', 'wo'),
(132, NULL, NULL, NULL, 'Xhosa', 'xh'),
(133, NULL, NULL, NULL, 'Yoruba', 'yo'),
(134, NULL, NULL, NULL, 'Chinese', 'zh'),
(135, NULL, NULL, NULL, 'Zulu', 'zu');

-- --------------------------------------------------------

--
-- Table structure for table `learning_languages`
--

CREATE TABLE `learning_languages` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` int(10) UNSIGNED DEFAULT NULL,
  `language_id` int(10) UNSIGNED DEFAULT NULL,
  `level` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `native_languages`
--

CREATE TABLE `native_languages` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` int(10) UNSIGNED DEFAULT NULL,
  `language_id` int(10) UNSIGNED DEFAULT NULL,
  `level` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(150) DEFAULT NULL,
  `surname` varchar(150) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `birthday` varchar(255) DEFAULT NULL,
  `country_id` int(11) DEFAULT NULL,
  `gender_id` int(11) DEFAULT NULL,
  `description` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `surname`, `email`, `password`, `birthday`, `country_id`, `gender_id`, `description`) VALUES
(1, '2021-07-21 08:01:09', '2021-07-21 08:01:09', NULL, 'Pedro', 'Gimenez', 'manuel@manuel.es', '$2a$10$MKuujPGWdVjuBrlsYhobM.4v2ktzuSdQ8.YvHnVxEoYlhk5kmWYr.', '24/04/1992', 2, 1, 'text description'),
(7, '2021-07-21 08:09:28', '2021-07-21 08:09:28', NULL, 'Pedro', 'Gimenez', 'manuel@mhgghanuel.es', '$2a$10$dRjbtvsgRTFT3GK3j1Zz3eTkW811e1gwKrJvFPKX1fEqqDER4vnym', '24/04/1992', 2, 1, 'text description'),
(8, '2021-07-21 08:11:10', '2021-07-21 08:11:10', NULL, 'Pedro', 'Gimenez', 'manuel@mhggggghanuel.es', '$2a$10$efQpx8zkCuu.r7lVFRZ.LuwHtKld4F0dohPZQKLkwF7bg9eAOc7/S', '24/04/1992', 2, 1, 'text description'),
(9, '2021-07-21 08:11:58', '2021-07-21 08:11:58', NULL, 'Pedro', 'Gimenez', 'manuel@mhgggghanuel.es', '$2a$10$mvwvzCBC8IMWGSPKvdMuhOY2p9dtedRq7kLPwbT3IldVZLEU/YPri', '24/04/1992', 2, 1, 'text description'),
(10, '2021-07-21 08:12:42', '2021-07-21 08:12:42', NULL, 'Pedro', 'Gimenez', 'manuel@mgggghanuel.es', '$2a$10$5WweCPU32mg4iGkRn9a7se6L3gYu0fx0iC9eCKaxr70nPV.6D4d3a', '24/04/1992', 2, 1, 'text description'),
(11, '2021-07-21 08:13:45', '2021-07-21 08:13:45', NULL, 'Pedro', 'Gimenez', 'manuel@mghjhgghggghanuel.es', '$2a$10$YjPx5CjHRun/sEranfk5hO3A8b.Wli8se9STllIOcI9fZGQe6h2Ea', '24/04/1992', 2, 1, 'text description'),
(12, '2021-07-21 08:13:50', '2021-07-21 08:13:50', NULL, 'Pedro', 'Gimenez', 'manuel@mghjhgghggghhghhganuel.es', '$2a$10$UsaDHYDCTE52le4taO7b4.dTKIv06FQDwahknsbCpBqWYdMOckjdG', '24/04/1992', 2, 1, 'text description'),
(13, '2021-07-21 08:13:52', '2021-07-21 08:13:52', NULL, 'Pedro', 'Gimenez', 'manuel@mghjhgghggghhghhgaghghhgnuel.es', '$2a$10$pGEBy/Zp2p8YuBvxmiqHe.zd9zHN2SILXk2//j//7CV/7DRaOwd/m', '24/04/1992', 2, 1, 'text description'),
(14, '2021-07-21 08:13:58', '2021-07-21 08:13:58', NULL, 'Pedro', 'Gimenez', 'manhgghghuel@mghjhgghggghhghhgaghghhgnuel.es', '$2a$10$wXbe43GR/5riiELm5KRNRO5WxXINjV.0Glps0Xh6ckFKpIWCDQNMa', '24/04/1992', 2, 1, 'text description'),
(15, '2021-07-21 08:14:06', '2021-07-21 08:14:06', NULL, 'Pedro', 'Gimenez', 'manhgghghuel@mghjhgghggghhghhgaghghhgghghgnuel.es', '$2a$10$HE/9J.wluF0UN7mKxw3FJej48qRLzXgMBkX26Oxuamzj9MW1dah9O', '24/04/1992', 2, 1, 'text description'),
(16, '2021-07-21 14:02:48', '2021-07-21 14:02:48', NULL, 'Pedro', 'Gimenez', 'manhgghghuel@mghjhgghggghhghhgaghghhgghghghhghggnuel.es', '$2a$10$8hB0GTtv30UXRlTYmWUUVO4sLZ01bxZjBjPYwUPzwQkrFx126/pYa', '24/04/1992', 2, 1, 'text description'),
(17, '2021-07-21 14:02:52', '2021-07-21 14:02:52', NULL, 'Pedro', 'Gimenez', 'manhgffgfgghghuel@mghjhgghggghhghhgaghghhgghghghhghggnuel.es', '$2a$10$wdKSaVyW9.ldgjxybJVq4.X8y9t7U7fwFIoi3CK.eGWQ28.nPkvYO', '24/04/1992', 2, 1, 'text description'),
(18, '2021-07-21 14:18:21', '2021-07-21 14:18:21', NULL, 'Pedro', 'Gimenez', 'manhgffgfgghghuel@mghjhgghghggnuel.es', '$2a$10$st9cwUtPiZXrQ60MWbqt7ucYONNRDefKXLsWc9p9FogBCsMSPiBJe', '24/04/1992', 2, 1, 'text description'),
(19, '2021-07-21 14:19:53', '2021-07-21 14:19:53', NULL, 'Pedro', 'Gimenez', 'manhgffgfgghghuel@mghjhgghhjhjjhghggnuel.es', '$2a$10$3Js0NrN3HD9MWFeKrcx3B.Wwpq.WFTfuF9HrNcmQ3JtV8u.AR1n.2', '24/04/1992', 2, 1, 'text description'),
(20, '2021-07-21 14:20:40', '2021-07-21 14:20:40', NULL, 'Pedro', 'Gimenez', 'manhgffgffgfgghghuel@mghjhgghhjhjjhghggnuel.es', '$2a$10$IajE0M6cd27Gm9oH419PiekZ15MFLQQqyg8WCpYrs0PnhhOIVFE22', '24/04/1992', 2, 1, 'text description'),
(21, '2021-07-21 17:41:57', '2021-07-21 17:41:57', NULL, 'Pedro', 'Gimenez', 'nolan@nolan.es', '$2a$10$JxPLeZWgNKml1buIBjiSZOd9NZM2g1FBLOWqxFTiKE38VlRSE/rju', '24/04/1992', 2, 1, 'text description'),
(22, '2021-07-24 22:52:37', '2021-07-24 22:52:37', NULL, 'Pedro', 'Gimenez', 'nosdsdsdlan@nolan.es', '$2a$10$y8AGck7yEHvkwcU3abR3Ouc/QCaog5.udRbd0wFK4rIurRyFiwS/K', '24/04/1992', 2, 1, 'text description'),
(23, '2021-07-24 22:52:42', '2021-07-24 22:52:42', NULL, 'Pedro', 'Gimenez', 'nosdsdsdlan@nosddslan.es', '$2a$10$eaib7oFMk0MQUUJNwtEqi.5F43D6BtTv2HnMSinGTBru.xojb3/GK', '24/04/1992', 2, 1, 'text description'),
(24, '2021-07-24 22:52:54', '2021-07-24 22:52:54', NULL, 'Pedsdsdsdro', 'Gimenez', 'nosdsdsdlan@nossdsdsdddslan.es', '$2a$10$MZVlpsxrDOlWe/o.eWL6sew.cijxY/BZTmiYiQ8HZ5BXaLOr1HsjC', '24/04/1992', 2, 1, 'text descrsdsdsdsdiption'),
(25, '2021-07-24 23:19:59', '2021-07-24 23:19:59', NULL, 'Pedsdsdsdro', 'Gimenez', 'noan@nossdsdsdddslan.es', '$2a$10$X7YeJG/kCu7qxvz./D6Xx.Sz3aEJPuZGNCp.jwDnwQh8G16kO8pb6', '24/04/1992', 2, 1, 'text descrsdsdsdsdiption'),
(26, '2021-07-24 23:32:30', '2021-07-25 00:02:24', NULL, 'again Pedrito edited', 'again Edited Gimenez', 'ito@aghgghahghhginedit.es', 'nolannolan.es', '27/08/2005', 8, 3, 'descp edited'),
(28, '2021-07-24 23:56:30', '2021-07-24 23:56:30', NULL, 'again Pedrito edited', 'again Edited Gimenez', 'Pedrito@againedit.es', 'nolannolan.es', '27/08/2005', 8, 3, 'descp edited'),
(29, '2021-07-24 23:57:01', '2021-07-24 23:57:01', NULL, 'again Pedrito edited', 'again Edited Gimenez', 'Peddffdfrito@againedit.es', 'nolannolan.es', '27/08/2005', 8, 3, 'descp edited'),
(30, '2021-07-24 23:57:22', '2021-07-24 23:57:22', NULL, 'again Pedrito edited', 'again Edited Gimenez', 'Peddffdfrito@aghgghainedit.es', 'nolannolan.es', '27/08/2005', 8, 3, 'descp edited'),
(31, '2021-07-25 00:00:08', '2021-07-25 00:11:20', NULL, 'test', 'test', 'nolan@nolan.com', 'nolannolan.es', '27/08/2005', 8, 3, 'descp edited');

-- --------------------------------------------------------

--
-- Table structure for table `user_languages`
--

CREATE TABLE `user_languages` (
  `user_id` int(10) UNSIGNED NOT NULL,
  `language_id` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `countries`
--
ALTER TABLE `countries`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_countries_deleted_at` (`deleted_at`);

--
-- Indexes for table `genders`
--
ALTER TABLE `genders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_genders_deleted_at` (`deleted_at`);

--
-- Indexes for table `languages`
--
ALTER TABLE `languages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_languages_deleted_at` (`deleted_at`);

--
-- Indexes for table `learning_languages`
--
ALTER TABLE `learning_languages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_learning_languages_deleted_at` (`deleted_at`),
  ADD KEY `fk_language_id_learning_languages` (`language_id`),
  ADD KEY `fk_user_id_learning_languages` (`user_id`);

--
-- Indexes for table `native_languages`
--
ALTER TABLE `native_languages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_native_languages_deleted_at` (`deleted_at`),
  ADD KEY `fk_language_id_native_languages` (`language_id`),
  ADD KEY `fk_user_id_native_languages` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- Indexes for table `user_languages`
--
ALTER TABLE `user_languages`
  ADD PRIMARY KEY (`user_id`,`language_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `countries`
--
ALTER TABLE `countries`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=247;

--
-- AUTO_INCREMENT for table `genders`
--
ALTER TABLE `genders`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `languages`
--
ALTER TABLE `languages`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=136;

--
-- AUTO_INCREMENT for table `learning_languages`
--
ALTER TABLE `learning_languages`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `native_languages`
--
ALTER TABLE `native_languages`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=32;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `learning_languages`
--
ALTER TABLE `learning_languages`
  ADD CONSTRAINT `fk_language_id_learning_languages` FOREIGN KEY (`language_id`) REFERENCES `languages` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_user_id_learning_languages` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `native_languages`
--
ALTER TABLE `native_languages`
  ADD CONSTRAINT `fk_language_id_native_languages` FOREIGN KEY (`language_id`) REFERENCES `languages` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_user_id_native_languages` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
