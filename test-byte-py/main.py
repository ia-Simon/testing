hex_string = "021f30323030fefb4601a8e1f20a000000000004000431363533353739383531353032333336303930303030303030303030303030303031393830303030303030303031393830303030303030303031393830383237323233313435363130303030303036313030303030303930333633363139333132393038323730383237303832373533303030353130303030393939383632343531343130393030303030323435313337353335373938353135303233333630393d323731323230313030303030393030303030303030383237363737313338323049543331332020203030313338363137373038202020205041472a5370616e69202020202020202020202020202053414f205041554c4f2020202020425241303339523337333430313131303030303032333138313130333135303031333836313737303820202020393836393836393836107d05c1218d2c553133329f2608f14a050ea71704819f2701809f10120314a04003220000000000000000000000ff9f3704e921ed129f36020003950500000480009a032108279c01009f02060000000001985f2a020986820239009f1a0200769f34030203008407a00000000430609f03060000000000009f3501229f090200029f4104000598269f3303e0d0c830323630303030303030303030333031303736303438363530303520203031324d5332313333383930383839303037303830334f4c52303135303036503341303033484953464e58"

str_out = ""

for i in range(0, len(hex_string), 2):
  str_out += f"\\x{hex_string[i:i+2]}"

print(str_out)

b = bytes.fromhex(hex_string)

with open("./out", "wb") as f:
  f.write(b)