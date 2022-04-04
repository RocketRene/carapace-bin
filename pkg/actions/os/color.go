package os

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionHexColors completes hex color codes
// source: https://www.ditig.com/256-colors-cheat-sheet
//   #0000ff (Blue1)
//   #00d75f (SpringGreen3)
func ActionHexColors() carapace.Action {
	return carapace.ActionStyledValuesDescribed(
		"#000000", "Black (SYSTEM)", style.XTerm256Color(0),
		"#800000", "Maroon (SYSTEM)", style.XTerm256Color(1),
		"#008000", "Green (SYSTEM)", style.XTerm256Color(2),
		"#808000", "Olive (SYSTEM)", style.XTerm256Color(3),
		"#000080", "Navy (SYSTEM)", style.XTerm256Color(4),
		"#800080", "Purple (SYSTEM)", style.XTerm256Color(5),
		"#008080", "Teal (SYSTEM)", style.XTerm256Color(6),
		"#c0c0c0", "Silver (SYSTEM)", style.XTerm256Color(7),
		"#808080", "Grey (SYSTEM)", style.XTerm256Color(8),
		"#ff0000", "Red (SYSTEM)", style.XTerm256Color(9),
		"#00ff00", "Lime (SYSTEM)", style.XTerm256Color(10),
		"#ffff00", "Yellow (SYSTEM)", style.XTerm256Color(11),
		"#0000ff", "Blue (SYSTEM)", style.XTerm256Color(12),
		"#ff00ff", "Fuchsia (SYSTEM)", style.XTerm256Color(13),
		"#00ffff", "Aqua (SYSTEM)", style.XTerm256Color(14),
		"#ffffff", "White (SYSTEM)", style.XTerm256Color(15),
		//"#000000", "Grey0", style.XTerm256Color(16),
		"#00005f", "NavyBlue", style.XTerm256Color(17),
		"#000087", "DarkBlue", style.XTerm256Color(18),
		"#0000af", "Blue3", style.XTerm256Color(19),
		"#0000d7", "Blue3", style.XTerm256Color(20),
		//"#0000ff", "Blue1", style.XTerm256Color(21),
		"#005f00", "DarkGreen", style.XTerm256Color(22),
		"#005f5f", "DeepSkyBlue4", style.XTerm256Color(23),
		"#005f87", "DeepSkyBlue4", style.XTerm256Color(24),
		"#005faf", "DeepSkyBlue4", style.XTerm256Color(25),
		"#005fd7", "DodgerBlue3", style.XTerm256Color(26),
		"#005fff", "DodgerBlue2", style.XTerm256Color(27),
		"#008700", "Green4", style.XTerm256Color(28),
		"#00875f", "SpringGreen4", style.XTerm256Color(29),
		"#008787", "Turquoise4", style.XTerm256Color(30),
		"#0087af", "DeepSkyBlue3", style.XTerm256Color(31),
		"#0087d7", "DeepSkyBlue3", style.XTerm256Color(32),
		"#0087ff", "DodgerBlue1", style.XTerm256Color(33),
		"#00af00", "Green3", style.XTerm256Color(34),
		"#00af5f", "SpringGreen3", style.XTerm256Color(35),
		"#00af87", "DarkCyan", style.XTerm256Color(36),
		"#00afaf", "LightSeaGreen", style.XTerm256Color(37),
		"#00afd7", "DeepSkyBlue2", style.XTerm256Color(38),
		"#00afff", "DeepSkyBlue1", style.XTerm256Color(39),
		"#00d700", "Green3", style.XTerm256Color(40),
		"#00d75f", "SpringGreen3", style.XTerm256Color(41),
		"#00d787", "SpringGreen2", style.XTerm256Color(42),
		"#00d7af", "Cyan3", style.XTerm256Color(43),
		"#00d7d7", "DarkTurquoise", style.XTerm256Color(44),
		"#00d7ff", "Turquoise2", style.XTerm256Color(45),
		//"#00ff00", "Green1", style.XTerm256Color(46),
		"#00ff5f", "SpringGreen2", style.XTerm256Color(47),
		"#00ff87", "SpringGreen1", style.XTerm256Color(48),
		"#00ffaf", "MediumSpringGreen", style.XTerm256Color(49),
		"#00ffd7", "Cyan2", style.XTerm256Color(50),
		//"#00ffff", "Cyan1", style.XTerm256Color(51),
		"#5f0000", "DarkRed", style.XTerm256Color(52),
		"#5f005f", "DeepPink4", style.XTerm256Color(53),
		"#5f0087", "Purple4", style.XTerm256Color(54),
		"#5f00af", "Purple4", style.XTerm256Color(55),
		"#5f00d7", "Purple3", style.XTerm256Color(56),
		"#5f00ff", "BlueViolet", style.XTerm256Color(57),
		"#5f5f00", "Orange4", style.XTerm256Color(58),
		"#5f5f5f", "Grey37", style.XTerm256Color(59),
		"#5f5f87", "MediumPurple4", style.XTerm256Color(60),
		"#5f5faf", "SlateBlue3", style.XTerm256Color(61),
		"#5f5fd7", "SlateBlue3", style.XTerm256Color(62),
		"#5f5fff", "RoyalBlue1", style.XTerm256Color(63),
		"#5f8700", "Chartreuse4", style.XTerm256Color(64),
		"#5f875f", "DarkSeaGreen4", style.XTerm256Color(65),
		"#5f8787", "PaleTurquoise4", style.XTerm256Color(66),
		"#5f87af", "SteelBlue", style.XTerm256Color(67),
		"#5f87d7", "SteelBlue3", style.XTerm256Color(68),
		"#5f87ff", "CornflowerBlue", style.XTerm256Color(69),
		"#5faf00", "Chartreuse3", style.XTerm256Color(70),
		"#5faf5f", "DarkSeaGreen4", style.XTerm256Color(71),
		"#5faf87", "CadetBlue", style.XTerm256Color(72),
		"#5fafaf", "CadetBlue", style.XTerm256Color(73),
		"#5fafd7", "SkyBlue3", style.XTerm256Color(74),
		"#5fafff", "SteelBlue1", style.XTerm256Color(75),
		"#5fd700", "Chartreuse3", style.XTerm256Color(76),
		"#5fd75f", "PaleGreen3", style.XTerm256Color(77),
		"#5fd787", "SeaGreen3", style.XTerm256Color(78),
		"#5fd7af", "Aquamarine3", style.XTerm256Color(79),
		"#5fd7d7", "MediumTurquoise", style.XTerm256Color(80),
		"#5fd7ff", "SteelBlue1", style.XTerm256Color(81),
		"#5fff00", "Chartreuse2", style.XTerm256Color(82),
		"#5fff5f", "SeaGreen2", style.XTerm256Color(83),
		"#5fff87", "SeaGreen1", style.XTerm256Color(84),
		"#5fffaf", "SeaGreen1", style.XTerm256Color(85),
		"#5fffd7", "Aquamarine1", style.XTerm256Color(86),
		"#5fffff", "DarkSlateGray2", style.XTerm256Color(87),
		"#870000", "DarkRed", style.XTerm256Color(88),
		"#87005f", "DeepPink4", style.XTerm256Color(89),
		"#870087", "DarkMagenta", style.XTerm256Color(90),
		"#8700af", "DarkMagenta", style.XTerm256Color(91),
		"#8700d7", "DarkViolet", style.XTerm256Color(92),
		"#8700ff", "Purple", style.XTerm256Color(93),
		"#875f00", "Orange4", style.XTerm256Color(94),
		"#875f5f", "LightPink4", style.XTerm256Color(95),
		"#875f87", "Plum4", style.XTerm256Color(96),
		"#875faf", "MediumPurple3", style.XTerm256Color(97),
		"#875fd7", "MediumPurple3", style.XTerm256Color(98),
		"#875fff", "SlateBlue1", style.XTerm256Color(99),
		"#878700", "Yellow4", style.XTerm256Color(100),
		"#87875f", "Wheat4", style.XTerm256Color(101),
		"#878787", "Grey53", style.XTerm256Color(102),
		"#8787af", "LightSlateGrey", style.XTerm256Color(103),
		"#8787d7", "MediumPurple", style.XTerm256Color(104),
		"#8787ff", "LightSlateBlue", style.XTerm256Color(105),
		"#87af00", "Yellow4", style.XTerm256Color(106),
		"#87af5f", "DarkOliveGreen3", style.XTerm256Color(107),
		"#87af87", "DarkSeaGreen", style.XTerm256Color(108),
		"#87afaf", "LightSkyBlue3", style.XTerm256Color(109),
		"#87afd7", "LightSkyBlue3", style.XTerm256Color(110),
		"#87afff", "SkyBlue2", style.XTerm256Color(111),
		"#87d700", "Chartreuse2", style.XTerm256Color(112),
		"#87d75f", "DarkOliveGreen3", style.XTerm256Color(113),
		"#87d787", "PaleGreen3", style.XTerm256Color(114),
		"#87d7af", "DarkSeaGreen3", style.XTerm256Color(115),
		"#87d7d7", "DarkSlateGray3", style.XTerm256Color(116),
		"#87d7ff", "SkyBlue1", style.XTerm256Color(117),
		"#87ff00", "Chartreuse1", style.XTerm256Color(118),
		"#87ff5f", "LightGreen", style.XTerm256Color(119),
		"#87ff87", "LightGreen", style.XTerm256Color(120),
		"#87ffaf", "PaleGreen1", style.XTerm256Color(121),
		"#87ffd7", "Aquamarine1", style.XTerm256Color(122),
		"#87ffff", "DarkSlateGray1", style.XTerm256Color(123),
		"#af0000", "Red3", style.XTerm256Color(124),
		"#af005f", "DeepPink4", style.XTerm256Color(125),
		"#af0087", "MediumVioletRed", style.XTerm256Color(126),
		"#af00af", "Magenta3", style.XTerm256Color(127),
		"#af00d7", "DarkViolet", style.XTerm256Color(128),
		"#af00ff", "Purple", style.XTerm256Color(129),
		"#af5f00", "DarkOrange3", style.XTerm256Color(130),
		"#af5f5f", "IndianRed", style.XTerm256Color(131),
		"#af5f87", "HotPink3", style.XTerm256Color(132),
		"#af5faf", "MediumOrchid3", style.XTerm256Color(133),
		"#af5fd7", "MediumOrchid", style.XTerm256Color(134),
		"#af5fff", "MediumPurple2", style.XTerm256Color(135),
		"#af8700", "DarkGoldenrod", style.XTerm256Color(136),
		"#af875f", "LightSalmon3", style.XTerm256Color(137),
		"#af8787", "RosyBrown", style.XTerm256Color(138),
		"#af87af", "Grey63", style.XTerm256Color(139),
		"#af87d7", "MediumPurple2", style.XTerm256Color(140),
		"#af87ff", "MediumPurple1", style.XTerm256Color(141),
		"#afaf00", "Gold3", style.XTerm256Color(142),
		"#afaf5f", "DarkKhaki", style.XTerm256Color(143),
		"#afaf87", "NavajoWhite3", style.XTerm256Color(144),
		"#afafaf", "Grey69", style.XTerm256Color(145),
		"#afafd7", "LightSteelBlue3", style.XTerm256Color(146),
		"#afafff", "LightSteelBlue", style.XTerm256Color(147),
		"#afd700", "Yellow3", style.XTerm256Color(148),
		"#afd75f", "DarkOliveGreen3", style.XTerm256Color(149),
		"#afd787", "DarkSeaGreen3", style.XTerm256Color(150),
		"#afd7af", "DarkSeaGreen2", style.XTerm256Color(151),
		"#afd7d7", "LightCyan3", style.XTerm256Color(152),
		"#afd7ff", "LightSkyBlue1", style.XTerm256Color(153),
		"#afff00", "GreenYellow", style.XTerm256Color(154),
		"#afff5f", "DarkOliveGreen2", style.XTerm256Color(155),
		"#afff87", "PaleGreen1", style.XTerm256Color(156),
		"#afffaf", "DarkSeaGreen2", style.XTerm256Color(157),
		"#afffd7", "DarkSeaGreen1", style.XTerm256Color(158),
		"#afffff", "PaleTurquoise1", style.XTerm256Color(159),
		"#d70000", "Red3", style.XTerm256Color(160),
		"#d7005f", "DeepPink3", style.XTerm256Color(161),
		"#d70087", "DeepPink3", style.XTerm256Color(162),
		"#d700af", "Magenta3", style.XTerm256Color(163),
		"#d700d7", "Magenta3", style.XTerm256Color(164),
		"#d700ff", "Magenta2", style.XTerm256Color(165),
		"#d75f00", "DarkOrange3", style.XTerm256Color(166),
		"#d75f5f", "IndianRed", style.XTerm256Color(167),
		"#d75f87", "HotPink3", style.XTerm256Color(168),
		"#d75faf", "HotPink2", style.XTerm256Color(169),
		"#d75fd7", "Orchid", style.XTerm256Color(170),
		"#d75fff", "MediumOrchid1", style.XTerm256Color(171),
		"#d78700", "Orange3", style.XTerm256Color(172),
		"#d7875f", "LightSalmon3", style.XTerm256Color(173),
		"#d78787", "LightPink3", style.XTerm256Color(174),
		"#d787af", "Pink3", style.XTerm256Color(175),
		"#d787d7", "Plum3", style.XTerm256Color(176),
		"#d787ff", "Violet", style.XTerm256Color(177),
		"#d7af00", "Gold3", style.XTerm256Color(178),
		"#d7af5f", "LightGoldenrod3", style.XTerm256Color(179),
		"#d7af87", "Tan", style.XTerm256Color(180),
		"#d7afaf", "MistyRose3", style.XTerm256Color(181),
		"#d7afd7", "Thistle3", style.XTerm256Color(182),
		"#d7afff", "Plum2", style.XTerm256Color(183),
		"#d7d700", "Yellow3", style.XTerm256Color(184),
		"#d7d75f", "Khaki3", style.XTerm256Color(185),
		"#d7d787", "LightGoldenrod2", style.XTerm256Color(186),
		"#d7d7af", "LightYellow3", style.XTerm256Color(187),
		"#d7d7d7", "Grey84", style.XTerm256Color(188),
		"#d7d7ff", "LightSteelBlue1", style.XTerm256Color(189),
		"#d7ff00", "Yellow2", style.XTerm256Color(190),
		"#d7ff5f", "DarkOliveGreen1", style.XTerm256Color(191),
		"#d7ff87", "DarkOliveGreen1", style.XTerm256Color(192),
		"#d7ffaf", "DarkSeaGreen1", style.XTerm256Color(193),
		"#d7ffd7", "Honeydew2", style.XTerm256Color(194),
		"#d7ffff", "LightCyan1", style.XTerm256Color(195),
		//"#ff0000", "Red1", style.XTerm256Color(196),
		"#ff005f", "DeepPink2", style.XTerm256Color(197),
		"#ff0087", "DeepPink1", style.XTerm256Color(198),
		"#ff00af", "DeepPink1", style.XTerm256Color(199),
		"#ff00d7", "Magenta2", style.XTerm256Color(200),
		//"#ff00ff", "Magenta1", style.XTerm256Color(201),
		"#ff5f00", "OrangeRed1", style.XTerm256Color(202),
		"#ff5f5f", "IndianRed1", style.XTerm256Color(203),
		"#ff5f87", "IndianRed1", style.XTerm256Color(204),
		"#ff5faf", "HotPink", style.XTerm256Color(205),
		"#ff5fd7", "HotPink", style.XTerm256Color(206),
		"#ff5fff", "MediumOrchid1", style.XTerm256Color(207),
		"#ff8700", "DarkOrange", style.XTerm256Color(208),
		"#ff875f", "Salmon1", style.XTerm256Color(209),
		"#ff8787", "LightCoral", style.XTerm256Color(210),
		"#ff87af", "PaleVioletRed1", style.XTerm256Color(211),
		"#ff87d7", "Orchid2", style.XTerm256Color(212),
		"#ff87ff", "Orchid1", style.XTerm256Color(213),
		"#ffaf00", "Orange1", style.XTerm256Color(214),
		"#ffaf5f", "SandyBrown", style.XTerm256Color(215),
		"#ffaf87", "LightSalmon1", style.XTerm256Color(216),
		"#ffafaf", "LightPink1", style.XTerm256Color(217),
		"#ffafd7", "Pink1", style.XTerm256Color(218),
		"#ffafff", "Plum1", style.XTerm256Color(219),
		"#ffd700", "Gold1", style.XTerm256Color(220),
		"#ffd75f", "LightGoldenrod2", style.XTerm256Color(221),
		"#ffd787", "LightGoldenrod2", style.XTerm256Color(222),
		"#ffd7af", "NavajoWhite1", style.XTerm256Color(223),
		"#ffd7d7", "MistyRose1", style.XTerm256Color(224),
		"#ffd7ff", "Thistle1", style.XTerm256Color(225),
		//"#ffff00", "Yellow1", style.XTerm256Color(226),
		"#ffff5f", "LightGoldenrod1", style.XTerm256Color(227),
		"#ffff87", "Khaki1", style.XTerm256Color(228),
		"#ffffaf", "Wheat1", style.XTerm256Color(229),
		"#ffffd7", "Cornsilk1", style.XTerm256Color(230),
		//"#ffffff", "Grey100", style.XTerm256Color(231),
		"#080808", "Grey3", style.XTerm256Color(232),
		"#121212", "Grey7", style.XTerm256Color(233),
		"#1c1c1c", "Grey11", style.XTerm256Color(234),
		"#262626", "Grey15", style.XTerm256Color(235),
		"#303030", "Grey19", style.XTerm256Color(236),
		"#3a3a3a", "Grey23", style.XTerm256Color(237),
		"#444444", "Grey27", style.XTerm256Color(238),
		"#4e4e4e", "Grey30", style.XTerm256Color(239),
		"#585858", "Grey35", style.XTerm256Color(240),
		"#626262", "Grey39", style.XTerm256Color(241),
		"#6c6c6c", "Grey42", style.XTerm256Color(242),
		"#767676", "Grey46", style.XTerm256Color(243),
		//"#808080", "Grey50", style.XTerm256Color(244),
		"#8a8a8a", "Grey54", style.XTerm256Color(245),
		"#949494", "Grey58", style.XTerm256Color(246),
		"#9e9e9e", "Grey62", style.XTerm256Color(247),
		"#a8a8a8", "Grey66", style.XTerm256Color(248),
		"#b2b2b2", "Grey70", style.XTerm256Color(249),
		"#bcbcbc", "Grey74", style.XTerm256Color(250),
		"#c6c6c6", "Grey78", style.XTerm256Color(251),
		"#d0d0d0", "Grey82", style.XTerm256Color(252),
		"#dadada", "Grey85", style.XTerm256Color(253),
		"#e4e4e4", "Grey89", style.XTerm256Color(254),
		"#eeeeee", "Grey93", style.XTerm256Color(255),
	)
}

// ActionXtermColorNames completes xterm color names
//   Green
//   Olive
func ActionXtermColorNames() carapace.Action {
	return carapace.ActionValues(
		"Black",
		"Maroon",
		"Green",
		"Olive",
		"Navy",
		"Purple",
		"Teal",
		"Silver",
		"Grey",
		"Red",
		"Lime",
		"Yellow",
		"Blue",
		"Fuchsia",
		"Aqua",
		"White",
	)
}
