package GeneralHashFunctions

import "testing"

func TestGeneralHashFunctions(t *testing.T) {
	key := "abcdefghijklmnopqrstuvwxyz1234567890"
	t.Log("1. RS-Hash Function:", RSHash(key))
	t.Log("2. JS-Hash Function Value:", JSHash(key))
	t.Log("3. PJW-Hash Function Value:", PJWHash(key))
	t.Log("4. ELF-Hash Function Value:", ELFHash(key))
	t.Log("5. BKDR-Hash Function Value:", BKDRHash(key))
	t.Log("6. SDBM-Hash Function Value:", SDBHash(key))
	t.Log("7. DJB-Hash Function Value:", DJBHash(key))
	t.Log("8. DEK-Hash Function Value:", DEKHash(key))
	t.Log("9. BP-Hash Function Value:", BPHash(key))
	t.Log("10. FNV-Hash Function Value:", FNVHash(key))
	t.Log("11. AP-Hash Function Value:", APHash(key))
}
