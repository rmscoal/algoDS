package main

import (
	"competitive/interviews"
	"competitive/leetcode"
	"competitive/vidio"
	"fmt"
)

/*
Struct Function (without pointer) will not change the value.

Slice are a reference to an array. Passing slice to a function
parameter and the functions modifies the slice, hence the original
slice are also modified.

Passing non-primitive data types to function parameter in Golang are
likely to modify the original data.
Non-primitive:
- Slice
- Map
- Channel
*/

func main() {
	R := &leetcode.RomanToInt{}
	fmt.Println("Solver:", R.Solver("XIX"))
	fmt.Println("Fast Solver:", R.FastSolver("XIX"))

	T := &leetcode.TwoSum{}
	fmt.Println("Two sum", T.SlowSolver([]int{3, 2, 4}, 6))
	fmt.Println("Two sum", T.Solver([]int{2, 7, 11, 15}, 9))
	fmt.Println("Two sum", T.Solver([]int{2, 1, -5, 11, 15}, -3))

	P := &leetcode.PalindromeNumber{}
	fmt.Println("PalindromeNumber", P.Solver(123))

	L := &leetcode.LongestCommonPrefix{}
	fmt.Println("Longest Common Prefix", L.Solver([]string{"flower", "flow", "flight"}))
	fmt.Println("Longest Common Prefix", L.Solver([]string{"dog", "racecar", "car"}))
	fmt.Println("Longest Common Prefix", L.Solver([]string{"ab", "a"}))
	fmt.Println("Longest Common Prefix", L.Solver([]string{"aaa", "aa", "aaa"}))
	fmt.Println("Longest Common Prefix", L.Solver([]string{"124", "12", "123"}))

	S := &leetcode.SearchInsertPosition{}
	fmt.Println("Search Insert Position", S.Solver([]int{1, 3, 5, 6}, 7))
	fmt.Println("Search Insert Position", S.Solver([]int{1, 3, 5, 6}, 5))
	fmt.Println("Search Insert Position", S.Solver([]int{1, 3, 5, 6}, 2))
	fmt.Println("Search Insert Position", S.Solver([]int{1}, 2))
	fmt.Println("Search Insert Position", S.Solver([]int{1}, 0))
	fmt.Println("Search Insert Position", S.Solver([]int{1, 3}, 2))

	V := &leetcode.ValidParenthesis{}
	fmt.Println("Valid Parenthesis", V.Solver("()"))
	fmt.Println("Valid Parenthesis", V.Solver(")("))
	fmt.Println("Valid Parenthesis", V.Solver("()[]{}"))
	fmt.Println("Valid Parenthesis", V.Solver("(]"))
	fmt.Println("Valid Parenthesis", V.Solver("{[]}"))

	RE := &leetcode.RemoveElement{}
	fmt.Println("Remove Element ----")
	fmt.Println(RE.Solver([]int{1, 2}, 2))
	fmt.Println(RE.Solver([]int{3, 2, 2, 3}, 2))

	RD := &leetcode.RemoveDuplicate{}
	fmt.Println("Remove Duplicate ----")
	fmt.Println(RD.Solver([]int{1, 1, 2}))
	fmt.Println(RD.Solver([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(RD.FastSolution([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	LL := &leetcode.LengthLastWord{}
	fmt.Println("Length of Last World", LL.Solver("Hello World"))
	fmt.Println("Length of Last World", LL.Solver("   fly me   to   the moon  "))
	fmt.Println("Length of Last World", LL.Solver("moon"))

	PO := &leetcode.PlusOne{}
	fmt.Println("Plus One", PO.FastSolver([]int{1, 2, 9}))

	A := &leetcode.AddBinary{}
	fmt.Println("Plus One", A.Solver("1", "11"))
	fmt.Println("Plus One", A.FastSolver("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))

	MP := &leetcode.MaxProfit{}
	fmt.Println("Max Profit", MP.Solver([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("Max Profit", MP.Solver([]int{7, 6, 4, 3, 1}))
	fmt.Println("Max Profit", MP.Solver([]int{2, 4, 1}))
	fmt.Println("Max Profit", MP.FastSolver([]int{225, 130, 661, 808, 236, 798, 659, 987, 564, 244, 886, 158, 304, 705, 365, 344, 560, 424, 11, 876, 12, 803, 805, 82, 484, 117, 491, 962, 709, 2, 377, 566, 380, 844, 670, 278, 455, 59, 829, 719, 980, 647, 12, 751, 938, 234, 399, 472, 635, 634, 78, 31, 363, 201, 453, 368, 548, 267, 387, 823, 654, 71, 178, 273, 727, 657, 58, 845, 487, 698, 337, 188, 390, 177, 256, 724, 958, 745, 152, 665, 101, 534, 8, 784, 442, 954, 763, 2, 104, 846, 628, 904, 19, 193, 370, 188, 964, 778, 408, 831, 2, 374, 381, 678, 474, 28, 514, 344, 478, 948, 186, 844, 438, 912, 287, 611, 618, 247, 161, 542, 878, 75, 530, 706, 952, 69, 988, 279, 445, 377, 186, 350, 788, 242, 815, 878, 829, 489, 557, 341, 906, 691, 975, 513, 295, 79, 348, 650, 436, 927, 472, 379, 48, 208, 612, 375, 278, 653, 67, 294, 522, 982, 812, 299, 862, 716, 652, 389, 107, 358, 656, 54, 43, 687, 515, 115, 980, 440, 406, 322, 4, 116, 587, 807, 452, 571, 704, 526, 264, 853, 521, 345, 671, 747, 888, 313, 12, 965, 527, 419, 894, 250, 31, 298, 171, 687, 125, 752, 986, 375, 430, 215, 491, 993, 615, 884, 111, 355, 396, 177, 890, 272, 505, 563, 949, 554, 572, 345, 329, 766, 701, 228, 409, 104, 200, 347, 544, 165, 982, 971, 277, 626, 240, 44, 159, 164, 214, 565, 580, 505, 528, 983, 281, 570, 631, 889, 524, 513, 455, 523, 579, 66, 66, 15, 67, 416, 185, 440, 767, 274, 670, 79, 559, 274, 194, 991, 519, 515, 159, 336, 704, 53, 449, 693, 890, 833, 507, 604, 849, 645, 645, 408, 430, 302, 764, 516, 252, 584, 814, 27, 824, 41, 305, 330, 709, 722, 218, 202, 511, 78, 227, 517, 476, 588, 488, 60, 813, 472, 235, 415, 278, 235, 522, 102, 992, 487, 134, 103, 697, 816, 500, 483, 310, 119, 947, 48, 628, 373, 373, 694, 323, 587, 665, 285, 677, 928, 709, 273, 892, 141, 534, 513, 625, 502, 459, 840, 145, 863, 171, 198, 477, 428, 250, 141, 479, 511, 131, 791, 44, 925, 657, 373, 437, 748, 330, 623, 938, 742, 169, 402, 953, 229, 73, 759, 878, 398, 915, 493, 909, 207, 423, 405, 900, 158, 263, 403, 340, 291, 612, 19, 254, 562, 162, 488, 41, 237, 518, 149, 827, 210, 87, 333, 394, 52, 871, 119, 732, 35, 917, 234, 725, 710, 146, 913, 45, 902, 295, 665, 767, 650, 605, 571, 645, 426, 11, 995, 901, 55, 4, 955, 936, 626, 447, 995, 726, 938, 28, 872, 560, 601, 470, 920, 183, 127, 591, 57, 471, 0, 306, 205, 978, 669, 492, 585, 103, 697, 210, 923, 512, 306, 736, 850, 316, 78, 484, 507, 111, 258, 954, 338, 816, 138, 149, 181, 515, 805, 362, 274, 530, 875, 272, 955, 267, 404, 230, 598, 730, 30, 68, 348, 655, 499, 729, 715, 168, 121, 885, 386, 171, 693, 860, 168, 203, 997, 382, 219, 493, 511, 500, 581, 223, 185, 535, 467, 969, 297, 178, 82, 807, 896, 26, 886, 909, 772, 200, 90, 462, 373, 682, 672, 101, 35, 592, 359, 742, 945, 172, 660, 823, 986, 547, 514, 169, 602, 963, 624, 898, 4, 376, 697, 42, 877, 175, 126, 185, 813, 516, 205, 238, 520, 501, 640, 485, 293, 811, 302, 292, 858, 481, 685, 749, 786, 224, 176, 190, 156, 339, 475, 190, 805, 841, 24, 135, 428, 841, 824, 708, 861, 769, 552, 707, 27, 469, 168, 278, 964, 4, 612, 558, 187, 410, 286, 228, 566, 368, 182, 252, 443, 140, 11, 390, 944, 724, 613, 327, 688, 47, 519, 510, 945, 814, 679, 487, 93, 543, 502, 119, 5, 309, 989, 802, 576, 814, 787, 942, 89, 194, 646, 954, 663, 74, 819, 770, 977, 235, 777, 386, 326, 751, 937, 557, 985, 887, 67, 412, 924, 146, 473, 168, 596, 845, 146, 341, 315, 412, 616, 599, 501, 731, 676, 47, 414, 474, 644, 388, 232, 924, 110, 635, 703, 27, 6, 432, 622, 519, 494, 870, 576, 834, 563, 807, 261, 746, 550, 943, 540, 63, 925, 724, 818, 110, 786, 696, 959, 334, 135, 71, 29, 666, 101, 645, 640, 443, 676, 33, 296, 368, 956, 529, 300, 733, 198, 265, 70, 222, 570, 829, 136, 179, 397, 333, 831, 153, 417, 722, 667, 283, 390, 821, 745, 753, 140, 479, 475, 136, 715, 549, 455, 207, 350, 318, 160, 910, 527, 778, 641, 233, 601, 244, 580, 685, 600, 854, 165, 941, 916, 73, 138, 320, 140, 987, 948, 558, 425, 330, 663, 198, 716, 687, 242, 439, 56, 608, 446, 173, 237, 113, 712, 170, 274, 285, 894, 464, 403, 331, 159, 312, 836, 439, 984, 301, 39, 71, 645, 233, 399, 155, 51, 280, 543, 547, 480, 557, 314, 626, 12, 432, 686, 159, 569, 762, 882, 412, 17, 241, 415, 334, 138, 499, 798, 22, 522, 873, 779, 187, 763, 997, 305, 765, 722, 376, 775, 327, 313, 220, 441, 208, 877, 518, 910, 293, 211, 616, 595, 482, 123, 405, 836, 127, 250, 648, 925, 99, 995, 84, 471, 326, 964, 764, 302, 761, 149, 766, 856, 256, 883, 543, 437, 33, 244, 16, 318, 406, 239, 429, 402, 990, 628, 339, 657, 211, 654, 181, 477, 770, 820, 857, 992, 193, 640, 622, 621, 910, 481, 723, 686, 622, 456, 147, 556, 888, 282, 330, 116, 401, 677, 842, 487, 526, 520, 766, 862, 650, 987, 915, 388, 942, 599, 797, 958, 760, 953, 308, 786, 804, 533, 450, 431, 106, 901, 777, 426, 14, 38, 694, 548, 992, 479, 352, 411, 774, 606, 428, 793, 111, 266, 403, 151, 374, 819, 59, 466, 610, 449, 277, 29, 28, 989, 747, 368, 16, 146, 281, 55, 429, 189, 465, 315, 147, 42, 672, 960, 82, 926, 830, 947, 698, 284, 949, 792, 75, 194, 855, 395, 867, 512, 850, 459, 994, 815, 159, 971, 937, 856, 223, 148, 478, 655, 458, 747, 690, 647, 669, 992, 834, 927, 876, 883, 330, 983, 780, 960, 252, 807, 646, 677, 668, 861, 111, 645, 691, 152, 85, 619, 47, 65, 453, 903, 692, 141, 427, 190, 188, 527, 421, 103, 69, 764, 2, 257, 170, 640, 178, 483, 508, 105, 19, 916, 614, 863, 548, 741, 367, 541, 701, 835, 448, 989, 588, 498, 765, 591, 439, 335, 42, 583, 447, 728, 148, 406, 983, 282, 927, 880, 212, 392, 11, 457, 216, 35, 762, 873, 962, 543, 32, 879, 74, 776, 81, 61, 268, 806, 786, 241, 213, 860, 411, 463, 749, 954, 578, 777, 991, 846, 758, 735, 113, 628, 406, 304, 178, 971, 158, 426, 380, 408, 866, 36, 644, 280, 143, 821, 447, 465, 441, 382, 201, 620, 259, 115, 785, 928, 642, 394, 719, 898, 943, 73, 89, 67, 779, 623, 621, 968, 197, 794, 383, 971, 837, 280, 854, 765, 840, 70, 733, 375, 275, 562, 10, 210, 325, 633, 857, 848, 593, 771, 966, 31, 945, 44, 795, 696, 24, 232, 284, 418, 773, 3, 252, 486, 824, 781, 948, 894, 853, 406, 552, 938, 971, 453, 637, 463, 610, 565, 991, 500, 56, 443, 318, 226, 995, 580, 734, 890, 368, 817, 57, 520, 197, 783, 912, 780, 25, 802, 4, 599, 843, 746, 862, 224, 606, 140, 740, 637, 166, 490, 118, 822, 113, 477, 663, 996, 576, 781, 388, 990, 573, 396, 192, 132, 343, 118, 351, 67, 727, 12, 128, 602, 664, 764, 338, 122, 299, 769, 251, 272, 177, 285, 410, 509, 547, 726, 422, 64, 662, 912, 323, 58, 728, 347, 321, 789, 269, 274, 673, 310, 357, 384, 969, 231, 945, 769, 685, 53, 202, 854, 133, 30, 868, 633, 727, 955, 141, 880, 173, 762, 889, 617, 174, 277, 641, 287, 99, 109, 606, 706, 754, 838, 274, 203, 126, 505, 913, 736, 517, 369, 990, 8, 203, 228, 575, 158, 550, 147, 340, 287, 411, 462, 339, 54, 288, 611, 943, 451, 566, 418, 407, 317, 966, 308, 885, 311, 290, 330, 801, 259, 576, 363, 268, 469, 925, 977, 534, 566, 482, 985, 556, 465, 400, 20, 633, 330, 621, 304, 409, 732, 535, 788, 142, 607, 220, 346, 558, 790, 249, 912, 557, 194, 993, 913, 101, 787, 323, 246, 739, 512, 151, 247, 573, 828, 529, 347, 428, 738, 94, 256, 263, 940, 954, 9, 121, 958, 46, 331, 965, 65, 159, 837, 920, 120, 4, 627, 552, 501, 327, 493, 618, 28, 704, 588, 336, 578, 137, 255, 228, 182, 970, 648, 377, 18, 719, 571, 934, 415, 314, 455, 286, 832, 613, 401, 370, 421, 832, 409, 269, 939, 985, 249, 411, 62, 477, 452, 405, 524, 581, 845, 724, 573, 525, 408, 250, 347, 837, 436, 641, 884, 183, 75, 302, 488, 107, 368, 266, 536, 294, 78, 909, 421, 345, 523, 226, 711, 711, 804, 652, 906, 389, 895, 250, 347, 768, 938, 312, 762, 709, 234, 892, 407, 626, 68, 563, 366, 888, 628, 992, 53, 46, 727, 23, 297, 294, 839, 721, 808, 978, 45, 258, 241, 549, 796, 179, 619, 135, 204, 63, 542, 172, 559, 945, 179, 709, 671, 703, 516, 899, 390, 192, 718, 587, 26, 384, 41, 523, 37, 888, 436, 689, 761, 841, 50, 108, 121, 663, 0, 871, 670, 912, 280, 422, 83, 76, 331, 469, 702, 77, 583, 904, 813, 104, 934, 794, 208, 766, 849, 261, 805, 149, 793, 647, 387, 73, 428, 525, 957, 744, 201, 116, 371, 485, 939, 131, 826, 839, 571, 305, 542, 277, 695, 320, 257, 990, 170, 602, 554, 259, 773, 727, 567, 243, 535, 97, 84, 463, 954, 585, 81, 948, 32, 810, 289, 870, 124, 603, 480, 765, 483, 986, 280, 456, 939, 888, 826, 809, 748, 304, 29, 426, 533, 502, 482, 374, 325, 432, 624, 231, 461, 39, 283, 274, 334, 181, 999, 930, 613, 471, 431, 577, 777, 213, 496, 449, 254, 658, 585, 267, 658, 217, 763, 186, 54, 198, 275, 228, 958, 903, 48, 323, 115, 361, 198, 778, 105, 113, 948, 301, 234, 407, 75, 108, 799, 360, 116, 315, 634, 432, 857, 815, 458, 201, 10, 28, 724, 343, 877, 907, 128, 94, 570, 102, 549, 680, 437, 364, 111, 653, 326, 674, 132, 486, 471, 788, 161, 207, 773, 415, 531, 948, 102, 709, 165, 486, 433, 320, 784, 178, 625, 553, 360, 87, 209, 43, 39, 795, 840, 493, 803, 513, 316, 582, 94, 547, 702, 709, 207, 28, 494, 266, 857, 669, 719, 294, 457, 240, 934, 648, 809, 741, 78, 983, 288, 631, 438, 518, 253, 575, 8, 335, 198, 858, 536, 781, 401, 267, 236, 2, 16, 90, 790, 120, 499, 853, 498, 715, 927, 383, 597, 425, 923, 65, 463, 250, 948, 186, 642, 32, 265, 374, 336, 849, 314, 104, 317, 97, 741, 74, 589, 317, 668, 607, 721, 431, 690, 185, 953, 894, 549, 119, 410, 54, 383, 72, 148, 154, 907, 172, 294, 620, 589, 640, 941, 257, 98, 688, 852, 388, 489, 795, 348, 135, 719, 984, 831, 123, 537, 240, 503, 185, 632, 824, 880, 508, 932, 714, 340, 362, 484, 447, 108, 32, 960, 576, 472, 788, 695, 121, 191, 457, 671, 13, 294, 442, 133, 592, 847, 494, 997, 315, 398, 206, 197, 250, 35, 992, 422, 463, 655, 826, 143, 807, 866, 868, 697, 321, 758, 845, 560, 974, 462, 109, 398, 93, 776, 293, 79, 416, 328, 250, 431, 337, 797, 914, 720, 829, 998, 795, 424, 325, 698, 721, 932, 751, 525, 874, 108, 310, 294, 323, 812, 933, 220, 349, 492, 331, 399, 715, 593, 333, 840, 805, 160, 493, 619, 570, 831, 904, 311, 544, 417, 421, 440, 622, 441, 818, 658, 738, 202, 407, 661, 207, 668, 376, 436, 661, 965, 291, 778, 97, 192, 233, 389, 490, 408, 299, 215, 577, 419, 127, 643, 325, 501, 82, 472, 716, 728, 977, 720, 240, 82, 481, 549, 486, 209, 631, 783, 677, 537, 143, 353, 773, 933, 829, 270, 276, 628, 105, 816, 716, 727, 887, 940, 26, 584, 783, 992, 346, 314, 701, 298, 966, 975, 248, 902, 997, 507, 438, 536, 498, 135, 552, 786, 581, 16, 106, 477, 11, 997, 758, 38, 849, 279, 777, 268, 416, 708, 160, 3, 759, 471, 739, 101, 774, 8, 80, 479, 546, 739, 949, 800, 963, 370, 839, 427, 570, 120, 787, 508, 709, 72, 428, 592, 914, 655, 294, 389, 833, 174, 84, 198, 80, 663, 214, 666, 265, 23, 963, 769, 414, 55, 897, 70, 694, 428, 988, 494, 557, 826, 342, 418, 349, 567, 233, 716, 718, 52, 256, 632, 720, 135, 38, 394, 421, 175, 664, 548, 543, 91, 785, 525, 736, 945, 113, 483, 914, 843, 135, 629, 837, 337, 714, 236, 633, 620, 795, 89, 431, 856, 187, 377, 942, 829, 28, 32, 630, 297, 318, 812, 146, 865, 6, 704, 540, 787, 577, 389, 307, 879, 851, 51, 814, 20, 124, 707, 141, 741, 151, 484, 949, 781, 531, 919, 46, 130, 331, 2, 541, 289, 139, 359, 859, 75, 93, 361, 691, 445, 729, 862, 360, 318, 726, 454, 973, 175, 817, 494, 976, 115, 414, 990, 898, 614, 560, 641, 513, 47, 45, 377, 563, 932, 60, 282, 276, 156, 790, 581, 344, 897, 199, 711, 544, 468, 364, 784, 215, 638, 207, 573, 632, 826, 894, 250, 684, 325, 747, 239, 18, 404, 873, 209, 199, 245, 581, 444, 597, 948, 799, 768, 958, 891, 348, 497, 348, 378, 220, 215, 7, 37, 783, 495, 495, 939, 811, 126, 601, 75, 340, 222, 692, 993, 997, 546, 410, 44, 272, 214, 471, 849, 196, 5, 815, 439, 530, 568, 72, 662, 222, 391, 617, 210, 799, 403, 483, 776, 424, 555, 767, 111, 313, 706, 564, 374, 254, 44, 334, 856, 632, 121, 673, 170, 573, 841, 115, 974, 462, 288, 902, 380, 397, 667, 125, 669, 624, 934, 447, 41, 620, 116, 767, 673, 226, 599, 239, 525, 508, 206, 107, 943, 41, 814, 231, 722, 216, 597, 409, 880, 252, 189, 724, 902, 207, 768, 46, 846, 176, 983, 44, 671, 218, 922, 744, 761, 461, 822, 279, 459, 765, 764, 446, 259, 558, 64, 604, 375, 106, 392, 993, 491, 602, 1, 610, 236, 762, 137, 386, 421, 958, 519, 61, 111, 979, 94, 812, 555, 437, 637, 200, 38, 640, 186, 459, 346, 7, 698, 687, 685, 497, 19, 815, 709, 825, 897, 454, 935, 110, 967, 795, 12, 85, 481, 880, 680, 821, 524, 433, 963, 5, 844, 697, 226, 792, 837, 814, 981, 324, 788, 23, 493, 84, 725, 811, 27, 258, 0, 462, 372, 232, 631, 184, 648, 649, 405, 343, 742, 244, 254, 280, 540, 672, 696, 653, 107, 344, 238, 77, 791, 893, 411, 332, 685, 961, 88, 940, 802, 19, 0, 942, 605, 633, 147, 458, 146, 550, 940, 841, 503, 910, 377, 17, 186, 928, 189, 148, 691, 290, 244, 272, 4, 864, 862, 963, 151, 749, 865, 829, 251, 836, 426, 294, 658, 476, 933, 561, 808, 433, 68, 127, 270, 354, 614, 465, 453, 431, 593, 275, 202, 56, 613, 714, 672, 738, 257, 60, 304, 540, 523, 79, 831, 158, 223, 112, 252, 300, 989, 903, 514, 721, 835, 426, 919, 246, 902, 129, 943, 198, 944, 440, 66, 966, 128, 422, 809, 719, 611, 984, 702, 386, 20, 533, 607, 269, 488, 593, 526, 945, 493, 950, 780, 803, 203, 835, 358, 785, 171, 20, 569, 149, 736, 668, 925, 338, 525, 976, 531, 273, 268, 724, 613, 641, 997, 815, 402, 725, 240, 138, 660, 561, 76, 376, 311, 568, 59, 588, 515, 113, 129, 405, 655, 697, 575, 624, 98, 674, 716, 948, 442, 347, 172, 676, 321, 320, 457, 475, 972, 93, 986, 454, 47, 201, 103, 321, 674, 518, 804, 779, 59, 2, 212, 293, 733, 633, 48, 341, 816, 252, 800, 590, 575, 163, 913, 722, 305, 145, 284, 331, 349, 138, 323, 257, 233, 703, 346, 175, 916, 880, 759, 921, 28, 992, 876, 870, 417, 138, 424, 470, 814, 857, 803, 259, 895, 933, 423, 545, 251, 842, 67, 410, 455, 914, 644, 869, 587, 426, 851, 269, 668, 789, 325, 203, 41, 64, 638, 804, 386, 764, 622, 286, 366, 13, 299, 911, 851, 107, 655, 778, 909, 202, 487, 542, 333, 113, 531, 608, 39, 377, 805, 269, 550, 513, 827, 495, 618, 198, 527, 223, 919, 149, 781, 318, 49, 759, 677, 881, 350, 824, 64, 131, 865, 927, 633, 787, 28, 418, 697, 902, 130, 831, 321, 729, 763, 808, 982, 126, 678, 517, 970, 441, 58, 690, 424, 53, 61, 876, 624, 355, 711, 780, 577, 316, 207, 51, 759, 872, 824, 311, 230, 650, 463, 547, 285, 158, 696, 519, 764, 751, 583, 137, 911, 332, 721, 529, 595, 894, 330, 80, 237, 979, 136, 888, 380, 431, 968, 700, 287, 547, 521, 266, 345, 75, 56, 739, 163, 664, 710, 983, 304, 35, 867, 961, 618, 865, 930, 271, 387, 824, 116, 405, 914, 531, 275, 48, 848, 361, 518, 116, 980, 735, 490, 429, 319, 502, 964, 319, 486, 948, 211, 634, 796, 118, 755, 699, 101, 658, 791, 153, 252, 59, 16, 442, 226, 524, 550, 186, 837, 935, 76, 55, 350, 992, 389, 433, 195, 868, 651, 978, 97, 678, 495, 361, 466, 385, 93, 751, 559, 855, 317, 361, 841, 393, 648, 307, 84, 94, 842, 499, 955, 402, 867, 631, 878, 703, 615, 745, 526, 493, 263, 409, 460, 894, 192, 192, 900, 497, 422, 239, 255, 430, 809, 333, 562, 43, 357, 485, 747, 24, 134, 29, 171, 882, 713, 956, 88, 692, 625, 293, 753, 164, 270, 260, 311, 915, 82, 700, 229, 37, 323, 358, 200, 297, 778, 477, 105, 100, 962, 238, 82, 158, 715, 332, 641, 913, 8, 978, 863, 556, 891, 860, 705, 359, 417, 801, 23, 776, 683, 244, 111, 204, 423, 403, 519, 138, 855, 776, 692, 189, 220, 225, 96, 116, 972, 854, 49, 852, 746, 646, 385, 956, 674, 429, 112, 185, 529, 357, 225, 218, 747, 472, 600, 744, 897, 966, 888, 717, 718, 344, 972, 723, 712, 90, 739, 284, 118, 383, 855, 679, 961, 784, 53, 475, 172, 445, 763, 587, 722, 527, 547, 381, 947, 890, 694, 944, 213, 201, 282, 882, 546, 655, 315, 226, 621, 136, 535, 101, 410, 993, 695, 513, 131, 715, 155, 232, 811, 501, 457, 48, 720, 210, 734, 888, 343, 850, 951, 887, 474, 383, 692, 375, 403, 53, 119, 166, 890, 296, 427, 95, 301, 805, 87, 318, 809, 843, 172, 770, 904, 631, 135, 138, 401, 762, 136, 599, 40, 80, 653, 591, 232, 237, 139, 505, 524, 303, 136, 550, 981, 618, 775, 951, 589, 82, 346, 125, 941, 644, 475, 692, 656, 739, 615, 675, 478, 58, 84, 113, 924, 384, 235, 470, 408, 46, 152, 133, 791, 994, 306, 900, 348, 257, 552, 188, 139, 948, 86, 518, 204, 674, 754, 547, 116, 777, 981, 921, 724, 351, 916, 940, 806, 310, 829, 232, 680, 427, 268, 717, 28, 867, 520, 975, 704, 60, 526, 850, 729, 536, 250, 466, 851, 446, 355, 727, 888, 377, 208, 320, 985, 356, 218, 927, 317, 19, 149, 410, 864, 267, 425, 862, 162, 339, 768, 941, 622, 519, 855, 318, 159, 402, 259, 563, 249, 781, 186, 417, 190, 868, 553, 700, 891, 272, 543, 873, 610, 921, 44, 734, 483, 876, 595, 337, 805, 826, 40, 332, 373, 972, 304, 183, 268, 764, 267, 448, 280, 336, 527, 937, 995, 273, 231, 407, 573, 721, 143, 402, 839, 848, 268, 678, 733, 365, 999, 215, 763, 93, 13, 852, 884, 633, 518, 671, 994, 871, 376, 411, 82, 884, 709, 68, 485, 927, 700, 692, 62, 81, 885, 707, 535, 252, 415, 474, 98, 482, 343, 968, 270, 786, 594, 8, 897, 272, 267, 512, 25, 370, 698, 253, 791, 394, 667, 814, 944, 944, 688, 821, 225, 658, 214, 245, 412, 59, 625, 422, 183, 681, 706, 270, 457, 890, 831, 26, 630, 188, 726, 539, 482, 567, 16, 323, 652, 589, 150, 961, 931, 834, 165, 552, 183, 880, 267, 814, 835, 131, 51, 27, 286, 296, 357, 775, 63, 38, 570, 624, 946, 264, 254, 568, 147, 535, 125, 910, 339, 638, 865, 321, 97, 341, 156, 95, 196, 645, 978, 841, 533, 526, 400, 650, 664, 753, 472, 345, 54, 672, 144, 66, 246, 458, 835, 646, 193, 134, 56, 715, 126, 671, 957, 782, 850, 384, 917, 179, 182, 69, 547, 276, 585, 665, 672, 5, 539, 18, 710, 553, 680, 385, 375, 727, 423, 96, 507, 806, 85, 203, 232, 457, 618, 356, 682, 405, 690, 320, 824, 72, 843, 774, 727, 944, 537, 559, 711, 119, 392, 839, 86, 937, 514, 608, 442, 525, 427, 180, 600, 613, 563, 428, 323, 679, 978, 671, 669, 866, 804, 61, 855, 407, 510, 248, 369, 648, 227, 264, 250, 845, 115, 306, 206, 222, 501, 584, 837, 181, 786, 559, 325, 353, 337, 671, 821, 122, 706, 123, 993, 451, 118, 81, 929, 372, 224, 104, 780, 482, 822, 568, 167, 135, 881, 66, 891, 778, 74, 492, 642, 375, 698, 603, 637, 55, 995, 703, 788, 402, 277, 341, 24, 591, 583, 745, 637, 128, 27, 719, 189, 678, 870, 720, 549, 649, 858, 566, 86, 532, 10, 421, 498, 633, 461, 60, 888, 596, 223, 24, 885, 609, 822, 875, 430, 738, 382, 961, 892, 953, 940, 730, 15, 50, 492, 46, 556, 850, 325, 965, 859, 280, 911, 15, 416, 658, 661, 818, 878, 770, 710, 243, 444, 750, 187, 701, 99, 955, 906, 870, 715, 814, 255}))
	fmt.Println("Max Profit", MP.FastSolver([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("Max Profit", MP.FastSolver([]int{2, 4, 5, 6, 1, 2, 3}))

	AB := &vidio.BinarySquare{}
	AB.Solver(5)

	SQRT := &leetcode.SQRT{}
	fmt.Println("SQRT", SQRT.Solver(10))

	CS := &leetcode.ClimbingStairs{}
	fmt.Println("Climbing Stairs", CS.Solver(3))
	fmt.Println("Climbing Stairs", CS.Solver(5))
	fmt.Println("Climbing Stairs", CS.FastSolver(4))
	fmt.Println("Climbing Stairs", CS.FastSolver(7))

	RSL := &leetcode.RemoveDuplicateSortedList{}
	l3 := leetcode.ListNode{
		Val:  2,
		Next: nil,
	}
	l2 := leetcode.ListNode{
		Val:  1,
		Next: &l3,
	}
	l1 := &leetcode.ListNode{
		Val:  1,
		Next: &l2,
	}

	fmt.Println("Remove Duplicated Sorted List", RSL.Solver(l1))
	fmt.Println("Remove Duplicated Sorted List", RSL.Solver(nil))
	fmt.Println("Remove Duplicated Sorted List", RSL.FastSolver(nil))
	fmt.Println("Remove Duplicated Sorted List", RSL.FastSolver(l1))

	PDB := &leetcode.PartitioningDeciBinary{}
	fmt.Println("Partitioning Deci-Binary", PDB.Solver("32"))
	fmt.Println("Partitioning Deci-Binary", PDB.FastSolver("32"))
	fmt.Println("Partitioning Deci-Binary", PDB.FastSolver("82734"))
	fmt.Println("Partitioning Deci-Binary", PDB.FastSolver("27346209830709182346"))

	IP := &leetcode.IPAddress{}
	// var input string
	fmt.Print("Input for IPAddress: ")
	// fmt.Scanf("%s", &input)
	fmt.Println("IP Address", IP.Solver("1.1"))
	fmt.Println("IP Address", IP.FastSolver("1.1"))

	NL := &leetcode.NewLanguage{}
	fmt.Println("New Language:", NL.Solver([]string{"--X", "X++", "X++"}))

	LS := &leetcode.LongestSentence{}
	fmt.Println("Longest Sentence", LS.Solver([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))

	GP := &leetcode.GoalParser{}
	fmt.Println("Goal Parser", GP.Solver("G()(al)"))
	fmt.Println("Goal Parser", GP.Solver("(al)G(al)()()G"))

	// EncDecURL := &leetcode.EncodeAndDecodeURL{}
	// fmt.Println("Encode and Decode URL:")
	// EncDecURL.Solver("https://leetcode.com/problems/design-tinyurl")

	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Binary Search", leetcode.BinarySearch(arr, 3, 0, len(arr)-1))

	SS := &leetcode.ShuffleString{}
	fmt.Println("Shuffle String", SS.Solver("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
	fmt.Println("Shuffle String", SS.FastSolver("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
	fmt.Println("Shuffle String", SS.FastSolver("ab", []int{1, 0}))

	ES := &leetcode.ExcelSheet{}
	fmt.Println("Excel String", ES.WrongSolver("A9:B3"))
	fmt.Println("Excel String", ES.Solver("A7:K8"))

	MBB := &leetcode.MoveBallToBox{}
	fmt.Println("Move Ball To Box", MBB.Solver("110"))
	fmt.Println("Move Ball To Box", MBB.Solver("001011"))
	fmt.Println("Move Ball To Box", MBB.FastSolver("001011"))

	AP := &leetcode.ArrayPermuatation{}
	fmt.Println("Array Permutation", AP.Solver([]int{0, 2, 1, 5, 3, 4}))
	fmt.Println("Array Permutation", AP.Solver([]int{5, 0, 1, 2, 3, 4}))
	fmt.Println("Array Permutation", AP.OtherSolver([]int{5, 0, 1, 2, 3, 4}))

	RSA := &leetcode.RunningSumArray{}
	fmt.Println("Running Sum Array", RSA.Solver([]int{1, 2, 3, 4}))
	fmt.Println("Running Sum Array", RSA.Solver([]int{1, 1, 1, 1, 1}))
	fmt.Println("Running Sum Array", RSA.Solver([]int{1}))

	SA := &leetcode.ShuffleArray{}
	fmt.Println("Shuffle Array", SA.Solver([]int{2, 5, 1, 3, 4, 7}, 3))

	CG := &leetcode.CollectGarbage{}
	fmt.Println("Collect Garbage", CG.Solver([]string{"G", "M", "P"}, []int{1, 3}))
	fmt.Println("Collect Garbage", CG.Solver([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}))

	DM := &leetcode.DecodeMessage{}
	fmt.Println("Decode Message Solver:", DM.Solver("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
	fmt.Println("Decode Message Solver:", DM.Solver("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"))
	fmt.Println("Decode Message Solver:", DM.FastSolver("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"))

	KCS := &leetcode.KeepCitySkyline{}
	fmt.Println("Keep City Skyline Solver:", KCS.Solver([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}))
	fmt.Println("Keep City Skyline Solver:", KCS.Solver([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}))

	DXORedArray := &leetcode.DecodeXORedArray{}
	fmt.Println("Decode XORed Array Solver:", DXORedArray.Solver([]int{1, 2, 3}, 1))
	fmt.Println("Decode XORed Array Solver:", DXORedArray.Solver([]int{6, 2, 7, 3}, 4))

	SIG := &leetcode.StayingInGrid{}
	fmt.Println("Stay In Grid:", SIG.Solver(3, []int{0, 1}, "RRDDLU"))

	MMTS := &leetcode.MinMovesToSeat{}
	fmt.Println("Min Moves To Seat:", MMTS.Solver([]int{3, 1, 5}, []int{2, 7, 4}))
	fmt.Println("Min Moves To Seat:", MMTS.Solver([]int{4, 1, 5, 9}, []int{1, 3, 2, 6}))
	fmt.Println("Min Moves To Seat:", MMTS.Solver([]int{2, 2, 6, 6}, []int{1, 3, 2, 6}))

	fmt.Println("------------------Interview-------------------")
	Test := &interviews.Test{}
	fmt.Println("Test:", Test.Tes())

	B := &interviews.Beli{}
	B.MulaiBeli([]string{"minyak", "bawang merah", "minyak"})

}
