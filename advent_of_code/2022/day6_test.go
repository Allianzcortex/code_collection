package adventofcode

import "testing"

var puzzle6 = "plhlsssjsrscspsffmrffwvfvrvvmbbnjnrnrfnndlnlznlznlnccdbbhvbvgvmvzmzvmzmbbrcclsslzslzzsztzftzzhggnjgnjnhnmnqnqqfdfnnrwnwsnwwvgvqqgpptvtrvvfmmzjzmmjssmwsmmhzzvvwzwcwrcrllpbppdgdvvwqvvsnsrnsncnwntnggwqgqhqrqprqrddjvddqsqhshchfffzddswdsshhcnhnqqfjffvlvwvssdqqwrwvrvhrvrzzgwggcjcgclcwchhzvzmzrrjwjqwwvbwwrmmvmpmzpzgpgsghgrrtmmlfmmnzzmpzzvmvjjsqqshqshsqqgtqggpvvrtttwbbhnbnlbnlnhnthtjhjshhrmrjjlclpprmrnrsrwrbwrwwjnwnbwbhbmbggdbgghllcvlvddzbdbdvdwvdvjvqvcvczvcvclvvrggrngrrwcwbcczgghnhznhnttbcbvcvttrrpbrppndppvvvgvtgvvhfvhhttppjmjllznlnldnldlwlnwlnlttgzzcfcggwmgmqgmmshmhqqdfqfpfqppprzrhrnrhnrrtsrrgpgnpngnqnmmrtrvtrrfccszszffvlfvlvssvdvvpggvcvscvvpmpgpqpfqqhttrhhsbhbqhhzggzrgrqrfqffwllggrgqqjzqzzsgglvgllsgllptltblbggvrggctczzllvsvcscrrzjrzzjnnbvvtntpntnvtntcntcnnwsnnnvjnnsccsddcvvgzgwggbnbwnwbwmwttzzsgzzjpzztwztwwhhzggplgplggwwwphpmmhchsccmwmttvjtjftffzbfffjljbbqvvstsggbbqrbqrbrhbbrmbmppvrpphptpggqgddtmdtmdmbbbcdbbgssmzzthhtjjrgjrgrzzchzhttgddjnjrjmjllrcrqqsvsjjjhvvphhgjhgjhjccwmcmjcmjmhmzznmmcbcscddcsssnppsnpnmnqmmtztftqtjtqqgvgvjjfnjffqbqbfqqwfqfcfsfttvccnssjvvpfvpvttvpvccwhwfwlwlclljqqststllgqgzqqfpflpplpssnntstqqpnnsrnsrrsvrsvsrvsrrtztmtptpcpssgnnfvflffscsqszzsppfnftfptfpfnfbbwzbzfbbtggrzrnntztnznhnnlbnbrbdddjhjchhdshhnzhnnsrnrrqwrrdsdlsstdsslqsqdqvqzvqvwqqlmmwzmwzwztzlzczhqfrclvgvnlchrggsrjhntctdbpfdcffwjngdvdrjmgvwvptlvlhvhshqphmrdznqbtchcvrwfrpvwhzmrwlcjwnrsgrqcbsgpwjthstvqzlwjjmqbvhhdfdsmqnmnswmwjtpgjhdpgcnvmlcjwjzrjhmrqmqnrqrmgdnbdgznwhgzncmcbzpntcdflvrbfdzwgpnqjqmrcqpbrzwhwdgtgshhmrwvhnrwslvcswjvdgglfrdvmqdspppwmvfzvdbvpcnhmvgfqwnvjvvzrvttwvbjrbjlllmwtlcltvqmwshnqsdtjrptvqjvdjgzwgzzhcdbwjzhdgsptfrtmmqvhsnsnpgwbncbnnvwmmrrjgfccbzcpcjmqvqsbvjrstzblsrngphwndfdswjnnnfdgpcbslvbjglqqnbbtjljsmdgcslmwlvgwpsqthlmmqfgpgmcvrpvvtzcjdrwcjgrbwthblwpwpbzvjvhzsphmfhfwvsthlfnhhfcmpsnmgrvrntlzpdvqwtrghnslnfhcjwrsvrngqqtwvcsfhbjwmsnmsmgvdhnzjgljtchbtwlfppvbtdclbdjdmwzcntvgfjlcwdplnjwqnzqhnfgcnbrgftqpdmqzrrhglbzzvjcdnbtfnvmsrbjdzhbqmhnsmprgvjzzgvllhqnzgpstqzcnlgsrcwzlhwqvcjlwnnjslmdtwqcpbrntrmrmmtscjwwtzjhghtqvvpldzvhtmspzlnlfgrfhmsndbdgpgvphwwgqhrtlwztgqqsrwnrnqphqfsrtbztqbrgmfbtpjwhhrglhbzmmjptppnfdzlpbqfcbdwzdtqbrvfmtdzdjlnzvqfnzmttpqgzgmrqwmdtmdzrttffpdlgwdhnlrhnnztphvrbzcrlvcpswlngcjhdzqwwwpzdmhhwpnzwgjsdsdbdvpfsrwmwvsggpcqbchwjpgljnbtpvjzbbvgsbsbjmtwtbjtzzsfvrmfvnmcngvvdsvljvjbrlfgstjrhtjplttzhjfbmphvbqdsmwfwspqpcvmgzgjnqlphshlgdwcvtmpwcfgdcbqwpnshfgrfjvlrtqbffwwtbnwtvjlsdgwgfmhlrsmfrjzcwccdzfwwdwhwdsjbllhjsqmnqvngvdmbvbssfjtjfbtngrdwgldcrtpmvrbrmpvwwsfrlnbqsqzfnftpbhslqccnbqwgbdbfpblpmwbgjzmnptnhdjzqjhdrhqbtfhsrdqwlmwzqlsmmslgfzpvtgtsrlszvqrhrzclbqhzzwwfmhrrvsrnsjbdjsqqlsmgdmgbdtmvfjmsbwjqmqrvrhqbchpqrwvlvwpvhjlbdzfjvrwmchccrsrfvhzmpfjqwnrbvqmgwjcbndjdtfgnrzrqwgzhrdvghdrvgtplcrthgchmvwtdrchfwpdszzzhqpmrlzvfdnfnlwghmmwvsscbrdbchhgttsnbzdbqgddqfvcgvqwltnqtcwrmhtftnlwvlglsvvctccnntznsjnmmgqlzmplsdspnjtmzlbvrfzfclnhgjzmvntdwhspqwtpgndspnjqjqwrpwhjhpvjwfptpndnwvcjdwsvdtcrtwpsprgmrspgmcsdgtjbbgsgvcjhrcldlvgvqwqwthplzgwbzmszjfrlnznvgphnqzcwsztvvljlrlzrndbzccstprhntnlmshhclnrsmsvvvsbmpfdjsmspwcqtmlrrdmfzjjjhmsmdfqddzpgtbzbsnhhhpsfrdrdvpvpnjmvnhdrzrqggrpqdcmctvqfrtfmjjqjwgzzbrdfplhzjbnqlmjlcgvmsbpgdlfjmbgqtrvzzdgtlmbqthjrdtlstqtzqvfjvmmstsmtsbnjvstjjvrrjqcbjvhfpslpvjmdznrcnsvlpbpzmslqtpczmvhdwzrhwbwtfvrrmbszvrhwsjrclcscgngwvblbbrqprgshwzhlqgwmpfsmqsvpjbdccdmtnnhqfwvlgjlszmmmdmtmpzwhplzsjztrnwngbvspqqbmghwzgvfjdrblfmtwcvnczsrflmsjmrsvzldmttjwcmnvwcbjfvznhgntnqfbfchcqqshhjldgltqhdlqldlpfjnjvtbvbntzdjzstcqbzdnmsvcdgvjmmvtdfvdplqfndqqlzlspmjgdfbgsvwzqzsvvbbldltbtzwzpqrzfmfzgdbpqnwtrfcgwmlbpzrgscbjvfdwnjzjdfzltsbppnljzrgggplmttpmgwnhdlwfhwzsrcflnrqqzwsbqllwjqlrgwbhvcvqdvjvpbzgnfbbbtccvplzggplbrsbldllwmttwtvltsfljfbbprtvlfshhwdhdgvzfzjttvnphpnjnzsbvfwflfpqwnhvwjdsrtbwsjzqhgfldnssfbbzzqqrwtjvwjschmndgqzjtpsbfhwtmqbtfstrbghgtnldjqtshdnrwzvwddchhvdbsfjnqzfjdpvhvwwjspftgbtgwzdfgzzhpvjpdlrrdfnpftshthwzjzmzdghnfdqcbmjhfdgzrcgrzbrjtmhwbjhcpgjdsmnqzncdlwhqqzqblgdbbdsmrqgbdbmdczvvpnbbjdwrlmrwgnnbbzpcnsjgcmshgzwnjzwjlrdmvmhvjrzphgpczppqvwjthcdphprhhrggjdpzmgtpjjfvpczzrvsssfrrptnzlstrhmbhvzmwjnddshrrtgspbllvqlsptrtvtldsgnjjbwtfmtbjdvmgbptjlzhpttjmvpgnjphswhtdq"

func TestDay6Part1(t *testing.T) {
	expectedRes := 1
	if res := day6Part1(puzzle6); res != expectedRes {
		t.Errorf("expected res is %v,what we get is %v", expectedRes, res)
	}
}
