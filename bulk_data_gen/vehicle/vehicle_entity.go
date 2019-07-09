package vehicle

import (
	. "github.com/caict-benchmark/BDC-TS/bulk_data_gen/common"
	"math/rand"
	"time"
)

var (
	EntityByteString      = []byte("vehicle")       // heap optimization
	EntityTotalByteString = []byte("vehicle-total") // heap optimization
)

var (
	// Field keys for 'vehicle entity' points.
	EntityFieldKeys = [][]byte{
		[]byte("value1"),
		[]byte("value2"),
		[]byte("value3"),
		[]byte("value4"),
		[]byte("value5"),
		[]byte("value6"),
		[]byte("value7"),
		[]byte("value8"),
		[]byte("value9"),
		[]byte("value10"),
		[]byte("value11"),
		[]byte("value12"),
		[]byte("value13"),
		[]byte("value14"),
		[]byte("value15"),
		[]byte("value16"),
		[]byte("value17"),
		[]byte("value18"),
		[]byte("value19"),
		[]byte("value20"),
		[]byte("value21"),
		[]byte("value22"),
		[]byte("value23"),
		[]byte("value24"),
		[]byte("value25"),
		[]byte("value26"),
		[]byte("value27"),
		[]byte("value28"),
		[]byte("value29"),
		[]byte("value30"),
		[]byte("value31"),
		[]byte("value32"),
		[]byte("value33"),
		[]byte("value34"),
		[]byte("value35"),
		[]byte("value36"),
		[]byte("value37"),
		[]byte("value38"),
		[]byte("value39"),
		[]byte("value40"),
		[]byte("value41"),
		[]byte("value42"),
		[]byte("value43"),
		[]byte("value44"),
		[]byte("value45"),
		[]byte("value46"),
		[]byte("value47"),
		[]byte("value48"),
		[]byte("value49"),
		[]byte("value50"),
		[]byte("value51"),
		[]byte("value52"),
		[]byte("value53"),
		[]byte("value54"),
		[]byte("value55"),
		[]byte("value56"),
		[]byte("value57"),
		[]byte("value58"),
		[]byte("value59"),
		[]byte("value60"),
	}
)

type EntityMeasurement struct {
	timestamp time.Time
	//distributions []Distribution
	values []int64
}

func NewEntityMeasurement(start time.Time) *EntityMeasurement {
	//distributions := make([]Distribution, len(EntityFieldKeys))
	//for i := range distributions {
	//	distributions[i] = &ClampedRandomWalkDistribution{
	//		State: rand.Float64() * 100.0,
	//		Min:   0.0,
	//		Max:   100.0,
	//		Step: &NormalDistribution{
	//			Mean:   0.0,
	//			StdDev: 1.0,
	//		},
	//	}
	//}
	values := make([]int64, len(EntityFieldKeys))
	//for i := range values {
	//	index := rand.Intn(100)
	//	values[i] = randomNumbers[i][index]
	//}
	return &EntityMeasurement{
		timestamp: start,
		//distributions: distributions,
		values: values,
	}
}

func (m *EntityMeasurement) Tick(d time.Duration) {
	m.timestamp = m.timestamp.Add(d)
	//for i := range m.distributions {
	//	m.distributions[i].Advance()
	//}
}

func (m *EntityMeasurement) ToPoint(p *Point) bool {
	p.SetMeasurementName(EntityByteString)
	p.SetTimestamp(&m.timestamp)

	//for i := range m.distributions {
	//	p.AppendField(EntityFieldKeys[i], m.distributions[i].Get())
	//}
	for i := range m.values {
		index := rand.Intn(100)
		p.AppendField(EntityFieldKeys[i], randomNumbers[i][index])
	}
	return true
}

var randomNumbers = [][]int64{
	{1, 254, 3, 1, 254, 1, 1, 254, 3, 3, 255, 255, 254, 2, 3, 1, 1, 3, 254, 2, 1, 254, 255, 1, 255, 254, 2, 3, 1, 1, 3, 2, 254, 3, 255, 3, 255, 3, 255, 255, 1, 2, 2, 2, 2, 3, 2, 1, 254, 2, 1, 2, 254, 255, 3, 3, 3, 255, 254, 255, 2, 2, 3, 3, 3, 254, 255, 254, 3, 254, 1, 254, 3, 2, 2, 2, 2, 1, 1, 254, 1, 2, 1, 255, 1, 254, 254, 3, 254, 2, 2, 255, 2, 2, 1, 3, 1, 2, 3, 254},
	{1558, 1068, 205, 449, 692, 30, 686, 208, 280, 463, 2102, 1431, 411, 1628, 523, 305, 1366, 1968, 1198, 1327, 2038, 1545, 775, 2029, 1076, 1690, 1255, 1857, 1312, 359, 524, 594, 1900, 2118, 1156, 2121, 1338, 1186, 497, 633, 1360, 411, 1308, 163, 1780, 1712, 537, 871, 1121, 1332, 1116, 161, 1132, 165, 1672, 1242, 2200, 1221, 2076, 855, 1795, 2091, 1163, 2106, 2057, 352, 1742, 820, 1817, 1029, 1118, 1383, 488, 1324, 1138, 122, 1902, 35, 405, 1933, 961, 964, 968, 1480, 437, 667, 1921, 68, 90, 1198, 1759, 2001, 1502, 2054, 561, 33, 520, 659, 311, 1623},
	{161022, 961521, 699795, 621574, 365161, 281439, 931743, 261748, 755780, 42548, 168473, 439381, 357019, 745135, 42087, 295696, 786214, 441052, 29113, 492886, 686203, 601142, 240898, 816918, 795133, 228396, 499786, 606708, 754287, 819332, 600225, 231346, 531318, 278661, 459486, 145255, 962184, 711024, 298107, 817417, 749023, 145171, 115024, 745097, 264567, 515485, 963446, 400089, 154673, 722082, 829130, 851441, 951428, 530657, 821601, 224197, 815946, 477060, 795095, 194770, 466650, 733669, 871021, 385221, 421804, 566885, 80313, 721184, 835679, 278363, 314021, 213963, 926245, 575265, 502013, 492310, 94406, 634208, 386461, 192029, 733075, 795727, 384261, 207285, 275461, 359983, 88154, 227403, 859672, 175037, 35112, 297428, 782195, 576237, 34654, 886290, 633479, 203318, 204003, 675547},
	{9409, 706, 3174, 9264, 1680, 7184, 7092, 5012, 4110, 2547, 3537, 1331, 7056, 9573, 3850, 8917, 8549, 1699, 3266, 7178, 1097, 4984, 7126, 3964, 8016, 5503, 3224, 9414, 7235, 58, 1733, 6851, 8493, 6690, 31, 9597, 8925, 8875, 1911, 6338, 6864, 6721, 5710, 9705, 701, 9400, 6296, 2850, 5126, 9449, 6456, 9667, 4363, 7296, 7546, 4872, 4423, 7602, 7101, 1472, 3533, 1747, 4629, 6061, 1986, 8093, 114, 2068, 6771, 1199, 5502, 9675, 8199, 3046, 6797, 3233, 2176, 7121, 6803, 4043, 6424, 82, 5041, 45, 2921, 5924, 9030, 1590, 6635, 2620, 3579, 2862, 1369, 7791, 4913, 3705, 8525, 3650, 3151, 1610},
	{14342, 14789, 14753, 13694, 4813, 15372, 19002, 9481, 13507, 3390, 9661, 4013, 4185, 6007, 15949, 9208, 11553, 11388, 1126, 2867, 16, 2034, 4923, 14957, 8456, 18488, 935, 7919, 4991, 13296, 10333, 14633, 10919, 9072, 19814, 3332, 17204, 8410, 15743, 10100, 11764, 12670, 5974, 5200, 5462, 15177, 14703, 16676, 13203, 3601, 11774, 17568, 16997, 4532, 9376, 8990, 1897, 18053, 14605, 15499, 8726, 2925, 15275, 17283, 2372, 403, 6480, 5002, 2760, 6666, 16341, 9026, 13623, 3680, 2255, 7403, 8557, 17305, 10854, 17737, 11705, 19524, 65535, 3154, 11520, 10666, 6244, 9766, 16750, 12454, 3176, 3629, 13804, 339, 10970, 7587, 5109, 2977, 12992, 2164},
	{71, 68, 64, 66, 66, 96, 254, 72, 64, 1, 46, 97, 82, 64, 87, 83, 58, 85, 72, 27, 21, 1, 77, 98, 21, 95, 43, 5, 73, 95, 21, 6, 84, 42, 92, 74, 80, 0, 86, 73, 62, 5, 72, 26, 3, 32, 98, 29, 10, 28, 92, 62, 29, 0, 39, 72, 62, 53, 57, 33, 74, 14, 5, 77, 44, 21, 64, 21, 9, 83, 28, 32, 42, 24, 16, 9, 65, 50, 2, 62, 79, 89, 38, 29, 28, 24, 46, 85, 93, 97, 32, 5, 39, 8, 12, 74, 57, 4, 61, 54},
	{255, 2, 254, 255, 2, 1, 1, 255, 255, 1, 1, 255, 1, 254, 254, 1, 1, 255, 1, 2, 2, 254, 255, 254, 255, 254, 1, 254, 255, 2, 255, 1, 2, 255, 255, 255, 255, 255, 2, 2, 1, 1, 1, 1, 254, 2, 254, 254, 1, 254, 255, 1, 2, 254, 2, 255, 255, 255, 1, 254, 1, 255, 1, 1, 255, 255, 255, 255, 254, 1, 1, 255, 1, 2, 1, 2, 2, 254, 254, 1, 2, 255, 1, 1, 255, 1, 1, 2, 2, 255, 1, 2, 1, 1, 1, 1, 2, 255, 1, 254},
	{35, 63, 40, 19, 28, 28, 54, 49, 44, 6, 33, 14, 23, 55, 13, 25, 40, 47, 21, 38, 58, 59, 57, 50, 15, 15, 4, 36, 6, 2, 62, 31, 46, 50, 56, 8, 40, 14, 40, 32, 26, 5, 55, 55, 30, 17, 7, 32, 38, 27, 14, 14, 23, 46, 47, 21, 44, 44, 23, 0, 27, 37, 12, 20, 23, 4, 12, 32, 26, 61, 1, 47, 29, 12, 54, 55, 60, 47, 44, 58, 13, 4, 63, 47, 34, 49, 48, 12, 62, 49, 50, 56, 63, 16, 0, 27, 55, 51, 46, 33},
	{43343, 51210, 34528, 34536, 19648, 54738, 34056, 17668, 37940, 11390, 39842, 9666, 25225, 5662, 3787, 39998, 15439, 8645, 7109, 33870, 25499, 22253, 25730, 12004, 32820, 35797, 48945, 21679, 55960, 3094, 24328, 49008, 21839, 16778, 40055, 36403, 54503, 48786, 45136, 58117, 39506, 57986, 48772, 35113, 24699, 58056, 30415, 5421, 1493, 11907, 45159, 21674, 45333, 6127, 52958, 51609, 41894, 8140, 42066, 42005, 48404, 16525, 21069, 37514, 34052, 29214, 43165, 33423, 5390, 20065, 32144, 15631, 14518, 13998, 43763, 58147, 51859, 51510, 25463, 47490, 14740, 50164, 49484, 13477, 44641, 59728, 58549, 47820, 23744, 43914, 11020, 42797, 36742, 45930, 36220, 38003, 48295, 32600, 55197, 2979},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{67, 29, 193, 33, 65, 141, 27, 191, 89, 198, 20, 107, 157, 86, 100, 13, 75, 195, 52, 22, 180, 213, 191, 194, 217, 189, 167, 89, 219, 56, 234, 211, 136, 26, 36, 146, 23, 43, 101, 87, 134, 20, 231, 249, 166, 103, 23, 59, 43, 231, 114, 73, 160, 76, 233, 221, 246, 235, 116, 242, 95, 15, 118, 147, 99, 186, 19, 156, 176, 114, 74, 233, 194, 237, 140, 11, 231, 217, 59, 119, 215, 124, 64, 176, 169, 104, 225, 5, 89, 59, 143, 193, 175, 68, 52, 139, 22, 4, 159, 109},
	{19261, 7258, 8514, 44337, 12175, 2630, 3108, 2198, 12041, 59475, 15603, 2257, 45020, 9743, 42789, 34176, 30316, 33686, 12503, 36313, 49362, 64377, 22792, 3985, 36771, 3794, 11281, 40181, 59042, 2, 50498, 28454, 10041, 21596, 56112, 27280, 26935, 39557, 42897, 12139, 62696, 63643, 44143, 34529, 4964, 3435, 63634, 15172, 16027, 26252, 45373, 45005, 59404, 27123, 4323, 61084, 59684, 9748, 62858, 7514, 4923, 14559, 40205, 22005, 31666, 53498, 59349, 46905, 47409, 27204, 41139, 29125, 8499, 51015, 52991, 18388, 16076, 34351, 63824, 28789, 28211, 40800, 59887, 50184, 3957, 9301, 30146, 54861, 52525, 2927, 39927, 8147, 19384, 22413, 5511, 8935, 11297, 11364, 57961, 61705},
	{5, 220, 36, 165, 125, 138, 57, 189, 122, 186, 243, 23, 140, 196, 13, 46, 115, 36, 126, 253, 17, 32, 233, 234, 127, 6, 101, 226, 21, 35, 33, 250, 18, 253, 25, 135, 239, 59, 182, 28, 7, 76, 171, 170, 114, 61, 129, 71, 183, 15, 239, 127, 4, 108, 26, 149, 74, 61, 134, 55, 195, 64, 32, 3, 151, 221, 193, 115, 9, 106, 143, 155, 29, 62, 182, 122, 97, 199, 87, 174, 153, 94, 248, 19, 51, 67, 82, 60, 169, 27, 79, 110, 145, 60, 97, 117, 123, 60, 113, 109},
	{255, 2, 255, 4, 3, 4, 2, 255, 2, 255, 1, 2, 4, 2, 254, 2, 4, 4, 254, 3, 255, 255, 4, 3, 3, 254, 3, 254, 254, 3, 3, 4, 4, 2, 254, 3, 4, 1, 4, 255, 4, 255, 1, 255, 254, 2, 255, 255, 2, 255, 4, 2, 2, 4, 3, 2, 2, 3, 2, 2, 2, 254, 3, 2, 255, 254, 2, 1, 4, 254, 4, 2, 254, 254, 255, 4, 3, 254, 255, 3, 2, 254, 4, 254, 255, 255, 255, 4, 2, 4, 3, 1, 254, 255, 3, 255, 3, 3, 3, 3},
	{16, 165, 45, 241, 190, 220, 42, 69, 44, 235, 3, 213, 37, 187, 101, 218, 246, 83, 192, 112, 158, 34, 212, 50, 106, 180, 109, 207, 120, 149, 160, 189, 37, 180, 168, 184, 130, 127, 169, 77, 87, 94, 44, 233, 255, 214, 54, 128, 248, 18, 115, 210, 247, 232, 124, 78, 210, 155, 255, 233, 46, 43, 186, 48, 15, 219, 236, 182, 157, 22, 100, 200, 120, 200, 91, 41, 119, 247, 107, 153, 49, 55, 189, 203, 136, 22, 48, 230, 192, 172, 152, 68, 126, 105, 66, 116, 117, 249, 30, 220},
	{22406, 11038, 64459, 61919, 52681, 26313, 57992, 43979, 64868, 50833, 51770, 50388, 46916, 7892, 61508, 47839, 53491, 53514, 15044, 7769, 17780, 53067, 42854, 53468, 21998, 43423, 55341, 41682, 40363, 39531, 12345, 9203, 31137, 61737, 18332, 25553, 50000, 41248, 43103, 30102, 20829, 19515, 60848, 60873, 30854, 59173, 4334, 21128, 19546, 64343, 59881, 48270, 2195, 23348, 52929, 12320, 8624, 21720, 31848, 14065, 13825, 5277, 39743, 1867, 55860, 56164, 20691, 50491, 53621, 22262, 36339, 20183, 52521, 15012, 62585, 40166, 18772, 39292, 46491, 21100, 45652, 24463, 16490, 24294, 15531, 24039, 4921, 44893, 52916, 5145, 42630, 39631, 4208, 35868, 23078, 8631, 15462, 53289, 50661, 9062},
	{24293, 10513, 35817, 23015, 5826, 15346, 65151, 54630, 17660, 25550, 19436, 52391, 9308, 12707, 13069, 57368, 25786, 45656, 29226, 50400, 11133, 33580, 36744, 50583, 8767, 45681, 61615, 17676, 11010, 19123, 57709, 890, 33380, 45086, 6472, 13372, 25590, 51646, 9926, 38168, 58874, 58818, 25684, 53133, 13654, 4652, 27174, 20271, 16707, 57876, 3193, 31481, 53421, 60352, 26166, 12560, 54477, 21459, 54441, 58595, 21191, 7824, 28545, 35524, 41383, 6894, 57988, 16140, 31097, 1444, 62239, 41276, 44538, 63121, 21492, 47397, 17991, 19086, 55611, 36603, 31259, 8365, 64118, 30270, 36929, 51847, 25658, 2810, 37344, 20819, 53920, 11307, 46721, 45494, 61799, 17420, 48957, 5832, 47395, 40410},
	{217, 145, 97, 227, 71, 65, 248, 45, 12, 153, 175, 180, 191, 96, 181, 149, 23, 110, 141, 238, 92, 175, 90, 110, 40, 35, 79, 135, 28, 220, 250, 116, 183, 28, 11, 62, 226, 119, 18, 92, 75, 111, 191, 207, 5, 217, 128, 193, 136, 153, 223, 115, 169, 208, 110, 31, 162, 235, 248, 205, 138, 243, 114, 190, 155, 225, 232, 86, 220, 171, 180, 213, 67, 104, 216, 178, 1, 9, 223, 20, 231, 161, 4, 91, 119, 3, 104, 175, 0, 160, 45, 44, 162, 30, 34, 75, 128, 131, 11, 140},
	{53820, 47646, 6683, 21304, 41958, 47272, 58810, 30195, 5437, 9280, 31572, 57108, 57517, 26957, 9193, 9312, 10788, 54780, 24713, 24382, 31285, 3281, 6714, 4888, 30106, 35310, 14558, 26663, 4289, 29820, 16857, 14801, 33913, 15192, 17071, 25054, 6828, 26565, 59084, 19676, 9781, 36524, 15783, 49473, 38190, 19014, 25866, 17525, 24875, 37181, 47638, 25208, 31992, 10771, 18303, 13789, 9999, 17226, 42989, 28077, 41947, 20812, 56396, 58829, 59722, 22458, 9464, 17446, 4459, 27596, 39733, 11638, 44546, 15768, 27144, 49564, 4382, 21093, 55571, 18948, 35216, 38527, 25281, 49648, 26777, 45313, 15345, 52644, 15408, 34835, 28040, 49795, 44381, 2550, 22729, 56884, 10364, 33061, 16932, 24567},
	{3520, 12460, 15616, 14267, 195, 15446, 3559, 8579, 14400, 2410, 4229, 3140, 8174, 15850, 10793, 7475, 4558, 13275, 13769, 6842, 9495, 2012, 307, 5625, 4882, 10604, 8339, 10754, 19510, 8434, 8878, 7104, 8805, 6501, 19215, 3153, 3001, 5215, 14900, 1166, 4743, 5203, 9576, 13865, 14444, 7077, 16230, 19453, 14289, 3177, 17912, 12561, 11414, 16848, 12139, 16168, 5086, 5146, 2152, 15139, 11200, 16011, 15687, 9601, 5626, 4239, 18465, 8513, 1980, 12119, 16246, 17717, 6618, 3948, 3495, 17531, 7216, 5048, 18469, 14548, 5866, 5189, 6828, 41, 10937, 363, 16840, 354, 2350, 15660, 18125, 8670, 7420, 8658, 12961, 4051, 16555, 6832, 6784, 18274},
	{19031, 8719, 9462, 18786, 4925, 19186, 7387, 15273, 14319, 13545, 5213, 2963, 2963, 15872, 10946, 4719, 4299, 11281, 19884, 14604, 11583, 19606, 7217, 17566, 7930, 17330, 19826, 4975, 16271, 4815, 11660, 18767, 3959, 7566, 15902, 10654, 2187, 17782, 12409, 13637, 13081, 14055, 4625, 17441, 10067, 6818, 8972, 11035, 1427, 4249, 10231, 13353, 110, 4819, 106, 14357, 13987, 8072, 19617, 4640, 4036, 16238, 281, 3382, 10722, 4349, 14840, 16447, 16564, 12862, 16364, 15935, 566, 15413, 2795, 1807, 8699, 2069, 7123, 19238, 4642, 17831, 16386, 16225, 3425, 9874, 1021, 12632, 18543, 7265, 4358, 1371, 10389, 724, 2489, 3533, 7538, 6965, 13622, 5495},
	{18417, 17815, 10040, 19252, 14953, 6921, 9878, 4483, 8895, 18070, 2693, 7852, 19322, 13440, 856, 6396, 2169, 8183, 12957, 1867, 4080, 19411, 5414, 13250, 10148, 15346, 8924, 17281, 3423, 6260, 10328, 2751, 12038, 16466, 1117, 8545, 14867, 15345, 660, 7099, 16567, 12536, 15375, 13962, 2151, 3891, 11287, 4514, 13237, 2786, 12827, 1570, 2000, 16208, 14820, 12964, 819, 6796, 13704, 9991, 15832, 18254, 3074, 7367, 2407, 12928, 17722, 9631, 2648, 13789, 14723, 8449, 1037, 16227, 19236, 17448, 18838, 8668, 19826, 8402, 248, 6511, 17876, 3300, 2296, 10105, 1309, 1981, 9582, 6448, 19241, 16933, 15607, 10333, 8635, 2774, 12102, 5346, 7190, 8222},
	{10481, 43060, 21683, 1240, 41960, 32231, 19710, 57002, 49267, 4722, 23000, 14714, 46226, 8126, 499, 37145, 10570, 47263, 11215, 9421, 41055, 50634, 3813, 29326, 54091, 24019, 41742, 32813, 17287, 35537, 53642, 5253, 33440, 27130, 21114, 5665, 36277, 29871, 17540, 25332, 38884, 56420, 22932, 7494, 58920, 25147, 24644, 13879, 29973, 42696, 16297, 55994, 46431, 42290, 36513, 35175, 16227, 17538, 14754, 41927, 11159, 17667, 57617, 51988, 19306, 28762, 55846, 43405, 12161, 32278, 37086, 55509, 6511, 20614, 47175, 56727, 38759, 17438, 30989, 10634, 19776, 17174, 53008, 57712, 3506, 14750, 59736, 55740, 1355, 29726, 16125, 12192, 7797, 49305, 53234, 13995, 24121, 23973, 16373, 31460},
	{17504, 32342, 22494, 44227, 51607, 53208, 2916, 39310, 38972, 8378, 38866, 55005, 64652, 42379, 11880, 53232, 40789, 25270, 12976, 46929, 13654, 49474, 16860, 44120, 63767, 13300, 42810, 31387, 64635, 52750, 21693, 61867, 23195, 1807, 45403, 55604, 48115, 30530, 408, 27559, 32564, 29540, 32695, 18555, 56821, 11631, 61970, 12560, 21110, 5373, 24103, 12961, 59200, 58787, 16084, 21957, 59671, 35628, 2883, 41698, 52320, 44262, 48796, 18185, 52747, 34659, 40794, 35592, 14931, 62808, 44428, 40605, 29018, 47295, 46807, 22901, 57158, 15172, 40043, 6932, 63104, 62197, 33878, 2956, 22933, 43105, 28767, 5578, 33458, 40660, 2341, 30812, 39122, 20318, 1106, 46571, 58717, 16267, 19889, 53969},
	{90, 136, 220, 115, 197, 113, 73, 149, 44, 204, 166, 200, 201, 218, 232, 9, 133, 89, 62, 148, 98, 111, 226, 39, 94, 139, 226, 223, 50, 33, 172, 142, 180, 201, 155, 166, 83, 3, 214, 207, 53, 178, 139, 110, 189, 127, 12, 196, 129, 236, 65, 206, 55, 162, 58, 58, 92, 237, 14, 112, 185, 53, 111, 32, 48, 133, 175, 108, 187, 102, 11, 78, 111, 178, 132, 134, 107, 42, 163, 27, 77, 7, 101, 156, 51, 56, 18, 93, 76, 154, 199, 96, 193, 19, 139, 25, 204, 25, 184, 234},
	{1621, 1846, 764, 153, 245, 2262, 342, 298, 1839, 1930, 556, 2241, 1418, 691, 9, 1793, 1300, 972, 1790, 2078, 80, 163, 308, 782, 1134, 43, 659, 830, 620, 573, 614, 1666, 1824, 2067, 1117, 129, 857, 1053, 1783, 1018, 521, 165, 1725, 231, 676, 2093, 2244, 1472, 828, 979, 2375, 1911, 2091, 1444, 1927, 80, 1211, 1462, 1138, 1475, 2324, 784, 2210, 2364, 52, 447, 1089, 223, 849, 1551, 581, 1277, 707, 1394, 1406, 929, 2289, 1626, 353, 2340, 2327, 1931, 878, 1053, 1164, 1852, 1252, 2301, 1774, 1467, 2133, 698, 1203, 1461, 490, 1874, 25, 2166, 258, 105},
	{140, 251, 85, 244, 165, 146, 120, 7, 138, 181, 236, 61, 75, 177, 4, 70, 223, 94, 222, 149, 198, 119, 223, 39, 147, 1, 227, 197, 148, 193, 235, 237, 132, 52, 155, 255, 73, 124, 158, 104, 124, 85, 89, 37, 48, 43, 163, 9, 26, 133, 180, 145, 192, 158, 111, 171, 252, 74, 13, 203, 235, 192, 245, 254, 161, 30, 217, 211, 5, 54, 36, 224, 173, 119, 1, 9, 7, 53, 171, 132, 155, 77, 127, 243, 93, 248, 105, 160, 29, 120, 53, 45, 207, 130, 12, 84, 103, 59, 245, 99},
	{23124, 57175, 36485, 45150, 6912, 45809, 25239, 8795, 10970, 34685, 33291, 18387, 47189, 7524, 14411, 43674, 27741, 55188, 37736, 37085, 18634, 22223, 11882, 42474, 42826, 50782, 26901, 10654, 30921, 9369, 58360, 17736, 25837, 32218, 13750, 47936, 50654, 48090, 31179, 38299, 20571, 11175, 36336, 51664, 10382, 604, 21049, 13475, 6598, 17280, 1306, 17413, 48916, 42204, 55593, 46275, 13260, 58128, 48745, 32369, 36848, 38336, 21804, 55255, 22597, 33180, 39390, 14961, 39641, 58648, 14547, 33215, 43954, 42526, 42036, 48627, 41012, 53812, 7512, 22526, 44804, 56063, 44590, 24661, 18481, 55913, 57011, 2272, 9052, 59607, 6930, 55441, 47804, 47890, 41567, 18856, 40592, 14033, 28559, 3947},
	{116, 8, 65, 87, 154, 206, 142, 191, 141, 155, 13, 16, 27, 216, 32, 179, 188, 65, 44, 175, 151, 182, 61, 120, 243, 88, 101, 28, 211, 33, 65, 129, 26, 218, 78, 181, 175, 206, 43, 4, 99, 24, 54, 23, 209, 212, 164, 165, 173, 233, 222, 229, 244, 252, 148, 25, 135, 45, 201, 35, 187, 205, 183, 234, 102, 83, 2, 45, 254, 20, 205, 64, 148, 255, 62, 112, 159, 176, 9, 200, 37, 74, 56, 144, 82, 204, 12, 59, 39, 62, 68, 211, 228, 3, 184, 219, 51, 100, 30, 114},
	{48, 453, 660, 603, 798, 336, 869, 836, 457, 566, 609, 123, 53, 164, 292, 602, 686, 435, 558, 489, 757, 213, 497, 235, 273, 478, 710, 155, 696, 815, 846, 191, 985, 573, 850, 104, 390, 248, 918, 314, 572, 840, 768, 811, 705, 400, 241, 909, 305, 346, 811, 305, 975, 849, 676, 340, 520, 943, 990, 498, 435, 491, 813, 396, 757, 105, 696, 724, 368, 397, 689, 454, 486, 3, 975, 992, 577, 91, 748, 38, 937, 146, 9, 994, 475, 383, 737, 558, 130, 877, 656, 712, 495, 921, 613, 686, 975, 195, 348, 165},
	{109, 129, 225, 20, 43, 136, 30, 128, 25, 178, 191, 27, 43, 18, 186, 123, 173, 228, 26, 185, 167, 50, 194, 27, 23, 103, 4, 240, 234, 53, 27, 110, 71, 227, 53, 127, 55, 136, 110, 118, 87, 65, 27, 19, 119, 177, 64, 36, 227, 35, 137, 180, 150, 116, 177, 209, 140, 176, 41, 131, 42, 107, 119, 83, 85, 15, 7, 185, 207, 42, 120, 243, 55, 124, 84, 33, 11, 93, 54, 99, 90, 150, 22, 254, 99, 233, 116, 206, 57, 30, 200, 37, 139, 41, 74, 93, 215, 20, 180, 219},
	{1, 2, 2, 2, 254, 255, 1, 254, 255, 1, 1, 255, 2, 255, 254, 254, 2, 255, 254, 2, 255, 1, 1, 254, 1, 1, 1, 2, 254, 1, 1, 254, 255, 255, 2, 1, 254, 1, 1, 255, 1, 254, 254, 2, 1, 254, 2, 1, 1, 255, 1, 254, 1, 255, 1, 254, 1, 255, 1, 255, 1, 1, 254, 254, 255, 2, 2, 1, 254, 254, 255, 254, 255, 255, 2, 2, 255, 2, 2, 1, 1, 1, 254, 255, 2, 254, 1, 255, 1, 254, 254, 1, 255, 254, 1, 254, 2, 2, 254, 254},
	{255, 255, 254, 254, 255, 254, 1, 254, 254, 254, 2, 1, 254, 254, 2, 2, 255, 2, 254, 255, 254, 2, 1, 254, 1, 254, 254, 1, 255, 2, 2, 254, 1, 2, 255, 2, 254, 255, 255, 255, 255, 1, 2, 254, 1, 1, 254, 254, 254, 1, 254, 254, 255, 255, 1, 254, 255, 254, 255, 1, 255, 2, 254, 254, 254, 2, 254, 1, 2, 1, 255, 1, 255, 2, 2, 254, 1, 255, 254, 1, 255, 254, 254, 1, 255, 1, 255, 2, 254, 255, 2, 1, 254, 2, 1, 1, 2, 254, 255, 2},
	{36502, 55060, 46769, 25654, 58832, 28677, 15664, 18130, 39489, 45326, 23774, 26917, 50131, 29438, 36641, 169, 15342, 41115, 26128, 46420, 31827, 26960, 56210, 14320, 31987, 59364, 31836, 32726, 30907, 46655, 19659, 50900, 30373, 29247, 19130, 59066, 14947, 10587, 43112, 3727, 1988, 50245, 53867, 5159, 9007, 23733, 19216, 31721, 2581, 35863, 20792, 25594, 17816, 44230, 53506, 32125, 3642, 2448, 4304, 17658, 5861, 16868, 3840, 25163, 1566, 17617, 51503, 5296, 11228, 13667, 20415, 52124, 199, 22185, 39864, 28400, 29434, 56826, 29551, 15515, 12417, 9765, 25139, 26463, 44131, 21760, 16490, 1136, 48402, 36056, 33445, 14104, 13913, 31825, 42161, 28680, 54871, 44539, 45271, 25988},
	{12310, 52490, 53337, 24530, 52079, 28377, 421, 12836, 20091, 34203, 56038, 1817, 8238, 35302, 48200, 59765, 28558, 20114, 51855, 44953, 766, 5871, 36501, 32450, 16823, 39076, 33996, 42104, 1038, 19668, 47049, 44177, 47780, 18013, 28427, 32255, 1347, 42103, 30560, 43044, 13983, 33516, 44071, 37140, 40347, 53394, 26346, 22232, 48419, 24306, 46802, 55234, 37806, 5779, 56450, 25867, 16699, 15366, 2465, 51950, 4745, 37817, 7959, 8593, 9121, 11339, 59937, 29015, 34234, 13556, 53193, 39103, 30653, 58937, 1881, 48725, 13887, 26582, 1890, 13168, 21328, 44243, 45918, 3031, 4941, 47536, 13527, 53493, 41681, 4367, 48820, 18998, 20768, 44491, 56487, 20322, 41351, 39008, 44985, 52971},
	{5, 6, 4, 6, 2, 5, 1, 7, 3, 7, 6, 7, 6, 6, 7, 1, 7, 1, 7, 6, 7, 0, 0, 6, 6, 4, 4, 1, 5, 2, 5, 7, 7, 1, 1, 0, 3, 7, 1, 5, 7, 1, 5, 3, 7, 2, 2, 1, 5, 0, 7, 6, 1, 2, 7, 0, 7, 4, 1, 6, 6, 1, 7, 6, 5, 2, 6, 6, 7, 4, 1, 2, 7, 0, 0, 7, 2, 1, 5, 3, 3, 5, 4, 5, 1, 2, 0, 6, 3, 3, 0, 7, 2, 7, 5, 0, 1, 4, 7, 1},
	{41389606, 32499518, 155989136, 151094440, 111209504, 175762285, 170901293, 142315697, 56389062, 162426257, 157852470, 8616493, 61485705, 124646359, 38464455, 59697439, 151891082, 82842260, 58479072, 177253912, 159537591, 124449832, 101174738, 7749849, 126967962, 106998553, 81065542, 54019181, 106633175, 142352123, 170429688, 165283999, 3620265, 74990485, 119180389, 10436500, 28189903, 23874894, 19295553, 166895020, 86436933, 36477740, 78589870, 83075663, 122393195, 78019696, 29165614, 42646603, 164996097, 145274800, 140417845, 152302799, 63594980, 130300361, 90973163, 28628173, 5614790, 67943552, 108108186, 102441233, 13213687, 4455220, 3817471, 36011410, 28104392, 39987468, 80155390, 158181732, 89461480, 167091808, 58293734, 102848042, 151694647, 675619, 61850545, 49256028, 119830543, 1123689, 114577483, 28018634, 62720590, 70570295, 34585697, 73340376, 131241822, 163174239, 91046594, 109952810, 162425760, 158629790, 110580804, 110675626, 131838120, 94746125, 80789204, 74146446, 9582184, 581104, 54847576, 91551452},
	{82024028, 56377525, 81585636, 31630848, 87553736, 66620076, 84355551, 70524855, 6482877, 55469751, 58305583, 79840111, 51852155, 85055437, 49169297, 12375276, 76476546, 16326396, 9042604, 32184093, 63569264, 49729440, 46422447, 19498458, 69849740, 75006634, 48501790, 82801690, 83071921, 87143768, 62463973, 37248087, 20107311, 3340034, 66944681, 56721742, 35251774, 72943587, 80634615, 29036951, 8835256, 71351211, 50214800, 8692660, 16781390, 247167, 23399063, 9191893, 82482549, 40125466, 13048207, 56025028, 50624658, 70413326, 39422326, 4057237, 16097542, 1977663, 7505900, 43654481, 74911263, 1217134, 15988327, 81654949, 55921711, 80214490, 33717230, 41086067, 40750283, 77280238, 57663706, 74138134, 48353461, 21112461, 41538424, 59304556, 48072986, 80656449, 48798056, 67973059, 41228240, 74888948, 32828816, 7222851, 57832796, 37115420, 18462926, 58441750, 8254229, 63301596, 60968492, 39742199, 84775744, 22148976, 71572162, 68598549, 26269143, 87590732, 80382937, 902402},
	{166, 89, 93, 230, 163, 146, 237, 47, 60, 38, 86, 237, 60, 159, 219, 224, 65, 246, 182, 207, 78, 144, 176, 75, 30, 188, 156, 196, 131, 42, 189, 138, 150, 145, 102, 214, 210, 233, 127, 218, 45, 64, 180, 206, 132, 73, 213, 17, 36, 34, 186, 49, 101, 231, 79, 232, 67, 71, 23, 82, 238, 194, 81, 54, 78, 61, 85, 102, 144, 73, 176, 130, 57, 38, 141, 70, 170, 200, 210, 116, 49, 187, 84, 181, 207, 40, 147, 236, 130, 168, 242, 12, 60, 172, 208, 135, 224, 205, 29, 137},
	{92, 202, 246, 10, 82, 220, 197, 213, 146, 72, 136, 178, 88, 208, 65, 248, 225, 156, 121, 81, 234, 65, 85, 82, 180, 20, 131, 224, 9, 240, 25, 201, 66, 159, 186, 74, 137, 101, 207, 100, 155, 163, 128, 164, 152, 149, 184, 160, 209, 142, 142, 113, 135, 172, 121, 193, 206, 22, 211, 7, 195, 177, 155, 41, 71, 165, 22, 23, 27, 73, 216, 198, 58, 57, 10, 102, 133, 97, 118, 41, 136, 154, 97, 199, 20, 120, 22, 177, 204, 233, 164, 96, 194, 211, 202, 217, 145, 76, 119, 78},
	{7558, 8223, 8793, 2468, 3057, 2685, 10083, 869, 5093, 8425, 4511, 14606, 1377, 4701, 9623, 8557, 3943, 2813, 3178, 6813, 3224, 14745, 4856, 3066, 14621, 8905, 8649, 65, 10759, 11298, 11489, 8715, 8557, 4336, 14791, 13503, 7181, 1445, 8221, 14505, 6770, 9260, 7385, 9058, 1011, 4963, 5982, 12573, 9950, 10117, 8723, 5452, 387, 2554, 840, 2885, 7193, 1421, 10709, 1775, 14381, 8154, 9007, 11790, 11314, 3266, 7494, 797, 8131, 8914, 1008, 7230, 8840, 14012, 5408, 8027, 13815, 1097, 14862, 14797, 500, 2748, 1916, 13698, 10140, 4748, 2605, 6962, 1900, 27, 3660, 14860, 9781, 9775, 6267, 11318, 4700, 7061, 3904, 12016},
	{106, 249, 25, 205, 43, 250, 232, 39, 250, 44, 9, 150, 11, 95, 213, 71, 71, 49, 68, 198, 1, 12, 159, 118, 19, 200, 98, 154, 25, 27, 64, 160, 173, 187, 121, 194, 137, 226, 234, 135, 213, 227, 133, 12, 183, 146, 35, 69, 230, 188, 86, 227, 127, 42, 27, 147, 57, 101, 230, 64, 85, 140, 246, 61, 80, 34, 72, 119, 9, 232, 111, 212, 97, 167, 236, 246, 100, 61, 57, 227, 104, 129, 246, 88, 239, 29, 116, 114, 26, 87, 8, 206, 196, 245, 65, 229, 33, 197, 157, 87},
	{122, 167, 56, 91, 36, 124, 235, 217, 176, 238, 75, 49, 155, 145, 56, 49, 237, 16, 81, 225, 230, 44, 63, 135, 147, 157, 102, 148, 180, 96, 224, 119, 1, 143, 160, 161, 7, 104, 160, 147, 159, 58, 116, 52, 32, 7, 128, 165, 135, 192, 81, 229, 92, 195, 89, 41, 145, 240, 250, 31, 100, 219, 23, 38, 60, 151, 181, 186, 96, 211, 200, 14, 71, 3, 214, 33, 254, 249, 16, 117, 237, 236, 135, 78, 194, 198, 147, 106, 106, 132, 87, 118, 54, 128, 26, 142, 18, 184, 193, 194},
	{1391, 4821, 12379, 488, 8952, 5368, 7403, 1791, 11544, 8958, 11335, 13813, 14640, 12881, 7497, 12178, 13392, 13707, 6802, 10129, 2720, 12851, 13906, 4707, 10151, 11591, 7610, 12019, 5110, 12216, 9065, 4026, 754, 12752, 4430, 13509, 8644, 5069, 4429, 14722, 273, 5513, 1837, 1142, 1892, 9985, 10259, 5753, 10121, 2104, 14750, 10091, 4010, 10904, 10683, 14988, 2699, 7866, 13214, 10679, 1338, 8679, 3362, 10698, 13424, 14, 9944, 3770, 1599, 2124, 9304, 8371, 418, 11239, 8344, 10348, 12699, 10520, 13489, 12885, 8564, 13271, 4559, 13180, 9609, 5520, 4182, 11240, 10567, 3320, 8592, 11786, 881, 877, 2891, 5169, 2398, 2473, 9373, 13054},
	{254, 31, 4, 105, 236, 187, 74, 151, 115, 71, 213, 96, 30, 150, 97, 170, 195, 107, 136, 63, 86, 236, 24, 115, 237, 248, 108, 59, 140, 142, 218, 77, 33, 10, 191, 163, 254, 103, 135, 183, 254, 234, 203, 198, 101, 219, 109, 177, 75, 222, 201, 156, 185, 22, 245, 126, 113, 18, 108, 186, 197, 81, 36, 239, 227, 160, 84, 83, 100, 30, 36, 38, 20, 85, 77, 220, 220, 90, 243, 27, 70, 15, 71, 246, 114, 237, 24, 45, 108, 72, 162, 124, 125, 111, 165, 151, 77, 211, 12, 2},
	{205, 165, 4, 191, 227, 200, 178, 248, 188, 73, 124, 34, 18, 188, 17, 1, 213, 22, 184, 5, 198, 142, 171, 223, 83, 233, 223, 63, 112, 212, 4, 107, 162, 240, 222, 255, 178, 59, 142, 89, 100, 45, 147, 130, 215, 56, 92, 108, 227, 103, 190, 213, 1, 63, 193, 60, 180, 44, 78, 238, 16, 94, 189, 207, 174, 231, 20, 242, 151, 13, 228, 89, 131, 136, 127, 9, 17, 158, 219, 106, 225, 207, 69, 112, 206, 188, 190, 68, 202, 46, 236, 132, 8, 220, 153, 134, 52, 203, 42, 13},
	{173, 147, 233, 181, 62, 186, 239, 226, 162, 101, 57, 28, 183, 195, 45, 75, 112, 111, 224, 108, 255, 197, 35, 126, 232, 37, 21, 77, 77, 169, 229, 126, 21, 35, 19, 86, 46, 84, 108, 128, 27, 177, 137, 97, 138, 229, 69, 64, 101, 64, 190, 69, 240, 59, 36, 70, 51, 2, 37, 27, 153, 101, 96, 70, 121, 50, 77, 80, 235, 229, 93, 64, 0, 131, 60, 47, 235, 157, 162, 224, 197, 142, 170, 120, 231, 47, 224, 211, 59, 102, 169, 46, 2, 102, 188, 189, 224, 76, 149, 108},
	{19, 78, 185, 208, 165, 128, 247, 64, 143, 7, 113, 106, 248, 163, 57, 210, 104, 28, 2, 37, 10, 241, 143, 130, 198, 121, 157, 170, 98, 10, 238, 154, 85, 39, 92, 136, 9, 74, 62, 49, 242, 92, 94, 43, 204, 228, 155, 31, 241, 49, 81, 3, 239, 11, 241, 50, 187, 203, 228, 115, 184, 129, 97, 101, 14, 5, 210, 139, 82, 158, 15, 75, 69, 224, 20, 210, 83, 1, 137, 206, 204, 230, 20, 159, 74, 98, 69, 116, 36, 128, 97, 19, 77, 177, 216, 144, 59, 63, 166, 186},
	{238, 17, 28, 11, 57, 150, 241, 163, 62, 255, 103, 229, 33, 115, 58, 134, 186, 77, 131, 62, 137, 132, 143, 115, 246, 73, 198, 81, 80, 217, 72, 57, 4, 25, 141, 192, 34, 131, 131, 188, 166, 52, 55, 58, 40, 176, 57, 204, 71, 21, 161, 197, 140, 18, 137, 217, 14, 98, 245, 173, 134, 170, 236, 249, 89, 168, 124, 197, 191, 64, 47, 19, 21, 175, 22, 86, 240, 110, 66, 219, 176, 184, 197, 48, 171, 49, 112, 140, 63, 216, 119, 108, 55, 60, 73, 32, 3, 154, 53, 214},
	{78, 58, 223, 31, 243, 114, 139, 195, 224, 245, 201, 14, 124, 243, 113, 88, 114, 40, 154, 128, 23, 87, 79, 167, 92, 89, 139, 23, 152, 47, 173, 166, 156, 25, 105, 146, 170, 57, 78, 160, 231, 164, 237, 1, 72, 161, 141, 212, 136, 203, 91, 181, 17, 24, 5, 77, 57, 254, 227, 247, 78, 87, 231, 180, 13, 141, 58, 213, 134, 70, 206, 133, 24, 87, 235, 75, 15, 61, 116, 50, 66, 78, 162, 148, 180, 130, 216, 132, 164, 67, 91, 70, 93, 123, 28, 109, 23, 172, 33, 106},
	{3, 1, 1, 1, 1, 3, 0, 2, 2, 2, 2, 2, 3, 255, 255, 1, 254, 1, 254, 3, 3, 254, 255, 1, 255, 2, 2, 3, 0, 0, 254, 1, 2, 1, 255, 3, 255, 0, 3, 1, 2, 255, 2, 255, 254, 2, 254, 3, 3, 2, 2, 255, 3, 255, 2, 2, 255, 2, 255, 1, 0, 1, 254, 255, 1, 255, 2, 1, 1, 0, 255, 2, 255, 255, 255, 2, 3, 0, 1, 3, 254, 1, 0, 0, 254, 3, 2, 254, 1, 254, 255, 3, 255, 255, 1, 255, 0, 3, 0, 0},
	{43, 212, 9, 197, 239, 157, 237, 124, 11, 4, 136, 19, 56, 159, 135, 110, 186, 138, 244, 63, 124, 129, 14, 199, 220, 38, 245, 15, 218, 58, 198, 180, 157, 174, 5, 92, 233, 219, 72, 152, 147, 138, 250, 101, 75, 210, 105, 185, 139, 12, 244, 251, 131, 177, 249, 156, 68, 55, 104, 195, 115, 68, 48, 138, 180, 184, 191, 107, 212, 164, 71, 104, 82, 119, 63, 255, 167, 210, 195, 180, 33, 227, 37, 19, 120, 95, 34, 144, 111, 248, 231, 13, 91, 197, 49, 34, 86, 224, 1, 52},
	{235, 246, 202, 183, 214, 87, 117, 84, 22, 205, 48, 169, 210, 191, 130, 242, 81, 85, 105, 197, 170, 120, 77, 102, 97, 59, 94, 196, 230, 13, 82, 149, 132, 15, 193, 52, 85, 159, 18, 0, 51, 104, 133, 202, 170, 24, 30, 15, 24, 178, 95, 168, 19, 169, 146, 5, 92, 210, 144, 202, 37, 122, 108, 140, 115, 248, 95, 115, 220, 100, 128, 118, 164, 201, 131, 204, 9, 167, 191, 120, 38, 48, 45, 187, 142, 112, 139, 184, 207, 133, 24, 44, 105, 28, 174, 99, 26, 156, 103, 243},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{117, 6, 227, 81, 75, 105, 77, 225, 208, 50, 184, 43, 1, 29, 192, 108, 197, 187, 162, 18, 77, 111, 102, 40, 69, 184, 201, 27, 8, 136, 195, 200, 52, 86, 30, 241, 76, 39, 4, 35, 151, 26, 76, 240, 72, 62, 210, 61, 132, 37, 92, 54, 143, 130, 251, 146, 21, 216, 10, 11, 222, 183, 34, 204, 46, 150, 29, 228, 234, 82, 11, 103, 73, 154, 182, 187, 165, 87, 24, 211, 51, 122, 31, 93, 115, 232, 188, 134, 125, 166, 255, 221, 208, 194, 138, 118, 169, 227, 45, 244},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{170, 42, 217, 249, 200, 138, 30, 0, 190, 241, 177, 109, 250, 149, 59, 127, 158, 207, 55, 7, 57, 8, 153, 32, 93, 53, 210, 89, 71, 254, 47, 151, 30, 116, 234, 81, 141, 116, 176, 91, 158, 91, 241, 238, 214, 142, 255, 83, 202, 136, 87, 234, 148, 241, 80, 53, 222, 101, 173, 165, 147, 167, 97, 25, 162, 247, 245, 6, 223, 160, 29, 13, 144, 137, 173, 170, 220, 66, 201, 156, 99, 190, 225, 251, 169, 111, 104, 51, 255, 110, 246, 248, 113, 236, 53, 28, 212, 96, 224, 184},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{156, 203, 133, 157, 6, 124, 50, 141, 251, 17, 170, 19, 210, 39, 116, 75, 173, 169, 97, 117, 95, 120, 75, 168, 86, 74, 13, 64, 46, 93, 172, 168, 194, 189, 243, 63, 220, 116, 251, 107, 77, 162, 254, 233, 174, 169, 125, 161, 196, 186, 140, 51, 120, 149, 152, 242, 149, 237, 9, 58, 236, 131, 211, 33, 67, 223, 250, 153, 235, 139, 192, 78, 83, 232, 174, 3, 68, 94, 132, 231, 254, 149, 180, 157, 110, 209, 121, 49, 22, 122, 92, 241, 232, 151, 229, 162, 34, 38, 53, 189},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}
