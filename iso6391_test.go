package iso6391

import "testing"

func TestFromCode(t *testing.T) {
	lang := FromCode("en")

	if lang.Name != "English" || lang.NativeName != "English" {
		t.Fatalf("Expected English, but was: %v", lang)
	}

	lang = FromCode("zh")

	if lang.Name != "Chinese" || lang.NativeName != "中文" {
		t.Fatalf("Expected Chinese, but was: %v", lang)
	}

	lang = FromCode("xx")

	if lang.Name != "" || lang.NativeName != "" {
		t.Fatalf("Expected no valid language, but was: %v", lang)
	}
}

func TestFromName(t *testing.T) {
	lang := FromName("English")

	if lang.Code != "en" || lang.NativeName != "English" {
		t.Fatalf("Expected English, but was: %v", lang)
	}

	lang = FromName("Chinese")

	if lang.Code != "zh" || lang.NativeName != "中文" {
		t.Fatalf("Expected Chinese, but was: %v", lang)
	}

	lang = FromName("Neverheardof")

	if lang.Code != "" || lang.NativeName != "" {
		t.Fatalf("Expected no valid language, but was: %v", lang)
	}
}

func TestFromNativeName(t *testing.T) {
	lang := FromNativeName("English")

	if lang.Code != "en" || lang.Name != "English" {
		t.Fatalf("Expected English, but was: %v", lang)
	}

	lang = FromNativeName("中文")

	if lang.Code != "zh" || lang.Name != "Chinese" {
		t.Fatalf("Expected Chinese, but was: %v", lang)
	}

	lang = FromNativeName("Neverheardof")

	if lang.Code != "" || lang.Name != "" {
		t.Fatalf("Expected no valid language, but was: %v", lang)
	}
}

func TestName(t *testing.T) {
	name := Name("en")

	if name != "English" {
		t.Fatalf("Expected English, but was: %v", name)
	}

	name = Name("xx")

	if name != "" {
		t.Fatalf("Expected no valid name, but was: %v", name)
	}
}

func TestNativeName(t *testing.T) {
	name := NativeName("zh")

	if name != "中文" {
		t.Fatalf("Expected Chinese, but was: %v", name)
	}

	name = NativeName("xx")

	if name != "" {
		t.Fatalf("Expected no valid native name, but was: %v", name)
	}
}

func TestCodeForName(t *testing.T) {
	code := CodeForName("English")

	if code != "en" {
		t.Fatalf("Expected English, but was: %v", code)
	}

	code = CodeForName("Chinese")

	if code != "zh" {
		t.Fatalf("Expected Chinese, but was: %v", code)
	}

	code = CodeForName("Neverheardof")

	if code != "" {
		t.Fatalf("Expected no valid code, but was: %v", code)
	}
}

func TestCodeForNativeName(t *testing.T) {
	code := CodeForNativeName("English")

	if code != "en" {
		t.Fatalf("Expected English, but was: %v", code)
	}

	code = CodeForNativeName("中文")

	if code != "zh" {
		t.Fatalf("Expected Chinese, but was: %v", code)
	}

	code = CodeForNativeName("Neverheardof")

	if code != "" {
		t.Fatalf("Expected no valid code, but was: %v", code)
	}
}

func TestValidCode(t *testing.T) {
	if !ValidCode("en") {
		t.Fatal("Code must be valid")
	}

	if !ValidCode("zh") {
		t.Fatal("Code must be valid")
	}

	if ValidCode("xx") {
		t.Fatal("Code must be invalid")
	}
}

func TestValidName(t *testing.T) {
	if !ValidName("English") {
		t.Fatal("Name must be valid")
	}

	if !ValidName("Chinese") {
		t.Fatal("Name must be valid")
	}

	if ValidName("Neverheardof") {
		t.Fatal("Name must be invalid")
	}
}

func TestValidNativeName(t *testing.T) {
	if !ValidNativeName("English") {
		t.Fatal("Native name must be valid")
	}

	if !ValidNativeName("中文") {
		t.Fatal("Native name must be valid")
	}

	if ValidNativeName("Neverheardof") {
		t.Fatal("Native name must be invalid")
	}
}
