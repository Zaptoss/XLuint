package XLuint

import "testing"

func TestINV(t *testing.T) {
	var a1 XLuint

	a1.SetHex("6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b")
	if INV(a1).GetHex() != "94794d8c00cb031e62947fb100a5c0a8b8525b155dd0e2b63fe1ad224878a4b4" {
		t.Errorf("Error 1")
	}

	a1.SetHex("d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35")
	if INV(a1).GetHex() != "2b8ca1c5d9a1e9111fc0a68e7464a2fcfe63f827493ae06f25c5999113ec54ca" {
		t.Errorf("Error 2")
	}

	a1.SetHex("4e07408562bedb8b60ce05c1decfe3ad16b72230967de01f640b7e4729b49fce")
	if INV(a1).GetHex() != "b1f8bf7a9d4124749f31fa3e21301c52e948ddcf69821fe09bf481b8d64b6031" {
		t.Errorf("Error 3")
	}

	a1.SetHex("4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a")
	if INV(a1).GetHex() != "b4dd88882b22e039e39077b0b79be2fd4b2ede2c02cd734f74aace0353254075" {
		t.Errorf("Error 4")
	}

	a1.SetHex("ef2d127de37b942baad06145e54b0c619a1f22327b2ebbcfbec78f5564afe39d")
	if INV(a1).GetHex() != "10d2ed821c846bd4552f9eba1ab4f39e65e0ddcd84d14430413870aa9b501c62" {
		t.Errorf("Error 5")
	}
}

func TestXOR(t *testing.T) {
	var a1, a2 XLuint

	a1.SetHex("4fc82b26aecb47d2868c4efbe3581732a3e7cbcc6c2efb32062c08170a05eeb8")
	a2.SetHex("6b51d431df5d7f141cbececcf79edf3dd861c3b4069f0b11661a3eefacbba918")
	if XOR(a1, a2).GetHex() != "2499ff17719638c69a32803714c6c80f7b8608786ab1f023603636f8a6be47a0" {
		t.Errorf("Error 1")
	}

	a1.SetHex("6f4b6612125fb3a0daecd2799dfd6c9c299424fd920f9b308110a2c1fbd8f443")
	a2.SetHex("785f3ec7eb32f30b90cd0fcf3657d388b5ff4297f2f9716ff66e9b69c05ddd09")
	if XOR(a1, a2).GetHex() != "171458d5f96d40ab4a21ddb6abaabf149c6b666a60f6ea5f777e39a83b85294a" {
		t.Errorf("Error 2")
	}

	a1.SetHex("eb1e33e8a81b697b75855af6bfcdbcbf7cbbde9f94962ceaec1ed8af21f5a50f")
	a2.SetHex("e29c9c180c6279b0b02abd6a1801c7c04082cf486ec027aa13515e4f3884bb6b")
	if XOR(a1, a2).GetHex() != "982aff0a47910cbc5afe79ca7cc7b7f3c3911d7fa560b40ff4f86e019711e64" {
		t.Errorf("Error 3")
	}

	a1.SetHex("3d914f9348c9cc0ff8a79716700b9fcd4d2f3e711608004eb8f138bcba7f14d9")
	a2.SetHex("73475cb40a568e8da8a045ced110137e159f890ac4da883b6b17dc651b3a8049")
	if XOR(a1, a2).GetHex() != "4ed61327429f42825007d2d8a11b8cb358b0b77bd2d28875d3e6e4d9a1459490" {
		t.Errorf("Error 4")
	}

	a1.SetHex("031b4af5197ec30a926f48cf40e11a7dbc470048a21e4003b7a3c07c5dab1baa")
	a2.SetHex("41cfc0d1f2d127b04555b7246d84019b4d27710a3f3aff6e7764375b1e06e05d")
	if XOR(a1, a2).GetHex() != "42d48a24ebafe4bad73affeb2d651be6f16071429d24bf6dc0c7f72743adfbf7" {
		t.Errorf("Error 5")
	}
}

func TestOR(t *testing.T) {
	var a1, a2 XLuint

	a1.SetHex("4fc82b26aecb47d2868c4efbe3581732a3e7cbcc6c2efb32062c08170a05eeb8")
	a2.SetHex("6b51d431df5d7f141cbececcf79edf3dd861c3b4069f0b11661a3eefacbba918")
	if OR(a1, a2).GetHex() != "6fd9ff37ffdf7fd69ebecefff7dedf3ffbe7cbfc6ebffb33663e3effaebfefb8" {
		t.Errorf("Error 1")
	}

	a1.SetHex("6f4b6612125fb3a0daecd2799dfd6c9c299424fd920f9b308110a2c1fbd8f443")
	a2.SetHex("785f3ec7eb32f30b90cd0fcf3657d388b5ff4297f2f9716ff66e9b69c05ddd09")
	if OR(a1, a2).GetHex() != "7f5f7ed7fb7ff3abdaeddfffbfffff9cbdff66fff2fffb7ff77ebbe9fbddfd4b" {
		t.Errorf("Error 2")
	}

	a1.SetHex("eb1e33e8a81b697b75855af6bfcdbcbf7cbbde9f94962ceaec1ed8af21f5a50f")
	a2.SetHex("e29c9c180c6279b0b02abd6a1801c7c04082cf486ec027aa13515e4f3884bb6b")
	if OR(a1, a2).GetHex() != "eb9ebff8ac7b79fbf5affffebfcdffff7cbbdfdffed62feaff5fdeef39f5bf6f" {
		t.Errorf("Error 3")
	}

	a1.SetHex("3d914f9348c9cc0ff8a79716700b9fcd4d2f3e711608004eb8f138bcba7f14d9")
	a2.SetHex("73475cb40a568e8da8a045ced110137e159f890ac4da883b6b17dc651b3a8049")
	if OR(a1, a2).GetHex() != "7fd75fb74adfce8ff8a7d7def11b9fff5dbfbf7bd6da887ffbf7fcfdbb7f94d9" {
		t.Errorf("Error 4")
	}

	a1.SetHex("031b4af5197ec30a926f48cf40e11a7dbc470048a21e4003b7a3c07c5dab1baa")
	a2.SetHex("41cfc0d1f2d127b04555b7246d84019b4d27710a3f3aff6e7764375b1e06e05d")
	if OR(a1, a2).GetHex() != "43dfcaf5fbffe7bad77fffef6de51bfffd67714abf3eff6ff7e7f77f5faffbff" {
		t.Errorf("Error 5")
	}
}

func TestAND(t *testing.T) {
	var a1, a2 XLuint

	a1.SetHex("4fc82b26aecb47d2868c4efbe3581732a3e7cbcc6c2efb32062c08170a05eeb8")
	a2.SetHex("6b51d431df5d7f141cbececcf79edf3dd861c3b4069f0b11661a3eefacbba918")
	if AND(a1, a2).GetHex() != "4b4000208e494710048c4ec8e31817308061c384040e0b10060808070801a818" {
		t.Errorf("Error 1")
	}

	a1.SetHex("6f4b6612125fb3a0daecd2799dfd6c9c299424fd920f9b308110a2c1fbd8f443")
	a2.SetHex("785f3ec7eb32f30b90cd0fcf3657d388b5ff4297f2f9716ff66e9b69c05ddd09")
	if AND(a1, a2).GetHex() != "684b26020212b30090cc024914554088219400959209112080008241c058d401" {
		t.Errorf("Error 2")
	}

	a1.SetHex("eb1e33e8a81b697b75855af6bfcdbcbf7cbbde9f94962ceaec1ed8af21f5a50f")
	a2.SetHex("e29c9c180c6279b0b02abd6a1801c7c04082cf486ec027aa13515e4f3884bb6b")
	if AND(a1, a2).GetHex() != "e21c10080802693030001862180184804082ce08048024aa0010580f2084a10b" {
		t.Errorf("Error 3")
	}

	a1.SetHex("3d914f9348c9cc0ff8a79716700b9fcd4d2f3e711608004eb8f138bcba7f14d9")
	a2.SetHex("73475cb40a568e8da8a045ced110137e159f890ac4da883b6b17dc651b3a8049")
	if AND(a1, a2).GetHex() != "31014c9008408c0da8a005065000134c050f08000408000a281118241a3a0049" {
		t.Errorf("Error 4")
	}

	a1.SetHex("031b4af5197ec30a926f48cf40e11a7dbc470048a21e4003b7a3c07c5dab1baa")
	a2.SetHex("41cfc0d1f2d127b04555b7246d84019b4d27710a3f3aff6e7764375b1e06e05d")
	if AND(a1, a2).GetHex() != "10b40d11050030000450004408000190c070008221a4002372000581c020008" {
		t.Errorf("Error 5")
	}
}

func TestADD(t *testing.T) {
	var a1, a2 XLuint

	a1.SetHex("4fc82b26aecb47d2868c4efbe3581732a3e7cbcc6c2efb32062c08170a05eeb8")
	a2.SetHex("6b51d431df5d7f141cbececcf79edf3dd861c3b4069f0b11661a3eefacbba918")
	if ADD(a1, a2).GetHex() != "bb19ff588e28c6e6a34b1dc8daf6f6707c498f8072ce06436c464706b6c197d0" {
		t.Errorf("Error 1")
	}

	a1.SetHex("6f4b6612125fb3a0daecd2799dfd6c9c299424fd920f9b308110a2c1fbd8f443")
	a2.SetHex("785f3ec7eb32f30b90cd0fcf3657d388b5ff4297f2f9716ff66e9b69c05ddd09")
	if ADD(a1, a2).GetHex() != "e7aaa4d9fd92a6ac6bb9e248d4554024df93679585090ca0777f3e2bbc36d14c" {
		t.Errorf("Error 2")
	}

	a1.SetHex("eb1e33e8a81b697b75855af6bfcdbcbf7cbbde9f94962ceaec1ed8af21f5a50f")
	a2.SetHex("e29c9c180c6279b0b02abd6a1801c7c04082cf486ec027aa13515e4f3884bb6b")
	if ADD(a1, a2).GetHex() != "1cdbad000b47de32c25b01860d7cf847fbd3eade803565494ff7036fe5a7a607a" {
		t.Errorf("Error 3")
	}

	a1.SetHex("3d914f9348c9cc0ff8a79716700b9fcd4d2f3e711608004eb8f138bcba7f14d9")
	a2.SetHex("73475cb40a568e8da8a045ced110137e159f890ac4da883b6b17dc651b3a8049")
	if ADD(a1, a2).GetHex() != "b0d8ac4753205a9da147dce5411bb34b62cec77bdae2888a24091521d5b99522" {
		t.Errorf("Error 4")
	}

	a1.SetHex("031b4af5197ec30a926f48cf40e11a7dbc470048a21e4003b7a3c07c5dab1baa")
	a2.SetHex("41cfc0d1f2d127b04555b7246d84019b4d27710a3f3aff6e7764375b1e06e05d")
	if ADD(a1, a2).GetHex() != "44eb0bc70c4feabad7c4fff3ae651c19096e7152e1593f722f07f7d77bb1fc07" {
		t.Errorf("Error 5")
	}
}

func TestSUB(t *testing.T) {
	var a1, a2 XLuint

	a1.SetHex("4fc82b26aecb47d2868c4efbe3581732a3e7cbcc6c2efb32062c08170a05eeb8")
	a2.SetHex("6b51d431df5d7f141cbececcf79edf3dd861c3b4069f0b11661a3eefacbba918")
	if SUB(a1, a2).GetHex() != "e47656f4cf6dc8be69cd802eebb937f4cb860818658ff020a011c9275d4a45a0" {
		t.Errorf("Error 1")
	}

	a1.SetHex("6f4b6612125fb3a0daecd2799dfd6c9c299424fd920f9b308110a2c1fbd8f443")
	a2.SetHex("785f3ec7eb32f30b90cd0fcf3657d388b5ff4297f2f9716ff66e9b69c05ddd09")
	if SUB(a1, a2).GetHex() != "f6ec274a272cc0954a1fc2aa67a599137394e2659f1629c08aa207583b7b173a" {
		t.Errorf("Error 2")
	}

	a1.SetHex("eb1e33e8a81b697b75855af6bfcdbcbf7cbbde9f94962ceaec1ed8af21f5a50f")
	a2.SetHex("e29c9c180c6279b0b02abd6a1801c7c04082cf486ec027aa13515e4f3884bb6b")
	if SUB(a1, a2).GetHex() != "88197d09bb8efcac55a9d8ca7cbf4ff3c390f5725d60540d8cd7a5fe970e9a4" {
		t.Errorf("Error 3")
	}

	a1.SetHex("3d914f9348c9cc0ff8a79716700b9fcd4d2f3e711608004eb8f138bcba7f14d9")
	a2.SetHex("73475cb40a568e8da8a045ced110137e159f890ac4da883b6b17dc651b3a8049")
	if SUB(a1, a2).GetHex() != "ca49f2df3e733d82500751479efb8c4f378fb566512d78134dd95c579f449490" {
		t.Errorf("Error 4")
	}

	a1.SetHex("031b4af5197ec30a926f48cf40e11a7dbc470048a21e4003b7a3c07c5dab1baa")
	a2.SetHex("41cfc0d1f2d127b04555b7246d84019b4d27710a3f3aff6e7764375b1e06e05d")
	if SUB(a1, a2).GetHex() != "c14b8a2326ad9b5a4d1991aad35d18e26f1f8f3e62e34095403f89213fa43b4d" {
		t.Errorf("Error 5")
	}
}

func TestMUL(t *testing.T) {
	var a1, a2 XLuint

	a1.SetHex("4fc82b26aecb47d2868c4efbe358173")
	a2.SetHex("6b51d431df5d7f141cbececcf79edf3")
	if MUL(a1, a2).GetHex() != "21722a8020f51167758dcd016869bcf492554aee4b6ee4cfbd67a341fc5729" {
		t.Errorf("Error 1")
	}

	a1.SetHex("6f4b6612125fb3a0daecd2799dfd6c9")
	a2.SetHex("785f3ec7eb32f30b90cd0fcf3657d38")
	if MUL(a1, a2).GetHex() != "3454c01e85e6c34ca4af62126da34563022ecccea0e57b895f22d87d2420f8" {
		t.Errorf("Error 2")
	}

	a1.SetHex("eb1e33e8a81b697b75855af6bfcdbcb")
	a2.SetHex("e29c9c180c6279b0b02abd6a1801c7c")
	if MUL(a1, a2).GetHex() != "d0207f97840523d60fdec63572b7a999c897f8dda4c7b33f8f0a228704aa54" {
		t.Errorf("Error 3")
	}

	a1.SetHex("3d914f9348c9cc0ff8a79716700b9fc")
	a2.SetHex("73475cb40a568e8da8a045ced110137")
	if MUL(a1, a2).GetHex() != "1bb97057c0cd368fa0673c078a9513d9fb8406785b14985f1fd28ec79df124" {
		t.Errorf("Error 4")
	}

	a1.SetHex("031b4af5197ec30a926f48cf40e11a7")
	a2.SetHex("41cfc0d1f2d127b04555b7246d84019")
	if MUL(a1, a2).GetHex() != "cc7370dc1b28dd77b9bcab7a708d354568d0b071167c4a981cd7e8b1794f" {
		t.Errorf("Error 5")
	}
}
