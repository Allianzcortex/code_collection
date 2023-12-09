package adventofcode

import "testing"

var input9 = []string{
	"23 34 50 94 210 482 1077 2327 4864 9829 19209 36438 67557 123497 224466 408029 743317 1354940 2461668 4436846 7899892",
	"7 9 15 26 43 67 99 140 191 253 327 414 515 631 763 912 1079 1265 1471 1698 1947",
	"12 32 74 154 291 503 814 1298 2210 4286 9343 21393 48633 106944 226041 460406 909128 1752848 3324280 6245160 11693791",
	"5 18 46 106 235 508 1066 2167 4287 8327 16052 31043 60763 120970 244922 502099 1035401 2133556 4369562 8861092 17752611",
	"27 50 81 131 223 400 743 1399 2619 4806 8573 14811 24767 40132 63139 96671 144379 210810 301545 423347 584319",
	"7 20 53 118 222 367 553 793 1158 1879 3542 7421 16003 33768 68296 131782 243049 430158 733723 1211048 1941212",
	"7 5 12 34 74 135 240 495 1245 3405 9085 22675 52646 114551 236281 467922 901195 1708405 3219482 6070032 11475986",
	"5 13 33 65 109 165 233 313 405 509 625 753 893 1045 1209 1385 1573 1773 1985 2209 2445",
	"8 -2 -17 -37 -52 -23 148 661 1876 4388 9121 17443 31304 53399 87358 137965 211408 315562 460307 657883 923284",
	"15 21 44 102 223 460 922 1844 3736 7677 15877 32763 67136 136539 276119 556377 1117965 2238201 4454922 8791220 17153221",
	"3 16 34 67 142 305 618 1152 1977 3157 4788 7200 11632 22063 49558 121622 300847 719834 1643270 3573485 7423202",
	"11 17 36 88 213 494 1087 2251 4375 8024 14095 24327 42724 79084 157087 331858 726650 1607251 3529578 7627561 16167592",
	"-1 0 21 78 198 427 846 1608 3017 5678 10763 20486 39010 74316 142220 275028 539771 1075368 2165681 4378150 8818318",
	"18 22 38 82 181 381 753 1405 2517 4425 7789 13889 25102 45622 82494 147042 256780 437904 728472 1182388 1874315",
	"0 1 21 81 210 445 831 1421 2276 3465 5065 7161 9846 13221 17395 22485 28616 35921 44541 54625 66330",
	"9 27 58 113 216 424 873 1870 4069 8807 18742 39037 79480 158128 307321 583238 1080569 1954363 3451690 5956433 10051312",
	"0 15 43 82 139 256 556 1322 3137 7150 15616 33032 68536 140880 288411 588317 1192156 2388619 4709831 9104648 17212191",
	"4 2 12 55 158 353 677 1181 1958 3197 5262 8782 14720 24366 39171 60306 87792 119004 146304 153505 110810",
	"6 21 45 69 80 61 -9 -155 -406 -795 -1359 -2139 -3180 -4531 -6245 -8379 -10994 -14155 -17931 -22395 -27624",
	"19 32 45 58 71 84 97 110 123 136 149 162 175 188 201 214 227 240 253 266 279",
	"10 13 24 50 101 199 404 873 1976 4501 9988 21240 43067 83327 154336 274727 471846 784781 1268128 1996606 3070641",
	"0 3 7 18 51 129 282 546 962 1575 2433 3586 5085 6981 9324 12162 15540 19499 24075 29298 35191",
	"17 31 45 59 73 87 101 115 129 143 157 171 185 199 213 227 241 255 269 283 297",
	"8 10 12 14 16 18 20 22 24 26 28 30 32 34 36 38 40 42 44 46 48",
	"1 10 39 114 273 568 1066 1847 3012 4748 7561 12914 24758 52917 120106 275652 620853 1353401 2841486 5742580 11189623",
	"14 17 17 14 8 -1 -13 -28 -46 -67 -91 -118 -148 -181 -217 -256 -298 -343 -391 -442 -496",
	"10 18 23 25 24 20 13 3 -10 -26 -45 -67 -92 -120 -151 -185 -222 -262 -305 -351 -400",
	"7 21 35 40 27 5 46 380 1577 4877 12782 30145 66241 138771 281571 559245 1094594 2119788 4072116 7774237 14774706",
	"16 23 49 115 254 527 1061 2136 4362 9015 18660 38309 77606 155029 306115 599783 1171024 2286635 4477211 8800303 17358594",
	"6 12 17 21 24 26 27 27 26 24 21 17 12 6 -1 -9 -18 -28 -39 -51 -64",
	"8 13 38 102 241 526 1092 2186 4257 8134 15369 28861 53924 100017 183417 331187 586870 1018427 1729032 2871440 4666755",
	"18 30 55 98 178 340 676 1361 2711 5278 10014 18564 33789 60676 107865 190115 332144 574414 981593 1654614 2747468",
	"18 20 21 22 34 89 257 674 1586 3414 6845 12954 23362 40435 67529 109286 171986 263960 396069 582254 840162",
	"20 35 60 100 175 341 727 1591 3407 7019 13956 27133 52437 102235 202842 409760 837515 1715946 3493279 7021299 13886128",
	"15 25 44 79 155 343 810 1895 4220 8855 17578 33324 61032 109319 193790 343400 612197 1100085 1988056 3595765 6472489",
	"19 38 71 118 179 254 343 446 563 694 839 998 1171 1358 1559 1774 2003 2246 2503 2774 3059",
	"10 24 49 99 199 402 824 1713 3579 7433 15232 30745 61319 121575 241175 481021 966676 1956580 3979811 8111053 16511955",
	"19 31 57 103 175 279 421 607 843 1135 1489 1911 2407 2983 3645 4399 5251 6207 7273 8455 9759",
	"0 -1 4 24 87 253 640 1478 3206 6627 13136 25036 45957 81393 139372 231274 372812 585191 896460 1343072 1971667",
	"10 36 75 127 192 270 361 465 582 712 855 1011 1180 1362 1557 1765 1986 2220 2467 2727 3000",
	"-6 -5 -2 12 69 245 687 1656 3603 7306 14121 26447 48580 88239 159197 285682 509691 903614 1594019 2811443 5002638",
	"-1 -1 0 4 21 82 254 660 1510 3152 6155 11439 20470 35541 60163 99593 161529 257005 401522 616454 930771",
	"14 21 40 91 194 361 591 871 1192 1618 2508 5092 12744 33476 84391 199070 439114 911299 1792008 3360753 6044662",
	"11 22 47 97 194 371 672 1152 1877 2924 4381 6347 8932 12257 16454 21666 28047 35762 44987 55909 68726",
	"1 -6 -20 -31 -5 141 579 1630 3858 8230 16404 31252 57788 104755 187233 330765 577663 996354 1694860 2839779 4682449",
	"7 8 10 20 45 88 156 295 667 1684 4214 9874 21425 43284 82168 147885 254287 420400 671746 1041872 1574101",
	"20 40 65 95 130 170 215 265 320 380 445 515 590 670 755 845 940 1040 1145 1255 1370",
	"21 31 40 48 53 57 85 218 635 1647 3701 7367 13449 23654 42799 84428 183024 417769 955047 2119794 4510308",
	"6 25 58 105 166 241 330 433 550 681 826 985 1158 1345 1546 1761 1990 2233 2490 2761 3046",
	"19 34 70 153 321 623 1112 1827 2763 3834 4847 5535 5753 6022 8685 19939 52788 130387 288488 576099 1059512",
	"0 -7 -16 -22 -9 59 247 658 1442 2805 5018 8426 13457 20631 30569 44002 61780 84881 114420 151658 198011",
	"8 10 17 34 73 164 376 847 1827 3757 7444 14462 28051 55078 110200 224438 462238 954184 1957407 3963142 7881761",
	"20 34 57 89 130 180 239 307 384 470 565 669 782 904 1035 1175 1324 1482 1649 1825 2010",
	"3 4 9 17 27 38 49 59 67 72 73 69 59 42 17 -17 -61 -116 -183 -263 -357",
	"-3 -3 4 30 91 219 486 1046 2204 4519 8951 17097 31687 57858 106570 203428 410155 873797 1933322 4339084 9684335",
	"10 25 40 55 70 85 100 115 130 145 160 175 190 205 220 235 250 265 280 295 310",
	"-4 7 33 80 157 272 422 571 607 266 -992 -4173 -11026 -24431 -48925 -91396 -161978 -275183 -451309 -718166 -1113165",
	"8 16 39 93 198 378 661 1079 1668 2468 3523 4881 6594 8718 11313 14443 18176 22584 27743 33733 40638",
	"3 7 20 59 156 383 906 2082 4613 9771 19708 37865 69494 122307 207266 339528 539559 834431 1259316 1859191 2690768",
	"9 30 55 79 102 142 254 566 1367 3319 7914 18366 41259 89580 188492 386780 780041 1557464 3094013 6126141 12074720",
	"11 22 51 109 212 373 595 886 1342 2384 5301 13362 33956 82595 190387 418201 884049 1816682 3660479 7278227 14335104",
	"19 22 26 35 62 136 318 749 1775 4227 9978 22954 50842 107815 218682 424970 793555 1428580 2487530 4202477 6907662",
	"21 31 45 68 100 139 194 325 737 1960 5143 12466 27624 56293 106580 190008 325237 550617 959681 1788700 3611684",
	"-4 5 32 90 203 415 814 1586 3126 6256 12634 25483 50825 99472 190104 353853 640912 1129799 1940028 3249072 5314647",
	"13 34 66 106 155 238 434 929 2114 4764 10362 21685 43875 86463 167418 321730 620219 1210847 2407738 4877617 10018867",
	"15 40 88 175 333 636 1247 2495 4992 9801 18667 34324 60892 104379 173304 279458 438821 672654 1008786 1483117 2141359",
	"9 15 37 84 164 289 479 771 1250 2130 3924 7753 15855 32366 64456 123914 229287 408689 703407 1172442 1898134",
	"9 26 55 117 251 530 1096 2234 4517 9074 18070 35550 68900 131341 246150 453809 824268 1478552 2628474 4653755 8260732",
	"17 18 15 4 -13 -10 90 449 1359 3325 7268 15036 30585 62440 128345 263285 532179 1048390 1998963 3678322 6537405",
	"21 28 35 42 49 56 63 70 77 84 91 98 105 112 119 126 133 140 147 154 161",
	"22 41 69 120 231 472 954 1844 3416 6207 11424 21894 44129 92630 198667 428027 916737 1942612 4067342 8415103 17210070",
	"10 4 9 47 159 425 996 2137 4282 8113 14703 25819 44590 76974 134963 243555 455734 883874 1763329 3576095 7281317",
	"0 9 28 57 96 145 204 273 352 441 540 649 768 897 1036 1185 1344 1513 1692 1881 2080",
	"4 4 10 27 72 179 400 804 1476 2512 3989 5857 7651 7858 2727 -15596 -61125 -153822 -309483 -500209 -540831",
	"2 11 36 92 206 423 812 1472 2538 4187 6644 10188 15158 21959 31068 43040 58514 78219 102980 133724 171486",
	"7 13 27 73 191 453 991 2042 4034 7776 14888 28737 56364 112236 225186 450672 891561 1732103 3291689 6107481 11058165",
	"12 40 92 177 299 453 621 768 838 750 394 -373 -1731 -3901 -7149 -11790 -18192 -26780 -38040 -52523 -70849",
	"2 10 36 91 202 431 909 1893 3855 7616 14545 26853 48025 83449 141320 233919 379392 604182 946298 1459639 2219628",
	"16 29 65 143 282 494 787 1185 1768 2732 4467 7650 13350 23143 39237 64610 103168 159935 241293 355297 512098",
	"10 5 12 38 96 220 487 1059 2282 4931 10788 23917 53323 118277 258683 554867 1164838 2393797 4822956 9546940 18608510",
	"11 24 50 99 183 327 594 1138 2302 4788 9954 20347 40665 79456 152009 285077 524313 945636 1672246 2899791 4933431",
	"1 7 33 99 232 470 883 1622 3013 5733 11152 22022 43880 87868 176248 352827 701969 1382067 2682539 5117917 9578800",
	"8 4 -2 -3 27 147 473 1223 2795 5897 11776 22662 42674 79652 148739 279218 527591 1003263 1920647 3706984 7225286",
	"8 9 11 16 45 163 513 1361 3153 6581 12649 22723 38542 62161 95793 141516 200814 273929 359015 451106 540939",
	"13 16 14 10 9 14 23 34 84 384 1667 5941 17938 47682 114784 255368 533109 1057155 2012660 3716109 6720753",
	"13 23 42 77 147 297 614 1246 2441 4667 8966 17865 37461 81790 181409 399474 860801 1802863 3662826 7221761 13835595",
	"24 33 39 50 87 186 405 841 1671 3258 6416 13015 27236 57965 123052 256464 519738 1019599 1934157 3550744 6319205",
	"12 25 46 90 192 414 860 1709 3277 6126 11266 20576 37758 70530 135517 268659 546296 1125970 2322222 4742431 9522654",
	"0 1 -3 -14 -33 -60 -94 -133 -174 -213 -245 -264 -263 -234 -168 -55 116 357 681 1102 1635",
	"-4 7 26 50 83 141 270 590 1387 3309 7796 18018 40870 91096 199609 429975 910670 1897655 3894826 7883742 15756517",
	"21 24 37 69 127 230 435 875 1809 3684 7209 13441 23883 40594 66311 104583 159917 237936 345549 491133 684727",
	"6 4 -1 1 34 148 444 1112 2482 5088 9745 17639 30430 50368 80422 124422 187214 274828 394659 555661 768554",
	"23 40 66 101 145 198 260 331 411 500 598 705 821 946 1080 1223 1375 1536 1706 1885 2073",
	"14 16 29 70 165 350 675 1231 2244 4314 8934 19546 43664 97163 212917 457876 965814 1997888 4050487 8040434 15612411",
	"4 13 43 102 195 328 530 917 1841 4196 9996 23406 52516 112337 229827 452313 861586 1597375 2896065 5153678 9026609",
	"-8 -1 25 77 164 297 492 782 1244 2047 3527 6295 11384 20441 35970 61632 102608 166031 261493 401633 602812",
	"24 37 55 79 115 180 314 604 1234 2601 5588 12166 26609 57758 123022 255360 515860 1017797 1975233 3803805 7328983",
	"12 14 23 61 175 445 998 2039 3924 7327 13593 25419 48056 91256 172164 319258 577398 1014674 1734881 2909539 4869847",
	"12 10 19 51 124 273 579 1231 2648 5714 12236 25864 54002 111875 231267 479212 997404 2083634 4356277 9078801 18785272",
	"-2 -11 -17 -13 5 49 169 507 1380 3397 7615 15739 30371 55313 95929 159571 256074 398325 602911 890851 1288417",
	"3 1 -4 -3 32 153 441 1011 2017 3657 6178 9881 15126 22337 32007 44703 61071 81841 107832 139957 179228",
	"3 5 24 76 193 433 900 1791 3497 6802 13259 25907 50693 99390 195638 387278 770904 1539404 3071735 6098927 12005782",
	"7 15 42 95 182 316 512 779 1116 1541 2240 4056 9795 27268 75688 198084 483877 1107795 2394006 4917853 9662028",
	"3 15 33 60 99 152 224 335 555 1095 2507 6064 14403 32516 69162 138743 263635 476887 825093 1371100 2196035",
	"14 19 29 53 108 219 416 728 1174 1751 2419 3083 3572 3615 2814 614 -3730 -11189 -22999 -40703 -66196",
	"19 47 96 171 270 385 522 756 1339 2881 6626 14847 31386 62367 117112 209292 358347 591211 944380 1466363 2220558",
	"13 17 32 67 148 344 820 1942 4475 9944 21289 44066 88659 174301 336168 637393 1188485 2177203 3912210 6883485 11841026",
	"11 22 46 93 177 329 615 1155 2138 3825 6522 10483 15664 21188 24293 18415 -10099 -87352 -260282 -609319 -1266917",
	"5 16 42 93 188 369 722 1418 2810 5668 11714 24759 53000 113516 240922 503944 1037329 2104085 4218930 8396808 16657278",
	"13 22 30 34 33 26 4 -57 -179 -287 81 2511 11446 38112 109202 285643 702073 1647364 3723532 8148980 17318427",
	"-10 -16 -15 9 80 230 499 935 1594 2540 3845 5589 7860 10754 14375 18835 24254 30760 38489 47585 58200",
	"25 48 79 118 165 220 283 354 433 520 615 718 829 948 1075 1210 1353 1504 1663 1830 2005",
	"17 33 66 114 169 221 276 393 742 1698 4035 9394 21409 48259 108122 240351 527839 1141439 2428576 5091846 10552481",
	"17 21 25 29 33 37 41 45 49 53 57 61 65 69 73 77 81 85 89 93 97",
	"24 31 39 48 54 49 21 -46 -172 -381 -701 -1164 -1806 -2667 -3791 -5226 -7024 -9241 -11937 -15176 -19026",
	"24 32 35 33 26 14 -3 -25 -52 -84 -121 -163 -210 -262 -319 -381 -448 -520 -597 -679 -766",
	"-4 -11 -26 -48 -57 -1 217 752 1830 3761 6952 11920 19305 29883 44579 64480 90848 125133 168986 224272 293083",
	"25 49 86 151 280 550 1119 2294 4635 9103 17260 31529 55522 94444 155581 248880 387629 589245 876178 1276939 1827260",
	"21 40 80 166 333 631 1135 1960 3281 5358 8566 13430 20665 31221 46333 67576 96925 136820 190236 260758 352661",
	"1 14 42 90 162 255 349 393 287 -140 -1156 -3156 -6696 -12531 -21657 -35357 -55251 -83350 -122114 -174514 -244098",
	"11 36 72 123 194 294 458 797 1583 3374 7182 14685 28482 52388 91764 153875 248267 387152 585788 862839 1240698",
	"0 0 15 55 130 250 425 665 980 1380 1875 2475 3190 4030 5005 6125 7400 8840 10455 12255 14250",
	"8 18 36 57 73 67 16 -82 -125 294 2321 8719 25582 65633 154559 343189 730803 1509265 3045463 6033724 11776940",
	"-6 8 45 117 255 527 1073 2181 4453 9156 18932 39171 80556 163599 326438 637799 1217892 2271160 4136295 7360841 12810093",
	"-5 -7 6 61 210 541 1199 2438 4747 9126 17628 34331 66985 129775 248135 467680 871677 1616059 2999641 5603532 10565076",
	"16 20 26 45 102 241 537 1116 2189 4114 7509 13460 23936 42715 77591 145645 285394 583456 1230208 2630623 5615774",
	"16 27 47 76 114 161 217 282 356 439 531 632 742 861 989 1126 1272 1427 1591 1764 1946",
	"26 44 80 156 300 550 965 1643 2746 4532 7394 11906 18876 29406 44959 67433 99242 143404 203636 284456 391292",
	"3 3 17 56 143 329 710 1445 2775 5043 8715 14402 22883 35129 52328 75911 107579 149331 203493 272748 360167",
	"8 25 61 139 309 661 1334 2526 4517 7724 12814 20908 33916 55050 89569 145817 236622 381131 607163 954169 1476895",
	"17 30 42 44 27 -8 -27 76 541 1851 4949 11704 25947 55791 118750 252688 538313 1143420 2406244 4986244 10126859",
	"27 41 55 69 83 97 111 125 139 153 167 181 195 209 223 237 251 265 279 293 307",
	"19 38 72 139 263 485 903 1751 3536 7282 15003 30691 62449 127086 259832 534432 1103901 2281810 4701898 9626407 19531031",
	"7 14 32 58 84 97 79 7 -147 -416 -838 -1456 -2318 -3477 -4991 -6923 -9341 -12318 -15932 -20266 -25408",
	"-1 13 36 65 99 139 188 251 335 449 604 813 1091 1455 1924 2519 3263 4181 5300 6649 8259",
	"19 29 31 27 21 19 30 75 231 764 2438 7135 19006 46526 106073 227986 466399 914279 1724577 3136463 5501023",
	"8 14 38 95 203 383 659 1058 1610 2348 3308 4529 6053 7925 10193 12908 16124 19898 24290 29363 35183",
	"-9 -18 -19 -2 42 134 336 799 1827 3947 7969 15027 26651 45129 74986 127739 234931 480152 1071709 2497940 5843812",
	"1 2 6 20 57 149 367 848 1829 3688 6992 12552 21485 35283 55889 85780 128057 186542 265882 371660 510513",
	"19 26 49 115 277 640 1410 2974 6028 11792 22384 41477 75463 135555 241668 429671 764877 1366666 2452191 4411515 7932621",
	"18 18 12 -2 -26 -62 -112 -178 -262 -366 -492 -642 -818 -1022 -1256 -1522 -1822 -2158 -2532 -2946 -3402",
	"20 47 92 163 281 493 889 1625 2954 5267 9146 15431 25303 40385 62863 95629 142448 208151 298856 422219 587717",
	"8 25 63 142 290 541 933 1506 2300 3353 4699 6366 8374 10733 13441 16482 19824 23417 27191 31054 34890",
	"-3 10 33 61 89 112 125 123 101 54 -23 -135 -287 -484 -731 -1033 -1395 -1822 -2319 -2891 -3543",
	"0 6 34 97 214 418 769 1377 2440 4302 7536 13057 22270 37258 61015 97729 153120 234838 352926 520353 753622",
	"3 6 21 61 161 390 865 1780 3476 6607 12522 24112 47605 96194 197051 404364 822757 1648136 3237095 6222123 11698781",
	"-3 6 24 52 102 210 465 1062 2386 5146 10618 21150 41290 80351 158174 317705 651388 1354197 2825776 5860072 11987885",
	"-6 1 25 68 128 196 258 321 490 1123 3092 8203 19921 44787 95443 197247 402536 820630 1679579 3451173 7100776",
	"16 28 52 87 138 238 491 1150 2759 6421 14320 30749 64118 130805 262403 519191 1015136 1964740 3772294 7204873 13731480",
	"17 29 45 63 93 169 357 759 1519 2853 5169 9451 18313 38575 86999 202117 467101 1051637 2285101 4779391 9628013",
	"17 24 30 38 66 166 448 1107 2455 4974 9440 17236 31101 56807 106721 207074 412313 832592 1684894 3385342 6711128",
	"5 12 21 37 65 110 177 271 397 560 765 1017 1321 1682 2105 2595 3157 3796 4517 5325 6225",
	"15 33 67 137 276 542 1055 2084 4236 8841 18684 39306 81179 163152 317662 598301 1090421 1925537 3300345 5501199 8934878",
	"17 33 56 88 146 275 564 1183 2471 5124 10576 21774 44787 92171 189935 391680 806673 1654423 3368668 6791624 13529572",
	"14 20 29 41 57 85 155 359 952 2587 6822 17155 41070 94032 207286 443115 924585 1892777 3814446 7578544 14842554",
	"19 33 47 61 75 89 103 117 131 145 159 173 187 201 215 229 243 257 271 285 299",
	"9 35 74 126 191 269 360 464 581 711 854 1010 1179 1361 1556 1764 1985 2219 2466 2726 2999",
	"4 1 -3 -8 -14 -21 -29 -38 -48 -59 -71 -84 -98 -113 -129 -146 -164 -183 -203 -224 -246",
	"8 22 49 111 240 488 962 1896 3776 7544 14922 28916 54583 100178 178876 311485 531164 892610 1491347 2505101 4281102",
	"19 25 35 58 109 213 416 814 1624 3340 7055 15121 32534 69903 149820 320260 681865 1441415 3012595 6198864 12513843",
	"16 24 31 39 50 66 89 121 164 220 291 379 486 614 765 941 1144 1376 1639 1935 2266",
	"-3 2 26 80 175 322 532 816 1185 1650 2222 2912 3731 4690 5800 7072 8517 10146 11970 14000 16247",
	"3 23 67 152 306 588 1137 2277 4718 9906 20588 41671 81467 153429 278496 488178 828525 1365137 2189385 3426026 5242408",
	"4 14 34 69 135 276 589 1253 2558 4930 8948 15349 25017 38952 58215 83845 116744 157526 206326 262565 324667",
	"-2 5 32 87 178 313 500 747 1062 1453 1928 2495 3162 3937 4828 5843 6990 8277 9712 11303 13058",
	"10 14 18 22 26 30 34 38 42 46 50 54 58 62 66 70 74 78 82 86 90",
	"-4 7 31 65 109 181 344 747 1693 3774 8172 17357 36698 78069 167555 361054 774146 1636229 3383653 6812237 13315587",
	"4 3 -1 0 27 122 368 939 2212 4999 11017 23842 50835 106942 221926 453569 910780 1792467 3450595 6489188 11915283",
	"-5 4 19 33 35 10 -61 -201 -437 -800 -1325 -2051 -3021 -4282 -5885 -7885 -10341 -13316 -16877 -21095 -26045",
	"-1 -2 3 28 89 197 347 503 579 416 -245 -1794 -4787 -9985 -18397 -31327 -50425 -77742 -115789 -167600 -236799",
	"4 16 37 72 137 274 567 1154 2230 4036 6829 10828 16131 22598 29695 36294 40424 38968 27301 -1136 -55331",
	"2 13 42 101 214 443 933 1986 4193 8694 17720 35731 71758 144096 289487 580743 1159992 2300323 4517950 8772093 16813322",
	"19 28 42 73 149 325 716 1577 3466 7551 16174 33877 69240 138092 268946 511891 952661 1734206 3088826 5384809 9192551",
	"2 12 44 108 210 343 470 500 267 -468 -1961 -4318 -7079 -8198 -1557 28951 123080 374933 1000696 2486171 5904680",
	"12 21 55 131 278 562 1132 2306 4725 9607 19136 37035 69422 126176 223308 387327 663438 1130801 1930337 3314239 5732378",
	"3 25 68 152 326 684 1385 2692 5067 9394 17455 32857 62708 120511 231130 439616 827811 1547071 2883854 5389714 10131642",
	"2 15 40 71 109 177 336 710 1539 3290 6867 13972 27680 53302 99621 180597 317648 542625 901610 1459677 2306767",
	"2 0 9 39 112 282 671 1538 3403 7244 14765 28683 52886 92176 151206 232390 333580 449312 586517 820327 1443696",
	"3 -2 -7 -12 -17 -22 -27 -32 -37 -42 -47 -52 -57 -62 -67 -72 -77 -82 -87 -92 -97",
	"-4 3 32 100 230 468 924 1844 3714 7393 14277 26519 47387 81978 138840 233850 399463 706075 1309191 2548656 5147939",
	"10 12 8 -1 -4 25 132 388 894 1786 3240 5477 8768 13439 19876 28530 39922 54648 73384 96891 126020",
	"-1 -2 -1 15 76 237 590 1278 2507 4563 7879 13270 22570 40087 75597 150131 306718 629679 1278058 2541040 4924779",
	"6 11 37 93 189 349 636 1199 2368 4855 10167 21409 44779 92290 186698 370476 722452 1389717 2651819 5056495 9717067",
	"18 24 33 47 67 97 163 363 978 2701 7081 17332 39723 85844 176135 345170 649306 1177438 2065745 3517469 5828939",
	"18 35 65 129 273 595 1288 2705 5466 10653 20177 37451 68565 124233 222868 395239 691274 1189695 2011305 3336893 5430881",
	"8 16 22 39 105 300 779 1846 4105 8744 18060 36461 72466 142823 281093 555533 1106071 2219890 4483805 9086663 18409344",
	"-2 -6 -10 -14 -18 -22 -26 -30 -34 -38 -42 -46 -50 -54 -58 -62 -66 -70 -74 -78 -82",
	"17 32 61 125 265 555 1131 2256 4457 8803 17454 34711 68956 136139 265959 512819 975397 1831886 3403551 6268573 11463055",
	"4 20 45 80 138 264 573 1311 2944 6280 12629 24006 43382 74988 124677 200349 312444 474508 703837 1022204 1456674",
	"19 35 60 97 147 205 256 271 203 -17 -484 -1323 -2693 -4791 -7856 -12173 -18077 -25957 -36260 -49495 -66237",
	"15 26 56 115 223 439 909 1947 4181 8823 18169 36525 71936 139470 267591 510765 974681 1866797 3596898 6974250 13586527",
	"15 20 22 31 65 150 327 680 1422 3125 7280 17566 42555 101170 233169 518440 1111279 2299649 4605667 8953841 16950205",
	"6 10 13 20 47 146 449 1238 3054 6878 14462 28969 56209 106944 200990 374179 689670 1255626 2251915 3969258 6865147",
	"12 20 28 45 98 239 551 1158 2254 4193 7746 14769 29806 63702 141370 317903 711064 1565202 3374044 7111009 14650730",
	"7 27 74 167 335 632 1161 2107 3779 6661 11472 19235 31355 49706 76727 115527 169999 244943 346198 480783 657047",
	"7 14 30 80 206 479 1019 2029 3848 7028 12452 21545 36696 62115 105534 181588 318906 576315 1082463 2134847 4443259",
	"22 30 37 42 55 117 328 891 2183 4871 10104 19841 37454 68967 125892 232163 441415 879483 1845822 4046547 9108744",
	"11 29 62 129 272 570 1155 2238 4162 7508 13289 23276 40509 70055 120084 203343 339117 555775 894008 1410875 2184782",
	"7 20 48 97 183 345 664 1311 2667 5589 11946 25630 54378 112952 228561 449940 862323 1611794 2944352 5267719 9247757",
	"11 10 20 57 143 310 603 1080 1820 2985 5049 9416 19824 45236 105488 242048 538249 1156912 2409250 4878527 9636677",
}

func TestDay9Part1(t *testing.T) {
	expectedRes := 2756160
	if res := day9Part1(input9); res != expectedRes {
		t.Errorf("expected res is %v,what we get is %v", expectedRes, res)
	}
}

func TestDay9Part2(t *testing.T) {
	expectedRes := 34788142
	if res := day9Part2(input9); res != expectedRes {
		t.Errorf("expected res is %v,what we get is %v", expectedRes, res)
	}
}
